package lcd

import (
	"fmt"
	"github.com/davecheney/gpio"
	"github.com/davecheney/gpio/rpi"
	"time"
)

type LCDPanel struct {
	rsPin    gpio.Pin   // the pin for 'rs'
	enPin    gpio.Pin   // the pin for 'en'
	dataPins []gpio.Pin // array for data pins
}

func (l *LCDPanel) SetRsPin(pin gpio.Pin) {
	l.rsPin = pin
}

func (l *LCDPanel) SetEnablePin(pin gpio.Pin) {
	l.enPin = pin
}

func (l *LCDPanel) AddDataPin(pin gpio.Pin) {
	l.dataPins = append(l.dataPins, pin)
}

func (l *LCDPanel) CommandString(str string) {
	byteArray := []byte(str)
	for _, b := range byteArray {
		l.Message(b)
	}
}

func (l *LCDPanel) Command(val byte) {
	l.send(val, 0)
}

func (l *LCDPanel) Message(val byte) {
	l.send(val, 1)
}

/* low level bit writes */
func (l *LCDPanel) send(val byte, mode int) {

	//fmt.Printf("Begin :send val:%v, sending %d to rsPin %d\n", val, mode, l.rsPin)

	// set pin mode
	if mode > 0 {
		l.rsPin.Set()
	} else {
		l.rsPin.Clear()
	}

	// write the 4 bits sequentially
	l.Write4Bits(val >> 4)
	l.Write4Bits(val)

	//fmt.Printf("\nEnd :send for %v\n", val)
}

func (l *LCDPanel) Write4Bits(val byte) {

	DelayMicroseconds(1000)

	// pulse enable at end of function
	defer l.pulseEnable()

	// clear all data pins
	for _, pin := range l.dataPins {
		pin.Clear()
	}

	var mask uint8 = 1 << 3
	for i := 0; i < 4; i++ {
		if val&mask >= 1 {
			l.dataPins[i].Set()
		}
		mask >>= 1
	}
}

func (l *LCDPanel) pulseEnable() {
	l.enPin.Clear()
	DelayMicroseconds(10)
	l.enPin.Set()
	DelayMicroseconds(10)
	l.enPin.Clear()
	DelayMicroseconds(10)
}

func Delay(duration int) {
	time.Sleep(time.Duration(duration) * time.Millisecond)
}

func DelayMicroseconds(duration int) {
	time.Sleep(time.Duration(duration) * time.Microsecond)
}

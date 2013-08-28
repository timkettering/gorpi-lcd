package lcd

import (
	"github.com/davecheney/gpio"
	"time"
)

const (
	NEWLINE = 0xC0
)

type LCDPanel struct {
	rsPin    gpio.Pin   // the pin for 'rs'
	enPin    gpio.Pin   // the pin for 'en'
	dataPins []gpio.Pin // array for data pins
}

/* high level functions */
func (l *LCDPanel) Message(str string) {

	for _, c := range str {
		if c == '\n' {
			l.CommandByte(NEWLINE)
		} else {
			l.MessageByte(byte(c))
		}
	}
}

/* mid level functions */
func (l *LCDPanel) CommandByte(val byte) {
	l.send(val, 0)
}

func (l *LCDPanel) MessageByte(val byte) {
	l.send(val, 1)
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

/* low level bit writes */
func (l *LCDPanel) send(val byte, mode int) {

	// set pin mode
	if mode > 0 {
		l.rsPin.Set()
	} else {
		l.rsPin.Clear()
	}

	// write the 4 bits sequentially
	l.Write4Bits(val >> 4)
	l.Write4Bits(val)
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

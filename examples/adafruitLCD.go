package main

import (
	"github.com/davecheney/gpio"
	"github.com/davecheney/gpio"
	lcd "github.com/timkettering/gorpi-lcd"
)

func main() {

	lcdPanel := new(LCDPanel)

	// set up pins
	// enable pin
	gpio24, _ := rpi.OpenPin(rpi.GPIO24, gpio.ModeOutput)
	// rs pin
	gpio25, _ := rpi.OpenPin(rpi.GPIO25, gpio.ModeOutput)
	// data channel pins
	gpio23, _ := rpi.OpenPin(rpi.GPIO23, gpio.ModeOutput)
	gpio17, _ := rpi.OpenPin(rpi.GPIO17, gpio.ModeOutput)
	gpio27, _ := rpi.OpenPin(rpi.GPIO27, gpio.ModeOutput)
	gpio22, _ := rpi.OpenPin(rpi.GPIO22, gpio.ModeOutput)

	// setup lcdpanel struct
	lcdPanel.SetRsPin(gpio25)
	lcdPanel.SetEnablePin(gpio24)
	lcdPanel.AddDataPin(gpio22)
	lcdPanel.AddDataPin(gpio27)
	lcdPanel.AddDataPin(gpio17)
	lcdPanel.AddDataPin(gpio23)

	// start init
	DelayMicroseconds(50000)

	lcdPanel.Command(0x33) // $33 8-bit mode
	lcdPanel.Command(0x32) // $32 8-bit mode
	lcdPanel.Command(0x28) // $28 8-bit mode
	lcdPanel.Command(0x0C) // $0C 8-bit mode
	lcdPanel.Command(0x06) // $06 8-bit mode

	lcdPanel.Command(0x04 | 0x02 | 0x00)

	lcdPanel.Command(0x01)
	DelayMicroseconds(3000)

	// send text
	lcdPanel.CommandString("TEST")
}

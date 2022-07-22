// This example is intended to be used with the Adafruit Circuitplay Bluefruit board.
// It allows you to control the color of the built-in NeoPixel LEDS while they animate
// in a circular pattern.
//
package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/bluetooth"
	"tinygo.org/x/drivers/ws2812"
)

var adapter = bluetooth.DefaultAdapter

// TODO: use atomics to access this value.
var ledColor = [3]byte{0xff, 0x00, 0x00} // start out with red
var stringMsg = [10]string{}
var leds [10]color.RGBA

var (
	serviceUUID = [16]byte{0x28, 0x75, 0x99, 0x68, 0xf4, 0x50, 0x11, 0xec, 0xa9, 0xd6, 0xeb, 0x6f, 0x98, 0x8b, 0x5b, 0x69}
	charUUID    = [16]byte{0x28, 0x75, 0x99, 0x90, 0xf4, 0x50, 0x11, 0xec, 0xa9, 0xd7, 0x87, 0x73, 0xad, 0x8e, 0x89, 0x62}
)

var neo machine.Pin = machine.NEOPIXELS
var led machine.Pin = machine.LED
var ws ws2812.Device
var rg bool

var connected bool
var disconnected bool = true

func main() {
	println("starting")

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws = ws2812.New(neo)

	adapter.SetConnectHandler(func(d bluetooth.Addresser, c bool) {
		connected = c

		if !connected && !disconnected {
			clearLEDS()
			disconnected = true
		}

		if connected {
			disconnected = false
		}
	})

	must("enable BLE stack", adapter.Enable())
	adv := adapter.DefaultAdvertisement()
	must("config adv", adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: "GopherHelmet dev1.3",
	}))
	must("start adv", adv.Start())

	var ledColorCharacteristic bluetooth.Characteristic
	must("add service", adapter.AddService(&bluetooth.Service{
		UUID: bluetooth.NewUUID(serviceUUID),
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &ledColorCharacteristic,
				UUID:   bluetooth.NewUUID(charUUID),
				Value:  ledColor[:],
				Flags:  bluetooth.CharacteristicReadPermission | bluetooth.CharacteristicWritePermission,
				WriteEvent: func(client bluetooth.Connection, offset int, value []byte) {
					if offset != 0 || len(value) != 3 {
						return
					}
					println(value[0])
					println(value[1])
					println(value[2])
					ledColor[0] = value[0]
					ledColor[1] = value[1]
					ledColor[2] = value[2]
				},
			},
		},
	}))

	var stringCharacteristic bluetooth.Characteristic
	must("add service", adapter.AddService(&bluetooth.Service{
		UUID: bluetooth.NewUUID(serviceUUID),
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &stringCharacteristic,
				UUID:   bluetooth.NewUUID(charUUID),
				Value:  []byte(stringMsg[0]),
				Flags:  bluetooth.CharacteristicReadPermission | bluetooth.CharacteristicWritePermission,
				WriteEvent: func(client bluetooth.Connection, offset int, value []byte) {
					stringMsg[offset] = string(value)
					println(offset, string(value), stringMsg[0])
				},
			},
		},
	}))

	for {
		rg = !rg
		if connected {
			writeLEDS()
		}
		led.Set(rg)
		time.Sleep(100 * time.Millisecond)
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}

func writeLEDS() {
	for i := range leds {
		rg = !rg
		if rg {
			leds[i] = color.RGBA{R: ledColor[0], G: ledColor[1], B: ledColor[2]}
		} else {
			leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
		}
	}

	ws.WriteColors(leds[:])
}

func clearLEDS() {
	for i := range leds {
		leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
	}

	ws.WriteColors(leds[:])
}

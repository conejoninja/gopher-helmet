//go:build circuitplay_bluefruit
// +build circuitplay_bluefruit

package gopherhelmet

import (
	"machine"
	"tinygo.org/x/bluetooth"
	"tinygo.org/x/drivers/servo"
)

var (
	servo1 servo.PWM = machine.PWM1
	servo2 servo.PWM = machine.PWM2

	adapter      = bluetooth.DefaultAdapter
	serviceUUID  = [16]byte{0x28, 0x75, 0x99, 0x68, 0xf4, 0x50, 0x11, 0xec, 0xa9, 0xd6, 0xeb, 0x6f, 0x98, 0x8b, 0x5b, 0x69}
	charUUID     = [16]byte{0x28, 0x75, 0x99, 0x90, 0xf4, 0x50, 0x11, 0xec, 0xa9, 0xd7, 0x87, 0x73, 0xad, 0x8e, 0x89, 0x62}
	connected    = false
	disconnected = true

	alt       = true
	timeCount = 0
)

func InitBLE(device *BackpackDevice, fn BLECallback) {
	/*println("starting")

	adapter.SetConnectHandler(func(d bluetooth.Addresser, c bool) {
		connected = c

		if !connected && !disconnected {
			device.Clear()
			timeCount = 0
			disconnected = true
		}

		if connected {
			timeCount = 0
			disconnected = false
		}
	})

	must("enable BLE stack", adapter.Enable())
	adv := adapter.DefaultAdvertisement()
	must("config adv", adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: "GopherHelmet dev1.3",
	}))
	must("start adv", adv.Start())

	var stringCharacteristic bluetooth.Characteristic
	var stringMsg string
	must("add service", adapter.AddService(&bluetooth.Service{
		UUID: bluetooth.NewUUID(serviceUUID),
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &stringCharacteristic,
				UUID:   bluetooth.NewUUID(charUUID),
				Value:  []byte(stringMsg),
				Flags:  bluetooth.CharacteristicReadPermission | bluetooth.CharacteristicWritePermission,
				WriteEvent: func(client bluetooth.Connection, offset int, value []byte) {
					stringMsg = string(value)
					fn(stringMsg)
					println(offset, string(value), stringMsg)
				},
			},
		},
	}))

	for {
		alt = !alt

		if alt && timeCount < 30 {
			timeCount++
			if connected {
				device.Green()
			} else {
				device.Red()
			}
		} else {
			device.Clear()
		}

		time.Sleep(100 * time.Millisecond)
	} */
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}

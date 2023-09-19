package main

import (
	"github.com/conejoninja/gopher-helmet/fonts"
	"image/color"
	"math/rand"
	"strconv"
	"time"
	"tinygo.org/x/tinyfont"

	"tinygo.org/x/drivers/scd4x"

	"machine"

	gopherhelmet "github.com/conejoninja/gopher-helmet"
)

type Msg struct {
	text string
	c    color.RGBA
}

const (
	UseCO2Sensor = false

	Bootup uint8 = iota
	Demo
	Eyes
	CO2
	Axis
	Message
	Beer
	TinyGlobo
	Countdown

	Idle uint8 = iota
	Swipe

	RGB = iota
	Fading
	Beeping
)

var (
	accel    *gopherhelmet.AccelerometerDevice
	visor    *gopherhelmet.VisorDevice
	speaker  *gopherhelmet.SpeakerDevice
	antenna  *gopherhelmet.AntennaDevice
	backpack *gopherhelmet.BackpackDevice
	ears     *gopherhelmet.EarsDevice

	left  *gopherhelmet.ButtonDevice
	right *gopherhelmet.ButtonDevice

	co2sensor *scd4x.Device

	visorMode    uint8 = Demo
	earsMode     uint8 = Idle
	antennaMode  uint8 = RGB
	antennaColor color.RGBA

	msgs = [5]Msg{
		{"WWW.TINYGO.ORG", gopherhelmet.Red},
		{"ASK ME ABOUT GOPHER BADGES", gopherhelmet.Blue},
		{"@DEADPROGRAM TECHNOLOGIST FOR HIRE", gopherhelmet.Magenta},
		{"TALK - TINYGO: GETTING THE UPPER HEN BY DONIA CHAIEHLOUDJ", gopherhelmet.Yellow},
		{"FREE PINS AND STICKERS", gopherhelmet.Green},
	}
	visorStep = 0

	msgColored = gopherhelmet.TextColorSequence{
		{Text: "WWW", Color: gopherhelmet.Green},
		{Text: ".", Color: gopherhelmet.Yellow},
		{Text: "TINYGO", Color: gopherhelmet.Blue},
		{Text: ".", Color: gopherhelmet.Yellow},
		{Text: "ORG", Color: gopherhelmet.Magenta},
	}

	msgColoredRonTalk = gopherhelmet.TextColorSequence{
		{Text: "LOOKING", Color: gopherhelmet.Green},
		{Text: " FOR ", Color: gopherhelmet.Yellow},
		{Text: "FUNDING", Color: gopherhelmet.Blue},
	}

	msgColoredConejoTalk = gopherhelmet.TextColorSequence{
		{Text: "GOPHER", Color: gopherhelmet.Green},
		{Text: "CON", Color: gopherhelmet.Yellow},
		{Text: "EU ", Color: gopherhelmet.Blue},
		{Text: " SOOO ", Color: gopherhelmet.Red},
		{Text: " COOL", Color: gopherhelmet.Magenta},
	}

	msgLaunch = gopherhelmet.TextColorSequence{
		{Text: "LAUNCH !! ", Color: gopherhelmet.Green},
		{Text: "GO ", Color: gopherhelmet.Yellow},
		{Text: "TINY", Color: gopherhelmet.Blue},
		{Text: "GLOBO ", Color: gopherhelmet.Red},
		{Text: "GO", Color: gopherhelmet.Magenta},
	}

	msgColoredTinyGlobo = gopherhelmet.TextColorSequence{
		{Text: "TINY", Color: gopherhelmet.Green},
		{Text: "GLOBO ", Color: gopherhelmet.Yellow},
		{Text: "LAUNCH PARTY ", Color: gopherhelmet.Blue},
		{Text: "TODAY 12:30 ", Color: gopherhelmet.Red},
		{Text: "FRONT DOOR", Color: gopherhelmet.Magenta},
	}
)

func main() {

	//time.Sleep(2 * time.Second)
	accel = gopherhelmet.Accelerometer()
	visor = gopherhelmet.Visor()
	speaker = gopherhelmet.Speaker()

	antenna = gopherhelmet.Antenna()
	backpack = gopherhelmet.Backpack()
	ears = gopherhelmet.Ears()

	left = gopherhelmet.LeftButton()
	right = gopherhelmet.RightButton()

	if UseCO2Sensor {
		machine.I2C0.Configure(machine.I2CConfig{SCL: machine.SCL_PIN, SDA: machine.SDA_PIN})
		co2sensor = scd4x.New(machine.I2C0)
		co2sensor.Configure()

		if err := co2sensor.StartPeriodicMeasurement(); err != nil {
			println(err)
		}
	}

	/*go gopherhelmet.InitBLE(backpack, func(str string) {
		println("CALLBACK", str, str[:7])
		if str[:7] == "TINYGO" {
			if str[7:10] == "EAR" {
				a, _ := strconv.Atoi(str[11:])
				ears.Set(str[10], int16(a)-1)
			} else if str[7:10] == "MOD" {
				s := str[11:]
				if s == "demo" {
					visorMode = Demo
				} else if s == "eyes" {
					visorMode = Eyes
				} else if s == "co2" {
					visorMode = CO2
				} else if s == "axis" {
					visorMode = Axis
				} else if s == "message" {
					visorMode = Message
				}
			} else if str[7:10] == "MSG" {
				b, err := hex.DecodeString(str[11:17] + "FF")
				if err != nil {
					log.Fatal(err)
				}

				msgs[str[10]] = Msg{
					c:    color.RGBA{b[0], b[1], b[2], b[3]},
					text: str[17:],
				}
				visorStep = int(str[10])
				visorMode = Message
			} else if str[7:10] == "EYE" {

			}
		}
	})  */
	//CountdownF()

	ears.Set(2, 90)
	visor.BootUp()
	go earsLoop()
	go antennaLoop()
	go buttonsLoop()
	visorLoop()
}

func visorLoop() {
	var step uint8 = 0
	for {
		switch visorMode {
		case Demo:
			switch step {
			case 0:
				visor.MarqueeColored(msgColored)
				//visor.Marquee(msgs[0].text, msgs[0].c)
				break
			case 1:
				co2Marquee()
				break
			case 2:
				visor.MarqueeColored(msgColoredConejoTalk)
				//visor.Marquee(msgs[1].text, msgs[1].c)
				break
			case 3:
				lookingSides(gopherhelmet.White)
				break
			case 4:
				visor.Marquee(msgs[2].text, msgs[2].c)
				break
			case 5:
				lookingSuspicious(gopherhelmet.Red)
				break
			case 6:
				visor.MarqueeColored(msgColoredRonTalk)
				//visor.Marquee(msgs[3].text, msgs[3].c)
				break
			case 7:
				demoAxis()
				break
			case 8:
				visor.Marquee(msgs[4].text, msgs[4].c)
				break
			case 9:
				visor.MarqueeColored(msgColoredTinyGlobo)
				break
			}
			step++

			if step > 9 {
				step = 0
			}
			break
		case Eyes:
			visorMode = Demo
			break
		case CO2:
			co2Marquee()
			visorMode = Demo
			break
		case Axis:
			demoAxis()
			visorMode = Demo
			break
		case Beer:
			visor.MarqueeColored(msgLaunch)
			visorMode = Demo
			break
		case TinyGlobo:
			visor.MarqueeColored(msgColoredTinyGlobo)
			visorMode = Demo
			break
		case Message:
			visor.Marquee(msgs[visorStep].text, msgs[visorStep].c)
			visorMode = Demo
			break
		case Countdown:
			CountdownF()
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func earsLoop() {
	angle := int16(90)
	forward := true
	for {
		switch earsMode {
		case Idle:
			ears.Off()
			if rand.Int31n(1000) == 1 {
				angle = 90
				forward = true
				earsMode = Swipe
			}
			break
		case Swipe:
			if forward {
				println("FORWARD", angle)
				angle -= 10
				if angle <= 0 {
					angle = 0
					forward = false
				}
				ears.Set(2, angle)
			} else {
				angle += 10
				println("BACK", angle)
				if angle >= 90 {
					angle = 90
					forward = true
					earsMode = Idle
				}
				ears.Set(2, angle)
			}
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func antennaLoop() {
	var i uint8
	var f uint8 = 2
	for {
		switch antennaMode {
		case RGB:
			antenna.Rainbow(i)
			i++
			break
		case Fading:
			i += f
			if i >= 255 {
				i = 255
				f = -f
			}
			if i <= 0 {
				i = 0
				f = -f
			}
			antenna.SetColor(color.RGBA{i, 0, 0, 255})
			break
		case Beeping:
			i++
			if i >= 20 {
				i = 0
			}
			if i < 10 {
				antenna.SetColor(antennaColor)
			} else {
				antenna.SetColor(gopherhelmet.Black)
			}
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func demoAxis() {
	val := make([]int32, 3)
	c := gopherhelmet.Red
	for i := 0; i < 400; i++ {
		val[0], val[1], val[2], _ = accel.ReadAcceleration()
		visor.Clear()
		visor.SetPixel(8, 0, gopherhelmet.Red)
		visor.SetPixel(8, 2, gopherhelmet.Green)
		visor.SetPixel(8, 4, gopherhelmet.Blue)
		for k := int16(0); k < 3; k++ {
			val[k] = val[k] / 125000
			if k == 0 {
				c = gopherhelmet.Red
			} else if k == 1 {
				c = gopherhelmet.Green
			} else {
				c = gopherhelmet.Blue
			}
			if val[k] < 0 {
				for i := int16(8 + val[k]); i < 8; i++ {
					visor.SetPixel(i, k*2, c)
				}
			} else {
				for i := int16(8 + val[k]); i > 8; i-- {
					visor.SetPixel(i, k*2, c)
				}
			}
		}

		visor.Display()
		time.Sleep(50 * time.Millisecond)
	}
}

func buttonsLoop() {
	for {
		if left.Pushed() {
			visorMode = Countdown
		}
		if right.Pushed() {
			visorMode = TinyGlobo
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func co2Marquee() {
	if !UseCO2Sensor {
		return
	}
	var co2 int32
	var err error
	for i := 0; i < 5; i++ {
		co2, err = co2sensor.ReadCO2()
		value, _ := co2sensor.ReadTemperature()
		println("TEMP", value)
		value, _ = co2sensor.ReadHumidity()
		println("HUM", value)
		if err != nil {
			println(err)
		}
		println(co2)
		if co2 != 0 {
			break
		} else {
			time.Sleep(200 * time.Millisecond)
		}
	}
	switch {
	case co2 < 800:
		antennaColor = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
	case co2 < 1500:
		antennaColor = color.RGBA{R: 0xff, G: 0xff, B: 0x00}
	default:
		antennaColor = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
		go beepBeep()
	}
	antennaMode = Beeping
	visor.Marquee("CO2 LEVELS: "+strconv.Itoa(int(co2))+" PPM", antennaColor)
	antennaMode = RGB
}

func beepBeep() {
	speaker.Bloop()
	time.Sleep(50 * time.Millisecond)
	speaker.Blip()
	time.Sleep(50 * time.Millisecond)
	speaker.Bleep()
	time.Sleep(50 * time.Millisecond)
	speaker.Bloop()
	time.Sleep(50 * time.Millisecond)
	speaker.Blip()
	time.Sleep(50 * time.Millisecond)
	speaker.Bleep()
}

func CountdownF() {
	t := time.Duration(932)
	visor.Clear()
	tinyfont.WriteLine(visor.Device, &fonts.TomThumb, 5, 5, "10", gopherhelmet.White)
	visor.Display()
	speaker.Blip()
	time.Sleep(t * time.Millisecond)

	visor.Clear()
	tinyfont.WriteLine(visor.Device, &fonts.TomThumb, 6, 5, "9", gopherhelmet.White)
	visor.Display()
	speaker.Blip()
	time.Sleep(t * time.Millisecond)

	visor.Clear()
	tinyfont.WriteLine(visor.Device, &fonts.TomThumb, 6, 5, "8", gopherhelmet.White)
	visor.Display()
	speaker.Blip()
	time.Sleep(t * time.Millisecond)

	visor.Clear()
	tinyfont.WriteLine(visor.Device, &fonts.TomThumb, 6, 5, "7", gopherhelmet.White)
	visor.Display()
	speaker.Blip()
	time.Sleep(t * time.Millisecond)

	visor.Clear()
	tinyfont.WriteLine(visor.Device, &fonts.TomThumb, 6, 5, "6", gopherhelmet.White)
	visor.Display()
	speaker.Blip()
	time.Sleep(t * time.Millisecond)

	visor.Clear()
	tinyfont.WriteLine(visor.Device, &fonts.TomThumb, 6, 5, "5", gopherhelmet.White)
	visor.Display()
	speaker.Blip()
	time.Sleep(t * time.Millisecond)

	visor.Clear()
	tinyfont.WriteLine(visor.Device, &fonts.TomThumb, 6, 5, "4", gopherhelmet.White)
	visor.Display()
	speaker.Blip()
	time.Sleep(t * time.Millisecond)

	visor.Clear()
	tinyfont.WriteLine(visor.Device, &fonts.TomThumb, 6, 5, "3", gopherhelmet.White)
	visor.Display()
	speaker.Blip()
	time.Sleep(t * time.Millisecond)

	visor.Clear()
	tinyfont.WriteLine(visor.Device, &fonts.TomThumb, 6, 5, "2", gopherhelmet.White)
	visor.Display()
	speaker.Blip()
	time.Sleep(t * time.Millisecond)

	visor.Clear()
	tinyfont.WriteLine(visor.Device, &fonts.TomThumb, 6, 5, "1", gopherhelmet.White)
	visor.Display()
	speaker.Blip()
	time.Sleep(t * time.Millisecond)

	visor.Clear()
	tinyfont.WriteLine(visor.Device, &fonts.TomThumb, 6, 5, "0", gopherhelmet.White)
	visor.Display()
	speaker.Blip()
	time.Sleep(30 * time.Millisecond)
	speaker.Blip()
	time.Sleep(30 * time.Millisecond)
	speaker.Blip()
	time.Sleep(30 * time.Millisecond)

	visor.MarqueeColored(msgLaunch)
	visorMode = Demo
}

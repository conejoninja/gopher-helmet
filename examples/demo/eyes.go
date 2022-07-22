package main

import (
	"image/color"
	"time"
)

var eyeFrame uint8
var eyeWait uint8
var eyeWaitTime uint8

func lookingSides(c color.RGBA) {
	eyeFrame = 0
	eyeWait = 0
	for {
		visor.Clear()
		switch eyeFrame {
		case 0:
			eye(0, 0, c)
			eye(12, 0, c)
			eyeWaitTime = 40
			break
		case 1:
			eye(0, -1, c)
			eye(12, -1, c)
			eyeWaitTime = 12
			break
		case 2:
			eye(0, 1, c)
			eye(12, 1, c)
			eyeWaitTime = 12
			break
		case 3:
			eye(0, -1, c)
			eye(12, -1, c)
			eyeWaitTime = 12
			break
		case 4:
			eye(0, 1, c)
			eye(12, 1, c)
			eyeWaitTime = 12
			break
		case 5:
			eye(0, 0, c)
			eye(12, 0, c)
			eyeWaitTime = 40
			break
		case 6:
			eye_close(0, 0, c)
			eye_close(12, 0, c)
			eyeWaitTime = 2
			break
		case 7:
			eye_fully_close(0, 0, c)
			eye_fully_close(12, 0, c)
			eyeWaitTime = 2
			break
		case 8:
			eye_close(0, 0, c)
			eye_close(12, 0, c)
			eyeWaitTime = 12
			break
		case 9:
			return
		}
		eyeWait++
		if eyeWait > eyeWaitTime {
			eyeWait = 0
			eyeFrame++
		}
		visor.Display()
		time.Sleep(50 * time.Millisecond)
	}
}

func eye(x, z int16, c color.RGBA) {
	visor.SetPixel(x, 1, c)
	visor.SetPixel(x, 2, c)
	visor.SetPixel(x, 3, c)

	visor.SetPixel(x+1, 0, c)
	visor.SetPixel(x+2, 0, c)
	visor.SetPixel(x+3, 0, c)

	visor.SetPixel(x+4, 1, c)
	visor.SetPixel(x+4, 2, c)
	visor.SetPixel(x+4, 3, c)

	visor.SetPixel(x+1, 4, c)
	visor.SetPixel(x+2, 4, c)
	visor.SetPixel(x+3, 4, c)

	visor.SetPixel(x+2+z, 2, c)
}

func eye_close(x, z int16, c color.RGBA) {
	visor.SetPixel(x, 2, c)

	visor.SetPixel(x+1, 1, c)
	visor.SetPixel(x+2, 1, c)
	visor.SetPixel(x+3, 1, c)

	visor.SetPixel(x+4, 2, c)

	visor.SetPixel(x+1, 3, c)
	visor.SetPixel(x+2, 3, c)
	visor.SetPixel(x+3, 3, c)
}

func eye_fully_close(x, z int16, c color.RGBA) {
	visor.SetPixel(x, 2, c)
	visor.SetPixel(x+1, 2, c)
	visor.SetPixel(x+2, 2, c)
	visor.SetPixel(x+3, 2, c)
	visor.SetPixel(x+4, 2, c)
}

package main

import (
	"bytes"
	"image"
	"image/png"
	"os"
)

func CreateJuliaSet(myImg *image.RGBA, cReal, cImaginary, zoom, movex, movey float64, index int, name string, byteChan ...chan ImageInput) {
	var newReal, newImaginary, oldReal, oldImaginary float64
	maxIterations := 255 // 255 for RGB
	height := float64(myImg.Bounds().Size().Y)
	width := float64(myImg.Bounds().Size().X)

	for i := 0; i < myImg.Bounds().Size().X; i++ {
		for j := 0; j < myImg.Bounds().Size().Y; j++ {

			newReal = 1.5*(float64(j)-width/2)/(0.5*zoom*width) + movex
			newImaginary = (float64(i)-height/2)/(0.5*zoom*height) + movey

			iterations := 0
			for ; iterations < maxIterations; iterations++ {
				oldReal = newReal
				oldImaginary = newImaginary
				// z is a complex number described by newReal and newImaginary
				// a+bi where a = newReal, b = newImaginary
				newReal = oldReal*oldReal - oldImaginary*oldImaginary + cReal
				newImaginary = 2*oldReal*oldImaginary + cImaginary
				if (newReal*newReal + newImaginary*newImaginary) > 2 {
					break
				}
			}

			myImg.Set(i, j, colorPalette[mapToRange(float64(iterations))])
		}
	}

	if len(byteChan) == 1 && byteChan[0] != nil {
		// auto ffmpeg
		b := new(bytes.Buffer)
		png.Encode(b, myImg)
		byteChan[0] <- ImageInput{
			index: index,
			bytes: b.Bytes(),
		}
	} else {
		out, err := os.Create(name)
		if err != nil {
			panic(err)
		}
		png.Encode(out, myImg)
		out.Close()
	}
}

package main

func init() {
	// setup color gradient, add additional hex values as desired.
	// transitions between each color value are done in 2048 steps, which makes the length of the color
	// palette equal to 2048 * (numberOfColors)
	colorPalette = generateColorPalette("#000764", "#206acb", "#edffff", "#ffaa00", "#0002000")
}

// windows ffmpeg command to create mp4 from pngs
// ffmpeg -framerate 30 -i img%04d.png -c:v libx264 -pix_fmt yuv420p out.mp4
func main() {

	workerCount := 4
	endRange := .005
	increment := 0.00001
	cReal := 0.280
	cImaginary := 0.01

	imageWidth := 1600
	imageHeight := 1600
	zoom := 1.
	moveX := 0.
	moveY := 0.

	constructor := WorkerPoolConstructor{
		WorkerCount: workerCount,
		endRange:    endRange,
		increment:   increment,
		InitialCondition: InitialCondition{
			cReal:      cReal,
			cImaginary: cImaginary,
		},
		ImageProperties: ImageProperties{
			imageHeight: imageHeight,
			imageWidth:  imageWidth,
		},
		CameraModifiers: CameraModifiers{
			zoom:  zoom,
			moveX: moveX,
			moveY: moveY,
		},
	}

	StartWorkers(constructor.CreateWorkerPool())
}

package main

import (
	"context"
	"flag"
)

func run() {
	constructor, isWindows := flags()
	if !isWindows {
		ctx, cancel := context.WithCancel(context.Background())
		ffmpegProcessor := ImageProcessor{
			Input:  make(chan ImageInput),
			endCtx: cancel,
		}
		StartWorkers(constructor.CreateWorkerPool(&ffmpegProcessor))
		<-ctx.Done()
	} else {
		StartWorkers(constructor.CreateWorkerPool())
	}
}

func flags() (WorkerPoolConstructor, bool) {
	isWindows := flag.Bool("windows", false, "underlying OS is Windows (default: false)")
	width := flag.Int("image-width", 1600, "width of resulting video or images")
	height := flag.Int("image-height", 1600, "height of resulting video or images")
	workerCount := flag.Int("worker-count", 2, "number of threads to use")
	cReal := flag.Float64("constant-real", 0.280, "constant real value to be used in processing of julia set")
	cImaginary := flag.Float64("constant-imaginary", 0.01, "constant imaginary value to be used in processing of julia set")
	endRange := flag.Float64("end-range", 0.005, "worker loop range")
	increment := flag.Float64("increment", 0.0001, "loop increment (smaller is slower, larger is faster)")
	zoom := flag.Float64("zoom", 1., "zoom value")
	flag.Parse()
	return WorkerPoolConstructor{
		WorkerCount: *workerCount,
		endRange:    *endRange,
		increment:   *increment,
		InitialCondition: InitialCondition{
			cReal:      *cReal,
			cImaginary: *cImaginary,
		},
		ImageProperties: ImageProperties{
			imageHeight: *height,
			imageWidth:  *width,
		},
		CameraModifiers: CameraModifiers{
			zoom: *zoom,
		},
	}, *isWindows
}

package main

import (
	"fmt"
	"image"
	"math"
	"sync"
	"time"
)

// The Ranges struct specifies the loop range that a particular worker
// will operate on. For example, if we wish to compute a julia set function and increase its cReal
// until it is 0.25 greater than its initial value (so as to animate the set),
// and we increment the cReal each time by some constant value such as 0.1 (to control the speed of the animation),
// and we use 5 workers, the first Ranges struct would look like,
// {
//  beginRange: 0.
//  endRange: 0.5
//  increment: 0.1
// }
// and the worker would in turn produce 5 images.
type Ranges struct {
	beginRange float64
	endRange   float64
	increment  float64
}

// CameraModifiers controls camera panning across the X and Y axis,
// as well as camera zoom.
type CameraModifiers struct {
	zoom  float64
	moveX float64
	moveY float64
}

// ImageProperties defines the size of the julia set image,
// which directly relates to the size of the
// matrices used within the computation.
type ImageProperties struct {
	imageHeight int
	imageWidth  int
}

// InitialCondition specifies the cReal and cImaginary values
// that should be used.
type InitialCondition struct {
	cReal      float64
	cImaginary float64
}

// Worker represents a julia set worker, capable of producing images
type Worker struct {
	ID int
	InitialCondition
	Ranges
	ImageProperties
	CameraModifiers
	imageNumberIndex int
	index            int // use zero value
	amount           int
}

// WorkerPoolConstructor is a utility type for creating worker pools which
// have been properly initialized.
type WorkerPoolConstructor struct {
	WorkerCount int
	endRange    float64
	increment   float64
	InitialCondition
	ImageProperties
	CameraModifiers
}

func (w WorkerPoolConstructor) CreateWorkerPool() []Worker {
	var workers []Worker
	totalItems := math.Ceil(w.endRange / w.increment)
	workItemsPerWorker := int(totalItems) / w.WorkerCount
	for i := 0; i < w.WorkerCount; i++ {
		workerBeginOffset := float64(i) * (float64(workItemsPerWorker))
		workerEndOffset := float64(i+1) * (float64(workItemsPerWorker))
		workers = append(workers, Worker{
			ID:               i,
			InitialCondition: w.InitialCondition,
			Ranges: Ranges{
				beginRange: workerBeginOffset * w.increment,
				endRange:   workerEndOffset * w.increment,
				increment:  w.increment,
			},
			ImageProperties:  w.ImageProperties,
			CameraModifiers:  w.CameraModifiers,
			imageNumberIndex: int(workerBeginOffset),
			index:            0,
			amount:           workItemsPerWorker,
		})
	}
	return workers
}

func StartWorkers(workers []Worker) {
	start := time.Now()
	fmt.Println("Starting Workers")
	wg := &sync.WaitGroup{}
	for _, w := range workers {
		go func(w Worker, group *sync.WaitGroup) {
			wg.Add(1)
			w.CreateJuliaSetImage()
			wg.Done()
		}(w, wg)
	}
	time.Sleep(500 * time.Millisecond)
	wg.Wait()
	fmt.Println("Done. Total runtime = ", time.Since(start))
}

func (w *Worker) CreateJuliaSetImage() {
	n := time.Now()
	k := 0
	for i := w.beginRange; i < w.endRange; i = i + w.increment {
		fmt.Println("Worker ", w.ID, ": ", w.index, "/", w.amount, " start")
		CreateJuliaSet(
			image.NewRGBA(image.Rect(0, 0, w.imageWidth, w.imageHeight)),
			w.cReal+i, // real constant
			w.cImaginary,
			w.zoom,
			w.moveX,
			w.moveY,
			GetImageFilePath(w.imageNumberIndex+k),
		)
		w.index++
		k++
		fmt.Println("Worker ", w.ID, ": ", w.index, "/", w.amount, " complete")
	}
	fmt.Println("ID: ", w.ID, " Total Runtime: ", time.Since(n), " DONE")
}

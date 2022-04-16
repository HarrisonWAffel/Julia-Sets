package main

import (
	"bytes"
	"context"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"sort"
)

type ImageProcessor struct {
	ExpectedReceives int // how many byte arrays will be sent
	Input            chan ImageInput
	endCtx           context.CancelFunc
}

type ImageInput struct {
	index int
	bytes []byte
}

// processVideo listens for incoming
// image files and converts them into
// an mp4 file using FFMPEG.
func (b *ImageProcessor) processVideo() {
	images := make([]ImageInput, b.ExpectedReceives)
	received := 0
	for {
		if received == b.ExpectedReceives {
			fmt.Println("got ", b.ExpectedReceives, " byte arrays")
			break
		}
		select {
		case input := <-b.Input:
			images = append(images, input)
			received++
		}
	}

	sort.Slice(images, func(i, j int) bool {
		return images[i].index > images[j].index
	})

	i := bytes.Buffer{}
	for _, img := range images {
		i.Write(img.bytes)
	}

	if ffmpeg.Input("pipe:").
		Output("juliaSet.mp4", ffmpeg.KwArgs{
			"c:v":     "libx264",
			"pix_fmt": "yuv420p",
		}).
		OverWriteOutput().
		ErrorToStdOut().
		WithInput(bytes.NewReader(i.Bytes())).
		Run() != nil {
		panic("error running FFMPEG")
	}

	fmt.Println("video generated!")
	os.Exit(0)
}

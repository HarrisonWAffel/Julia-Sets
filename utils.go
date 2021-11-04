package main

import (
	"fmt"
	"github.com/muesli/gamut"
	"image/color"
	"math"
	"path/filepath"
	"runtime"
	"strings"
)

// maps one range to another, e.g. maps a float64 range of [0 ,1] to any other range such as [0,360]
func mapToRange(input float64) int {
	inputStart := 0.
	inputEnd := 255.
	outputStart := 0.
	outputEnd := 2048.

	slope := 1.0 * (outputEnd - outputStart) / (inputEnd - inputStart)
	return int(outputStart + math.Round(slope*(input-inputStart)))
}

var colorPalette []color.Color

func generateColorPalette(hexColors ...string) []color.Color {
	var cp []color.Color
	for i := 0; i < len(hexColors)-1; i++ {
		cp = append(cp, gamut.Blends(gamut.Hex(hexColors[i]), gamut.Hex(hexColors[i+1]), 2048/(len(hexColors)-1)+1)...)
	}
	return cp
}

func GetImageFilePath(ID int) string {
	_, filename, _, _ := runtime.Caller(1)
	curFilename := strings.Split(filename, "/")

	windowsPath := strings.ReplaceAll(filename, curFilename[len(curFilename)-1], fmt.Sprintf("movie/img%04d.png", ID))
	x := fmt.Sprintf(filepath.FromSlash(windowsPath))
	return x
}

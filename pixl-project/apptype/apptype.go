package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType = int

type PxCanvasconfig struct {
	DrawingArea fyne.Size
	CanvaOffset fyne.Position
	PxRows, Pxcols int
	PxSize int
}

type State struct {
	BrushColor color.Color
	BrushType int
	SwatchSelected int
	FilePath string
}

func(state *State) SetFilePath(path string)
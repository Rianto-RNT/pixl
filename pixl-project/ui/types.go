package ui

import (
	"rnt/pixl/apptype"
	"rnt/pixl/pxcanvas"
	"rnt/pixl/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
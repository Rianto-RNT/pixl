package pxcanvas

import (
	"image"
	"image/color"
	"rnt/pixl/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type pxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer *PxCanvasRenderer
	pixelData image.Image
	MouseState PxCanvasMouseState
	appState *apptype.State
	reloadImage bool
}

func (pxCanvas *PxCanvas) Bounds() image.Rectangle {
	x0 := int(pxCanvas.CanvasOffset.x)
	y0 := int(pxCanvas.CanvasOffset.y)
	x1 := int(pxCanvas.PxCols * pxCanvas.PxSize + int(pxCanvas.CanvasOffset.x))
	y1 := int(pxCanvas.PxRows * pxCanvas.PxSize + int(pxCanvas.CanvasOffset.y))
	return image.Rect(x0,y0,x1,y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) &&
	pos.X < float32(bounds.Max.X) &&
	pos.Y >= float32(bounds.Min.Y)&&
	pos.Y < float32(bounds.Min.Y) {
		return true
	}
	return false
}

func NewBlankImage(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0,0,cols, rows))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x,y,c)
		}
	}
	return img
}

func NewPxCanvas( state *apptype.State, config apptype.PxCanvasConfig) *PxCanvas {
	pxCanvas := &PxCanvas {
		PxCanvasConfig : config,
		appState: state,
	}
	pxCanvas.PixelData = NewBlankImage(pxCanvas.PxCols, pxCanvas.PxRows, color.NRGBA{128,128,128,255})
	pxCanvas.ExtendBaseWidget(pxCanvas)
	return pxCanvas
} 

func (pxCanvas *PxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100,100,100,255}
		canvasBorder[i].StrokeWidth = 2
	}
	renderer := &PxCanvasRenderer{
		pxCanvas: pxCanvas,
		canvasImage: canvasImage,
		canvasBorder: canvasBorder,
	}

	pxCanvas.renderer = renderer
	return renderer
}




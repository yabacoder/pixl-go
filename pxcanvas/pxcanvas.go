package pxcanvas

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"yabacoder.com/pixl/apptype"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer *PxCanvasRenderer
	PixelData image.Image
	mouseState PxCanvasMouseState
	appState *apptype.State
	reloadImage bool
}

func (pxCanvas *PxCanvas) Bounds() image.Rectangle {
	x0 := int(pxCanvas.CanvasOffset.X) // The beginning of the canvas drawing X
	y0 := int(pxCanvas.CanvasOffset.Y) // the Y
	// Placement of the Pixel canvas within the drawable canvas 
	x1 := int(pxCanvas.PxCols *pxCanvas.PxSize + int(pxCanvas.CanvasOffset.X))
	
	y1 := int(pxCanvas.PxRows *pxCanvas.PxSize + int(pxCanvas.CanvasOffset.Y))
	return image.Rect(x0, y0, x1, y1)
}

/**
This function tracks the mouse position
*/
  func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	// Tracks the position of the mouse, ensure it stays within the box boundary
	if pos.X >= float32(bounds.Min.X) && 
		pos.Y < float32(bounds.Max.X) &&
		pos.Y >= float32(bounds.Min.X) &&
		pos.Y < float32(bounds.Max.Y) {
			return true
		}
		return false
  }

// Let's create an image to be displayed when the app loads
func NewBlankImage(cols, rows int, c color.Color) image.Image {
	// We use NRGBA to create image with an alpha filter independent of the set colors
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows)) 
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x, y, c) // set the coords color
		}
	}
	return img
}

func NewPxCanvas(state *apptype.State, config apptype.PxCanvasConfig) *PxCanvas {
	pxCanvas := &PxCanvas {
		PxCanvasConfig: config,
		appState: state,
	}
	pxCanvas.PixelData = NewBlankImage(pxCanvas.PxCols, pxCanvas.PxRows, color.NRGBA{128, 128, 125, 255})
	pxCanvas.ExtendBaseWidget(pxCanvas)
	return pxCanvas
}

func (pxCanvas *PxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++{
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

/*
We try to pan the image/. This is dependent on the systems permission to do so.
It may require pressing additional button to effect the panning.
*/
func (pxCanvas *PxCanvas) TryPan(previousCoord *fyne.PointEvent, ev *desktop.MouseEvent) {
	// Allow panning ONLY through the scrollwheel of the mouse.
	if previousCoord != nil && ev.Button == desktop.MouseButtonTertiary {
		pxCanvas.Pan(*previousCoord, ev.PointEvent)
	}
	/*
		If no scroll wheel, use below
		if previousCoord != nil && ev.Button == desktop.MouseButtonPrimary {
        pxCanvas.Pan(*previousCoord, ev.PointEvent)
    }
	*/
}

// Brushable interface
func (pxCanvas PxCanvas) SetColor(c color.Color, x, y int) {
	// access the interface to check the underlaying image type
	if nrgba, ok := pxCanvas.PixelData.(*image.NRGBA); ok { 
		nrgba.Set(x,y,c) // If it's then set the color of the nrgba
	}
	
	if rgba, ok := pxCanvas.PixelData.(*image.RGBA); ok { 
		rgba.Set(x,y,c) // If it's then set the color of the rgba
	}
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int) {
	bounds := pxCanvas.Bounds()
	// Get the location of the mouse and ensure it's on top of the image canvas, if not. Dont do anything
	if !InBounds(ev.Position, bounds) {
		return nil,nil
	}
	
	// copy information needed for calculation
	pxSize := float32(pxCanvas.PxSize)
	xOffset := pxCanvas.CanvasOffset.X
	yOffset := pxCanvas.CanvasOffset.Y

	x := int((ev.Position.X - xOffset) / pxSize)
	y := int((ev.Position.Y - yOffset) / pxSize)

	return &x, &y

}

/**
This function loads image onto the canvas
*/
func (pxCanvas *PxCanvas) LoadImage(img image.Image) {
	dimensions := img.Bounds()

	pxCanvas.PxCanvasConfig.PxCols = dimensions.Dx()
	pxCanvas.PxCanvasConfig.PxRows = dimensions.Dy()

	pxCanvas.PixelData = img
	pxCanvas.reloadImage = true
	pxCanvas.Refresh()
}

func (pxCanvas PxCanvas) NewDrawing(cols, rows int) {
	pxCanvas.appState.SetFilePath("")
	pxCanvas.PxCols = cols
	pxCanvas.PxRows = rows
	pixelData := NewBlankImage(cols, rows, color.NRGBA{128,128,128, 255})
	pxCanvas.LoadImage(pixelData)
}
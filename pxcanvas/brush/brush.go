package brush

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"yabacoder.com/pixl/apptype"
)

const (
	Pixel = iota
)

func Cursor(config apptype.PxCanvasConfig, brush apptype.BrushType, ev *desktop.MouseEvent, x int, y int) []fyne.CanvasObject {
	var objects []fyne.CanvasObject
	switch  {
	case brush == Pixel:
		pxSize := float32(config.PxSize)
		xOrigin := (float32(x) * pxSize) + config.CanvasOffset.X // Upper left corner of our virtual object
		yOrigin := (float32(y) * pxSize) + config.CanvasOffset.Y

		cursorColor := color.NRGBA{80, 80, 80, 255}

		//we create boxed line around our above

		left := canvas.NewLine(cursorColor)
		left.StrokeWidth = 3
		left.Position1 = fyne.NewPos(xOrigin, yOrigin) // start of a line
		left.Position2 = fyne.NewPos(xOrigin, yOrigin+pxSize) // end if the line with the size of the pixel picked.
		
		top := canvas.NewLine(cursorColor)
		top.StrokeWidth = 3
		top.Position1 = fyne.NewPos(xOrigin, yOrigin) // start of a line
		top.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin) // end if the line with the size of the pixel picked.
		
		right := canvas.NewLine(cursorColor)
		right.StrokeWidth = 3
		right.Position1 = fyne.NewPos(xOrigin+pxSize, yOrigin) // start of a line
		right.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin+pxSize) // end if the line with the size of the pixel picked.
		
		bottom := canvas.NewLine(cursorColor)
		bottom.StrokeWidth = 3
		bottom.Position1 = fyne.NewPos(xOrigin, yOrigin+pxSize) // start of a line
		bottom.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin+pxSize) // end if the line with the size of the pixel picked.
		
		objects = append(objects, left, top, right, bottom) //surround the object with these lines.
	}
	// return the object
	return objects
}
// Attempt to brush on the canvas
func TryBrush(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	switch  {
	case appState.BrushType == Pixel: //Check the selected brush
		return TryPaintPixel(appState, canvas, ev)
	default:
		return false
	}
}

func TryPaintPixel(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	x,y := canvas.MouseToCanvasXY(ev)
	if x == nil && y != nil && ev.Button == desktop.MouseButtonPrimary { // check if left side of the mouse was clicked
		canvas.SetColor(appState.BrushColor, *x, *y) //Update the appstate with the new color
		return true //Return true
	}
	return false
}
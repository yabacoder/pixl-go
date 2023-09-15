package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	pxCanvas *PxCanvas
	canvasImage *canvas.Image
	canvasBorder []canvas.Line
	canvasCursor []fyne.CanvasObject // cursor
}
// set our canvas cursor to whatever object we supllied as parameters.
func (renderer *PxCanvasRenderer) SetCursor(objects []fyne.CanvasObject) {
	renderer.canvasCursor = objects
}

// WidgetRenderer interface implementation
func(renderer *PxCanvasRenderer) MinSize() fyne.Size {
	return renderer.pxCanvas.DrawingArea
}

// WidgetRenderer interface implementation
func(renderer *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)
	for i := 0; i < len(renderer.canvasBorder); i++ {
		// Draw line across the border and append to the objects
		objects = append(objects, &renderer.canvasBorder[i]) 
	}
	// the 4 borders and the image are added to the object and drawn to the screen
	objects = append(objects, renderer.canvasImage)
	objects = append(objects, renderer.canvasCursor...)
	return objects
}

// WidgetRenderer interface implementation
func(renderer *PxCanvasRenderer) Destroy() {
}

// WidgetRenderer interface implementation // 3 diff implementations are require for the border,
// 1. is the main layout. 2. is for the border 3. for image display
func(renderer *PxCanvasRenderer) Layout(size fyne.Size) {
	renderer.LayoutCanvas(size)
	renderer.LayoutBorder(size)
}

// Refresh function reloads the image and display the modified data supplied
func (renderer *PxCanvasRenderer) Refresh() {
	if renderer.pxCanvas.reloadImage {
		renderer.canvasImage = canvas.NewImageFromImage(renderer.pxCanvas.PixelData)
		renderer.canvasImage.ScaleMode = canvas.ImageScalePixels // This allows us to flexibly scale our image to the best resolution
		renderer.canvasImage.FillMode = canvas.ImageFillContain // allows image to be contained within the size we specified.
		renderer.pxCanvas.reloadImage = false
	}
	renderer.Layout(renderer.pxCanvas.Size())
	canvas.Refresh(renderer.canvasImage)
}
// Image layout implementation
func(renderer *PxCanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgPxWidth := renderer.pxCanvas.PxCols
	imgPxHeight := renderer.pxCanvas.PxRows
	pxSize := renderer.pxCanvas.PxSize
	// Mover the image canvas to the right location inside the drawing canvas
	renderer.canvasImage.Move(fyne.NewPos(renderer.pxCanvas.CanvasOffset.X, renderer.pxCanvas.CanvasOffset.Y))
	// Move it to the right spot we want it to be.
	renderer.canvasImage.Resize(fyne.NewSize(float32(imgPxWidth*pxSize), float32(imgPxHeight*pxSize)))
}

// Border layout implementation
func(renderer *PxCanvasRenderer) LayoutBorder(size fyne.Size) {
	//We will draw our border to surrond the main canvas while leaving the image layout intact.
	offset := renderer.pxCanvas.CanvasOffset
	imgHeight := renderer.canvasImage.Size().Height
	imgWidth := renderer.canvasImage.Size().Width

	// Drawing the borders

	//Left border
	left := &renderer.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y) // the position of the border
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight) // the height of the border
	
	//Top border
	top := &renderer.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y) // the position of the border
	top.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y) // the height of the border

	//Right border
	right := &renderer.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X+imgWidth, offset.Y) // the position of the border
	right.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight) // the height of the border

	//Bottom border
	bottom := &renderer.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y+imgHeight) // the position of the border
	bottom.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight) // the height of the border
}
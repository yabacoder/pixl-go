package pxcanvas

import "fyne.io/fyne/v2"

/**
Whenever this function is called, it picks the current coordinate, 
@params Previous coordinate stored in the state
@params Current/active coordinate as mouse pointer. 
*/
func (pxCanvas *PxCanvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	// We will calcalute the difference between the 2 to get where we should move the canvas
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	// Once we know the difference, it's good to change the offset.
	pxCanvas.CanvasOffset.X = xDiff
	pxCanvas.CanvasOffset.Y = yDiff
	pxCanvas.Refresh() //ensure this is reflected on the page

}

func (pxCanvas *PxCanvas) scale(direction int) {
	// track the direction of the mouse
	// positive or negative
	switch {
	case direction > 0:
		// if the number is position incrase the pixel by 1 unit
		pxCanvas.PxSize += 1
	case direction < 0:
		if pxCanvas.PxSize > 2 { // only reduce the size when it's greater than 2.
									// this ensures that the pixel size it's never les than 1
			pxCanvas.PxSize -= 1 // if it's less than zero, decrease it by 1
		}
	default:
		pxCanvas.PxSize = 10 // if it's zero make it 10 as default
	}
}
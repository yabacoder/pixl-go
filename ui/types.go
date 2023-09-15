package ui

import (
	"fyne.io/fyne/v2"
	"yabacoder.com/pixl/apptype"
	"yabacoder.com/pixl/pxcanvas"
	"yabacoder.com/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State *apptype.State
	Swatches []*swatch.Swatch
}

package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"yabacoder.com/pixl/swatch"
)

func BuildSwatches(app *AppInit) *fyne.Container {
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)
	for i := 0; i < cap(app.Swatches); i++ {
		initialColor := color.NRGBA{255,255,255,255}
		s := swatch.NewSwatch(app.State, initialColor, i, func (s *swatch.Swatch) {
			// Remove all borders around the swatches
			for j := 0; j < len(app.Swatches); j++ {
				app.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}
			// Update the app state to show that a swatch is select and return the index
			app.State.SwatchSelected = s.SwatchIndex
			// Brush color is changed to the selected color.
			app.State.BrushColor = s.Color

		})
		if i == 0 {
			// set the selcted swatch as default when the app is opened
			s.Selected = true
			app.State.SwatchSelected = 0
			s.Refresh() // refreshes the swatch to update the change
		}
		app.Swatches = append(app.Swatches, s)
		canvasSwatches = append(canvasSwatches, s)
	}

	// GridWraps only operate on canvas objects
	return container.NewGridWrap(fyne.NewSize(20,20), canvasSwatches...)
}
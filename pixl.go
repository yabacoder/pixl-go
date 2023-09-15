package main

import (
	"image/color"

	"yabacoder.com/pixl/apptype"
	"yabacoder.com/pixl/pxcanvas"
	"yabacoder.com/pixl/swatch"
	"yabacoder.com/pixl/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("Pixl - Yabacoder")

	// Set the current state and color for the app
	state := apptype.State {
		BrushColor: color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	//Pixl Canvas configuration
	pixlCanvasConfig := apptype.PxCanvasConfig {
		DrawingArea: fyne.NewSize(600,600),
		CanvasOffset: fyne.NewPos(0,0),
		PxRows: 10,
		PxCols: 10,
		PxSize: 30,
	}

	pixlCanvas := pxcanvas.NewPxCanvas(&state, pixlCanvasConfig)

	appInit := ui.AppInit{
		PixlCanvas: pixlCanvas,
		PixlWindow: pixlWindow,
		State: &state,
		Swatches: make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun() 
}
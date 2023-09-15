package ui

import (
	"fyne.io/fyne/v2/container"

	// "yabacoder.com/pixl/apptype"
	// "yabacoder.com/pixl/swatch"
)

func Setup(app *AppInit) {
	SetupMenus(app)
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)
	// Use app.PixlCanvas to extract the container for the main editor area. 
	// this has been predefined in the app config file in pixl.go
	appLayout := container.NewBorder(nil, swatchesContainer,nil, colorPicker, app.PixlCanvas)


	app.PixlWindow.SetContent(appLayout)
}
package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	swatchContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)

	appLayout := container.NewBorder(nil, swatchContainer, nil, colorPicker)

	app.PixlWindow.SetContent(appLayout)
}
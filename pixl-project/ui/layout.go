package ui

func Setup(app *AppInit) {
	swatchContainer := BuildSwatches(app)

	app.PixlWindow.SetContent(swatchContainer)
}
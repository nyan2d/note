package app

func (app *App) hWfoldersAdd() {
	app.ShowAlert(app.foldersWindow.window, "Hello?")
}

func (app *App) hWfoldersRemove() {
	app.ShowAlert(app.foldersWindow.window, "Hello?")
}

func (app *App) hWfoldersRename() {
	app.ShowAlert(app.foldersWindow.window, "Hello?")
}

func (app *App) hWfoldersClose() {
	app.foldersWindow.Hide()
}

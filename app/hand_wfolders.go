package app

import (
	"fmt"
	"strconv"
)

func (app *App) hWfoldersAdd() {
	s := fmt.Sprintf("Canvas: %+v\nContent: %+v\nContainer: %+v\nList: %+v\n",
		app.foldersWindow.window.Canvas().Size(),
		app.foldersWindow.window.Content().Size(),
		app.foldersWindow.container.Size(),
		app.foldersWindow.FoldersList.Size(),
	)
	app.ShowAlert(app.foldersWindow.window, s)
}

func (app *App) hWfoldersRemove() {
	// app.ShowAlert(app.foldersWindow.window, "Hello?")
	app.ShowAlert(app.foldersWindow.window, strconv.Itoa(app.foldersWindow.FoldersList.SelectedID()))
}

func (app *App) hWfoldersRename() {
	app.ShowAlert(app.foldersWindow.window, "Hello?")
	app.foldersWindow.FoldersList.Refresh()
}

func (app *App) hWfoldersClose() {
	app.foldersWindow.Hide()
}

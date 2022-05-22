package app

import (
	"github.com/nyan2d/note/model"
	"github.com/nyan2d/smartbolt"
)

func (a *App) hMenuFileQuit() {
	a.App.Quit()
}

func (a *App) hMenuEditFolders() {
	a.foldersWindow.Show()
}

func (a *App) hMenuEditRemove() {
	if a.navigationPage.selectedNoteID > -1 {
		notes := smartbolt.OpenBucket[int, model.Note](a.DB, "notes")
		notes.Delete(a.navigationPage.selectedNoteID)
		a.navigationPage.NotesList.Refresh()
		a.SwitchPage(a.homePage)
	}
}

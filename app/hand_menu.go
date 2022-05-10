package app

import (
	"github.com/nyan2d/note/model"
	"github.com/nyan2d/smartbolt"
)

func (a *App) hMenuFileQuit() {
	a.app.Quit()
}

func (a *App) hMenuEditFolders() {
	a.foldersWindow.Show()
}

func (a *App) hMenuEditRemove() {
	if a.navigationPage.selectedNoteID > -1 {
		notes := smartbolt.OpenBucket[int, model.Note](a.db, "notes")
		notes.Delete(a.navigationPage.selectedNoteID)
        a.navigationPage.NotesList.Refresh()
        a.SwitchPage(a.homePage)
	}
}

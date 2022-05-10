package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/nyan2d/note/model"
	"github.com/nyan2d/smartbolt"
)

type App struct {
	// general
	app        fyne.App
	mainWindow fyne.Window
	mainMenu   *MainMenu

	// database
	db *smartbolt.Smartbolt

	// items
	folders *model.Folders
	notes   *model.Notes

	// pages
	navigationPage *NavigationPage
	homePage       *HomePage
	previewPage    *PreviewPage
	editPage       *EditPage

	// windows
	foldersWindow *FoldersWindow
}

func NewApp(database *smartbolt.Smartbolt) *App {
	x := &App{}

	x.app = app.New()
	x.mainWindow = x.app.NewWindow("Note")
	x.mainMenu = NewMainMenu()
	x.db = database

	x.folders = x.readAllFolders()
	x.notes = x.readAllNotes()
	x.navigationPage = NewNavigationPage(x.folders, x.notes)
	x.homePage = NewHomePage()
	x.previewPage = NewPreviewPage()
	x.editPage = NewEditPage(*x.folders)

	x.foldersWindow = NewFoldersWindow(x.app)

	x.bindHandlers()

	x.mainWindow.SetMainMenu(x.mainMenu.Menu())
	x.SwitchPage(x.homePage)

	return x
}

func (a *App) Run() {
	a.mainWindow.ShowAndRun()
}

func (a *App) SwitchPage(page Page) {
	x := container.NewHSplit(a.navigationPage.CanvasObject(), page.CanvasObject())
	a.mainWindow.SetContent(x)
}

func (a *App) ShowAlert(text string) {
	label := widget.NewLabel(text)
	popup := widget.NewPopUp(label, a.mainWindow.Canvas())
	popup.Show()

	windowSize := a.mainWindow.Canvas().Size()
	labelSize := label.Size()
	popupPos := fyne.NewPos(windowSize.Width/2-labelSize.Width/2, windowSize.Height/2-labelSize.Height/2)
	popup.Move(popupPos)
}

func (a *App) bindHandlers() {
	a.navigationPage.CreateButton.OnTapped = a.hNavCreate
	a.navigationPage.SearchEntry.OnChanged = a.hNavSearchTextChanged
	a.navigationPage.NotesList.OnSelected = a.hNavListElementSelected

	a.mainMenu.FileQuitMenuItem.Action = a.hMenuFileQuit
	a.mainMenu.EditFoldersMenuItem.Action = a.hMenuEditFolders
	a.mainMenu.EditRemoveMenuItem.Action = a.hMenuEditRemove

	a.editPage.SaveButton.OnTapped = a.hEditpageSave
	a.editPage.CancelButton.OnTapped = a.hEditpageCancel
}

func (a *App) readAllFolders() *model.Folders {
	fbucket := smartbolt.OpenBucket[int, model.Folder](a.db, "folders")
	folders, _ := fbucket.GetAll()
	if len(folders) == 0 {
		generalID := fbucket.NextID()
		trashID := fbucket.NextID()
		fbucket.PutBulk(map[int]model.Folder{
			generalID: {ID: generalID, Name: "General"},
			trashID:   {ID: trashID, Name: "Trash"},
		})
		return a.readAllFolders()
	}
	fders := model.Folders(folders)
	return &fders
}

func (a *App) readAllNotes() *model.Notes {
	return model.NewNotes(a.db)
}

package app

import "fyne.io/fyne/v2"

type Page interface {
	CanvasObject() fyne.CanvasObject
}

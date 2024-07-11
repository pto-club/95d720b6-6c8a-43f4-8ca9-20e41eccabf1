package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type gui struct {
	win fyne.Window
	fileName widget.Label
}


func (w *gui) Header() *fyne.Container{
	open_file := widget.NewToolbarAction(theme.FolderOpenIcon(), w.open_file_onActivate)
	toolbar := widget.NewToolbar(open_file)

	logo := Logo()

	nameApp := widget.NewLabelWithStyle("Catalogul informativ de prețuri la materiale, \n echipamente și utilaje aplicate în achizițiile publice", 
	fyne.TextAlignCenter,widget.RichTextStyleCodeBlock.TextStyle)

	header := container.NewStack(toolbar, logo)
	vertical := container.NewVBox(header, nameApp)
	return vertical
}

func (w *gui) Center() *fyne.Container{
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Search string...hhhhh")
	entry.Resize(fyne.NewSize(30, 20))
	entry.OnSubmitted = func(s string) {} // string entered in textbox

	table := widget.NewTableWithHeaders(func() (rows int, cols int) {
		return 10, 9
	}, func() fyne.CanvasObject {
		return widget.NewLabel("test")
	}, func(tci widget.TableCellID, co fyne.CanvasObject) {

	},
)
table.SetColumnWidth(2, 200)
table.StickyRowCount = 10

	cont := container.NewPadded(table)

	center := container.NewVBox(entry, cont)	
	return center
}



func (w *gui) open_file_onActivate() {
	fd := dialog.NewFileOpen(
		func(file fyne.URIReadCloser, err error) {
			if err != nil {
				fmt.Printf("Error to open file... %s ", file.URI().Name())
				return
			}
			if file == nil {
				return //if press cancel
			}
			w.win.SetTitle(file.URI().Name())
			w.fileName.SetText(file.URI().Name())
		}, w.win)
	//fd.SetFilter(storage.NewExtensionFileFilter([]string{".db"}))
	fd.Show()
}

func Logo() *canvas.Image{
	logo := canvas.NewImageFromResource(resourceLogoPng)
	logo.FillMode = canvas.ImageFillContain
	return logo
}
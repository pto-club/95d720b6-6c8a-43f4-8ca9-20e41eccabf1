package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	application.Settings().SetTheme(NewSmet4ikTheme())

	window := application.NewWindow("Price List smet4ik.CLUB")
	window.Resize(fyne.NewSize(900, 500))
	window.CenterOnScreen()
	window.SetFixedSize(true)

	mainWindows := &gui{win: window}

	header := mainWindows.Header()	
	center := mainWindows.Center()

	space := widget.NewLabel("                                            ")
	
	main_container := container.NewBorder(header, nil, space, space, center)
	
	

	mainWindows.win.SetContent(main_container)
	mainWindows.win.Show()
	application.Run()

}


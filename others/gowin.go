package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	apps := app.New()

	w := apps.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Eric!"),
		widget.NewButton("Quit", func() {
			apps.Quit()
		}),
	))

	w.ShowAndRun()
}

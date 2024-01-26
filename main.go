package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Open Wallet")
	w.Resize(fyne.NewSize(300, 300))

	introduction := canvas.NewText("Hello Fyne!", nil)
	introduction.TextSize = 24
	introduction.Alignment = fyne.TextAlignCenter
	// introduction
	// introduction.
	// introduction.Alignment = fyne.TextAlignCenter
	// introduction.TextStyle
	w.SetContent(container.NewVBox(
		introduction,
		widget.NewButton("Hi!", func() {
			introduction.Text = "Welcome :)"
			introduction.Refresh()
		}),
	))

	w.ShowAndRun()
}

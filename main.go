package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	mainWindow := app.NewWindow("gemix - API fuzztester")

	radio := widget.NewRadioGroup([]string{"POST", "GET", "PUT", "DELETE"}, func(selected string) {
		fmt.Println(selected)
	})
	radio.Horizontal = true
	radio.Required = true

	text := canvas.NewText("Choose HTTP method:", color.White)

	mainContent := container.NewVBox(
		container.NewCenter(text),
		container.NewCenter(radio),
	)

	mainWindow.SetContent(mainContent)
	mainWindow.Resize(fyne.NewSize(500, 300))
	mainWindow.ShowAndRun()
}

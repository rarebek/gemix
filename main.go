package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var selectedHttpMethod string

func main() {
	app := app.New()
	mainWindow := app.NewWindow("gemix - API fuzztester")

	radio := widget.RadioGroup{
		DisableableWidget: widget.DisableableWidget{},
		Horizontal:        true,
		Required:          true,
		OnChanged: func(selected string) {
			selectedHttpMethod = selected
			fmt.Println(selectedHttpMethod)
		},
		Options:  []string{"POST", "GET", "PUT", "DELETE"},
		Selected: "",
	}

	label := widget.NewLabel("Select HTTP method:")

	mainWindow.SetContent(container.NewVBox(label, &radio))
	mainWindow.Resize(fyne.NewSize(500, 300))
	mainWindow.ShowAndRun()
}

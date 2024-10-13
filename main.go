package main

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TestResult struct {
	ID      int
	Status  string
	Message string
}

type APIFuzzTester struct {
	Method      string
	File        fyne.URI
	Loading     bool
	Results     []TestResult
	GeminiToken string
	TempToken   string
	MainWindow  fyne.Window
	TokenDialog dialog.Dialog
}

func NewAPIFuzzTester(w fyne.Window) *APIFuzzTester {
	return &APIFuzzTester{
		Method:      "GET",
		MainWindow:  w,
		Results:     []TestResult{},
		Loading:     false,
		GeminiToken: "",
	}
}

func (a *APIFuzzTester) renderMethodSelection() *widget.RadioGroup {
	return widget.NewRadioGroup([]string{"GET", "POST", "PUT", "DELETE"}, func(value string) {
		a.Method = value
	})
}

func (a *APIFuzzTester) handleFileSelect(file fyne.URIReadCloser, err error) {
	if err != nil {
		dialog.ShowError(err, a.MainWindow)
		return
	}
	if file != nil {
		a.File = file.URI()
		fmt.Println("Selected file:", a.File)
	}
}

func (a *APIFuzzTester) showFileDialog() {
	dialog.ShowFileOpen(a.handleFileSelect, a.MainWindow)
}

func (a *APIFuzzTester) handleSubmit() {
	if a.GeminiToken == "" {
		dialog.ShowInformation("Error", "Please set your Gemini token.", a.MainWindow)
		return
	}

	a.Loading = true
	a.Results = []TestResult{}
	go func() {
		time.Sleep(2 * time.Second) 

		mockResults := []TestResult{}
		for i := 1; i <= 10; i++ {
			result := TestResult{
				ID:      i,
				Status:  randomStatus(),
				Message: randomMessage(),
			}
			mockResults = append(mockResults, result)
		}
		a.Loading = false
		a.Results = mockResults

		a.MainWindow.Content().Refresh()
	}()
}

func randomStatus() string {
	if rand.Float32() > 0.3 {
		return "success"
	}
	return "failure"
}

func randomMessage() string {
	if rand.Float32() > 0.3 {
		return "Test passed successfully"
	}
	return "Unexpected response received"
}

func (a *APIFuzzTester) handleSetToken() {
	a.GeminiToken = a.TempToken
	a.TokenDialog.Hide()
}

func (a *APIFuzzTester) showTokenDialog() {
	content := widget.NewForm(
		widget.NewFormItem("Token", widget.NewEntry()),
	)
	dialog := dialog.NewForm("Set Gemini Token", "Set", "Cancel", content.Items, func(b bool) {
		if b {
			a.handleSetToken()
		}
	}, a.MainWindow)
	dialog.Resize(fyne.NewSize(400, 200))
	dialog.Show()
	a.TokenDialog = dialog
}

func (a *APIFuzzTester) renderFuzzTestResults() *fyne.Container {
	if len(a.Results) == 0 {
		return container.NewCenter(canvas.NewText("No test results yet. Start fuzztesting to see results here.", theme.ForegroundColor()))
	}

	var resultLabels []fyne.CanvasObject
	for _, result := range a.Results {
		resultText := fmt.Sprintf("ID: %d - %s - %s", result.ID, result.Status, result.Message)
		resultLabels = append(resultLabels, widget.NewLabel(resultText))
	}
	return container.NewVBox(resultLabels...)
}

func main() {
	a := app.New()
	w := a.NewWindow("API Fuzz Tester")

	tester := NewAPIFuzzTester(w)

	methodLabel := widget.NewLabel("Select HTTP Method")
	methodSelection := tester.renderMethodSelection()
	fileUploadButton := widget.NewButton("Upload a file", func() {
		tester.showFileDialog()
	})
	setTokenButton := widget.NewButtonWithIcon("Set Gemini Token", theme.ConfirmIcon(), func() {
		tester.showTokenDialog()
	})
	startButton := widget.NewButton("Start Fuzztesting", func() {
		tester.handleSubmit()
	})
	startButton.Importance = widget.HighImportance

	uploadDescription := widget.NewLabel("JSON file up to 10MB")
	fileUploadContainer := container.NewVBox(
		layout.NewSpacer(),
		fileUploadButton,
		uploadDescription,
		layout.NewSpacer(),
	)

	leftSide := container.NewVBox(
		methodLabel,
		methodSelection,
		widget.NewLabel("Upload JSON Request Body"),
		fileUploadContainer,
		startButton,
		setTokenButton,
	)

	resultsLabel := canvas.NewText("Fuzz Test Results", theme.ForegroundColor())
	resultsContainer := tester.renderFuzzTestResults()
	rightSide := container.NewBorder(resultsLabel, nil, nil, nil, resultsContainer)

	content := container.NewHSplit(
		leftSide,
		rightSide,
	)
	content.Offset = 0.3

	w.SetContent(content)

	w.Resize(fyne.NewSize(800, 500))
	w.ShowAndRun()
}

package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainWindow struct {
	app    fyne.App
	window fyne.Window
}

func NewMainWindow() *MainWindow {
	a := app.New()
	w := a.NewWindow("Ellie")
	w.Resize(fyne.NewSize(600, 400))

	// Set custom theme
	a.Settings().SetTheme(EllieTheme{})

	return &MainWindow{
		app:    a,
		window: w,
	}
}

func (mw *MainWindow) Show() {
	// Create command input
	commandInput := widget.NewEntry()
	commandInput.SetPlaceHolder("Enter command...")

	// Create results area
	results := widget.NewTextGrid()
	results.SetText("Welcome to Ellie")

	// Create command buttons
	buttons := container.NewGridWithColumns(3,
		widget.NewButton("System", func() { mw.showSystemCommands() }),
		widget.NewButton("Git", func() { mw.showGitCommands() }),
		widget.NewButton("Services", func() { mw.showServiceCommands() }),
		widget.NewButton("Network", func() { mw.showNetworkCommands() }),
		widget.NewButton("Todo", func() { mw.showTodoCommands() }),
		widget.NewButton("Projects", func() { mw.showProjectCommands() }),
	)

	// Create main content
	content := container.NewVBox(
		commandInput,
		buttons,
		results,
	)

	mw.window.SetContent(content)
	mw.window.Show()
	mw.app.Run()
}

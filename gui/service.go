package gui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (mw *MainWindow) showServiceCommands() {
	// Create service selection
	serviceSelect := widget.NewSelect([]string{"Apache", "MySQL", "PostgreSQL", "All"}, func(s string) {})
	serviceSelect.SetSelected("Apache")

	// Create action buttons
	actions := container.NewHBox(
		widget.NewButton("Start", func() {
			mw.executeCommand("start", serviceSelect.Selected)
		}),
		widget.NewButton("Stop", func() {
			mw.executeCommand("stop", serviceSelect.Selected)
		}),
		widget.NewButton("Restart", func() {
			mw.executeCommand("restart", serviceSelect.Selected)
		}),
	)

	// Create status display
	status := widget.NewTextGrid()
	status.SetText("Service status will be displayed here")

	// Create refresh button
	refresh := widget.NewButton("Refresh Status", func() {
		mw.executeCommand("status", serviceSelect.Selected)
	})

	// Create main content
	content := container.NewVBox(
		widget.NewLabel("Select Service:"),
		serviceSelect,
		actions,
		widget.NewSeparator(),
		status,
		refresh,
	)

	// Show dialog
	dialog.ShowCustom("Service Management", "Close", content, mw.window)
}

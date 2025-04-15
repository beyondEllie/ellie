package gui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func (mw *MainWindow) showProjectCommands() {
	// Create project list
	projectList := widget.NewTextGrid()
	projectList.SetText("Loading projects...")

	// Create action buttons
	actions := container.NewHBox(
		widget.NewButton("Add Project", func() {
			mw.showAddProjectInput()
		}),
		widget.NewButton("Delete Project", func() {
			mw.showDeleteProjectInput()
		}),
		widget.NewButton("Switch Project", func() {
			mw.showSwitchProjectInput()
		}),
	)

	// Create search
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Search projects...")
	searchButton := widget.NewButton("Search", func() {
		mw.executeCommand("project", "search", searchEntry.Text)
	})
	searchBox := container.NewBorder(nil, nil, nil, searchButton, searchEntry)

	// Create refresh button
	refresh := widget.NewButton("Refresh List", func() {
		mw.executeCommand("project", "list")
	})

	// Create main content
	content := container.NewVBox(
		searchBox,
		projectList,
		widget.NewSeparator(),
		actions,
		refresh,
	)

	// Show dialog
	dialog.ShowCustom("Project Management", "Close", content, mw.window)
}

func (mw *MainWindow) showAddProjectInput() {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter project name...")

	pathEntry := widget.NewEntry()
	pathEntry.SetPlaceHolder("Enter project path...")

	descEntry := widget.NewMultiLineEntry()
	descEntry.SetPlaceHolder("Enter project description...")

	tagsEntry := widget.NewEntry()
	tagsEntry.SetPlaceHolder("Enter tags (comma-separated)...")

	content := container.NewVBox(
		widget.NewLabel("Add New Project:"),
		nameEntry,
		pathEntry,
		descEntry,
		tagsEntry,
	)

	dialog.ShowCustomConfirm("Add Project", "Add", "Cancel", content,
		func(b bool) {
			if b {
				mw.executeCommand("project", "add", nameEntry.Text, pathEntry.Text, descEntry.Text, tagsEntry.Text)
			}
		}, mw.window)
}

func (mw *MainWindow) showDeleteProjectInput() {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter project name...")

	content := container.NewVBox(
		widget.NewLabel("Delete Project:"),
		nameEntry,
	)

	dialog.ShowCustomConfirm("Delete Project", "Delete", "Cancel", content,
		func(b bool) {
			if b {
				mw.executeCommand("project", "delete", nameEntry.Text)
			}
		}, mw.window)
}

func (mw *MainWindow) showSwitchProjectInput() {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter project name...")

	content := container.NewVBox(
		widget.NewLabel("Switch to Project:"),
		nameEntry,
	)

	dialog.ShowCustomConfirm("Switch Project", "Switch", "Cancel", content,
		func(b bool) {
			if b {
				mw.executeCommand("switch", nameEntry.Text)
			}
		}, mw.window)
}

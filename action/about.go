package actions

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ShowAboutWindow displays the about window for Ellie
func ShowAboutWindow(args []string) {
	// Create a new Fyne application
	a := app.New()
	w := a.NewWindow("About Ellie")

	// Set window size
	w.Resize(fyne.NewSize(600, 400))

	// Create content
	title := widget.NewLabel("Ellie - The AI-Powered CLI Companion")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	version := widget.NewLabel("Version: 0.0.11")
	version.Alignment = fyne.TextAlignCenter

	description := widget.NewLabel("Your all-in-one terminal buddy for system management, Git workflows, and productivity hacks.")
	description.Wrapping = fyne.TextWrapWord
	description.Alignment = fyne.TextAlignCenter

	features := widget.NewLabel("Core Features:\n" +
		"• System Management\n" +
		"• Git Workflows\n" +
		"• Todo Management\n" +
		"• Project Management\n" +
		"• Network Management\n" +
		"• AI Integration")
	features.Wrapping = fyne.TextWrapWord

	author := widget.NewLabel("Built with ❤️ by Tachera Sasi")
	author.Alignment = fyne.TextAlignCenter

	// Create a scrollable container for the content
	content := container.NewVBox(
		title,
		version,
		description,
		widget.NewSeparator(),
		features,
		widget.NewSeparator(),
		author,
	)

	// Create a scroll container
	scroll := container.NewScroll(content)
	scroll.Resize(fyne.NewSize(580, 380))

	// Set the content
	w.SetContent(scroll)

	// Show and run the window
	w.Show()
	a.Run()
}

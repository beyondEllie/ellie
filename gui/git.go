package gui

// import (
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/dialog"
// 	"fyne.io/fyne/v2/widget"
// )

// func (mw *MainWindow) showGitCommands() {
// 	// Create dialog content
// 	content := container.NewVBox(
// 		widget.NewButton("Setup Git", func() {
// 			mw.executeCommand("setup-git")
// 		}),
// 		widget.NewButton("Git Status", func() {
// 			mw.executeCommand("git", "status")
// 		}),
// 		widget.NewButton("Git Push", func() {
// 			mw.executeCommand("git", "push")
// 		}),
// 		widget.NewButton("Git Pull", func() {
// 			mw.executeCommand("git", "pull")
// 		}),
// 		widget.NewButton("Create Commit", func() {
// 			mw.showCommitInput()
// 		}),
// 	)

// 	// Show dialog
// 	dialog.ShowCustom("Git Commands", "Close", content, mw.window)
// }

// func (mw *MainWindow) showCommitInput() {
// 	entry := widget.NewEntry()
// 	entry.SetPlaceHolder("Enter commit message...")

// 	content := container.NewVBox(
// 		widget.NewLabel("Enter commit message:"),
// 		entry,
// 	)

// 	dialog.ShowCustomConfirm("Create Commit", "Commit", "Cancel", content,
// 		func(b bool) {
// 			if b {
// 				mw.executeCommand("git", "commit", entry.Text)
// 			}
// 		}, mw.window)
// }

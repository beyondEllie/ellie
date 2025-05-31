package gui

// import (
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/dialog"
// 	"fyne.io/fyne/v2/widget"
// )

// func (mw *MainWindow) showSystemCommands() {
// 	content := container.NewVBox(
// 		widget.NewButton("System Info", func() { mw.executeCommand("sysinfo") }),
// 		widget.NewButton("Open Explorer", func() { mw.executeCommand("open-explorer") }),
// 		widget.NewButton("List Dir", func() { mw.showDirectoryInput() }),
// 		widget.NewButton("Create File", func() { mw.showCreateFileInput() }),
// 		widget.NewButton("Install", func() { mw.showPackageInput() }),
// 		widget.NewButton("Update", func() { mw.executeCommand("update") }),
// 	)

// 	dialog.ShowCustom("System", "Close", content, mw.window)
// }

// func (mw *MainWindow) showDirectoryInput() {
// 	entry := widget.NewEntry()
// 	entry.SetPlaceHolder("Enter path...")

// 	content := container.NewVBox(
// 		widget.NewLabel("Path:"),
// 		entry,
// 	)

// 	dialog.ShowCustomConfirm("List Dir", "List", "Cancel", content,
// 		func(b bool) {
// 			if b {
// 				mw.executeCommand("list", entry.Text)
// 			}
// 		}, mw.window)
// }

// func (mw *MainWindow) showCreateFileInput() {
// 	entry := widget.NewEntry()
// 	entry.SetPlaceHolder("Enter path...")

// 	content := container.NewVBox(
// 		widget.NewLabel("Path:"),
// 		entry,
// 	)

// 	dialog.ShowCustomConfirm("Create File", "Create", "Cancel", content,
// 		func(b bool) {
// 			if b {
// 				mw.executeCommand("create-file", entry.Text)
// 			}
// 		}, mw.window)
// }

// func (mw *MainWindow) showPackageInput() {
// 	entry := widget.NewEntry()
// 	entry.SetPlaceHolder("Enter name...")

// 	content := container.NewVBox(
// 		widget.NewLabel("Package:"),
// 		entry,
// 	)

// 	dialog.ShowCustomConfirm("Install", "Install", "Cancel", content,
// 		func(b bool) {
// 			if b {
// 				mw.executeCommand("install", entry.Text)
// 			}
// 		}, mw.window)
// }

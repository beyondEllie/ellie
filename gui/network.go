package gui

// import (
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/dialog"
// 	"fyne.io/fyne/v2/widget"
// )

// func (mw *MainWindow) showNetworkCommands() {
// 	// Create dialog content
// 	content := container.NewVBox(
// 		widget.NewButton("Network Status", func() {
// 			mw.executeCommand("network-status")
// 		}),
// 		widget.NewButton("Connect WiFi", func() {
// 			mw.showWiFiInput()
// 		}),
// 	)

// 	// Show dialog
// 	dialog.ShowCustom("Network Commands", "Close", content, mw.window)
// }

// func (mw *MainWindow) showWiFiInput() {
// 	ssidEntry := widget.NewEntry()
// 	ssidEntry.SetPlaceHolder("Enter WiFi SSID...")

// 	passwordEntry := widget.NewPasswordEntry()
// 	passwordEntry.SetPlaceHolder("Enter WiFi password...")

// 	content := container.NewVBox(
// 		widget.NewLabel("Enter WiFi credentials:"),
// 		ssidEntry,
// 		passwordEntry,
// 	)

// 	dialog.ShowCustomConfirm("Connect WiFi", "Connect", "Cancel", content,
// 		func(b bool) {
// 			if b {
// 				mw.executeCommand("connect-wifi", ssidEntry.Text, passwordEntry.Text)
// 			}
// 		}, mw.window)
// }

package gui

// import (
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/dialog"
// 	"fyne.io/fyne/v2/widget"
// )

// func (mw *MainWindow) showTodoCommands() {
// 	// Create todo list
// 	todoList := widget.NewTextGrid()
// 	todoList.SetText("Loading todo items...")

// 	// Create action buttons
// 	actions := container.NewHBox(
// 		widget.NewButton("Add Todo", func() {
// 			mw.showAddTodoInput()
// 		}),
// 		widget.NewButton("Complete Todo", func() {
// 			mw.showCompleteTodoInput()
// 		}),
// 		widget.NewButton("Delete Todo", func() {
// 			mw.showDeleteTodoInput()
// 		}),
// 	)

// 	// Create refresh button
// 	refresh := widget.NewButton("Refresh List", func() {
// 		mw.executeCommand("todo", "list")
// 	})

// 	// Create main content
// 	content := container.NewVBox(
// 		todoList,
// 		widget.NewSeparator(),
// 		actions,
// 		refresh,
// 	)

// 	// Show dialog
// 	dialog.ShowCustom("Todo Management", "Close", content, mw.window)
// }

// func (mw *MainWindow) showAddTodoInput() {
// 	taskEntry := widget.NewEntry()
// 	taskEntry.SetPlaceHolder("Enter task description...")

// 	categoryEntry := widget.NewEntry()
// 	categoryEntry.SetPlaceHolder("Enter category (optional)...")

// 	prioritySelect := widget.NewSelect([]string{"Low", "Medium", "High"}, func(s string) {})
// 	prioritySelect.SetSelected("Medium")

// 	content := container.NewVBox(
// 		widget.NewLabel("Add New Todo:"),
// 		taskEntry,
// 		categoryEntry,
// 		prioritySelect,
// 	)

// 	dialog.ShowCustomConfirm("Add Todo", "Add", "Cancel", content,
// 		func(b bool) {
// 			if b {
// 				mw.executeCommand("todo", "add", taskEntry.Text, categoryEntry.Text, prioritySelect.Selected)
// 			}
// 		}, mw.window)
// }

// func (mw *MainWindow) showCompleteTodoInput() {
// 	idEntry := widget.NewEntry()
// 	idEntry.SetPlaceHolder("Enter todo ID...")

// 	content := container.NewVBox(
// 		widget.NewLabel("Complete Todo:"),
// 		idEntry,
// 	)

// 	dialog.ShowCustomConfirm("Complete Todo", "Complete", "Cancel", content,
// 		func(b bool) {
// 			if b {
// 				mw.executeCommand("todo", "complete", idEntry.Text)
// 			}
// 		}, mw.window)
// }

// func (mw *MainWindow) showDeleteTodoInput() {
// 	idEntry := widget.NewEntry()
// 	idEntry.SetPlaceHolder("Enter todo ID...")

// 	content := container.NewVBox(
// 		widget.NewLabel("Delete Todo:"),
// 		idEntry,
// 	)

// 	dialog.ShowCustomConfirm("Delete Todo", "Delete", "Cancel", content,
// 		func(b bool) {
// 			if b {
// 				mw.executeCommand("todo", "delete", idEntry.Text)
// 			}
// 		}, mw.window)
// }

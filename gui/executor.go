package gui

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2/dialog"
)

func (mw *MainWindow) executeCommand(cmd string, args ...string) {
	// Create the command
	allArgs := append([]string{cmd}, args...)
	command := exec.Command(os.Args[0], allArgs...)

	// Create buffers for stdout and stderr
	var stdout, stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr

	// Run the command
	err := command.Run()

	// Handle the output
	if err != nil {
		// Show error dialog
		dialog.ShowError(err, mw.window)
		return
	}

	// Get combined output
	output := strings.TrimSpace(stdout.String())
	if output == "" {
		output = strings.TrimSpace(stderr.String())
	}

	// Show success dialog with output
	if output != "" {
		dialog.ShowInformation("Command Output", output, mw.window)
	}
}

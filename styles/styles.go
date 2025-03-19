package styles

import "github.com/fatih/color"

var (
	// Basic colors
	Cyan    = color.New(color.FgCyan)
	Green   = color.New(color.FgGreen)
	Red     = color.New(color.FgRed)
	Yellow  = color.New(color.FgYellow)
	Magenta = color.New(color.FgMagenta)
	White   = color.New(color.FgWhite)
	Blue    = color.New(color.FgBlue)

	// Text styles
	Bold      = color.New(color.Bold)
	Underline = color.New(color.Underline)

	// Custom styles
	HeaderStyle  = color.New(color.FgHiGreen, color.Bold)
	ErrorStyle   = color.New(color.FgHiRed, color.Bold)
	SuccessStyle = color.New(color.FgHiGreen, color.Bold)
	WarningStyle = color.New(color.FgHiYellow, color.Bold)
	InfoStyle    = color.New(color.FgHiCyan, color.Bold)
	DebugStyle   = color.New(color.FgHiMagenta, color.Bold)
	Highlight    = color.New(color.FgHiBlue, color.Bold)
	DimText      = color.New(color.FgHiWhite, color.Faint)
)
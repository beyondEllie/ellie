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

	// Ellie styles
	HeaderStyle   = color.New(color.FgGreen, color.Bold)         // Titles & headers
	ErrorStyle    = color.New(color.FgRed, color.Bold, color.BgBlack) // Errors
	SuccessStyle  = color.New(color.FgHiGreen, color.Bold)       // Success messages
	WarningStyle  = color.New(color.FgYellow, color.Bold)        // Warnings
	InfoStyle     = color.New(color.FgCyan, color.Bold)          // Information
	DebugStyle    = color.New(color.FgMagenta, color.Bold)       // Debug messages
	InputPrompt   = color.New(color.FgBlue, color.Bold)          // Input prompts
	Highlight     = color.New(color.FgHiBlue, color.Bold)        // Highlights key info
	DimText       = color.New(color.FgWhite, color.Faint)        // Subtle text
	InvertedStyle = color.New(color.FgBlack, color.BgWhite)      // Inverted text

)

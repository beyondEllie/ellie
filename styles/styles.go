package styles

import (
	"os"
	"strings"

	"github.com/fatih/color"
)

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
	HeaderStyle   = color.New(color.FgGreen, color.Bold)              // Titles & headers
	ErrorStyle    = color.New(color.FgRed, color.Bold, color.BgBlack) // Errors
	SuccessStyle  = color.New(color.FgHiGreen, color.Bold)            // Success messages
	WarningStyle  = color.New(color.FgYellow, color.Bold)             // Warnings
	InfoStyle     = color.New(color.FgCyan, color.Bold)               // Information
	DebugStyle    = color.New(color.FgMagenta, color.Bold)            // Debug messages
	InputPrompt   = color.New(color.FgBlue, color.Bold)               // Input prompts
	Highlight     = color.New(color.FgHiBlue, color.Bold)             // Highlights key info
	DimText       = color.New(color.FgWhite, color.Faint)             // Subtle text
	InvertedStyle = color.New(color.FgBlack, color.BgWhite)           // Inverted text

	// ThemeType represents the available themes
	// "auto" = detect, "light" = light, "dark" = dark
	ThemeType = "auto"

	// AlwaysDarkTerminals is a list of terminals that are always dark
	AlwaysDarkTerminals = []string{"ghost"}
)

// SetTheme sets the current theme ("light", "dark", or "auto")
func SetTheme(theme string) {
	ThemeType = theme
}

// GetTheme returns the current theme, auto-detecting if needed
func GetTheme() string {
	if ThemeType == "auto" {
		if isAlwaysDarkTerminal() {
			return "dark"
		}
		// You can add more detection logic here (env vars, etc.)
		return "light" // default to light if not detected
	}
	return ThemeType
}

// isAlwaysDarkTerminal checks if the terminal is always dark (e.g., Ghost)
func isAlwaysDarkTerminal() bool {
	term := os.Getenv("TERM_PROGRAM")
	if term == "" {
		term = os.Getenv("TERM")
	}
	term = strings.ToLower(term)
	for _, t := range AlwaysDarkTerminals {
		if strings.Contains(term, t) {
			return true
		}
	}
	return false
}

// Dynamic style getters (use these in new code)
func GetHeaderStyle() *color.Color {
	if GetTheme() == "dark" {
		return color.New(color.FgHiGreen, color.Bold)
	}
	return color.New(color.FgGreen, color.Bold)
}

func GetErrorStyle() *color.Color {
	if GetTheme() == "dark" {
		return color.New(color.FgHiRed, color.Bold)
	}
	return color.New(color.FgRed, color.Bold, color.BgBlack)
}

func GetSuccessStyle() *color.Color {
	if GetTheme() == "dark" {
		return color.New(color.FgGreen, color.Bold)
	}
	return color.New(color.FgHiGreen, color.Bold)
}

func GetWarningStyle() *color.Color {
	if GetTheme() == "dark" {
		return color.New(color.FgHiYellow, color.Bold)
	}
	return color.New(color.FgYellow, color.Bold)
}

func GetInfoStyle() *color.Color {
	if GetTheme() == "dark" {
		return color.New(color.FgHiCyan, color.Bold)
	}
	return color.New(color.FgCyan, color.Bold)
}

func GetHighlightStyle() *color.Color {
	if GetTheme() == "dark" {
		return color.New(color.FgHiBlue, color.Bold)
	}
	return color.New(color.FgHiBlue, color.Bold)
}

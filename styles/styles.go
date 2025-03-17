package styles

import "github.com/fatih/color"

var (
	Cyan         = color.New(color.FgCyan)
	Green        = color.New(color.FgGreen)
	Red          = color.New(color.FgRed)
	Yellow       = color.New(color.FgYellow)
	Magenta      = color.New(color.FgMagenta)
	Bold         = color.New(color.Bold)
	HeaderStyle  = color.New(color.FgGreen, color.Bold)
	ErrorStyle   = color.New(color.FgRed, color.Bold, color.BgBlack)
	SuccessStyle = color.New(color.FgHiGreen, color.Bold)
)
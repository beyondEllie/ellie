package actions

import (
	"runtime"

	"github.com/tacheraSasi/ellie/styles"
)

// Displays the size of directories and files
func Size(){
	switch runtime.GOOS {
	case "darwin":
		SizeMac()
	default:
		styles.ErrorStyle.Println("Unknown OS") //TODO: will updated this to a better message later
	}
}

func SizeMac(){

}
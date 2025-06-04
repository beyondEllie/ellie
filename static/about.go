package static

import "fmt"


func GetAbout() string {
	readme,err := GetStaticFile("ABOUT.md")
	if err != nil {
		// log.Println("Error loading ABOUT.md:", err)
		return "Error loading about content."
	}
	fmt.Print(readme)
	return readme
}
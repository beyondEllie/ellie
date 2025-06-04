package static


func GetReadme() string {
	readme,err := GetStaticFile("README.md")
	if err != nil {
		return "Error loading instructions"
	}
	return readme
}
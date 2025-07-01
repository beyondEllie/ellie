package static

func Icon()any{
	icon, err := GetStaticFile("icon.png")
	if err != nil {
		panic(err)
	}
	return icon
}
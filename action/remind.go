package actions

import (
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

func Remind(){
	styles.Cyan.Print("ellie remind")
	title, err := getTitle()
	if err != nil {
		utils.Error("Something when wron failed to get the title.")
		return 
	}
}

func getTitle()(string,error){
	for {
		subject, err := utils.GetInput("What do you want to remind yourself?")
		if err == nil && subject != "" {
			return subject,nil
		}
		styles.ErrorStyle.Println("🚫 Title cannot be empty")
	}
}
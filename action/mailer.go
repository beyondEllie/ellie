package actions

import (
	// "bufio"
	// "os"

	"fmt"

	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

func Mailer(){
	
	styles.Cyan.Print("Send an email")
	styles.DimText.Println(" (powered by ekilirelay) ")
	
	subject := getSubject()
	fmt.Println("The subject entered",subject)
	
	
}

func getSubject() string{
	subject,err := utils.GetInput("Enter the subject")
	
	if err != nil{
		styles.ErrorStyle.Println("Error:",err)
	}
	
	if subject == ""{
		styles.ErrorStyle.Println("🚫 Subject can not be empty")
	}
	return subject
}
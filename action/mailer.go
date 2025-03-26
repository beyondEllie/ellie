package actions

import (
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

	for {
		if subject != ""{
			return subject
		}
		styles.ErrorStyle.Println("🚫 Subject can not be empty")
	}
	
}

func getEmail() string{
	recipientEmail,err := utils.GetInput("Enter the recipient email")
	
	if err != nil{
		styles.ErrorStyle.Println("Error:",err)
	}

	for {
		if recipientEmail != ""{
			return recipientEmail
		}
		styles.ErrorStyle.Println("🚫 recipient email can not be empty")
	}
}

func getMessage()string{
	message,err := utils.GetInput("Enter the message")
	
	if err != nil{
		styles.ErrorStyle.Println("Error:",err)
	}

	for {
		if message != ""{
			return message
		}
		styles.ErrorStyle.Println("🚫 message can not be empty")
	}

}


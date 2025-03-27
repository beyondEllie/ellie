package actions

import (
	"fmt"

	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

type EmailRequest struct {
	APIKey  string `json:"apikey"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
	Headers string `json:"headers"`
}

func Mailer(){
	
	styles.Cyan.Print("Send an email")
	styles.DimText.Println(" (powered by ekilirelay) ")
	
	subject := getSubject()
	fmt.Println("The subject entered",subject)
	
	
}

func getSubject() string {
	subject, err := utils.GetInput("Enter the subject")
	if err != nil {
		styles.ErrorStyle.Println("Error:", err)
	}
	for subject == "" {
		styles.ErrorStyle.Println("🚫 Subject cannot be empty")
		subject, _ = utils.GetInput("Enter the subject")
	}
	return subject
}

func getEmail() string {
	email, err := utils.GetInput("Enter the recipient email")
	if err != nil {
		styles.ErrorStyle.Println("Error:", err)
	}
	for email == "" {
		styles.ErrorStyle.Println("🚫 Recipient email cannot be empty")
		email, _ = utils.GetInput("Enter the recipient email")
	}
	return email
}

func getMessage() string {
	message, err := utils.GetInput("Enter the message")
	if err != nil {
		styles.ErrorStyle.Println("Error:", err)
	}
	for message == "" {
		styles.ErrorStyle.Println("🚫 Message cannot be empty")
		message, _ = utils.GetInput("Enter the message")
	}
	return message
}
package actions

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/tacheraSasi/ellie/configs"
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

func Mailer() {
	styles.Cyan.Print("Send an email")
	styles.DimText.Println(" (powered by ekilirelay) ")

	apiKey := getAPIKey()
	if apiKey == "" {
		return
	}

	to := getEmail()
	subject := getSubject()
	message := getMessage()
	headers := "From: Ellie Mailer"

	requestBody := EmailRequest{
		APIKey:  apiKey,
		To:      to,
		Subject: subject,
		Message: message,
		Headers: headers,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		styles.ErrorStyle.Println("Error encoding JSON:", err)
		return
	}

	resp, err := http.Post("https://relay.ekilie.com/api/index.php", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		styles.ErrorStyle.Println("Error sending email:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		styles.SuccessStyle.Println("✅ Email sent successfully!")
	} else {
		styles.ErrorStyle.Println("🚫 Failed to send email. Status:", resp.Status)
	}
}

func getSubject() string {
	for {
		subject, err := utils.GetInput("Enter the subject")
		if err == nil && subject != "" {
			return subject
		}
		styles.ErrorStyle.Println("🚫 Subject cannot be empty")
	}
}

func getEmail() string {
	for {
		email, err := utils.GetInput("Enter the recipient email")
		if err == nil && email != "" {
			return email
		}
		styles.ErrorStyle.Println("🚫 Recipient email cannot be empty")
	}
}

func getMessage() string {
	for {
		message, err := utils.GetInput("Enter the message")
		if err == nil && message != "" {
			return message 
		} 
		styles.ErrorStyle.Println("🚫 Message cannot be empty")
	}
}

func getAPIKey() string {
	apiKey := configs.GetEnv("API_KEY")
	if apiKey == "" {
		styles.ErrorStyle.Println("❌ Missing required configuration: API_KEY")
		styles.DimText.Println("💡 Run 'ellie config' to reconfigure (Get the API key at https://relay.ekilie.com/console )")
	}
	return apiKey
}

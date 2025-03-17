package configs

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

// Default configuration directory & file
var configDir string
var configPath string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		styles.ErrorStyle.Println("❌ Error: Unable to determine user home directory")
		os.Exit(1)
	}
	
	configDir = filepath.Join(homeDir, "ellie")         // Set directory
	configPath = filepath.Join(configDir, ".ellie.env") // Set full .env path

	// Ensure the directory exists
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		styles.ErrorStyle.Println("❌ Error: Failed to create config directory:", err)
		os.Exit(1)
	}

	// Initialize config
	Init()
}

// Init loads config or creates it if missing
func Init() {
	// If .env doesn't exist, create it
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		createDefaultConfig()
	}

	// Load .env file
	if err := godotenv.Load(configPath); err != nil {
		styles.WarningStyle.Println("⚠️ Warning: Failed to load config file:", err)
	}
}

// createDefaultConfig asks for user input & writes .env file
func createDefaultConfig() {
	styles.InfoStyle.Println("🔧 Setting up Ellie CLI configuration...")

	username, err := utils.GetInput("-> Enter your username: ")
	if err != nil {
		styles.ErrorStyle.Printf("❌ Error: %v\n", err)
		return
	}

	openaiKey, err := utils.GetInput("-> Enter your OpenAI API key: ")
	if err != nil {
		styles.ErrorStyle.Printf("❌ Error: %v\n", err)
		return
	}

	email, err := utils.GetInput("-> Enter your Email (Optional): ")
	if err != nil {
		styles.ErrorStyle.Printf("❌ Error: %v\n", err)
		return
	}

	// Save config
	envData := map[string]string{
		"USERNAME":       username,
		"EMAIL":       email,
		"OPENAI_API_KEY": openaiKey,
	}

	// Ensure .env file is written correctly
	err = godotenv.Write(envData, configPath)
	if err != nil {
		styles.ErrorStyle.Println("❌ Error: Failed to create config file:", err)
		os.Exit(1)
	}

	styles.SuccessStyle.Println("✅ Configuration saved successfully at", configPath)
	styles.InfoStyle.Println("🔧 Want to edit it? Open:", configPath)
}

// GetEnv fetches environment variables from the .env file
func GetEnv(key string) string {
	return os.Getenv(key)
}

// GetConfigPath returns the config file path
func GetConfigPath() string {
	return configPath
}

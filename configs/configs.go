package configs

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

// Default configuration file path in the user's home directory
var configPath string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		styles.ErrorStyle.Println("Error: Unable to determine user home directory")
		os.Exit(1)
	}
	configPath = filepath.Join(homeDir, "ellie/.ellie.env")

	// Initializes configuration
	Init()
}

// Init checks if the config file exists; if not, prompts user input and creates one
func Init() {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		createDefaultConfig()
	}

	// Load .env file
	if err := godotenv.Load(configPath); err != nil {
		styles.WarningStyle.Println("Warning: Failed to load config file:", err)
	}
}

// createDefaultConfig prompts user input and writes a .env file
func createDefaultConfig() {
	styles.InfoStyle.Println("🔧 Setting up Ellie CLI configuration...")

	username, err := utils.GetInput("-> Enter your username: ")
	if err != nil {
		styles.ErrorStyle.Printf("Error: %v\n", err)
		return
	}

	openaiKey, err := utils.GetInput("-> Enter your OpenAI API key: ")
	if err != nil {
		styles.ErrorStyle.Printf("Error: %v\n", err)
		return
	}

	pat, err := utils.GetInput("-> Enter your PAT (Personal Access Token): ")
	if err != nil {
		styles.ErrorStyle.Printf("Error: %v\n", err)
		return
	}

	// Save to .env file
	envData := map[string]string{
		"USERNAME":       username,
		"OPENAI_API_KEY": openaiKey,
		"PAT":            pat,
	}

	err = godotenv.Write(envData, configPath)
	if err != nil {
		styles.ErrorStyle.Println("Error: Failed to create config file:", err)
		os.Exit(1)
	}

	styles.SuccessStyle.Println("✅ Configuration saved successfully at", configPath)
	styles.InfoStyle.Println("Wanna edit the configurations?, Open ", configPath)
}

// GetEnv fetches environment variables from the config file
func GetEnv(key string) string {
	return os.Getenv(key)
}

// GetConfigPath returns the config file path
func GetConfigPath() string {
	return configPath
}

package configs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Default configuration file path in the user's home directory
var configPath string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error: Unable to determine user home directory")
		os.Exit(1)
	}
	configPath = filepath.Join(homeDir, ".ellie.env")

	// Initialize configuration
	Init()
}

// Init checks if the config file exists; if not, creates one with defaults
func Init() {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		createDefaultConfig()
	}

	// Load .env file
	if err := godotenv.Load(configPath); err != nil {
		fmt.Println("Warning: Failed to load config file", err)
	}
}

// createDefaultConfig writes a default .env file
func createDefaultConfig() {
	defaultConfig := `# Ellie CLI Configuration
OPENAI_API_KEY=
PAT=
USERNAME=
`
	err := os.WriteFile(configPath, []byte(defaultConfig), 0644)
	if err != nil {
		fmt.Println("Error: Failed to create default config file", err)
		os.Exit(1)
	}
	fmt.Println("Created default config at", configPath)
}

// GetEnv fetches environment variables from the config file
func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return ""
	}
	return value
}

// GetConfigPath returns the config file path
func GetConfigPath() string {
	return configPath
}

package actions

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/styles"
)

type Alias struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

var aliases []Alias

func init() {
	loadAliases()
}

func loadAliases() {
	aliasFile := filepath.Join(configs.GetEllieDir(), "aliases.json")
	if _, err := os.Stat(aliasFile); os.IsNotExist(err) {
		aliases = []Alias{}
		return
	}

	data, err := os.ReadFile(aliasFile)
	if err != nil {
		styles.ErrorStyle.Println("Error reading aliases:", err)
		return
	}

	if err := json.Unmarshal(data, &aliases); err != nil {
		styles.ErrorStyle.Println("Error parsing aliases:", err)
		aliases = []Alias{}
	}
}

func saveAliases() {
	aliasFile := filepath.Join(configs.GetEllieDir(), "aliases.json")
	data, err := json.MarshalIndent(aliases, "", "  ")
	if err != nil {
		styles.ErrorStyle.Println("Error saving aliases:", err)
		return
	}

	if err := os.WriteFile(aliasFile, data, 0644); err != nil {
		styles.ErrorStyle.Println("Error writing aliases:", err)
	}
}

func AliasAdd(args []string) {
	if len(args) < 2 {
		styles.ErrorStyle.Println("Usage: ellie alias add <name>=\"<command>\"")
		return
	}

	// Parse the alias definition
	parts := strings.SplitN(args[1], "=", 2)
	if len(parts) != 2 {
		styles.ErrorStyle.Println("Invalid alias format. Use: name=\"command\"")
		return
	}

	name := parts[0]
	command := strings.Trim(parts[1], "\"")

	// Check if alias already exists
	for i, a := range aliases {
		if a.Name == name {
			aliases[i].Command = command
			saveAliases()
			styles.SuccessStyle.Printf("Updated alias '%s'\n", name)
			return
		}
	}

	// Add new alias
	aliases = append(aliases, Alias{Name: name, Command: command})
	saveAliases()
	styles.SuccessStyle.Printf("Added alias '%s'\n", name)
}

func AliasList(args []string) {
	if len(aliases) == 0 {
		styles.InfoStyle.Println("No aliases defined")
		return
	}

	styles.InfoStyle.Println("Defined aliases:")
	for _, a := range aliases {
		fmt.Printf("  %s = %s\n", a.Name, a.Command)
	}
}

func AliasDelete(args []string) {
	if len(args) < 2 {
		styles.ErrorStyle.Println("Usage: ellie alias delete <name>")
		return
	}

	name := args[1]
	for i, a := range aliases {
		if a.Name == name {
			aliases = append(aliases[:i], aliases[i+1:]...)
			saveAliases()
			styles.SuccessStyle.Printf("Deleted alias '%s'\n", name)
			return
		}
	}

	styles.ErrorStyle.Printf("Alias '%s' not found\n", name)
}

func ExecuteAlias(name string) bool {
	for _, a := range aliases {
		if a.Name == name {
			// Split the command into parts
			parts := strings.Fields(a.Command)
			if len(parts) > 0 {
				// Execute the command
				// handleCommand(parts)
			}
			return true
		}
	}
	return false
}

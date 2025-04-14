package actions

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/styles"
)

type Project struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

var projects []Project

func init() {
	loadProjects()
}

func loadProjects() {
	projectFile := filepath.Join(configs.GetEllieDir(), "projects.json")
	if _, err := os.Stat(projectFile); os.IsNotExist(err) {
		projects = []Project{}
		return
	}

	data, err := os.ReadFile(projectFile)
	if err != nil {
		styles.ErrorStyle.Println("Error reading projects:", err)
		return
	}

	if err := json.Unmarshal(data, &projects); err != nil {
		styles.ErrorStyle.Println("Error parsing projects:", err)
		projects = []Project{}
	}
}

func saveProjects() {
	projectFile := filepath.Join(configs.GetEllieDir(), "projects.json")
	data, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		styles.ErrorStyle.Println("Error saving projects:", err)
		return
	}

	if err := os.WriteFile(projectFile, data, 0644); err != nil {
		styles.ErrorStyle.Println("Error writing projects:", err)
	}
}

func ProjectAdd(args []string) {
	if len(args) < 3 {
		styles.ErrorStyle.Println("Usage: ellie project add <name> <path>")
		return
	}

	name := args[1]
	path := args[2]

	// Convert path to absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		styles.ErrorStyle.Printf("Invalid path: %s\n", err)
		return
	}

	// Check if project already exists
	for i, p := range projects {
		if p.Name == name {
			projects[i].Path = absPath
			saveProjects()
			styles.SuccessStyle.Printf("Updated project '%s'\n", name)
			return
		}
	}

	// Add new project
	projects = append(projects, Project{Name: name, Path: absPath})
	saveProjects()
	styles.SuccessStyle.Printf("Added project '%s'\n", name)
}

func ProjectList(args []string) {
	if len(projects) == 0 {
		styles.InfoStyle.Println("No projects defined")
		return
	}

	styles.InfoStyle.Println("Defined projects:")
	for _, p := range projects {
		fmt.Printf("  %s -> %s\n", p.Name, p.Path)
	}
}

func ProjectDelete(args []string) {
	if len(args) < 2 {
		styles.ErrorStyle.Println("Usage: ellie project delete <name>")
		return
	}

	name := args[1]
	for i, p := range projects {
		if p.Name == name {
			projects = append(projects[:i], projects[i+1:]...)
			saveProjects()
			styles.SuccessStyle.Printf("Deleted project '%s'\n", name)
			return
		}
	}

	styles.ErrorStyle.Printf("Project '%s' not found\n", name)
}

func ProjectSwitch(args []string) {
	if len(args) < 2 {
		styles.ErrorStyle.Println("Usage: ellie switch <project-name>")
		return
	}

	name := args[1]
	for _, p := range projects {
		if p.Name == name {
			if err := os.Chdir(p.Path); err != nil {
				styles.ErrorStyle.Printf("Failed to switch to project '%s': %s\n", name, err)
				return
			}
			styles.SuccessStyle.Printf("Switched to project '%s'\n", name)
			return
		}
	}

	styles.ErrorStyle.Printf("Project '%s' not found\n", name)
}

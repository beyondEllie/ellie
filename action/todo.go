package actions

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/styles"
)

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "🔵 Low"
	case Medium:
		return "🟡 Medium"
	case High:
		return "🔴 High"
	default:
		return "⚪ Unknown"
	}
}

type Todo struct {
	ID          int       `json:"id"`
	Task        string    `json:"task"`
	Category    string    `json:"category"`
	Priority    Priority  `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
	Completed   bool      `json:"completed"`
}

var todos []Todo

func init() {
	loadTodos()
}

func loadTodos() {
	todoFile := filepath.Join(configs.GetEllieDir(), "todos.json")
	if _, err := os.Stat(todoFile); os.IsNotExist(err) {
		todos = []Todo{}
		return
	}

	data, err := os.ReadFile(todoFile)
	if err != nil {
		styles.ErrorStyle.Println("Error reading todos:", err)
		return
	}

	if err := json.Unmarshal(data, &todos); err != nil {
		styles.ErrorStyle.Println("Error parsing todos:", err)
		todos = []Todo{}
	}
}

func saveTodos() {
	todoFile := filepath.Join(configs.GetEllieDir(), "todos.json")
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		styles.ErrorStyle.Println("Error saving todos:", err)
		return
	}

	if err := os.WriteFile(todoFile, data, 0644); err != nil {
		styles.ErrorStyle.Println("Error writing todos:", err)
	}
}

func TodoAdd(args []string) {
	if len(args) < 2 {
		styles.ErrorStyle.Println("Usage: ellie todo add \"<task>\" [category] [priority]")
		styles.InfoStyle.Println("Priorities: low, medium, high (default: medium)")
		return
	}

	task := args[1]
	category := "general"
	priority := Medium

	if len(args) > 2 {
		category = args[2]
	}
	if len(args) > 3 {
		switch args[3] {
		case "low":
			priority = Low
		case "high":
			priority = High
		}
	}

	newID := 1
	if len(todos) > 0 {
		newID = todos[len(todos)-1].ID + 1
	}

	todo := Todo{
		ID:        newID,
		Task:      task,
		Category:  category,
		Priority:  priority,
		CreatedAt: time.Now(),
		Completed: false,
	}

	todos = append(todos, todo)
	saveTodos()
	styles.SuccessStyle.Printf("Added todo #%d: %s [%s] %s\n", todo.ID, todo.Task, todo.Category, todo.Priority)
}

func TodoList(args []string) {
	if len(todos) == 0 {
		styles.InfoStyle.Println("No todos found")
		return
	}

	// Sort todos by priority and creation date
	sort.Slice(todos, func(i, j int) bool {
		if todos[i].Priority != todos[j].Priority {
			return todos[i].Priority > todos[j].Priority
		}
		return todos[i].CreatedAt.Before(todos[j].CreatedAt)
	})

	// Group by category
	categories := make(map[string][]Todo)
	for _, todo := range todos {
		categories[todo.Category] = append(categories[todo.Category], todo)
	}

	styles.InfoStyle.Println("Your todos:")
	for category, categoryTodos := range categories {
		fmt.Printf("\n📁 %s:\n", category)
		for _, todo := range categoryTodos {
			status := "❌"
			if todo.Completed {
				status = "✅"
			}
			fmt.Printf("  %s #%d: %s %s\n", status, todo.ID, todo.Task, todo.Priority)
		}
	}
}

func TodoComplete(args []string) {
	if len(args) < 2 {
		styles.ErrorStyle.Println("Usage: ellie todo complete <id>")
		return
	}

	id := 0
	fmt.Sscanf(args[1], "%d", &id)

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true
			todos[i].CompletedAt = time.Now()
			saveTodos()
			styles.SuccessStyle.Printf("Completed todo #%d: %s\n", todo.ID, todo.Task)
			return
		}
	}

	styles.ErrorStyle.Printf("Todo #%d not found\n", id)
}

func TodoDelete(args []string) {
	if len(args) < 2 {
		styles.ErrorStyle.Println("Usage: ellie todo delete <id>")
		return
	}

	id := 0
	fmt.Sscanf(args[1], "%d", &id)

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			saveTodos()
			styles.SuccessStyle.Printf("Deleted todo #%d: %s\n", todo.ID, todo.Task)
			return
		}
	}

	styles.ErrorStyle.Printf("Todo #%d not found\n", id)
}

func TodoEdit(args []string) {
	if len(args) < 3 {
		styles.ErrorStyle.Println("Usage: ellie todo edit <id> <field> <value>")
		styles.InfoStyle.Println("Fields: task, category, priority")
		return
	}

	id := 0
	fmt.Sscanf(args[1], "%d", &id)

	for i, todo := range todos {
		if todo.ID == id {
			field := args[2]
			value := args[3]

			switch field {
			case "task":
				todos[i].Task = value
			case "category":
				todos[i].Category = value
			case "priority":
				switch value {
				case "low":
					todos[i].Priority = Low
				case "medium":
					todos[i].Priority = Medium
				case "high":
					todos[i].Priority = High
				default:
					styles.ErrorStyle.Println("Invalid priority. Use: low, medium, high")
					return
				}
			default:
				styles.ErrorStyle.Println("Invalid field. Use: task, category, priority")
				return
			}

			saveTodos()
			styles.SuccessStyle.Printf("Updated todo #%d\n", id)
			return
		}
	}

	styles.ErrorStyle.Printf("Todo #%d not found\n", id)
}

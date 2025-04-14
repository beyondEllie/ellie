package actions

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/styles"
)

type Todo struct {
	ID          int       `json:"id"`
	Task        string    `json:"task"`
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
		styles.ErrorStyle.Println("Usage: ellie todo add \"<task>\"")
		return
	}

	task := args[1]
	newID := 1
	if len(todos) > 0 {
		newID = todos[len(todos)-1].ID + 1
	}

	todo := Todo{
		ID:        newID,
		Task:      task,
		CreatedAt: time.Now(),
		Completed: false,
	}

	todos = append(todos, todo)
	saveTodos()
	styles.SuccessStyle.Printf("Added todo #%d: %s\n", todo.ID, todo.Task)
}

func TodoList(args []string) {
	if len(todos) == 0 {
		styles.InfoStyle.Println("No todos found")
		return
	}

	styles.InfoStyle.Println("Your todos:")
	for _, todo := range todos {
		status := "❌"
		if todo.Completed {
			status = "✅"
		}
		fmt.Printf("%s #%d: %s\n", status, todo.ID, todo.Task)
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

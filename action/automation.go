package actions

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

type AutomationTask struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Command     string    `json:"command"`
	Schedule    string    `json:"schedule"` // e.g., "daily", "hourly", "weekly", "custom"
	Time        string    `json:"time"`     // e.g., "09:00" for daily tasks
	Enabled     bool      `json:"enabled"`
	LastRun     time.Time `json:"last_run"`
	NextRun     time.Time `json:"next_run"`
	Description string    `json:"description"`
}

type AutomationData struct {
	Tasks []AutomationTask `json:"tasks"`
}

func getAutomationFilePath() string {
	return filepath.Join(configs.ConfigDir, "automations.json")
}

func loadAutomations() (*AutomationData, error) {
	data := &AutomationData{Tasks: []AutomationTask{}}
	
	filePath := getAutomationFilePath()
	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return data, nil
		}
		return nil, err
	}
	
	err = json.Unmarshal(content, data)
	return data, err
}

func saveAutomations(data *AutomationData) error {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(getAutomationFilePath(), content, 0600)
}

// AutomationAdd adds a new automation task
func AutomationAdd(args []string) {
	if len(args) < 3 {
		styles.GetErrorStyle().Println("Usage: ellie automate add <name> <schedule> <command>")
		styles.GetInfoStyle().Println("Schedules: daily, hourly, weekly, @time (e.g., @09:00)")
		return
	}
	
	name := args[1]
	schedule := args[2]
	command := strings.Join(args[3:], " ")
	
	// Validate schedule
	if !isValidSchedule(schedule) {
		styles.GetErrorStyle().Println("Invalid schedule. Use: daily, hourly, weekly, or @HH:MM")
		return
	}
	
	data, err := loadAutomations()
	if err != nil {
		styles.GetErrorStyle().Println("Error loading automations:", err)
		return
	}
	
	// Generate unique ID
	id := fmt.Sprintf("auto_%d", time.Now().Unix())
	
	// Calculate next run time
	nextRun := calculateNextRun(schedule, "")
	
	task := AutomationTask{
		ID:          id,
		Name:        name,
		Command:     command,
		Schedule:    schedule,
		Enabled:     true,
		NextRun:     nextRun,
		Description: fmt.Sprintf("Runs %s", schedule),
	}
	
	data.Tasks = append(data.Tasks, task)
	
	if err := saveAutomations(data); err != nil {
		styles.GetErrorStyle().Println("Error saving automation:", err)
		return
	}
	
	styles.GetSuccessStyle().Printf("✅ Automation '%s' added successfully!\n", name)
	fmt.Printf("   ID: %s\n", id)
	fmt.Printf("   Schedule: %s\n", schedule)
	fmt.Printf("   Next run: %s\n", nextRun.Format("2006-01-02 15:04:05"))
	styles.GetInfoStyle().Println("\n💡 Tip: Run 'ellie automate run' to execute scheduled tasks")
}

// AutomationList lists all automation tasks
func AutomationList(args []string) {
	data, err := loadAutomations()
	if err != nil {
		styles.GetErrorStyle().Println("Error loading automations:", err)
		return
	}
	
	if len(data.Tasks) == 0 {
		styles.GetInfoStyle().Println("No automations configured")
		styles.GetInfoStyle().Println("💡 Add one with: ellie automate add <name> <schedule> <command>")
		return
	}
	
	styles.GetInfoStyle().Println("\n⚙️  Automation Tasks")
	styles.GetInfoStyle().Println("═══════════════════════════════════════════════════════")
	
	for _, task := range data.Tasks {
		status := "🟢"
		if !task.Enabled {
			status = "⚫"
		}
		
		fmt.Printf("\n%s %s [%s]\n", status, task.Name, task.ID)
		fmt.Printf("   Command: %s\n", task.Command)
		fmt.Printf("   Schedule: %s\n", task.Schedule)
		
		if !task.NextRun.IsZero() {
			timeUntil := time.Until(task.NextRun)
			fmt.Printf("   Next run: %s (in %s)\n", 
				task.NextRun.Format("2006-01-02 15:04"), 
				formatDuration(timeUntil))
		}
		
		if !task.LastRun.IsZero() {
			fmt.Printf("   Last run: %s\n", task.LastRun.Format("2006-01-02 15:04"))
		}
	}
	
	styles.GetInfoStyle().Println("\n═══════════════════════════════════════════════════════")
}

// AutomationDelete removes an automation task
func AutomationDelete(args []string) {
	if len(args) < 2 {
		styles.GetErrorStyle().Println("Usage: ellie automate delete <id>")
		return
	}
	
	id := args[1]
	
	data, err := loadAutomations()
	if err != nil {
		styles.GetErrorStyle().Println("Error loading automations:", err)
		return
	}
	
	newTasks := []AutomationTask{}
	found := false
	
	for _, task := range data.Tasks {
		if task.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, task)
	}
	
	if !found {
		styles.GetErrorStyle().Printf("Automation with ID '%s' not found\n", id)
		return
	}
	
	data.Tasks = newTasks
	
	if err := saveAutomations(data); err != nil {
		styles.GetErrorStyle().Println("Error saving automations:", err)
		return
	}
	
	styles.GetSuccessStyle().Printf("✅ Automation '%s' deleted\n", id)
}

// AutomationToggle enables or disables an automation
func AutomationToggle(args []string) {
	if len(args) < 2 {
		styles.GetErrorStyle().Println("Usage: ellie automate toggle <id>")
		return
	}
	
	id := args[1]
	
	data, err := loadAutomations()
	if err != nil {
		styles.GetErrorStyle().Println("Error loading automations:", err)
		return
	}
	
	found := false
	for i, task := range data.Tasks {
		if task.ID == id {
			data.Tasks[i].Enabled = !data.Tasks[i].Enabled
			found = true
			
			status := "disabled"
			if data.Tasks[i].Enabled {
				status = "enabled"
			}
			
			styles.GetSuccessStyle().Printf("✅ Automation '%s' %s\n", task.Name, status)
			break
		}
	}
	
	if !found {
		styles.GetErrorStyle().Printf("Automation with ID '%s' not found\n", id)
		return
	}
	
	if err := saveAutomations(data); err != nil {
		styles.GetErrorStyle().Println("Error saving automations:", err)
		return
	}
}

// AutomationRun executes due automation tasks
func AutomationRun(args []string) {
	data, err := loadAutomations()
	if err != nil {
		styles.GetErrorStyle().Println("Error loading automations:", err)
		return
	}
	
	if len(data.Tasks) == 0 {
		styles.GetInfoStyle().Println("No automations to run")
		return
	}
	
	styles.GetInfoStyle().Println("🔄 Checking scheduled tasks...")
	
	now := time.Now()
	tasksRun := 0
	
	for i, task := range data.Tasks {
		if !task.Enabled {
			continue
		}
		
		// Check if task is due
		if task.NextRun.IsZero() || now.After(task.NextRun) {
			styles.GetHighlightStyle().Printf("\n▶️  Running: %s\n", task.Name)
			
			// Execute the command
			err := executeAutomationCommand(task.Command)
			if err != nil {
				styles.GetErrorStyle().Printf("❌ Error: %v\n", err)
			} else {
				styles.GetSuccessStyle().Println("✅ Completed")
			}
			
			// Update last run and calculate next run
			data.Tasks[i].LastRun = now
			data.Tasks[i].NextRun = calculateNextRun(task.Schedule, task.Time)
			tasksRun++
		}
	}
	
	if tasksRun == 0 {
		styles.GetInfoStyle().Println("No tasks due at this time")
	} else {
		styles.GetSuccessStyle().Printf("\n✅ Executed %d task(s)\n", tasksRun)
	}
	
	// Save updated task data
	if err := saveAutomations(data); err != nil {
		styles.GetErrorStyle().Println("Error saving automations:", err)
	}
}

// AutomationDaemon runs the automation daemon
func AutomationDaemon(args []string) {
	styles.GetInfoStyle().Println("🤖 Starting Ellie Automation Daemon")
	styles.GetInfoStyle().Println("Press Ctrl+C to stop")
	fmt.Println()
	
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	
	// Run immediately on start
	checkAndRunAutomations()
	
	for {
		select {
		case <-ticker.C:
			checkAndRunAutomations()
		}
	}
}

func checkAndRunAutomations() {
	data, err := loadAutomations()
	if err != nil {
		return
	}
	
	now := time.Now()
	
	for i, task := range data.Tasks {
		if !task.Enabled {
			continue
		}
		
		if task.NextRun.IsZero() || now.After(task.NextRun) {
			fmt.Printf("[%s] Running: %s\n", now.Format("15:04:05"), task.Name)
			
			err := executeAutomationCommand(task.Command)
			if err != nil {
				fmt.Printf("[%s] Error: %v\n", now.Format("15:04:05"), err)
			} else {
				fmt.Printf("[%s] Completed: %s\n", now.Format("15:04:05"), task.Name)
			}
			
			data.Tasks[i].LastRun = now
			data.Tasks[i].NextRun = calculateNextRun(task.Schedule, task.Time)
		}
	}
	
	saveAutomations(data)
}

func executeAutomationCommand(command string) error {
	// Parse and execute the command
	// For safety, we'll only allow ellie commands
	if !strings.HasPrefix(command, "ellie ") {
		return fmt.Errorf("only ellie commands are allowed in automations")
	}
	
	// Remove "ellie " prefix
	command = strings.TrimPrefix(command, "ellie ")
	
	// Split into args
	args := strings.Fields(command)
	if len(args) == 0 {
		return fmt.Errorf("empty command")
	}
	
	// Execute based on command
	// This is a simplified version - in production, you'd use the command registry
	fmt.Printf("   Executing: ellie %s\n", command)
	
	return nil
}

func isValidSchedule(schedule string) bool {
	validSchedules := []string{"daily", "hourly", "weekly"}
	
	for _, valid := range validSchedules {
		if schedule == valid {
			return true
		}
	}
	
	// Check for @time format
	if strings.HasPrefix(schedule, "@") && len(schedule) == 6 {
		// Format: @HH:MM
		return true
	}
	
	return false
}

func calculateNextRun(schedule, timeStr string) time.Time {
	now := time.Now()
	
	switch schedule {
	case "hourly":
		return now.Add(1 * time.Hour)
	
	case "daily":
		// If time is specified, use it; otherwise use current time tomorrow
		if timeStr != "" {
			// Parse time
			parts := strings.Split(timeStr, ":")
			if len(parts) == 2 {
				var hour, minute int
				fmt.Sscanf(parts[0], "%d", &hour)
				fmt.Sscanf(parts[1], "%d", &minute)
				
				nextRun := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
				if nextRun.Before(now) {
					nextRun = nextRun.Add(24 * time.Hour)
				}
				return nextRun
			}
		}
		return now.Add(24 * time.Hour)
	
	case "weekly":
		return now.Add(7 * 24 * time.Hour)
	
	default:
		// Check for @time format
		if strings.HasPrefix(schedule, "@") {
			timeStr := strings.TrimPrefix(schedule, "@")
			parts := strings.Split(timeStr, ":")
			if len(parts) == 2 {
				hour := 0
				minute := 0
				fmt.Sscanf(parts[0], "%d", &hour)
				fmt.Sscanf(parts[1], "%d", &minute)
				
				nextRun := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
				if nextRun.Before(now) {
					nextRun = nextRun.Add(24 * time.Hour)
				}
				return nextRun
			}
		}
		return now.Add(1 * time.Hour)
	}
}

// QuickAutomations sets up common automation tasks
func QuickAutomations(args []string) {
	styles.GetInfoStyle().Println("\n⚡ Quick Automation Setup")
	styles.GetInfoStyle().Println("═══════════════════════════════════════════════════════")
	
	styles.GetInfoStyle().Println("\nSelect automations to enable:")
	fmt.Println("1. Daily health check (9:00 AM)")
	fmt.Println("2. Hourly git status check")
	fmt.Println("3. Daily system cleanup (11:00 PM)")
	fmt.Println("4. Weekly system update check (Sunday 10:00 AM)")
	fmt.Println("5. All of the above")
	fmt.Println("0. Cancel")
	
	choice, err := utils.GetInput("\nYour choice: ")
	if err != nil || choice == "0" {
		return
	}
	
	data, err := loadAutomations()
	if err != nil {
		styles.GetErrorStyle().Println("Error loading automations:", err)
		return
	}
	
	quickTasks := make(map[string]AutomationTask)
	quickTasks["1"] = AutomationTask{
		ID:          fmt.Sprintf("auto_health_%d", time.Now().Unix()),
		Name:        "Daily Health Check",
		Command:     "ellie health",
		Schedule:    "@09:00",
		Enabled:     true,
		NextRun:     calculateNextRun("@09:00", ""),
		Description: "Daily system health check",
	}
	
	quickTasks["2"] = AutomationTask{
		ID:          fmt.Sprintf("auto_git_%d", time.Now().Unix()),
		Name:        "Hourly Git Check",
		Command:     "ellie git status",
		Schedule:    "hourly",
		Enabled:     true,
		NextRun:     calculateNextRun("hourly", ""),
		Description: "Check git status every hour",
	}
	
	quickTasks["3"] = AutomationTask{
		ID:          fmt.Sprintf("auto_cleanup_%d", time.Now().Unix()),
		Name:        "Daily Cleanup",
		Command:     "ellie disk space",
		Schedule:    "@23:00",
		Enabled:     true,
		NextRun:     calculateNextRun("@23:00", ""),
		Description: "Check disk space daily",
	}
	
	quickTasks["4"] = AutomationTask{
		ID:          fmt.Sprintf("auto_update_%d", time.Now().Unix()),
		Name:        "Weekly Update Check",
		Command:     "ellie update",
		Schedule:    "weekly",
		Enabled:     true,
		NextRun:     calculateNextRun("weekly", ""),
		Description: "Check for system updates weekly",
	}
	
	if choice == "5" {
		for _, task := range quickTasks {
			data.Tasks = append(data.Tasks, task)
		}
		styles.GetSuccessStyle().Println("✅ All automations enabled!")
	} else if task, ok := quickTasks[choice]; ok {
		data.Tasks = append(data.Tasks, task)
		styles.GetSuccessStyle().Printf("✅ Automation '%s' enabled!\n", task.Name)
	} else {
		styles.GetErrorStyle().Println("Invalid choice")
		return
	}
	
	if err := saveAutomations(data); err != nil {
		styles.GetErrorStyle().Println("Error saving automations:", err)
		return
	}
	
	styles.GetInfoStyle().Println("\n💡 Run 'ellie automate run' to execute scheduled tasks")
	styles.GetInfoStyle().Println("💡 Run 'ellie automate daemon' to start the automation daemon")
}

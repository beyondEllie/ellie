package actions

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

// PomodoroSession represents a focus session
type PomodoroSession struct {
	WorkDuration            time.Duration
	BreakDuration           time.Duration
	LongBreakDuration       time.Duration
	SessionsBeforeLongBreak int
	CurrentSession          int
	TotalSessions           int
	IsActive                bool
	StartTime               time.Time
	EndTime                 time.Time
	Task                    string
}

// Focus implements the Pomodoro technique for productivity
func Focus(args []string) {
	styles.InfoStyle.Println("🍅 Activating Pomodoro Focus Mode...")

	// Parse arguments for custom durations
	workDuration := 25 * time.Minute
	breakDuration := 5 * time.Minute
	longBreakDuration := 15 * time.Minute
	sessionsBeforeLongBreak := 4

	if len(args) > 1 {
		// Allow custom work duration: focus 30 (30 minutes)
		if customWork, err := time.ParseDuration(args[1] + "m"); err == nil {
			workDuration = customWork
		}
	}

	// Get task description
	task, err := utils.GetInput("What are you working on? ")
	if err != nil {
		styles.ErrorStyle.Println("Error reading input:", err)
		return
	}

	if task == "" {
		task = "Focus Session"
	}

	session := &PomodoroSession{
		WorkDuration:            workDuration,
		BreakDuration:           breakDuration,
		LongBreakDuration:       longBreakDuration,
		SessionsBeforeLongBreak: sessionsBeforeLongBreak,
		CurrentSession:          1,
		TotalSessions:           0,
		IsActive:                true,
		StartTime:               time.Now(),
		Task:                    task,
	}

	styles.SuccessStyle.Printf("🎯 Starting Pomodoro session: %s\n", task)
	styles.InfoStyle.Printf("⏱️  Work duration: %s | Break duration: %s\n",
		formatDuration(workDuration), formatDuration(breakDuration))

	// Start the Pomodoro cycle
	runPomodoroCycle(session)
}

// runPomodoroCycle manages the complete Pomodoro workflow
func runPomodoroCycle(session *PomodoroSession) {
	for session.IsActive {
		// Work session
		styles.GetHighlightStyle().Printf("\n🚀 Starting Work Session %d\n", session.CurrentSession)
		styles.InfoStyle.Printf("📝 Task: %s\n", session.Task)

		if !runTimer(session.WorkDuration, "Work", session.Task) {
			break // User interrupted
		}

		session.TotalSessions++
		showWorkSessionComplete(session)

		// Check if it's time for a long break
		if session.CurrentSession%session.SessionsBeforeLongBreak == 0 {
			styles.GetHighlightStyle().Printf("\n🎉 Long Break Time! (%s)\n", formatDuration(session.LongBreakDuration))
			if !runTimer(session.LongBreakDuration, "Long Break", "Take a well-deserved long break!") {
				break
			}
		} else {
			// Regular break
			styles.GetHighlightStyle().Printf("\n☕ Break Time! (%s)\n", formatDuration(session.BreakDuration))
			if !runTimer(session.BreakDuration, "Break", "Take a short break!") {
				break
			}
		}

		// Ask if user wants to continue
		if !askContinue() {
			break
		}

		session.CurrentSession++
	}

	// Session ended
	session.EndTime = time.Now()
	session.IsActive = false
	showSessionSummary(session)
}

// runTimer runs a countdown timer with visual feedback
func runTimer(duration time.Duration, sessionType, message string) bool {
	endTime := time.Now().Add(duration)

	styles.InfoStyle.Printf("⏰ %s: %s\n", sessionType, message)
	styles.DimText.Println("Press Ctrl+C to stop early")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		remaining := endTime.Sub(time.Now())
		if remaining <= 0 {
			break
		}

		// Clear line and show progress
		fmt.Printf("\r%s %s remaining... ", sessionType, formatDuration(remaining))

		select {
		case <-ticker.C:
			continue
		case <-utils.GetInterruptChannel():
			fmt.Println() // New line after interrupt
			styles.WarningStyle.Println("⏹️  Session interrupted by user")
			return false
		}
	}

	fmt.Println() // New line after timer completes
	showTimerComplete(sessionType)
	return true
}

// showWorkSessionComplete displays completion message
func showWorkSessionComplete(session *PomodoroSession) {
	styles.SuccessStyle.Printf("✅ Work Session %d Complete!\n", session.CurrentSession)
	styles.InfoStyle.Printf("📊 Total sessions completed: %d\n", session.TotalSessions)

	// Play notification sound if available
	playNotificationSound()
}

// showTimerComplete displays timer completion with notification
func showTimerComplete(sessionType string) {
	styles.SuccessStyle.Printf("🔔 %s Complete!\n", sessionType)

	// Show system notification
	showSystemNotification(sessionType + " Complete!")

	// Play notification sound
	playNotificationSound()
}

// showSessionSummary displays final session statistics
func showSessionSummary(session *PomodoroSession) {
	duration := session.EndTime.Sub(session.StartTime)

	styles.GetHighlightStyle().Println("\n📈 Session Summary")
	styles.InfoStyle.Println("─────────────────")
	styles.InfoStyle.Printf("🎯 Task: %s\n", session.Task)
	styles.InfoStyle.Printf("⏱️  Total time: %s\n", formatDuration(duration))
	styles.InfoStyle.Printf("🚀 Work sessions: %d\n", session.TotalSessions)
	styles.InfoStyle.Printf("📅 Started: %s\n", session.StartTime.Format("15:04:05"))
	styles.InfoStyle.Printf("📅 Ended: %s\n", session.EndTime.Format("15:04:05"))

	// Calculate productivity metrics
	if session.TotalSessions > 0 {
		avgSessionTime := duration / time.Duration(session.TotalSessions)
		styles.InfoStyle.Printf("📊 Average session time: %s\n", formatDuration(avgSessionTime))
	}

	styles.SuccessStyle.Println("🎉 Great job! Keep up the productivity!")
}

// askContinue asks user if they want to continue with another session
func askContinue() bool {
	styles.InfoStyle.Print("Continue with another session? (y/n): ")
	var response string
	fmt.Scanln(&response)
	return strings.ToLower(strings.TrimSpace(response)) == "y" ||
		strings.ToLower(strings.TrimSpace(response)) == "yes"
}

// formatDuration formats a duration in a human-readable format
func formatDuration(d time.Duration) string {
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60

	if minutes >= 60 {
		hours := minutes / 60
		minutes = minutes % 60
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}

	if seconds > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	}

	return fmt.Sprintf("%dm", minutes)
}

// showSystemNotification displays a system notification
func showSystemNotification(message string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "Ellie Focus"`, message))
	case "linux":
		cmd = exec.Command("notify-send", "Ellie Focus", message)
	case "windows":
		cmd = exec.Command("powershell", "-Command", fmt.Sprintf(`New-BurntToastNotification -Text "%s" -Header "Ellie Focus"`, message))
	default:
		return // Skip notification on unsupported OS
	}

	// Run notification in background
	go func() {
		cmd.Run()
	}()
}

// playNotificationSound plays a notification sound
func playNotificationSound() {
	// Try to play a system sound or beep
	switch runtime.GOOS {
	case "darwin":
		exec.Command("afplay", "/System/Library/Sounds/Glass.aiff").Run()
	case "linux":
		exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga").Run()
	case "windows":
		// Windows doesn't have a simple command for this, but we can try
		exec.Command("powershell", "-Command", "[console]::beep(800,200)").Run()
	}
}

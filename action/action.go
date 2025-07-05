package actions

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

// Run executes system commands
func Run(args []string) {
	if len(args) < 2 {
		fmt.Println("Please specify a command to run")
		return
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("pwsh", "-Command", strings.Join(args[1:], " "))
	default:
		cmd = exec.Command(args[1], args[2:]...)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		styles.ErrorStyle.Printf("🚫 Error: %v\n", err)
		os.Exit(0)
		return
	}
	fmt.Printf("%s\n", output)
}

// Pwd prints working directory
func Pwd() {
	dir, err := os.Getwd()
	if err != nil {
		styles.ErrorStyle.Printf("🚫 Error: %v\n", err)
		return
	}
	fmt.Println(dir)
}

func GitSetup(pat, username string) {
	cmd := exec.Command("git", "status")
	output, err := cmd.CombinedOutput()
	if err != nil {
		styles.ErrorStyle.Printf("🚫 Error: %v\n", err)
		return
	}

	if len(output) > 0 {
		fmt.Printf("Output: %s\n", string(output))
	}
}

func ListFiles(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		styles.ErrorStyle.Println("Error reading directory:", err)
		return
	}
	fmt.Println("Files:")
	for _, file := range files {
		styles.Bold.Println("--", file.Name())
	}
}

func CreateFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		styles.ErrorStyle.Println("Error creating file:", err)
		return
	}
	file.Close()
	fmt.Printf("File %s created successfully.\n", filePath)
}

func NetworkStatus() {
	cmd := exec.Command("nmcli", "general", "status")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error checking network status:", err)
		return
	}
	fmt.Printf("Network Status:\n%s\n", string(output))
}

func ConnectWiFi(ssid, password string) {
	cmd := exec.Command("nmcli", "dev", "wifi", "connect", ssid, "password", password)
	output, err := cmd.CombinedOutput()
	if err != nil {
		styles.ErrorStyle.Printf("Error connecting to Wi-Fi %s: %s\n", ssid, err)
		return
	}
	fmt.Printf("Connected to Wi-Fi %s successfully:\n%s\n", ssid, string(output))
}

func StartApache() {
	styles.InfoStyle.Println("STARTING APACHE...")
	if err := controlService("apache2", "start"); err == nil {
		styles.SuccessStyle.Println("Apache server started successfully.")
	}
}

func StartMysql() {
	styles.InfoStyle.Println("STARTING MYSQL...")
	if err := controlService("mysql", "start"); err == nil {
		styles.SuccessStyle.Println("MySQL server started successfully.")
	}
}

func StartPostgres() {
	styles.InfoStyle.Println("STARTING POSTGRES...")
	if err := controlService("postgresql", "start"); err == nil {
		styles.SuccessStyle.Println("PostgreSQL server started successfully.")
	}
}

func StartAll() {
	StartApache()
	StartMysql()
	StartPostgres()
}

func StopApache() {
	styles.InfoStyle.Println("STOPPING APACHE...")
	if err := controlService("apache2", "stop"); err == nil {
		styles.SuccessStyle.Println("Apache server stopped successfully.")
	}
}

func StopMysql() {
	styles.InfoStyle.Println("STOPPING MYSQL...")
	if err := controlService("mysql", "stop"); err == nil {
		styles.SuccessStyle.Println("MySQL server stopped successfully.")
	}
}
func StopPostgres() {
	styles.InfoStyle.Println("STOPPING POSTGRES...")
	if err := controlService("postgresql", "stop"); err == nil {
		styles.SuccessStyle.Println("PostgreSQL server stopped successfully.")
	}
}

func StopAll() {
	StopApache()
	StopMysql()
	StopPostgres()
}

// SysInfo gets system information
func SysInfo() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("pwsh", "-Command", `
			Get-CimInstance Win32_OperatingSystem | Select-Object Caption, Version, OSArchitecture | Format-List;
			Get-ComputerInfo -Property 'OsTotalVisibleMemorySize', 'OsFreePhysicalMemory' | Format-List`)
	case "darwin":
		cmd = exec.Command("sh", "-c", `top -l 1 | head -n 10 && sysctl -n hw.memsize && df -h`)
	default: // Linux
		cmd = exec.Command("sh", "-c", `top -bn1 | grep load && free -m && df -h`)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		styles.ErrorStyle.Printf("Error getting system info: %v\n", err)
		return
	}
	fmt.Printf("System Info:\n%s\n", output)
}

// InstallPackage installs packages
func InstallPackage(pkg string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("winget", "install", pkg)
	case "darwin":
		cmd = exec.Command("brew", "install", pkg)
	default:
		cmd = exec.Command("sudo", "apt-get", "install", "-y", pkg)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		styles.ErrorStyle.Printf("Install error: %v\n", err)
		return
	}
	fmt.Printf("Installed %s:\n%s\n", pkg, output)
}

// UpdatePackages updates system packages
func UpdatePackages() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("winget", "update")
	case "darwin":
		cmd = exec.Command("brew", "update")
	default:
		cmd = exec.Command("sudo", "apt-get", "update")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		styles.ErrorStyle.Printf("Update error: %v\n", err)
		return
	}
	fmt.Printf("Updates:\n%s\n", output)
}

// Service control functions
func controlService(service, action string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("net", action, service)
	case "darwin":
		cmd = exec.Command("launchctl", action, service)
	default:
		cmd = exec.Command("sudo", "systemctl", action, service)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		styles.ErrorStyle.Printf("service control failed: %v\nOutput: %s", err, output)
		return fmt.Errorf("service control failed: %v\nOutput: %s", err, output)
	}
	fmt.Printf("Service %s %sed\nOutput: %s\n", service, action, output)
	return nil
}

// OpenExplorer opens file manager
func OpenExplorer(optionalPath ...string) {
	var path string = "."
	if len(optionalPath) > 0 {
		path = optionalPath[0]
	}
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", path)
	case "darwin":
		cmd = exec.Command("open", path)
	default:
		cmd = exec.Command("xdg-open", path)
	}

	if err := cmd.Start(); err != nil {
		styles.ErrorStyle.Printf("Error opening explorer: %v\n", err)
	}
}

func Play(args []string) {
	if len(args) < 2 {
		styles.ErrorStyle.Println("Please provide a file path to play.")
		return
	}

	audioPath := args[1]

	if runtime.GOOS == "linux" {
		// Try mpv first
		cmd := exec.Command("which", "mpv")
		if err := cmd.Run(); err == nil {
			command := []string{"mpv", audioPath}
			fmt.Println("Playing file with mpv...")
			utils.RunCommand(command, "Error playing the file:")
			return
		}
	}

	// Fallback to custom beep-based player
	fmt.Println("Playing file using Go beep...")
	f, err := os.Open(audioPath)
	if err != nil {
		styles.ErrorStyle.Printf("Error opening file: %v\n", err)
		return
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		styles.ErrorStyle.Printf("Error decoding file: %v\n", err)
		return
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

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

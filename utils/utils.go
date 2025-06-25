package utils

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/gen2brain/beeep"
	"github.com/tacheraSasi/ellie/styles"
)

// // Ads for random promotion messages
// var Ads []string = []string{
// 	"🚀 Boost your productivity with ekilie!",
// 	"🔥 Check out ekiliSense for smarter school management!",
// 	"💻 Need a project tracker? Try ekilie!",
// }

// RandNum generates a random number between 0 and 100.
func RandNum() int {
	return rand.IntN(100)
}

// RandNumRange generates a random number between min and max.
func RandNumRange(min, max int) int {
	return rand.IntN(max-min+1) + min
}

// IsEven checks if a number is even.
func IsEven(num int) bool {
	return num%2 == 0
}

// IsOdd checks if a number is odd.
func IsOdd(num int) bool {
	return num%2 != 0
}

// RunCommand executes a shell command and prints the output or error.
func RunCommand(cmdArgs []string, errMsg string) {
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s %s\n", errMsg, err)
		return
	}
	if len(output) > 0 {
		fmt.Printf("Output:\n%s\n", output)
	}
}

// IsLinux returns true if the OS is Linux.
func IsLinux() bool {
	return strings.Contains(runtime.GOOS, "linux")
}

// IsMac returns true if the OS is macOS.
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// IsWindows returns true if the OS is Windows.
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// RenderMarkdown renders Markdown input using Glamour.
func RenderMarkdown(input string) (string, error) {
	rendered, err := glamour.Render(input, "dark")
	if err != nil {
		return "", err
	}
	return rendered, nil
}

// Exists checks if a file or directory exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// ReadFile reads a file and returns its content as a string.
func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteFile writes a string to a file.
func WriteFile(filePath, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}

// AppendToFile appends content to a file.
func AppendToFile(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content + "\n")
	return err
}

// CurrentTimestamp returns the current timestamp as a formatted string.
func CurrentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Sleep pauses execution for a given number of seconds.
func Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

// ClearScreen clears the console screen based on the OS.
func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "linux", "darwin":
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// GetRandomAd returns a random promotional message from Ads.
// func GetRandomAd() string {
// 	return Ads[rand.IntN(len(Ads))]
// }

func IsErr(err error, msg string) bool {
	if err != nil {
		styles.ErrorStyle.Println(msg, err)
		return true
	}
	return false
}
func IsErrFatal(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}
func IsErrFatalWithMsg(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}

func GetOS() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "mac"
	case "linux":
		return "linux"
	case "windows":
		return "windows"
	default:
		return "unknown"
	}
}

func AskYesNo(question string, defaultYes bool) bool {
	options := "[Y/n]"
	if !defaultYes {
		options = "[y/N]"
	}

	styles.InfoStyle.Printf("%s %s ", question, options)
	var response string
	fmt.Scanln(&response)

	response = strings.ToLower(strings.TrimSpace(response))
	if response == "" {
		return defaultYes
	}
	return response == "y" || response == "yes"
}

// Shows a loading spinner in the terminal
func ShowLoadingSpinner(message string, done chan bool) {
	spinner := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	i := 0
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			fmt.Print("\r\033[K") // Clear the current line
			return
		case <-ticker.C:
			fmt.Printf("\r%s %s", message, spinner[i%len(spinner)])
			i++
		}
	}
}

// ShowLoadingSpinnerWithMessage shows a loading spinner with a custom message.
func ShowLoadingSpinnerWithMessage(message string) {
	styles.InfoStyle.Printf("%s ", message)
	spinner := []string{"|", "/", "-", "\\"}
	for i := 0; i < 10; i++ {
		fmt.Printf("\r%s %s", message, spinner[i%len(spinner)])
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("\r" + message + " done!")
}

// ShowLoadingSpinnerWithMessageAndDuration shows a loading spinner with a custom message and duration.
func ShowLoadingSpinnerWithMessageAndDuration(message string, duration time.Duration) {
	styles.InfoStyle.Printf("%s ", message)
	spinner := []string{"|", "/", "-", "\\"}
	endTime := time.Now().Add(duration)
	for time.Now().Before(endTime) {
		fmt.Printf("\r%s %s", message, spinner[time.Now().UnixNano()/100000000%int64(len(spinner))])
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("\r" + message + " done!")
}

// Error util for ellie
func Error(message string, details ...any) {
	styles.ErrorStyle.Printf("%s", message)
}

// Sends a desktop Notification
func Notify(message string) {
	err := beeep.Notify("🔔 Ellie Reminder", message, "")

	if err != nil {
		Error("❌ Failed to send notification: " + err.Error())
	} else {
		styles.InfoStyle.Printf("\n🔔 Reminder: %s\n", message)
	}
}

// Schedules a native reminder using the 'at' command.
func ScheduleNativeReminder(title string, durationMinutes int) {
	cmd := fmt.Sprintf(`echo "notify-send 'Ellie Reminder' '%s'" | at now + %d minutes`, title, durationMinutes)
	exec.Command("bash", "-c", cmd).Run()
}

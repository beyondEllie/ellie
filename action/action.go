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
	"github.com/tacheraSasi/ellie/elliecore"
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
	var output string

	switch runtime.GOOS {
	case "windows":
		// Windows: use ipconfig
		output = elliecore.RunCmd("ipconfig")
	case "darwin":
		// macOS: use networksetup and ifconfig
		output = elliecore.RunCmd("networksetup -listallnetworkservices")
		output += "\n\n" + elliecore.RunCmd("ifconfig | grep -E 'inet |status'")
	case "linux":
		// Linux: try nmcli first, fallback to ip
		output = elliecore.RunCmd("nmcli general status")
		if strings.Contains(output, "Error:") {
			output = elliecore.RunCmd("ip addr show")
		}
	default:
		// Fallback for other systems
		output = elliecore.RunCmd("ifconfig")
		if strings.Contains(output, "Error:") {
			output = elliecore.RunCmd("ip addr show")
		}
	}

	fmt.Printf("Network Status:\n%s\n", output)
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

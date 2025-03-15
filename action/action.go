package actions

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/tacheraSasi/ellie/utils"
)

func Run(args []string) {
	cmd := exec.Command(args[2], args[3:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if output != nil || len(output) == 0 {
		fmt.Printf("%s", output)
	}

}
func Pwd() {
	cmd := exec.Command("pwd")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	if output != nil || len(output) == 0 {
		fmt.Printf("Output: %s", output)
	}
}

func GitSetup(pat, username string) {
	cmd := exec.Command("git", "status")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	if len(output) > 0 {
		fmt.Printf("Output: %s\n", string(output))
	}
}



func ListFiles(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	fmt.Println("Files:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func CreateFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
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
		fmt.Printf("Error connecting to Wi-Fi %s: %s\n", ssid, err)
		return
	}
	fmt.Printf("Connected to Wi-Fi %s successfully:\n%s\n", ssid, string(output))
}

func StartApache() {
	fmt.Println("STARTING APACHE...")
	if err := controlService("apache2", "start"); err == nil {
		fmt.Println("Apache server started successfully.")
	}
}

func StartMysql() {
	fmt.Println("STARTING MYSQL...")
	if err := controlService("mysql", "start"); err == nil {
		fmt.Println("MySQL server started successfully.")
	}
}

func StartAll() {
	StartApache()
	StartMysql()
}

func StopApache() {
	fmt.Println("STOPPING APACHE...")
	if err := controlService("apache2", "stop"); err == nil {
		fmt.Println("Apache server stopped successfully.")
	}
}

func StopMysql() {
	fmt.Println("STOPPING MYSQL...")
	if err := controlService("mysql", "stop"); err == nil {
		fmt.Println("MySQL server stopped successfully.")
	}
}

func StopAll() {
	StopApache()
	StopMysql()
}

// SysInfo gets system information cross-platform
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
		fmt.Printf("Error getting system info: %v\n", err)
		return
	}
	fmt.Printf("System Info:\n%s\n", output)
}

// InstallPackage installs packages cross-platform
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
		fmt.Printf("Install error: %v\n", err)
		return
	}
	fmt.Printf("Installed %s:\n%s\n", pkg, output)
}

// UpdatePackages updates system packages cross-platform
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
		fmt.Printf("Update error: %v\n", err)
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
		return fmt.Errorf("service control failed: %v\nOutput: %s", err, output)
	}
	fmt.Printf("Service %s %sed\nOutput: %s\n", service, action, output)
	return nil
}

// OpenExplorer opens file manager cross-platform
func OpenExplorer() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", ".")
	case "darwin":
		cmd = exec.Command("open", ".")
	default:
		cmd = exec.Command("xdg-open", ".")
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error opening explorer: %v\n", err)
	}
}

func Play(args []string) {
	if runtime.GOOS != "linux"{
		fmt.Println("Play functinality is only supported on Linux for now.")
		return
	}

	// fmt.Println(args) //For debugging
	command := []string{"mpv",args[1]}//TODO: check if mpv is installed
	//TODO:Will create a custom way of playing files in the future
	fmt.Println("Playing file...")
	utils.RunCommand(command,"Error playing the file:")
}

func Focus(args []string){//Doesnot work properly
	fmt.Println(args)
	var cmd *exec.Cmd;
	if len(args) < 2{
		cmd = exec.Command(args[0])
	}else{
		cmd = exec.Command(args[0], args[1:]...)
	}
	output,err := cmd.CombinedOutput()
	if err != nil{
		fmt.Println("Error: ",err)
		return
	}
	if output != nil || len(output) == 0{
		fmt.Printf("%s",output)
	}

}

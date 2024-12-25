package action

import (
	"fmt"
	"os/exec"
)

func controlService(service, action string) error {
	cmd := exec.Command("pkexec", "systemctl", action, service)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}

	if output != nil || len(output) == 0 {
		fmt.Printf("Output:\n%s\n", output)
	}
	return nil
}

func Pwd(){
	cmd := exec.Command("pwd")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	if output != nil || len(output) == 0 {
		fmt.Printf("Output: %s", output)
	}
}

func GitSetup(pat,username string) {
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
	fmt.Println("STARTING ALL SERVICES...")
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
	fmt.Println("STOPPING ALL SERVICES...")
	StopApache()
	StopMysql()
}

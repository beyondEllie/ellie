package action

import (
	"fmt"
	"os/exec"
)

func controlService(service, action string) {
	cmd := exec.Command(fmt.Sprintf("pkexec systemctl %s %s", action, service))

	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	
	if output != nil{
		fmt.Printf("Output:\n%s\n", output)
	}

}
func StartApache() {
	fmt.Println("STARTING APACHE...")
	controlService("apache", "start")
}

func StartMysql() {

}

func StartAll() {

}

func Default() {

}

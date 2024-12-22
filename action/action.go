package action

import (
	"fmt"
	"os/exec"
)

func controlService(service, action string) {
	cmd := exec.Command(fmt.Sprintf("pkexec systemctl %s %s", action, service))
	
	output,err := cmd.Output()
	
	
}
func StartApache() {

}

func StartMysql() {

}

func StartAll() {

}

func Default() {

}

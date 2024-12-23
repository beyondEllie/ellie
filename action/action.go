package action

import (
	"fmt"
	"os/exec"
)

func controlService(service, action string) error {
	cmd := exec.Command(fmt.Sprintf("pkexec systemctl %s %s", action, service))

	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}
	
	if output != nil{
		fmt.Printf("Output:\n%s\n", output)
	}
	return nil
}
func StartApache() {
	fmt.Println("STARTING APACHE...")
	if err := controlService("mysql", "start"); err==nil{
		fmt.Println("Mysql server started successfully")
	}
	
}

func StartMysql() {
	fmt.Println("STARTING MYSQL...")
	if err := controlService("apache", "start"); err==nil{
		fmt.Println("Apache server started successfully")
	}
}

func StartAll() {

}

func Default() {

}

package actions

import (
	"os"
	"os/exec"

	"github.com/tacheraSasi/ellie/styles"
)

func DockerBuild(args []string) {
	cmdArgs := append([]string{"build"}, args...)
	cmd := exec.Command("docker", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		styles.GetErrorStyle().Println("Error building Docker image:", err)
	}
}

func DockerRun(args []string) {
	cmdArgs := append([]string{"run"}, args...)
	cmd := exec.Command("docker", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		styles.GetErrorStyle().Println("Error running Docker container:", err)
	}
}

func DockerPS(args []string) {
	cmdArgs := append([]string{"ps"}, args...)
	cmd := exec.Command("docker", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		styles.GetErrorStyle().Println("Error listing Docker containers:", err)
	}
}

func DockerCompose(args []string) {
	cmd := exec.Command("docker-compose", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		styles.GetErrorStyle().Println("Error running docker-compose command:", err)
	}
}
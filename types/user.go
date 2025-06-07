package types

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/tacheraSasi/ellie/configs"
)

// UserContext holds realtime information about the user and their environment
type UserContext struct {
	Username     string
	Hostname     string
	OS           string
	Shell        string
	CurrentDir   string
	GitBranch    string
	GitStatus    string
	TimeOfDay    string
	MemoryUsage  string
	CPUUsage     string
	LastCommand  string
	CommandCount int
	EllieDir     string
}

// NewUserContext creates a new UserContext with current system information
func NewUserContext() *UserContext {
	ctx := &UserContext{
		TimeOfDay: getTimeOfDay(),
	}

	// Get basic system info
	ctx.Username = getUsername()
	ctx.Hostname = getHostname()
	ctx.OS = runtime.GOOS
	ctx.Shell = getShell()
	ctx.CurrentDir = getCurrentDir()
	ctx.GitBranch = getGitBranch()
	ctx.GitStatus = getGitStatus()
	ctx.MemoryUsage = getMemoryUsage()
	ctx.CPUUsage = getCPUUsage()
	ctx.EllieDir = getEllieDir()

	return ctx
}

// GetContextString returns a formatted string with all context information
func (ctx *UserContext) GetContextString() string {
	return fmt.Sprintf(`
Current User Context:
- User: %s@%s
- OS: %s
- Shell: %s
- Current Directory: %s
- Git Branch: %s
- Git Status: %s
- Time of Day: %s
- Memory Usage: %s
- CPU Usage: %s
- Last Command: %s
- Command Count: %d
`, ctx.Username, ctx.Hostname, ctx.OS, ctx.Shell, ctx.CurrentDir, ctx.GitBranch, ctx.GitStatus,
		ctx.TimeOfDay, ctx.MemoryUsage, ctx.CPUUsage, ctx.LastCommand, ctx.CommandCount)
}

// UpdateContext updates the context with new information
func (ctx *UserContext) UpdateContext() {
	ctx.TimeOfDay = getTimeOfDay()
	ctx.CurrentDir = getCurrentDir()
	ctx.GitBranch = getGitBranch()
	ctx.GitStatus = getGitStatus()
	ctx.MemoryUsage = getMemoryUsage()
	ctx.CPUUsage = getCPUUsage()
}

// Helper functions to get system information
func getUsername() string {
	user, err := os.UserHomeDir()
	if err != nil {
		return "unknown"
	}
	return filepath.Base(user)
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}

func getShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "unknown"
	}
	return filepath.Base(shell)
}

func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "unknown"
	}
	return dir
}

func getEllieDir() string {
	homeDir,err := os.UserHomeDir()
	if err != nil{
		return ""
	}
	return homeDir + "/" + configs.ConfigDirName
}

func getGitBranch() string {
	cmd := exec.Command("git", "branch", "--show-current")
	output, err := cmd.Output()
	if err != nil {
		return "not a git repository"
	}
	return strings.TrimSpace(string(output))
}

func getGitStatus() string {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return "not a git repository"
	}
	if len(output) == 0 {
		return "clean"
	}
	return "modified"
}

func getTimeOfDay() string {
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		return "morning"
	case hour < 17:
		return "afternoon"
	default:
		return "evening"
	}
}

func getMemoryUsage() string {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("top", "-l", "1", "-stats", "mem")
	case "linux":
		cmd = exec.Command("free", "-h")
	default:
		return "unknown"
	}
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}

func getCPUUsage() string {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("top", "-l", "1", "-stats", "cpu")
	case "linux":
		cmd = exec.Command("top", "-bn1")
	default:
		return "unknown"
	}
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}

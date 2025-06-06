package static

import (
	"fmt"

	"github.com/tacheraSasi/ellie/types"
)

func Instructions(userCtx types.UserContext) string {
	prompt, err := GetStaticFile("./prompt.txt")
	if err != nil {
		return "Error loading instructions"
	}

	// Formats the prompt with actual user context values
	formattedPrompt := fmt.Sprintf(prompt,
		userCtx.Username, userCtx.Hostname, // User: <username>@<hostname>
		userCtx.OS,           // OS: <os>
		userCtx.Shell,        // Shell: <shell>
		userCtx.CurrentDir,   // Current Directory: <current_dir>
		userCtx.GitBranch,    // Git Branch: <git_branch>
		userCtx.GitStatus,    // Git Status: <git_status>
		userCtx.TimeOfDay,    // Time of Day: <time_of_day>
		userCtx.MemoryUsage,  // Memory Usage: <memory_usage>
		userCtx.CPUUsage,     // CPU Usage: <cpu_usage>
		userCtx.LastCommand,  // Last Command: <last_command>
		userCtx.CommandCount, // Command Count: <command_count>
		userCtx.OS,           // OS version in user_info
		userCtx.CurrentDir,   // Workspace path in user_info
		userCtx.Shell,        // Shell in user_info
	)

	return formattedPrompt
}

package elliecore

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"strings"
	"time"
)

// RunCmd executes a shell command and returns the output or error.
func RunCmd(cmd string) string {
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "Error: " + err.Error() + "\n" + string(out)
	}
	return string(out)
}

// RunCmdWithEnv executes a shell command with additional environment variables.
func RunCmdWithEnv(cmd string, envs string) string {
	envList := strings.Split(envs, ";")
	command := exec.Command("sh", "-c", cmd)
	command.Env = append(os.Environ(), envList...)
	out, err := command.CombinedOutput()
	if err != nil {
		return "Error: " + err.Error() + "\n" + string(out)
	}
	return string(out)
}

// ReadFile reads the content of a file.
func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return "Error: " + err.Error()
	}
	return string(data)
}

// WriteFile overwrites a file with given content.
func WriteFile(path, content string) string {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return "Error: " + err.Error()
	}
	return "OK"
}

// AppendFile appends content to a file, creates it if not exists.
func AppendFile(path, content string) string {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "Error: " + err.Error()
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		return "Error: " + err.Error()
	}
	return "OK"
}

// DeleteFile removes a file.
func DeleteFile(path string) string {
	err := os.Remove(path)
	if err != nil {
		return "Error: " + err.Error()
	}
	return "OK"
}

// ListDir lists files and directories in a path.
func ListDir(path string) string {
	files, err := os.ReadDir(path)
	if err != nil {
		return "Error: " + err.Error()
	}
	var builder strings.Builder
	for _, f := range files {
		builder.WriteString(f.Name() + "\n")
	}
	return builder.String()
}

// GetEnv returns the value of an environment variable.
func GetEnv(key string) string {
	return os.Getenv(key)
}

// SetEnv sets an environment variable in the current process.
func SetEnv(key, value string) string {
	err := os.Setenv(key, value)
	if err != nil {
		return "Error: " + err.Error()
	}
	return "OK"
}

// GetCwd returns the current working directory.
func GetCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		return "Error: " + err.Error()
	}
	return dir
}

// ChangeDir changes the current working directory.
func ChangeDir(path string) string {
	err := os.Chdir(path)
	if err != nil {
		return "Error: " + err.Error()
	}
	return "OK"
}

// FileExists checks if a file or dir exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// FileMetadata returns JSON with file size, readonly, modified time.
func FileMetadata(path string) string {
	info, err := os.Stat(path)
	if err != nil {
		return "Error: " + err.Error()
	}

	meta := map[string]interface{}{
		"size":     info.Size(),
		"readonly": info.Mode().Perm()&0200 == 0,
		"modified": info.ModTime().Format(time.RFC3339),
		"is_dir":   info.IsDir(),
	}

	jsonData, _ := json.MarshalIndent(meta, "", "  ")
	return string(jsonData)
}

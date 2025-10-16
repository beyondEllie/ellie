package actions

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/tacheraSasi/ellie/elliecore"
	"github.com/tacheraSasi/ellie/styles"
)

// Disk displays disk usage information
func Disk(args []string) {
	// Determine path to check
	path := "."
	if len(args) > 1 {
		path = args[1]
	}

	styles.InfoStyle.Println("💾 Disk Usage Information")
	styles.InfoStyle.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// Get disk usage using elliecore
	diskInfo := elliecore.GetDiskUsage(path)

	if strings.HasPrefix(diskInfo, "Error:") {
		styles.ErrorStyle.Printf("❌ %s\n", diskInfo)
		return
	}

	// Format and display the output
	formatDiskOutput(diskInfo, runtime.GOOS, path)
}

// formatDiskOutput formats disk usage information for better readability
func formatDiskOutput(rawOutput, osType, path string) {
	lines := strings.Split(rawOutput, "\n")

	switch osType {
	case "windows":
		// Windows dir /-c output
		styles.Bold.Printf("\n📁 Path: %s\n\n", path)
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}
			if strings.Contains(line, "bytes") {
				styles.SuccessStyle.Println("  " + line)
			} else {
				fmt.Println("  " + line)
			}
		}

	case "darwin", "linux":
		// Unix df -h output
		if len(lines) > 0 {
			// Display header with styling
			header := lines[0]
			styles.Bold.Println("\n" + header)
			styles.InfoStyle.Println(strings.Repeat("─", len(header)))
		}

		// Display data rows
		for i := 1; i < len(lines); i++ {
			line := strings.TrimSpace(lines[i])
			if line == "" {
				continue
			}

			// Parse the line to extract usage information
			fields := strings.Fields(line)
			if len(fields) >= 5 {
				// fields: [Filesystem, Size, Used, Available, Use%, MountedOn]
				filesystem := fields[0]
				size := fields[1]
				used := fields[2]
				available := fields[3]
				usePercent := fields[4]
				mountPoint := strings.Join(fields[5:], " ")

				// Color code based on usage percentage
				usageNum := strings.TrimSuffix(usePercent, "%")
				var colorStyle func(...interface{}) string

				// Determine color based on usage
				if len(usageNum) > 0 {
					if usageNum[0] >= '0' && usageNum[0] <= '9' {
						if strings.HasPrefix(usageNum, "9") || strings.HasPrefix(usageNum, "10") {
							colorStyle = styles.ErrorStyle.Sprint
						} else if strings.HasPrefix(usageNum, "7") || strings.HasPrefix(usageNum, "8") {
							colorStyle = styles.WarningStyle.Sprint
						} else {
							colorStyle = styles.SuccessStyle.Sprint
						}
					} else {
						colorStyle = styles.SuccessStyle.Sprint
					}
				} else {
					colorStyle = styles.SuccessStyle.Sprint
				}

				// Format the output
				fmt.Printf("%-25s %8s %8s %8s %10s  %s\n",
					filesystem,
					size,
					used,
					available,
					colorStyle(usePercent),
					mountPoint,
				)
			} else {
				// Fallback for lines that don't match expected format
				fmt.Println(line)
			}
		}

		fmt.Println()
		displayDiskLegend()

	default:
		// Fallback for unknown OS
		fmt.Println(rawOutput)
	}

	fmt.Println()
}

// displayDiskLegend shows a legend explaining the color coding
func displayDiskLegend() {
	styles.InfoStyle.Println("\n📊 Usage Legend:")
	styles.SuccessStyle.Println("  🟢 Green: < 70% - Healthy")
	styles.WarningStyle.Println("  🟡 Yellow: 70-89% - Warning")
	styles.ErrorStyle.Println("  🔴 Red: ≥ 90% - Critical")
}

// DiskAll displays disk usage for all mounted filesystems
func DiskAll() {
	styles.InfoStyle.Println("💾 All Disk Partitions")
	styles.InfoStyle.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	var diskInfo string
	switch runtime.GOOS {
	case "windows":
		diskInfo = elliecore.RunCmd("wmic logicaldisk get caption,size,freespace")
	case "darwin":
		diskInfo = elliecore.RunCmd("df -h")
	case "linux":
		diskInfo = elliecore.RunCmd("df -h --total")
	default:
		diskInfo = elliecore.RunCmd("df -h")
	}

	if strings.HasPrefix(diskInfo, "Error:") {
		styles.ErrorStyle.Printf("❌ %s\n", diskInfo)
		return
	}

	formatDiskOutput(diskInfo, runtime.GOOS, "all")
}

// DiskSpace displays available and used space summary
func DiskSpace() {
	styles.InfoStyle.Println("💾 Disk Space Summary")
	styles.InfoStyle.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	cwd, err := os.Getwd()
	if err != nil {
		styles.ErrorStyle.Printf("❌ Error getting current directory: %v\n", err)
		return
	}

	diskInfo := elliecore.GetDiskUsage(cwd)

	if strings.HasPrefix(diskInfo, "Error:") {
		styles.ErrorStyle.Printf("❌ %s\n", diskInfo)
		return
	}

	// Parse and display a simplified summary
	lines := strings.Split(diskInfo, "\n")

	switch runtime.GOOS {
	case "darwin", "linux":
		for i, line := range lines {
			if i == 0 {
				continue // Skip header
			}
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			fields := strings.Fields(line)
			if len(fields) >= 5 {
				size := fields[1]
				used := fields[2]
				available := fields[3]
				usePercent := fields[4]

				fmt.Println()
				styles.Bold.Printf("📂 Current Directory: %s\n\n", cwd)
				styles.InfoStyle.Printf("  Total Space:     %s\n", size)
				styles.SuccessStyle.Printf("  Available Space: %s\n", available)
				styles.WarningStyle.Printf("  Used Space:      %s\n", used)

				// Color code usage percentage
				usageNum := strings.TrimSuffix(usePercent, "%")
				if strings.HasPrefix(usageNum, "9") || strings.HasPrefix(usageNum, "10") {
					styles.ErrorStyle.Printf("  Usage:           %s\n", usePercent)
				} else if strings.HasPrefix(usageNum, "7") || strings.HasPrefix(usageNum, "8") {
					styles.WarningStyle.Printf("  Usage:           %s\n", usePercent)
				} else {
					styles.SuccessStyle.Printf("  Usage:           %s\n", usePercent)
				}
				fmt.Println()
				break
			}
		}
	default:
		fmt.Println(diskInfo)
	}
}

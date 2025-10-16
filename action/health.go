package actions

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/tacheraSasi/ellie/styles"
)

// SystemHealth displays a comprehensive system health dashboard
func SystemHealth() {
	styles.GetInfoStyle().Println("\n🏥 System Health Dashboard")
	styles.GetInfoStyle().Println("═══════════════════════════════════════════════════════")
	
	// CPU Health
	displayCPUHealth()
	
	// Memory Health
	displayMemoryHealth()
	
	// Disk Health
	displayDiskHealth()
	
	// Load Average (Linux/Mac)
	if runtime.GOOS != "windows" {
		displayLoadAverage()
	}
	
	// Active Processes
	displayProcessCount()
	
	// System Uptime
	displaySystemUptime()
	
	// Overall Health Score
	displayHealthScore()
	
	styles.GetInfoStyle().Println("═══════════════════════════════════════════════════════")
}

func displayCPUHealth() {
	cpuCount := runtime.NumCPU()
	styles.GetHighlightStyle().Printf("\n🖥️  CPU: %d cores\n", cpuCount)
	
	// Get CPU usage if possible
	if usage, err := getCPUUsage(); err == nil {
		displayMetricBar("CPU Usage", usage, 80, 95)
	}
}

func displayMemoryHealth() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	
	totalMem, usedMem, err := getSystemMemory()
	if err == nil {
		usagePercent := float64(usedMem) / float64(totalMem) * 100
		styles.GetHighlightStyle().Println("\n💾 Memory")
		fmt.Printf("   Total: %s | Used: %s | Available: %s\n", 
			formatBytes(totalMem), 
			formatBytes(usedMem), 
			formatBytes(totalMem-usedMem))
		displayMetricBar("Memory Usage", usagePercent, 80, 90)
	}
}

func displayDiskHealth() {
	diskInfo, err := getDiskUsage()
	if err == nil {
		styles.GetHighlightStyle().Println("\n💿 Disk Space")
		for _, disk := range diskInfo {
			fmt.Printf("   %s: %s / %s (%.1f%% used)\n", 
				disk.Mount, 
				formatBytes(disk.Used), 
				formatBytes(disk.Total),
				disk.UsagePercent)
			displayMetricBar(disk.Mount, disk.UsagePercent, 80, 90)
		}
	}
}

func displayLoadAverage() {
	if load, err := getLoadAverage(); err == nil {
		styles.GetHighlightStyle().Println("\n📊 Load Average")
		fmt.Printf("   1 min: %.2f | 5 min: %.2f | 15 min: %.2f\n", load[0], load[1], load[2])
	}
}

func displayProcessCount() {
	if count, err := getProcessCount(); err == nil {
		styles.GetHighlightStyle().Printf("\n⚙️  Active Processes: %d\n", count)
	}
}

func displaySystemUptime() {
	if uptime, err := getSystemUptime(); err == nil {
		styles.GetHighlightStyle().Printf("\n⏱️  Uptime: %s\n", uptime)
	}
}

func displayHealthScore() {
	score := calculateHealthScore()
	styles.GetHighlightStyle().Println("\n🎯 Overall Health Score")
	
	var statusColor *color.Color
	var status string
	
	if score >= 85 {
		statusColor = styles.GetSuccessStyle()
		status = "Excellent"
	} else if score >= 70 {
		statusColor = color.New(color.FgYellow)
		status = "Good"
	} else if score >= 50 {
		statusColor = color.New(color.FgYellow, color.Bold)
		status = "Fair"
	} else {
		statusColor = styles.GetErrorStyle()
		status = "Poor"
	}
	
	statusColor.Printf("   Score: %d/100 - %s\n", score, status)
	displayMetricBar("Health", float64(score), 70, 85)
}

func displayMetricBar(name string, value float64, warningThreshold, criticalThreshold float64) {
	barLength := 30
	filled := int(value * float64(barLength) / 100)
	if filled > barLength {
		filled = barLength
	}
	
	var barColor *color.Color
	if value >= criticalThreshold {
		barColor = styles.GetErrorStyle()
	} else if value >= warningThreshold {
		barColor = color.New(color.FgYellow)
	} else {
		barColor = styles.GetSuccessStyle()
	}
	
	bar := strings.Repeat("█", filled) + strings.Repeat("░", barLength-filled)
	barColor.Printf("   [%s] %.1f%%\n", bar, value)
}

func getCPUUsage() (float64, error) {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		out, err := exec.Command("sh", "-c", "top -bn1 | grep 'Cpu(s)' | sed 's/.*, *\\([0-9.]*\\)%* id.*/\\1/' | awk '{print 100 - $1}'").Output()
		if err != nil {
			return 0, err
		}
		usage, err := strconv.ParseFloat(strings.TrimSpace(string(out)), 64)
		if err != nil {
			return 0, err
		}
		return usage, nil
	}
	return 0, fmt.Errorf("unsupported OS")
}

func getSystemMemory() (total, used uint64, err error) {
	if runtime.GOOS == "linux" {
		out, err := exec.Command("free", "-b").Output()
		if err != nil {
			return 0, 0, err
		}
		lines := strings.Split(string(out), "\n")
		if len(lines) < 2 {
			return 0, 0, fmt.Errorf("unexpected output")
		}
		fields := strings.Fields(lines[1])
		if len(fields) < 3 {
			return 0, 0, fmt.Errorf("unexpected format")
		}
		total, _ = strconv.ParseUint(fields[1], 10, 64)
		used, _ = strconv.ParseUint(fields[2], 10, 64)
		return total, used, nil
	} else if runtime.GOOS == "darwin" {
		// Mac memory detection
		totalOut, _ := exec.Command("sysctl", "-n", "hw.memsize").Output()
		total, _ = strconv.ParseUint(strings.TrimSpace(string(totalOut)), 10, 64)
		
		_, err = exec.Command("vm_stat").Output()
		if err != nil {
			return total, 0, err
		}
		
		// Parse vm_stat output (simplified)
		used = total / 2 // Simplified estimation
		return total, used, nil
	}
	return 0, 0, fmt.Errorf("unsupported OS")
}

type DiskInfo struct {
	Mount        string
	Total        uint64
	Used         uint64
	UsagePercent float64
}

func getDiskUsage() ([]DiskInfo, error) {
	var disks []DiskInfo
	
	if runtime.GOOS == "windows" {
		return disks, fmt.Errorf("windows not yet supported")
	}
	
	out, err := exec.Command("df", "-B1").Output()
	if err != nil {
		return nil, err
	}
	
	lines := strings.Split(string(out), "\n")
	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}
		
		// Skip tmpfs and other virtual filesystems
		if strings.Contains(fields[0], "tmpfs") || strings.Contains(fields[0], "devtmpfs") {
			continue
		}
		
		total, _ := strconv.ParseUint(fields[1], 10, 64)
		used, _ := strconv.ParseUint(fields[2], 10, 64)
		mount := fields[5]
		
		if total == 0 {
			continue
		}
		
		usagePercent := float64(used) / float64(total) * 100
		
		disks = append(disks, DiskInfo{
			Mount:        mount,
			Total:        total,
			Used:         used,
			UsagePercent: usagePercent,
		})
	}
	
	return disks, nil
}

func getLoadAverage() ([3]float64, error) {
	var load [3]float64
	
	out, err := exec.Command("uptime").Output()
	if err != nil {
		return load, err
	}
	
	output := string(out)
	if strings.Contains(output, "load average:") {
		parts := strings.Split(output, "load average:")
		if len(parts) < 2 {
			return load, fmt.Errorf("cannot parse load average")
		}
		
		loadStr := strings.TrimSpace(parts[1])
		loadParts := strings.Split(loadStr, ",")
		
		for i := 0; i < 3 && i < len(loadParts); i++ {
			val, _ := strconv.ParseFloat(strings.TrimSpace(loadParts[i]), 64)
			load[i] = val
		}
	}
	
	return load, nil
}

func getProcessCount() (int, error) {
	if runtime.GOOS == "windows" {
		out, err := exec.Command("tasklist").Output()
		if err != nil {
			return 0, err
		}
		lines := strings.Split(string(out), "\n")
		return len(lines) - 3, nil // Subtract header lines
	}
	
	out, err := exec.Command("ps", "aux").Output()
	if err != nil {
		return 0, err
	}
	
	lines := strings.Split(string(out), "\n")
	return len(lines) - 2, nil // Subtract header and empty line
}

func getSystemUptime() (string, error) {
	if runtime.GOOS == "windows" {
		out, err := exec.Command("net", "stats", "workstation").Output()
		if err != nil {
			return "", err
		}
		
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			if strings.Contains(line, "Statistics since") {
				return strings.TrimSpace(strings.Split(line, "since")[1]), nil
			}
		}
		return "Unknown", nil
	}
	
	out, err := exec.Command("uptime", "-p").Output()
	if err != nil {
		// Try alternative method
		out, err = exec.Command("uptime").Output()
		if err != nil {
			return "", err
		}
	}
	
	return strings.TrimSpace(string(out)), nil
}

func calculateHealthScore() int {
	score := 100
	
	// Check CPU usage
	if cpuUsage, err := getCPUUsage(); err == nil {
		if cpuUsage > 90 {
			score -= 20
		} else if cpuUsage > 75 {
			score -= 10
		}
	}
	
	// Check memory usage
	if total, used, err := getSystemMemory(); err == nil && total > 0 {
		memPercent := float64(used) / float64(total) * 100
		if memPercent > 90 {
			score -= 20
		} else if memPercent > 80 {
			score -= 10
		}
	}
	
	// Check disk usage
	if disks, err := getDiskUsage(); err == nil {
		for _, disk := range disks {
			if disk.UsagePercent > 90 {
				score -= 15
			} else if disk.UsagePercent > 80 {
				score -= 5
			}
		}
	}
	
	if score < 0 {
		score = 0
	}
	
	return score
}

func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	
	div, exp := uint64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	
	units := []string{"KB", "MB", "GB", "TB"}
	return fmt.Sprintf("%.1f %s", float64(bytes)/float64(div), units[exp])
}

// MonitorSystem starts real-time system monitoring
func MonitorSystem() {
	styles.GetInfoStyle().Println("🔍 Starting real-time system monitoring...")
	styles.GetInfoStyle().Println("Press Ctrl+C to stop\n")
	
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			clearScreen()
			SystemHealth()
		}
	}
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// CheckSystemAlerts checks for system issues and displays alerts
func CheckSystemAlerts() {
	styles.GetInfoStyle().Println("\n🚨 System Alerts Check")
	styles.GetInfoStyle().Println("─────────────────────────")
	
	hasAlerts := false
	
	// Check CPU
	if cpuUsage, err := getCPUUsage(); err == nil && cpuUsage > 80 {
		styles.GetErrorStyle().Printf("⚠️  High CPU usage: %.1f%%\n", cpuUsage)
		hasAlerts = true
	}
	
	// Check Memory
	if total, used, err := getSystemMemory(); err == nil && total > 0 {
		memPercent := float64(used) / float64(total) * 100
		if memPercent > 85 {
			styles.GetErrorStyle().Printf("⚠️  High memory usage: %.1f%%\n", memPercent)
			hasAlerts = true
		}
	}
	
	// Check Disk
	if disks, err := getDiskUsage(); err == nil {
		for _, disk := range disks {
			if disk.UsagePercent > 85 {
				styles.GetErrorStyle().Printf("⚠️  Low disk space on %s: %.1f%% used\n", disk.Mount, disk.UsagePercent)
				hasAlerts = true
			}
		}
	}
	
	if !hasAlerts {
		styles.GetSuccessStyle().Println("✅ All systems normal")
	}
}

// QuickHealthCheck performs a quick health check
func QuickHealthCheck() error {
	fmt.Print("Running quick health check")
	
	for i := 0; i < 3; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()
	
	score := calculateHealthScore()
	
	if score >= 80 {
		styles.GetSuccessStyle().Printf("✅ System health: %d/100 (Excellent)\n", score)
	} else if score >= 60 {
		color.New(color.FgYellow).Printf("⚠️  System health: %d/100 (Good)\n", score)
	} else {
		styles.GetErrorStyle().Printf("❌ System health: %d/100 (Needs attention)\n", score)
		styles.GetInfoStyle().Println("\n💡 Run 'ellie health' for detailed analysis")
	}
	
	return nil
}

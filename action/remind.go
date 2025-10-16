package actions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

func Remind() {
	styles.Cyan.Print("💡 Set a reminder")
	styles.DimText.Println(" (Schedule desktop notifications)")

	title := getReminderTitle()
	if title == "" {
		return
	}

	duration := getReminderDuration()
	if duration == 0 {
		return
	}

	setReminder(title, duration)
}

func getReminderTitle() string {
	for {
		title, err := utils.GetInput("What do you want to remind yourself?")
		if err == nil && title != "" {
			return title
		}
		styles.ErrorStyle.Println("🚫 Reminder title cannot be empty.")
	}
}

func getReminderDuration() time.Duration {
	for {
		input, err := utils.GetInput("⏳ When should I remind you? (e.g., 10s, 5m, 2h, 3d, 1w)")
		if err != nil {
			styles.ErrorStyle.Println("🚫 Failed to read input. Please try again.")
			continue
		}

		duration, err := parseDuration(input)
		if err == nil && duration > 0 {
			return duration
		}

		styles.ErrorStyle.Println("🚫 Invalid duration. Try formats like '10s', '5m', '2h', '3d', '1w'.")
	}
}

func parseDuration(input string) (time.Duration, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, fmt.Errorf("empty input")
	}

	// Try standard time.ParseDuration first (handles s, m, h)
	duration, err := time.ParseDuration(input)
	if err == nil {
		return duration, nil
	}

	// Parse custom formats: days (d) and weeks (w)
	re := regexp.MustCompile(`^(\d+(?:\.\d+)?)(d|w)$`)
	matches := re.FindStringSubmatch(strings.ToLower(input))

	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid duration format")
	}

	value, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, err
	}

	unit := matches[2]
	switch unit {
	case "d":
		return time.Duration(value * 24 * float64(time.Hour)), nil
	case "w":
		return time.Duration(value * 7 * 24 * float64(time.Hour)), nil
	default:
		return 0, fmt.Errorf("unsupported time unit: %s", unit)
	}
}

func setReminder(title string, duration time.Duration) {
	// Format duration for display
	durationStr := formatDuration(duration)
	styles.SuccessStyle.Printf("✅ Reminder set! I will remind you in %s.\n", durationStr)

	time.AfterFunc(duration, func() {
		// Send Desktop Notification
		err := beeep.Notify("🔔 Ellie Reminder", title, "")
		if err != nil {
			utils.Error("❌ Failed to send notification: " + err.Error())
		} else {
			styles.InfoStyle.Printf("\n🔔 Reminder: %s\n", title)
		}
	})
}

func formatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%d seconds", int(d.Seconds()))
	} else if d < time.Hour {
		return fmt.Sprintf("%d minutes", int(d.Minutes()))
	} else if d < 24*time.Hour {
		return fmt.Sprintf("%.1f hours", d.Hours())
	} else if d < 7*24*time.Hour {
		return fmt.Sprintf("%.1f days", d.Hours()/24)
	} else {
		return fmt.Sprintf("%.1f weeks", d.Hours()/(24*7))
	}
}

package actions

import (
	"time"

	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

func Remind() {
	styles.Cyan.Print("💡 ellie remind")

	title, err := getTitle()
	if err != nil {
		utils.Error("❌ Something went wrong, failed to get the title.")
		return
	}

	duration, err := getDuration()
	if err != nil {
		utils.Error("❌ Failed to get reminder duration.")
		return
	}

	setReminder(title, duration)
}

func getTitle() (string, error) {
	for {
		subject, err := utils.GetInput("📝 What do you want to remind yourself?")
		if err == nil && subject != "" {
			return subject, nil
		}
		styles.ErrorStyle.Println("🚫 Title cannot be empty.")
	}
}

func getDuration() (time.Duration, error) {
	for {
		input, err := utils.GetInput("⏳ In how many seconds/minutes/hours? (e.g., 10s, 5m, 2h)")
		if err != nil {
			return 0, err
		}
		duration, err := time.ParseDuration(input)
		if err == nil && duration > 0 {
			return duration, nil
		}
		styles.ErrorStyle.Println("🚫 Invalid duration. Try formats like '10s', '5m', '2h'.")
	}
}

func setReminder(title string, duration time.Duration) {
	styles.SuccessStyle.Printf("✅ Reminder set! I will remind you in %v.\n", duration)

	time.AfterFunc(duration, func() {
		// styles.ReminderStyle.Printf("\n🔔 Reminder: %s\n", title)
	})
}

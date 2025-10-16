package actions

import (
	"fmt"
	"time"

	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/styles"
)

// ShowWelcome displays an impressive welcome message
func ShowWelcome() {
	username := configs.GetEnv("USERNAME")
	if username == "" {
		username = "Friend"
	}
	
	hour := time.Now().Hour()
	var greeting string
	var emoji string
	
	switch {
	case hour < 12:
		greeting = "Good morning"
		emoji = "🌅"
	case hour < 18:
		greeting = "Good afternoon"
		emoji = "☀️"
	case hour < 22:
		greeting = "Good evening"
		emoji = "🌆"
	default:
		greeting = "Hello"
		emoji = "🌙"
	}
	
	styles.GetInfoStyle().Println("\n╔════════════════════════════════════════════════════════════╗")
	styles.GetHighlightStyle().Printf("║  %s %s, %s!%*s║\n", emoji, greeting, username, 42-len(greeting)-len(username), "")
	styles.GetInfoStyle().Println("╚════════════════════════════════════════════════════════════╝")
	
	// Show quick status
	fmt.Println()
	styles.GetInfoStyle().Print("⚡ Quick Status: ")
	
	// Quick health check
	score := calculateHealthScore()
	if score >= 80 {
		styles.GetSuccessStyle().Print("System Healthy")
	} else if score >= 60 {
		styles.GetWarningStyle().Print("System Fair")
	} else {
		styles.GetErrorStyle().Print("Needs Attention")
	}
	
	// Check for Git repo
	if isGitRepo() {
		fmt.Print(" | ")
		if hasUncommittedChanges() {
			styles.GetWarningStyle().Print("📝 Uncommitted changes")
		} else {
			styles.GetSuccessStyle().Print("✓ Git clean")
		}
	}
	
	fmt.Println()
	fmt.Println()
	
	// Show time-based suggestions
	showQuickSuggestions(hour)
	
	styles.GetInfoStyle().Println("\n💡 Type 'ellie assist' for context-aware help")
	styles.GetInfoStyle().Println("💡 Type 'ellie --help' to see all commands")
	fmt.Println()
}

func showQuickSuggestions(hour int) {
	styles.GetHighlightStyle().Println("🚀 Quick Actions:")
	
	if hour >= 6 && hour < 10 {
		fmt.Println("   • ellie start-day - Start your development day")
		fmt.Println("   • ellie health - Check system health")
	} else if hour >= 10 && hour < 14 {
		fmt.Println("   • ellie focus - Activate focus mode")
		fmt.Println("   • ellie suggest - Get smart suggestions")
	} else if hour >= 14 && hour < 18 {
		fmt.Println("   • ellie git status - Check your progress")
		fmt.Println("   • ellie todo list - Review tasks")
	} else {
		fmt.Println("   • ellie git commit - Save your work")
		fmt.Println("   • ellie health - Final system check")
	}
}

// ShowBanner displays a stylish banner
func ShowBanner() {
	styles.GetHighlightStyle().Println(`
 ███████╗██╗     ██╗     ██╗███████╗
 ██╔════╝██║     ██║     ██║██╔════╝
 █████╗  ██║     ██║     ██║█████╗  
 ██╔══╝  ██║     ██║     ██║██╔══╝  
 ███████╗███████╗███████╗██║███████╗
 ╚══════╝╚══════╝╚══════╝╚═╝╚══════╝`)
	
	styles.GetInfoStyle().Printf("    Your AI-Powered CLI Companion v%s\n", configs.VERSION)
	styles.GetSuccessStyle().Println("    Built with ❤️  for developers")
	fmt.Println()
}

// ShowFirstRunWelcome displays welcome message for first-time users
func ShowFirstRunWelcome() {
	ShowBanner()
	
	styles.GetHighlightStyle().Println("🎉 Welcome to Ellie!")
	fmt.Println()
	
	styles.GetInfoStyle().Println("Ellie is your personal command-line companion designed to make")
	styles.GetInfoStyle().Println("system management and automation effortless.")
	fmt.Println()
	
	styles.GetHighlightStyle().Println("✨ Key Features:")
	fmt.Println("   🏥 System Health Monitoring - Real-time insights")
	fmt.Println("   🤖 Smart Assistant - Context-aware suggestions")
	fmt.Println("   ⚙️  Automation - Schedule routine tasks")
	fmt.Println("   🔀 Git Mastery - Conventional commits made easy")
	fmt.Println("   📝 Todo & Project Management - Stay organized")
	fmt.Println("   🎯 Focus Mode - Minimize distractions")
	fmt.Println()
	
	styles.GetHighlightStyle().Println("🚀 Quick Start:")
	fmt.Println("   • ellie assist - Get context-aware help")
	fmt.Println("   • ellie health - Check system status")
	fmt.Println("   • ellie suggest - Get smart suggestions")
	fmt.Println("   • ellie --help - See all commands")
	fmt.Println()
	
	styles.GetSuccessStyle().Println("Let's get started! 🎊")
	fmt.Println()
}

// ShowInspiringQuote shows a random inspiring quote
func ShowInspiringQuote() {
	quotes := []string{
		"💎 'Code is like humor. When you have to explain it, it's bad.' - Cory House",
		"🚀 'First, solve the problem. Then, write the code.' - John Johnson",
		"⚡ 'Simplicity is the soul of efficiency.' - Austin Freeman",
		"🎯 'Make it work, make it right, make it fast.' - Kent Beck",
		"🌟 'The best error message is the one that never shows up.' - Thomas Fuchs",
		"💡 'Programming isn't about what you know; it's about what you can figure out.' - Chris Pine",
		"🔥 'Clean code always looks like it was written by someone who cares.' - Robert C. Martin",
		"✨ 'Any fool can write code that a computer can understand. Good programmers write code that humans can understand.' - Martin Fowler",
	}
	
	now := time.Now()
	index := now.Day() % len(quotes)
	
	styles.GetInfoStyle().Println("\n" + quotes[index])
	fmt.Println()
}

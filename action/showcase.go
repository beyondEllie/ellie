package actions

import (
	"fmt"
	"time"

	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/styles"
)

// ShowcaseFeatures demonstrates all the impressive features
func ShowcaseFeatures() {
	ShowBanner()
	
	styles.GetHighlightStyle().Println("🎪 Welcome to the Ellie Feature Showcase!")
	styles.GetInfoStyle().Println("═══════════════════════════════════════════════════════")
	fmt.Println()
	
	// Feature 1: System Health Monitoring
	showcaseHealthMonitoring()
	waitForUser()
	
	// Feature 2: Smart Assistant
	showcaseSmartAssistant()
	waitForUser()
	
	// Feature 3: Automation
	showcaseAutomation()
	waitForUser()
	
	// Feature 4: Productivity Tools
	showcaseProductivity()
	waitForUser()
	
	// Feature 5: Git Mastery
	showcaseGitFeatures()
	waitForUser()
	
	// Finale
	showFinale()
}

func showcaseHealthMonitoring() {
	styles.GetHighlightStyle().Println("\n🏥 System Health Monitoring")
	styles.GetInfoStyle().Println("─────────────────────────────────────")
	fmt.Println()
	
	fmt.Println("Ellie provides real-time system monitoring with:")
	fmt.Println("  ✓ CPU, Memory, and Disk usage tracking")
	fmt.Println("  ✓ Load average monitoring")
	fmt.Println("  ✓ Process count tracking")
	fmt.Println("  ✓ Health score calculation")
	fmt.Println("  ✓ Proactive alerts for issues")
	fmt.Println()
	
	styles.GetSuccessStyle().Println("Running quick health check...")
	QuickHealthCheck()
	
	fmt.Println()
	styles.GetInfoStyle().Println("💡 Try: ellie health, ellie monitor, ellie alerts")
}

func showcaseSmartAssistant() {
	styles.GetHighlightStyle().Println("\n🤖 Smart Assistant")
	styles.GetInfoStyle().Println("─────────────────────────────────────")
	fmt.Println()
	
	fmt.Println("Ellie's AI-powered assistant provides:")
	fmt.Println("  ✓ Context-aware command suggestions")
	fmt.Println("  ✓ Project type detection (Go, Node.js, Python, etc.)")
	fmt.Println("  ✓ Git repository status awareness")
	fmt.Println("  ✓ Time-based recommendations")
	fmt.Println("  ✓ Workflow analysis and insights")
	fmt.Println()
	
	styles.GetSuccessStyle().Println("Analyzing current context...")
	time.Sleep(500 * time.Millisecond)
	
	cwd := "/home/user/projects/awesome-app"
	fmt.Printf("  📂 Location: %s\n", cwd)
	fmt.Println("  🔀 Git Repository: Yes")
	fmt.Println("  🚀 Project Type: Go")
	fmt.Println()
	
	fmt.Println("Suggested actions:")
	fmt.Println("  • go build - Build the project")
	fmt.Println("  • ellie git status - Check repository")
	fmt.Println("  • ellie health - Monitor system")
	fmt.Println()
	
	styles.GetInfoStyle().Println("💡 Try: ellie suggest, ellie assist, ellie workflow")
}

func showcaseAutomation() {
	styles.GetHighlightStyle().Println("\n⚙️  Automation Scheduler")
	styles.GetInfoStyle().Println("─────────────────────────────────────")
	fmt.Println()
	
	fmt.Println("Automate your routine tasks with ease:")
	fmt.Println("  ✓ Schedule tasks (daily, hourly, weekly)")
	fmt.Println("  ✓ Custom time-based scheduling (@09:00)")
	fmt.Println("  ✓ Background daemon mode")
	fmt.Println("  ✓ Quick setup for common tasks")
	fmt.Println()
	
	styles.GetSuccessStyle().Println("Example automations:")
	fmt.Println()
	
	automations := []struct {
		name     string
		schedule string
		command  string
	}{
		{"Morning Health Check", "@09:00", "ellie health"},
		{"Hourly Git Status", "hourly", "ellie git status"},
		{"Daily Cleanup", "@23:00", "ellie disk space"},
		{"Weekly Updates", "weekly", "ellie update"},
	}
	
	for _, auto := range automations {
		fmt.Printf("  🟢 %s\n", auto.name)
		fmt.Printf("     Schedule: %s | Command: %s\n", auto.schedule, auto.command)
	}
	
	fmt.Println()
	styles.GetInfoStyle().Println("💡 Try: ellie automate quick, ellie automate add")
}

func showcaseProductivity() {
	styles.GetHighlightStyle().Println("\n📝 Productivity Tools")
	styles.GetInfoStyle().Println("─────────────────────────────────────")
	fmt.Println()
	
	fmt.Println("Stay organized and focused:")
	fmt.Println("  ✓ Todo management with categories and priorities")
	fmt.Println("  ✓ Project management with quick switching")
	fmt.Println("  ✓ Focus mode to minimize distractions")
	fmt.Println("  ✓ Daily routine automation")
	fmt.Println("  ✓ Command aliases for efficiency")
	fmt.Println()
	
	styles.GetSuccessStyle().Println("Example workflow:")
	fmt.Println()
	fmt.Println("  1. ellie start-day - Start your dev day")
	fmt.Println("  2. ellie focus - Activate focus mode")
	fmt.Println("  3. ellie todo list - Check your tasks")
	fmt.Println("  4. ellie switch api - Jump to project")
	fmt.Println("  5. ellie git commit - Perfect commits")
	fmt.Println()
	
	styles.GetInfoStyle().Println("💡 Try: ellie todo add, ellie project add, ellie focus")
}

func showcaseGitFeatures() {
	styles.GetHighlightStyle().Println("\n🔀 Git Mastery")
	styles.GetInfoStyle().Println("─────────────────────────────────────")
	fmt.Println()
	
	fmt.Println("Professional Git workflows made easy:")
	fmt.Println("  ✓ Conventional commits with guided prompts")
	fmt.Println("  ✓ Smart status checking")
	fmt.Println("  ✓ Branch management")
	fmt.Println("  ✓ Stash operations")
	fmt.Println("  ✓ Enhanced logging and diffs")
	fmt.Println()
	
	styles.GetSuccessStyle().Println("Example commit:")
	fmt.Println()
	fmt.Println("  📝 Type: feat")
	fmt.Println("  🎯 Scope: auth")
	fmt.Println("  📌 Description: Add OAuth2 support")
	fmt.Println("  💬 Body: Implemented Google and GitHub providers")
	fmt.Println()
	fmt.Println("  Result:")
	styles.GetSuccessStyle().Println("  feat(auth): Add OAuth2 support")
	fmt.Println()
	
	styles.GetInfoStyle().Println("💡 Try: ellie git commit, ellie git status")
}

func showFinale() {
	fmt.Println()
	styles.GetHighlightStyle().Println("═══════════════════════════════════════════════════════")
	styles.GetHighlightStyle().Println("           ✨ Ellie - Your CLI Companion ✨")
	styles.GetHighlightStyle().Println("═══════════════════════════════════════════════════════")
	fmt.Println()
	
	styles.GetSuccessStyle().Println("🎉 You've seen what Ellie can do!")
	fmt.Println()
	
	fmt.Println("Key highlights:")
	fmt.Println("  🏥 Intelligent system monitoring")
	fmt.Println("  🤖 Context-aware smart assistant")
	fmt.Println("  ⚙️  Powerful automation scheduler")
	fmt.Println("  📝 Productivity tools")
	fmt.Println("  🔀 Professional Git workflows")
	fmt.Println()
	
	styles.GetHighlightStyle().Println("Ready to get started?")
	fmt.Println()
	
	fmt.Println("Next steps:")
	fmt.Println("  1. ellie assist - Get personalized help")
	fmt.Println("  2. ellie suggest - See what you can do")
	fmt.Println("  3. ellie automate quick - Set up automation")
	fmt.Println("  4. ellie start-day - Begin your workflow")
	fmt.Println()
	
	styles.GetSuccessStyle().Println("Welcome to the future of command-line productivity! 🚀")
	fmt.Println()
	
	ShowInspiringQuote()
}

func waitForUser() {
	fmt.Println()
	styles.GetInfoStyle().Print("Press Enter to continue...")
	fmt.Scanln()
}

// ImpressMe runs an impressive demo of all features
func ImpressMe() {
	ShowBanner()
	
	username := configs.GetEnv("USERNAME")
	if username == "" {
		username = "there"
	}
	
	styles.GetHighlightStyle().Printf("\n🎭 Impressive Features Demo for %s!\n", username)
	styles.GetInfoStyle().Println("═══════════════════════════════════════════════════════")
	fmt.Println()
	
	// Show system health
	styles.GetHighlightStyle().Println("1️⃣  System Health Monitoring")
	fmt.Println()
	QuickHealthCheck()
	
	fmt.Println()
	
	// Show smart suggestions
	styles.GetHighlightStyle().Println("2️⃣  Smart Context-Aware Suggestions")
	fmt.Println()
	TimeBasedSuggestions()
	
	fmt.Println()
	
	// Show available automations
	styles.GetHighlightStyle().Println("3️⃣  Automation Capabilities")
	fmt.Println()
	fmt.Println("Available automation options:")
	fmt.Println("  • Daily health checks")
	fmt.Println("  • Hourly git status")
	fmt.Println("  • Weekly system updates")
	fmt.Println("  • Custom scheduled tasks")
	styles.GetInfoStyle().Println("\n💡 Run 'ellie automate quick' to set them up!")
	
	fmt.Println()
	
	// Show Git capabilities
	if isGitRepo() {
		styles.GetHighlightStyle().Println("4️⃣  Git Intelligence")
		fmt.Println()
		
		if hasUncommittedChanges() {
			styles.GetWarningStyle().Println("  ⚠️  Uncommitted changes detected")
			fmt.Println("  💡 Use 'ellie git commit' for guided conventional commits")
		} else {
			styles.GetSuccessStyle().Println("  ✅ Repository is clean")
			fmt.Println("  💡 Use 'ellie git status' for detailed insights")
		}
		
		fmt.Println()
	}
	
	// Final message
	styles.GetHighlightStyle().Println("═══════════════════════════════════════════════════════")
	styles.GetSuccessStyle().Println("\n✨ Ellie is your personal command-line companion!")
	fmt.Println()
	fmt.Println("Explore more:")
	fmt.Println("  • ellie --help - See all commands")
	fmt.Println("  • ellie assist - Get context-aware help")
	fmt.Println("  • ellie showcase - Full feature tour")
	fmt.Println()
	
	ShowInspiringQuote()
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"sprouted.dev/weather"
)

func main() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	command := os.Args[1]
	
	switch command {
	case "weather":
		handleWeatherCommand(os.Args[2:])
	case "farm":
		handleFarmCommand(os.Args[2:])
	case "validate-seed":
		handleValidateSeedCommand(os.Args[2:])
	case "init":
		handleInitCommand(os.Args[2:])
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showUsage()
	}
}

func showUsage() {
	fmt.Println("Sprout CLI - Weather Context Preservation System")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  sprout weather                Show current weather context")
	fmt.Println("  sprout weather --for-ai       Show AI-friendly JSON context")
	fmt.Println("  sprout weather --onboard-ai   Show comprehensive AI assistant onboarding context")
	fmt.Println("  sprout weather --raw          Show raw weather context JSON")
	fmt.Println("  sprout weather recent         Show recent progress summary")
	fmt.Println("  sprout weather --suggest-docs  Show documentation suggestions")
	fmt.Println("  sprout weather emit-event     Emit event to farm orchestrator")
	fmt.Println("  sprout weather context-status Show context usage and handoff advice")
	fmt.Println("  sprout farm process           Process farm-level events")
	fmt.Println("  sprout farm weather           Show farm-level weather")
	fmt.Println("  sprout validate-seed [path]   Validate a documentation seed")
	fmt.Println("  sprout init --with-claude     Initialize workspace with Claude integration")
	fmt.Println()
}

func handleWeatherCommand(args []string) {
	// Find garden directory (look for .garden folder)
	gardenPath, err := findGardenPath()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("Make sure you're in a garden directory (contains .garden folder)")
		return
	}

	// Handle git monitoring commands first
	if len(args) > 0 {
		switch args[0] {
		case "--update-from-commit":
			if len(args) < 2 {
				fmt.Println("Error: --update-from-commit requires commit hash")
				return
			}
			handleUpdateFromCommit(gardenPath, args[1])
			return
		case "--update-from-branch-change":
			if len(args) < 4 {
				fmt.Println("Error: --update-from-branch-change requires prevHead newHead branchFlag")
				return
			}
			handleUpdateFromBranchChange(gardenPath, args[1], args[2], args[3])
			return
		case "--install-hooks":
			handleInstallHooks(gardenPath)
			return
		}
	}

	cm := weather.NewContextManager(gardenPath)
	context, err := cm.LoadContext()
	if err != nil {
		fmt.Printf("Error loading weather context: %v\n", err)
		return
	}

	// Parse command arguments
	if len(args) == 0 {
		showWeatherContext(context)
		return
	}

	switch args[0] {
	case "--for-ai":
		showAIContext(context)
	case "--onboard-ai":
		showAIOnboardingContext(context)
	case "--raw":
		showRawContext(context)
	case "recent":
		showRecentProgress(context)
	case "--suggest-docs":
		showDocumentationSuggestions(gardenPath, context)
	case "emit-event":
		if len(args) < 2 {
			fmt.Println("Error: emit-event requires event type")
			return
		}
		handleEmitEvent(gardenPath, args[1:])
	case "context-status":
		showContextStatus(gardenPath, context)
	default:
		fmt.Printf("Unknown weather option: %s\n", args[0])
		showUsage()
	}
}

func findGardenPath() (string, error) {
	// Start from current directory and walk up
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	for {
		gardenDir := filepath.Join(dir, ".garden")
		if _, err := os.Stat(gardenDir); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root directory
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("no garden found (no .garden directory)")
}

func showWeatherContext(context *weather.WeatherContext) {
	fmt.Println("🌦️  Current Development Weather")
	fmt.Println()
	
	// Current Focus
	fmt.Printf("🎯 Current Focus: %s", context.CurrentFocus.Area)
	if context.CurrentFocus.Confidence < 1.0 {
		fmt.Printf(" (%.0f%% confidence)", context.CurrentFocus.Confidence*100)
	}
	fmt.Println()
	
	// Recent Progress
	if context.RecentProgress.Summary != "" {
		fmt.Printf("📈 Recent Progress: %s", context.RecentProgress.Summary)
		if context.RecentProgress.Timespan != "" {
			fmt.Printf(" (%s)", context.RecentProgress.Timespan)
		}
		fmt.Println()
	}
	
	// Weather Conditions
	fmt.Printf("🌡️  %d°F | %s", context.Weather.Temperature, getWeatherEmoji(context.Weather.Condition))
	if context.Weather.Pressure > 70 {
		fmt.Printf(" | 🔴 High Pressure")
	} else if context.Weather.Pressure > 40 {
		fmt.Printf(" | 🟡 Medium Pressure")
	} else {
		fmt.Printf(" | 🟢 Low Pressure")
	}
	fmt.Println()
	
	// Git Context
	if context.Git.CurrentBranch != "" {
		fmt.Printf("🌿 Branch: %s", context.Git.CurrentBranch)
		if context.Git.UncommittedChanges {
			fmt.Printf(" (uncommitted changes)")
		}
		fmt.Println()
	}
	
	// Next Steps
	if len(context.NextSteps.Suggestions) > 0 {
		fmt.Println()
		fmt.Println("⚡ Next Steps:")
		for i, step := range context.NextSteps.Suggestions {
			if i >= 3 { // Limit to top 3 suggestions
				break
			}
			fmt.Printf("   %d. %s\n", i+1, step)
		}
	}
	
	fmt.Println()
	fmt.Printf("Last updated: %s\n", context.Updated.Format("2006-01-02 15:04:05"))
}

func showAIContext(context *weather.WeatherContext) {
	aiContext := context.ToAIContext()
	jsonData, err := json.MarshalIndent(aiContext, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting AI context: %v\n", err)
		return
	}
	
	fmt.Println("AI-Friendly Weather Context:")
	fmt.Println(string(jsonData))
}

func showAIOnboardingContext(context *weather.WeatherContext) {
	onboardingContext := context.ToAIOnboardingContext()
	jsonData, err := json.MarshalIndent(onboardingContext, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting AI onboarding context: %v\n", err)
		return
	}
	
	fmt.Println("Comprehensive AI Assistant Onboarding Context:")
	fmt.Println(string(jsonData))
}

func showRawContext(context *weather.WeatherContext) {
	jsonData, err := context.ToJSON()
	if err != nil {
		fmt.Printf("Error formatting raw context: %v\n", err)
		return
	}
	
	fmt.Println("Raw Weather Context:")
	fmt.Println(string(jsonData))
}

func showRecentProgress(context *weather.WeatherContext) {
	fmt.Println("📈 Recent Development Progress")
	fmt.Println()
	
	if context.RecentProgress.Summary != "" {
		fmt.Printf("Summary: %s\n", context.RecentProgress.Summary)
		if context.RecentProgress.Timespan != "" {
			fmt.Printf("Timespan: %s\n", context.RecentProgress.Timespan)
		}
		fmt.Printf("Momentum: %d/100\n", context.RecentProgress.Momentum)
	} else {
		fmt.Println("No recent progress tracked yet.")
		fmt.Println("Make some commits to start seeing progress summaries!")
	}
	
	if len(context.RecentProgress.Commits) > 0 {
		fmt.Println()
		fmt.Println("Recent Commits:")
		for i, commit := range context.RecentProgress.Commits {
			if i >= 5 { // Limit to last 5 commits
				break
			}
			fmt.Printf("  • %s", commit.Message)
			if commit.SmartSummary != "" && commit.SmartSummary != commit.Message {
				fmt.Printf(" (%s)", commit.SmartSummary)
			}
			fmt.Println()
		}
	}
}

func getWeatherEmoji(condition weather.WeatherCondition) string {
	switch condition {
	case weather.WeatherSunny:
		return "☀️ Sunny"
	case weather.WeatherPartlyCloudy:
		return "⛅ Partly Cloudy"
	case weather.WeatherCloudy:
		return "☁️ Cloudy"
	case weather.WeatherStormy:
		return "⛈️ Stormy"
	case weather.WeatherFoggy:
		return "🌫️ Foggy"
	default:
		return "🌤️ " + strings.Title(string(condition))
	}
}

// Git monitoring handlers

func handleUpdateFromCommit(gardenPath, commitHash string) {
	gm := weather.NewGitMonitor(gardenPath)
	if err := gm.UpdateFromCommit(commitHash); err != nil {
		fmt.Printf("Error updating weather from commit: %v\n", err)
		return
	}
	// Display first 8 characters if hash is long enough, otherwise show full hash
	displayHash := commitHash
	if len(commitHash) > 8 {
		displayHash = commitHash[:8]
	}
	fmt.Printf("🌦️ Weather updated from commit %s\n", displayHash)
}

func handleUpdateFromBranchChange(gardenPath, prevHead, newHead, branchFlag string) {
	gm := weather.NewGitMonitor(gardenPath)
	if err := gm.UpdateFromBranchChange(prevHead, newHead, branchFlag); err != nil {
		fmt.Printf("Error updating weather from branch change: %v\n", err)
		return
	}
	fmt.Println("🌦️ Weather updated from branch change")
}

func handleInstallHooks(gardenPath string) {
	gm := weather.NewGitMonitor(gardenPath)
	if err := gm.InstallGitHooks(); err != nil {
		fmt.Printf("Error installing git hooks: %v\n", err)
		return
	}
	fmt.Println("✅ Git hooks installed successfully")
	fmt.Println("Weather will now update automatically on commits and branch changes!")
}

// New handlers for documentation suggestions and event emission

func showDocumentationSuggestions(gardenPath string, context *weather.WeatherContext) {
	w := &weather.Weather{
		RepoPath: gardenPath,
		Context:  context,
	}
	
	suggestions := w.DetectDocumentationGaps()
	
	fmt.Println("🌦️  Documentation Intelligence")
	fmt.Println()
	
	if len(suggestions.MissingDecisions) == 0 && 
	   len(suggestions.UncapturedLessons) == 0 && 
	   len(suggestions.ProcessGaps) == 0 {
		fmt.Println("✅ No documentation gaps detected")
		return
	}
	
	fmt.Println("📋 Missing Documentation Detected:")
	fmt.Println()
	
	// Show missing decisions
	for _, need := range suggestions.MissingDecisions {
		showDocumentationNeed(need, "🔍 MISSING DECISION")
	}
	
	// Show uncaptured lessons
	for _, need := range suggestions.UncapturedLessons {
		showDocumentationNeed(need, "🚨 CRITICAL LESSON")
	}
	
	// Show process gaps
	for _, need := range suggestions.ProcessGaps {
		showDocumentationNeed(need, "📚 PROCESS GAP")
	}
	
	if len(suggestions.RecentActivities) > 0 {
		fmt.Println()
		fmt.Println("⏰ RECENT ACTIVITY SUGGESTIONS:")
		for _, activity := range suggestions.RecentActivities {
			fmt.Printf("   • %s (%s) → Should document as %s\n", 
				activity.Activity, activity.Timeframe, activity.SuggestedDocType)
		}
	}
}

func showDocumentationNeed(need weather.DocumentationNeed, prefix string) {
	fmt.Printf("%s (Confidence: %.0f%%)\n", prefix, need.Confidence*100)
	fmt.Printf("   Title: \"%s\"\n", need.Title)
	fmt.Printf("   Type: %s\n", need.Type)
	fmt.Printf("   Location: %s\n", need.SuggestedLocation)
	fmt.Printf("   Detected: %s\n", strings.Join(need.DetectedFrom, ", "))
	fmt.Println()
}

func handleEmitEvent(gardenPath string, args []string) {
	w := &weather.Weather{RepoPath: gardenPath}
	emitter, err := weather.NewEventEmitter(w)
	if err != nil {
		fmt.Printf("Error creating event emitter: %v\n", err)
		return
	}
	
	eventType := weather.EventType(args[0])
	payload := make(map[string]interface{})
	
	// Parse additional arguments as key=value pairs
	for i := 1; i < len(args); i++ {
		parts := strings.SplitN(args[i], "=", 2)
		if len(parts) == 2 {
			payload[parts[0]] = parts[1]
		}
	}
	
	if err := emitter.Emit(eventType, payload); err != nil {
		fmt.Printf("Error emitting event: %v\n", err)
		return
	}
	
	fmt.Printf("✅ Event emitted: %s\n", eventType)
}

func handleFarmCommand(args []string) {
	if len(args) == 0 {
		fmt.Println("Farm commands:")
		fmt.Println("  sprout farm process    Process pending events")
		fmt.Println("  sprout farm weather    Show farm-level weather")
		return
	}
	
	// Find farm path (parent of garden)
	gardenPath, err := findGardenPath()
	if err != nil {
		// Not in a garden, try current directory as farm
		gardenPath = "."
	}
	farmPath := filepath.Dir(gardenPath)
	
	switch args[0] {
	case "process":
		handleFarmProcess(farmPath)
	case "weather":
		handleFarmWeather(farmPath)
	default:
		fmt.Printf("Unknown farm command: %s\n", args[0])
	}
}

func handleFarmProcess(farmPath string) {
	fmt.Printf("Processing events in farm: %s\n", farmPath)
	orchestrator := weather.NewFarmOrchestrator(farmPath)
	if err := orchestrator.ProcessEvents(); err != nil {
		fmt.Printf("Error processing events: %v\n", err)
		return
	}
	
	fmt.Println("✅ Farm events processed")
}

func handleFarmWeather(farmPath string) {
	weatherPath := filepath.Join(farmPath, ".farm", "weather", "current.json")
	data, err := os.ReadFile(weatherPath)
	if err != nil {
		fmt.Printf("No farm weather data found. Run 'sprout farm process' first.\n")
		return
	}
	
	var farmWeather weather.FarmWeather
	if err := json.Unmarshal(data, &farmWeather); err != nil {
		fmt.Printf("Error reading farm weather: %v\n", err)
		return
	}
	
	fmt.Println("🌦️  Farm-Level Weather")
	fmt.Println()
	fmt.Printf("🌡️  Overall: %.0f°F\n", farmWeather.OverallTemp)
	fmt.Printf("🌿 Active Gardens: %s\n", strings.Join(farmWeather.ActiveGardens, ", "))
	fmt.Println()
	
	if len(farmWeather.Correlations) > 0 {
		fmt.Println("🔗 Cross-Garden Patterns:")
		for _, corr := range farmWeather.Correlations {
			fmt.Printf("   • %s: %s (%.0f%% confidence)\n", 
				strings.Join(corr.Gardens, " + "), 
				corr.Pattern, 
				corr.Confidence*100)
			if corr.Suggestion != "" {
				fmt.Printf("     💡 %s\n", corr.Suggestion)
			}
		}
		fmt.Println()
	}
	
	if len(farmWeather.Suggestions) > 0 {
		fmt.Println("📋 Documentation Suggestions:")
		for _, sugg := range farmWeather.Suggestions {
			fmt.Printf("   • %s: %s\n", sugg.Type, sugg.Title)
		}
	}
}

func handleValidateSeedCommand(args []string) {
	// Get path to validate (default to current directory)
	path := "."
	if len(args) > 0 {
		path = args[0]
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Error resolving path: %v\n", err)
		return
	}

	fmt.Printf("🌱 Validating Seed at: %s\n\n", absPath)

	// Create validator
	validator := weather.NewSeedValidator(absPath)
	
	// Run validation
	result, err := validator.Validate()
	if err != nil {
		fmt.Printf("Error during validation: %v\n", err)
		return
	}

	// Display results
	fmt.Printf("📊 Validation Results\n")
	fmt.Printf("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
	fmt.Printf("Level: %s\n", strings.Title(result.Level))
	fmt.Printf("Valid: %s\n", formatBool(result.Valid))
	fmt.Printf("Score: %d/100\n", result.Score)
	fmt.Println()

	// Display structure found
	fmt.Println("📁 Structure Found:")
	fmt.Printf("   • Has docs/: %s\n", formatBool(result.Structure.HasDocs))
	fmt.Printf("   • Has README: %s\n", formatBool(result.Structure.HasReadme))
	fmt.Printf("   • Has .git: %s\n", formatBool(result.Structure.HasGitRepo))
	fmt.Printf("   • Weather initialized: %s\n", formatBool(result.Structure.HasWeatherContext))
	
	if len(result.Structure.Directories) > 0 {
		fmt.Printf("   • Subdirectories: %s\n", strings.Join(result.Structure.Directories, ", "))
	}
	if len(result.Structure.SpecialFiles) > 0 {
		fmt.Printf("   • Special files: %s\n", strings.Join(result.Structure.SpecialFiles, ", "))
	}
	fmt.Println()

	// Display issues
	if len(result.Issues) > 0 {
		fmt.Println("⚠️  Issues Found:")
		for _, issue := range result.Issues {
			icon := "ℹ️"
			if issue.Severity == "error" {
				icon = "❌"
			} else if issue.Severity == "warning" {
				icon = "⚠️"
			}
			fmt.Printf("   %s %s: %s\n", icon, issue.Path, issue.Message)
			if issue.Suggestion != "" {
				fmt.Printf("      💡 %s\n", issue.Suggestion)
			}
		}
		fmt.Println()
	}

	// Display suggestions
	if len(result.Suggestions) > 0 {
		fmt.Println("💡 Suggestions:")
		for _, suggestion := range result.Suggestions {
			fmt.Printf("   • %s\n", suggestion)
		}
		fmt.Println()
	}

	// Summary
	if result.Valid {
		if result.Score >= 80 {
			fmt.Println("✅ Excellent! Your seed is well-structured and ready to grow.")
		} else if result.Score >= 60 {
			fmt.Println("✅ Good seed structure. Consider the suggestions to improve further.")
		} else {
			fmt.Println("✅ Valid seed, but could be improved. Check suggestions above.")
		}
	} else {
		fmt.Println("❌ Seed has critical issues that should be addressed.")
	}
}

func formatBool(b bool) string {
	if b {
		return "✅ Yes"
	}
	return "❌ No"
}

func showContextStatus(gardenPath string, context *weather.WeatherContext) {
	// Create a monitor (in real usage, this would be persistent)
	monitor := weather.NewContextMonitor()
	
	// Create weather instance for analysis
	w := &weather.Weather{
		RepoPath: gardenPath,
		Context:  context,
	}
	
	// Get status
	status, err := monitor.GetStatus(w)
	if err != nil {
		fmt.Printf("Error checking context status: %v\n", err)
		return
	}
	
	fmt.Println("🤖 Claude Context Status")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	
	// Show usage with visual bar
	fmt.Printf("Usage: %d%% ", status.UsagePercent)
	showUsageBar(status.UsagePercent)
	fmt.Println()
	
	// Show urgency
	urgencyIcon := "✅"
	urgencyColor := ""
	switch status.Urgency {
	case "critical":
		urgencyIcon = "🚨"
		urgencyColor = "\033[0;31m" // Red
	case "high":
		urgencyIcon = "⚠️"
		urgencyColor = "\033[1;33m" // Yellow
	case "medium":
		urgencyIcon = "📊"
		urgencyColor = "\033[1;33m" // Yellow
	default:
		urgencyIcon = "✅"
		urgencyColor = "\033[0;32m" // Green
	}
	
	fmt.Printf("%sUrgency: %s %s\033[0m\n", urgencyColor, urgencyIcon, status.Urgency)
	
	// Show handoff analysis
	if status.IsLogicalHandoff {
		fmt.Println("\n✅ This is a good handoff point!")
		fmt.Println("Reasons:")
		for _, reason := range status.HandoffReasons {
			fmt.Printf("   • %s\n", reason)
		}
	} else {
		fmt.Println("\n⏳ Not an ideal handoff point yet")
	}
	
	// Show recommendation
	fmt.Printf("\n💡 Recommendation: %s\n", status.Recommendation)
	
	// Show handoff command if appropriate
	if status.UsagePercent >= 70 || status.IsLogicalHandoff {
		fmt.Println("\n📋 To prepare handoff:")
		fmt.Println("   .claude/commands/onboard-next-assistant")
	}
}

func showUsageBar(percent int) {
	barWidth := 40
	filled := (percent * barWidth) / 100
	
	fmt.Print("[")
	for i := 0; i < barWidth; i++ {
		if i < filled {
			if percent >= 90 {
				fmt.Print("\033[0;31m█\033[0m") // Red
			} else if percent >= 80 {
				fmt.Print("\033[1;33m█\033[0m") // Yellow
			} else {
				fmt.Print("\033[0;32m█\033[0m") // Green
			}
		} else {
			fmt.Print("░")
		}
	}
	fmt.Print("]")
}

func handleInitCommand(args []string) {
	// Check for --with-claude flag
	withClaude := false
	for _, arg := range args {
		if arg == "--with-claude" {
			withClaude = true
			break
		}
	}
	
	if !withClaude {
		fmt.Println("Usage: sprout init --with-claude")
		fmt.Println("This initializes Claude AI integration in your workspace")
		return
	}
	
	// Get current directory
	workspacePath, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}
	
	fmt.Println("🤖 Initializing Claude integration...")
	
	// Create integration
	integration := weather.NewClaudeIntegration(workspacePath)
	
	// Run setup
	if err := integration.Setup(); err != nil {
		fmt.Printf("Error setting up Claude integration: %v\n", err)
		return
	}
	
	// Detect workspace type and show appropriate message
	fmt.Println()
	fmt.Printf("✅ Claude integration installed for %s workspace!\n", integration.WorkspaceType)
	fmt.Println()
	fmt.Println("🚀 Quick Start:")
	fmt.Println("   1. Check context: .claude/commands/context-monitor check")
	fmt.Println("   2. Start work as normal")
	fmt.Println("   3. Monitor will alert when handoff is recommended")
	fmt.Println()
	fmt.Println("📋 Commands available:")
	fmt.Println("   • sprout weather context-status - Check context usage")
	fmt.Println("   • .claude/commands/onboard-next-assistant - Prepare handoff")
	fmt.Println("   • .claude/commands/context-monitor monitor - Run background monitor")
}
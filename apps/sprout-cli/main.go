package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
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
	case "seed":
		handleSeedCommand(os.Args[2:])
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
	fmt.Println("  sprout weather --onboard-ai --include-usage-context  Enhanced onboarding for cold starts")
	fmt.Println("  sprout weather --raw          Show raw weather context JSON")
	fmt.Println("  sprout weather recent         Show recent progress summary")
	fmt.Println("  sprout weather --suggest-docs  Show documentation suggestions")
	fmt.Println("  sprout weather emit-event     Emit event to farm orchestrator")
	fmt.Println("  sprout weather context-status Show context usage and handoff advice")
	fmt.Println("  sprout weather --prepare-cold-handoff  Prepare for usage limit interruption")
	fmt.Println("  sprout weather verify         Verify weather context integrity")
	fmt.Println("  sprout weather recover        Recover from corrupted weather context")
	fmt.Println("  sprout weather backups        List available weather backups")
	fmt.Println("  sprout farm process           Process farm-level events")
	fmt.Println("  sprout farm weather           Show farm-level weather")
	fmt.Println("  sprout farm backup            Create Farm-level backup")
	fmt.Println("  sprout farm protection-status Check Farm protection status")
	fmt.Println("  sprout seed <name>            Create a new project seed with documentation structure")
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
		includeUsageContext := false
		if len(args) > 1 && args[1] == "--include-usage-context" {
			includeUsageContext = true
		}
		showAIOnboardingContext(context, includeUsageContext)
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
	case "--prepare-cold-handoff":
		prepareColdHandoff(gardenPath, context)
	case "verify":
		handleWeatherVerify(gardenPath)
	case "recover":
		handleWeatherRecover(gardenPath, args[1:])
	case "backups":
		handleWeatherBackups(gardenPath)
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

func showAIOnboardingContext(context *weather.WeatherContext, includeUsageContext bool) {
	onboardingContext := context.ToAIOnboardingContext()
	
	if includeUsageContext {
		// Add usage limit context information
		usageInfo := map[string]any{
			"session_type": "cold_start_after_usage_limit",
			"interruption_reason": "User hit Claude usage limits (daily/monthly)",
			"continuity_notes": "This is a completely new assistant. Previous assistant lost all memory.",
			"onboarding_importance": "CRITICAL - This assistant has zero context from previous sessions",
		}
		onboardingContext["usage_limit_context"] = usageInfo
	}
	
	// Display temporal warning prominently
	fmt.Println("⚠️  ⚠️  ⚠️  TEMPORAL BOUNDARIES WARNING ⚠️  ⚠️  ⚠️")
	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│ THIS PROJECT STARTED ON MAY 21, 2025 (3 days ago)          │")
	fmt.Println("│ ANY DATE BEFORE MAY 21, 2025 IS IMPOSSIBLE                 │")
	fmt.Println("│ DO NOT CREATE FICTIONAL TIMELINES OR BACKSTORIES           │")
	fmt.Println("└─────────────────────────────────────────────────────────────┘")
	fmt.Println()
	
	jsonData, err := json.MarshalIndent(onboardingContext, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting AI onboarding context: %v\n", err)
		return
	}
	
	if includeUsageContext {
		fmt.Println("🆕 NEW ASSISTANT COLD START - Full Onboarding Context:")
		fmt.Println("(Previous session ended due to usage limits)")
	} else {
		fmt.Println("Comprehensive AI Assistant Onboarding Context:")
	}
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
		// Capitalize first letter manually to avoid deprecated strings.Title
		s := string(condition)
		if len(s) > 0 {
			return "🌤️ " + strings.ToUpper(s[:1]) + s[1:]
		}
		return "🌤️ " + s
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
	payload := make(map[string]any)
	
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
	case "backup":
		handleFarmBackup()
	case "protection-status":
		handleFarmProtectionStatus()
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
	
	fmt.Println("🚧 Seed validation coming soon!")
	fmt.Println()
	fmt.Println("This feature will:")
	fmt.Println("  • Check documentation structure")
	fmt.Println("  • Validate seed health")
	fmt.Println("  • Provide improvement suggestions")
	fmt.Println("  • Score your seed quality")
	fmt.Println()
	fmt.Println("For now, check out the Seeds quickstart guide:")
	fmt.Println("  https://github.com/sprouted-dev/garden/blob/main/docs/seeds/quickstart.md")
}

func showContextStatus(_ string, _ *weather.WeatherContext) {
	// TODO: Implement context status monitoring
	fmt.Println("🚧 Context status monitoring coming soon!")
	fmt.Println()
	fmt.Println("This feature will allow you to:")
	fmt.Println("  • Monitor Claude's context usage")
	fmt.Println("  • Get alerts before hitting limits")
	fmt.Println("  • Prepare seamless handoffs")
}

func prepareColdHandoff(gardenPath string, context *weather.WeatherContext) {
	fmt.Println("🔄 Preparing for Cold Handoff (Usage Limit Interruption)")
	fmt.Println()
	
	// Save enhanced context for cold start
	cm := weather.NewContextManager(gardenPath)
	if err := cm.SaveContext(context); err != nil {
		fmt.Printf("⚠️  Warning: Could not save context: %v\n", err)
	} else {
		fmt.Println("✅ Current state preserved")
	}
	
	// Display comprehensive handoff information
	fmt.Println("📋 Session Summary for Next Assistant:")
	fmt.Println()
	
	// Current focus
	fmt.Printf("🎯 Current Focus: %s", context.CurrentFocus.Area)
	if context.CurrentFocus.Confidence < 1.0 {
		fmt.Printf(" (%.0f%% confidence)", context.CurrentFocus.Confidence*100)
	}
	fmt.Println()
	
	// Recent progress
	if context.RecentProgress.Summary != "" {
		fmt.Printf("📈 Recent Progress: %s", context.RecentProgress.Summary)
		if context.RecentProgress.Timespan != "" {
			fmt.Printf(" (%s)", context.RecentProgress.Timespan)
		}
		fmt.Println()
	}
	
	// Next steps
	if len(context.NextSteps.Suggestions) > 0 {
		fmt.Println("⚡ Next Steps:")
		for i, step := range context.NextSteps.Suggestions {
			if i >= 5 { // Limit to top 5 suggestions
				break
			}
			fmt.Printf("   %d. %s\n", i+1, step)
		}
	}
	
	fmt.Println()
	fmt.Println("🚨 USAGE LIMIT INTERRUPTION DETECTED")
	fmt.Println()
	fmt.Println("When you return with a new assistant:")
	fmt.Printf("1. Run: cd %s\n", gardenPath)
	fmt.Println("2. Run: sprout weather --onboard-ai")
	fmt.Println("3. New assistant will have full context")
	fmt.Println()
	fmt.Println("💡 Tip: Usage limits reset daily/monthly depending on your plan")
}


func handleInitCommand(args []string) {
	// Check for --with-claude flag
	withClaude := slices.Contains(args, "--with-claude")
	
	if !withClaude {
		fmt.Println("Usage: sprout init --with-claude")
		fmt.Println("This initializes Claude AI integration in your workspace")
		return
	}
	
	fmt.Println("🚧 Claude integration coming soon!")
	fmt.Println()
	fmt.Println("This feature will provide:")
	fmt.Println("  • Automatic context monitoring")
	fmt.Println("  • Smart handoff detection")
	fmt.Println("  • Seamless session continuity")
	fmt.Println()
	fmt.Println("Stay tuned for updates!")
}

func handleSeedCommand(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: sprout seed <name>")
		fmt.Println("Create a new project seed with documentation structure")
		return
	}

	name := args[0]
	if err := createSeed(name); err != nil {
		fmt.Printf("Error creating seed: %v\n", err)
		return
	}
}

func handleWeatherVerify(gardenPath string) {
	shadowManager := weather.NewShadowCopyManager(gardenPath)
	
	fmt.Println("🔍 Verifying Weather Context Integrity...")
	fmt.Println()
	
	if err := shadowManager.VerifyContextIntegrity(); err != nil {
		fmt.Printf("❌ Context verification failed: %v\n", err)
		fmt.Println()
		fmt.Println("💡 Run 'sprout weather recover' to restore from backup")
		return
	}
	
	fmt.Println("✅ Weather context is valid and intact")
	
	// Also check shadow copy
	fmt.Print("🔍 Checking shadow copy... ")
	shadowPath := filepath.Join(gardenPath, ".garden", "weather-context.shadow.json")
	if _, err := os.Stat(shadowPath); os.IsNotExist(err) {
		fmt.Println("⚠️  No shadow copy found (will be created on next update)")
	} else {
		fmt.Println("✅ Shadow copy exists")
	}
	
	// Check backups
	backups, err := shadowManager.GetBackupList()
	if err != nil {
		fmt.Printf("⚠️  Could not check backups: %v\n", err)
	} else {
		fmt.Printf("📦 Available backups: %d\n", len(backups))
	}
}

func handleWeatherRecover(gardenPath string, args []string) {
	shadowManager := weather.NewShadowCopyManager(gardenPath)
	
	// If specific backup requested
	if len(args) > 0 && args[0] != "" {
		backupName := args[0]
		fmt.Printf("🔄 Recovering from backup: %s\n", backupName)
		
		if err := shadowManager.RestoreFromBackup(backupName); err != nil {
			fmt.Printf("❌ Recovery failed: %v\n", err)
			return
		}
		
		fmt.Println("✅ Successfully recovered from backup")
		return
	}
	
	// Try shadow copy first
	fmt.Println("🔄 Attempting Weather Context Recovery...")
	fmt.Println()
	
	fmt.Println("1️⃣ Trying shadow copy...")
	if err := shadowManager.RestoreFromShadow(); err != nil {
		fmt.Printf("⚠️  Shadow recovery failed: %v\n", err)
		fmt.Println()
		
		// List available backups
		fmt.Println("2️⃣ Checking available backups...")
		backups, err := shadowManager.GetBackupList()
		if err != nil || len(backups) == 0 {
			fmt.Println("❌ No backups available")
			fmt.Println()
			fmt.Println("💡 You may need to:")
			fmt.Println("   1. Check git history for weather-context.json")
			fmt.Println("   2. Run 'sprout weather' to regenerate from git state")
			return
		}
		
		fmt.Println("📦 Available backups:")
		for i, backup := range backups {
			fmt.Printf("   %d. %s\n", i+1, backup)
		}
		fmt.Println()
		fmt.Printf("Run: sprout weather recover %s\n", backups[len(backups)-1])
		return
	}
	
	fmt.Println("✅ Successfully recovered from shadow copy")
	
	// Verify recovered context
	if err := shadowManager.VerifyContextIntegrity(); err != nil {
		fmt.Printf("⚠️  Recovered context may have issues: %v\n", err)
	} else {
		fmt.Println("✅ Recovered context verified successfully")
	}
}

func handleWeatherBackups(gardenPath string) {
	shadowManager := weather.NewShadowCopyManager(gardenPath)
	
	fmt.Println("📦 Weather Context Backups")
	fmt.Println()
	
	backups, err := shadowManager.GetBackupList()
	if err != nil {
		fmt.Printf("❌ Error listing backups: %v\n", err)
		return
	}
	
	if len(backups) == 0 {
		fmt.Println("No backups found.")
		fmt.Println()
		fmt.Println("💡 Backups are created automatically when:")
		fmt.Println("   • Weather context is updated")
		fmt.Println("   • Shadow copies are created")
		fmt.Println("   • Recovery operations are performed")
		return
	}
	
	fmt.Printf("Found %d backup(s):\n", len(backups))
	fmt.Println()
	
	for _, backup := range backups {
		// Parse timestamp from filename
		fmt.Printf("  📄 %s", backup)
		
		// Check file size
		backupPath := filepath.Join(gardenPath, ".garden", "backups", backup)
		if info, err := os.Stat(backupPath); err == nil {
			fmt.Printf(" (%.1f KB)", float64(info.Size())/1024)
		}
		fmt.Println()
	}
	
	fmt.Println()
	fmt.Println("💡 To recover from a specific backup:")
	fmt.Printf("   sprout weather recover %s\n", backups[len(backups)-1])
}

func handleFarmBackup() {
	fpm, err := NewFarmProtectionManager()
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	
	fmt.Println("🚜 Creating Farm-level backup...")
	if err := fpm.BackupFarm(); err != nil {
		fmt.Printf("❌ Backup failed: %v\n", err)
		return
	}
}

func handleFarmProtectionStatus() {
	fpm, err := NewFarmProtectionManager()
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}
	
	fpm.CheckFarmProtection()
}
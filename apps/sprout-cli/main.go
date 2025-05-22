package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sprouted-dev/garden/libs/weather"
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
	fmt.Println("  sprout weather --raw          Show raw weather context JSON")
	fmt.Println("  sprout weather recent         Show recent progress summary")
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
	case "--raw":
		showRawContext(context)
	case "recent":
		showRecentProgress(context)
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
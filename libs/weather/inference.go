package weather

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// DefaultInferenceEngine implements smart analysis and inference
type DefaultInferenceEngine struct{}

// NewDefaultInferenceEngine creates a new inference engine
func NewDefaultInferenceEngine() *DefaultInferenceEngine {
	return &DefaultInferenceEngine{}
}

// InferScope determines the scope/area of work from changed files
func (ie *DefaultInferenceEngine) InferScope(files []string) string {
	if len(files) == 0 {
		return "general"
	}
	
	// Count file patterns to determine primary scope
	scopeCounts := make(map[string]int)
	
	for _, file := range files {
		file = strings.ToLower(file)
		
		// Categorize by file patterns - check test files first since they're more specific
		switch {
		case strings.Contains(file, "test") || strings.HasSuffix(file, "_test.go") || strings.HasSuffix(file, ".test."):
			scopeCounts["testing"]++
		case strings.Contains(file, "auth"):
			scopeCounts["authentication"]++
		case strings.Contains(file, "api") || strings.Contains(file, "handler") || strings.Contains(file, "route"):
			scopeCounts["api"]++
		case strings.Contains(file, "ui") || strings.Contains(file, "component") || strings.Contains(file, "frontend"):
			scopeCounts["frontend"]++
		case strings.Contains(file, "doc") || strings.HasSuffix(file, ".md"):
			scopeCounts["documentation"]++
		case strings.Contains(file, "config") || strings.Contains(file, "setting"):
			scopeCounts["configuration"]++
		case strings.Contains(file, "db") || strings.Contains(file, "database") || strings.Contains(file, "migration"):
			scopeCounts["database"]++
		default:
			// Try to infer from directory structure
			parts := strings.Split(file, "/")
			if len(parts) > 1 {
				scopeCounts[parts[0]]++
			} else {
				scopeCounts["general"]++
			}
		}
	}
	
	// Find the most common scope
	maxCount := 0
	primaryScope := "general"
	for scope, count := range scopeCounts {
		if count > maxCount {
			maxCount = count
			primaryScope = scope
		}
	}
	
	return primaryScope
}

// GenerateSmartSummary creates human-readable summary from commit message
func (ie *DefaultInferenceEngine) GenerateSmartSummary(message string) string {
	// Remove common prefixes and technical jargon
	summary := message
	
	// Handle conventional commit format
	conventionalRegex := regexp.MustCompile(`^(feat|fix|docs|style|refactor|test|chore)(\(.+\))?\s*:\s*(.+)`)
	if matches := conventionalRegex.FindStringSubmatch(summary); len(matches) >= 4 {
		commitType := matches[1]
		scope := matches[2]
		description := matches[3]
		
		switch commitType {
		case "feat":
			summary = fmt.Sprintf("Added %s", description)
		case "fix":
			summary = fmt.Sprintf("Fixed %s", description)  
		case "refactor":
			summary = fmt.Sprintf("Refactored %s", description)
		case "docs":
			summary = fmt.Sprintf("Updated documentation for %s", description)
		case "test":
			summary = fmt.Sprintf("Added tests for %s", description)
		case "chore":
			summary = fmt.Sprintf("Updated %s", description)
		default:
			summary = description
		}
		
		if scope != "" {
			scopeName := strings.Trim(scope, "()")
			summary = fmt.Sprintf("%s (%s)", summary, scopeName)
		}
	}
	
	// Capitalize first letter
	if len(summary) > 0 {
		summary = strings.ToUpper(string(summary[0])) + summary[1:]
	}
	
	// Limit length
	if len(summary) > 100 {
		summary = summary[:100-3] + "..."
	}
	
	return summary
}

// InferFocusArea determines focus area from commit and branch context  
func (ie *DefaultInferenceEngine) InferFocusArea(commit GitCommit, currentBranch string) FocusArea {
	scope := ie.InferScope(commit.FilesChanged)
	
	// Generate focus area description
	var focusArea string
	var confidence float64 = 0.85 // High confidence from commit activity
	
	// Check branch context first
	if strings.Contains(currentBranch, "/") {
		parts := strings.Split(currentBranch, "/")
		if len(parts) >= 2 {
			branchType := parts[0]
			branchName := strings.Join(parts[1:], "/")
			
			switch branchType {
			case "feature", "feat":
				focusArea = fmt.Sprintf("feature: %s", branchName)
			case "fix", "bugfix", "hotfix":
				focusArea = fmt.Sprintf("bug fix: %s", branchName)
			case "refactor":
				focusArea = fmt.Sprintf("refactoring: %s", branchName)
			default:
				focusArea = fmt.Sprintf("working on: %s", branchName)
			}
			confidence = 0.9 // Higher confidence from branch naming
		}
	}
	
	// Fall back to scope-based inference
	if focusArea == "" {
		switch scope {
		case "authentication":
			focusArea = "authentication system"
		case "api":
			focusArea = "API development"
		case "frontend":
			focusArea = "user interface"
		case "testing":
			focusArea = "testing and quality assurance"
		case "documentation":
			focusArea = "documentation"
		case "database":
			focusArea = "database layer"
		case "configuration":
			focusArea = "system configuration"
		default:
			focusArea = fmt.Sprintf("%s development", scope)
		}
	}
	
	return FocusArea{
		Area:         focusArea,
		Confidence:   confidence,
		LastActive:   time.Now(),
		InferredFrom: fmt.Sprintf("recent commit to %s", scope),
	}
}

// CalculateWeatherConditions determines weather based on commit patterns
func (ie *DefaultInferenceEngine) CalculateWeatherConditions(commits []GitCommit) WeatherConditions {
	now := time.Now()
	
	if len(commits) == 0 {
		return WeatherConditions{
			Temperature: 20,
			Condition:   WeatherSunny,
			Pressure:    30,
			LastUpdate:  now,
		}
	}
	
	// Calculate momentum based on recent commits
	recentCommitCount := 0
	for _, commit := range commits {
		if now.Sub(commit.Timestamp) < 2*time.Hour {
			recentCommitCount++
		}
	}
	momentum := minInt(recentCommitCount*25, 100)
	
	// Temperature based on momentum
	temperature := minInt(20+momentum, 95)
	
	// Analyze commit messages for condition indicators
	fixCount := 0
	featCount := 0
	for _, commit := range commits {
		msg := strings.ToLower(commit.Message)
		if strings.Contains(msg, "fix") || strings.Contains(msg, "bug") {
			fixCount++
		}
		if strings.Contains(msg, "feat") || strings.Contains(msg, "add") {
			featCount++
		}
	}
	
	// Determine weather condition
	var condition WeatherCondition
	totalCommits := len(commits)
	if fixCount > totalCommits/2 {
		condition = WeatherCloudy // Lots of fixes = some challenges
	} else if featCount > totalCommits/2 {
		condition = WeatherSunny // Lots of features = smooth progress
	} else {
		condition = WeatherPartlyCloudy // Mixed activity
	}
	
	return WeatherConditions{
		Temperature: temperature,
		Condition:   condition,
		Pressure:    50, // Default pressure, could be enhanced with deadline detection
		LastUpdate:  now,
	}
}

// GenerateNextSteps creates actionable suggestions based on context
func (ie *DefaultInferenceEngine) GenerateNextSteps(focus FocusArea, progress ProgressSummary) NextStepsSuggestion {
	var suggestions []string
	
	// Base suggestions on current focus area
	scope := strings.ToLower(focus.Area)
	
	switch {
	case strings.Contains(scope, "authentication"):
		suggestions = []string{
			"Add JWT token validation",
			"Implement user session management", 
			"Add password reset functionality",
		}
	case strings.Contains(scope, "api"):
		suggestions = []string{
			"Add API endpoint validation",
			"Implement error handling middleware",
			"Add API documentation",
		}
	case strings.Contains(scope, "user interface") || strings.Contains(scope, "frontend"):
		suggestions = []string{
			"Add responsive styling",
			"Implement form validation",
			"Add loading states and error handling",
		}
	case strings.Contains(scope, "testing"):
		suggestions = []string{
			"Add integration tests",
			"Improve test coverage",
			"Add end-to-end testing",
		}
	default:
		suggestions = []string{
			"Add unit tests for recent changes",
			"Update documentation",
			"Consider refactoring for better maintainability",
		}
	}
	
	priority := minInt(50+progress.Momentum/2, 90)
	
	return NextStepsSuggestion{
		Suggestions: suggestions,
		Priority:    priority,
		BasedOn:     fmt.Sprintf("current focus on %s", focus.Area),
	}
}

// Helper function for min
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
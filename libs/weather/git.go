package weather

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// GitMonitor handles git activity monitoring and smart inference
type GitMonitor struct {
	gardenPath string
	cm         *ContextManager
}

// NewGitMonitor creates a new git monitor for the given garden path
func NewGitMonitor(gardenPath string) *GitMonitor {
	return &GitMonitor{
		gardenPath: gardenPath,
		cm:         NewContextManager(gardenPath),
	}
}

// InstallGitHooks installs git hooks for automatic weather updates
func (gm *GitMonitor) InstallGitHooks() error {
	gitHooksDir := filepath.Join(gm.gardenPath, ".git", "hooks")
	
	// Check if .git directory exists
	if _, err := os.Stat(filepath.Join(gm.gardenPath, ".git")); os.IsNotExist(err) {
		return fmt.Errorf("not a git repository")
	}
	
	// Create hooks directory if it doesn't exist
	if err := os.MkdirAll(gitHooksDir, 0755); err != nil {
		return fmt.Errorf("failed to create hooks directory: %w", err)
	}
	
	// Install post-commit hook
	postCommitHook := filepath.Join(gitHooksDir, "post-commit")
	postCommitScript := `#!/bin/sh
# Weather context update after commit
sprout weather --update-from-commit HEAD 2>/dev/null || true
`
	if err := os.WriteFile(postCommitHook, []byte(postCommitScript), 0755); err != nil {
		return fmt.Errorf("failed to create post-commit hook: %w", err)
	}
	
	// Install post-checkout hook for branch changes
	postCheckoutHook := filepath.Join(gitHooksDir, "post-checkout")
	postCheckoutScript := `#!/bin/sh
# Weather context update after branch change
sprout weather --update-from-branch-change "$1" "$2" "$3" 2>/dev/null || true
`
	if err := os.WriteFile(postCheckoutHook, []byte(postCheckoutScript), 0755); err != nil {
		return fmt.Errorf("failed to create post-checkout hook: %w", err)
	}
	
	return nil
}

// UpdateFromCommit updates weather context based on a specific commit
func (gm *GitMonitor) UpdateFromCommit(commitHash string) error {
	commit, err := gm.getCommitInfo(commitHash)
	if err != nil {
		return fmt.Errorf("failed to get commit info: %w", err)
	}
	
	return gm.cm.UpdateContext(func(ctx *WeatherContext) {
		// Update git context
		ctx.Git.LastCommit = *commit
		ctx.Git.CurrentBranch = gm.getCurrentBranch()
		ctx.Git.UncommittedChanges = gm.hasUncommittedChanges()
		
		// Update recent progress
		gm.updateRecentProgress(ctx, *commit)
		
		// Update current focus based on commit
		gm.updateCurrentFocus(ctx, *commit)
		
		// Update weather conditions
		gm.updateWeatherConditions(ctx)
		
		// Update next steps
		gm.updateNextSteps(ctx)
	})
}

// UpdateFromBranchChange updates weather context when branches change
func (gm *GitMonitor) UpdateFromBranchChange(prevHead, newHead, branchFlag string) error {
	return gm.cm.UpdateContext(func(ctx *WeatherContext) {
		currentBranch := gm.getCurrentBranch()
		ctx.Git.CurrentBranch = currentBranch
		
		// Update focus area when switching branches
		if strings.Contains(currentBranch, "/") {
			parts := strings.Split(currentBranch, "/")
			if len(parts) >= 2 {
				branchType := parts[0]
				branchName := strings.Join(parts[1:], "/")
				
				switch branchType {
				case "feature", "feat":
					ctx.CurrentFocus.Area = fmt.Sprintf("feature: %s", branchName)
					ctx.CurrentFocus.InferredFrom = fmt.Sprintf("switched to feature branch %s", currentBranch)
				case "fix", "bugfix", "hotfix":
					ctx.CurrentFocus.Area = fmt.Sprintf("bug fix: %s", branchName)
					ctx.CurrentFocus.InferredFrom = fmt.Sprintf("switched to fix branch %s", currentBranch)  
				case "refactor":
					ctx.CurrentFocus.Area = fmt.Sprintf("refactoring: %s", branchName)
					ctx.CurrentFocus.InferredFrom = fmt.Sprintf("switched to refactor branch %s", currentBranch)
				default:
					ctx.CurrentFocus.Area = fmt.Sprintf("working on: %s", branchName)
					ctx.CurrentFocus.InferredFrom = fmt.Sprintf("switched to branch %s", currentBranch)
				}
				ctx.CurrentFocus.Confidence = 0.8
			}
		}
		
		ctx.CurrentFocus.LastActive = time.Now()
		
		// Update weather - branch switching can affect temperature
		ctx.Weather.Temperature = min(ctx.Weather.Temperature+5, 100) // Slight activity increase
		ctx.Weather.LastUpdate = time.Now()
	})
}

// GetRecentCommits gets recent commits for context analysis
func (gm *GitMonitor) GetRecentCommits(limit int) ([]GitCommit, error) {
	cmd := exec.Command("git", "log", fmt.Sprintf("-%d", limit), "--pretty=format:%H|%s|%an|%at", "--name-only")
	cmd.Dir = gm.gardenPath
	
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get git log: %w", err)
	}
	
	return gm.parseGitLog(string(output))
}

// getCommitInfo gets detailed information about a specific commit
func (gm *GitMonitor) getCommitInfo(commitHash string) (*GitCommit, error) {
	// Get commit details
	cmd := exec.Command("git", "show", "--pretty=format:%H|%s|%an|%at", "--name-only", commitHash)
	cmd.Dir = gm.gardenPath
	
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get commit details: %w", err)
	}
	
	commits, err := gm.parseGitLog(string(output))
	if err != nil {
		return nil, err
	}
	
	if len(commits) == 0 {
		return nil, fmt.Errorf("no commit found")
	}
	
	commit := &commits[0]
	commit.SmartSummary = gm.generateSmartSummary(commit.Message)
	commit.InferredScope = gm.inferScope(commit.FilesChanged)
	
	return commit, nil
}

// parseGitLog parses git log output into GitCommit structs
func (gm *GitMonitor) parseGitLog(output string) ([]GitCommit, error) {
	var commits []GitCommit
	lines := strings.Split(strings.TrimSpace(output), "\n")
	
	i := 0
	for i < len(lines) {
		if lines[i] == "" {
			i++
			continue
		}
		
		// Parse commit header: hash|subject|author|timestamp
		parts := strings.Split(lines[i], "|")
		if len(parts) != 4 {
			i++
			continue
		}
		
		timestamp, err := strconv.ParseInt(parts[3], 10, 64)
		if err != nil {
			i++
			continue
		}
		
		commit := GitCommit{
			Hash:      parts[0],
			Message:   parts[1],
			Author:    parts[2],
			Timestamp: time.Unix(timestamp, 0),
		}
		
		// Parse file list (following lines until next commit or end)
		i++
		var files []string
		for i < len(lines) && !strings.Contains(lines[i], "|") && lines[i] != "" {
			files = append(files, lines[i])
			i++
		}
		commit.FilesChanged = files
		
		commits = append(commits, commit)
	}
	
	return commits, nil
}

// generateSmartSummary creates human-readable summary from commit message
func (gm *GitMonitor) generateSmartSummary(message string) string {
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

// inferScope determines the scope/area of work from changed files
func (gm *GitMonitor) inferScope(files []string) string {
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

// updateRecentProgress updates the recent progress summary
func (gm *GitMonitor) updateRecentProgress(ctx *WeatherContext, commit GitCommit) {
	// Get recent commits for better context
	recentCommits, err := gm.GetRecentCommits(5)
	if err != nil {
		recentCommits = []GitCommit{commit}
	}
	
	// Update commits list
	ctx.RecentProgress.Commits = recentCommits
	
	// Generate summary based on recent activity
	if len(recentCommits) == 1 {
		ctx.RecentProgress.Summary = commit.SmartSummary
		ctx.RecentProgress.Timespan = "last commit"
	} else {
		// Analyze patterns in recent commits
		scopes := make(map[string]int)
		for _, c := range recentCommits {
			scope := gm.inferScope(c.FilesChanged)
			scopes[scope]++
		}
		
		// Find primary area of work
		var primaryScope string
		maxCount := 0
		for scope, count := range scopes {
			if count > maxCount {
				maxCount = count
				primaryScope = scope
			}
		}
		
		ctx.RecentProgress.Summary = fmt.Sprintf("Working on %s system", primaryScope)
		ctx.RecentProgress.Timespan = fmt.Sprintf("last %d commits", len(recentCommits))
	}
	
	// Update momentum based on commit frequency
	now := time.Now()
	recentCommitCount := 0
	for _, c := range recentCommits {
		if now.Sub(c.Timestamp) < 2*time.Hour {
			recentCommitCount++
		}
	}
	
	// Calculate momentum: more recent commits = higher momentum
	ctx.RecentProgress.Momentum = min(recentCommitCount*25, 100)
}

// updateCurrentFocus updates the current focus area
func (gm *GitMonitor) updateCurrentFocus(ctx *WeatherContext, commit GitCommit) {
	scope := commit.InferredScope
	
	// Generate focus area description
	var focusArea string
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
	
	ctx.CurrentFocus.Area = focusArea
	ctx.CurrentFocus.Confidence = 0.85 // High confidence from commit activity
	ctx.CurrentFocus.LastActive = time.Now()
	ctx.CurrentFocus.InferredFrom = fmt.Sprintf("recent commit to %s", scope)
}

// updateWeatherConditions updates weather based on git activity
func (gm *GitMonitor) updateWeatherConditions(ctx *WeatherContext) {
	// Temperature based on momentum
	ctx.Weather.Temperature = min(20+ctx.RecentProgress.Momentum, 95)
	
	// Condition based on recent activity patterns
	recentCommits := ctx.RecentProgress.Commits
	if len(recentCommits) == 0 {
		ctx.Weather.Condition = WeatherSunny
	} else {
		// Analyze commit messages for indicators
		fixCount := 0
		featCount := 0
		for _, commit := range recentCommits {
			msg := strings.ToLower(commit.Message)
			if strings.Contains(msg, "fix") || strings.Contains(msg, "bug") {
				fixCount++
			}
			if strings.Contains(msg, "feat") || strings.Contains(msg, "add") {
				featCount++
			}
		}
		
		// Determine weather condition
		totalCommits := len(recentCommits)
		if fixCount > totalCommits/2 {
			ctx.Weather.Condition = WeatherCloudy // Lots of fixes = some challenges
		} else if featCount > totalCommits/2 {
			ctx.Weather.Condition = WeatherSunny // Lots of features = smooth progress
		} else {
			ctx.Weather.Condition = WeatherPartlyCloudy // Mixed activity
		}
	}
	
	// Pressure remains relatively stable (would be updated by external factors)
	ctx.Weather.LastUpdate = time.Now()
}

// updateNextSteps generates next step suggestions
func (gm *GitMonitor) updateNextSteps(ctx *WeatherContext) {
	var suggestions []string
	
	// Base suggestions on current focus and recent activity
	scope := strings.ToLower(ctx.CurrentFocus.Area)
	
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
	
	ctx.NextSteps.Suggestions = suggestions
	ctx.NextSteps.Priority = min(50+ctx.RecentProgress.Momentum/2, 90)
	ctx.NextSteps.BasedOn = fmt.Sprintf("current focus on %s", ctx.CurrentFocus.Area)
}

// getCurrentBranch gets the current git branch
func (gm *GitMonitor) getCurrentBranch() string {
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = gm.gardenPath
	
	output, err := cmd.Output()
	if err != nil {
		return "main" // fallback
	}
	
	return strings.TrimSpace(string(output))
}

// hasUncommittedChanges checks if there are uncommitted changes
func (gm *GitMonitor) hasUncommittedChanges() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = gm.gardenPath
	
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	
	return len(strings.TrimSpace(string(output))) > 0
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
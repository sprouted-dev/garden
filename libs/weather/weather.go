package weather

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Weather represents the core weather system with context and intelligence
type Weather struct {
	RepoPath      string
	Context       *WeatherContext
	Config        *Config
	CurrentBranch string
}

// NewWeather creates a new Weather instance
func NewWeather(repoPath string) (*Weather, error) {
	cm := NewContextManager(repoPath)
	context, err := cm.LoadContext()
	if err != nil {
		return nil, fmt.Errorf("failed to load context: %w", err)
	}
	
	configMgr := NewConfigManager(repoPath)
	config, err := configMgr.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	
	return &Weather{
		RepoPath:      repoPath,
		Context:       context,
		Config:        config,
		CurrentBranch: context.Git.CurrentBranch,
	}, nil
}

// detectMissingDecisions looks for decisions in commits that aren't documented
func (w *Weather) detectMissingDecisions(existingDocs map[string]bool) []DocumentationNeed {
	var needs []DocumentationNeed
	
	// Patterns that indicate decisions
	decisionPatterns := []*regexp.Regexp{
		regexp.MustCompile(`(?i)(decided|chose|selected|opted|picked)\s+to\s+`),
		regexp.MustCompile(`(?i)(instead of|rather than|over)\s+`),
		regexp.MustCompile(`(?i)(fix|solve|resolve|handle)\s+by\s+`),
		regexp.MustCompile(`(?i)(implement|use|adopt)\s+\w+\s+(for|to|as)`),
	}
	
	// Check recent commits
	for _, commit := range w.Context.RecentProgress.Commits {
		for _, pattern := range decisionPatterns {
			if pattern.MatchString(commit.Message) {
				// Extract potential decision title
				title := extractDecisionTitle(commit.Message)
				
				// Check if already documented
				if !existingDocs[strings.ToLower(title)] {
					needs = append(needs, DocumentationNeed{
						Type:              "decision",
						Title:             title,
						Description:       fmt.Sprintf("Decision made in commit %s: %s", commit.Hash[:8], commit.Message),
						SuggestedLocation: "docs/decisions/",
						Confidence:        0.75,
						DetectedFrom:      []string{"commit message patterns"},
					})
				}
				break
			}
		}
	}
	
	return needs
}

// detectLessonsLearned identifies crisis resolutions that should be documented
func (w *Weather) detectLessonsLearned(existingDocs map[string]bool) []DocumentationNeed {
	var needs []DocumentationNeed
	
	// Crisis indicators
	crisisPatterns := []*regexp.Regexp{
		regexp.MustCompile(`(?i)(crisis|critical|emergency|urgent|broken|failed)`),
		regexp.MustCompile(`(?i)(fix|fixed|resolve|resolved|repair|repaired)\s+(critical|urgent|breaking|broken)`),
		regexp.MustCompile(`(?i)(hotfix|emergency\s+fix|quick\s+fix)`),
	}
	
	resolutionPatterns := []*regexp.Regexp{
		regexp.MustCompile(`(?i)(working|fixed|resolved|successful|passes)`),
		regexp.MustCompile(`(?i)(now\s+works|back\s+online|restored)`),
	}
	
	// Look for crisis → resolution patterns in recent commits
	var crisisFound bool
	var crisisDescription string
	
	for i, commit := range w.Context.RecentProgress.Commits {
		// Check for crisis
		if !crisisFound {
			for _, pattern := range crisisPatterns {
				if pattern.MatchString(commit.Message) {
					crisisFound = true
					crisisDescription = commit.Message
					break
				}
			}
		} else {
			// Look for resolution
			for _, pattern := range resolutionPatterns {
				if pattern.MatchString(commit.Message) {
					// Found crisis → resolution pattern
					title := fmt.Sprintf("Crisis Resolution: %s", extractCrisisTitle(crisisDescription))
					
					if !existingDocs[strings.ToLower(title)] {
						needs = append(needs, DocumentationNeed{
							Type:              "lesson",
							Title:             title,
							Description:       fmt.Sprintf("Crisis detected and resolved between commits %d and %d", i-1, i),
							SuggestedLocation: "docs/lessons/",
							Confidence:        0.85,
							DetectedFrom:      []string{"crisis resolution patterns", "commit sequence analysis"},
						})
					}
					
					crisisFound = false
					break
				}
			}
		}
	}
	
	return needs
}

// detectProcessGaps looks for repeated patterns that indicate missing process documentation
func (w *Weather) detectProcessGaps(existingDocs map[string]bool) []DocumentationNeed {
	var needs []DocumentationNeed
	
	// Count activity patterns
	activityCounts := make(map[string]int)
	
	for _, commit := range w.Context.RecentProgress.Commits {
		// Extract activity type
		activity := extractActivityType(commit.Message)
		if activity != "" {
			activityCounts[activity]++
		}
	}
	
	// Look for repeated activities that might need process docs
	for activity, count := range activityCounts {
		if count >= 3 { // Repeated 3+ times suggests a process
			title := fmt.Sprintf("%s Process", strings.Title(activity))
			
			if !existingDocs[strings.ToLower(title)] {
				needs = append(needs, DocumentationNeed{
					Type:              "process",
					Title:             title,
					Description:       fmt.Sprintf("Activity '%s' repeated %d times, suggesting need for process documentation", activity, count),
					SuggestedLocation: "docs/processes/",
					Confidence:        float64(count) / 10.0, // More repetitions = higher confidence
					DetectedFrom:      []string{"activity pattern analysis", "repetition detection"},
				})
			}
		}
	}
	
	return needs
}

// suggestRecentActivityDocumentation analyzes recent weather for documentation opportunities
func (w *Weather) suggestRecentActivityDocumentation() []ActivitySuggestion {
	var suggestions []ActivitySuggestion
	
	// Check temperature spikes (high activity)
	if w.Context.Weather.Temperature > 90 {
		suggestions = append(suggestions, ActivitySuggestion{
			Activity:         "High development activity period",
			Timeframe:        "current session",
			ShouldDocument:   true,
			SuggestedDocType: "process",
			Reasoning:        "High temperature indicates significant activity that may contain valuable patterns",
		})
	}
	
	// Check for recent focus changes
	if w.Context.CurrentFocus.Area != "" && w.Context.CurrentFocus.Confidence > 0.8 {
		suggestions = append(suggestions, ActivitySuggestion{
			Activity:         fmt.Sprintf("Focus on %s", w.Context.CurrentFocus.Area),
			Timeframe:        "current",
			ShouldDocument:   false, // Only if it leads to decisions
			SuggestedDocType: "decision",
			Reasoning:        "Strong focus area might involve architectural decisions",
		})
	}
	
	// Check for branch changes (feature work)
	if w.Context.Git.CurrentBranch != "main" && w.Context.Git.CurrentBranch != "master" {
		suggestions = append(suggestions, ActivitySuggestion{
			Activity:         fmt.Sprintf("Feature branch work: %s", w.Context.Git.CurrentBranch),
			Timeframe:        "current branch",
			ShouldDocument:   true,
			SuggestedDocType: "spec",
			Reasoning:        "Feature branches often contain undocumented requirements",
		})
	}
	
	return suggestions
}

// loadConversationalCaptures loads saved conversational knowledge
func (w *Weather) loadConversationalCaptures() []ConversationCapture {
	var captures []ConversationCapture
	
	// For now, return the example conversation we captured
	// In a full implementation, this would load from .garden/conversations/
	
	captures = append(captures, ConversationCapture{
		Timestamp:        time.Now(),
		UserPrompt:       "make sure you remember the current issue with the Farm level invisibility",
		AIResponse:       "Git hooks in individual gardens can't see or coordinate across the farm boundary",
		ConversationType: "architectural_discovery",
		KeyInsights: []string{
			"Farm root invisibility requires orchestration layer",
			"Multiple AI assistants rediscovering same constraints",
			"Weather System experiencing the problem it solves",
		},
		Outcome:       "Identified need for orchestration layer",
		ShouldPersist: true,
		RelatedFiles:  []string{"docs/specs/farm-orchestration-layer.md"},
		Context:       "Weather System gap analysis",
	})
	
	return captures
}

// Helper functions

func extractDecisionTitle(message string) string {
	// Extract key decision from commit message
	message = strings.ToLower(message)
	
	if idx := strings.Index(message, "decided to "); idx >= 0 {
		rest := message[idx+11:]
		if endIdx := strings.IndexAny(rest, ".,;"); endIdx > 0 {
			return strings.Title(rest[:endIdx]) + " decision"
		}
	}
	
	if idx := strings.Index(message, "chose "); idx >= 0 {
		rest := message[idx+6:]
		if endIdx := strings.IndexAny(rest, ".,;"); endIdx > 0 {
			return strings.Title(rest[:endIdx]) + " choice"
		}
	}
	
	// Fallback: use first part of message
	parts := strings.Fields(message)
	if len(parts) > 3 {
		return strings.Title(strings.Join(parts[:3], " ")) + " decision"
	}
	
	return "Undocumented decision"
}

func extractCrisisTitle(message string) string {
	// Extract crisis type from message
	message = strings.ToLower(message)
	
	crisisTypes := []string{"installation", "deployment", "build", "test", "security", "performance"}
	
	for _, crisisType := range crisisTypes {
		if strings.Contains(message, crisisType) {
			return strings.Title(crisisType) + " issue"
		}
	}
	
	return "System issue"
}

func extractActivityType(message string) string {
	// Extract activity type from commit message
	message = strings.ToLower(message)
	
	// Common activity prefixes
	activities := map[string]string{
		"feat:":     "feature",
		"fix:":      "bugfix",
		"docs:":     "documentation",
		"test:":     "testing",
		"refactor:": "refactoring",
		"chore:":    "maintenance",
		"deploy:":   "deployment",
	}
	
	for prefix, activity := range activities {
		if strings.HasPrefix(message, prefix) {
			return activity
		}
	}
	
	// Check for common verbs
	verbs := map[string]string{
		"add":      "addition",
		"remove":   "removal",
		"update":   "update",
		"fix":      "bugfix",
		"improve":  "improvement",
		"optimize": "optimization",
	}
	
	words := strings.Fields(message)
	if len(words) > 0 {
		if activity, ok := verbs[words[0]]; ok {
			return activity
		}
	}
	
	return ""
}
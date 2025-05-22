// Package weather provides automatic context preservation and intelligence
// for development workflows, serving as the heartbeat of the Sprouted ecosystem.
package weather

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// WeatherContext represents the complete development context state
// automatically maintained through git activity monitoring and smart inference.
type WeatherContext struct {
	// Metadata
	Updated    time.Time `json:"updated"`
	SessionID  string    `json:"sessionId"`
	GardenPath string    `json:"gardenPath"`
	Version    string    `json:"version"`

	// Core Intelligence
	CurrentFocus   FocusArea           `json:"currentFocus"`
	RecentProgress ProgressSummary     `json:"recentProgress"`
	NextSteps      NextStepsSuggestion `json:"nextSteps"`
	Weather        WeatherConditions   `json:"weather"`

	// Git Integration
	Git GitContext `json:"git"`

	// AI Assistant Onboarding Context
	Onboarding OnboardingContext `json:"onboarding"`
}

// FocusArea represents what the developer is currently working on
type FocusArea struct {
	Area        string    `json:"area"`        // e.g., "authentication system"
	Confidence  float64   `json:"confidence"`  // 0-1 confidence in inference
	LastActive  time.Time `json:"lastActive"`
	InferredFrom string   `json:"inferredFrom"` // e.g., "recent commits to auth/ directory"
}

// ProgressSummary captures recent development progress
type ProgressSummary struct {
	Summary   string      `json:"summary"`   // e.g., "Implemented user login and JWT validation"
	Commits   []GitCommit `json:"commits"`   // Recent commits with smart summaries
	Timespan  string      `json:"timespan"`  // e.g., "last 2 hours"
	Momentum  int         `json:"momentum"`  // 0-100 activity level
}

// NextStepsSuggestion provides AI-generated next actions
type NextStepsSuggestion struct {
	Suggestions []string `json:"suggestions"` // AI-generated next actions
	Priority    int      `json:"priority"`    // 0-100 urgency
	BasedOn     string   `json:"basedOn"`     // What the suggestions are inferred from
}

// WeatherConditions represents the current development "weather"
type WeatherConditions struct {
	Temperature int             `json:"temperature"` // 0-100 activity level
	Condition   WeatherCondition `json:"condition"`
	Pressure    int             `json:"pressure"`    // Deadline/urgency pressure
	LastUpdate  time.Time       `json:"lastUpdate"`
}

// GitContext captures git repository state
type GitContext struct {
	CurrentBranch       string      `json:"currentBranch"`
	LastCommit          GitCommit   `json:"lastCommit"`
	UncommittedChanges  bool        `json:"uncommittedChanges"`
	RecentBranches      []string    `json:"recentBranches"`
}

// GitCommit represents a git commit with smart analysis
type GitCommit struct {
	Hash          string    `json:"hash"`
	Message       string    `json:"message"`
	SmartSummary  string    `json:"smartSummary"`  // AI-generated human-readable summary
	Timestamp     time.Time `json:"timestamp"`
	FilesChanged  []string  `json:"filesChanged"`
	InferredScope string    `json:"inferredScope"` // e.g., "frontend/auth", "backend/api" 
	Author        string    `json:"author"`
}

// WeatherCondition represents weather state
type WeatherCondition string

const (
	WeatherSunny        WeatherCondition = "sunny"         // Smooth progress, no blockers
	WeatherPartlyCloudy WeatherCondition = "partly-cloudy" // Some minor issues or complexity
	WeatherCloudy       WeatherCondition = "cloudy"        // Multiple challenges or slow progress
	WeatherStormy       WeatherCondition = "stormy"        // Major blockers or critical issues
	WeatherFoggy        WeatherCondition = "foggy"         // Unclear direction or exploration phase
)

// OnboardingContext provides comprehensive context for AI assistant onboarding
type OnboardingContext struct {
	// Development Methodology
	Methodology  DevelopmentMethodology `json:"methodology"`
	
	// Project Vision and Mission
	Vision       ProjectVision          `json:"vision"`
	
	// Active Work Context
	ActiveWork   ActiveWorkContext      `json:"activeWork"`
	
	// Architectural Context
	Architecture ArchitecturalContext   `json:"architecture"`
	
	// Quick Start Guidance
	QuickStart   QuickStartGuide        `json:"quickStart"`
	
	// Last Updated
	LastScanned  time.Time              `json:"lastScanned"`
}

// DevelopmentMethodology describes how work is done in this project
type DevelopmentMethodology struct {
	ProcessName     string   `json:"processName"`     // e.g., "spec-driven development"
	Hierarchy       []string `json:"hierarchy"`       // e.g., ["Vision", "Specs", "Tasks", "Implementation"]
	RequiredDocs    []string `json:"requiredDocs"`    // Documentation requirements
	Templates       []string `json:"templates"`       // Available templates
	Principles      []string `json:"principles"`      // Key development principles
	WorkflowGuide   string   `json:"workflowGuide"`   // Path to workflow documentation
}

// ProjectVision captures the high-level project context
type ProjectVision struct {
	Name        string `json:"name"`        // Project name
	Mission     string `json:"mission"`     // Core mission statement  
	Problem     string `json:"problem"`     // Problem being solved
	Solution    string `json:"solution"`    // High-level solution approach
	CoreSystem  string `json:"coreSystem"`  // Central system/component
	Ecosystem   string `json:"ecosystem"`   // Broader ecosystem context
}

// ActiveWorkContext describes current active work
type ActiveWorkContext struct {
	CurrentPhase    string              `json:"currentPhase"`    // e.g., "Weather MVP Phase 1"
	Timeline        string              `json:"timeline"`        // e.g., "Week 2 of 4-week timeline"
	ActiveSpecs     []DocumentSummary   `json:"activeSpecs"`     // Current specs being worked on
	ActiveTasks     []DocumentSummary   `json:"activeTasks"`     // Current tasks in progress
	Priority        string              `json:"priority"`        // Current priority focus
	ImplementationContext string        `json:"implementationContext"` // Where in the implementation process
}

// ArchitecturalContext captures key architectural decisions and patterns
type ArchitecturalContext struct {
	Language        string            `json:"language"`        // Primary language (Go, TypeScript, etc.)
	Architecture    string            `json:"architecture"`    // e.g., "monorepo", "microservices"
	KeyPatterns     []string          `json:"keyPatterns"`     // Important patterns being used
	Conventions     []string          `json:"conventions"`     // Coding conventions and standards
	KeyDecisions    []ArchDecision    `json:"keyDecisions"`    // Important architectural decisions
	DirectoryStructure map[string]string `json:"directoryStructure"` // Key directories and their purposes
}

// ArchDecision represents an architectural decision
type ArchDecision struct {
	Decision    string    `json:"decision"`    // What was decided
	Rationale   string    `json:"rationale"`   // Why it was decided
	Impact      string    `json:"impact"`      // Impact on development
	Timestamp   time.Time `json:"timestamp"`   // When decision was made
}

// QuickStartGuide provides essential getting-started information
type QuickStartGuide struct {
	EssentialDocs   []string `json:"essentialDocs"`   // Must-read documents
	KeyCommands     []string `json:"keyCommands"`     // Important CLI commands
	FirstSteps      []string `json:"firstSteps"`      // Recommended first actions
	CommonPatterns  []string `json:"commonPatterns"`  // Frequently used patterns
	AvoidPatterns   []string `json:"avoidPatterns"`   // Patterns to avoid
}

// DocumentSummary provides a summary of an important document
type DocumentSummary struct {
	Path        string    `json:"path"`        // File path
	Title       string    `json:"title"`       // Document title
	Type        string    `json:"type"`        // vision, spec, task, etc.
	Summary     string    `json:"summary"`     // Brief content summary
	Status      string    `json:"status"`      // current status if applicable
	LastUpdated time.Time `json:"lastUpdated"` // Last modification time
}

// ContextManager handles weather context persistence and operations
type ContextManager struct {
	gardenPath string
}

// NewContextManager creates a new context manager for the given garden path
func NewContextManager(gardenPath string) *ContextManager {
	return &ContextManager{
		gardenPath: gardenPath,
	}
}

// GetContextPath returns the path to the weather context file
func (cm *ContextManager) GetContextPath() string {
	return filepath.Join(cm.gardenPath, ".garden", "weather-context.json")
}

// LoadContext loads the weather context from disk, creating a default if none exists
func (cm *ContextManager) LoadContext() (*WeatherContext, error) {
	contextPath := cm.GetContextPath()
	
	// Check if context file exists
	if _, err := os.Stat(contextPath); os.IsNotExist(err) {
		// Create default context
		context := cm.createDefaultContext()
		if err := cm.SaveContext(&context); err != nil {
			return nil, fmt.Errorf("failed to create default context: %w", err)
		}
		return &context, nil
	}
	
	// Read existing context
	data, err := os.ReadFile(contextPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read context file: %w", err)
	}
	
	var context WeatherContext
	if err := json.Unmarshal(data, &context); err != nil {
		// Context file corrupted, create new one
		fmt.Printf("Warning: corrupted weather context file, creating new one\n")
		newContext := cm.createDefaultContext()
		if err := cm.SaveContext(&newContext); err != nil {
			return nil, fmt.Errorf("failed to recreate context after corruption: %w", err)
		}
		return &newContext, nil
	}
	
	// Refresh onboarding context if it's empty or stale (older than 1 hour)
	if context.Onboarding.LastScanned.IsZero() || time.Since(context.Onboarding.LastScanned) > time.Hour {
		if err := cm.updateOnboardingContext(&context); err != nil {
			// Log error but continue with existing context
			fmt.Printf("Warning: failed to update onboarding context: %v\n", err)
		} else {
			// Save updated context with new onboarding data
			if err := cm.SaveContext(&context); err != nil {
				fmt.Printf("Warning: failed to save updated onboarding context: %v\n", err)
			}
		}
	}
	
	return &context, nil
}

// SaveContext saves the weather context to disk using atomic writes
func (cm *ContextManager) SaveContext(context *WeatherContext) error {
	contextPath := cm.GetContextPath()
	
	// Ensure .garden directory exists
	gardenDir := filepath.Dir(contextPath)
	if err := os.MkdirAll(gardenDir, 0755); err != nil {
		return fmt.Errorf("failed to create .garden directory: %w", err)
	}
	
	// Update metadata
	context.Updated = time.Now()
	context.GardenPath = cm.gardenPath
	context.Version = "1.0.0"
	
	// Marshal to JSON with indentation for readability
	data, err := json.MarshalIndent(context, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal context: %w", err)
	}
	
	// Atomic write using temporary file
	tempPath := contextPath + ".tmp"
	if err := os.WriteFile(tempPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write temporary context file: %w", err)
	}
	
	// Atomic rename
	if err := os.Rename(tempPath, contextPath); err != nil {
		os.Remove(tempPath) // Clean up temp file
		return fmt.Errorf("failed to rename temporary context file: %w", err)
	}
	
	return nil
}

// UpdateContext updates the weather context and saves it
func (cm *ContextManager) UpdateContext(updater func(*WeatherContext)) error {
	context, err := cm.LoadContext()
	if err != nil {
		return fmt.Errorf("failed to load context for update: %w", err)
	}
	
	updater(context)
	
	if err := cm.SaveContext(context); err != nil {
		return fmt.Errorf("failed to save updated context: %w", err)
	}
	
	return nil
}

// createDefaultContext creates a new default weather context
func (cm *ContextManager) createDefaultContext() WeatherContext {
	now := time.Now()
	
	context := WeatherContext{
		Updated:    now,
		SessionID:  generateSessionID(),
		GardenPath: cm.gardenPath,
		Version:    "1.0.0",
		
		CurrentFocus: FocusArea{
			Area:        "initializing garden",
			Confidence:  1.0,
			LastActive:  now,
			InferredFrom: "new garden created",
		},
		
		RecentProgress: ProgressSummary{
			Summary:  "Garden initialized",
			Commits:  []GitCommit{},
			Timespan: "just now",
			Momentum: 0,
		},
		
		NextSteps: NextStepsSuggestion{
			Suggestions: []string{
				"Make your first commit to start tracking progress",
				"Begin working on your first feature",
			},
			Priority: 50,
			BasedOn:  "new garden initialization",
		},
		
		Weather: WeatherConditions{
			Temperature: 20, // Cool - just starting
			Condition:   WeatherSunny,
			Pressure:    30, // Low pressure - no deadlines yet
			LastUpdate:  now,
		},
		
		Git: GitContext{
			CurrentBranch:      "main",
			LastCommit:         GitCommit{},
			UncommittedChanges: false,
			RecentBranches:     []string{"main"},
		},
	}
	
	// Populate onboarding context
	cm.updateOnboardingContext(&context)
	
	return context
}

// generateSessionID creates a unique session identifier
func generateSessionID() string {
	return fmt.Sprintf("session_%d", time.Now().Unix())
}

// Validate checks if the weather context is valid and consistent
func (wc *WeatherContext) Validate() error {
	if wc.Version == "" {
		return fmt.Errorf("weather context missing version")
	}
	
	if wc.GardenPath == "" {
		return fmt.Errorf("weather context missing garden path")
	}
	
	if wc.CurrentFocus.Confidence < 0 || wc.CurrentFocus.Confidence > 1 {
		return fmt.Errorf("focus confidence must be between 0 and 1")
	}
	
	if wc.Weather.Temperature < 0 || wc.Weather.Temperature > 100 {
		return fmt.Errorf("temperature must be between 0 and 100")
	}
	
	if wc.Weather.Pressure < 0 || wc.Weather.Pressure > 100 {
		return fmt.Errorf("pressure must be between 0 and 100")
	}
	
	return nil
}

// ToJSON converts the weather context to JSON for AI consumption
func (wc *WeatherContext) ToJSON() ([]byte, error) {
	return json.MarshalIndent(wc, "", "  ")
}

// ToAIContext converts the weather context to a format optimized for AI assistants
func (wc *WeatherContext) ToAIContext() map[string]interface{} {
	return map[string]interface{}{
		"project_status": map[string]interface{}{
			"current_focus":    wc.CurrentFocus.Area,
			"confidence":       wc.CurrentFocus.Confidence,
			"recent_progress":  wc.RecentProgress.Summary,
			"momentum":         wc.RecentProgress.Momentum,
			"next_steps":       wc.NextSteps.Suggestions,
		},
		"development_context": map[string]interface{}{
			"current_branch":       wc.Git.CurrentBranch,
			"uncommitted_changes":  wc.Git.UncommittedChanges,
			"last_commit_message":  wc.Git.LastCommit.Message,
			"last_commit_summary":  wc.Git.LastCommit.SmartSummary,
			"files_changed":        wc.Git.LastCommit.FilesChanged,
		},
		"weather_conditions": map[string]interface{}{
			"temperature":  wc.Weather.Temperature,
			"condition":    string(wc.Weather.Condition),
			"pressure":     wc.Weather.Pressure,
			"last_update":  wc.Weather.LastUpdate.Format(time.RFC3339),
		},
		"metadata": map[string]interface{}{
			"garden_path":  wc.GardenPath,
			"session_id":   wc.SessionID,
			"last_updated": wc.Updated.Format(time.RFC3339),
		},
	}
}

// ToAIOnboardingContext converts the weather context to comprehensive AI assistant onboarding format
func (wc *WeatherContext) ToAIOnboardingContext() map[string]interface{} {
	// Start with basic AI context
	context := wc.ToAIContext()
	
	// Add comprehensive onboarding information
	context["development_methodology"] = map[string]interface{}{
		"process_name":     wc.Onboarding.Methodology.ProcessName,
		"hierarchy":        wc.Onboarding.Methodology.Hierarchy,
		"required_docs":    wc.Onboarding.Methodology.RequiredDocs,
		"templates":        wc.Onboarding.Methodology.Templates,
		"principles":       wc.Onboarding.Methodology.Principles,
		"workflow_guide":   wc.Onboarding.Methodology.WorkflowGuide,
	}
	
	context["project_vision"] = map[string]interface{}{
		"name":         wc.Onboarding.Vision.Name,
		"mission":      wc.Onboarding.Vision.Mission,
		"problem":      wc.Onboarding.Vision.Problem,
		"solution":     wc.Onboarding.Vision.Solution,
		"core_system":  wc.Onboarding.Vision.CoreSystem,
		"ecosystem":    wc.Onboarding.Vision.Ecosystem,
	}
	
	context["active_work"] = map[string]interface{}{
		"current_phase":           wc.Onboarding.ActiveWork.CurrentPhase,
		"timeline":               wc.Onboarding.ActiveWork.Timeline,
		"active_specs":           wc.Onboarding.ActiveWork.ActiveSpecs,
		"active_tasks":           wc.Onboarding.ActiveWork.ActiveTasks,
		"priority":               wc.Onboarding.ActiveWork.Priority,
		"implementation_context": wc.Onboarding.ActiveWork.ImplementationContext,
	}
	
	context["architectural_context"] = map[string]interface{}{
		"language":             wc.Onboarding.Architecture.Language,
		"architecture":         wc.Onboarding.Architecture.Architecture,
		"key_patterns":         wc.Onboarding.Architecture.KeyPatterns,
		"conventions":          wc.Onboarding.Architecture.Conventions,
		"key_decisions":        wc.Onboarding.Architecture.KeyDecisions,
		"directory_structure":  wc.Onboarding.Architecture.DirectoryStructure,
	}
	
	context["onboarding_essentials"] = map[string]interface{}{
		"essential_docs":   wc.Onboarding.QuickStart.EssentialDocs,
		"key_commands":     wc.Onboarding.QuickStart.KeyCommands,
		"first_steps":      wc.Onboarding.QuickStart.FirstSteps,
		"common_patterns":  wc.Onboarding.QuickStart.CommonPatterns,
		"avoid_patterns":   wc.Onboarding.QuickStart.AvoidPatterns,
		"last_scanned":     wc.Onboarding.LastScanned.Format(time.RFC3339),
	}
	
	return context
}

// updateOnboardingContext refreshes the onboarding context by scanning documentation
func (cm *ContextManager) updateOnboardingContext(context *WeatherContext) error {
	docIntel := NewDocumentIntelligence(cm.gardenPath)
	
	onboarding, err := docIntel.ScanDocumentation()
	if err != nil {
		// If scanning fails, use minimal default onboarding context
		context.Onboarding = OnboardingContext{
			LastScanned: time.Now(),
			Methodology: DevelopmentMethodology{
				ProcessName: "spec-driven development", 
				Hierarchy:   []string{"Vision", "Specs", "Tasks", "Implementation"},
				Principles:  []string{"documentation-driven", "traceability"},
			},
			Vision: ProjectVision{
				Name:       "Garden Ecosystem",
				CoreSystem: "Weather System",
			},
			Architecture: ArchitecturalContext{
				Language:     "Go",
				Architecture: "monorepo",
			},
		}
		return nil
	}
	
	context.Onboarding = *onboarding
	return nil
}
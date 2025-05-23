package weather

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SeedValidator validates documentation seeds at various levels
type SeedValidator struct {
	Path string
}

// ValidationResult represents the outcome of seed validation
type ValidationResult struct {
	Valid       bool
	Level       string // "garden", "farm", or "minimal"
	Score       int    // 0-100 quality score
	Issues      []ValidationIssue
	Suggestions []string
	Structure   SeedStructure
}

// ValidationIssue represents a specific problem found
type ValidationIssue struct {
	Severity    string // "error", "warning", "info"
	Path        string
	Message     string
	Suggestion  string
}

// SeedStructure represents what was found
type SeedStructure struct {
	HasDocs           bool
	HasReadme         bool
	HasGitRepo        bool
	HasWeatherContext bool
	Directories       []string
	SpecialFiles      []string
}

// NewSeedValidator creates a validator for the given path
func NewSeedValidator(path string) *SeedValidator {
	return &SeedValidator{
		Path: path,
	}
}

// Validate performs comprehensive seed validation
func (sv *SeedValidator) Validate() (*ValidationResult, error) {
	result := &ValidationResult{
		Valid:       true,
		Issues:      []ValidationIssue{},
		Suggestions: []string{},
	}

	// Detect structure
	structure, err := sv.detectStructure()
	if err != nil {
		return nil, fmt.Errorf("failed to detect structure: %w", err)
	}
	result.Structure = structure

	// Determine level
	result.Level = sv.determineLevel(structure)

	// Validate based on level
	switch result.Level {
	case "farm":
		sv.validateFarm(result)
	case "garden":
		sv.validateGarden(result)
	default:
		sv.validateMinimal(result)
	}

	// Calculate score
	result.Score = sv.calculateScore(result)

	// Set overall validity
	for _, issue := range result.Issues {
		if issue.Severity == "error" {
			result.Valid = false
			break
		}
	}

	return result, nil
}

// detectStructure examines the directory structure
func (sv *SeedValidator) detectStructure() (SeedStructure, error) {
	structure := SeedStructure{
		Directories:  []string{},
		SpecialFiles: []string{},
	}

	// Check for git repository
	gitPath := filepath.Join(sv.Path, ".git")
	if info, err := os.Stat(gitPath); err == nil && info.IsDir() {
		structure.HasGitRepo = true
	}

	// Check for docs directory
	docsPath := filepath.Join(sv.Path, "docs")
	if info, err := os.Stat(docsPath); err == nil && info.IsDir() {
		structure.HasDocs = true

		// Check for README
		readmePath := filepath.Join(docsPath, "README.md")
		if _, err := os.Stat(readmePath); err == nil {
			structure.HasReadme = true
		}

		// Scan subdirectories
		entries, err := os.ReadDir(docsPath)
		if err == nil {
			for _, entry := range entries {
				if entry.IsDir() {
					structure.Directories = append(structure.Directories, entry.Name())
				} else if strings.HasSuffix(entry.Name(), ".md") {
					structure.SpecialFiles = append(structure.SpecialFiles, entry.Name())
				}
			}
		}
	}

	// Check for weather context
	weatherPath := filepath.Join(sv.Path, ".garden", "weather-context.json")
	if _, err := os.Stat(weatherPath); err == nil {
		structure.HasWeatherContext = true
	}

	// Check for CLAUDE.md or AI.md
	for _, aiFile := range []string{"CLAUDE.md", "AI.md", "ai.md", "claude.md"} {
		if _, err := os.Stat(filepath.Join(sv.Path, aiFile)); err == nil {
			structure.SpecialFiles = append(structure.SpecialFiles, aiFile)
		}
		if _, err := os.Stat(filepath.Join(docsPath, aiFile)); err == nil {
			structure.SpecialFiles = append(structure.SpecialFiles, "docs/"+aiFile)
		}
	}

	return structure, nil
}

// determineLevel identifies whether this is a farm, garden, or minimal seed
func (sv *SeedValidator) determineLevel(structure SeedStructure) string {
	// Farm: Has multiple gardens or .farm directory
	farmPath := filepath.Join(sv.Path, ".farm")
	if _, err := os.Stat(farmPath); err == nil {
		return "farm"
	}

	// Check for multiple git repos (farm indicator)
	gitCount := 0
	entries, err := os.ReadDir(sv.Path)
	if err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				gitPath := filepath.Join(sv.Path, entry.Name(), ".git")
				if _, err := os.Stat(gitPath); err == nil {
					gitCount++
				}
			}
		}
	}
	if gitCount >= 2 {
		return "farm"
	}

	// Garden: Has .git repository
	if structure.HasGitRepo {
		return "garden"
	}

	// Otherwise minimal
	return "minimal"
}

// validateFarm checks farm-level requirements
func (sv *SeedValidator) validateFarm(result *ValidationResult) {
	// Must have docs
	if !result.Structure.HasDocs {
		result.Issues = append(result.Issues, ValidationIssue{
			Severity:   "error",
			Path:       "docs/",
			Message:    "Farm-level documentation directory missing",
			Suggestion: "Create docs/ directory with README.md explaining the farm structure",
		})
	}

	// Should have README
	if result.Structure.HasDocs && !result.Structure.HasReadme {
		result.Issues = append(result.Issues, ValidationIssue{
			Severity:   "warning",
			Path:       "docs/README.md",
			Message:    "Farm README missing",
			Suggestion: "Add README.md explaining how gardens relate to each other",
		})
	}

	// Suggestions for farms
	result.Suggestions = append(result.Suggestions, 
		"Consider adding docs/architecture/ for cross-garden design decisions",
		"Document garden relationships and dependencies",
		"Use event system for cross-garden coordination tracking",
	)
}

// validateGarden checks garden-level requirements
func (sv *SeedValidator) validateGarden(result *ValidationResult) {
	// Must have docs
	if !result.Structure.HasDocs {
		result.Issues = append(result.Issues, ValidationIssue{
			Severity:   "error",
			Path:       "docs/",
			Message:    "Documentation directory missing",
			Suggestion: "Create docs/ directory with README.md",
		})
	}

	// Should have README
	if result.Structure.HasDocs && !result.Structure.HasReadme {
		result.Issues = append(result.Issues, ValidationIssue{
			Severity:   "warning",
			Path:       "docs/README.md",
			Message:    "README missing in docs",
			Suggestion: "Add README.md explaining the project",
		})
	}

	// Check for Weather installation
	if !result.Structure.HasWeatherContext {
		result.Issues = append(result.Issues, ValidationIssue{
			Severity:   "info",
			Path:       ".garden/",
			Message:    "Weather System not initialized",
			Suggestion: "Run 'sprout weather' to start tracking",
		})
	}

	// Check for common patterns
	hasSpecs := false
	hasDecisions := false
	for _, dir := range result.Structure.Directories {
		if strings.Contains(strings.ToLower(dir), "spec") {
			hasSpecs = true
		}
		if strings.Contains(strings.ToLower(dir), "decision") || dir == "ADR" {
			hasDecisions = true
		}
	}

	if !hasSpecs {
		result.Suggestions = append(result.Suggestions, 
			"Consider adding specs/ directory for feature planning",
		)
	}
	if !hasDecisions {
		result.Suggestions = append(result.Suggestions,
			"Consider adding decisions/ directory for architectural choices",
		)
	}
}

// validateMinimal checks minimal requirements
func (sv *SeedValidator) validateMinimal(result *ValidationResult) {
	// Even minimal seeds benefit from docs
	if !result.Structure.HasDocs {
		result.Issues = append(result.Issues, ValidationIssue{
			Severity:   "info",
			Path:       "docs/",
			Message:    "No documentation found",
			Suggestion: "Even minimal projects benefit from a docs/README.md",
		})
	}

	result.Suggestions = append(result.Suggestions,
		"Start simple - just add docs/README.md when ready",
		"Weather System can still track git activity without docs",
	)
}

// calculateScore generates a quality score
func (sv *SeedValidator) calculateScore(result *ValidationResult) int {
	score := 100

	// Deduct for issues
	for _, issue := range result.Issues {
		switch issue.Severity {
		case "error":
			score -= 20
		case "warning":
			score -= 10
		case "info":
			score -= 5
		}
	}

	// Bonus points for good patterns
	if result.Structure.HasReadme {
		score += 10
	}
	if len(result.Structure.Directories) > 0 {
		score += 5
	}
	for _, file := range result.Structure.SpecialFiles {
		if strings.Contains(strings.ToUpper(file), "CLAUDE") || strings.Contains(strings.ToUpper(file), "AI") {
			score += 10
			break
		}
	}

	// Cap between 0-100
	if score < 0 {
		score = 0
	}
	if score > 100 {
		score = 100
	}

	return score
}
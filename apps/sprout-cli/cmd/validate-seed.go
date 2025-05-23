package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SeedValidation represents the validation results
type SeedValidation struct {
	Valid       bool                 `json:"valid"`
	Score       int                  `json:"score"`
	Level       string               `json:"level"` // "minimal", "basic", "advanced", "expert"
	Issues      []ValidationIssue    `json:"issues"`
	Suggestions []string             `json:"suggestions"`
	Structure   SeedStructure        `json:"structure"`
}

// ValidationIssue represents a problem found
type ValidationIssue struct {
	Severity string `json:"severity"` // "error", "warning", "info"
	Message  string `json:"message"`
	Path     string `json:"path,omitempty"`
}

// SeedStructure represents what was found
type SeedStructure struct {
	HasDocs       bool     `json:"has_docs"`
	HasReadme     bool     `json:"has_readme"`
	DocCount      int      `json:"doc_count"`
	Directories   []string `json:"directories"`
	SpecialFiles  []string `json:"special_files"`
	EstimatedType string   `json:"estimated_type"` // "minimal", "team", "enterprise"
}

// validateSeed checks if a directory has a valid Seed structure
func validateSeed(rootPath string) (*SeedValidation, error) {
	validation := &SeedValidation{
		Valid:       true,
		Score:       0,
		Issues:      []ValidationIssue{},
		Suggestions: []string{},
	}

	// Check for docs directory
	docsPath := filepath.Join(rootPath, "docs")
	if _, err := os.Stat(docsPath); os.IsNotExist(err) {
		// Try alternative
		docsPath = filepath.Join(rootPath, "documentation")
		if _, err := os.Stat(docsPath); os.IsNotExist(err) {
			validation.Valid = false
			validation.Issues = append(validation.Issues, ValidationIssue{
				Severity: "error",
				Message:  "No docs/ or documentation/ directory found",
			})
			return validation, nil
		}
	}
	validation.Structure.HasDocs = true
	validation.Score += 20

	// Check for README
	readmePath := filepath.Join(docsPath, "README.md")
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		validation.Valid = false
		validation.Issues = append(validation.Issues, ValidationIssue{
			Severity: "error",
			Message:  "No README.md found in docs directory",
			Path:     docsPath,
		})
	} else {
		validation.Structure.HasReadme = true
		validation.Score += 20

		// Check README content
		content, err := os.ReadFile(readmePath)
		if err == nil {
			readmeStr := string(content)
			
			// Check for required sections
			if !strings.Contains(readmeStr, "How We Work") && 
			   !strings.Contains(readmeStr, "how we work") {
				validation.Issues = append(validation.Issues, ValidationIssue{
					Severity: "warning",
					Message:  "README should explain 'How We Work'",
					Path:     readmePath,
				})
			} else {
				validation.Score += 10
			}

			if !strings.Contains(readmeStr, "Directory Structure") && 
			   !strings.Contains(readmeStr, "Structure") {
				validation.Issues = append(validation.Issues, ValidationIssue{
					Severity: "warning",
					Message:  "README should document directory structure",
					Path:     readmePath,
				})
			} else {
				validation.Score += 10
			}
		}
	}

	// Scan documentation structure
	err := filepath.Walk(docsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		relPath, _ := filepath.Rel(docsPath, path)
		
		if info.IsDir() && path != docsPath {
			validation.Structure.Directories = append(validation.Structure.Directories, relPath)
		} else if strings.HasSuffix(path, ".md") {
			validation.Structure.DocCount++
			
			// Check for special files
			base := filepath.Base(path)
			switch strings.ToLower(base) {
			case "claude.md", "ai.md":
				validation.Structure.SpecialFiles = append(validation.Structure.SpecialFiles, "AI instructions")
				validation.Score += 5
			case "contributing.md":
				validation.Structure.SpecialFiles = append(validation.Structure.SpecialFiles, "Contribution guide")
				validation.Score += 5
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Analyze structure type
	analyzeStructureType(validation)
	
	// Generate suggestions
	generateSuggestions(validation)
	
	// Determine level
	determineLevel(validation)

	return validation, nil
}

func analyzeStructureType(v *SeedValidation) {
	dirs := v.Structure.Directories
	
	// Check for enterprise patterns
	for _, dir := range dirs {
		if strings.Contains(dir, "compliance") || strings.Contains(dir, "audit") {
			v.Structure.EstimatedType = "enterprise"
			return
		}
	}
	
	// Check for team patterns
	for _, dir := range dirs {
		if strings.Contains(dir, "specs") || strings.Contains(dir, "decisions") ||
		   strings.Contains(dir, "retrospectives") {
			v.Structure.EstimatedType = "team"
			return
		}
	}
	
	// Default to minimal
	if v.Structure.DocCount <= 3 {
		v.Structure.EstimatedType = "minimal"
	} else {
		v.Structure.EstimatedType = "basic"
	}
}

func generateSuggestions(v *SeedValidation) {
	// Based on what's missing
	if v.Structure.DocCount == 1 {
		v.Suggestions = append(v.Suggestions, 
			"Consider adding more documentation as your project grows")
	}
	
	// Check for common patterns
	hasSpecs := false
	hasDecisions := false
	for _, dir := range v.Structure.Directories {
		if strings.Contains(dir, "spec") {
			hasSpecs = true
		}
		if strings.Contains(dir, "decision") || strings.Contains(dir, "adr") {
			hasDecisions = true
		}
	}
	
	if !hasSpecs && v.Structure.DocCount > 5 {
		v.Suggestions = append(v.Suggestions, 
			"Consider adding a specs/ directory for feature planning")
	}
	
	if !hasDecisions && v.Structure.DocCount > 10 {
		v.Suggestions = append(v.Suggestions, 
			"Consider tracking decisions in a decisions/ or ADR/ directory")
	}
	
	// AI instructions
	hasAI := false
	for _, special := range v.Structure.SpecialFiles {
		if special == "AI instructions" {
			hasAI = true
			break
		}
	}
	if !hasAI {
		v.Suggestions = append(v.Suggestions, 
			"Consider adding CLAUDE.md for AI assistant context")
	}
}

func determineLevel(v *SeedValidation) {
	switch {
	case v.Score >= 80:
		v.Level = "expert"
	case v.Score >= 60:
		v.Level = "advanced"
	case v.Score >= 40:
		v.Level = "basic"
	default:
		v.Level = "minimal"
	}
}

// CLI command implementation
func cmdValidateSeed(args []string) error {
	// Default to current directory
	path := "."
	if len(args) > 0 {
		path = args[0]
	}
	
	// Validate the seed
	validation, err := validateSeed(path)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	
	// Output results
	if jsonOutput {
		data, err := json.MarshalIndent(validation, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		return nil
	}
	
	// Human-friendly output
	fmt.Println("🌱 Seed Validation Report")
	fmt.Println("========================")
	fmt.Printf("📁 Path: %s\n", path)
	fmt.Printf("✓ Valid: %v\n", validation.Valid)
	fmt.Printf("📊 Score: %d/100\n", validation.Score)
	fmt.Printf("🏆 Level: %s\n", validation.Level)
	fmt.Printf("🌿 Type: %s seed\n", validation.Structure.EstimatedType)
	
	if len(validation.Issues) > 0 {
		fmt.Println("\n⚠️  Issues Found:")
		for _, issue := range validation.Issues {
			icon := "ℹ️"
			if issue.Severity == "error" {
				icon = "❌"
			} else if issue.Severity == "warning" {
				icon = "⚠️"
			}
			fmt.Printf("  %s %s\n", icon, issue.Message)
		}
	}
	
	if len(validation.Suggestions) > 0 {
		fmt.Println("\n💡 Suggestions:")
		for _, suggestion := range validation.Suggestions {
			fmt.Printf("  • %s\n", suggestion)
		}
	}
	
	fmt.Printf("\n📈 Structure: %d docs in %d directories\n", 
		validation.Structure.DocCount, 
		len(validation.Structure.Directories))
	
	if !validation.Valid {
		return fmt.Errorf("seed validation failed")
	}
	
	return nil
}
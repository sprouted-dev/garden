// Package weather provides automatic context preservation and intelligence
// for development workflows, serving as the heartbeat of the Sprouted ecosystem.
package weather

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// DocumentIntelligence handles scanning and analyzing project documentation
type DocumentIntelligence struct {
	gardenPath string
}

// NewDocumentIntelligence creates a new document intelligence system
func NewDocumentIntelligence(gardenPath string) *DocumentIntelligence {
	return &DocumentIntelligence{
		gardenPath: gardenPath,
	}
}

// ScanDocumentation scans the docs/ hierarchy and generates onboarding context
func (di *DocumentIntelligence) ScanDocumentation() (*OnboardingContext, error) {
	onboarding := &OnboardingContext{
		LastScanned: time.Now(),
	}

	// Scan for methodology information
	if err := di.scanMethodology(onboarding); err != nil {
		return nil, fmt.Errorf("failed to scan methodology: %w", err)
	}

	// Scan for project vision
	if err := di.scanVision(onboarding); err != nil {
		return nil, fmt.Errorf("failed to scan vision: %w", err)
	}

	// Scan for active work
	if err := di.scanActiveWork(onboarding); err != nil {
		return nil, fmt.Errorf("failed to scan active work: %w", err)
	}

	// Scan for architectural context
	if err := di.scanArchitecture(onboarding); err != nil {
		return nil, fmt.Errorf("failed to scan architecture: %w", err)
	}

	// Generate quick start guide
	if err := di.generateQuickStart(onboarding); err != nil {
		return nil, fmt.Errorf("failed to generate quick start: %w", err)
	}

	return onboarding, nil
}

// scanMethodology scans for development methodology information
func (di *DocumentIntelligence) scanMethodology(onboarding *OnboardingContext) error {
	// Check for CLAUDE.md (project instructions)
	claudeFile := filepath.Join(di.gardenPath, "CLAUDE.md")
	if content, err := di.readFileContent(claudeFile); err == nil {
		onboarding.Methodology = di.extractMethodologyFromClaude(content)
	}

	// Scan workflows directory
	workflowsDir := filepath.Join(di.gardenPath, "docs", "workflows")
	if _, err := os.Stat(workflowsDir); err == nil {
		if err := di.scanWorkflowsDirectory(workflowsDir, onboarding); err != nil {
			return err
		}
	}

	// Scan templates directory
	templatesDir := filepath.Join(di.gardenPath, "templates")
	if templates, err := di.scanTemplatesDirectory(templatesDir); err == nil {
		onboarding.Methodology.Templates = templates
	}

	return nil
}

// scanVision scans for project vision information
func (di *DocumentIntelligence) scanVision(onboarding *OnboardingContext) error {
	visionDir := filepath.Join(di.gardenPath, "docs", "vision")
	
	if _, err := os.Stat(visionDir); os.IsNotExist(err) {
		// Try to extract vision from README.md
		readmeFile := filepath.Join(di.gardenPath, "README.md")
		if content, err := di.readFileContent(readmeFile); err == nil {
			onboarding.Vision = di.extractVisionFromReadme(content)
		}
		return nil
	}

	// Scan vision directory for vision documents
	visionDocs, err := di.scanDocumentDirectory(visionDir, "vision")
	if err != nil {
		return err
	}

	// Extract vision information from the documents
	onboarding.Vision = di.extractVisionFromDocs(visionDocs)
	return nil
}

// scanActiveWork scans for active specs and tasks
func (di *DocumentIntelligence) scanActiveWork(onboarding *OnboardingContext) error {
	activeWork := ActiveWorkContext{}

	// Scan specs directory
	specsDir := filepath.Join(di.gardenPath, "docs", "specs")
	if specs, err := di.scanDocumentDirectory(specsDir, "spec"); err == nil {
		activeWork.ActiveSpecs = specs
	}

	// Scan active tasks directory
	tasksDir := filepath.Join(di.gardenPath, "docs", "tasks", "active")
	if tasks, err := di.scanDocumentDirectory(tasksDir, "task"); err == nil {
		activeWork.ActiveTasks = tasks
		// Infer current phase and context from tasks
		activeWork.CurrentPhase = di.inferCurrentPhase(tasks)
		activeWork.Timeline = di.inferTimeline(tasks)
		activeWork.Priority = di.inferPriority(tasks)
	}

	onboarding.ActiveWork = activeWork
	return nil
}

// scanArchitecture scans for architectural context
func (di *DocumentIntelligence) scanArchitecture(onboarding *OnboardingContext) error {
	arch := ArchitecturalContext{
		DirectoryStructure: make(map[string]string),
	}

	// Infer language from directory structure and files
	arch.Language = di.inferPrimaryLanguage()
	arch.Architecture = di.inferArchitectureType()
	
	// Scan for architectural patterns in code
	arch.KeyPatterns = di.scanCodePatterns()
	arch.Conventions = di.scanCodingConventions()

	// Build directory structure map
	di.scanDirectoryStructure(&arch)

	onboarding.Architecture = arch
	return nil
}

// generateQuickStart generates quick start guidance
func (di *DocumentIntelligence) generateQuickStart(onboarding *OnboardingContext) error {
	quickStart := QuickStartGuide{}

	// Essential docs based on what we found
	quickStart.EssentialDocs = di.identifyEssentialDocs(onboarding)
	
	// Key commands based on project structure
	quickStart.KeyCommands = di.identifyKeyCommands()
	
	// First steps based on methodology and current work
	quickStart.FirstSteps = di.generateFirstSteps(onboarding)
	
	// Common patterns from architecture analysis
	quickStart.CommonPatterns = onboarding.Architecture.KeyPatterns
	
	onboarding.QuickStart = quickStart
	return nil
}

// Helper functions for document scanning and analysis

func (di *DocumentIntelligence) readFileContent(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (di *DocumentIntelligence) scanDocumentDirectory(dirPath, docType string) ([]DocumentSummary, error) {
	var docs []DocumentSummary
	
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return docs, nil
	}

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		if d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		doc, err := di.analyzeDocument(path, docType)
		if err != nil {
			// Log error but continue
			fmt.Printf("Warning: failed to analyze document %s: %v\n", path, err)
			return nil
		}
		
		docs = append(docs, doc)
		return nil
	})

	return docs, err
}

func (di *DocumentIntelligence) analyzeDocument(filePath, docType string) (DocumentSummary, error) {
	stat, err := os.Stat(filePath)
	if err != nil {
		return DocumentSummary{}, err
	}

	content, err := di.readFileContent(filePath)
	if err != nil {
		return DocumentSummary{}, err
	}

	title := di.extractTitleFromContent(content)
	summary := di.generateDocumentSummary(content, docType)
	status := di.extractStatusFromContent(content)

	// Make path relative to garden root
	relPath, _ := filepath.Rel(di.gardenPath, filePath)

	return DocumentSummary{
		Path:        relPath,
		Title:       title,
		Type:        docType,
		Summary:     summary,
		Status:      status,
		LastUpdated: stat.ModTime(),
	}, nil
}

func (di *DocumentIntelligence) extractTitleFromContent(content string) string {
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	return "Untitled"
}

func (di *DocumentIntelligence) generateDocumentSummary(content, docType string) string {
	// Extract first few paragraphs or key sections for summary
	lines := strings.Split(content, "\n")
	var summaryLines []string
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "##") {
			break // Stop at first subsection
		}
		summaryLines = append(summaryLines, line)
		if len(summaryLines) >= 3 { // Limit summary length
			break
		}
	}
	
	summary := strings.Join(summaryLines, " ")
	if len(summary) > 200 {
		summary = summary[:200] + "..."
	}
	
	return summary
}

func (di *DocumentIntelligence) extractStatusFromContent(content string) string {
	// Look for status indicators in the content
	lowerContent := strings.ToLower(content)
	
	if strings.Contains(lowerContent, "in progress") || strings.Contains(lowerContent, "active") {
		return "active"
	}
	if strings.Contains(lowerContent, "completed") || strings.Contains(lowerContent, "done") {
		return "completed"
	}
	if strings.Contains(lowerContent, "planned") || strings.Contains(lowerContent, "pending") {
		return "planned"
	}
	
	return "unknown"
}

// Methodology extraction functions

func (di *DocumentIntelligence) extractMethodologyFromClaude(content string) DevelopmentMethodology {
	methodology := DevelopmentMethodology{
		ProcessName: "spec-driven development",
		Hierarchy:   []string{"Vision", "Specs", "Tasks", "Phases", "Implementation"},
		Principles:  []string{"spec-driven", "documentation hierarchy", "traceability"},
	}

	// Extract specific details from CLAUDE.md
	if strings.Contains(content, "spec-driven development") {
		methodology.ProcessName = "spec-driven development"
	}
	
	if strings.Contains(content, "agentic-development.md") {
		methodology.WorkflowGuide = "/docs/workflows/agentic-development.md"
	}

	return methodology
}

func (di *DocumentIntelligence) scanWorkflowsDirectory(workflowsDir string, onboarding *OnboardingContext) error {
	// Look for workflow documentation
	agenticWorkflow := filepath.Join(workflowsDir, "agentic-development.md")
	if _, err := os.Stat(agenticWorkflow); err == nil {
		onboarding.Methodology.WorkflowGuide = "docs/workflows/agentic-development.md"
	}
	return nil
}

func (di *DocumentIntelligence) scanTemplatesDirectory(templatesDir string) ([]string, error) {
	var templates []string
	
	if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
		return templates, nil
	}

	entries, err := os.ReadDir(templatesDir)
	if err != nil {
		return templates, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
			templates = append(templates, entry.Name())
		}
	}

	return templates, nil
}

// Vision extraction functions

func (di *DocumentIntelligence) extractVisionFromReadme(content string) ProjectVision {
	vision := ProjectVision{
		Name: "Garden Ecosystem",
	}
	
	// Extract key information from README
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(strings.ToLower(line), "weather") && strings.Contains(strings.ToLower(line), "context") {
			vision.CoreSystem = "Weather System - intelligent context preservation"
		}
	}
	
	return vision
}

func (di *DocumentIntelligence) extractVisionFromDocs(docs []DocumentSummary) ProjectVision {
	vision := ProjectVision{
		Name: "Garden Ecosystem",
		CoreSystem: "Weather System",
	}
	
	// Analyze vision documents to extract key information
	for _, doc := range docs {
		if strings.Contains(strings.ToLower(doc.Title), "weather") {
			vision.CoreSystem = "Weather System - intelligent context preservation"
		}
		if strings.Contains(strings.ToLower(doc.Summary), "context loss") {
			vision.Mission = "Eliminate context loss that destroys developer flow state"
		}
	}
	
	return vision
}

// Active work analysis functions

func (di *DocumentIntelligence) inferCurrentPhase(tasks []DocumentSummary) string {
	for _, task := range tasks {
		if strings.Contains(strings.ToLower(task.Title), "weather") && strings.Contains(strings.ToLower(task.Title), "mvp") {
			return "Weather MVP Phase 1"
		}
	}
	return "Development Phase"
}

func (di *DocumentIntelligence) inferTimeline(tasks []DocumentSummary) string {
	for _, task := range tasks {
		if strings.Contains(task.Summary, "4-week") || strings.Contains(task.Summary, "week") {
			return "4-week Weather MVP timeline"
		}
	}
	return "Active development"
}

func (di *DocumentIntelligence) inferPriority(tasks []DocumentSummary) string {
	for _, task := range tasks {
		if strings.Contains(strings.ToLower(task.Title), "weather") {
			return "Weather System foundation"
		}
	}
	return "Current implementation"
}

// Architecture analysis functions

func (di *DocumentIntelligence) inferPrimaryLanguage() string {
	// Check for Go files
	if _, err := os.Stat(filepath.Join(di.gardenPath, "go.mod")); err == nil {
		return "Go"
	}
	
	// Check for TypeScript/JavaScript
	if _, err := os.Stat(filepath.Join(di.gardenPath, "package.json")); err == nil {
		return "TypeScript/JavaScript"
	}
	
	return "Multiple"
}

func (di *DocumentIntelligence) inferArchitectureType() string {
	// Check for monorepo structure
	if _, err := os.Stat(filepath.Join(di.gardenPath, "apps")); err == nil {
		if _, err := os.Stat(filepath.Join(di.gardenPath, "libs")); err == nil {
			return "monorepo"
		}
	}
	
	return "standard"
}

func (di *DocumentIntelligence) scanCodePatterns() []string {
	patterns := []string{}
	
	// Look for common patterns in Go code
	if di.inferPrimaryLanguage() == "Go" {
		patterns = append(patterns, "interface-based design", "dependency injection", "error wrapping")
	}
	
	return patterns
}

func (di *DocumentIntelligence) scanCodingConventions() []string {
	conventions := []string{}
	
	// Check for common convention files
	if _, err := os.Stat(filepath.Join(di.gardenPath, ".golangci.yml")); err == nil {
		conventions = append(conventions, "Go linting with golangci-lint")
	}
	
	return conventions
}

func (di *DocumentIntelligence) scanDirectoryStructure(arch *ArchitecturalContext) {
	dirs := map[string]string{
		"apps":      "Applications and executables",
		"libs":      "Shared libraries and packages", 
		"docs":      "Documentation hierarchy",
		"templates": "Document templates",
		"tools":     "Development tools and utilities",
	}
	
	for dir, desc := range dirs {
		if _, err := os.Stat(filepath.Join(di.gardenPath, dir)); err == nil {
			arch.DirectoryStructure[dir] = desc
		}
	}
}

// Quick start functions

func (di *DocumentIntelligence) identifyEssentialDocs(onboarding *OnboardingContext) []string {
	docs := []string{}
	
	// Add methodology guide if available
	if onboarding.Methodology.WorkflowGuide != "" {
		docs = append(docs, onboarding.Methodology.WorkflowGuide)
	}
	
	// Add CLAUDE.md if it exists
	if _, err := os.Stat(filepath.Join(di.gardenPath, "CLAUDE.md")); err == nil {
		docs = append(docs, "CLAUDE.md")
	}
	
	// Add README.md
	if _, err := os.Stat(filepath.Join(di.gardenPath, "README.md")); err == nil {
		docs = append(docs, "README.md")
	}
	
	return docs
}

func (di *DocumentIntelligence) identifyKeyCommands() []string {
	commands := []string{}
	
	// Check for sprout CLI
	if _, err := os.Stat(filepath.Join(di.gardenPath, "apps", "sprout-cli")); err == nil {
		commands = append(commands, "sprout weather", "sprout weather --for-ai")
	}
	
	// Check for Go
	if di.inferPrimaryLanguage() == "Go" {
		commands = append(commands, "go build", "go test")
	}
	
	return commands
}

func (di *DocumentIntelligence) generateFirstSteps(onboarding *OnboardingContext) []string {
	steps := []string{}
	
	// Based on methodology
	if onboarding.Methodology.ProcessName == "spec-driven development" {
		steps = append(steps, "Review active specs and tasks", "Follow spec-driven development process")
	}
	
	// Based on current work
	if strings.Contains(onboarding.ActiveWork.CurrentPhase, "Weather") {
		steps = append(steps, "Check weather context with `sprout weather --for-ai`", "Review Weather MVP implementation plan")
	}
	
	return steps
}

// Documentation gap detection methods

// DocumentationGaps represents detected gaps in documentation
type DocumentationGaps struct {
	MissingDecisions        []DocumentationNeed      `json:"missing_decisions"`
	UncapturedLessons       []DocumentationNeed      `json:"uncaptured_lessons"`
	ProcessGaps             []DocumentationNeed      `json:"process_gaps"`
	ContextGaps             []DocumentationNeed      `json:"context_gaps"`
	RecentActivities        []ActivitySuggestion     `json:"recent_activities"`
	ConversationalKnowledge []ConversationCapture   `json:"conversational_knowledge"`
}

// DocumentationNeed represents a detected documentation gap
type DocumentationNeed struct {
	Type              string   `json:"type"`           // "decision", "lesson", "process", "context"
	Title             string   `json:"title"`          // "Security vulnerability handling decision"
	Description       string   `json:"description"`    // "Decision made about 18 dependency vulnerabilities"
	SuggestedLocation string   `json:"suggested_location"` // "farm-docs/decisions/"
	Confidence        float64  `json:"confidence"`     // 0-1 confidence in suggestion
	DetectedFrom      []string `json:"detected_from"`  // ["commit messages", "discussion patterns"]
}

// ActivitySuggestion represents recent activity that might need documentation
type ActivitySuggestion struct {
	Activity         string `json:"activity"`       // "installation crisis resolution"
	Timeframe        string `json:"timeframe"`      // "last 2 hours"
	ShouldDocument   bool   `json:"should_document"`
	SuggestedDocType string `json:"suggested_doc_type"` // "lesson", "decision", "process"
	Reasoning        string `json:"reasoning"`      // "Major fix that others should learn from"
}

// DetectDocumentationGaps analyzes weather context for missing documentation
func (w *Weather) DetectDocumentationGaps() DocumentationGaps {
	gaps := DocumentationGaps{}
	
	// Use DocumentIntelligence for scanning
	di := NewDocumentIntelligence(w.RepoPath)
	
	// Load existing documentation to avoid duplicate suggestions
	existingDocs := di.scanExistingDocumentation()
	
	// Detect missing decisions from commits
	gaps.MissingDecisions = w.detectMissingDecisions(existingDocs)
	
	// Detect uncaptured lessons from crisis patterns
	gaps.UncapturedLessons = w.detectLessonsLearned(existingDocs)
	
	// Detect process gaps from workflow patterns  
	gaps.ProcessGaps = w.detectProcessGaps(existingDocs)
	
	// Analyze recent activities for documentation needs
	gaps.RecentActivities = w.suggestRecentActivityDocumentation()
	
	// Load conversational knowledge if available
	gaps.ConversationalKnowledge = w.loadConversationalCaptures()
	
	return gaps
}

// scanExistingDocumentation finds what's already documented
func (di *DocumentIntelligence) scanExistingDocumentation() map[string]bool {
	docs := make(map[string]bool)
	
	// Look for documentation in common locations
	docPaths := []string{
		filepath.Join(di.gardenPath, "docs"),
		filepath.Join(di.gardenPath, "..", "docs"), // Farm level
		filepath.Join(di.gardenPath, "documentation"),
	}
	
	for _, path := range docPaths {
		_ = filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}
			if strings.HasSuffix(file, ".md") {
				// Extract title from file
				content, _ := os.ReadFile(file)
				lines := strings.Split(string(content), "\n")
				for _, line := range lines {
					if strings.HasPrefix(line, "# ") {
						title := strings.TrimPrefix(line, "# ")
						docs[strings.ToLower(title)] = true
						break
					}
				}
			}
			return nil
		})
	}
	
	return docs
}
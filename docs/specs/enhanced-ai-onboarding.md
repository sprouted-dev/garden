# Spec: Enhanced AI Onboarding System

## Vision Reference
Related to: [Weather Context Preservation Vision](../vision/weather-context-preservation.md)

## Requirements

### Functional Requirements
- [ ] Auto-discover all documentation files in project hierarchy
- [ ] Parse and categorize documentation by type (vision, specs, tasks, architecture)
- [ ] Generate structured briefings for AI consumption
- [ ] Provide tiered onboarding commands with different depth levels
- [ ] Synthesize project methodology and conventions from documentation
- [ ] Extract architectural decisions and constraints
- [ ] Combine documentation insights with current weather context
- [ ] Support multi-repository documentation scanning (farm-aware)
- [ ] **Proactive Documentation Gap Detection**: Automatically identify missing documentation needs
- [ ] **Context Capture Suggestions**: Suggest when activities should be documented
- [ ] **Decision Point Detection**: Identify when decisions should be recorded
- [ ] **Lessons Learned Recognition**: Detect crisis resolutions and process improvements that should be captured
- [ ] **Conversational Intelligence**: Automatically capture meaningful real-time discussions that don't result in specs/commits
- [ ] **Interaction Type Classification**: Distinguish between git commits, documentation, and conversational knowledge
- [ ] **Real-Time Discussion Preservation**: Capture problem-solving, architectural insights, and process discoveries as they happen

### Non-Functional Requirements
- Performance: Documentation scan and synthesis <5 seconds for typical projects
- Security: No sensitive data exposure in AI briefings
- Accessibility: Human-readable output that can also be copied/shared

## API/Interface Design
```go
// Enhanced weather commands
type OnboardingLevel int

const (
    QuickContext OnboardingLevel = iota  // --for-ai
    FullOnboarding                       // --onboard-ai  
    DocsOnly                            // --docs-brief
    Comprehensive                       // --comprehensive-brief
)

type DocumentationBrief struct {
    ProjectOverview         ProjectSummary           `json:"project_overview"`
    Architecture           ArchitecturalInfo        `json:"architecture"`
    Methodology            WorkflowInfo             `json:"methodology"`
    CurrentContext         WeatherContext           `json:"current_context"`
    KeyInsights            []string                 `json:"key_insights"`
    DocumentationMap       FileHierarchy            `json:"documentation_map"`
    DocumentationSuggestions DocumentationGaps      `json:"documentation_suggestions"`
}

type DocumentationGaps struct {
    MissingDecisions       []DocumentationNeed      `json:"missing_decisions"`
    UncapturedLessons      []DocumentationNeed      `json:"uncaptured_lessons"`
    ProcessGaps            []DocumentationNeed      `json:"process_gaps"`
    ContextGaps            []DocumentationNeed      `json:"context_gaps"`
    RecentActivities       []ActivitySuggestion     `json:"recent_activities"`
    ConversationalKnowledge []ConversationCapture   `json:"conversational_knowledge"`
}

// NEW: Capture the "invisible knowledge layer"
type ConversationCapture struct {
    Timestamp              time.Time                `json:"timestamp"`
    UserPrompt             string                   `json:"user_prompt"`          // "the github button is still broken"
    AIResponse             string                   `json:"ai_response"`          // Analysis and solution
    ConversationType       string                   `json:"conversation_type"`    // "debugging", "architecture", "meta", "process"
    KeyInsights            []string                 `json:"key_insights"`         // Main takeaways
    Outcome                string                   `json:"outcome"`              // "Problem solved", "Architecture refined", etc.
    ShouldPersist          bool                     `json:"should_persist"`       // Auto-detection if this needs long-term capture
    RelatedFiles           []string                 `json:"related_files"`        // Files discussed/modified
    Context                string                   `json:"context"`              // "pre-launch testing", "farm architecture planning"
}

type DocumentationNeed struct {
    Type                   string                   `json:"type"`           // "decision", "lesson", "process", "context"
    Title                  string                   `json:"title"`          // "Security vulnerability handling decision"
    Description            string                   `json:"description"`    // "Decision made about 18 dependency vulnerabilities"
    SuggestedLocation      string                   `json:"suggested_location"` // "farm-docs/decisions/"
    Confidence             float64                  `json:"confidence"`     // 0-1 confidence in suggestion
    DetectedFrom           []string                 `json:"detected_from"`  // ["commit messages", "discussion patterns"]
}

type ActivitySuggestion struct {
    Activity               string                   `json:"activity"`       // "installation crisis resolution"
    Timeframe              string                   `json:"timeframe"`      // "last 2 hours"
    ShouldDocument         bool                     `json:"should_document"`
    SuggestedDocType       string                   `json:"suggested_doc_type"` // "lesson", "decision", "process"
    Reasoning              string                   `json:"reasoning"`      // "Major fix that others should learn from"
}

type ProjectSummary struct {
    Purpose           string             `json:"purpose"`
    Goals             []string           `json:"goals"`
    CurrentPhase      string             `json:"current_phase"`
    BusinessContext   string             `json:"business_context"`
    SuccessMetrics    []string           `json:"success_metrics"`
}

type ArchitecturalInfo struct {
    KeyDecisions      []Decision         `json:"key_decisions"`
    TechStack         []Technology       `json:"tech_stack"`
    Patterns          []Pattern          `json:"patterns"`
    Constraints       []string           `json:"constraints"`
    ModuleStructure   map[string]string  `json:"module_structure"`
}

type WorkflowInfo struct {
    DevelopmentProcess string            `json:"development_process"`
    CodeConventions   []string          `json:"code_conventions"`
    TestingStrategy   string            `json:"testing_strategy"`
    AICollaboration   string            `json:"ai_collaboration"`
    DocumentationFlow string            `json:"documentation_flow"`
}

// Enhanced weather commands
func (w *Weather) GenerateAIBrief(level OnboardingLevel) (*DocumentationBrief, error)
func (w *Weather) ScanDocumentation(rootPath string) (*FileHierarchy, error)
func (w *Weather) ParseDocumentationHierarchy(files []string) (*DocumentationBrief, error)
```

## Test Scenarios
- Test case 1: New AI assistant joins mid-project and gets complete context in <10 seconds
- Test case 2: Documentation changes are automatically reflected in next AI briefing
- Test case 3: Multi-repository projects provide unified briefing across all repos
- Test case 4: Sensitive business docs are excluded from AI briefings appropriately
- Test case 5: Large projects (100+ docs) still generate briefings within performance limits

## Acceptance Criteria
- [ ] `sprout weather --onboard-ai` provides complete project understanding without manual doc reading
- [ ] Documentation auto-discovery finds all relevant markdown files in hierarchy
- [ ] AI briefing includes project vision, architecture, methodology, and current context
- [ ] Briefing generation completes within 5 seconds for typical projects
- [ ] Output is structured and consumable by AI assistants
- [ ] Supports both single-repo and multi-repo (farm) scenarios
- [ ] Respects privacy boundaries (excludes private business docs when appropriate)

## Implementation Details

### Documentation Discovery Algorithm
```go
func discoverDocumentation(rootPath string) DocumentationTree {
    patterns := []string{
        "README.md", "CLAUDE.md", 
        "docs/**/*.md", 
        "*/docs/**/*.md",
        "**/vision/*.md",
        "**/specs/*.md", 
        "**/tasks/*.md",
        "**/decisions/*.md",
        "**/discussions/*.md",
        "**/lessons/*.md"
    }
    
    // Categorize by file path and content analysis
    // Prioritize by recency and importance
    // Build hierarchical understanding
}
```

### Proactive Documentation Intelligence
```go
func analyzeDocumentationGaps(weather WeatherContext, docs DocumentationTree) DocumentationGaps {
    gaps := DocumentationGaps{}
    
    // Detect missing decisions
    gaps.MissingDecisions = detectMissingDecisions(weather.RecentCommits, docs)
    
    // Detect uncaptured lessons from crisis patterns
    gaps.UncapturedLessons = detectLessonsLearned(weather.RecentActivity, docs)
    
    // Detect process gaps from workflow patterns  
    gaps.ProcessGaps = detectProcessGaps(weather.SessionHistory, docs)
    
    // Detect context gaps for AI onboarding
    gaps.ContextGaps = detectContextGaps(weather.AIInteractions, docs)
    
    // Analyze recent activities for documentation needs
    gaps.RecentActivities = suggestRecentActivityDocumentation(weather)
    
    return gaps
}

func detectMissingDecisions(commits []GitCommit, docs DocumentationTree) []DocumentationNeed {
    var needs []DocumentationNeed
    
    // Pattern: Multiple commit messages discussing options/choices
    // Pattern: Crisis resolution commits without corresponding decision docs
    // Pattern: Architecture changes without decision documentation
    
    for _, commit := range commits {
        if isDecisionPattern(commit.Message) && !hasCorrespondingDecisionDoc(commit, docs) {
            needs = append(needs, DocumentationNeed{
                Type: "decision",
                Title: extractDecisionTitle(commit.Message),
                Description: fmt.Sprintf("Decision made in commit %s", commit.Hash[:8]),
                SuggestedLocation: "farm-docs/decisions/",
                Confidence: calculateConfidence(commit),
                DetectedFrom: []string{"commit message patterns", "crisis resolution activity"},
            })
        }
    }
    
    return needs
}

func detectLessonsLearned(activity []ActivityEvent, docs DocumentationTree) []DocumentationNeed {
    var needs []DocumentationNeed
    
    // Pattern: Crisis resolution followed by successful fix
    // Pattern: Major debugging sessions with resolution
    // Pattern: Process improvements discovered during work
    
    crises := identifyCrisisResolutionPatterns(activity)
    for _, crisis := range crises {
        if !hasCorrespondingLessonDoc(crisis, docs) {
            needs = append(needs, DocumentationNeed{
                Type: "lesson",
                Title: fmt.Sprintf("%s resolution", crisis.Type),
                Description: crisis.Description,
                SuggestedLocation: "farm-docs/lessons/",
                Confidence: crisis.Confidence,
                DetectedFrom: []string{"crisis resolution patterns", "debugging sessions"},
            })
        }
    }
    
    return needs
}

// NEW: Conversational Intelligence Engine
func captureConversationalKnowledge(interactions []AIInteraction) []ConversationCapture {
    var captures []ConversationCapture
    
    for _, interaction := range interactions {
        // Detect meaningful conversations that should be preserved
        if isMeaningfulConversation(interaction) {
            capture := ConversationCapture{
                Timestamp:       interaction.Timestamp,
                UserPrompt:      interaction.UserInput,
                AIResponse:      interaction.AIResponse,
                ConversationType: classifyConversationType(interaction),
                KeyInsights:     extractKeyInsights(interaction),
                Outcome:         determineOutcome(interaction),
                ShouldPersist:   shouldPersistLongTerm(interaction),
                RelatedFiles:    findRelatedFiles(interaction),
                Context:         inferConversationContext(interaction),
            }
            captures = append(captures, capture)
        }
    }
    
    return captures
}

func isMeaningfulConversation(interaction AIInteraction) bool {
    // Patterns that indicate meaningful conversations:
    // - Problem statement → solution discovery
    // - Architecture discussion → design insights  
    // - Process observation → improvement suggestions
    // - Meta-conversation about development methodology
    // - Debugging session with resolution
    // - Requirements clarification leading to specs
    
    meaningfulPatterns := []string{
        "problem.*solution", "issue.*fix", "broken.*working",
        "architecture.*design", "structure.*organization", 
        "process.*improvement", "workflow.*efficiency",
        "should.*document", "capture.*context", "missing.*gap",
    }
    
    text := strings.ToLower(interaction.UserInput + " " + interaction.AIResponse)
    for _, pattern := range meaningfulPatterns {
        if matched, _ := regexp.MatchString(pattern, text); matched {
            return true
        }
    }
    
    // Also check for substantial AI responses (indicating analysis/solutions)
    return len(interaction.AIResponse) > 200 && containsAnalysis(interaction.AIResponse)
}

func classifyConversationType(interaction AIInteraction) string {
    text := strings.ToLower(interaction.UserInput + " " + interaction.AIResponse)
    
    if containsPatterns(text, []string{"bug", "broken", "error", "fix", "debug"}) {
        return "debugging"
    }
    if containsPatterns(text, []string{"architecture", "design", "structure", "organize"}) {
        return "architecture"  
    }
    if containsPatterns(text, []string{"process", "workflow", "methodology", "capture", "document"}) {
        return "process"
    }
    if containsPatterns(text, []string{"interaction", "conversation", "capture", "weather system"}) {
        return "meta"
    }
    
    return "general"
}
```

### Synthesis Engine
```go
func synthesizeProjectBrief(docs DocumentationTree) ProjectSummary {
    // Extract vision from vision/*.md files
    // Parse methodology from workflow docs
    // Identify architectural decisions from specs
    // Combine with current weather context
}
```

### Command Structure
```bash
# Tiered onboarding system
sprout weather --for-ai          # Quick: current focus + immediate context
sprout weather --onboard-ai      # Full: complete briefing with docs synthesis  
sprout weather --docs-brief      # Docs: documentation-focused analysis only
sprout weather --comprehensive   # Everything: includes code structure analysis

# Proactive documentation intelligence
sprout weather --suggest-docs    # Show documentation gaps and suggestions
sprout weather --missing-context # Identify missing decision/lesson documentation
sprout weather --post-session    # Review session for documentation needs
```

## Dependencies
- Core Weather System context tracking
- Git activity monitoring
- Documentation parsing libraries
- File system traversal utilities

## Related Documents
- Vision: [Weather Context Preservation](../vision/weather-context-preservation.md)
- Tasks: [Enhanced AI Onboarding Implementation](../tasks/active/enhanced-ai-onboarding-implementation.md) *(to be created)*
- Implementation: Weather System codebase (`libs/weather/`)

## Real-World Example

Based on the installation crisis resolution session, the enhanced system would automatically suggest:

```bash
$ sprout weather --suggest-docs

🌦️  Documentation Intelligence

📋 Missing Documentation Detected:

🚨 CRITICAL LESSON (Confidence: 95%)
   Title: "Pre-launch installation crisis resolution"  
   Type: Lesson learned
   Location: farm-docs/lessons/
   Detected: Crisis resolution pattern + successful fix + timing critical to launch
   
🔍 MISSING DECISION (Confidence: 80%)
   Title: "Security vulnerability handling strategy"
   Type: Decision
   Location: farm-docs/decisions/  
   Detected: 18 vulnerabilities mentioned, no decision doc found
   
📚 PROCESS GAP (Confidence: 75%)
   Title: "Website deployment workflow learning"
   Type: Process documentation
   Location: farm-docs/processes/
   Detected: Multiple deployment cycles with context switching

⏰ RECENT ACTIVITY SUGGESTIONS:
   • Installation crisis (last 2 hours) → Should document as lesson
   • Go module routing fix → Should document as technical decision  
   • Release creation process → Should document as process improvement

💬 CONVERSATIONAL KNOWLEDGE CAPTURED:
   • GitHub button debugging session → CSS analysis + solution (45 min ago)
   • Farm architecture UX discussion → Root visibility insights (2 hours ago)  
   • Interaction types meta-analysis → Weather System enhancement ideas (just now)
   • Security vulnerability context → Missing decision documentation discovered

🧠 INVISIBLE KNOWLEDGE LAYER:
   These meaningful discussions didn't result in specs/commits but contain critical insights:
   - UI debugging methodologies and common CSS inheritance issues
   - Farm root directory UX challenges and developer workflow impacts
   - Weather System interaction modeling and automatic capture requirements
   - Context gaps in existing documentation structure

Would you like me to create template documentation for any of these? [y/N]
```

This transforms the interaction from reactive documentation to proactive context preservation.

## Success Metrics
- **AI Onboarding Time**: <10 seconds from command to full project understanding
- **Context Completeness**: 95% of architectural decisions captured in briefings
- **Documentation Coverage**: All project markdown files automatically included
- **AI Effectiveness**: New AI assistants immediately productive without manual context
- **Developer Adoption**: 90% of teams use enhanced onboarding for AI collaboration
- **Documentation Proactivity**: 80% of important decisions/lessons automatically detected and suggested for documentation
- **Conversational Knowledge Capture**: 90% of meaningful discussions preserved without manual intervention
- **Context Continuity**: AI assistants can reference previous problem-solving sessions and architectural discussions

## Future Enhancements
- Code structure analysis integration
- Team workflow and convention extraction
- Multi-language documentation support
- Visual documentation generation
- Integration with external documentation systems
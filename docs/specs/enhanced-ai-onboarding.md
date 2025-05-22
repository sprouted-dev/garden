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
    ProjectOverview    ProjectSummary      `json:"project_overview"`
    Architecture       ArchitecturalInfo  `json:"architecture"`
    Methodology        WorkflowInfo       `json:"methodology"`
    CurrentContext     WeatherContext     `json:"current_context"`
    KeyInsights        []string           `json:"key_insights"`
    DocumentationMap   FileHierarchy      `json:"documentation_map"`
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
        "**/tasks/*.md"
    }
    
    // Categorize by file path and content analysis
    // Prioritize by recency and importance
    // Build hierarchical understanding
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

## Success Metrics
- **AI Onboarding Time**: <10 seconds from command to full project understanding
- **Context Completeness**: 95% of architectural decisions captured in briefings
- **Documentation Coverage**: All project markdown files automatically included
- **AI Effectiveness**: New AI assistants immediately productive without manual context
- **Developer Adoption**: 90% of teams use enhanced onboarding for AI collaboration

## Future Enhancements
- Code structure analysis integration
- Team workflow and convention extraction
- Multi-language documentation support
- Visual documentation generation
- Integration with external documentation systems
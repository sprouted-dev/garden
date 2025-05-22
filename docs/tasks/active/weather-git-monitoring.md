# Task: Weather Git Activity Monitoring System

## Spec Reference
Implementation of: [Weather Automatic Intelligence MVP](/docs/specs/weather-automatic-intelligence-mvp.md)

## Description
Build the automatic git integration system that monitors all git activity and triggers weather context updates. This is the core intelligence system that enables zero-maintenance context preservation by learning from developer activity patterns.

## Subtasks
- [ ] Implement git commit detection and parsing system
- [ ] Create automatic git hook installation for gardens
- [ ] Build commit message analysis and smart summary generation
- [ ] Add branch change detection and context switching
- [ ] Implement file change pattern analysis for focus detection
- [ ] Create git history analysis for context reconstruction
- [ ] Add uncommitted changes detection and staging area monitoring
- [ ] Build activity pattern recognition for session boundary detection

## Definition of Done
- [ ] Git hooks automatically install when weather system initializes
- [ ] Every commit triggers automatic weather context update
- [ ] Commit information is accurately parsed and analyzed
- [ ] Branch changes update current focus and context
- [ ] File path analysis correctly identifies development areas
- [ ] Smart summaries are generated from commit messages
- [ ] System works with existing git workflows without disruption
- [ ] Processing completes in <1 second per commit
- [ ] Unit tests cover all git integration scenarios
- [ ] Integration tests validate real git repository workflows
- [ ] Code reviewed and approved

## Dependencies
- Task: Weather Core Data Model (for context updates)
- Git repository with standard structure
- File system permissions for git hook installation

## Estimated Effort
3-4 days

## Status
- [x] Not Started
- [ ] In Progress
- [ ] Under Review
- [ ] Completed

## Technical Implementation Notes

### Git Hook Integration
```bash
# .git/hooks/post-commit
#!/bin/sh
# Automatically update weather context after each commit
sprout weather --update-from-commit HEAD

# .git/hooks/post-checkout  
#!/bin/sh
# Update context when switching branches
sprout weather --update-from-branch-change "$1" "$2" "$3"
```

### Commit Analysis System
```go
type GitCommit struct {
    Hash         string    `json:"hash"`
    Message      string    `json:"message"`
    SmartSummary string    `json:"smartSummary"`
    Timestamp    time.Time `json:"timestamp"`
    FilesChanged []string  `json:"filesChanged"`
    InferredScope string   `json:"inferredScope"`
    Author       string    `json:"author"`
}

// Smart summary generation from commit messages
func GenerateSmartSummary(commit GitCommit) string {
    // Extract action verbs and object nouns
    // Remove technical jargon and file paths
    // Generate human-readable summary
    // Examples:
    // "feat: add user authentication API" -> "Added user authentication system"
    // "fix: resolve login validation bug" -> "Fixed login validation issue"
}
```

### Focus Area Detection
```go
func DetectFocusArea(recentCommits []GitCommit) FocusArea {
    // Analyze file paths to identify primary development areas
    // Weight by recency and frequency of changes
    // Examples:
    // auth/*.go files -> "authentication system"
    // frontend/components/*.tsx -> "user interface components"
    // api/handlers/*.go -> "API endpoints"
}
```

### Activity Pattern Recognition
- Commit frequency analysis for temperature calculation
- File change patterns for focus area detection
- Time gap analysis for session boundary detection
- Commit message sentiment analysis for weather conditions

## Test Cases
1. **Hook Installation**: Git hooks install correctly in any garden
2. **Commit Processing**: Each commit triggers context update within 1 second
3. **Branch Changes**: Switching branches updates focus and context appropriately
4. **Focus Detection**: File path analysis correctly identifies development areas
5. **Smart Summaries**: Commit messages generate readable progress summaries
6. **Session Boundaries**: Activity gaps correctly trigger session detection
7. **Performance**: Large repositories process commits within time limits
8. **Error Handling**: System handles git errors and missing information gracefully

## Integration Points
- Weather context updates via data model from previous task
- CLI commands will consume git analysis results
- Session detection feeds into weather condition mapping
- Focus area detection drives next steps prediction

## Related Documents
- Spec: [Weather Automatic Intelligence MVP](/docs/specs/weather-automatic-intelligence-mvp.md)
- Vision: [Weather Context Preservation System](/docs/vision/weather-context-preservation.md)
- Implementation Plan: [Weather MVP Implementation Plan](/docs/tasks/active/weather-mvp-implementation-plan.md)
- Previous Task: [Weather Core Data Model](/docs/tasks/active/weather-core-data-model.md)
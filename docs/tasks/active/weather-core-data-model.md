# Task: Weather Core Data Model Implementation

## Spec Reference
Implementation of: [Weather Automatic Intelligence MVP](/docs/specs/weather-automatic-intelligence-mvp.md)

## Description
Implement the foundational WeatherContext data structures and JSON storage system that serves as the backbone for all weather intelligence features. This is the critical foundation that enables automatic context preservation and AI integration.

## Subtasks
- [ ] Define WeatherContext Go struct matching TypeScript interface from spec
- [ ] Implement JSON marshaling/unmarshaling with proper field validation
- [ ] Create weather context file management system (create, read, write, update)
- [ ] Add atomic file writes to prevent corruption during updates
- [ ] Implement automatic file validation and corruption recovery
- [ ] Create default context generation for new gardens
- [ ] Add context file migration system for future schema changes

## Definition of Done
- [ ] WeatherContext struct fully implemented with all required fields
- [ ] JSON serialization/deserialization working perfectly
- [ ] Atomic writes prevent file corruption during updates
- [ ] System gracefully handles missing or corrupted context files
- [ ] Default context automatically created for new gardens
- [ ] All data operations complete in <50ms
- [ ] Unit tests achieve 100% coverage for data operations
- [ ] Code reviewed and approved
- [ ] Integration with `.garden/` directory structure confirmed

## Dependencies
- Garden directory structure (`.garden/` folder)
- File system read/write permissions

## Estimated Effort
2-3 days

## Status
- [x] Not Started
- [ ] In Progress  
- [ ] Under Review
- [ ] Completed

## Technical Implementation Notes

### WeatherContext Go Struct
```go
type WeatherContext struct {
    // Metadata
    Updated   time.Time `json:"updated"`
    SessionID string    `json:"sessionId"`
    GardenPath string   `json:"gardenPath"`
    Version   string    `json:"version"`
    
    // Core Intelligence
    CurrentFocus   FocusArea        `json:"currentFocus"`
    RecentProgress ProgressSummary  `json:"recentProgress"`
    NextSteps      NextStepsSuggestion `json:"nextSteps"`
    Weather        WeatherConditions `json:"weather"`
    
    // Git Integration
    Git GitContext `json:"git"`
}
```

### File Operations
- Location: `.garden/weather-context.json`
- Atomic writes using temporary files
- Automatic backup of last good state
- Validation on read with automatic repair

### Error Handling
- Graceful degradation when file is missing/corrupted
- Automatic regeneration from git history when possible
- Comprehensive logging for debugging

## Test Cases
1. **Create new context**: System creates valid default context in new garden
2. **Read existing context**: Correctly deserialize all field types
3. **Update context**: Atomic updates preserve data integrity
4. **Handle corruption**: Recover from various file corruption scenarios
5. **Performance**: All operations complete within performance requirements

## Related Documents
- Spec: [Weather Automatic Intelligence MVP](/docs/specs/weather-automatic-intelligence-mvp.md)
- Vision: [Weather Context Preservation System](/docs/vision/weather-context-preservation.md)
- Implementation Plan: [Weather MVP Implementation Plan](/docs/tasks/active/weather-mvp-implementation-plan.md)
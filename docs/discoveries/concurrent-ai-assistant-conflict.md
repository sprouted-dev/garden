# Architectural Discovery: Concurrent AI Assistant Conflict

**Date**: 2025-05-22
**Discovered By**: Real-world usage pattern
**Discovery Type**: Critical Gap
**Impact Level**: High

## The Discovery

Multiple AI assistants running simultaneously can corrupt Weather context and create conflicting states. This "crazy person" scenario (running 2+ assistants) is actually a common power-user pattern that our architecture doesn't handle.

## Evidence/Examples

### Conflict Scenarios

1. **Weather Context Corruption**
```
Assistant A: Reads weather-context.json
Assistant B: Reads weather-context.json
Assistant A: Updates temperature to 85°F
Assistant B: Updates temperature to 72°F
Result: Last write wins, context lost
```

2. **Event Queue Chaos**
```
Assistant A: Emits discovery event
Assistant B: Emits conflicting discovery event
Orchestrator: Processes both, creates nonsensical correlation
```

3. **Git Branch Confusion**
```
Assistant A: Working on feature-x
Assistant B: Switches to feature-y
Assistant A: Suddenly on wrong branch, commits to feature-y
```

4. **Conversation Capture Overlap**
```
Assistant A: Captures architectural discussion
Assistant B: Captures implementation details
Both: Write to same conversation file
Result: Interleaved, incoherent capture
```

## Why This Matters

1. **Power Users**: Your most engaged users likely run multiple assistants
2. **Productivity Pattern**: Architect + Implementation assistants is natural
3. **Data Integrity**: Corrupted context undermines core value prop
4. **Trust Erosion**: Mysterious state changes confuse users

## Current Architecture Vulnerabilities

```go
// No locking mechanism
func (w *Weather) UpdateContext() error {
    // Read (Assistant A & B simultaneously)
    context, _ := w.contextRepo.Load()
    
    // Modify (Both assistants)
    context.Temperature = newTemp
    
    // Write (Last one wins)
    return w.contextRepo.Save(context)
}
```

## Solution Architecture

### 1. Assistant Identity & Coordination

```go
type AssistantIdentity struct {
    ID          string    `json:"id"`          // Unique per session
    Type        string    `json:"type"`        // "architect", "implementation", etc
    LaunchTime  time.Time `json:"launch_time"`
    LastActive  time.Time `json:"last_active"`
    WorkingOn   string    `json:"working_on"`  // Current focus
}

type AssistantRegistry struct {
    Active map[string]*AssistantIdentity
    Lock   sync.RWMutex
}
```

### 2. Cooperative Locking

```go
// File-based locking for weather context
type WeatherLock struct {
    Path      string
    Assistant string
    Acquired  time.Time
    TTL       time.Duration // Auto-release if stale
}

func (w *Weather) UpdateContextSafe(assistantID string) error {
    lock := w.AcquireLock(assistantID, 30*time.Second)
    defer lock.Release()
    
    // Safe to update
    return w.UpdateContext()
}
```

### 3. Assistant Roles & Boundaries

```yaml
# .garden/assistant-roles.yaml
roles:
  architect:
    can_modify: ["docs/", "specs/"]
    can_read: ["*"]
    prevents: ["code_changes"]
    
  implementation:
    can_modify: ["libs/", "apps/", "tests/"]
    can_read: ["*"]
    prevents: ["spec_changes"]
    
  reviewer:
    can_modify: []
    can_read: ["*"]
    prevents: []
```

### 4. Event Attribution

```go
type Event struct {
    EventID     string `json:"event_id"`
    AssistantID string `json:"assistant_id"`     // NEW
    AssistantType string `json:"assistant_type"` // NEW
    // ... existing fields
}
```

### 5. Conflict Resolution Strategies

```go
type ConflictStrategy string

const (
    LastWriteWins ConflictStrategy = "last_write"      // Current behavior
    MergeChanges  ConflictStrategy = "merge"           // Combine updates
    AssistantPriority ConflictStrategy = "priority"    // By role
    Interactive   ConflictStrategy = "interactive"      // Ask user
)
```

## Implementation Approach

### Phase 1: Detection & Warning
```bash
# Detect concurrent assistants
sprout assistants list
> ⚠️  Multiple assistants detected:
> - Architect (Claude-1): Working on documentation
> - Implementation (Claude-2): Modifying weather.go
> 
> Recommendation: Use different working directories
```

### Phase 2: Basic Protection
- File-based locking for weather context
- Assistant registration in `.garden/assistants/`
- Warning when conflicts detected

### Phase 3: Intelligent Coordination
- Role-based boundaries
- Automatic conflict resolution
- Event stream partitioning
- Collaborative mode

## User Experience Improvements

### Clear Modes
```bash
# Launch with explicit role
sprout assistant --role=architect
sprout assistant --role=implementation

# Weather shows assistant activity
sprout weather
> 🌤️  Weather Report
> Active Assistants: 2
> - Architect: Designing resilience layer
> - Implementation: Building event system
```

### Conflict Prevention
```bash
# Assistant-aware operations
sprout weather update --assistant-id=claude-1
> ✅ Context updated by Architect assistant

# Other assistant sees update
sprout weather --assistant-id=claude-2
> ℹ️  Context updated by Architect (2s ago)
```

## The Opportunity

This "bug" reveals an advanced usage pattern we should embrace:
1. **Team Simulation**: One developer, multiple AI assistants = team productivity
2. **Specialization**: Different assistants for different tasks
3. **Parallel Work**: True concurrent development
4. **Weather Station Feature**: "Team Weather" for assistant coordination

## Related Insights

- Power users will push boundaries in unexpected ways
- Concurrent access is not edge case but advanced feature
- Assistant coordination mirrors human team coordination
- This pattern validates Weather Station team features

---

*Sometimes "crazy" usage patterns reveal the next level of product evolution*
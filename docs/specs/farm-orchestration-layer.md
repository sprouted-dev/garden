# Spec: Farm Orchestration Layer

## Vision Reference
Related to: [Weather Context Preservation Vision](../vision/weather-context-preservation.md)
Addresses: [Farm Invisibility Discovery](../../../docs/discoveries/farm-invisibility-orchestration-need.md)

## Problem Statement

Individual gardens (repositories) operate in isolation due to IDE workflow constraints. Developers open specific repositories, not farm roots, making cross-garden coordination invisible to traditional tools like git hooks. This prevents the Weather System from providing workspace-level intelligence.

## Requirements

### Functional Requirements
- [ ] Event emission protocol for garden-to-farm communication
- [ ] Event aggregation service that correlates cross-garden activities  
- [ ] Farm-level weather synthesis from garden weather contexts
- [ ] Persistent event queue that survives garden restarts
- [ ] Cross-garden pattern detection and alerting
- [ ] Documentation suggestion engine with farm-wide awareness
- [ ] Lightweight garden clients that emit events without blocking
- [ ] Farm-level conversation and decision tracking
- [ ] Event replay capability for AI onboarding
- [ ] Privacy-aware event filtering (public vs private repos)

### Non-Functional Requirements
- Performance: Event emission <10ms to avoid git hook delays
- Reliability: Queue persistence prevents event loss
- Scalability: Handle 1000+ events/day across multiple gardens
- Security: No credential or sensitive data in events

## Architecture Design

### Event Flow
```
Garden Git Hook → Local Event Creation → Event Emission → Farm Queue
                           ↓                                ↓
                    Local Weather Update            Farm Orchestrator
                                                          ↓
                                                   Event Correlation
                                                          ↓
                                                    Farm Weather
                                                    Documentation
                                                    AI Context
```

### Directory Structure
```
farm-root/
├── .farm/                          # Farm orchestration directory
│   ├── orchestrator.db            # SQLite event store
│   ├── events/                    # Event queue directory
│   │   ├── pending/              # Unprocessed events
│   │   └── processed/            # Historical events
│   ├── weather/                  # Farm-level weather
│   │   ├── current.json         # Current farm weather
│   │   └── history/             # Weather snapshots
│   └── conversations/           # Farm-level captures
├── garden/                      # Individual repository
│   ├── .garden/
│   │   ├── weather-context.json
│   │   └── farm-client.json    # Farm connection config
└── sprouted-website/           # Another repository
```

### Event Schema
```json
{
  "event_id": "uuid",
  "timestamp": "2025-05-22T17:30:00Z",
  "garden": "garden",
  "event_type": "commit|documentation|conversation|decision",
  "payload": {
    // Type-specific data
  },
  "context": {
    "branch": "main",
    "weather_temp": 95,
    "session_id": "optional"
  },
  "correlation_hints": ["issue-123", "feature-x"]
}
```

## Implementation Approach

### Phase 1: Event Infrastructure
```go
// Garden-side event emitter
type EventEmitter interface {
    Emit(event Event) error
    QueueEvent(event Event) error  // For offline operation
}

// Farm-side event processor  
type Orchestrator interface {
    ProcessEvents() error
    CorrelateEvents(window time.Duration) []Correlation
    SynthesizeWeather() FarmWeather
}
```

### Phase 2: Git Hook Integration
```bash
#!/bin/bash
# .git/hooks/post-commit
# Lightweight, non-blocking event emission

# Create event
event=$(sprout weather emit-event \
  --type=commit \
  --garden=$(basename $(pwd)) \
  --data=@-)

# Queue for farm processing (returns immediately)
echo "$event" >> ../.farm/events/pending/$(date +%s)-commit.json
```

### Phase 3: Cross-Garden Intelligence
```go
type Correlation struct {
    Gardens      []string
    EventType    string  
    Pattern      string
    Confidence   float64
    Suggestion   string
}

// Example: Detect related changes
// If garden/ commits API change and sprouted-website/ updates API calls
// → Correlation detected, suggest documentation
```

## Test Scenarios

1. **Multi-Garden Commit Correlation**: Changes in `garden/` trigger related updates in `sprouted-website/`
2. **Farm-Level Documentation**: Architectural decision affects multiple gardens
3. **Conversation Preservation**: AI discussion spans multiple repositories  
4. **Offline Operation**: Garden continues working when farm orchestrator is down
5. **Event Replay**: New AI assistant gets full context from event history

## Acceptance Criteria

- [ ] Gardens can emit events without blocking git operations
- [ ] Farm orchestrator correlates events across gardens within 1 minute
- [ ] Weather System provides unified view across all repositories
- [ ] Documentation suggestions consider cross-garden patterns
- [ ] AI onboarding includes farm-level context and patterns
- [ ] System continues functioning when individual gardens are offline

## Security Considerations

- Events contain no credentials or sensitive data
- Private repository events can be filtered
- Event queue has appropriate file permissions
- No external network calls required

## Future Enhancements

- Web dashboard for farm-level visualization
- GitHub integration for issue correlation
- Real-time event streaming via websockets
- Machine learning for pattern detection
- Distributed orchestration for large farms

---

*This specification addresses the core architectural challenge discovered through Weather System development*
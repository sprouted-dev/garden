# Spec: Farm Architecture Planning

**Related to**: [Weather Automatic Intelligence MVP](/docs/specs/weather-automatic-intelligence-mvp.md)  
**Implementation of**: [Sprouted Ecosystem Consolidation Plan](/Users/nutmeg/sprouted/docs/business-strategy/sprouted-ecosystem-consolidation.md)

## Problem Statement

The current Weather System operates within single git repositories (Gardens), but real-world usage has revealed the need for workspace-level intelligence across multiple related repositories. The Sprouted ecosystem itself demonstrates this pattern:

- **sprouted/** (Farm) - Parent workspace with public/private separation
- **garden/** (Garden) - Weather System core repository  
- **sprouted-website/** (Garden) - Website repository
- **docs/** (Private) - Business strategy materials

**Current Limitation**: Weather System only scans within garden git repositories, missing workspace-level context and cross-project relationships.

**Technical Constraint**: Farm directories are not git repositories, so existing git hook-based weather updates don't work at the workspace level.

## Vision: Event-Based Farm Weather System

Transform the current single-repository Weather System into a distributed, event-based architecture that provides intelligent context preservation across multi-repository workspaces (Farms).

### Terminology Evolution

- **Garden**: Single repository with Weather System (current implementation)
- **Farm**: Multi-repository workspace with coordinated weather intelligence
- **Co-Op**: Community-shared patterns, seeds, and weather templates

## Architecture Design

### Core Principles

1. **Event-Driven Coordination**: Gardens emit events, Farm aggregates and correlates
2. **No Git Dependency**: Farm weather works without requiring git hooks at workspace level
3. **Distributed Resilience**: Farm can rebuild state from garden weather files
4. **Incremental Enhancement**: Extends existing garden weather without breaking changes
5. **Developer Workflow Integration**: Farm context must be discoverable from individual repository workflows

### Event-Based Communication

#### Event Sources (Garden Level)
```bash
# Automatic events from garden git hooks
sprout garden event --emit="commit" --data="weather-context.json"
sprout garden event --emit="branch-switch" --data="context-delta.json"
sprout garden event --emit="merge" --data="integration-summary.json"

# AI collaboration events
sprout garden event --emit="ai-session" --data="onboarding-context.json"
sprout garden event --emit="ai-handoff" --data="session-summary.json"

# Manual milestone events
sprout garden event --emit="milestone" --data="phase-complete.json"
sprout garden event --emit="spec-complete" --data="deliverable-summary.json"
```

#### Event Handling (Farm Level)
```bash
# Farm daemon for real-time coordination
sprout farm weather --daemon         # Background process listening for events
sprout farm weather --subscribe       # Real-time weather updates across gardens

# Manual farm operations
sprout farm weather                   # Current workspace-level context
sprout farm weather --replay          # Rebuild farm context from event history
sprout farm weather --sync            # Force synchronization with all gardens
```

### File System Structure

```
sprouted-farm/                       # Farm workspace root
├── .farm/                           # Farm coordination directory
│   ├── config.json                  # Farm configuration and garden registry
│   ├── weather-farm-context.json    # Aggregated workspace weather
│   ├── events/                      # Event log storage
│   │   ├── garden-core/
│   │   │   ├── 2025-05-22-commit-abc123.json
│   │   │   ├── 2025-05-22-ai-session-xyz.json
│   │   │   └── 2025-05-22-milestone-mvp.json
│   │   ├── sprouted-website/
│   │   │   ├── 2025-05-22-deploy-def456.json
│   │   │   └── 2025-05-22-feature-ghi789.json
│   │   └── shared/
│   │       └── 2025-05-22-cross-garden-sync.json
│   └── snapshots/                   # Periodic farm state snapshots
│       ├── 2025-05-22-morning-state.json
│       └── 2025-05-22-evening-state.json
├── garden-core/                     # Individual garden (git repository)
│   ├── .garden/
│   │   └── weather-context.json     # Garden-specific weather
│   └── ...
├── sprouted-website/                # Individual garden (git repository)
│   ├── .garden/
│   │   └── weather-context.json     # Garden-specific weather
│   └── ...
└── docs/                           # Private materials (not a garden)
    └── business-strategy/
```

### Farm Weather Context Schema

```json
{
  "updated": "2025-05-22T02:00:00Z",
  "farmId": "sprouted-ecosystem",
  "farmPath": "/Users/nutmeg/sprouted",
  "version": "1.0.0",
  
  "workspaceIntelligence": {
    "primaryFocus": "open source launch preparation",
    "crossGardenRelationships": [
      {
        "type": "deployment",
        "source": "garden-core",
        "target": "sprouted-website",
        "description": "Weather System showcased on website"
      }
    ],
    "strategicMomentum": 95,
    "coordinationNeeds": [
      "synchronize documentation updates",
      "align release timelines"
    ]
  },
  
  "gardens": {
    "garden-core": {
      "path": "./garden",
      "role": "primary-system",
      "weatherSummary": "Weather MVP development, high momentum",
      "lastSync": "2025-05-22T01:45:00Z",
      "publicPrivacy": "public"
    },
    "sprouted-website": {
      "path": "./sprouted-website", 
      "role": "showcase-platform",
      "weatherSummary": "Website deployment ready",
      "lastSync": "2025-05-22T01:30:00Z",
      "publicPrivacy": "public"
    }
  },
  
  "eventStream": {
    "recentEvents": 15,
    "lastEventTime": "2025-05-22T01:45:00Z",
    "eventVelocity": "high",
    "crossGardenCorrelations": [
      {
        "trigger": "garden-core spec update",
        "effects": ["sprouted-website docs sync needed"]
      }
    ]
  }
}
```

## Command Interface Design

### Farm Management Commands
```bash
# Farm initialization and setup
sprout farm init                      # Initialize farm in current directory
sprout farm add-garden <path> <role>  # Register garden with farm
sprout farm remove-garden <path>      # Unregister garden from farm
sprout farm list-gardens              # Show all registered gardens

# Farm weather operations  
sprout farm weather                   # Show workspace-level weather context
sprout farm weather --for-ai          # AI-friendly farm context
sprout farm weather --daemon          # Start farm weather daemon
sprout farm weather --sync            # Force synchronization with gardens
sprout farm weather --history         # Show farm weather evolution

# Cross-garden operations
sprout farm sync-docs                 # Synchronize documentation across gardens
sprout farm deploy-sequence           # Coordinate multi-garden deployments
sprout farm ai-briefing               # Comprehensive AI onboarding across workspace
```

### Enhanced Garden Commands  
```bash
# Garden registration with farm
sprout garden register-farm <farm-path>    # Connect garden to farm
sprout garden emit-event <type> [data]     # Manually emit farm event

# Farm-aware garden operations
sprout garden weather --include-farm       # Garden weather + farm context
sprout garden sync-to-farm                 # Push garden changes to farm
```

## Critical Design Challenge: Farm Root Visibility

### Problem: Developer Workflow Disconnect
**Issue Identified**: Farm root directories are not git repositories, making them invisible to standard developer workflows.

**Current Challenge**:
```
/sprouted/                    # Farm root - invisible to IDE workflow
├── docs/                     # Private decisions, discussions - hard to access
├── garden/                   # Dev opens THIS in IDE ✓
└── sprouted-website/         # Dev opens THIS separately ✓
```

**Developer Impact**:
- IDEs naturally open individual repositories (`garden/`, `sprouted-website/`)
- Farm-level context (decisions, discussions, cross-repo planning) lives outside daily workflow
- Important context gets "buried" at a level developers rarely see
- Extra cognitive cycles required to access farm-level information

### Solution: Farm Context Surfacing

#### Enhanced Documentation Architecture
```
/sprouted/
├── farm-docs/                # Clear farm-level documentation
│   ├── decisions/           # Technical decisions (security vulns, tech choices)
│   ├── discussions/         # Cross-repo planning sessions  
│   ├── lessons/            # What we learned (like launch prep insights)
│   └── strategy/           # Business docs (private)
├── garden/                  # Main repo with enhanced farm awareness
└── sprouted-website/        # Website repo with farm context links
```

#### Smart Farm Context Discovery
From within any garden, the Weather System should surface relevant farm context:

```bash
# From within garden/ directory
sprout weather --farm-context
> 🌱 Garden: weather-system (garden/)
> 🚜 Farm: sprouted (/path/to/sprouted)
> 📋 Recent farm decisions: 
>   - Security vulnerabilities: Address post-launch (../farm-docs/decisions/security-vulns.md)
>   - Launch strategy: Sequential rollout planned (../farm-docs/decisions/launch-timeline.md)
> 🗣️  Active discussions:
>   - Enhanced AI onboarding implementation
>   - Farm architecture UX challenges
```

#### Garden-Level Integration
Each garden should maintain awareness of its farm context:

```bash
# Garden commands with farm awareness
sprout weather --onboard-ai        # Includes farm-level decisions and context
sprout weather --deployment-status # Knows about farm-wide deployment patterns
sprout garden farm-sync            # Pull latest farm context into garden briefing
```

### Implementation Requirements

1. **Farm Discovery**: Gardens must auto-detect their parent farm
2. **Context Bridging**: Farm-level docs must be accessible from garden workflows
3. **Smart Filtering**: Respect privacy boundaries (exclude private strategy docs when appropriate)
4. **Workflow Integration**: Farm context appears in normal garden weather commands
5. **Documentation Organization**: Clear folder structure for different types of farm-level information

## Implementation Phases

### Phase 1: Foundation (Week 1-2)
- [ ] Farm directory structure and configuration
- [ ] Enhanced documentation architecture (`farm-docs/` structure)
- [ ] Basic event emission from gardens
- [ ] Farm weather aggregation (polling-based)
- [ ] `sprout farm init` and `sprout farm weather` commands

### Phase 2: Event System (Week 3-4)  
- [ ] Event log storage and replay system
- [ ] Farm daemon for real-time event processing
- [ ] Cross-garden relationship detection
- [ ] Enhanced AI onboarding with farm context

### Phase 3: Intelligence Layer (Month 2)
- [ ] Strategic momentum calculation across gardens
- [ ] Coordination need detection and recommendations
- [ ] Cross-garden dependency mapping
- [ ] Farm-level weather conditions and metaphors

### Phase 4: Advanced Features (Month 3+)
- [ ] Farm snapshots and time-travel debugging
- [ ] Multi-farm coordination (Co-Op features)
- [ ] Community sharing of farm patterns
- [ ] Integration with sprouted.dev platform

## Migration Strategy

### Current Setup → Farm Architecture

1. **Preserve Existing Functionality**
   - Garden weather continues working unchanged
   - No breaking changes to existing commands
   - Gradual migration of workspace features

2. **Incremental Enhancement**
   - Add farm commands alongside existing garden commands
   - Opt-in farm features for existing gardens
   - Backward compatibility maintained

3. **Real-World Validation**
   - Use Sprouted ecosystem as reference implementation
   - Validate farm architecture with actual multi-repo workflow
   - Refine based on daily usage patterns

## Technical Considerations

### Event Storage
- **JSON files**: Simple, portable, version-controllable
- **Event sourcing**: Complete audit trail of farm evolution
- **Compression**: Periodic snapshot + recent events for performance

### Synchronization Strategy
- **Eventually consistent**: Gardens can operate independently
- **Conflict resolution**: Last-writer-wins with manual override
- **Network resilience**: Offline-capable with sync on reconnection

### Privacy and Security
- **Public/private boundaries**: Respect garden privacy settings
- **Event filtering**: Exclude sensitive information from farm events
- **Access control**: Farm operations respect garden permissions

## Success Metrics

### Technical Metrics
- **Context restoration time**: <10 seconds for farm-level context
- **Event processing latency**: <1 second from garden to farm
- **Synchronization accuracy**: 99% consistency across gardens

### User Experience Metrics  
- **AI onboarding effectiveness**: Complete workspace context in single command
- **Cross-garden coordination**: Reduced manual synchronization effort
- **Development velocity**: Faster context switching between related projects

### Ecosystem Metrics
- **Farm adoption**: Number of workspaces using farm architecture
- **Pattern sharing**: Community reuse of successful farm configurations
- **Platform integration**: Farm data feeding sprouted.dev insights

## Next Steps

1. **Document current Sprouted setup** as first farm reference implementation
2. **Design farm configuration schema** based on real workspace needs  
3. **Implement basic event emission** from existing gardens
4. **Build farm weather aggregation** to prove concept
5. **Validate with daily usage** of Sprouted ecosystem development

This specification captures the evolution from single-repository Gardens to multi-repository Farms, enabling the next phase of the Sprouted ecosystem while maintaining the weather-centric approach that makes context preservation so effective.

## Related Documents

- [Weather Automatic Intelligence MVP](/docs/specs/weather-automatic-intelligence-mvp.md)
- [Sprouted Ecosystem Consolidation Plan](/Users/nutmeg/sprouted/docs/business-strategy/sprouted-ecosystem-consolidation.md)
- [Agentic Development Workflow](/docs/workflows/agentic-development.md)
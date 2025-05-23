# Architectural Discovery: Seeds as Innovation Foundations

**Date**: 2025-05-22
**Discovered By**: Architecture discussion
**Discovery Type**: Pattern Evolution
**Impact Level**: Critical

## The Discovery

Seeds are not templates - they are starter structures that allow developers to define their own documentation organization based on their workflow. Each seed contains a docs directory with any structure the developer chooses, requiring only a README explaining their approach. This transforms rigid templates into flexible innovation foundations.

## Context That Led to Discovery

While analyzing documentation patterns across the farm, we recognized that forcing developers into predefined templates constrains innovation. Different workflows need different structures, but the Weather System needs to work with all of them.

## Evidence/Examples

Current rigid approach:
```
templates/
├── spec-template.md
├── task-template.md
└── vision-template.md
```

Seeds approach:
```
seeds/
├── agile-sprint-seed/
│   └── docs/
│       ├── README.md (explains sprint-based workflow)
│       ├── sprints/
│       ├── retrospectives/
│       └── velocity/
├── research-driven-seed/
│   └── docs/
│       ├── README.md (explains research-first approach)
│       ├── hypotheses/
│       ├── experiments/
│       └── findings/
└── rapid-prototype-seed/
    └── docs/
        ├── README.md (explains fail-fast methodology)
        ├── prototypes/
        ├── iterations/
        └── pivots/
```

## Why This Matters

1. **Developer Autonomy**: Teams can work in their natural flow
2. **Innovation Preservation**: New workflows become shareable seeds
3. **Weather Adaptation**: System learns from diverse approaches
4. **Seed Exchange**: Future marketplace of proven workflows

## Weather System Requirements for Seeds

### Minimum Seed Structure
```
seed-name/
└── docs/
    └── README.md (required)
```

### Required README Sections
```markdown
# [Seed Name] Workflow

## How We Work
[Description of the workflow/methodology]

## Directory Structure
[Explanation of the docs organization]

## Key Concepts
[Core ideas that drive this workflow]

## Weather Integration Points
- Where progress happens: [directories/files]
- What indicates momentum: [patterns]
- How decisions are captured: [location/format]
- Where conversations matter: [integration points]
```

### Weather System Discovery Patterns

The Weather System needs to detect:

1. **Activity Hotspots** - Where most changes occur
2. **Decision Points** - Where choices are documented
3. **Progress Indicators** - What shows forward movement
4. **Knowledge Artifacts** - Where insights accumulate
5. **Conversation Anchors** - Where discussions reference

### Seed Metadata Schema
```json
{
  "seed": {
    "name": "agile-sprint-seed",
    "version": "1.0.0",
    "workflow_type": "iterative",
    "author": "team-name",
    "weather_hints": {
      "progress_indicators": ["sprints/current/*", "velocity.md"],
      "decision_locations": ["retrospectives/", "architecture/decisions/"],
      "conversation_anchors": ["discussions/", "planning/"],
      "momentum_patterns": {
        "high": "daily updates in current sprint",
        "low": "no sprint activity for 3 days"
      }
    }
  }
}
```

## Implications

### Immediate Implications
- Weather System must become structure-agnostic
- Need seed validation/testing framework
- Documentation detection must be pattern-based, not path-based
- README parsing becomes critical capability

### Long-term Implications
- Seed Exchange becomes innovation marketplace
- Weather learns optimal patterns from successful seeds
- AI assistants adapt to seed-specific workflows
- Cross-pollination of methodologies

## Related Decisions Needed

- [ ] Define minimum viable seed specification
- [ ] Design seed discovery/registration protocol
- [ ] Create Weather adapter interface for seeds
- [ ] Build seed validation framework
- [ ] Plan Seed Exchange architecture

## Implementation Strategy

### Phase 1: Seed Specification
```go
type Seed struct {
    Name        string            `json:"name"`
    Version     string            `json:"version"`
    ReadmePath  string            `json:"readme_path"`
    WeatherHints WeatherHints     `json:"weather_hints"`
    Metadata    map[string]interface{} `json:"metadata"`
}

type WeatherHints struct {
    ProgressIndicators []string          `json:"progress_indicators"`
    DecisionLocations  []string          `json:"decision_locations"`
    ConversationAnchors []string         `json:"conversation_anchors"`
    MomentumPatterns   map[string]string `json:"momentum_patterns"`
}
```

### Phase 2: Weather Adaptation Layer
```go
type SeedAdapter interface {
    DetectSeedType(repoPath string) (*Seed, error)
    ParseReadme(readmePath string) (*WorkflowDescription, error)
    FindActivityHotspots(seed *Seed) ([]string, error)
    ExtractProgress(seed *Seed) (*ProgressSummary, error)
    LocateDecisions(seed *Seed) ([]Decision, error)
}
```

### Phase 3: Seed Exchange Protocol
- Seed publishing/versioning
- Usage analytics
- Rating/feedback system
- Forking/evolution tracking

## Preservation Notes

**Why document this**: Seeds represent a fundamental shift from prescriptive to descriptive documentation systems. This enables true innovation while maintaining intelligibility.

**Cross-references**: 
- [Human-Centered Development Vision](../future-considerations/human-centered-development-vision.md)
- [Weather Context Preservation](../garden/docs/vision/weather-context-preservation.md)
- [Farm Orchestration Layer](../garden/docs/specs/farm-orchestration-layer.md)

**Keywords**: seeds, templates, innovation, workflows, documentation-structure, weather-adaptation, seed-exchange

---

*This discovery reveals how flexibility in structure enables innovation in process*
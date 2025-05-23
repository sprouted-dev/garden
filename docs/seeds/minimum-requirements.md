# Minimum Seed Requirements for Weather System

## What is a Seed?

A Seed is the documentation and workflow foundation that enables the Weather System to understand and track your project. Think of it as the "soil" in which your project grows - it provides the structure and nutrients (context) that the Weather System needs to function.

## The Farm/Garden/Seed Architecture

We use a three-level workspace architecture:

1. **Farm**: Multi-repository workspace (no `.git` directory)
2. **Garden**: Individual repository with `.git` tracking
3. **Seed**: Documentation structure within each workspace level

This architecture allows context preservation at multiple levels, which is why we built the event-based system.

## Minimum Requirements by Level

### For Gardens (Git Repositories)
```
garden/
├── .git/                    # Git repository for activity tracking
└── docs/                    # Documentation directory
    └── README.md            # Garden overview (required)
```

### For Farms (Multi-Repo Workspaces)
```
farm/
├── docs/                    # Farm-level documentation
│   └── README.md            # Farm overview (required)
├── garden-1/                # Individual gardens
├── garden-2/
└── .farm/                   # Farm event tracking (created by Weather)
    └── events/              # Cross-garden event log
```

### For Seeds (Documentation Structure)
```
any-workspace/
└── docs/
    └── README.md            # Explains your workflow
```

## Why These Are Required

1. **For Gardens**: `.git/` enables automatic activity tracking via git hooks
2. **For Farms**: Event system tracks cross-garden activities and coordination
3. **For All Levels**: `docs/` directory provides context and workflow documentation

## Lessons Learned (From Our Own Experience)

### 1. Farm-Level Documentation is Critical
We discovered that workspace-level documentation often becomes invisible to normal IDE workflows. This is why the Weather System specifically scans for Farm-level docs.

### 2. Event System Fills the Git Gap
Without `.git` at the Farm level, we need the event system to track:
- Cross-garden coordination
- Farm-level decisions
- Workspace-wide context changes

### 3. Seeds Must Be Flexible
Our own experience shows that rigid templates don't work. Each level needs its own documentation style:
- **Farm Seeds**: Strategic vision, cross-project coordination
- **Garden Seeds**: Implementation details, technical specs
- **Personal Seeds**: Individual workflow preferences

## Best Practices

### 1. Start Where You Are
- If you have a single repo, start with a Garden Seed
- If you have multiple repos, create a Farm to coordinate them
- Don't force a structure that doesn't match your reality

### 2. Document Your Actual Workflow
```markdown
# How We Actually Work (Not How We Think We Should)

## Real Workflow
- Where decisions really get made
- How work actually flows between people/repos
- What we check every morning

## Ideal vs Reality
- What we aspire to: [vision]
- What we do now: [current]
- Steps to bridge: [plan]
```

### 3. Use Events for Farm Coordination
Since Farms lack git tracking, emit events for significant activities:
```go
// Example: Cross-garden decision
event := FarmEvent{
    Type: "decision",
    Garden: "weather-service",
    Impact: []string{"api-gateway", "frontend"},
    Description: "Switched to event-based architecture",
}
```

## Common Patterns We've Discovered

### Farm Patterns
- `docs/` - Strategic documentation
- `gardens/` or direct garden directories
- `.farm/events/` - Event log (auto-created)
- `CLAUDE.md` - AI context at farm level

### Garden Patterns
- `docs/` - Technical documentation
- `apps/` and `libs/` - Code organization
- `.garden/` - Garden-specific Weather data

### Seed Evolution Patterns
1. **Sprouting**: Minimal docs/ with README
2. **Growing**: Add workflow-specific directories
3. **Mature**: Full documentation hierarchy
4. **Reproducing**: Patterns become new Seeds for other projects

## Testing Your Multi-Level Setup

```bash
# At Farm level (no git)
cd ~/my-farm
sprout farm status  # Future command

# At Garden level (with git)
cd ~/my-farm/garden-1
sprout weather

# Check cross-garden context
sprout weather --farm-context  # Future enhancement
```

## Current Reality Check

**What Works Now**:
- Garden-level Weather tracking (git-based)
- Basic event system for Farms
- Documentation intelligence at all levels

**What's In Progress**:
- Full Farm orchestration
- Cross-garden correlation
- Event-based activity tracking

**What's Planned**:
- Automated Farm/Garden discovery
- Seed template generation
- Multi-level context aggregation

## Next Steps

1. **For Existing Projects**: Add docs/README.md at your current level
2. **For Multi-Repo Projects**: Create a Farm structure with events
3. **For Teams**: Document your actual workflow, not the ideal one
4. **For Everyone**: Let your Seeds evolve with your needs

Remember: We're building this while using it. The Weather System adapts to how developers actually work, including complex multi-repository setups like our own.
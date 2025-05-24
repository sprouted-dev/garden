# Weather/Seed Separation Implementation Plan

*Created: May 24, 2025*
*Tornado: seed-separation-20250524*

## Mission

Transform Weather System from Sprouted-specific to universal infrastructure, with Sprouted Seed as the reference implementation.

## Current State Analysis

### Weather System (libs/weather/)
Currently contains Sprouted-specific:
- `velocity.go` - Features/hour tracking (Sprouted metric)
- Bamboo growth patterns (Sprouted philosophy)
- Partner terminology (Sprouted convention)
- WEMA concepts (Sprouted methodology)

### Generic Capabilities (Keep in Weather)
- Git activity monitoring
- Context preservation (JSON)
- File change tracking
- Backup/recovery systems
- Event streaming
- Temporal anchoring
- Documentation discovery

## Implementation Phases

### Phase 1: Create Seed Structure (Today)
```
garden/
├── .seed/
│   ├── config.json           # Sprouted methodology
│   ├── patterns/
│   │   ├── branches.json     # tornado/*, storm/*
│   │   └── metrics.json      # features/hour
│   ├── prompts/
│   │   ├── storm-detect.md   # "What if..." detection
│   │   └── partner.md        # Partner onboarding
│   └── philosophy/
│       └── README.md         # Links to our philosophy
```

### Phase 2: Extract Sprouted-Specific (Today)
Move from Weather to Seed:
- Velocity tracking → `.seed/plugins/velocity/`
- WEMA concepts → `.seed/plugins/wema/`
- Partner terminology → `.seed/config.json`
- Storm patterns → `.seed/patterns/`

### Phase 3: Create Seed Interface (Today)
```go
// weather/interfaces.go
type Seed interface {
    GetConfig() SeedConfig
    GetPatterns() PatternConfig
    GetMetrics() MetricsConfig
    HandleEvent(Event) error
}

// weather/seed_loader.go
func LoadSeed(path string) (Seed, error)
```

### Phase 4: Weather Adapts to Seeds (Tomorrow)
- Read seed configuration on startup
- Apply seed-specific patterns
- Load seed plugins if present
- Use seed prompts for AI

### Phase 5: Test with Alternative Seeds (Tomorrow)
Create mock seeds to validate:
- `test-seeds/traditional-agile/`
- `test-seeds/kanban-flow/`
- `test-seeds/chaos-startup/`

## File-by-File Changes

### In libs/weather/

**Remove from weather.go:**
- Velocity tracking initialization
- Partner-specific terminology
- WEMA references

**Add to weather.go:**
```go
type WeatherSystem struct {
    ctx    context.Context
    config Config
    seed   Seed  // Loaded from .seed/
}
```

**New file: seed_system.go**
- Seed loading logic
- Plugin system
- Pattern matching
- Metrics routing

### In .seed/ (New)

**config.json:**
```json
{
  "name": "sprouted",
  "version": "1.0.0",
  "methodology": "storm-driven-development",
  "plugins": ["velocity", "wema"],
  "patterns": {
    "tornado": "tornado/*",
    "storm": "storm/*"
  }
}
```

**plugins/velocity/** (Moved from weather)
- All velocity tracking code
- Features/hour calculation
- Momentum tracking

## Cross-Repo Coordination

### weather-station/
- Add seed configuration endpoint
- Remove hardcoded velocity metrics
- Make dashboard seed-aware

### sprouted-website/
- Update "How It Works" section
- Add "Weather vs Seed" explanation
- Create "Build Your Own Seed" guide

## Testing Strategy

1. **Unit Tests**: Ensure Weather works without any Seed
2. **Integration Tests**: Load Sprouted Seed, verify behavior
3. **Alternative Seeds**: Test with different methodologies
4. **Cross-Repo**: Ensure all repos work together

## Success Metrics

- [ ] Weather has zero compile-time Sprouted dependencies
- [ ] Sprouted Seed can be deleted without breaking Weather
- [ ] Alternative Seeds load and function correctly
- [ ] All repos coordinate through Seed configuration
- [ ] Documentation clearly separates infrastructure from methodology

## Today's Goals

1. Create `.seed/` structure in garden
2. Move velocity.go to seed plugin
3. Create basic seed loader
4. Update weather.go to use seed
5. Document in .storm/ directory

## Risk Mitigation

- **Over-abstraction**: Keep it simple, Seeds are JSON + plugins
- **Breaking changes**: Maintain compatibility layer temporarily
- **Complex plugins**: Start with simple file-based plugins
- **Documentation lag**: Update as we go

This plan separates concerns while maintaining functionality, creating a truly universal Weather System.
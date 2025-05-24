# Spec: Farm-Level Weather Protection

## Current Gap

The Weather System currently protects at the Garden (repository) level but not across the Farm (organization) level. This means:

- ✅ Each Garden's context is protected
- ❌ Farm-wide coordination state is not protected
- ❌ Cross-Garden relationships are not backed up

## Proposed Farm-Level Protection

### 1. Farm Weather State Protection
```
.farm/
├── weather/
│   ├── current.json          # Current farm weather
│   ├── current.shadow.json   # Shadow copy
│   └── backups/              # Historical farm states
└── events/
    └── pending/              # Unprocessed events
```

### 2. Cross-Garden Resilience
- Backup Farm-level correlations
- Preserve cross-repository relationships
- Maintain organization-wide context

### 3. Hierarchical Recovery
```bash
# Garden-level (existing)
sprout weather recover

# Farm-level (proposed)
sprout farm recover

# Cascade recovery
sprout farm recover --cascade  # Recovers farm + all gardens
```

## Implementation Considerations

### Pros
- Complete organizational protection
- Cross-project state preservation
- True multi-repo resilience

### Cons
- Added complexity
- Farm directory requirement
- Permission considerations

## Decision Point

Do we need Farm-level protection for Phase 2, or is Garden-level sufficient?

Current thinking: **Garden-level is sufficient for now** because:
1. Most developers work in single repositories
2. Farm orchestration is still experimental
3. Adds complexity without proven need

We can add Farm protection in Phase 3 if users request it.

---

*Status: PROPOSED - Not implemented in Phase 2*
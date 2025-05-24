# Tornado Coordination Protocol

*Created: May 24, 2025*

## Overview

When a creative tornado spans multiple repositories, we need coordinated branch management to ensure consistency and proper integration.

## Pre-Tornado Checklist

Before creating tornado branches:

1. **Ensure Clean State**
   ```bash
   # For each repo:
   git status  # Should be clean
   git pull origin main  # Latest main
   ```

2. **Create Tornado Branches**
   ```bash
   # Naming convention: tornado/[concept]-[date]
   git checkout -b tornado/seed-separation-20250524
   ```

3. **Document Storm Path**
   Create `.storm/` directory at Farm level for cross-repo documentation

## Multi-Repo Tornado Structure

```
sprouted/ (farm)
├── .storm/
│   ├── active/
│   │   └── seed-separation-20250524.md
│   └── completed/
├── garden/ (tornado/seed-separation-20250524)
├── sprouted-website/ (tornado/seed-separation-20250524)
├── weather-station/ (tornado/seed-separation-20250524)
└── docs/ (main - business docs stay stable)
```

## Tornado Phases

### Phase 1: Storm Formation (Current)
- Identify cross-repo impacts
- Create coordinated branches
- Set up .storm/ tracking

### Phase 2: Creative Destruction
- Rapid experimentation
- Break existing patterns
- Document insights in .storm/

### Phase 3: Eye of the Storm
- Pause and reflect
- Categorize changes
- Plan integration

### Phase 4: Recovery (WEMA)
- Cherry-pick valuable changes
- Update documentation
- Merge or archive branches

## Today's Tornado: Seed Separation

### Scope
Separating Weather System (generic infrastructure) from Sprouted Seed (our methodology)

### Affected Repos

**garden/**
- Refactor libs/weather/ to remove Sprouted-specific code
- Create .seed/ directory for Sprouted methodology
- Update CLI to read seed configuration

**weather-station/**
- Ensure MCP server works with any Seed
- Remove hardcoded Sprouted patterns
- Add seed configuration support

**sprouted-website/**
- Update documentation to explain Weather vs Seed
- Show Sprouted as example, not only way
- Add "Create Your Own Seed" section

**docs/** (private)
- Keep business strategy separate
- Document tornado outcomes
- Track multi-repo coordination

## Implementation Order

1. **Set up .storm/ directory** (Farm level)
2. **Create tornado branches** (all repos)
3. **Document current state** (.storm/active/)
4. **Begin refactoring** (Weather genericization)
5. **Create Sprouted Seed** (reference implementation)
6. **Update integrations** (cross-repo compatibility)
7. **Test with mock Seeds** (validate genericness)

## Success Criteria

- Weather System has zero Sprouted-specific code
- Sprouted Seed fully defines our methodology
- Other Seeds can be created without changing Weather
- All repos remain coordinated
- Clear documentation of changes

## Risk Management

- **Branch Divergence**: Daily sync meetings
- **Integration Conflicts**: Test cross-repo early
- **Scope Creep**: Focus on Weather/Seed separation only
- **Documentation Lag**: Update .storm/ in real-time

## Communication

- All decisions documented in .storm/
- Major changes discussed before implementation
- Progress tracked in tornado branch commits
- Final integration planned collaboratively

This protocol ensures our tornado creates focused innovation while maintaining cross-repo coherence.
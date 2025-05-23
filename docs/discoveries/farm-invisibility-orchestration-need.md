# Architectural Discovery: Farm Root Invisibility & Orchestration Need

**Date**: December 2024 (Rediscovered May 2025)
**Discovered By**: Multiple AI Assistants + Human Developer
**Discovery Type**: Architectural Gap
**Impact Level**: Critical

## The Discovery

Farm root directories are invisible to normal developer workflows because IDEs naturally open individual repositories (gardens). This creates a fundamental coordination challenge where git hooks and Weather Systems in individual gardens cannot see or coordinate across repository boundaries.

## Context That Led to Discovery

While designing the Weather System's proactive documentation features, we realized that git hooks in individual gardens would be unable to:
- Detect cross-repository patterns
- Coordinate documentation at the farm level
- Share context between gardens
- Aggregate weather data across the workspace

## Evidence/Examples

1. **Developer Workflow Reality**: Developers use `code garden/` or `code sprouted-website/`, never `code .` from farm root
2. **Git Hook Isolation**: Hooks in `garden/.git/hooks/` can't access `../sprouted-website/`
3. **Multiple Rediscoveries**: At least 2 AI assistants independently identified this same issue
4. **Context Loss**: Each new AI session risks redesigning solutions without awareness of this constraint

## Why This Matters

- **Breaks Multi-Repo Coordination**: Can't track related changes across repositories
- **Limits Weather System**: Cannot provide workspace-level intelligence
- **Documentation Gaps**: Farm-level decisions and patterns go undocumented
- **Repeated Work**: Same architectural insights rediscovered in each session

## Previous Attempts

- December 2024: First assistant identified need for "enhanced documentation architecture"
- May 2025: Current assistant rediscovered same issue when designing Weather intelligence
- Both arrived at similar conclusion: need for orchestration layer

## Implications

### Immediate Implications
- Cannot implement farm-aware git hooks without orchestration layer
- Weather System needs event-based architecture from the start
- Documentation system must be designed with farm coordination in mind

### Long-term Implications
- Need persistent orchestration service or daemon
- Event-driven architecture becomes core, not optional
- Farm-level intelligence requires new architectural patterns

## Related Decisions Needed

- [ ] Event protocol design for garden-to-farm communication
- [ ] Orchestration layer architecture (daemon vs. on-demand)
- [ ] Farm-level weather aggregation strategy
- [ ] Cross-garden event correlation approach
- [ ] Storage mechanism for farm-level data

## Preservation Notes

**Why document this**: This is a fundamental architectural constraint that shapes the entire Weather System design. Without preserving this discovery, every new contributor (human or AI) wastes time rediscovering it.

**Cross-references**: 
- Weather System Vision
- Enhanced AI Onboarding Spec
- Farm Architecture Planning

**Keywords**: farm visibility, orchestration, multi-repo, git hooks, workspace coordination, event architecture

---

*This discovery is part of the Weather System's knowledge preservation layer*
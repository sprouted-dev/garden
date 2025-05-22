# Task: Enhanced AI Onboarding Implementation

## Spec Reference
Implementation of: [Enhanced AI Onboarding System](../../specs/enhanced-ai-onboarding.md)

## Description
Implement the enhanced AI onboarding system that automatically discovers, parses, and synthesizes all project documentation to provide comprehensive AI briefings. This eliminates the manual "read all docs" step and ensures AI assistants get complete project understanding instantly.

## Subtasks
- [ ] **Documentation Discovery Engine**
  - [ ] File pattern matching for markdown discovery
  - [ ] Hierarchical categorization (vision, specs, tasks, etc.)
  - [ ] Multi-repository support for farm scenarios
  - [ ] Privacy boundary detection (exclude private docs)

- [ ] **Documentation Parser and Synthesizer**
  - [ ] Extract project vision and goals from vision docs
  - [ ] Parse architectural decisions from specs
  - [ ] Identify methodology from workflow documents
  - [ ] Generate structured project summary

- [ ] **Enhanced Command Interface**
  - [ ] Extend `sprout weather --onboard-ai` with documentation synthesis
  - [ ] Add `sprout weather --docs-brief` for docs-only analysis
  - [ ] Add `sprout weather --comprehensive` for full briefing
  - [ ] Maintain backward compatibility with existing commands

- [ ] **Structured Output Generation**
  - [ ] JSON format for programmatic consumption
  - [ ] Human-readable markdown format
  - [ ] AI-optimized briefing format
  - [ ] Performance optimization for large documentation sets

- [ ] **Integration with Current Weather**
  - [ ] Combine documentation insights with current weather context
  - [ ] Cross-reference current focus with documented architecture
  - [ ] Provide contextual next steps based on documentation

## Definition of Done
- [ ] All documentation files automatically discovered and categorized
- [ ] Comprehensive AI briefing generated in <5 seconds
- [ ] New AI assistants get complete project understanding without manual steps
- [ ] Enhanced commands integrate seamlessly with existing weather interface
- [ ] Tests cover discovery, parsing, synthesis, and output generation
- [ ] Documentation updated with new command usage
- [ ] Performance requirements met for typical projects (100+ docs)
- [ ] Privacy boundaries respected for sensitive documentation

## Dependencies
- Core Weather System data model (context.go)
- Git monitoring infrastructure (git.go)
- Documentation intelligence system (docs_intelligence.go)
- File system utilities and markdown parsing libraries

## Estimated Effort
**2-3 weeks** (can be done in parallel with other MVP features)

### Breakdown:
- Documentation discovery: 3-4 days
- Parser and synthesizer: 5-7 days
- Command interface: 2-3 days
- Testing and optimization: 3-4 days

## Status
- [x] Spec completed
- [ ] Not Started
- [ ] In Progress
- [ ] Under Review
- [ ] Completed

## Implementation Priority
**High** - This feature significantly enhances the value proposition for launch and addresses real AI collaboration pain points identified during development.

## Success Criteria
- Demo: New AI assistant gets full project understanding in one command
- Performance: Documentation synthesis completes within 5 seconds
- Coverage: 95% of project architectural decisions captured in briefings
- Adoption: Feature becomes standard part of AI onboarding workflow

## Related Documents
- Spec: [Enhanced AI Onboarding System](../../specs/enhanced-ai-onboarding.md)
- Vision: [Weather Context Preservation](../../vision/weather-context-preservation.md)
- Phase: [Weather MVP Implementation Plan](../weather-mvp-implementation-plan.md)

## Notes
This feature directly addresses the gap identified during AI assistant onboarding and will be a key differentiator for the public launch. Implementation can proceed in parallel with core Weather MVP features.
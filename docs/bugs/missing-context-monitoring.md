# Context Monitoring Integration Gap

## Discovery Date
May 23, 2025

## Issue Description
The Weather System lacks integration with the existing context monitoring system, resulting in no proactive warnings about low context.

## Current State
- Manual `--prepare-cold-handoff` command exists in Weather System
- Recovery mechanism `--include-usage-context` works perfectly
- Backup/resilience infrastructure is in place
- **Context monitor EXISTS** at `.claude/commands/context-monitor`
- **Handoff script EXISTS** at `.claude/commands/onboard-next-assistant`
- **MISSING**: Integration between Weather System and monitoring scripts

## Impact
Users must:
1. Know about the separate monitoring scripts
2. Manually run them in addition to Weather commands
3. Remember to start the monitor at session beginning

Without integration, the Weather System's context preservation features aren't automatically invoked when needed.

## Case Study Context
This was discovered during the "Assistant Who Knew Its Own Story" incident where:
1. New assistant was reading about its own successful onboarding via Weather System
2. At 1% context remaining, no warning was given
3. The monitoring scripts existed but weren't running
4. Manual intervention required to prepare handoff

## Root Cause
- Context monitoring exists at Farm level (parent directory)
- Weather System operates at Garden level
- No automatic integration between the two systems
- Monitor uses heuristics (time-based) rather than actual context metrics

## Existing Integration Point
The `sprout init --with-claude` command already exists as a placeholder for this exact integration! Currently shows "coming soon" message but promises:
- Automatic context monitoring  
- Smart handoff detection
- Seamless session continuity

## Proposed Implementation
Complete the `sprout init --with-claude` feature by:

### Phase 1: Link Existing Scripts
- Copy/symlink `.claude/commands/` scripts during init
- Add `sprout weather monitor` wrapper command
- Start monitor automatically with Weather System

### Phase 2: Enhanced Integration  
- Store monitor state in `.garden/claude/` directory
- Add monitor status to `sprout weather` output
- Integrate with existing resilience features

### Phase 3: Native Implementation
- Port bash logic to Go for better integration
- Use actual context metrics if available
- Auto-trigger Weather System handoff preparation

## Technical Notes
- Monitor script location: `/Users/nutmeg/sprouted/.claude/commands/context-monitor`
- Uses time-based heuristics (1% per 5 minutes + activity bonus)
- Checks for logical handoff points (completed todos, major commits, branch switches)
- Auto-triggers handoff preparation at 90% threshold
- **Scripts need updates**: Not aware of `--prepare-cold-handoff` or `--include-usage-context`
- **Path issues**: Scripts assume fixed ~/sprouted path, need to be directory-agnostic

## Priority
High - This integration gap undermines the Weather System's core value proposition of "automatic" context preservation.
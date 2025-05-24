# Discovery: Weather System Dogfooding Gap

## The Irony 🐕
The Weather System, designed to preserve context for AI assistants, failed to preserve context about AI assistant tools!

## What Happened
- Claude integration tools exist at Farm level (`.claude/commands/`)
- Weather System operates at Garden level
- New AI assistants onboarded via Weather don't learn about Claude-specific tools
- User had to manually inform assistant about context monitoring

## Why This Matters
This is a textbook dogfooding failure:
1. We built a system to solve context loss
2. We experienced context loss about our context loss tools
3. The very tools that could have prevented the issue weren't known

## Root Causes

### 1. Scope Boundaries
- Weather System scans Garden-level docs
- Claude tools live at Farm level
- No cross-boundary context preservation

### 2. Tool Discovery Gap
```
Farm Level (.claude/commands/)
├── context-monitor         # Unknown to Weather
└── onboard-next-assistant  # Unknown to Weather

Garden Level (weather system)
├── --prepare-cold-handoff  # Known
└── --include-usage-context # Known
```

### 3. Documentation Assumption
Weather assumes all relevant tools are documented in Garden docs, but operational tools may live elsewhere.

## Solution: Enhanced Onboarding Context

### Quick Fix
Add to Weather's onboarding scan:
- Check parent directories for `.claude/` 
- Include tool discovery in context
- Document external integrations

### Better Fix
Complete `sprout init --with-claude` to formally integrate these tools into Weather's awareness.

### Best Fix
Weather should periodically ask: "What tools are you using that I should know about?"

## Lessons Learned
1. **Eat your own dog food** - Use Weather to document Weather workflows
2. **Cross-boundary awareness** - Don't assume all context lives in one place
3. **Tool discovery** - Actively scan for development tools, not just docs

## The Meta-Lesson
If your context preservation system doesn't know about its own helper tools, you're not truly dogfooding. The assistant who reads this should know about:
- `.claude/commands/context-monitor`
- `.claude/commands/onboard-next-assistant`
- The relationship between Weather and Claude tools
- Why this gap existed

## Action Items
- [ ] Add Farm-level tool scanning to Weather
- [ ] Document all Weather-adjacent tools
- [ ] Include "tool discovery" in onboarding
- [ ] Add 🐕 as official dogfooding indicator

Remember: A context system that loses context about itself needs more dogfooding!
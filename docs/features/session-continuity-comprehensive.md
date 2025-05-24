# Session Continuity - Comprehensive Guide

## Two Types of Session Interruptions

The Weather System handles two distinct types of session interruptions that require different approaches:

### 1. Context Limit Interruptions
**What happens:** Claude runs out of context space within the same conversation
**Characteristics:**
- Same assistant, same session
- Predictable based on message count and content length
- Can be monitored and anticipated
- Handoff preserves full conversation memory

**Current tools:**
- `sprout weather context-status` - monitors context usage
- `.claude/commands/onboard-next-assistant` - handoff command
- Weather context preservation automatically

### 2. Usage Limit Interruptions  
**What happens:** User hits daily/monthly usage limits, conversation ends completely
**Characteristics:**
- Completely new assistant starts next session
- Unpredictable timing based on user's overall Claude usage
- Cannot be monitored from within the session
- Requires cold-start onboarding of brand new assistant

**Current tools:**
- `sprout weather --onboard-ai` - comprehensive AI onboarding
- Weather System state preservation
- Manual session restoration

## Enhanced Monitoring Strategy

### For Context Limits (Existing)
```bash
# Monitor current session context
sprout weather context-status

# Prepare handoff when approaching limits
.claude/commands/onboard-next-assistant
```

### For Usage Limits (New)
```bash
# Comprehensive state capture before ending any session
sprout weather --prepare-cold-handoff

# Full onboarding for completely new assistant
sprout weather --onboard-ai --include-usage-context
```

## Best Practices

### Before Each Session Ends
1. **Always capture current state:** Run weather context preservation
2. **Document recent progress:** Update weather with latest accomplishments  
3. **Prepare next steps:** Clear todo list and roadmap status

### When Starting New Sessions
1. **Determine interruption type:** Was this context limit or usage limit?
2. **Run appropriate onboarding:** Use context handoff vs full onboarding
3. **Verify continuity:** Check that assistant understands current state

### Proactive Session Management
1. **Monitor both limits:** Context usage and overall session health
2. **Plan handoffs strategically:** Complete logical units of work
3. **Document extensively:** Assume every session might be the last

## Implementation Status

- ✅ Context limit monitoring and handoff
- 🚧 Usage limit detection and preparation  
- 🚧 Enhanced onboarding for cold starts
- 🚧 Automated state preservation
- ❌ Predictive handoff recommendations

## Usage Scenarios

### Scenario 1: Context Limit Approaching
```bash
# Assistant detects context approaching limit
sprout weather context-status
# Shows: "⚠️ Context at 85% - prepare handoff soon"

# User or assistant triggers handoff
.claude/commands/onboard-next-assistant
# Next assistant seamlessly continues
```

### Scenario 2: Usage Limit Hit Unexpectedly
```bash
# Session ends abruptly, new assistant starts
# User runs comprehensive onboarding:
sprout weather --onboard-ai

# New assistant gets full context and can continue
```

### Scenario 3: Planned Session End
```bash
# Before closing (either limit type)
sprout weather --prepare-cold-handoff
# Captures comprehensive state for any future assistant
```

## Technical Implementation Notes

- Weather System already preserves state across all interruption types
- Context monitoring exists, usage monitoring needs implementation
- Onboarding commands work for both scenarios
- Enhanced documentation needed for usage limit scenarios

---

*This document ensures no work is lost regardless of why Claude sessions end*
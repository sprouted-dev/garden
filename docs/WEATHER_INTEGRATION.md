# Weather System Integration Points

## Claude Integration Tools 🐕

The Weather System works alongside Claude-specific tools located at the Farm level:

### Context Monitoring
```bash
# Location: ~/.claude/commands/context-monitor
# Purpose: Monitors context usage and suggests handoffs

# Run continuous monitoring
~/.claude/commands/context-monitor monitor

# Quick status check  
~/.claude/commands/context-monitor status

# Check if good handoff point
~/.claude/commands/context-monitor check
```

### Assistant Handoff
```bash
# Location: ~/.claude/commands/onboard-next-assistant
# Purpose: Prepares context for new assistant after limits

# Captures Weather context and creates handoff package
~/.claude/commands/onboard-next-assistant
```

## Integration Status

### Current State
- Tools exist separately at Farm level
- Weather System has placeholders for integration
- Manual coordination required

### Planned Integration
```bash
# Coming soon - will integrate everything
sprout init --with-claude
```

## Why This Document Exists 🐕
This is a dogfooding fix! The Weather System should preserve knowledge about ALL tools in the development workflow, including those that help Weather itself. Future AI assistants onboarded via Weather will now know about these Claude integration tools.

## For New Assistants
If you're reading this after being onboarded via Weather:
1. Check if `~/.claude/commands/` exists in the Farm
2. These tools complement Weather's context preservation
3. Use them until `sprout init --with-claude` is implemented
4. Remember: Good systems dogfood themselves!
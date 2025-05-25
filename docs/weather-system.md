# Weather System Documentation

*The intelligent context preservation system for development*

## What is the Weather System?

The Weather System automatically tracks and preserves your development context, ensuring seamless continuity when switching between AI assistants or recovering from interruptions. Think of it as a "magic notebook" that remembers everything about your project, even when your AI assistant forgets.

### The Problem It Solves

When working with AI assistants, you face several "storms":
- **Context Storms**: Assistant runs out of memory mid-task
- **Usage Limit Storms**: Hit daily/monthly limits, new assistant has zero context  
- **Hallucination Storms**: AI invents incorrect dates or facts about your project
- **Documentation Storms**: Same information explained multiple times

The Weather System navigates these storms by maintaining accurate, real-time context that any assistant can instantly access.

## How It Works

### Automatic Tracking
The Weather System monitors your git activity and maintains a comprehensive context file (`.weather/context.json`) that includes:
- Project structure and active work
- Recent changes and progress
- Development methodology and patterns
- Temporal boundaries to prevent hallucination
- Session handoff information

### Weather Conditions
Your development activity is represented as weather:
- **Temperature** (0-95°F): Velocity of changes
  - 0-32°F: Frozen/no activity
  - 60-75°F: Comfortable pace  
  - 85-95°F: Rapid development
- **Conditions**: sunny (smooth), cloudy (complex), stormy (issues), foggy (blocked)
- **Pressure**: Deadline/urgency indicators

## Commands & Usage

### For Humans
```bash
# View current weather (human-readable)
sprout weather

# Update weather for specific commit
sprout weather -c abc123

# Update for current work
sprout weather --update
```

### For AI Assistants
```bash
# Get JSON context for AI consumption
sprout weather --for-ai

# Full onboarding for new assistant
sprout weather --onboard-ai

# Include usage limit context (after hitting limits)
sprout weather --onboard-ai --include-usage-context

# Prepare for usage limit interruption
sprout weather --prepare-cold-handoff
```

### Git Integration
```bash
# Install git hooks for automatic tracking
cd your-project
sprout init

# This installs hooks that update weather on:
# - commits
# - branch switches  
# - merges
# - rebases
```

## Architecture

The Weather System is part of a larger ecosystem:

```
Weather Service™ Ecosystem
├── Weather System (Local tracking) - ✓ SHIPPED
├── Weather Station (Cloud dashboard) - ✓ BUILT  
└── Seed Exchange (Methodology sharing) - 🔄 DESIGNED
```

### Core Components

1. **Context Engine**: Monitors git and file changes
2. **Weather Calculator**: Translates activity into conditions
3. **Handoff System**: Preserves context across sessions
4. **Temporal Validator**: Prevents date hallucination

## Best Practices

### Solo Development
- Run `sprout init` in every project
- Let automatic tracking handle context
- Use `--prepare-cold-handoff` before ending sessions

### AI-Assisted Development  
- Start sessions with `sprout weather --onboard-ai`
- Include `--include-usage-context` after limit interruptions
- Trust the temporal boundaries to prevent hallucination

### Team Development
- Weather Station provides shared visibility (coming soon)
- Each developer maintains local Weather System
- Automatic sync through git

## Integration with Claude

Current integration happens through command-line tools. Future `sprout init --with-claude` will provide deeper integration.

For now, use these Farm-level tools:
- `.claude/commands/context-monitor` - Monitor context size
- `.claude/commands/onboard-next-assistant` - Handoff helper

## Understanding Weather Reports

```json
{
  "temperature": 72,
  "condition": "sunny", 
  "pressure": 30,
  "context": {
    "current_work": "Consolidating documentation",
    "recent_changes": ["Removed redundant files", "Created unified guides"],
    "active_branch": "main"
  }
}
```

This indicates:
- Comfortable development pace (72°F)
- Smooth progress (sunny)
- Normal urgency (30 pressure)
- Clear current focus

## Troubleshooting

### Weather Not Updating
1. Check git hooks: `ls .git/hooks/`
2. Manually update: `sprout weather --update`
3. Verify permissions on `.weather/` directory

### Context Too Large
- Archive old weather data: `sprout weather --archive`
- Focus on recent context: `--days 7`

### Temporal Issues
If AI mentions impossible dates (before project start), the Weather System will warn. Always trust Weather System timestamps over AI assumptions.

## Philosophy

The Weather System embodies "Weather Not Roadmaps" thinking:
- Adapt to actual conditions, not fictional plans
- Preserve what happened, not what should happen  
- Enable flow state, don't interrupt it
- Make context preservation automatic, not manual

By tracking the "weather" of development, we work with natural patterns instead of against them.
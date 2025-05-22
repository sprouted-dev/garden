# Weather System Usage Guide

The Weather Context Preservation System is the heartbeat of your Garden workspace, automatically tracking your development progress and providing instant context restoration.

## Quick Start

### 1. Install Git Hooks (One-time setup)
```bash
sprout weather --install-hooks
```
This enables automatic weather updates on every commit and branch change.

### 2. Check Your Weather
```bash
sprout weather
```
Shows your current development context with:
- 🎯 **Current Focus**: What you're working on (auto-detected)
- 📈 **Recent Progress**: Summary of recent work
- 🌡️ **Temperature**: Activity level (higher = more active)
- 🌿 **Git Status**: Current branch and changes
- ⚡ **Next Steps**: AI-suggested next actions

## Core Commands

### `sprout weather`
Displays beautiful, human-readable weather context:
```
🌦️  Current Development Weather

🎯 Current Focus: authentication system (85% confidence)
📈 Recent Progress: Working on user login (last 3 commits)
🌡️  75°F | ☀️ Sunny | 🟢 Low Pressure
🌿 Branch: feature/auth (uncommitted changes)

⚡ Next Steps:
   1. Add JWT token validation
   2. Implement user session management
   3. Add password reset functionality

Last updated: 2025-05-21 15:30:45
```

### `sprout weather --for-ai`
Outputs structured JSON perfect for AI assistants like Claude:
```json
{
  "project_status": {
    "current_focus": "authentication system",
    "confidence": 0.85,
    "recent_progress": "Working on user login",
    "momentum": 75,
    "next_steps": ["Add JWT token validation", "..."]
  },
  "development_context": {
    "current_branch": "feature/auth",
    "uncommitted_changes": true,
    "last_commit_message": "Add login form validation",
    "files_changed": ["auth/login.go", "auth/middleware.go"]
  },
  "weather_conditions": {
    "temperature": 75,
    "condition": "sunny",
    "pressure": 45
  }
}
```

### `sprout weather recent`
Shows detailed recent activity:
```
📈 Recent Development Progress

Summary: Working on authentication system (last 2 hours)
Momentum: 75/100

Recent Commits:
  • Add login form validation (Fixed user input validation)
  • Implement JWT middleware (Added JWT authentication middleware)
  • Create user authentication routes (Added user authentication routes)
```

### `sprout weather --raw`
Shows the complete weather context JSON (useful for debugging).

## Weather Intelligence

### Automatic Focus Detection
Weather automatically detects what you're working on from:
- **File patterns**: `auth/` files → "authentication system"
- **Commit messages**: "feat: add API endpoint" → "API development"  
- **Branch names**: `feature/ui-redesign` → "UI redesign"

### Temperature Scale
- **🥶 0-20°F**: Cold (inactive, no recent commits)
- **😐 21-60°F**: Cool to Mild (steady progress)
- **🔥 61-95°F**: Warm to Hot (active development)
- **🌡️ 95°F+**: Very Hot (intense activity)

### Weather Conditions
- **☀️ Sunny**: Smooth progress, no blockers
- **⛅ Partly Cloudy**: Some minor challenges
- **☁️ Cloudy**: Multiple challenges or slower progress
- **⛈️ Stormy**: Major blockers or critical issues
- **🌫️ Foggy**: Unclear direction or exploration phase

## AI Integration Workflows

### Starting a New AI Session
```bash
# Get comprehensive context for AI
sprout weather --for-ai

# Copy the JSON output and share with Claude:
# "Here's my current development context: [paste JSON]"
```

### During Development
```bash
# Quick context check
sprout weather

# See what's been accomplished
sprout weather recent
```

### After Long Breaks
```bash
# Instantly restore where you left off
sprout weather

# The weather shows:
# - What you were working on
# - Recent progress made
# - Suggested next steps
# - Current git status
```

## Troubleshooting

### Weather Not Updating Automatically
```bash
# Reinstall git hooks
sprout weather --install-hooks

# Manually update from latest commit
sprout weather --update-from-commit HEAD
```

### Inaccurate Focus Detection
The system learns from your patterns. After a few commits in a new area, focus detection accuracy improves. You can help by:
- Using descriptive commit messages
- Following conventional commit format (`feat:`, `fix:`, etc.)
- Organizing files in logical directories

### Performance Issues
Weather operations are optimized for <200ms response times. If you experience slowness:
- Check if you're in a very large git repository
- Ensure `.garden/weather-context.json` isn't corrupted
- Try `sprout weather --raw` to see if data loads correctly

## Advanced Usage

### Manual Weather Updates
```bash
# Update from specific commit
sprout weather --update-from-commit abc1234

# Update from branch change
sprout weather --update-from-branch-change old-branch new-branch 1
```

### Understanding Weather Data
Weather context is stored in `.garden/weather-context.json` and includes:
- Current focus area with confidence score
- Recent commit analysis and smart summaries
- Activity momentum and temperature calculation
- Predicted next steps based on patterns
- Git repository state and branch information

## Best Practices

### For Solo Development
1. **Check weather at session start**: `sprout weather`
2. **Let git hooks handle updates**: Just commit normally
3. **Use AI integration**: Share context with AI assistants
4. **Review recent progress**: `sprout weather recent` after breaks

### For Team Collaboration
1. **Consistent commit messages**: Helps weather intelligence
2. **Logical branch naming**: `feature/`, `fix/`, `refactor/` prefixes
3. **Regular commits**: Better activity tracking and temperature
4. **Share weather context**: Include in handoff discussions

### For AI-Assisted Development
1. **Start sessions with context**: `sprout weather --for-ai`
2. **Update AI mid-session**: Share progress with `sprout weather recent`
3. **Leverage next steps**: AI can help implement suggested actions
4. **Use weather for planning**: Temperature and momentum guide work scheduling

## Integration with Garden Workflow

Weather integrates seamlessly with the Garden development process:

1. **Vision Phase**: Weather tracks high-level project direction
2. **Spec Phase**: Focus shifts to detailed requirement areas  
3. **Task Phase**: Granular task-level progress tracking
4. **Implementation**: Real-time activity monitoring and context preservation

The weather system ensures you never lose context, whether working solo or with AI assistants, making development flow seamless and productive.
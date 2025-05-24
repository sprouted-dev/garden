# Weather System - Context-Aware Development

Automatically track your development progress and eliminate context loss. Never lose your place again.

🌐 **[Try Weather System at sprouted.dev](https://sprouted.dev)** - Install in 30 seconds  
📖 **[Complete Documentation](https://sprouted.dev)** - Getting started guides and examples

[![Usage Limit Recovery Certified™](https://img.shields.io/badge/Usage%20Limit%20Recovery-Certified%20%E2%9C%93-brightgreen?style=for-the-badge)](docs/case-studies/usage-limit-recovery-live.md) [![Battle Tested](https://img.shields.io/badge/Battle%20Tested-In%20Production-blue?style=for-the-badge)](docs/case-studies/usage-limit-recovery-live.md)

## What is the Weather System?

The Weather System automatically preserves your development context - what you're working on, recent progress, and where you're headed next. Like meteorology tracks atmospheric conditions, Weather tracks your development "atmosphere" and makes it instantly shareable with AI assistants.

## Quick Start

Install and try Weather System in any Git repository:

```bash
# Install Weather System (macOS, Linux, Windows)
curl -fsSL https://sprouted.dev/install.sh | sh

# See your current development context
sprout weather

# Get AI-friendly context for assistants like Claude
sprout weather --for-ai
```

**🔥 Works in any Git repository** - no special structure required!

**See complete documentation and examples at [sprouted.dev](https://sprouted.dev)**

## Key Features

✨ **Automatic Context Tracking** - Monitors git activity and development patterns  
🤖 **AI Assistant Onboarding** - Complete project understanding in seconds  
🌡️ **Development Weather** - Visual progress tracking with temperature and conditions  
⚡ **Flow State Preservation** - Never lose your place across sessions  
🔄 **Team Handoffs** - Share rich context instantly  
🔒 **Privacy First** - All data stored locally by default  

## How It Works

Weather System understands your development context by analyzing:
- Git commit patterns and branch activity
- File modification patterns and focus areas  
- Development momentum and project health
- Architectural decisions and progress indicators

This creates a rich, shareable context that AI assistants can understand immediately - no lengthy explanations needed.

### Claude Context Management (New!)

Weather System now includes intelligent context management for Claude sessions:

```bash
# Check Claude's context usage
sprout weather context-status

# Monitor context usage in background
.claude/commands/context-monitor monitor

# Prepare handoff before hitting context limits
.claude/commands/onboard-next-assistant
```

Features:
- **Automatic Detection** - Knows when Claude is approaching context limits
- **Intelligent Handoffs** - Preserves critical context between sessions
- **Background Monitoring** - Alerts you before hitting limits
- **Universal Compatibility** - Works in any workspace type

[Learn more about Claude Context Management →](docs/features/claude-context-management.md)

## For Contributors

This repository contains the Weather System source code and follows a spec-driven development process.

### Repository Structure
- `libs/weather/` - Core Weather System library (Go)
- `apps/sprout-cli/` - Command-line interface 
- `docs/` - Specifications and development documentation
- `templates/` - Document templates for spec-driven development

### The Ultimate Validation Story 🏆

On January 23, 2025, we hit Claude's usage limits while implementing usage limit recovery. The Weather System we were building saved its own development session. [Read the full story →](docs/case-studies/usage-limit-recovery-live.md)

**"The only time in history someone was excited about hitting usage limits!"**

### Contributing
1. Check out our [Contributing Guide](CONTRIBUTING.md)
2. Browse [open issues](https://github.com/sprouted-dev/garden/issues) 
3. Join discussions in [GitHub Discussions](https://github.com/sprouted-dev/garden/discussions)

**🌱 Built by developers, for developers - with the community at heart.**
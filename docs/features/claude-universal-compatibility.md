# Claude Context Management - Universal Compatibility

The Claude integration is designed to work with ALL types of workspaces, not just Sprouted projects.

## Workspace Types Supported

### 1. Sprouted Workspaces

#### Farms (Multi-repo)
- Full Weather System integration across gardens
- Cross-repository context tracking
- Event-based coordination
- Lower context thresholds (more aggressive handoffs)

#### Gardens (Git repos)
- Complete Weather System features
- Automatic git activity tracking
- Smart inference from commits
- Standard context thresholds

#### Seedlings (Small projects)
- Basic Weather features
- Simple context tracking
- Higher thresholds (less frequent handoffs)

### 2. Non-Sprouted Workspaces

#### Standard Git Repositories
```bash
# Works even without Weather System
cd any-git-repo
curl -fsSL https://sprouted.dev/claude-setup.sh | bash

# Features available:
# ✅ Context monitoring (time-based)
# ✅ Git activity detection
# ✅ Handoff preparation
# ❌ Weather System intelligence
```

#### Non-Git Projects
```bash
# Works in any directory
cd my-project
curl -fsSL https://sprouted.dev/claude-setup.sh | bash

# Features available:
# ✅ Basic context monitoring
# ✅ Time-based usage estimates
# ✅ Manual handoff triggers
# ❌ Git-based intelligence
```

#### Monorepos (Non-Sprouted)
```bash
# Detects multiple git repos
cd my-monorepo
curl -fsSL https://sprouted.dev/claude-setup.sh | bash

# Treated as "farm-like"
# ✅ Multi-repo awareness
# ✅ Aggregate context tracking
```

## Graceful Degradation

The system adapts based on available features:

### With Full Weather System
```javascript
// All features available
{
  "context_tracking": "weather-enhanced",
  "handoff_detection": "intelligent",
  "git_integration": "full",
  "inference": "smart"
}
```

### With Git Only
```javascript
// Git-based features
{
  "context_tracking": "git-based",
  "handoff_detection": "commit-patterns",
  "git_integration": "basic",
  "inference": "limited"
}
```

### Standalone Mode
```javascript
// Time-based only
{
  "context_tracking": "time-based",
  "handoff_detection": "manual",
  "git_integration": "none",
  "inference": "none"
}
```

## Installation Adaptations

### Detection Logic
```bash
# The setup script detects:
if [ -x "$(command -v sprout)" ]; then
    # Full Sprouted installation
elif [ -d ".git" ]; then
    # Git-based installation
else
    # Standalone installation
fi
```

### Feature Availability

| Feature | Sprouted | Git Only | Standalone |
|---------|----------|----------|------------|
| Context Monitoring | ✅ Weather-enhanced | ✅ Git-based | ✅ Time-based |
| Usage Estimation | ✅ Multi-factor | ✅ Commits+time | ✅ Time only |
| Handoff Detection | ✅ Intelligent | ✅ Commit patterns | ❌ Manual only |
| Context Capture | ✅ Full Weather | ✅ Git status | ✅ Basic info |
| Progress Tracking | ✅ Automatic | ✅ From commits | ❌ Manual |

## Usage Examples

### In a Python Project (No Sprouted)
```bash
cd my-django-app
curl -fsSL https://sprouted.dev/claude-setup.sh | bash

# Works with:
# - Time-based monitoring
# - Git commit detection
# - Basic handoff points
```

### In a Node.js Project
```bash
cd my-react-app  
npm install @sprouted/claude-tools  # Future NPM package
# OR
curl -fsSL https://sprouted.dev/claude-setup.sh | bash
```

### In Any Directory
```bash
cd ~/documents/research
curl -fsSL https://sprouted.dev/claude-setup.sh | bash

# Still provides:
# - Context usage estimates
# - Manual handoff preparation  
# - Session preservation
```

## Compatibility Matrix

| Environment | Context Monitor | Handoff Prep | Weather Integration | Auto-detect |
|-------------|----------------|--------------|-------------------|-------------|
| Sprouted Farm | ✅ | ✅ | ✅ | ✅ |
| Sprouted Garden | ✅ | ✅ | ✅ | ✅ |
| Git Repository | ✅ | ✅ | ❌ | ✅ |
| SVN/Mercurial | ✅ | ✅ | ❌ | ⚠️ Limited |
| Plain Directory | ✅ | ✅ | ❌ | ❌ |
| Docker Container | ✅ | ✅ | ❌ | ✅ |
| Cloud IDE | ✅ | ✅ | ❌ | ✅ |
| Jupyter Notebook | ✅ | ✅ | ❌ | ⚠️ Special |

## Special Considerations

### For Large Codebases
```bash
# Adjust thresholds for longer sessions
CONTEXT_WARNING_THRESHOLD=70
CONTEXT_CRITICAL_THRESHOLD=80
```

### For Research/Documentation
```bash
# Higher thresholds for text-heavy work
CONTEXT_WARNING_THRESHOLD=85
CONTEXT_CRITICAL_THRESHOLD=95
```

### For Debugging Sessions
```bash
# Lower thresholds for complex debugging
CONTEXT_WARNING_THRESHOLD=60
CONTEXT_CRITICAL_THRESHOLD=75
```

## Future Universal Features

### Planned Integrations
- VSCode extension
- JetBrains plugin  
- Vim/Neovim plugin
- Emacs package

### Language-Specific Adapters
- Python: Parse virtualenv activity
- JavaScript: Monitor node_modules changes
- Go: Track go.mod updates
- Rust: Monitor cargo activity

### IDE Integration
```javascript
// Future API
const claude = require('@sprouted/claude-context');

claude.monitor({
  threshold: 80,
  onWarning: () => console.log('Context getting full'),
  onCritical: () => claude.prepareHandoff()
});
```

## Minimum Requirements

- Bash shell (or compatible)
- Basic file system access
- 1MB disk space
- No network required (fully offline)

## Summary

The Claude Context Management system is designed as a **universal tool** that enhances any development environment, with additional features when Sprouted tools are available. It gracefully degrades to provide value even in the most basic environments while scaling up to leverage full Weather System intelligence when available.
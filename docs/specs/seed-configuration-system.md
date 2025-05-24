# Seed Configuration System Specification

*Created: May 24, 2025*

## Overview

Seeds define how teams work. Weather System reads Seeds to understand team patterns and adapt its behavior accordingly.

## Seed Structure

```
.seed/
├── config.json           # Core configuration
├── prompts/             # Custom AI prompts
│   ├── storm-detect.md  # Detect creative chaos
│   └── onboarding.md    # Team-specific onboarding
├── patterns/            # Methodology patterns
│   └── branches.json    # Branch naming rules
└── philosophy/          # Optional philosophy docs
    └── README.md        # How this team works
```

## Core Configuration Schema

```json
{
  "seed": {
    "name": "string",              // Unique identifier
    "version": "1.0.0",           // Seed version
    "methodology": "string",       // Named methodology
    "description": "string",       // What makes this unique
    
    "patterns": {
      "branches": {
        "experimental": "regex",   // Experimental work
        "feature": "regex",        // Feature branches
        "release": "regex",        // Release branches
        "hotfix": "regex"         // Emergency fixes
      },
      "commits": {
        "format": "string",       // Commit message format
        "types": ["feat", "fix"]  // Allowed types
      }
    },
    
    "metrics": {
      "velocity": {
        "unit": "string",         // points, features, etc
        "period": "string"        // sprint, hour, day
      },
      "custom": {}               // Team-specific metrics
    },
    
    "automation": {
      "triggers": [              // Event-based automation
        {
          "event": "branch.create",
          "pattern": "tornado/*",
          "action": "notify.storm"
        }
      ]
    },
    
    "ai": {
      "context": {
        "priority": ["docs", "pattern"], // Context priority
        "exclude": ["test", "vendor"]    // Ignore patterns
      },
      "prompts": {
        "onboarding": "prompts/onboarding.md",
        "detection": "prompts/storm-detect.md"
      }
    }
  }
}
```

## Example Seeds

### Sprouted Seed (Storm-Driven Development)
```json
{
  "seed": {
    "name": "sprouted",
    "methodology": "storm-driven-development",
    "patterns": {
      "branches": {
        "experimental": "tornado/*",
        "storm": "storm/*"
      }
    },
    "metrics": {
      "velocity": {
        "unit": "features",
        "period": "hour"
      }
    }
  }
}
```

### Enterprise Seed (Traditional Agile)
```json
{
  "seed": {
    "name": "enterprise-agile",
    "methodology": "scaled-agile",
    "patterns": {
      "branches": {
        "feature": "feature/JIRA-*",
        "release": "release/*"
      }
    },
    "metrics": {
      "velocity": {
        "unit": "story-points",
        "period": "sprint"
      }
    }
  }
}
```

### Startup Seed (Move Fast)
```json
{
  "seed": {
    "name": "startup-chaos",
    "methodology": "controlled-chaos",
    "patterns": {
      "branches": {
        "experimental": "try/*",
        "ship": "ship/*"
      }
    },
    "metrics": {
      "velocity": {
        "unit": "shipped",
        "period": "day"
      }
    }
  }
}
```

## Weather System Integration

Weather reads seed configuration and:

1. **Adapts Monitoring**
   - Tracks metrics in specified units
   - Recognizes branch patterns
   - Applies automation rules

2. **Customizes AI Context**
   - Uses team-specific prompts
   - Prioritizes relevant docs
   - Excludes noise

3. **Enables Methodology**
   - Activates relevant features
   - Applies team patterns
   - Preserves philosophy

## Benefits

- **Teams work their way** - No forced methodology
- **Weather stays generic** - Pure infrastructure
- **Seeds are shareable** - Learn from others
- **Evolution enabled** - Seeds can version and grow

## Implementation Notes

- Seeds are JSON for easy parsing
- Optional philosophy docs for humans
- Prompts customize AI behavior
- Patterns drive automation
- Metrics inform dashboards

This specification enables Weather System to support any development methodology while remaining unopinionated infrastructure.
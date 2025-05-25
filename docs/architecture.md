# Sprouted Architecture: Seeds, Gardens, and Farms

*A natural approach to project organization and growth*

## Core Concepts

The Sprouted ecosystem uses agricultural metaphors that match how projects naturally evolve:

### 🌱 Seeds (Individual Projects)
A Seed is any project with a `docs/` directory. That's it. No complex requirements, no forced structure - just your project and its documentation.

**Examples:**
- A startup's MVP (minimal seed)
- An open source library (community seed)  
- An enterprise application (structured seed)

**Key principle**: Seeds contain the DNA of how your team works, captured in documentation patterns.

### 🌿 Gardens (Shared Infrastructure)
When Seeds need to share resources (libraries, tools, configs), they grow together in a Garden. This is typically a monorepo or workspace.

**Examples:**
- Multiple microservices sharing utilities
- Website + API + mobile app in one repo
- Open source project with multiple packages

**What Gardens provide:**
- Shared Weather System tracking
- Common development tools
- Unified dependency management
- Cross-project visibility

### 🚜 Farms (Organization Level)
Farms coordinate multiple Gardens, providing organization-wide visibility and pattern recognition. This is where the big picture emerges.

**Examples:**
- A company with multiple product repositories
- An open source organization with many projects
- A consultancy managing client Gardens

**What Farms enable:**
- Cross-repository Weather tracking
- Organization-wide pattern detection
- Unified methodology evolution
- Privacy boundaries (public/private Gardens)

## Natural Evolution Patterns

Projects grow naturally through these stages:

```
Seed → Garden → Farm
 │      │        │
 │      │        └── Multiple Gardens need coordination
 │      └────────── Multiple Seeds share infrastructure  
 └───────────────── Single project with docs/
```

### Real-World Example: Startup Growth

1. **Seed Stage**: MVP in a single repo with basic docs
2. **Garden Stage**: Add mobile app, share component library
3. **Farm Stage**: Launch second product, need unified insights

### Open Source Example

1. **Seed**: Core library with README
2. **Garden**: Add CLI tool, documentation site, examples
3. **Farm**: Community plugins in separate repos

## Technical Implementation

### Seed Structure
```
my-project/
├── docs/
│   ├── README.md          # How we work
│   ├── decisions/         # Architectural choices
│   └── onboarding/        # Getting started
├── src/                   # Your actual code
└── .weather/              # Context tracking
```

### Garden Architecture
```
my-garden/
├── apps/
│   ├── web/              # Seed 1
│   └── mobile/           # Seed 2
├── libs/
│   └── shared/           # Shared resources
├── docs/                 # Garden-level patterns
└── .weather/             # Unified tracking
```

### Farm Coordination
```
~/mycompany/
├── .farm/
│   ├── events/           # Cross-garden activity
│   └── patterns/         # Detected methodologies
├── main-product/         # Garden 1
├── internal-tools/       # Garden 2  
└── oss-projects/         # Garden 3
```

## Working with the Architecture

### Quick Start (30 seconds)
```bash
# Create a Seed
mkdir my-project/docs
echo "# How We Work" > my-project/docs/README.md

# Initialize Weather System
cd my-project
sprout init

# That's it! You have a Seed
```

### Pattern Detection

The Weather System automatically detects:
- Documentation structure → Team methodology
- Commit patterns → Development velocity
- Branch strategies → Workflow preferences
- PR descriptions → Communication style

### Best Practices

**DO:**
- Start with a simple Seed
- Let patterns emerge naturally
- Share resources when it makes sense
- Use Farms when coordination helps

**DON'T:**
- Over-engineer from the start
- Force artificial boundaries
- Create Gardens before needed
- Add complexity without benefit

## Integration with Weather System

The Weather System operates at all levels:

- **Seed Level**: Tracks individual project context
- **Garden Level**: Aggregates shared activity
- **Farm Level**: Provides organization insights

```bash
# Seed weather
cd my-project && sprout weather

# Garden weather  
cd my-garden && sprout weather --garden

# Farm weather
cd ~ && sprout weather --farm
```

## Philosophy: Natural Growth

This architecture embraces how projects actually grow:

1. **Start small** - Just docs/ and code
2. **Share when needed** - Not before
3. **Coordinate when helpful** - Not by default
4. **Patterns emerge** - Don't force them

Unlike rigid frameworks, this approach:
- Requires no upfront planning
- Adapts to your actual needs
- Preserves existing workflows
- Scales naturally with growth

## Common Patterns

### The Minimal Seed
```
project/
└── docs/
    └── README.md  # Just explain what this is
```

### The Team Seed  
```
project/
└── docs/
    ├── README.md
    ├── decisions/     # Why we chose X
    ├── onboarding/    # How to get started
    └── workflows/     # How we work
```

### The Enterprise Seed
```
project/
└── docs/
    ├── architecture/
    ├── compliance/
    ├── operations/
    ├── governance/
    └── standards/
```

## Future: Event-Driven Architecture

Farms will use event streams for real-time coordination:

```json
{
  "event": "commit.created",
  "garden": "main-product",
  "seed": "web-app",
  "timestamp": "2025-05-24T10:30:00Z",
  "weather": {
    "temperature": 78,
    "condition": "sunny"
  }
}
```

This enables:
- Cross-garden activity streams
- Pattern detection algorithms
- Methodology evolution tracking
- Privacy-preserving insights

## Summary

Seeds, Gardens, and Farms aren't imposed structure - they're recognition of natural patterns. Start with a Seed (any project with docs/), grow into a Garden when sharing makes sense, coordinate through Farms when scale demands it.

The architecture grows with you, not ahead of you.
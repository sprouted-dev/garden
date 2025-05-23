# Farm/Garden/Seed Patterns

The Weather System uses a three-level hierarchy to organize documentation and project structure. This pattern provides flexibility while maintaining consistency across different project scales.

## Core Concepts

### 🌱 Seed - Individual Project
The atomic unit of documentation. Every repository should have a Seed.

**What it is:**
- A `docs/` directory with at least a README.md
- Documentation that lives with the code
- The source of truth for project-specific information

**Examples:**
- A single microservice
- A personal project
- A standalone library
- An individual app in a monorepo

### 🌾 Garden - Shared Infrastructure
The coordinating layer for multiple related projects.

**What it is:**
- Cross-cutting concerns documentation
- Shared tooling and libraries
- Standards and patterns used across Seeds
- The "how we work together" layer

**Examples:**
- This repository (`sprouted-dev/garden`)
- A company's shared component library
- Platform team documentation
- Design system repository

### 🚜 Farm - Organizational Level
The highest level of coordination, typically organization-wide.

**What it is:**
- Enterprise governance and compliance
- Organization-wide standards
- Strategic technical decisions
- Multi-Garden coordination

**Examples:**
- Company-wide engineering standards
- Compliance documentation
- Cross-team architectural decisions
- Strategic technology roadmaps

## Pattern Recognition

Weather System automatically detects these patterns:

### Seed Detection
```
my-project/
├── docs/
│   └── README.md      ← Minimal Seed
├── src/
└── package.json
```

### Garden Detection
```
platform/
├── docs/
│   ├── README.md
│   ├── standards/     ← Cross-project standards
│   ├── patterns/      ← Shared patterns
│   └── tooling/       ← Common tooling docs
├── libs/              ← Shared libraries
└── tools/             ← Shared tools
```

### Farm Detection
```
enterprise/
├── governance/
│   ├── compliance/
│   ├── security/
│   └── policies/
├── architecture/
│   ├── decisions/     ← Enterprise ADRs
│   └── principles/
└── standards/
    ├── engineering/
    └── operations/
```

## When to Use Each Level

### Use a Seed When:
- Starting any new project
- Documenting a specific codebase
- You need local, project-specific docs
- Working on individual features or services

### Use a Garden When:
- Multiple projects share infrastructure
- You need cross-project standards
- Building shared tooling or libraries
- Coordinating between related projects

### Use a Farm When:
- Managing enterprise-wide concerns
- Dealing with compliance requirements
- Setting organization standards
- Coordinating multiple teams/divisions

## Practical Examples

### Example 1: Startup Evolution

**Month 1 - Just Seeds**
```
cool-app/
└── docs/README.md     ← Single Seed

api/
└── docs/README.md     ← Another Seed
```

**Month 6 - Garden Emerges**
```
platform/              ← Garden
├── docs/
│   ├── standards/
│   └── guides/
├── cool-app/         ← Seed
└── api/              ← Seed
```

**Year 2 - Farm Structure**
```
company/               ← Farm
├── governance/
├── platform/          ← Garden
│   ├── cool-app/     ← Seed
│   └── api/          ← Seed
└── data-team/         ← Another Garden
    └── analytics/     ← Seed
```

### Example 2: Open Source Project

```
my-oss-project/        ← Acts as both Seed and Garden
├── docs/
│   ├── README.md     ← Project docs (Seed)
│   ├── contributing/ ← Community standards (Garden)
│   └── architecture/ ← Technical decisions (Garden)
├── packages/
│   ├── core/         ← Seed
│   └── cli/          ← Seed
└── examples/         ← Seeds
```

## Anti-Patterns to Avoid

### ❌ Over-Engineering
Don't create Garden/Farm structure before you need it.

### ❌ Rigid Hierarchy
These are patterns, not rules. Adapt to your needs.

### ❌ Documentation Theater
Don't create structure just to "look professional."

### ❌ Ignoring Evolution
Let structure emerge naturally as projects grow.

## Integration with Weather System

The Weather System understands these patterns:

```bash
# In a Seed
sprout weather
# Shows: Project-specific context

# In a Garden
sprout weather
# Shows: Cross-project context + infrastructure health

# In a Farm
sprout weather
# Shows: Organizational weather patterns
```

## Migration Paths

### Seed → Garden
When you find yourself:
- Copying docs between projects
- Sharing code/tools
- Needing consistent standards

### Garden → Farm
When you have:
- Multiple teams/divisions
- Compliance requirements
- Enterprise governance needs

## Best Practices

1. **Start with Seeds** - Every project gets a Seed
2. **Grow Gardens naturally** - When sharing emerges
3. **Farms are rare** - Most teams only need Seeds/Gardens
4. **Document reality** - Reflect how you actually work
5. **Evolve gradually** - Don't restructure everything at once

## Quick Reference

| Level | Scope | Examples | Key Files |
|-------|-------|----------|-----------|
| 🌱 Seed | Single project | App, service, library | `docs/README.md` |
| 🌾 Garden | Shared platform | Design system, tooling | `docs/standards/` |
| 🚜 Farm | Organization | Enterprise standards | `governance/` |

Remember: These patterns exist to help, not constrain. Use what makes sense for your context.
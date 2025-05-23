# Team Seed Example

A documentation structure optimized for team collaboration and decision tracking.

## How We Work

We follow a lightweight agile process with emphasis on:
- Clear specifications before implementation
- Documented decisions for future reference
- Regular retrospectives to improve

## Directory Structure

```
docs/
├── README.md           # You are here
├── specs/              # Feature specifications
│   ├── active/        # Currently being implemented
│   └── completed/     # Finished features
├── decisions/         # Architectural Decision Records (ADRs)
│   └── YYYY-MM-DD-decision-title.md
└── retrospectives/    # Team learnings
    └── YYYY-MM-sprint-N.md
```

## Key Concepts

### Specifications
Before building features, we write specs that cover:
- Problem statement
- Proposed solution
- Success criteria
- Technical approach

### Decision Records
Major technical decisions are documented with:
- Context and problem
- Options considered
- Decision made
- Consequences

### Retrospectives
Every sprint we capture:
- What went well
- What could improve
- Action items

## Weather Integration Points

- **Progress Tracking**: Active specs show current focus
- **Decision History**: ADRs provide context for choices
- **Team Learning**: Retrospectives show evolution
- **Momentum**: Regular updates to these docs indicate healthy velocity

## Getting Started

1. Check `specs/active/` for current work
2. Read recent decisions in `decisions/`
3. Review last retrospective for team context
4. Pick up a spec or create a new one
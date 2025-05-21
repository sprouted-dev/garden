# Garden Monorepo

A spec-driven development workspace for apps, libraries, and tools.

## Structure

```
garden/
├── apps/           # Applications (React, Next.js, Expo, Go, Rust, etc.)
├── libs/           # Shared libraries and packages
├── tools/          # Development tools and utilities
├── docs/           # Documentation system
│   ├── vision/     # Long-term vision documents
│   ├── specs/      # Detailed specifications
│   ├── tasks/      # Implementation tasks
│   ├── phases/     # Development phases
│   └── workflows/  # Development workflows
└── templates/      # Document templates
```

## Development Process

1. **Brainstorm/Plan** → Create vision document
2. **Spec** → Define detailed requirements
3. **Tests** → Write acceptance criteria
4. **Tasks** → Break down implementation
5. **Implementation** → Execute with AI assistance

## Quick Start

1. Use templates from `/templates/` to create new documents
2. Follow the hierarchy: Vision → Spec → Tasks → Implementation
3. Reference `/docs/workflows/agentic-development.md` for AI collaboration guidelines

## Documentation Guidelines

- All projects should have corresponding vision and spec documents
- Maintain traceability across all document types
- Use the time-based hierarchy to track progress
- Keep documentation up-to-date with implementation
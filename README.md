# Garden Monorepo

A spec-driven development workspace with automatic context preservation powered by the Weather System.

## Weather Context Preservation

The Weather System automatically tracks your development progress and provides instant context restoration:

```bash
# One-time setup
sprout weather --install-hooks

# Check your current development context
sprout weather

# Get AI-friendly context for Claude
sprout weather --for-ai
```

**See [Weather Usage Guide](docs/WEATHER_USAGE.md) for complete documentation.**

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

1. **Setup Weather System**: `sprout weather --install-hooks` (enables automatic context tracking)
2. **Check Your Context**: `sprout weather` (see current focus and progress)
3. **Create Documents**: Use templates from `/templates/` following Vision → Spec → Tasks → Implementation
4. **AI Collaboration**: Use `sprout weather --for-ai` to share context with AI assistants
5. **Reference**: See `/docs/workflows/agentic-development.md` for detailed AI collaboration guidelines

## Documentation Guidelines

- All projects should have corresponding vision and spec documents
- Maintain traceability across all document types
- Use the time-based hierarchy to track progress
- Keep documentation up-to-date with implementation
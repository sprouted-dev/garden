# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Structure

This is a spec-driven development monorepo called "Garden" with the following architecture:

- `apps/` - Applications (React, Next.js, Expo, Go, Rust, etc.)
- `libs/` - Shared libraries and packages
- `tools/` - Development tools and utilities
- `docs/` - Documentation system with time-based hierarchy
- `templates/` - Document templates for consistent project documentation

## Development Workflow

The repository follows a spec-driven development process with a strict documentation hierarchy:

1. **Vision** (`/docs/vision/`) → Long-term goals and direction
2. **Specs** (`/docs/specs/`) → Detailed requirements with acceptance criteria
3. **Tasks** (`/docs/tasks/`) → Actionable work items broken down from specs
4. **Phases** (`/docs/phases/`) → Implementation phases
5. **Implementation** → Code execution

## Documentation Requirements

- All new projects/features must have corresponding vision and spec documents
- Use templates from `/templates/` directory for consistency
- Maintain traceability: each document should reference related documents in other hierarchy levels
- Follow the agentic development workflow in `/docs/workflows/agentic-development.md`

## Key Principles

- Break large features into small, testable units
- Create clear acceptance criteria before implementation
- Always reference existing code patterns and conventions
- Use file paths and line numbers when discussing specific code
- Write tests before implementation when possible
- Validate implementation against specs

## Working with This Repository

- Start with vision documents for new features
- Create detailed specs with functional/non-functional requirements
- Break specs into actionable tasks
- Implement following existing project conventions
- Maintain documentation up-to-date with implementation
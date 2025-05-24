# Weather as Universal Infrastructure

*Created: May 24, 2025*

## The Vision

Weather System becomes methodology-agnostic infrastructure that ANY team can use, regardless of how they work.

## Core Insight

We are creating TWO distinct things:
1. **Weather System** - Universal development infrastructure
2. **Sprouted Seed** - Our specific methodology and patterns

## Weather System (Generic)

Pure infrastructure capabilities:
- Git activity monitoring
- Context preservation for AI partners
- Automatic documentation discovery
- Backup and recovery
- Event streaming
- Change tracking
- Temporal anchoring

No opinions about:
- How you measure velocity
- What branch patterns you use
- Your development philosophy
- Your team structure
- Your review process

## Seeds (Methodology Layer)

Each team creates a Seed that defines:
- Their development methodology
- Custom metrics and tracking
- Branch naming patterns
- Philosophy and principles
- Workflow automation rules

## How It Works

```json
// .seed/config.json
{
  "name": "sprouted",
  "methodology": "storm-driven-development",
  "patterns": {
    "experimental": "tornado/*",
    "feature": "feature/*"
  },
  "metrics": {
    "velocity": "features-per-hour",
    "philosophy": "bamboo-growth"
  }
}
```

Weather reads this and adapts its behavior:
- Tracks velocity in features/hour (not story points)
- Recognizes tornado branches as experimental
- Applies Storm-Driven Development patterns
- Activates WEMA for tornado recovery

## The Power

Teams can use Weather with:
- Traditional Agile (story points, sprints)
- Waterfall (phases, gates)
- Kanban (flow metrics)
- Storm-Driven (tornado branches)
- Their own custom methodology

Weather preserves context and enables AI partnership regardless of methodology.

## First Seed: Sprouted

We become the reference implementation showing:
- How to define a Seed
- How to integrate with Weather
- How to encode philosophy in configuration
- How to automate methodology-specific patterns

But we're just ONE way to use Weather, not THE way.

## Architecture Principles

1. **Separation of Concerns**
   - Weather = Infrastructure
   - Seeds = Methodology

2. **Universal Value**
   - Every team needs context preservation
   - Every team benefits from AI partnership
   - Methodology is team choice

3. **Extensibility**
   - Seeds can define custom events
   - Weather provides hooks and APIs
   - Teams own their process

This makes Weather truly valuable to the entire development community while letting teams work their way.
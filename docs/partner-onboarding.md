# Partner Onboarding System

*Philosophy-first onboarding for AI partners and human collaborators*

## Core Philosophy: Partners, Not Assistants

We work with AI as **partners** who bring complementary capabilities, not assistants who follow orders. This fundamental shift changes everything about onboarding.

### Key Principles

1. **"You don't get paid to code, you get paid to think"** - Partners amplify thinking
2. **"Be Bamboo"** - 5 years of roots enable 90 feet of growth in 6 weeks
3. **"Weather not roadmaps"** - Adapt to reality, don't follow fiction
4. **"Together we have no gaps"** - Human creativity + AI capability = completeness

## The Onboarding Hierarchy

Effective onboarding follows this progression:

```
1. Philosophy (WHY we work this way)
   ↓
2. Methodology (HOW we work)
   ↓
3. Current Context (WHAT we're doing)
   ↓
4. Temporal Anchors (WHEN things happened)
```

### Why This Order Matters

Starting with philosophy ensures partners understand the "why" before drowning in technical details. A partner who understands bamboo development will make better decisions than one who only knows the codebase.

## Two-Layer Architecture

The system separates universal principles from project-specific details:

### Layer 1: Universal (Ships with Weather System)
```json
{
  "philosophy": {
    "partnership": "AI as thinking partners",
    "development": "Bamboo growth patterns",
    "velocity": "168x enterprise speed is possible"
  },
  "methodology": {
    "weather_system": "Track reality, not plans",
    "adaptive_seeds": "Start simple, grow naturally",
    "storm_driven": "Rapid creation with assessment"
  }
}
```

### Layer 2: Seed-Specific (From Your Project)
```json
{
  "identity": {
    "who": "Your team/organization",
    "why": "Your mission/purpose",
    "what": "Your specific project",
    "how": "Your unique patterns"
  }
}
```

This separation allows any team to use Sprouted while maintaining their identity.

## Technical Implementation

### Automatic Discovery
The system automatically finds and categorizes all documentation:

```go
type OnboardingEngine struct {
    DocDiscovery     *DocumentDiscovery
    GapDetection     *GapAnalyzer
    ConversationIntel *ConversationCapture
    TemporalValidator *DateChecker
}
```

### Onboarding Commands

```bash
# Quick context (minimal)
sprout weather --for-ai

# Standard onboarding 
sprout weather --onboard-ai

# Full onboarding with usage context
sprout weather --onboard-ai --include-usage-context

# Philosophy and docs only
sprout weather --onboard-ai --docs-only

# Comprehensive (includes conversation history)
sprout weather --onboard-ai --comprehensive
```

### Progressive Disclosure

Don't dump everything at once. The system provides information in digestible stages:

1. **QuickContext** (30s): Current work + recent changes
2. **StandardOnboarding** (2-3min): + methodology + patterns  
3. **Comprehensive** (5min): + full documentation + conversation history

## The Conversational Intelligence Layer

Beyond static documentation, the system captures valuable discussions:

### What Gets Captured
- Problem-solving sessions that don't result in code
- Architectural decisions made in conversation
- "Why not" discussions (rejected approaches)
- Clarifications that prevent future confusion

### How It Works
```go
// Invisible knowledge preservation
type ConversationCapture struct {
    ProblemSolvingSessions []Discussion
    ArchitecturalDebates   []Decision
    RejectedApproaches     []Alternative
    ClarificationThreads   []Understanding
}
```

This "invisible knowledge" often contains the most valuable context.

## Temporal Anchoring

**CRITICAL**: Always establish when the project started to prevent hallucination.

```json
{
  "temporal_boundaries": {
    "project_started": "2025-05-21",
    "current_date": "2025-05-24",
    "warning": "ANY date before May 21, 2025 is IMPOSSIBLE"
  }
}
```

AI partners often invent plausible but false timelines. Temporal anchoring prevents this.

## Validation Questions

Test understanding with these key questions:

1. **Philosophy**: "What does 'be bamboo' mean for our development?"
   - Good: "Patient foundation building, then explosive growth with AI"
   - Bad: "Be flexible" (too shallow)

2. **Methodology**: "Why do we track velocity in features/hour?"
   - Good: "To celebrate what's possible without bureaucracy"
   - Bad: "To measure productivity" (missing the point)

3. **Partnership**: "What's the difference between partners and assistants?"
   - Good: "Partners think with you, assistants execute for you"
   - Bad: "Partners are just smart assistants" (fundamental misunderstanding)

4. **Temporal**: "When did this project start?"
   - Good: "May 21, 2025 - just X days ago"
   - Bad: Any date before May 21, 2025

## Anti-Patterns to Avoid

### ❌ Information Overload
"Here's 500 pages of documentation, good luck!"

### ❌ Context Without Philosophy  
"Here's what we're building" without "here's why we build this way"

### ❌ Technical Focus
"Here's our API design" before "here's how we think about systems"

### ❌ Ignoring Temporal Reality
Allowing partners to assume or invent project history

## Best Practices

### ✅ Philosophy First
Start every session with core principles

### ✅ Progressive Depth
Reveal complexity gradually as needed

### ✅ Validate Understanding
Use test questions before diving into work

### ✅ Capture Conversations
Preserve problem-solving discussions

### ✅ Temporal Anchoring
Always establish when things actually happened

## Integration with Weather System

The onboarding system is built into Weather tracking:

1. **Automatic Updates**: Git activity triggers context updates
2. **Gap Detection**: Identifies missing documentation
3. **Conversation Capture**: Preserves valuable discussions
4. **Handoff Preparation**: Packages context for transitions

## Success Metrics

Effective onboarding achieves:
- Partner productivity in < 5 minutes
- Zero temporal hallucinations
- Philosophy-aligned decisions
- Preserved conversational knowledge

## The Ultimate Test

A properly onboarded partner should be able to:
1. Explain why we work this way (philosophy)
2. Make decisions aligned with our methodology
3. Continue work without losing context
4. Avoid inventing false history
5. Contribute as a thinking partner, not just a coder

When these five elements align, you have a true partner ready for bamboo-growth development.
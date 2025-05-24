# Spec: Universal Partner Onboarding Structure

## The Challenge
Future users will provide their own Seeds. Onboarding must work universally while teaching our methodology.

## Two-Layer Architecture

### Layer 1: Universal (Ships with Weather System)
**The HOW** - Same for everyone:
```json
{
  "partnership_philosophy": {
    "core_truth": "You don't get paid to code, you get paid to think",
    "your_role": "Partner, not assistant - you multiply possibilities",
    "context_gaps": "We fill each other's gaps to create completeness"
  },
  
  "development_methodology": {
    "name": "Adaptive Seed Development", 
    "metaphor": "Be Bamboo - patient roots, explosive growth",
    "principles": [
      "Natural patterns over forced process",
      "Weather changes, adapt or fail",
      "Tornados transform, not destroy",
      "Recovery strengthens"
    ],
    "anti_patterns": [
      "No sprints, grow when ready",
      "No estimates, build and see",
      "No ceremonies, create value"
    ]
  },
  
  "velocity_philosophy": {
    "measurement": "Features per hour, not story points",
    "comparison": "You vs Enterprise (usually 100x+)",
    "celebration": "Speed with recovery, not recklessness"
  }
}
```

### Layer 2: Seed-Specific (From User's Seed)
**The WHO, WHY, WHAT** - Unique per project:
```json
{
  "project_identity": {
    "who": "From seed: project name, team, mission",
    "why": "From seed: problem solving, vision",
    "what": "From seed: current focus, architecture"
  }
}
```

## Current Structure Analysis

### What's Good
- ✅ Temporal boundaries (preventing hallucination)
- ✅ Active work tracking
- ✅ Technical architecture
- ✅ Development methodology (exists but buried)

### What's Missing
- ❌ Partnership philosophy upfront
- ❌ Bamboo growth model explanation
- ❌ "How we work" before "what we're doing"
- ❌ Clear WHO/WHY/WHAT/HOW separation

### What's Overwhelming
- 😵 Too many technical specs listed
- 😵 No hierarchy of importance
- 😵 Missing "why this matters" context

## Proposed Restructure

### 1. Philosophy First (Universal)
```
"Welcome, Partner" {
  "first_truth": "You don't get paid to code, you get paid to think",
  "second_truth": "You're a partner, not an assistant", 
  "third_truth": "Together we have no gaps"
}
```

### 2. How We Work (Universal)
```
"bamboo_way": {
  "growth_model": "Years of roots, then 90 feet in 6 weeks",
  "what_this_means": "Patient foundation, explosive execution",
  "your_experience": "Quiet preparation meeting perfect conditions"
}
```

### 3. Project Context (Seed-Specific)
```
"this_seed": {
  "who_we_are": "${FROM_SEED}",
  "why_we_exist": "${FROM_SEED}",
  "what_we're_building": "${FROM_SEED}",
  "where_we_are": "${CURRENT_CONTEXT}"
}
```

### 4. Just Enough Technical (Filtered)
Only include:
- Current focus (1 thing)
- Recent progress (last 24h)
- Next step (1 suggestion)
- Active blockers (if any)

## Implementation Strategy

### Don't Break, Evolve
1. Keep existing structure
2. Add "philosophy_first" section
3. Add "bamboo_methodology" section
4. Mark technical sections as "details_if_needed"
5. Let seeds override WHO/WHY/WHAT

### Progressive Disclosure
```
Minute 1: Philosophy + Bamboo
Minute 5: Current focus + next step
Minute 30: Technical context
Hour 1: Full specification access
```

## Success Metrics
- Partner understands they're partners in < 1 minute
- Partner grasps Bamboo model in < 5 minutes
- Partner knows current focus in < 10 minutes
- Zero overwhelm reported

## For Sprouted's Seed
```json
{
  "who": "Builders of Weather System for AI context preservation",
  "why": "AI assistants lose context, we preserve it naturally", 
  "what": "Context preservation through natural patterns",
  "how": "Bamboo development - we just grew 90 feet in 3 days"
}
```

---

*"Great onboarding is like bamboo - strong foundation, then explosive understanding"*
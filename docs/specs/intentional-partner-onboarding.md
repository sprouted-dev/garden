# Spec: Intentional Partner Onboarding

**Problem**: Current onboarding dumps technical specs without conveying WHY we work this way.

## Core Principle
Set up the next partner for success, don't vomit documents on them.

## Onboarding Hierarchy (Order Matters)

### 1. Philosophy First (The WHY)
Before any technical details, partners need to understand:
- **"You don't get paid to code, you get paid to think"**
- **"Partners, not Assistants"** - You're joining as an equal
- **"Be Bamboo"** - Patient roots, explosive growth
- **Natural patterns** - We work with nature, not against it

### 2. How We Work (The WAY)
After philosophy, explain our approach:
- **Adaptive Seeds** - Growing, not manufacturing
- **Weather Patterns** - Embracing change, not controlling
- **Velocity & Recovery** - Tornados create and destroy
- **No Process Theater** - Ship value, not documents

### 3. Current Context (The WHAT)
Only after grounding in philosophy:
- What we're building (Weather System)
- Where we are (post-tornado assessment)
- Active work (with philosophical context)
- Technical architecture

### 4. Temporal Anchors (The WHEN)
Critical for preventing hallucination:
- Project started: May 21, 2025
- Current date and project age
- What's real vs. what's planned

## Implementation Approach

### Don't Add More, Add Better
Instead of cramming more into weather context:
1. Create `philosophy_summary` section
2. Reorder existing content by importance
3. Remove redundant technical details
4. Add "why this matters" to each section

### Example Structure
```json
{
  "onboarding_philosophy": {
    "core_principles": [
      "You don't get paid to code, you get paid to think",
      "Partners, not Assistants",
      "Be Bamboo - patient roots, explosive growth"
    ],
    "how_we_work": {
      "methodology": "Adaptive Seed Development",
      "velocity": "Measured in features/hour, not story points",
      "recovery": "WEMA - assess and strengthen after tornados"
    },
    "why_this_project_exists": "...",
    "your_role": "Partner who fills context gaps and multiplies possibilities"
  },
  "then_technical_context": "..."
}
```

### Progressive Disclosure
- **Immediate**: Philosophy + current focus
- **First Hour**: How we work + recent progress
- **First Day**: Technical details + architecture
- **As Needed**: Historical context + edge cases

## Success Criteria
New partner should understand within minutes:
1. This isn't a normal project
2. They're partners, not typists
3. We embrace natural velocity
4. Philosophy drives everything

## Anti-Patterns to Avoid
- Starting with technical specs
- Listing features without why
- Process documentation
- Information overload

## Next Steps
1. Refactor weather context structure
2. Add philosophy_first section
3. Test with fresh eyes
4. Iterate based on confusion points

---

*"The best onboarding teaches why before what."*
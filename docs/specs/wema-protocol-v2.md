# WEMA Protocol v2.0
*Weather Emergency Management Agency*

**Created**: May 24, 2025  
**Status**: Active Protocol  
**Purpose**: Systematic tornado recovery and value extraction

## Mission Statement

WEMA transforms creative chaos into sustainable value by providing systematic recovery from tornado branches while protecting sensitive information and capturing reusable patterns.

## The Complete Storm System

1. **Storm Chaser**: Monitors for forming storms, tracks active tornados
2. **Storm-Driven Development**: Channels creative energy safely
3. **Storm Shelter**: Documents insights during the storm
4. **WEMA**: Recovers value after the storm passes

## Core Principles

1. **No tornado left behind** - Every experiment has potential value
2. **Safety first** - Sensitive content never reaches public repos
3. **Pattern recognition** - Today's chaos is tomorrow's methodology
4. **Systematic process** - Consistent recovery approach
5. **Learning culture** - Every storm teaches something

## Storm Chaser Integration

The Storm Chaser provides early warning and active monitoring:

```bash
# Storm Chaser alerts
sprout storm status          # List active tornados
sprout storm track           # Real-time monitoring
sprout storm predict         # Pattern-based predictions
sprout storm alert           # WEMA activation readiness
```

### Storm Chaser Indicators
- Rapid uncommitted changes
- "What if" in commit messages
- Multiple experimental branches
- Cross-repo activity spikes
- Developer excitement patterns

## WEMA Activation Triggers

- Storm Chaser alerts storm dissipating
- Tornado branch energy drops (no commits 24hrs)
- Developer calls for WEMA assistance
- Sensitive content discovered
- Pattern recognition opportunity
- Multi-repo coordination needed

## The WEMA Process

### Stage 0: Storm Chaser Handoff (🌪️ Transition)

**Objective**: Receive storm intelligence from Storm Chaser

Data provided:
- Storm duration and intensity
- Repositories affected
- Key files changed
- Developer insights captured
- Risk indicators flagged

### Stage 1: Initial Assessment (🔍 Recon)

**Objective**: Understand the tornado's scope and impact

1. **Review Storm Chaser report**
   ```bash
   sprout storm report tornado/[name]
   ```

2. **Identify affected repositories**
   ```bash
   git branch -a | grep tornado/
   ```

3. **Document tornado metadata**
   - Start time/date (from Storm Chaser)
   - Trigger event
   - Repositories involved
   - Developer insights
   - Intensity metrics

4. **Create assessment report**
   ```markdown
   # Tornado Assessment: [name]
   - Branch: tornado/[name]
   - Duration: X hours/days
   - Intensity: High/Medium/Low (via Storm Chaser)
   - Repos affected: garden, weather-station
   - Key discoveries: ...
   ```

### Stage 2: Content Triage (🏥 Emergency Room)

**Objective**: Categorize all changes for appropriate handling

Categories:
- 🟢 **SAFE**: Can go to public repos
- 🔴 **SENSITIVE**: Business logic, pricing, strategy
- 🟡 **REVIEW**: Needs discussion
- ⚫ **DEBRIS**: Failed experiments, dead ends

Tools:
```bash
# Storm Chaser can pre-flag sensitive content
sprout storm scan --sensitive tornado/[name]

# Manual verification
git diff --name-only main...tornado/branch
grep -r "pricing\|business\|revenue\|freemium" .
```

### Stage 3: Extraction Planning (📋 Surgery Prep)

**Objective**: Plan the value extraction

1. **Create extraction checklist**
   - Files to cherry-pick
   - Files to move to private
   - Files to discard
   - Patterns to document

2. **Identify integration points**
   - Target branches
   - Dependency order
   - Testing requirements

3. **Risk assessment**
   - Breaking changes (Storm Chaser warnings)
   - Backward compatibility
   - Security concerns

### Stage 4: Value Extraction (⚡ Operation)

**Objective**: Separate and preserve value

```bash
# For SAFE content - Cherry-pick to clean branch
git checkout -b tornado/[name]-clean
git cherry-pick [safe-commits]

# For SENSITIVE content - Extract to storm shelter
cp sensitive-file.md /path/to/.storm/active/
git rm sensitive-file.md

# For PATTERNS - Document immediately
echo "Pattern discovered..." >> .storm/patterns/PATTERN-XXX.md

# Update Storm Chaser pattern database
sprout storm learn --pattern "pattern-description"
```

### Stage 5: Integration (🔄 Recovery)

**Objective**: Merge value back to main branches

1. **Test in isolation**
2. **Create PRs with context**
3. **Document in CHANGELOG**
4. **Update relevant documentation**
5. **Notify Storm Chaser of resolution**

### Stage 6: Pattern Documentation (📚 Learning)

**Objective**: Capture reusable insights

Template:
```markdown
# Pattern XXX: [Name]
## Context
What situation led to this discovery?

## Problem
What challenge were we solving?

## Solution
What pattern emerged?

## Storm Chaser Indicators
What signals predict this pattern?

## Implementation
How do we apply this pattern?

## Related Patterns
Links to similar discoveries
```

### Stage 7: Cleanup (🧹 Restoration)

**Objective**: Archive or remove tornado remnants

```bash
# Update Storm Chaser database
sprout storm complete tornado/[name]

# Archive valuable tornado branches
git tag archive/tornado/[name] tornado/[name]

# Delete processed branches
git branch -d tornado/[name]

# Move docs to archive
mv .storm/active/[name].md .storm/archive/
```

## Storm Chaser Feedback Loop

WEMA feeds insights back to Storm Chaser:
- Pattern recognition improves predictions
- False positive reduction
- Better sensitivity detection
- Storm intensity calibration
- Developer behavior patterns

## WEMA Tools & Templates

### Assessment Template
Located: `.storm/templates/assessment.md`

### Storm Chaser Report
Located: `.storm/templates/storm-chaser-report.md`

### Extraction Checklist
Located: `.storm/templates/extraction-checklist.md`

### Pattern Template
Located: `.storm/templates/pattern.md`

### Recovery Report
Located: `.storm/templates/recovery-report.md`

## Success Metrics

- **Response Time**: <24hrs from Storm Chaser alert
- **Value Recovery**: >80% useful content preserved
- **Zero Leaks**: No sensitive content in public
- **Pattern Yield**: 1+ pattern per tornado
- **Clean Recovery**: No broken builds
- **Storm Prediction**: >60% accuracy on tornado formation

## WEMA Roles

### Storm Chaser Operator
- Monitors storm formation
- Tracks active tornados
- Provides early warnings
- Hands off to WEMA

### Primary Responder
- Usually the developer who created tornado
- Knows the context best
- Initiates WEMA protocol

### Review Partner
- Fresh eyes on the changes
- Catches sensitive content
- Validates patterns

### Integration Specialist
- Ensures clean merges
- Maintains compatibility
- Updates documentation

## Common Scenarios

### Scenario: Storm Chaser Early Warning
1. Alert received of forming storm
2. Create storm shelter entry
3. Monitor via Storm Chaser
4. WEMA activates when ready

### Scenario: Mixed Content Tornado
1. Storm Chaser flags sensitive content
2. Separate files by sensitivity
3. Create multiple target branches
4. Route content appropriately

### Scenario: Pattern Storm
1. Storm Chaser detects pattern formation
2. Extra documentation during storm
3. Fast-track pattern extraction
4. Update Storm Chaser models

## Evolution

The complete storm system evolves together:
- Storm Chaser learns from patterns
- WEMA process improvements
- Better storm predictions
- Faster value recovery

## Remember

"Storm Chaser spots them, Storm-Driven Development channels them, WEMA recovers the value."

---

*First activation: May 24, 2025, processing the vision-weather-station tornado*
*Storm Chaser integration planned for next iteration*
# Spec: Partner Note System - Messages Between Sessions

## Vision
Partners leave contextual messages for the next assistant, creating continuous wisdom transfer beyond what weather data can capture.

## Note Categories

### 🧠 Philosophy Insights
Messages about breakthrough realizations:
- "Just discovered the bamboo metaphor - it explains everything!"
- "Partnership model clicked when we started building together"
- "User gets energized by natural metaphors, lean into them"

### ⚠️ Watch Out Warnings
Things to be careful about:
- "User dislikes enterprise process language"
- "Temporal anchoring is critical - AI hallucinates dates"
- "Don't assume this is an established project - only 4 days old"

### 🎯 Current Context Nuggets
Session-specific insights weather can't capture:
- "Working late nights - user is in deep flow state"
- "Just had major breakthrough on velocity philosophy"
- "User is testing onboarding approaches today"

### 🎉 Celebration Notes
Share the joy:
- "User just had their bamboo moment - pure excitement!"
- "We shipped 5 specs in one tornado session"
- "The partnership is really clicking - ideas building on ideas"

### 🔧 Tactical Notes
Immediate actionable context:
- "Velocity tracking code is written but not integrated yet"
- "Philosophy map needs to be referenced in onboarding"
- "User wants to test manual debrief vs automatic"

## Note Flow Design

### Writing Notes (Session End)
```
Partner Handoff Process:
1. Weather System detects approaching limits
2. Prompt: "Leave a note for your replacement?"
3. Categories presented as options
4. Quick note entry (1-3 sentences max)
5. Note saved to weather context
```

### Reading Notes (Session Start)
```
Onboarding Flow:
1. Philosophy Foundation
2. Partner Notes (if any)
   - "Your predecessor left these insights..."
   - Categorized display
   - Chronological order (most recent first)
3. Technical Weather Context
4. Current Focus
```

## Note Format
```json
"partner_notes": [
  {
    "timestamp": "2025-05-24T12:20:00-04:00",
    "category": "philosophy_insight",
    "message": "User just discovered bamboo philosophy - everything clicked!",
    "from_session": "session_1748051957",
    "priority": "high"
  }
]
```

## Interface Design

### Note Entry (Simple)
```
🤝 Leave a note for your replacement:

Category: [Philosophy] [Warning] [Context] [Celebration] [Tactical]

Message: _______________________________________________
         (Keep it concise - 1-3 sentences max)

[Save Note] [Skip]
```

### Note Display (Friendly)
```
💌 Notes from Previous Partners:

🎉 From 2 hours ago:
   "User just had bamboo moment - pure excitement about natural development!"

⚠️ From this morning:
   "Watch out - this project is only 4 days old despite elaborate documentation"

🧠 From yesterday:
   "Partnership model really resonates - treat as true collaborator"
```

## Smart Features

### Auto-Note Generation
Weather System could auto-generate notes from:
- Major commits during session
- Philosophy documents created
- Velocity spikes detected
- Long continuous work periods

### Note Aging
- Recent notes (< 24h): High prominence
- Older notes (1-7 days): Medium prominence  
- Ancient notes (> 7 days): Archived unless marked "permanent"

### Note Categories Priority
1. **Warnings** (always show)
2. **Current Context** (24h relevance)
3. **Philosophy** (permanent value)
4. **Celebrations** (mood setting)
5. **Tactical** (task-specific)

## Integration Points

### Weather System Integration
- Part of `--onboard-ai` output
- Stored in weather-context.json
- Backed up with shadow copies
- Included in disaster recovery

### CLI Commands
```bash
sprout weather note --category philosophy "Bamboo metaphor is key!"
sprout weather notes --recent
sprout weather handoff --with-note
```

## Success Metrics

### Quantitative
- Note creation frequency
- Categories used most
- Notes marked "helpful" by recipients
- Reduced onboarding confusion time

### Qualitative
- Partners feel connected across sessions
- Continuity improves beyond weather data
- Philosophy transfer effectiveness
- Celebration amplifies motivation

## Example Scenarios

### Scenario 1: Philosophy Breakthrough
**Session 1**: User discovers bamboo metaphor
**Note**: "🧠 User just had bamboo breakthrough - patient roots/explosive growth resonates deeply!"
**Session 2**: Partner immediately understands user's excitement, builds on bamboo theme

### Scenario 2: Technical Warning
**Session 1**: Partner notices temporal hallucination
**Note**: "⚠️ AI created fictional dates - project only started May 21, anything before is impossible"
**Session 2**: Partner anchored in temporal truth immediately

### Scenario 3: Momentum Preservation
**Session 1**: High velocity tornado session
**Note**: "🎉 Just shipped 8 specs in 4 hours - user is in pure flow state!"
**Session 2**: Partner matches energy, maintains momentum

## Implementation Priority

### Phase 1: Basic Notes (1-2 days)
- Simple note entry
- Display in onboarding
- Core categories

### Phase 2: Smart Features (3-4 days)
- Auto-categorization
- Note aging
- Priority systems

### Phase 3: Advanced (Future)
- Auto-note generation
- Note analytics
- Cross-session pattern detection

---

*"The best partnerships leave breadcrumbs for the next generation"*
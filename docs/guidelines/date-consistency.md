# Date Consistency Guidelines

## Core Principle
Dates anchor us in reality. Without temporal truth, we drift into fiction.

## Standards

### 1. Always Use ISO 8601 Format
- **Dates**: YYYY-MM-DD (2025-05-24)
- **Timestamps**: YYYY-MM-DDTHH:MM:SS±TZ (2025-05-24T12:06:19-04:00)
- **Why**: Unambiguous, sortable, universal

### 2. Include Timezone Context
- Always specify timezone offset or name
- Use local time with offset (more human-readable)
- Example: "2025-05-24 12:06 AM EDT" or "2025-05-24T00:06:19-04:00"

### 3. Temporal Anchoring Requirements
Every document should include:
```markdown
*Created: YYYY-MM-DD*
*Last Updated: YYYY-MM-DD*
```

### 4. Reality Checks
Before writing any date:
1. Check project start date (git log)
2. Verify against current date
3. Ensure temporal possibility

### 5. Relative Time Guidelines
- **Avoid**: "Next month", "Last week"
- **Prefer**: "By 2025-06-24", "Since 2025-05-21"
- **Exception**: Very near-term ("Tomorrow", "Later today")

### 6. AI Session Reminders
When starting new AI sessions, include:
```
Current date: May 24, 2025
Project started: May 21, 2025
Project age: 3 days
Any date before May 21, 2025 is impossible
```

### 7. Historical Accuracy
- Never create fictional past events
- Document when things actually happened
- Preserve evidence of temporal errors as case studies

### 8. Update Obligations
When editing documents:
- Update "Last Updated" date
- Don't change "Created" date
- Note significant revisions

## Common Patterns to Avoid

### ❌ Enterprise Fiction
"Planning began in Q4 2024..." (when project started yesterday)

### ❌ Vague Futures  
"Coming in the next sprint..." (we don't have sprints)

### ❌ Relative Confusion
"As discussed last month..." (project is 3 days old)

### ✅ Temporal Honesty
"Created May 24, 2025 during tornado development"

## Implementation

### Git Hooks
Consider pre-commit hooks that:
- Validate dates aren't before project start
- Check ISO 8601 format
- Flag suspicious temporal claims

### Weather Integration
Weather System should:
- Track document creation dates
- Flag temporal impossibilities
- Provide current date context

## The Bamboo Connection

Like bamboo, some things take time underground:
- Ideas may percolate for years
- But documentation dates from manifestation
- Not from first thought

## Remember

> "Time is the only truth we can't refactor"

Every date we write becomes part of the historical record. Make it true.

---

*Created: 2025-05-24*
*Last Updated: 2025-05-24*
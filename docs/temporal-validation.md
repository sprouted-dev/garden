# Temporal Validation Guide

*Preventing AI hallucination through rigorous date verification*

## The Temporal Truth

**ABSOLUTE FACT**: The Sprouted project started on May 21, 2025 at 7:11 PM EDT.

Any reference to dates before May 21, 2025 in this codebase is a hallucination. This includes:
- December 2024 planning sessions (never happened)
- January 2025 launches (impossible)  
- Early 2025 development (fictional)
- Any backstory before May 21, 2025

## Why This Matters

AI assistants pattern-match to create plausible timelines. They see enterprise projects and assume months of planning. They see version 0.2.0 and invent a version history. This creates false documentation that pollutes the codebase with fiction.

### Real Example
An AI assistant wrote about "Usage limit recovery from January 23, 2025" when the event actually occurred on May 23, 2025. The assistant time-traveled 4 months into the past because January felt more plausible for a "mature" feature.

## Technical Implementation

### 1. Date Format Standards
Always use ISO 8601 format:
```
YYYY-MM-DD for dates: 2025-05-24
YYYY-MM-DDTHH:MM:SS±TZ for timestamps: 2025-05-24T18:07:00-04:00
```

### 2. Required Metadata
Every document must include:
```markdown
*Created: May 24, 2025*
*Last Updated: May 24, 2025*
```

### 3. Validation Functions
```go
func ValidateDate(date time.Time) error {
    projectStart := time.Date(2025, 5, 21, 19, 11, 0, 0, time.Local)
    if date.Before(projectStart) {
        return fmt.Errorf("TEMPORAL VIOLATION: Date %s is before project start (May 21, 2025)", 
            date.Format("2006-01-02"))
    }
    return nil
}
```

### 4. Git Integration
```bash
# Pre-commit hook to catch temporal violations
#!/bin/bash
# .git/hooks/pre-commit

# Check for impossible dates in staged files
INVALID_DATES=$(git diff --cached --name-only -z | \
    xargs -0 grep -E "202[0-4]|2025-0[1-4]|2025-05-[01][0-9]|2025-05-20" || true)

if [ -n "$INVALID_DATES" ]; then
    echo "❌ TEMPORAL VIOLATION: Found references to impossible dates"
    echo "$INVALID_DATES"
    exit 1
fi
```

## AI Context Requirements

### Onboarding Context
```json
{
  "temporal_boundaries": {
    "project_started": "2025-05-21T19:11:00-04:00",
    "current_date": "2025-05-24",
    "validation": {
      "before_start": "IMPOSSIBLE - This is a hallucination",
      "valid_range": ["2025-05-21", "current_date"]
    },
    "warning": "⚠️ ANY date before May 21, 2025 is FICTIONAL"
  }
}
```

### Temporal Anchoring
Start every AI session with:
```
This project started on May 21, 2025. Today is [current_date].
Any dates before May 21, 2025 are impossible hallucinations.
The project is only [X] days old.
```

## Documentation Guidelines

### When Writing Dates
1. **Check Reality**: Is this date actually possible?
2. **Verify with Git**: Does git history support this claim?
3. **Add Metadata**: Include created/updated timestamps
4. **Question Patterns**: Does this timeline make logical sense?

### Common Hallucination Patterns
- **"Months of planning"** → Reality: 4 days from idea to launch
- **"Version 0.1.0 in December"** → Reality: All versions after May 21
- **"Early 2025 development"** → Reality: Started May 21, 2025
- **"Q1 2025 milestones"** → Reality: No Q1, project didn't exist

### Updating Documents
When editing any document:
1. Update the "Last Updated" date
2. Verify all dates in the document
3. Fix any temporal violations found
4. Note corrections in commit message

## Case Studies to Preserve

Some temporal errors are valuable learning examples:
- Files claiming January 2025 events (show pattern-matching behavior)
- Version histories extending before May 21 (show assumption patterns)
- Planning documents referencing Q4 2024 (show enterprise mimicry)

Mark these clearly:
```markdown
⚠️ TEMPORAL HALLUCINATION PRESERVED FOR STUDY
[Original incorrect content]
✅ REALITY: [What actually happened]
```

## Validation Checklist

Before committing any date-containing content:

- [ ] All dates are May 21, 2025 or later
- [ ] Document includes created/updated metadata
- [ ] Dates verified against git history
- [ ] No fictional backstory added
- [ ] Temporal anchoring included for AI-generated content

## The Bottom Line

**When in doubt, check git history.** Git doesn't lie. AI assistants do, not from malice but from pattern-matching. Our job is to maintain temporal truth through validation, not assumption.

Remember: This project achieved in 4 days what others do in months. That's the real story - don't let AI assistants rewrite history to match their expectations.
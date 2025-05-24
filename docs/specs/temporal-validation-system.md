# Temporal Validation System Specification

## Problem Statement
AI assistants create fictional timelines despite Weather System having accurate timestamps. They invent plausible-sounding dates based on patterns rather than checking actual data.

## Solution Overview
Multi-layered temporal validation that makes it impossible for AI to hallucinate dates.

## Components

### 1. Enhanced AI Onboarding Output
When running `sprout weather --onboard-ai`, include:

```
⚠️  TEMPORAL BOUNDARIES ⚠️
- PROJECT STARTED: May 21, 2025 (3 days ago)
- FIRST COMMIT: 2025-05-21T19:11:39
- ANY DATES BEFORE May 21, 2025 ARE IMPOSSIBLE
- Current date: May 24, 2025
```

### 2. Temporal Facts in Weather Context
Add new section to weather-context.json:

```json
"temporalFacts": {
  "projectStarted": "2025-05-21T19:11:39-04:00",
  "projectAgeDays": 3,
  "currentDate": "2025-05-24",
  "impossibleDatesBefore": "2025-05-21",
  "validDateRange": {
    "start": "2025-05-21",
    "end": "2025-05-24"
  }
}
```

### 3. Date Validation Functions
Add to Weather System:

```go
func ValidateDate(dateStr string) error {
    // Compare against git history
    // Flag any date before first commit
    // Warn about future dates
}

func ValidateDocument(content string) []DateError {
    // Scan for date patterns
    // Validate each against temporal facts
    // Return list of violations
}
```

### 4. Git Pre-commit Hook
Scan changed files for temporal violations:

```bash
#!/bin/bash
# Check for impossible dates in committed files
# Reject commits with dates before project start
```

### 5. AI Context Warnings
Include in CLAUDE.md and onboarding:

```
CRITICAL: This project started on May 21, 2025.
- Do NOT create dates before this
- Do NOT reference December 2024 or January 2025
- Always check weather-context.json for temporal facts
```

## Expected Outcomes
1. AI assistants cannot create impossible dates without triggering warnings
2. Temporal hallucinations caught before they enter the codebase
3. Clear boundaries prevent pattern-matching to fictional timelines
4. Multiple validation layers ensure accuracy

## Implementation Priority
1. Update AI onboarding output (immediate)
2. Add temporal facts to weather context (high)
3. Create validation functions (medium)
4. Add git hooks (low)

## Success Criteria
- Zero fictional dates in new documentation
- AI assistants reference actual timestamps
- Temporal violations caught automatically
- Clear error messages guide corrections
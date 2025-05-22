# Spec: Weather Automatic Intelligence MVP

## Vision Reference
Related to: [Weather Context Preservation System](/docs/vision/weather-context-preservation.md)

## Overview

The Weather Automatic Intelligence MVP delivers the core value proposition: **zero-maintenance context preservation** that automatically tracks development progress and provides rich context for both developers and AI assistants. This MVP focuses on git-based intelligence with smart inference to eliminate manual maintenance overhead.

## Requirements

### Functional Requirements

#### Core Automatic Intelligence
- [ ] **Git Activity Monitoring**: Automatically detect and process all git commits, branch changes, and repository activity
- [ ] **Smart Context Inference**: Generate human-readable progress summaries from commit messages and patterns
- [ ] **Current Focus Detection**: Automatically infer what the developer is currently working on from recent activity
- [ ] **Session Boundary Detection**: Identify development session start/end based on activity patterns
- [ ] **Progress Tracking**: Maintain running history of recent accomplishments without manual input

#### Weather Context Generation
- [ ] **Automatic Weather Updates**: Update weather context after every git commit automatically
- [ ] **Temperature Calculation**: Automatically calculate activity level based on commit frequency and file changes
- [ ] **Condition Inference**: Determine weather conditions (sunny/cloudy/stormy) from project health indicators
- [ ] **Next Steps Prediction**: Generate suggested next actions based on recent progress patterns
- [ ] **Focus Area Tracking**: Identify which components/areas are currently being developed

#### CLI Interface
- [ ] **Instant Context Display**: `sprout weather` shows automatically-generated current conditions
- [ ] **AI-Friendly Output**: `sprout weather --for-ai` provides rich JSON context for AI assistants
- [ ] **Recent Activity**: `sprout weather recent` shows automatically-captured recent progress
- [ ] **Zero Configuration**: Works immediately in any garden/git repository without setup

#### Context Persistence
- [ ] **Automatic Storage**: Store weather context in `.garden/weather-context.json` without user intervention
- [ ] **Cross-Session Continuity**: Maintain context across development sessions and computer restarts
- [ ] **Garden Integration**: Seamlessly work with existing garden structure and commands

### Non-Functional Requirements

#### Performance
- **Startup Speed**: `sprout weather` command responds in <200ms
- **Background Processing**: Git hook processing completes in <1 second
- **Memory Usage**: <50MB memory footprint for background intelligence
- **File I/O**: Minimal disk writes, efficient JSON updates

#### Reliability
- **Fault Tolerance**: Continue working even if git hooks fail or context file is corrupted
- **Graceful Degradation**: Provide basic context even with limited git history
- **Automatic Recovery**: Self-heal corrupted weather context files

#### Usability
- **Zero Setup**: Works immediately without configuration
- **Transparent Operation**: Intelligence happens in background without interrupting workflow
- **Clear Output**: Weather information is immediately understandable and actionable

## API/Interface Design

### Core Weather Context Model
```typescript
interface WeatherContext {
  // Automatic metadata
  updated: Date;
  sessionId: string;
  gardenPath: string;
  
  // Git-inferred intelligence
  currentFocus: {
    area: string;           // e.g., "authentication system"
    confidence: number;     // 0-1 confidence in inference
    lastActivity: Date;
    inferredFrom: string;   // e.g., "recent commits to auth/ directory"
  };
  
  recentProgress: {
    summary: string;        // e.g., "Implemented user login and JWT validation"
    commits: GitCommit[];   // Recent commits with smart summaries
    timespan: string;       // e.g., "last 2 hours"
    momentum: number;       // 0-100 activity level
  };
  
  nextSteps: {
    suggestions: string[];  // AI-generated next actions
    priority: number;       // 0-100 urgency
    basedOn: string;        // What the suggestions are inferred from
  };
  
  weather: {
    temperature: number;    // 0-100 activity level
    condition: WeatherCondition;
    pressure: number;       // Deadline/urgency pressure
    lastUpdate: Date;
  };
  
  // Git integration
  git: {
    currentBranch: string;
    lastCommit: GitCommit;
    uncommittedChanges: boolean;
    recentBranches: string[];
  };
}

interface GitCommit {
  hash: string;
  message: string;
  smartSummary: string;    // AI-generated human-readable summary
  timestamp: Date;
  filesChanged: string[];
  inferredScope: string;   // e.g., "frontend/auth", "backend/api"
}

type WeatherCondition = 
  | "sunny"           // Smooth progress, no blockers
  | "partly-cloudy"   // Some minor issues or complexity
  | "cloudy"          // Multiple challenges or slow progress  
  | "stormy"          // Major blockers or critical issues
  | "foggy";          // Unclear direction or exploration phase
```

### CLI Commands
```bash
# Core command - show automatically-generated current conditions
sprout weather

# AI assistant format - rich JSON context
sprout weather --for-ai

# Recent activity summary
sprout weather recent

# Current focus and next steps
sprout weather focus

# Raw context data (for debugging)
sprout weather --raw
```

### Git Hook Integration
```bash
# Post-commit hook (automatic)
#!/bin/sh
sprout weather --update-from-commit "$1"

# Post-checkout hook (branch changes)
#!/bin/sh  
sprout weather --update-from-branch-change "$1" "$2"
```

## Core Intelligence Algorithms

### Focus Area Detection
```
1. Analyze recent commits (last 2-8 hours)
2. Extract file paths and group by directory/component
3. Weight by recency and frequency of changes
4. Generate human-readable focus area description
5. Confidence score based on consistency of patterns
```

### Progress Summary Generation
```
1. Collect recent commit messages
2. Extract action verbs and object nouns
3. Group related changes (e.g., auth-related commits)
4. Generate concise summary avoiding technical jargon
5. Include timespan and momentum indicators
```

### Next Steps Prediction
```
1. Analyze current progress patterns
2. Identify incomplete work (TODOs, failing tests, etc.)
3. Look for natural continuation of current focus
4. Generate 2-3 actionable next step suggestions
5. Rank by priority and logical sequence
```

### Weather Condition Mapping
```
Temperature (Activity Level):
- 0-20: Cold (inactive, no recent commits)
- 21-40: Cool (low activity, occasional commits)  
- 41-60: Mild (steady progress, regular commits)
- 61-80: Warm (active development, frequent commits)
- 81-100: Hot (intense activity, very frequent commits)

Conditions (Project Health):
- Sunny: Recent commits successful, no obvious blockers
- Partly Cloudy: Mixed progress, some complexity
- Cloudy: Slow progress, multiple revisions, uncertainty
- Stormy: Reverts, failed attempts, obvious blockers
- Foggy: Exploratory commits, unclear direction
```

## Test Scenarios

### Automatic Intelligence Accuracy
- **Test Case 1**: After 3 commits to `auth/` directory, weather should identify "authentication system" as current focus
- **Test Case 2**: After commit "Fix login bug", weather should update recent progress to reflect bug fixing activity
- **Test Case 3**: After 2-hour gap in commits, weather should detect session boundary and adjust temperature
- **Test Case 4**: After commit "Add user registration API", next steps should suggest related auth features

### CLI Interface Responsiveness  
- **Test Case 5**: `sprout weather` command returns context in <200ms on repository with 1000+ commits
- **Test Case 6**: Weather context updates automatically after `git commit` without user intervention
- **Test Case 7**: `sprout weather --for-ai` provides valid JSON that AI assistants can parse

### Cross-Session Continuity
- **Test Case 8**: Context persists correctly after computer restart
- **Test Case 9**: Context remains accurate after switching between multiple git repositories
- **Test Case 10**: Weather gracefully handles corrupted or missing context files

## Acceptance Criteria

### Core Intelligence
- [ ] **95% Focus Accuracy**: Current focus detection matches actual development area 95% of the time
- [ ] **Automatic Updates**: Weather context updates after every git commit without manual intervention
- [ ] **Session Detection**: Correctly identifies development session boundaries based on activity patterns
- [ ] **Progress Summaries**: Generates human-readable progress summaries from commit activity

### Developer Experience
- [ ] **Zero Setup**: Works immediately in any git repository without configuration
- [ ] **Instant Context**: `sprout weather` provides immediately useful context in <200ms
- [ ] **Transparent Operation**: Intelligence happens in background without interrupting normal workflow
- [ ] **Clear Next Steps**: Provides actionable next step suggestions based on current progress

### AI Integration
- [ ] **Rich Context Format**: `sprout weather --for-ai` provides comprehensive JSON context
- [ ] **Conversation Continuity**: AI assistants can understand project state from weather context alone
- [ ] **Dynamic Updates**: Context stays current as development progresses

### Technical Requirements
- [ ] **Performance**: Background processing completes in <1 second per commit
- [ ] **Reliability**: Continues working even with git hook failures or corrupted files
- [ ] **Cross-Platform**: Works on macOS, Linux, and Windows development environments

## Dependencies

### Core Dependencies
- **Git Integration**: Requires git repository and ability to install git hooks
- **Garden Structure**: Integrates with existing `.garden/` directory structure
- **File System Access**: Read/write permissions for weather context file
- **JSON Processing**: Reliable JSON parsing and generation for context storage

### Optional Enhancements
- **Network Access**: For future AI-powered inference improvements
- **Issue Tracker APIs**: For enhanced project health detection
- **CI/CD Integration**: For build status and project health indicators

## Implementation Phases

### Phase 1: Git Intelligence Foundation (2-3 weeks)
- Git commit monitoring and smart inference
- Basic weather context generation and storage
- Core CLI commands (`sprout weather`, `sprout weather --for-ai`)

### Phase 2: Advanced Intelligence (1-2 weeks)
- Session boundary detection and activity patterns
- Next steps prediction and focus area confidence scoring
- Weather condition mapping and temperature calculation

### Phase 3: Polish and Performance (1 week)
- Performance optimization and error handling
- Cross-platform compatibility and testing
- Documentation and integration with existing garden commands

## Success Metrics

### User Experience
- **Context Restoration Time**: <10 seconds to understand current state after any break
- **Intelligence Accuracy**: 95% of automatically-generated context matches actual development state
- **Adoption Rate**: Developers use `sprout weather` at start of 90% of development sessions

### Technical Performance
- **Response Time**: Weather commands respond in <200ms consistently
- **Resource Usage**: <50MB memory footprint for background intelligence
- **Reliability**: 99.9% uptime for automatic weather updates

## Related Documents
- Vision: [Weather Context Preservation System](/docs/vision/weather-context-preservation.md)
- Implementation: [Weather Intelligence Architecture](/docs/tasks/) (to be created)
- Testing: [Weather Intelligence Test Plan](/docs/specs/) (to be created)
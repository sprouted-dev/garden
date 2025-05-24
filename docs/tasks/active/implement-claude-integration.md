# Task: Implement `sprout init --with-claude`

## Status
Active - Implementation needed

## Context
The `sprout init --with-claude` command exists as a placeholder but needs implementation. Existing monitoring scripts at Farm level need integration and updates for new usage limit features.

## Current State
- Placeholder shows "coming soon" message in main.go:628-646
- Monitor scripts exist at `~/.claude/commands/`
- Scripts are NOT aware of new features:
  - `--prepare-cold-handoff`
  - `--include-usage-context`
- Scripts assume fixed paths (~/sprouted)

## Requirements

### 1. Directory-Agnostic Installation
- Init should work in ANY directory (not just gardens)
- Should work with any Seed type
- Create `.claude/` directory in current working directory
- Store paths relative to init location

### 2. Update Existing Scripts
The existing scripts need updates:

#### context-monitor updates:
- Use `--prepare-cold-handoff` instead of custom handoff
- Make paths relative to init directory
- Integrate with Weather System's resilience features

#### onboard-next-assistant updates:
- Add `--include-usage-context` flag for cold starts
- Use relative paths instead of hardcoded ~/sprouted
- Leverage Weather System's context preservation

### 3. Implementation Steps

```go
func handleInitCommand(args []string) {
    // 1. Get current working directory
    cwd, _ := os.Getwd()
    
    // 2. Create .claude directory structure
    claudeDir := filepath.Join(cwd, ".claude")
    // .claude/commands/
    // .claude/context/
    // .claude/config.json
    
    // 3. Copy/generate updated scripts
    // - Update paths to be relative
    // - Add new Weather System features
    // - Configure for current directory
    
    // 4. Create config with paths
    config := ClaudeConfig{
        InitPath: cwd,
        GardenPath: findNearestGarden(cwd),
        MonitoringEnabled: true,
    }
    
    // 5. Optionally start monitor
    if autoStart {
        startContextMonitor()
    }
}
```

### 4. Script Updates Needed

#### Updated context-monitor snippet:
```bash
# Auto-detect paths
INIT_DIR="$(dirname "$(dirname "$0")")"  # .claude/commands -> .claude
WORK_DIR="$(dirname "$INIT_DIR")"        # Directory where init was run

# Find garden if exists
GARDEN_PATH=$(find_garden_path "$WORK_DIR")

# Use new Weather System features
prepare_handoff() {
    if [ -n "$GARDEN_PATH" ]; then
        cd "$GARDEN_PATH"
        ./apps/sprout-cli/build/sprout weather --prepare-cold-handoff
    else
        # Fallback for non-garden directories
        create_basic_handoff
    fi
}
```

#### Updated onboard-next-assistant snippet:
```bash
# Capture with usage context for cold starts
if [ -n "$GARDEN_PATH" ]; then
    cd "$GARDEN_PATH"
    ./apps/sprout-cli/build/sprout weather --onboard-ai --include-usage-context > "$INIT_DIR/context/current-weather.json"
else
    # Basic context capture for non-gardens
    capture_basic_context
fi
```

## Success Criteria
- [ ] `sprout init --with-claude` works in any directory
- [ ] Scripts use relative paths from init location
- [ ] Scripts leverage new Weather System features
- [ ] Monitor auto-starts (optionally)
- [ ] Works with gardens and non-garden projects
- [ ] Proper documentation added

## Testing
1. Test in garden directory
2. Test in non-garden project
3. Test with different Seed types
4. Test handoff flow with new features
5. Test cold start recovery

## Notes
- Consider making monitor a daemon/service
- Could integrate with OS notifications
- Future: Real context metrics via Claude API
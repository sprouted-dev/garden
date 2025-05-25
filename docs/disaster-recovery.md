# Disaster Recovery - Weather System

The Weather System includes multiple layers of protection against data loss and corruption.

## Protection Layers

### 1. Atomic Writes
All context updates use atomic file operations to prevent partial writes:
- Write to temporary file first
- Atomic rename to final location
- No corruption from interrupted writes

### 2. Shadow Copies
Automatic backup created on every context save:
- `weather-context.shadow.json` - Latest known good copy
- Created after successful writes
- Used for quick recovery

### 3. Timestamped Backups
Automatic rotating backups:
- Stored in `.garden/backups/`
- Named with timestamp: `weather-context-20250523-150405.json`
- Keeps last 10 backups automatically
- Older backups auto-deleted

### 4. Integrity Verification
Built-in context validation:
- JSON structure verification
- Required field checking
- Timestamp validation

## Recovery Commands

### Verify Context Integrity
```bash
sprout weather verify
```

Checks:
- ✅ Current context validity
- ✅ Shadow copy existence
- ✅ Available backup count

### Automatic Recovery
```bash
sprout weather recover
```

Recovery process:
1. Tries shadow copy first (fastest)
2. Lists available backups if shadow fails
3. Preserves corrupted file for analysis

### Recover from Specific Backup
```bash
sprout weather backups              # List available
sprout weather recover <backup-name> # Restore specific
```

## Disaster Scenarios

### Scenario 1: Corrupted Context
```bash
# Context gets corrupted
$ sprout weather
Error loading weather context: invalid JSON

# Verify the issue
$ sprout weather verify
❌ Context verification failed: invalid JSON structure

# Recover automatically
$ sprout weather recover
✅ Successfully recovered from shadow copy
```

### Scenario 2: Missing Context
```bash
# Context file deleted
$ rm .garden/weather-context.json

# Weather system handles gracefully
$ sprout weather recover
✅ Successfully recovered from shadow copy
```

### Scenario 3: Both Files Corrupted
```bash
# Both main and shadow corrupted
$ sprout weather recover
⚠️ Shadow recovery failed
📦 Available backups:
   1. weather-context-20250523-150405.json
   2. weather-context-20250523-140302.json

# Recover from backup
$ sprout weather recover weather-context-20250523-150405.json
✅ Successfully recovered from backup
```

## Best Practices

### Prevention
1. **Regular Git Commits**: Weather context is tracked in git
2. **Automatic Backups**: Happen on every save
3. **Shadow Copies**: Always maintain redundancy

### Recovery Priority
1. **Shadow Copy**: Most recent, fastest recovery
2. **Timestamped Backups**: Historical states available
3. **Git History**: Last resort, check previous commits

### Monitoring
```bash
# Regular health check
sprout weather verify

# Check backup status  
sprout weather backups
```

## Architecture Benefits

- **Zero Configuration**: Works out of the box
- **Automatic Protection**: No manual backup needed
- **Fast Recovery**: Shadow copy enables instant restore
- **History Preservation**: Rotating backups maintain history
- **Graceful Degradation**: Multiple fallback options

## Integration with Usage Limits

This disaster recovery perfectly complements our usage limit recovery:
- **Usage Limits**: Protects against session interruption
- **Disaster Recovery**: Protects against data corruption
- **Together**: Complete protection for your development context

---

*Your work is protected at every level - from typing to commits to backups.*
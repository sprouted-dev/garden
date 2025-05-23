# Architectural Discovery: Event JSON Parsing Bug

**Date**: 2025-05-23
**Discovered By**: nutmeg & Claude during documentation consolidation
**Discovery Type**: Bug/Implementation Gap
**Impact Level**: High - Prevents Farm event processing

## The Discovery

While checking Farm-level events during our documentation consolidation work, we discovered that all pending events fail to process due to JSON unmarshaling errors. The error message indicates newline characters in commit messages are breaking JSON parsing.

## Context That Led to Discovery

1. Ran `sprout weather` and noticed we're using the event system
2. Checked for pending events: found 7 in `.farm/events/pending/`
3. Attempted to process with `sprout farm process`
4. All events failed with: `invalid character '\n' in string literal`

## Evidence/Examples

**Event file that fails** (`1747958096-commit-22522.json`):
```json
{
  "payload": {
    "message": "feat: implement Weather System v0.2.0 - Event-Driven Farm Orchestration\n\nMajor enhancements..."
  }
}
```

The multi-line commit messages contain literal newlines that break JSON parsing.

**Error output**:
```
Error unmarshaling event 1747958096-commit-22522.json: invalid character '\n' in string literal
```

## Why This Matters

1. **Events accumulate but never process** - The pending directory will grow indefinitely
2. **Farm orchestration is broken** - Cross-garden patterns can't be detected
3. **Core feature unusable** - The event system we built to solve Farm invisibility doesn't work
4. **Data loss risk** - Valuable cross-garden insights trapped in unprocessable events

## Root Cause Analysis

The git hook is emitting JSON with unescaped newlines in string fields. When Go's `json.Unmarshal` tries to parse these files, it fails because JSON strings must escape newlines as `\n`.

## Potential Solutions

### Option 1: Fix Event Emission (Recommended)
Update the git hook to properly escape JSON strings:
```bash
# In git hook
commit_message=$(echo "$commit_message" | jq -Rs .)
```

### Option 2: Fix Event Processing
Make the processor more resilient to malformed JSON:
```go
// Pre-process the JSON to fix known issues
content = strings.ReplaceAll(content, "\n", "\\n")
```

### Option 3: Use Different Format
Switch from JSON to a format that handles multi-line strings better (YAML, TOML).

## Immediate Workaround

Manually fix the pending events:
```bash
for f in .farm/events/pending/*.json; do
  jq . "$f" > "$f.tmp" && mv "$f.tmp" "$f"
done
```

## Related Discoveries

- This validates our "eating our own dog food" approach
- Shows why we need resilience in the Weather System
- Demonstrates the farm visibility challenge in action

## Action Items

- [ ] Fix git hook JSON emission
- [ ] Add JSON validation to event emission
- [ ] Create event processor error recovery
- [ ] Add tests for multi-line commit messages
- [ ] Document event format specification

## Lessons Learned

1. **Test with real data** - Multi-line commits are common
2. **Validate at boundaries** - Check JSON validity before writing
3. **Graceful degradation** - System should handle malformed events
4. **Monitor the monitors** - Need visibility into the event system itself

## Preservation Notes

**Why document this**: This bug blocks the entire Farm orchestration system and demonstrates the importance of robust event handling.

**Cross-references**: 
- [Farm Orchestration Layer Spec](../../garden/docs/specs/farm-orchestration-layer.md)
- [Event System Implementation](../../garden/libs/weather/events.go)

**Keywords**: event-system, json-parsing, farm-orchestration, bug, resilience

---

*This discovery shows that even our context preservation system needs context preservation for its own issues!*
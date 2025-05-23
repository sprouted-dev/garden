# Bug: Event JSON Parsing Fails on Multi-line Commit Messages

**Status**: 🔴 Active
**Severity**: High
**Component**: Event System / Farm Orchestration
**Discovered**: 2025-05-23

## Description

The Farm event processor fails to parse any events that contain multi-line strings (like commit messages). All events with newlines fail with:
```
Error unmarshaling event [filename]: invalid character '\n' in string literal
```

## Reproduction Steps

1. Make a commit with a multi-line message
2. Check `.farm/events/pending/` - event file is created
3. Run `sprout farm process`
4. Observe JSON parsing errors

## Current Impact

- 7 events stuck in pending state
- Farm orchestration completely broken
- Cross-garden pattern detection not working
- Events accumulating but never processing

## Root Cause

The git post-commit hook emits JSON with unescaped newlines:
```json
"message": "feat: implement Weather System\n\nThis breaks JSON"
```

Should be:
```json  
"message": "feat: implement Weather System\\n\\nThis escapes properly"
```

## Proposed Fix

Update `/tools/git-hooks/post-commit` to properly escape JSON strings:
```bash
# Current (broken)
commit_message=$(git log -1 --pretty=%B)

# Fixed
commit_message=$(git log -1 --pretty=%B | jq -Rs .)
```

## Workaround

Until fixed, manually process events:
```bash
cd .farm/events/pending
for f in *.json; do
  jq . "$f" > "$f.fixed" && mv "$f.fixed" "$f"
done
```

## Test Plan

1. Create commit with multi-line message
2. Verify event file has escaped newlines
3. Run event processor successfully
4. Add unit test for multi-line messages

## Related

- [Event JSON Parsing Bug Discovery](/docs/discoveries/event-json-parsing-bug.md)
- [Farm Orchestration Spec](/docs/specs/farm-orchestration-layer.md)
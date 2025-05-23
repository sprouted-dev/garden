#!/bin/bash
# Weather System Disaster Recovery
# Emergency recovery when Weather System data is lost

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo "🚨 Weather System Disaster Recovery"
echo "==================================="
echo ""

# Check what's missing
MISSING_FILES=()
GARDEN_PATH="."

if [ ! -f "$GARDEN_PATH/.garden/weather-context.json" ]; then
    MISSING_FILES+=("weather-context.json")
    echo "❌ Missing: .garden/weather-context.json"
fi

if [ ! -f "$GARDEN_PATH/weather.md" ]; then
    MISSING_FILES+=("weather.md")
    echo "❌ Missing: weather.md"
fi

if [ ${#MISSING_FILES[@]} -eq 0 ]; then
    echo "✅ All Weather System files present - no recovery needed"
    exit 0
fi

echo ""
echo "🔍 Attempting recovery strategies..."
echo ""

# Strategy 1: Shadow Copies
echo "1️⃣ Checking shadow copies..."
if [ -d "$GARDEN_PATH/.garden/shadows" ]; then
    SHADOW_FOUND=false
    
    if [ -f "$GARDEN_PATH/.garden/shadows/weather-context.shadow.json" ]; then
        echo "   ✅ Found shadow copy of weather-context.json"
        cp "$GARDEN_PATH/.garden/shadows/weather-context.shadow.json" \
           "$GARDEN_PATH/.garden/weather-context.json"
        SHADOW_FOUND=true
    fi
    
    if [ -f "$GARDEN_PATH/.garden/shadows/weather.shadow.md" ]; then
        echo "   ✅ Found shadow copy of weather.md"
        cp "$GARDEN_PATH/.garden/shadows/weather.shadow.md" \
           "$GARDEN_PATH/weather.md"
        SHADOW_FOUND=true
    fi
    
    if [ "$SHADOW_FOUND" = true ]; then
        echo -e "   ${GREEN}✅ Recovered from shadow copies${NC}"
        exit 0
    else
        echo "   ❌ No shadow copies found"
    fi
else
    echo "   ❌ No shadow directory"
fi

# Strategy 2: Recent Backups
echo ""
echo "2️⃣ Checking backups..."
if [ -d "$GARDEN_PATH/.garden/backups" ]; then
    LATEST_BACKUP=$(ls -t "$GARDEN_PATH/.garden/backups" 2>/dev/null | head -1)
    
    if [ -n "$LATEST_BACKUP" ]; then
        echo "   ✅ Found backup: $LATEST_BACKUP"
        BACKUP_DIR="$GARDEN_PATH/.garden/backups/$LATEST_BACKUP"
        
        if [ -f "$BACKUP_DIR/weather-context.json" ]; then
            cp "$BACKUP_DIR/weather-context.json" "$GARDEN_PATH/.garden/weather-context.json"
            echo "   ✅ Restored weather-context.json"
        fi
        
        if [ -f "$BACKUP_DIR/weather.md" ]; then
            cp "$BACKUP_DIR/weather.md" "$GARDEN_PATH/weather.md"
            echo "   ✅ Restored weather.md"
        fi
        
        echo -e "   ${GREEN}✅ Recovered from backup${NC}"
        exit 0
    else
        echo "   ❌ No backups found"
    fi
else
    echo "   ❌ No backup directory"
fi

# Strategy 3: Git History Reconstruction
echo ""
echo "3️⃣ Attempting git reconstruction..."
if [ -d "$GARDEN_PATH/.git" ]; then
    echo "   🔧 Creating minimal weather-context.json from git..."
    
    # Create minimal context
    cat > "$GARDEN_PATH/.garden/weather-context.json" << EOF
{
  "updated": "$(date -u +"%Y-%m-%dT%H:%M:%SZ")",
  "version": "1.0.0",
  "gardenPath": "$GARDEN_PATH",
  "currentFocus": {
    "area": "recovery",
    "confidence": 0.5,
    "lastActive": "$(date -u +"%Y-%m-%dT%H:%M:%SZ")",
    "inferredFrom": "disaster recovery"
  },
  "note": "Reconstructed after data loss"
}
EOF
    
    # Create minimal weather.md
    cat > "$GARDEN_PATH/weather.md" << EOF
# Weather Report

*Reconstructed from git history after data loss*

## Current Status
🔧 Recovery Mode - Minimal context available

## Recent Activity
Run \`sprout weather\` to regenerate full context from git history.

## Recovery Notes
- Original data was lost
- This is a minimal reconstruction
- Full context will rebuild over time
EOF
    
    echo -e "   ${YELLOW}⚠️  Created minimal files${NC}"
    echo "   💡 Run 'sprout weather' to rebuild full context"
    exit 0
else
    echo "   ❌ No git repository found"
fi

# Strategy 4: Manual Recovery Instructions
echo ""
echo -e "${RED}❌ Automatic recovery failed${NC}"
echo ""
echo "📋 Manual Recovery Options:"
echo ""
echo "1. Check for manual backups:"
echo "   - ~/weather-backup/"
echo "   - Cloud storage"
echo "   - Time Machine / System backups"
echo ""
echo "2. Reinitialize Weather System:"
echo "   mkdir -p .garden"
echo "   sprout weather init"
echo ""
echo "3. Check other gardens for examples:"
echo "   cp ../other-project/.garden/weather-context.json .garden/"
echo "   # Then modify for this project"
echo ""
echo "💡 Prevention Tips:"
echo "- Enable protection: sprout protect"
echo "- Regular backups: sprout backup"
echo "- Use git to track weather.md"

exit 1
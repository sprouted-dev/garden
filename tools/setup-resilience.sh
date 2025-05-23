#!/bin/bash
# Setup Weather System Resilience
# Configures automatic protection and recovery mechanisms

set -e

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo "🛡️  Weather System Resilience Setup"
echo "==================================="
echo ""

GARDEN_PATH="${1:-.}"

# Create resilience directories
echo "📁 Creating resilience directories..."
mkdir -p "$GARDEN_PATH/.garden/shadows"
mkdir -p "$GARDEN_PATH/.garden/backups"
mkdir -p "$GARDEN_PATH/.garden/journal"
echo "✅ Directories created"

# Install git hooks
echo ""
echo "🪝 Installing git hooks..."

# Post-commit hook for shadow copies
POST_COMMIT_HOOK="$GARDEN_PATH/.git/hooks/post-commit"
if [ -d "$GARDEN_PATH/.git" ]; then
    cat > "$POST_COMMIT_HOOK" << 'EOF'
#!/bin/bash
# Weather System Resilience - Post Commit Hook

# Create shadow copies after each commit
if [ -f ".garden/weather-context.json" ]; then
    cp ".garden/weather-context.json" ".garden/shadows/weather-context.shadow.json" 2>/dev/null || true
    
    # Rotate old shadows
    for i in 2 1 0; do
        if [ -f ".garden/shadows/weather-context.shadow.$i.json" ]; then
            mv ".garden/shadows/weather-context.shadow.$i.json" \
               ".garden/shadows/weather-context.shadow.$((i+1)).json" 2>/dev/null || true
        fi
    done
    
    # Move current to .0
    if [ -f ".garden/shadows/weather-context.shadow.json" ]; then
        cp ".garden/shadows/weather-context.shadow.json" \
           ".garden/shadows/weather-context.shadow.0.json" 2>/dev/null || true
    fi
fi

# Shadow weather.md if it exists
if [ -f "weather.md" ]; then
    cp "weather.md" ".garden/shadows/weather.shadow.md" 2>/dev/null || true
fi

# Log to journal
echo "$(date -u +"%Y-%m-%dT%H:%M:%SZ") - Shadow copies created" >> .garden/journal/resilience.log
EOF
    
    chmod +x "$POST_COMMIT_HOOK"
    echo "✅ Post-commit hook installed"
else
    echo "⚠️  No .git directory found - skipping git hooks"
fi

# Create cron job for regular backups
echo ""
echo "⏰ Setting up scheduled backups..."

BACKUP_SCRIPT="$GARDEN_PATH/.garden/backup-weather.sh"
cat > "$BACKUP_SCRIPT" << 'EOF'
#!/bin/bash
# Weather System Scheduled Backup

GARDEN_PATH="$(cd "$(dirname "$0")/../.." && pwd)"
cd "$GARDEN_PATH"

# Create timestamped backup
TIMESTAMP=$(date +"%Y%m%d-%H%M%S")
BACKUP_DIR=".garden/backups/$TIMESTAMP"
mkdir -p "$BACKUP_DIR"

# Backup files
[ -f ".garden/weather-context.json" ] && cp ".garden/weather-context.json" "$BACKUP_DIR/"
[ -f "weather.md" ] && cp "weather.md" "$BACKUP_DIR/"

# Create metadata
cat > "$BACKUP_DIR/backup-metadata.json" << EOL
{
  "timestamp": "$(date -u +"%Y-%m-%dT%H:%M:%SZ")",
  "type": "scheduled",
  "files": ["weather-context.json", "weather.md"]
}
EOL

# Clean old backups (keep last 7 days)
find .garden/backups -name "2*" -type d -mtime +7 -exec rm -rf {} + 2>/dev/null || true

# Log
echo "$(date -u +"%Y-%m-%dT%H:%M:%SZ") - Backup completed: $TIMESTAMP" >> .garden/journal/resilience.log
EOF

chmod +x "$BACKUP_SCRIPT"
echo "✅ Backup script created"

# Create recovery test script
echo ""
echo "🧪 Creating recovery test script..."

TEST_SCRIPT="$GARDEN_PATH/.garden/test-recovery.sh"
cat > "$TEST_SCRIPT" << 'EOF'
#!/bin/bash
# Test Weather System Recovery

echo "🧪 Testing Weather System Recovery"
echo "================================="
echo ""
echo "⚠️  This will temporarily remove weather files to test recovery!"
echo "Press Ctrl+C to cancel, or Enter to continue..."
read

# Backup current files
echo "📦 Backing up current files..."
mkdir -p .garden/test-backup
cp .garden/weather-context.json .garden/test-backup/ 2>/dev/null || true
cp weather.md .garden/test-backup/ 2>/dev/null || true

# Remove files
echo "🗑️  Removing weather files..."
rm -f .garden/weather-context.json
rm -f weather.md

# Test recovery
echo "🚨 Testing recovery..."
if ../tools/weather-recover.sh; then
    echo ""
    echo "✅ Recovery test PASSED"
    
    # Verify files exist
    if [ -f ".garden/weather-context.json" ] && [ -f "weather.md" ]; then
        echo "✅ Files successfully recovered"
    else
        echo "❌ Recovery reported success but files missing"
    fi
else
    echo "❌ Recovery test FAILED"
    
    # Restore from test backup
    echo "🔄 Restoring from test backup..."
    cp .garden/test-backup/* . 2>/dev/null || true
    cp .garden/test-backup/weather-context.json .garden/ 2>/dev/null || true
fi

# Cleanup
rm -rf .garden/test-backup
echo ""
echo "🧪 Test complete"
EOF

chmod +x "$TEST_SCRIPT"
echo "✅ Test script created"

# Create initial shadow copies
echo ""
echo "📋 Creating initial protection..."
if [ -f "$GARDEN_PATH/.garden/weather-context.json" ]; then
    cp "$GARDEN_PATH/.garden/weather-context.json" \
       "$GARDEN_PATH/.garden/shadows/weather-context.shadow.json"
    echo "✅ Shadow copy of weather-context.json created"
fi

if [ -f "$GARDEN_PATH/weather.md" ]; then
    cp "$GARDEN_PATH/weather.md" \
       "$GARDEN_PATH/.garden/shadows/weather.shadow.md"
    echo "✅ Shadow copy of weather.md created"
fi

# Run initial backup
echo ""
echo "💾 Creating initial backup..."
"$BACKUP_SCRIPT"
echo "✅ Initial backup completed"

# Summary
echo ""
echo -e "${GREEN}✅ Weather System Resilience Setup Complete${NC}"
echo ""
echo "🛡️  Protection Features Enabled:"
echo "   • Shadow copies on every commit"
echo "   • Backup script at .garden/backup-weather.sh"
echo "   • Recovery script at tools/weather-recover.sh"
echo "   • Test script at .garden/test-recovery.sh"
echo ""
echo "📋 Next Steps:"
echo "   1. Add to crontab for hourly backups:"
echo "      0 * * * * cd $GARDEN_PATH && .garden/backup-weather.sh"
echo ""
echo "   2. Test recovery:"
echo "      .garden/test-recovery.sh"
echo ""
echo "   3. Check protection status:"
echo "      sprout resilience status"
echo ""
echo -e "${BLUE}💡 Your Weather System is now protected against data loss!${NC}"
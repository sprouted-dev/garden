#!/bin/bash
# Emergency Farm-level protection setup

FARM_PATH="/Users/nutmeg/sprouted"
PROTECTION_DIR="$FARM_PATH/.farm-protection"

echo "🚜 Setting up Farm-level protection..."
echo

# Create protection directory
mkdir -p "$PROTECTION_DIR/backups"

# Create protection script
cat > "$PROTECTION_DIR/backup-farm.sh" << 'EOF'
#!/bin/bash
# Farm-level backup script

FARM_PATH="/Users/nutmeg/sprouted"
BACKUP_DIR="$FARM_PATH/.farm-protection/backups"
TIMESTAMP=$(date +%Y%m%d-%H%M%S)

echo "🚜 Creating Farm backup..."

# Critical directories to backup
CRITICAL_DIRS=(
    "docs"              # Private business docs
    ".farm"             # Farm-level state
    "*.md"              # Root level docs
)

# Create backup archive
BACKUP_FILE="$BACKUP_DIR/farm-backup-$TIMESTAMP.tar.gz"
cd "$FARM_PATH"

# Build tar command
TAR_CMD="tar -czf $BACKUP_FILE"
for dir in "${CRITICAL_DIRS[@]}"; do
    if [[ -e "$dir" ]]; then
        TAR_CMD="$TAR_CMD $dir"
    fi
done

# Execute backup
eval $TAR_CMD 2>/dev/null

echo "✅ Farm backup created: $BACKUP_FILE"

# Keep only last 5 backups
cd "$BACKUP_DIR"
ls -t farm-backup-*.tar.gz | tail -n +6 | xargs rm -f 2>/dev/null

echo "📦 Current backups:"
ls -lh farm-backup-*.tar.gz
EOF

chmod +x "$PROTECTION_DIR/backup-farm.sh"

# Create restore script
cat > "$PROTECTION_DIR/restore-farm.sh" << 'EOF'
#!/bin/bash
# Farm-level restore script

FARM_PATH="/Users/nutmeg/sprouted"
BACKUP_DIR="$FARM_PATH/.farm-protection/backups"

if [ -z "$1" ]; then
    echo "Usage: $0 <backup-file>"
    echo
    echo "Available backups:"
    ls -lh "$BACKUP_DIR"/farm-backup-*.tar.gz
    exit 1
fi

BACKUP_FILE="$1"
if [[ ! -f "$BACKUP_DIR/$BACKUP_FILE" ]]; then
    echo "❌ Backup file not found: $BACKUP_FILE"
    exit 1
fi

echo "🔄 Restoring from: $BACKUP_FILE"
echo "⚠️  This will overwrite existing files. Continue? (y/N)"
read -r confirm

if [[ "$confirm" != "y" ]]; then
    echo "Restore cancelled."
    exit 0
fi

# Create recovery backup first
TIMESTAMP=$(date +%Y%m%d-%H%M%S)
echo "Creating pre-restore backup..."
"$FARM_PATH/.farm-protection/backup-farm.sh"

# Restore
cd "$FARM_PATH"
tar -xzf "$BACKUP_DIR/$BACKUP_FILE"

echo "✅ Farm restored from backup"
EOF

chmod +x "$PROTECTION_DIR/restore-farm.sh"

# Create cron entry suggestion
cat > "$PROTECTION_DIR/setup-cron.txt" << 'EOF'
# Add to crontab for hourly Farm backups:
# crontab -e
# Then add:

0 * * * * /Users/nutmeg/sprouted/.farm-protection/backup-farm.sh >/dev/null 2>&1

# Or for daily backups at 2 AM:
0 2 * * * /Users/nutmeg/sprouted/.farm-protection/backup-farm.sh >/dev/null 2>&1
EOF

echo "✅ Farm protection scripts created!"
echo
echo "📁 Protection directory: $PROTECTION_DIR"
echo "🔧 Scripts created:"
echo "   - backup-farm.sh   (run manually or via cron)"
echo "   - restore-farm.sh  (disaster recovery)"
echo
echo "🏃 Creating initial backup..."
"$PROTECTION_DIR/backup-farm.sh"
echo
echo "💡 To setup automatic backups:"
echo "   cat $PROTECTION_DIR/setup-cron.txt"
echo
echo "🛡️ Your Farm is now protected!"
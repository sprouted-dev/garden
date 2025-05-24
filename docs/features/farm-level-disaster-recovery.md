# Farm-Level Disaster Recovery

## The Critical Vulnerability

In a typical Sprouted ecosystem setup:

```
/path/to/farm/               <- NOT IN GIT! 🚨
├── garden-repo/             <- In git ✅
├── another-repo/            <- In git ✅
├── docs/                    <- PRIVATE DOCS - NOT IN GIT! 🚨
└── .farm/                   <- Farm state - NOT IN GIT! 🚨
```

**The Risk**: If the Farm directory is corrupted or deleted:
- ❌ All private documentation is lost
- ❌ Cross-repository context disappears
- ❌ Organization-level decisions vanish
- ❌ Business-critical files gone forever

## The Solution: Farm Protection

### Quick Setup
```bash
# From any garden
./tools/setup-farm-protection.sh

# Creates protection infrastructure:
.farm-protection/
├── backup-farm.sh       # Backup script
├── restore-farm.sh      # Recovery script
└── backups/             # Backup storage
```

### Protection Commands

#### Check Protection Status
```bash
sprout farm protection-status
```

Shows:
- Protection setup status
- Number of backups
- Latest backup time
- Protected directories

#### Create Manual Backup
```bash
sprout farm backup
```

Creates timestamped backup of:
- `/docs` - Private documentation
- `/.farm` - Farm-level state
- Root-level markdown files

#### Automatic Backups
```bash
# Add to crontab for hourly backups
crontab -e
# Add: 0 * * * * /path/to/farm/.farm-protection/backup-farm.sh

# Or daily at 2 AM:
# 0 2 * * * /path/to/farm/.farm-protection/backup-farm.sh
```

### Recovery Process

#### List Available Backups
```bash
ls -lh /path/to/farm/.farm-protection/backups/
```

#### Restore from Backup
```bash
# Restore latest
/path/to/farm/.farm-protection/restore-farm.sh farm-backup-20250523-202811.tar.gz
```

## Real-World Example

During our usage limit recovery session:

1. **Identified Risk**: Parent directory not in git
2. **Created Manual Backup**: `backup-garden-docs-*`
3. **Implemented Protection**: Farm-level backup system
4. **Validated**: 70KB backup created successfully

## Protection Layers

### Garden Level (Existing)
- ✅ Weather context backups
- ✅ Shadow copies
- ✅ Git versioning

### Farm Level (New)
- ✅ Private docs backup
- ✅ Cross-repo state
- ✅ Business-critical files
- ✅ Automatic rotation

## Best Practices

### 1. Initial Setup
```bash
# Right after creating Farm structure
./tools/setup-farm-protection.sh
```

### 2. Regular Backups
- **Manual**: Before major changes
- **Automatic**: Hourly or daily via cron
- **Verification**: Weekly protection status check

### 3. Testing Recovery
```bash
# Periodically test restore process
cp -r docs docs.test
/path/to/farm/.farm-protection/restore-farm.sh <backup>
diff -r docs docs.test
```

## Why This Matters

The Farm level often contains:
- **Business Strategy**: Not suitable for public repos
- **Private Documentation**: Confidential information
- **Cross-Team Coordination**: Multi-repo context
- **Decision Records**: Organization-wide choices

Without protection, one `rm -rf` could destroy months of work.

## Integration with Weather System

```
Protection Hierarchy:
┌─────────────────────────┐
│     Farm Protection     │ <- Organization backup
├─────────────────────────┤
│   Garden Protection     │ <- Repository backup
├─────────────────────────┤
│    Seed Protection      │ <- Project docs
└─────────────────────────┘
```

Complete protection at every level!

---

*Remember: Gardens are in git, but Farms often aren't. Protect accordingly.*
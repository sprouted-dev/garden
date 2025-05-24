#!/bin/bash
# Setup Farm-level Storm Shelter for Storm-Driven Development
# First WEMA Protocol Implementation

echo "🌪️  Setting up Storm Shelter at Farm level..."

# Check if we're at Farm level
if [ ! -d "garden" ] || [ ! -d "weather-station" ]; then
    echo "❌ Error: Must run from Farm level (parent directory)"
    echo "   Current dir: $(pwd)"
    echo "   Please run from /Users/nutmeg/sprouted"
    exit 1
fi

# Create storm shelter structure
echo "📁 Creating .storm/ directory structure..."
mkdir -p .storm/{active,processed,patterns,archive}

# Create README for storm shelter
cat > .storm/README.md << 'EOF'
# 🌪️ Storm Shelter

Central location for all tornado documentation and unprocessed ideas.

## Structure

- `active/` - Currently active tornados across repos
- `processed/` - WEMA-processed documentation
- `patterns/` - Recurring patterns discovered
- `archive/` - Historical tornado records

## Principles

1. **Code in branches, docs in shelter**
2. **One shelter to rule them all** (no per-repo shelters)
3. **WEMA processes shelter contents** into appropriate locations
4. **Patterns become methodology**

## Current Process

1. Tornado forms (creative burst)
2. Create tornado branches in affected repos
3. Document insights in storm shelter
4. WEMA processes aftermath
5. Value integrated, debris discarded

*Created: May 24, 2025*
*First WEMA Response Active*
EOF

# Move existing storm content if it exists
if [ -d "garden/.storm" ]; then
    echo "📦 Moving existing storm content from garden..."
    cp -r garden/.storm/* .storm/ 2>/dev/null || true
    echo "   ✓ Content moved to Farm shelter"
    echo "   ! Remember to remove garden/.storm after verification"
fi

# Create WEMA protocol
cat > .storm/WEMA_PROTOCOL.md << 'EOF'
# WEMA Protocol v1.0
*Weather Emergency Management Agency*

## Purpose
Process tornado aftermath to extract value and discard debris

## Process

### 1. Initial Assessment
- List all affected repos
- Identify sensitive content
- Document valuable insights

### 2. Triage
- **Critical**: Remove sensitive content from public repos
- **Valuable**: Extract reusable patterns
- **Debris**: Document what didn't work

### 3. Integration
- Cherry-pick safe commits
- Move docs to appropriate locations
- Update repo documentation

### 4. Pattern Recognition
- Document recurring themes
- Update methodology if needed
- Share lessons learned

### 5. Archive
- Move processed content to archive/
- Keep patterns in patterns/
- Clear active/ for next tornado

## First Activation
- Date: May 24, 2025
- Tornado: vision-weather-station
- Status: In Progress
EOF

echo "✅ Storm Shelter created successfully!"
echo ""
echo "📋 Next steps:"
echo "1. cd /Users/nutmeg/sprouted"
echo "2. Review .storm/README.md"
echo "3. Move garden/.storm content"
echo "4. Begin WEMA processing"
echo ""
echo "🌪️  Storm-Driven Development infrastructure ready!"
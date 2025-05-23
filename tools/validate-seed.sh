#!/bin/bash
# Seed Validation Tool
# Checks if your documentation structure is Weather System ready

set -e

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Default to current directory
DIR="${1:-.}"

echo "🌱 Validating Seed structure in: $DIR"
echo "=================================="

# Check for git repository
if [ -d "$DIR/.git" ]; then
    echo "✅ Git repository found"
    SCORE=20
else
    echo "⚠️  No .git directory (required for Weather System)"
    SCORE=0
fi

# Check for docs directory
if [ -d "$DIR/docs" ] || [ -d "$DIR/documentation" ]; then
    echo "✅ Documentation directory found"
    SCORE=$((SCORE + 20))
    
    DOCS_DIR="$DIR/docs"
    [ -d "$DIR/documentation" ] && DOCS_DIR="$DIR/documentation"
    
    # Check for README
    if [ -f "$DOCS_DIR/README.md" ]; then
        echo "✅ README.md found"
        SCORE=$((SCORE + 20))
        
        # Check README content
        if grep -qi "how we work" "$DOCS_DIR/README.md"; then
            echo "✅ README explains workflow"
            SCORE=$((SCORE + 10))
        else
            echo "⚠️  README should explain 'How We Work'"
        fi
        
        if grep -qi "structure" "$DOCS_DIR/README.md"; then
            echo "✅ README documents structure"
            SCORE=$((SCORE + 10))
        else
            echo "⚠️  README should document directory structure"
        fi
    else
        echo "❌ No README.md in docs directory"
    fi
    
    # Count documentation
    DOC_COUNT=$(find "$DOCS_DIR" -name "*.md" -type f | wc -l | tr -d ' ')
    echo "📊 Found $DOC_COUNT documentation files"
    
    # Check for special files
    if [ -f "$DIR/CLAUDE.md" ] || [ -f "$DOCS_DIR/AI.md" ]; then
        echo "✅ AI context file found"
        SCORE=$((SCORE + 5))
    else
        echo "💡 Consider adding CLAUDE.md for AI assistants"
    fi
    
    # Check for common directories
    if [ -d "$DOCS_DIR/specs" ] || [ -d "$DOCS_DIR/specifications" ]; then
        echo "✅ Specs directory found"
        SCORE=$((SCORE + 5))
    fi
    
    if [ -d "$DOCS_DIR/decisions" ] || [ -d "$DOCS_DIR/ADR" ]; then
        echo "✅ Decision records found"
        SCORE=$((SCORE + 5))
    fi
    
else
    echo "❌ No docs/ or documentation/ directory found"
    echo "   Create one with: mkdir -p $DIR/docs"
fi

# Calculate level
echo ""
echo "📊 Seed Score: $SCORE/100"

if [ $SCORE -ge 80 ]; then
    echo "🏆 Level: ${GREEN}Expert Seed${NC}"
elif [ $SCORE -ge 60 ]; then
    echo "🏆 Level: ${GREEN}Advanced Seed${NC}"
elif [ $SCORE -ge 40 ]; then
    echo "🏆 Level: ${YELLOW}Basic Seed${NC}"
else
    echo "🏆 Level: ${RED}Minimal Seed${NC}"
fi

# Quick start if minimal
if [ $SCORE -lt 40 ]; then
    echo ""
    echo "🚀 Quick Start:"
    echo "   mkdir -p $DIR/docs"
    echo "   cat > $DIR/docs/README.md << 'EOF'"
    echo "# My Project"
    echo ""
    echo "## How We Work"
    echo "Describe your workflow here"
    echo ""
    echo "## Directory Structure"
    echo "- docs/ - Documentation"
    echo "EOF"
fi

# Weather System check
echo ""
if command -v sprout &> /dev/null; then
    echo "✅ Sprout CLI detected"
    echo "   Run 'sprout weather' to see your development weather"
else
    echo "💡 Install Sprout CLI for full Weather System features"
    echo "   Visit: https://sprouted.dev"
fi
#!/bin/bash
# Seed Health Monitor
# Combines validation and drift detection for continuous monitoring

set -e

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m'

# Get directory to check
DIR="${1:-.}"
REPORT_FILE="${2:-seed-health-report.json}"

echo "🏥 Seed Health Monitor"
echo "====================="
echo "📂 Checking: $DIR"
echo "📅 Date: $(date)"
echo ""

# Create report structure
cat > "$REPORT_FILE" << EOF
{
  "scan_date": "$(date -u +"%Y-%m-%dT%H:%M:%SZ")",
  "directory": "$DIR",
  "health_score": 0,
  "validation": {},
  "drift": {},
  "recommendations": []
}
EOF

# Run validation
echo "🌱 Running Seed Validation..."
echo "----------------------------"
VALIDATION_OUTPUT=$(./tools/validate-seed.sh "$DIR" 2>&1)
VALIDATION_SCORE=$(echo "$VALIDATION_OUTPUT" | grep "Seed Score" | grep -o '[0-9]\+' | head -1)
echo "$VALIDATION_OUTPUT" | grep -E "✅|❌|⚠️|💡|🏆"

# Run drift detection
echo ""
echo "🌊 Running Drift Detection..."
echo "----------------------------"
if command -v python3 &> /dev/null; then
    DRIFT_OUTPUT=$(python3 ./tools/detect-drift.py "$DIR" 2>&1)
    DRIFT_ISSUES=$(echo "$DRIFT_OUTPUT" | grep "Found" | grep -o '[0-9]\+' | head -1 || echo "0")
    echo "$DRIFT_OUTPUT" | grep -E "✅|⚠️|🔴|🟡|🔵" | head -10
else
    echo "⚠️  Python3 not found, skipping drift detection"
    DRIFT_ISSUES=0
fi

# Calculate overall health
HEALTH_SCORE=$VALIDATION_SCORE
if [ "$DRIFT_ISSUES" -gt 0 ]; then
    # Reduce score based on drift issues
    PENALTY=$((DRIFT_ISSUES * 5))
    HEALTH_SCORE=$((HEALTH_SCORE - PENALTY))
    [ $HEALTH_SCORE -lt 0 ] && HEALTH_SCORE=0
fi

# Generate recommendations
echo ""
echo "📊 Health Summary"
echo "-----------------"
echo "Validation Score: $VALIDATION_SCORE/100"
echo "Drift Issues: $DRIFT_ISSUES"
echo "Overall Health: $HEALTH_SCORE/100"

# Health level
if [ $HEALTH_SCORE -ge 80 ]; then
    echo -e "Status: ${GREEN}Healthy Seed ✨${NC}"
elif [ $HEALTH_SCORE -ge 60 ]; then
    echo -e "Status: ${YELLOW}Minor Issues 🌱${NC}"
elif [ $HEALTH_SCORE -ge 40 ]; then
    echo -e "Status: ${YELLOW}Needs Attention ⚠️${NC}"
else
    echo -e "Status: ${RED}Unhealthy Seed 🏥${NC}"
fi

# Recommendations
echo ""
echo "💡 Recommendations:"
if [ $VALIDATION_SCORE -lt 60 ]; then
    echo "  • Improve basic Seed structure (add README sections)"
fi
if [ "$DRIFT_ISSUES" -gt 5 ]; then
    echo "  • Review and update conflicting documentation"
fi
if [ $VALIDATION_SCORE -lt 40 ]; then
    echo "  • Create minimal docs/README.md to start"
fi
if [ "$DRIFT_ISSUES" -gt 0 ] && [ "$DRIFT_ISSUES" -lt 5 ]; then
    echo "  • Address the $DRIFT_ISSUES drift issues found"
fi
if [ $HEALTH_SCORE -ge 80 ]; then
    echo "  • Keep up the good work! Consider adding more docs"
fi

# Set up monitoring
echo ""
echo "⏰ Continuous Monitoring:"
echo "  Add to crontab for weekly checks:"
echo "  0 9 * * 1 cd $PWD && ./tools/monitor-seed-health.sh"
echo ""
echo "  Or add to git pre-commit hook for every commit"

# Exit code based on health
if [ $HEALTH_SCORE -lt 40 ]; then
    exit 1
else
    exit 0
fi
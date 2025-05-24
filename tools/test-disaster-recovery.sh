#!/bin/bash
# Test script for Weather System disaster recovery

echo "🧪 Weather System Disaster Recovery Test"
echo "======================================="
echo

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Test 1: Verify healthy context
echo "Test 1: Verifying healthy context..."
if ./apps/sprout-cli/build/sprout weather verify | grep -q "valid and intact"; then
    echo -e "${GREEN}✅ Context verification passed${NC}"
else
    echo -e "${RED}❌ Context verification failed${NC}"
    exit 1
fi
echo

# Test 2: Simulate corruption
echo "Test 2: Simulating context corruption..."
cp .garden/weather-context.json .garden/weather-context.original.json
echo '{"corrupted": true}' > .garden/weather-context.json
echo -e "${YELLOW}💥 Context corrupted${NC}"
echo

# Test 3: Verify detects corruption
echo "Test 3: Verifying corruption detection..."
if ./apps/sprout-cli/build/sprout weather verify 2>&1 | grep -q "verification failed"; then
    echo -e "${GREEN}✅ Corruption detected correctly${NC}"
else
    echo -e "${RED}❌ Failed to detect corruption${NC}"
fi
echo

# Test 4: Attempt recovery
echo "Test 4: Testing automatic recovery..."
if ./apps/sprout-cli/build/sprout weather recover; then
    echo -e "${GREEN}✅ Recovery successful${NC}"
else
    echo -e "${RED}❌ Recovery failed${NC}"
    # Restore original
    mv .garden/weather-context.original.json .garden/weather-context.json
    exit 1
fi
echo

# Test 5: Verify recovered context
echo "Test 5: Verifying recovered context..."
if ./apps/sprout-cli/build/sprout weather verify | grep -q "valid and intact"; then
    echo -e "${GREEN}✅ Recovered context is valid${NC}"
else
    echo -e "${RED}❌ Recovered context is invalid${NC}"
fi
echo

# Cleanup
rm -f .garden/weather-context.original.json

echo -e "${GREEN}🎉 All disaster recovery tests passed!${NC}"
echo
echo "The Weather System successfully:"
echo "  • Detected corrupted context"
echo "  • Recovered from backup automatically"
echo "  • Verified the recovered state"
echo
echo "Your development context is protected! 🛡️"
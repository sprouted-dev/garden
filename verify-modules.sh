#!/bin/bash

# Verify Go Module Configuration
# Tests that our custom module paths work correctly

set -e

echo "🔍 Verifying Go module configuration..."

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
RESET='\033[0m'

# Check if we're in the right directory
if [ ! -f "go.work" ] && [ ! -d "apps/sprout-cli" ]; then
    echo -e "${RED}❌ Please run this script from the Garden repository root${RESET}"
    exit 1
fi

echo -e "${YELLOW}📦 Checking module declarations...${RESET}"

# Check weather module
echo "Checking libs/weather module..."
cd libs/weather
if grep -q "module sprouted.dev/weather" go.mod; then
    echo -e "${GREEN}✅ Weather module correctly declared as sprouted.dev/weather${RESET}"
else
    echo -e "${RED}❌ Weather module not using sprouted.dev domain${RESET}"
    exit 1
fi

# Check sprout-cli module  
echo "Checking apps/sprout-cli module..."
cd ../../apps/sprout-cli
if grep -q "module sprouted.dev/sprout-cli" go.mod; then
    echo -e "${GREEN}✅ Sprout CLI module correctly declared as sprouted.dev/sprout-cli${RESET}"
else
    echo -e "${RED}❌ Sprout CLI module not using sprouted.dev domain${RESET}"
    exit 1
fi

# Check import statement
echo "Checking import statements..."
if grep -q '"sprouted.dev/weather"' main.go; then
    echo -e "${GREEN}✅ Import statement using sprouted.dev/weather${RESET}"
else
    echo -e "${RED}❌ Import statement not using sprouted.dev domain${RESET}"
    exit 1
fi

# Test build
echo -e "${YELLOW}🔨 Testing build...${RESET}"
cd ../..
if make build >/dev/null 2>&1; then
    echo -e "${GREEN}✅ Build successful with new module paths${RESET}"
else
    echo -e "${RED}❌ Build failed with new module paths${RESET}"
    exit 1
fi

# Test weather demo
echo -e "${YELLOW}🌦️ Testing weather system...${RESET}"
if make weather-demo >/dev/null 2>&1; then
    echo -e "${GREEN}✅ Weather system working with new module paths${RESET}"
else
    echo -e "${RED}❌ Weather system failed with new module paths${RESET}"
    exit 1
fi

echo ""
echo -e "${GREEN}🎉 All module configuration checks passed!${RESET}"
echo ""
echo -e "${YELLOW}Next steps:${RESET}"
echo "1. Deploy sprouted.dev with Go module redirects"
echo "2. Test with: go get sprouted.dev/weather"
echo "3. Test with: go get sprouted.dev/sprout-cli"
echo ""
echo -e "${YELLOW}Configuration files created:${RESET}"
echo "• docs/sprouted.dev-config.md - Server configuration guide"
echo "• docs/go-import-template.html - HTML template for Go imports"
echo ""
echo -e "${YELLOW}Module paths:${RESET}"
echo "• sprouted.dev/weather (instead of github.com/sprouted-dev/garden/libs/weather)"
echo "• sprouted.dev/sprout-cli (instead of github.com/sprouted-dev/garden/apps/sprout-cli)"
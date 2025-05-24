# Tornado Seed Separation - Immediate Actions

*Created: May 24, 2025*
*Tornado: seed-separation-20250524*

## Next Steps (In Order)

### 1. Create Storm Shelter Structure
```bash
# At farm level (/Users/nutmeg/sprouted/)
mkdir -p .storm/active
mkdir -p .storm/completed
mkdir -p .storm/patterns

# Create initial storm documentation
echo "# Seed Separation Tornado - May 24, 2025" > .storm/active/seed-separation-20250524.md
```

### 2. Create Tornado Branches in Other Repos
```bash
# For each repo (weather-station, sprouted-website):
cd [repo]
git checkout main
git pull origin main
git checkout -b tornado/seed-separation-20250524
```

### 3. Create Sprouted Seed Structure in Garden
```bash
# In garden repo (current tornado branch)
mkdir -p .seed/config
mkdir -p .seed/patterns
mkdir -p .seed/prompts
mkdir -p .seed/philosophy
mkdir -p .seed/plugins/velocity
mkdir -p .seed/plugins/wema
```

### 4. Start With Simple Seed Config
Create `.seed/config.json`:
```json
{
  "name": "sprouted",
  "version": "1.0.0",
  "methodology": "storm-driven-development",
  "description": "Bamboo growth patterns with tornado velocity",
  "metrics": {
    "velocity": {
      "unit": "features",
      "period": "hour"
    }
  },
  "patterns": {
    "branches": {
      "experimental": "tornado/*",
      "storm": "storm/*",
      "feature": "feature/*"
    }
  }
}
```

### 5. Move Velocity Code
- Move `libs/weather/velocity.go` → `.seed/plugins/velocity/velocity.go`
- Update imports and interfaces
- Create plugin loader

### 6. Document Progress in Storm Shelter
Update `.storm/active/seed-separation-20250524.md` with:
- Changes made
- Decisions taken
- Cross-repo impacts
- Next steps

## Order of Implementation

1. **Infrastructure First** (Storm Shelter)
2. **Branch Coordination** (All repos in tornado)
3. **Seed Structure** (Reference implementation)
4. **Code Migration** (Move Sprouted-specific)
5. **Interface Design** (Generic Weather API)
6. **Testing** (Validate separation)
7. **Documentation** (Update all repos)

## Key Decisions to Make

- Plugin architecture: Simple Go files or formal plugin system?
- Configuration format: JSON, YAML, or TOML?
- Backward compatibility: Temporary shim or clean break?
- Cross-repo coordination: Event-based or direct integration?

## Success Checkpoints

- [ ] All repos have tornado branches
- [ ] .storm/ directory is tracking progress
- [ ] .seed/ structure exists in garden
- [ ] velocity.go moved to seed plugin
- [ ] Weather compiles without Sprouted code
- [ ] Sprouted Seed loads correctly

Ready to start creating?
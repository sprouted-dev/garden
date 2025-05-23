# Seeds Quickstart Guide

Get your project working with the Weather System in under 5 minutes.

## What You Need

1. A git repository (or create one with `git init`)
2. The Sprout CLI installed
3. 5 minutes

## Quickest Start (30 seconds)

```bash
# In your project root
mkdir -p docs
echo "# My Project" > docs/README.md
sprout weather
```

That's it! The Weather System is now tracking your project.

## Understanding Seeds

A **Seed** is just a `docs/` directory with a README that explains how your project works. The Weather System reads this to understand your project's context.

## Three Seed Examples

### 1. Minimal Seed (Solo Developer)
Perfect for personal projects or prototypes.

```bash
mkdir -p docs
cat > docs/README.md << 'EOF'
# My Project

## What This Is
A brief description of your project.

## How to Use It
Basic usage instructions.
EOF
```

### 2. Team Seed (Small Teams)
Adds structure for collaboration.

```bash
mkdir -p docs/{specs,decisions,retrospectives}
cat > docs/README.md << 'EOF'
# Team Project

## How We Work
- Specs before code
- Document decisions  
- Regular retrospectives

## Where Things Go
- `specs/` - What we're building
- `decisions/` - Why we chose this way
- `retrospectives/` - What we learned
EOF
```

### 3. Enterprise Seed (Large Organizations)
Compliance and governance ready.

```bash
mkdir -p docs/{architecture,compliance,processes,standards}
# See examples/enterprise-seed for full structure
```

## Verify Your Seed

Check that Weather System understands your Seed:

```bash
# See basic weather
sprout weather

# Get detailed context (what Weather found)
sprout weather --onboard-ai

# Check for documentation gaps
sprout weather --suggest-docs
```

## Common Patterns

Weather System automatically recognizes these patterns:

- `CLAUDE.md` or `AI.md` - AI assistant instructions
- `decisions/` or `ADR/` - Decision records
- `specs/` or `requirements/` - Feature plans
- `TODO.md` or `tasks/` - Work tracking

Name them however makes sense for your project!

## Growing Your Seed

Seeds evolve naturally:

**Week 1**: Just `docs/README.md`
```
docs/
└── README.md
```

**Month 1**: Add structure as needed
```
docs/
├── README.md
├── decisions/
└── setup.md
```

**Month 6**: Mature structure
```
docs/
├── README.md
├── architecture/
├── decisions/
├── guides/
└── troubleshooting/
```

## Tips for Success

### Do ✓
- Start simple
- Document what you actually do
- Let structure emerge naturally
- Use names that make sense to you

### Don't ✗
- Over-engineer from the start
- Force a structure that doesn't fit
- Document for documentation's sake
- Worry about "doing it right"

## FAQ

**Q: What if I already have docs?**
A: Great! Weather System will discover and use them.

**Q: Can I use a different directory name?**
A: Weather also checks `documentation/` and `docs/`.

**Q: What if I have no documentation?**
A: Weather still tracks git activity, but context will be limited.

**Q: How do I know it's working?**
A: Run `sprout weather` - you'll see activity tracking.

## Next Steps

1. Create your minimal Seed
2. Run `sprout weather --onboard-ai`
3. See what Weather discovered
4. Add documentation as you need it
5. Let your Seed grow with your project

Remember: **The best Seed is one that matches how you actually work.**

## Getting Help

- Check examples in `/docs/seeds/examples/`
- Read lessons learned in `/docs/seeds/lessons-learned.md`
- For Farm/Garden setup, see `/docs/seeds/minimum-requirements.md`

Happy growing! 🌱
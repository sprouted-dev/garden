# Documentation Migration Checklist

## Complete Checklist for Confidence Language Update

### 🚨 Priority 1: Public-Facing (Do First)

- [ ] **garden/README.md** - Replace with confident version
- [ ] **Website landing page** (sprouted-website/src/app/(landing)/page.tsx)
- [ ] **Install script messages** (garden/website/install.sh)
- [ ] **CLI help text** (garden/apps/sprout-cli/)
- [ ] **GitHub repository description**

### 📋 Priority 2: Core Documentation

- [ ] **docs/current/ROADMAP.md**
  - [ ] Change "Only if Demanded" → "Early Access"
  - [ ] Update timelines with confidence
  - [ ] Add licensing information

- [ ] **docs/current/STATE.md**
  - [ ] Weather Station: "Not Implemented" → "Integration Pending"
  - [ ] Add "Built in vibes/mono" note
  - [ ] Update all component statuses

- [ ] **docs/current/ARCHITECTURE.md**
  - [ ] Move Weather Station to current architecture
  - [ ] Add integration patterns
  - [ ] Remove "future" sections

- [ ] **docs/current/README.md**
  - [ ] Add complete ecosystem overview
  - [ ] Update project status
  - [ ] Include licensing model

### 📁 Priority 3: Vision & Specs

- [ ] **docs/vision/weather-context-preservation.md**
  - [ ] Present tense for built features
  - [ ] Future tense only for genuine future work
  - [ ] Add implementation status badges

- [ ] **docs/specs/weather-automatic-intelligence-mvp.md**
  - [ ] Update to reflect current reality
  - [ ] Reference Weather Station implementation
  - [ ] Add integration specifications

- [ ] **docs/specs/farm-orchestration-layer.md**
  - [ ] Clarify this extends Weather Station
  - [ ] Position as roadmap item
  - [ ] Remove tentative language

### 📝 Priority 4: Feature Documentation

- [ ] **All files in docs/features/**
  - [ ] Add "Available in Weather Station" where applicable
  - [ ] Clarify free vs. premium features
  - [ ] Update status indicators

### 🌐 Priority 5: External Communications

- [ ] **LinkedIn Company Page**
  - [ ] Update description
  - [ ] Post announcement
  - [ ] Update specialties

- [ ] **Twitter/X Bio** (if exists)
- [ ] **Dev.to Articles** (if any)
- [ ] **Reddit Posts** (update any old ones)

### 🔍 Search & Replace Operations

Run these searches across the codebase:

```bash
# Find tentative language
grep -r "maybe\|possibly\|could be\|if demanded\|might\|potentially" docs/

# Find "not implemented"
grep -r "not implemented\|Not Implemented" docs/

# Find future tense about existing features
grep -r "will be\|plans to\|hopes to\|intends to" docs/
```

### 📊 Language Replacement Table

| Find | Replace With |
|------|--------------|
| "if demanded" | "in early access" |
| "not implemented" | "ready for integration" |
| "could include" | "includes" |
| "might have" | "features" |
| "possibly" | "available in" |
| "future possibility" | "premium feature" |
| "someday" | "Phase 3" |
| "may be" | "is" |

### ✅ Final Verification

- [ ] Run searches - no tentative language found
- [ ] Read through main README - sounds confident
- [ ] Check website - presents complete platform
- [ ] Review roadmap - clear timeline, no "ifs"
- [ ] Test messaging - consistent across all channels

### 🎉 Launch Announcement Draft

Once migration complete, announce:

```
🚀 Big Documentation Update!

We've updated all our docs to better reflect reality:

✅ Weather Station isn't a "future possibility" - it's built
✅ Clear roadmap with real timelines
✅ Transparent about what's free vs. premium
✅ Confident in our complete platform vision

The tentative language is gone. The reality is here.

Check out the new docs: [link]
```

---

**Remember**: We're not changing the facts, just how we present them. Be confident in what we've built!
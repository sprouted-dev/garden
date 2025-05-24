# Weather Station Migration Specification

## Overview

Weather Station already exists as a fully functional Mission Control Panel (MCP) server in the vibes/mono codebase. This spec outlines migrating and adapting it to become the official Sprouted Weather Station - transforming tentative "future possibilities" into working reality.

## Discovery Summary

**What We Found**: A complete real-time development dashboard with:
- WebSocket MCP server for push updates
- Real-time dashboards with weather metaphors
- Bug radar, task forecasts, seasonal tracking
- Alert system and activity feeds
- SQLite database with migrations
- Full REST + WebSocket APIs

**What This Means**: Weather Station isn't a dream - it's built and tested. We just need to connect it to the real Weather System.

## Migration Goals

1. **Preserve What Works**: Keep the proven MCP architecture and dashboard features
2. **Connect to Reality**: Replace mock data with actual Weather System integration
3. **Enhance for Sprouted**: Add cloud backup, multi-garden support, team features
4. **Update Positioning**: Transform docs from "maybe someday" to "available now"

## Technical Migration Plan

### Phase 1: Code Migration (Week 1)

1. **Copy Codebase**
   ```bash
   cp -r ~/vibes/mono/apps/weather-dashboard garden/apps/weather-station
   ```

2. **Update Module Paths**
   - Change imports from `vibes/mono` to `sprouted/garden`
   - Update go.mod dependencies
   - Fix any breaking references

3. **Rebrand Throughout**
   - `Weather Dashboard` → `Weather Station`
   - Update all UI text and documentation
   - Align with Sprouted visual identity

### Phase 2: Weather System Integration (Week 2)

1. **Connect to Real Data**
   ```go
   // Instead of mock data
   weather := weatherlib.LoadContext(gardenPath)
   
   // Stream real updates
   watcher := weatherlib.WatchForChanges(gardenPath)
   ```

2. **Integrate with Sprout CLI**
   - Add `sprout station` commands
   - Enable `sprout weather --push-to-station`
   - Support station configuration in weather.json

3. **Multi-Garden Support**
   - Scan for all gardens in a farm
   - Aggregate weather across projects
   - Show farm-level patterns

### Phase 3: Cloud Features (Week 3)

1. **Add Backup Functionality**
   - Implement weather.json syncing
   - Add recovery endpoints
   - Create backup scheduling

2. **Team Collaboration**
   - User authentication system
   - Permission-based garden access
   - Shared dashboard views

3. **Enhanced Analytics**
   - Historical weather patterns
   - Productivity insights
   - Team velocity tracking

## API Updates

### New Endpoints

```
POST   /api/backup              - Backup weather data
GET    /api/backup/restore      - Restore from backup
GET    /api/gardens/discover    - Auto-discover local gardens
GET    /api/analytics/personal  - Personal productivity metrics
GET    /api/analytics/team      - Team collaboration metrics
```

### WebSocket Messages

```json
// New message types
{
  "type": "backup_complete",
  "timestamp": "2025-01-24T10:00:00Z",
  "backup_id": "bkp_123"
}

{
  "type": "garden_discovered",
  "garden_id": "new_garden",
  "path": "/path/to/garden"
}
```

## UI/UX Updates

1. **Dashboard Enhancements**
   - Add backup status widget
   - Show sync indicators
   - Multi-garden selector
   - Team presence indicators

2. **New Views**
   - Personal analytics dashboard
   - Team collaboration view
   - Historical weather playback
   - Learning recommendations

3. **Branding Alignment**
   - Sprouted color scheme
   - Weather Service iconography
   - Consistent terminology

## Documentation Updates

### Remove Tentative Language

**Before**: "Weather Station (Coming Soon) - future cloud service if demanded"

**After**: "Weather Station - Your real-time development command center with cloud backup, team dashboards, and intelligent alerts"

### Update Architecture Docs

Show Weather Station as core component, not optional addition:

```
Sprouted Weather Service™
├── Weather System (Local Intelligence)
├── Weather Station (Mission Control) ← We are here!
└── Seed Exchange (Community Wisdom)
```

### Create Launch Materials

1. **Migration Announcement**
   - "Weather Station Lives: From Prototype to Production"
   - Show dashboard screenshots
   - Highlight real features

2. **Getting Started Guide**
   - Installation from existing Weather System
   - First dashboard setup
   - Team onboarding

## Success Metrics

- Weather Station running locally in 1 hour
- Connected to real Weather System in 1 day
- Cloud features operational in 1 week
- Documentation updated to reflect reality

## The Bigger Picture

This migration transforms Sprouted from "building toward a vision" to "delivering on proven solutions." Weather Station's existence validates:

1. The ecosystem architecture works
2. Real-time monitoring is valuable
3. The Weather metaphor resonates
4. We build what we need

## Next Steps

1. Start code migration immediately
2. Update all tentative documentation
3. Plan Weather Station launch announcement
4. Begin early access program

---

*"We didn't just dream it. We built it. Now let's ship it."*
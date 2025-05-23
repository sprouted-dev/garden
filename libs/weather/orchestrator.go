package weather

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// FarmOrchestrator processes events from multiple gardens
type FarmOrchestrator struct {
	FarmPath   string
	EventsPath string
	Gardens    map[string]*GardenInfo
}

// GardenInfo tracks information about a garden
type GardenInfo struct {
	Name         string    `json:"name"`
	LastActivity time.Time `json:"last_activity"`
	EventCount   int       `json:"event_count"`
	Temperature  float64   `json:"temperature"`
}

// FarmWeather represents aggregated weather across all gardens
type FarmWeather struct {
	Timestamp        time.Time              `json:"timestamp"`
	Gardens          map[string]GardenInfo  `json:"gardens"`
	OverallTemp      float64                `json:"overall_temperature"`
	ActiveGardens    []string               `json:"active_gardens"`
	RecentEvents     []Event                `json:"recent_events"`
	Correlations     []EventCorrelation     `json:"correlations"`
	Suggestions      []DocumentationNeed    `json:"suggestions"`
}

// EventCorrelation represents related events across gardens
type EventCorrelation struct {
	Gardens    []string  `json:"gardens"`
	EventTypes []string  `json:"event_types"`
	Pattern    string    `json:"pattern"`
	Confidence float64   `json:"confidence"`
	Timespan   string    `json:"timespan"`
	Suggestion string    `json:"suggestion,omitempty"`
}

// NewFarmOrchestrator creates a new orchestrator
func NewFarmOrchestrator(farmPath string) *FarmOrchestrator {
	return &FarmOrchestrator{
		FarmPath:   farmPath,
		EventsPath: filepath.Join(farmPath, ".farm", "events"),
		Gardens:    make(map[string]*GardenInfo),
	}
}

// ProcessEvents processes pending events from all gardens
func (o *FarmOrchestrator) ProcessEvents() error {
	pendingDir := filepath.Join(o.EventsPath, "pending")
	processedDir := filepath.Join(o.EventsPath, "processed")
	
	// Ensure directories exist
	if err := os.MkdirAll(processedDir, 0755); err != nil {
		return fmt.Errorf("failed to create processed directory: %w", err)
	}
	
	// Read pending events
	files, err := os.ReadDir(pendingDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No events to process
		}
		return err
	}
	
	var events []Event
	fmt.Printf("Found %d files in pending directory\n", len(files))
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		fmt.Printf("Processing event file: %s\n", file.Name())
		
		eventPath := filepath.Join(pendingDir, file.Name())
		data, err := os.ReadFile(eventPath)
		if err != nil {
			continue
		}
		
		var event Event
		if err := json.Unmarshal(data, &event); err != nil {
			fmt.Printf("Error unmarshaling event %s: %v\n", file.Name(), err)
			continue
		}
		
		events = append(events, event)
		
		// Update garden info
		if _, exists := o.Gardens[event.Garden]; !exists {
			o.Gardens[event.Garden] = &GardenInfo{Name: event.Garden}
		}
		o.Gardens[event.Garden].LastActivity = event.Timestamp
		o.Gardens[event.Garden].EventCount++
		o.Gardens[event.Garden].Temperature = event.Context.WeatherTemp
		
		// Move to processed
		processedPath := filepath.Join(processedDir, file.Name())
		if err := os.Rename(eventPath, processedPath); err != nil {
			return fmt.Errorf("failed to move event to processed: %w", err)
		}
	}
	
	// Correlate events
	fmt.Printf("Processing %d events\n", len(events))
	if len(events) > 0 {
		correlations := o.correlateEvents(events)
		farmWeather := o.synthesizeFarmWeather(events, correlations)
		
		// Save farm weather
		fmt.Printf("Saving farm weather to %s\n", o.FarmPath)
		return o.saveFarmWeather(farmWeather)
	}
	
	return nil
}

// correlateEvents finds patterns across garden events
func (o *FarmOrchestrator) correlateEvents(events []Event) []EventCorrelation {
	var correlations []EventCorrelation
	
	// Group events by time window (5 minutes)
	timeWindows := make(map[int64][]Event)
	for _, event := range events {
		window := event.Timestamp.Unix() / 300 // 5-minute windows
		timeWindows[window] = append(timeWindows[window], event)
	}
	
	// Look for patterns in each window
	for _, windowEvents := range timeWindows {
		if len(windowEvents) < 2 {
			continue
		}
		
		// Check for cross-garden activity
		gardens := make(map[string]bool)
		eventTypes := make(map[string]bool)
		
		for _, event := range windowEvents {
			gardens[event.Garden] = true
			eventTypes[string(event.EventType)] = true
		}
		
		if len(gardens) > 1 {
			// Multiple gardens active in same time window
			correlation := EventCorrelation{
				Gardens:    mapKeys(gardens),
				EventTypes: mapKeys(eventTypes),
				Pattern:    "cross-garden-activity",
				Confidence: 0.8,
				Timespan:   "5m",
			}
			
			// Check for specific patterns
			if eventTypes["commit"] && len(gardens) == 2 {
				correlation.Pattern = "coordinated-changes"
				correlation.Suggestion = "Consider documenting the relationship between these changes"
				correlation.Confidence = 0.9
			}
			
			correlations = append(correlations, correlation)
		}
	}
	
	return correlations
}

// synthesizeFarmWeather creates aggregated weather view
func (o *FarmOrchestrator) synthesizeFarmWeather(events []Event, correlations []EventCorrelation) FarmWeather {
	// Calculate overall temperature (activity level)
	var totalTemp float64
	activeGardens := []string{}
	
	for name, info := range o.Gardens {
		if time.Since(info.LastActivity) < 24*time.Hour {
			activeGardens = append(activeGardens, name)
			totalTemp += info.Temperature
		}
	}
	
	overallTemp := totalTemp / float64(len(activeGardens))
	if overallTemp == 0 {
		overallTemp = 72.0 // Default comfortable temperature
	}
	
	// Get recent events (last 10)
	sort.Slice(events, func(i, j int) bool {
		return events[i].Timestamp.After(events[j].Timestamp)
	})
	
	recentEvents := events
	if len(recentEvents) > 10 {
		recentEvents = recentEvents[:10]
	}
	
	// Generate suggestions based on patterns
	suggestions := o.generateSuggestions(events, correlations)
	
	return FarmWeather{
		Timestamp:     time.Now(),
		Gardens:       o.copyGardens(),
		OverallTemp:   overallTemp,
		ActiveGardens: activeGardens,
		RecentEvents:  recentEvents,
		Correlations:  correlations,
		Suggestions:   suggestions,
	}
}

// generateSuggestions creates documentation suggestions from patterns
func (o *FarmOrchestrator) generateSuggestions(events []Event, correlations []EventCorrelation) []DocumentationNeed {
	var suggestions []DocumentationNeed
	
	// Check for undocumented cross-garden changes
	for _, correlation := range correlations {
		if correlation.Pattern == "coordinated-changes" && correlation.Confidence > 0.8 {
			suggestions = append(suggestions, DocumentationNeed{
				Type:        "decision",
				Title:       fmt.Sprintf("Cross-garden coordination: %s", strings.Join(correlation.Gardens, " + ")),
				Description: "Multiple gardens changed in coordinated fashion, suggesting an architectural decision",
				SuggestedLocation: "farm-docs/decisions/",
				Confidence:  correlation.Confidence,
				DetectedFrom: []string{"event correlation", "timing patterns"},
			})
		}
	}
	
	return suggestions
}

// saveFarmWeather persists the farm weather state
func (o *FarmOrchestrator) saveFarmWeather(weather FarmWeather) error {
	weatherDir := filepath.Join(o.FarmPath, ".farm", "weather")
	if err := os.MkdirAll(weatherDir, 0755); err != nil {
		return fmt.Errorf("failed to create weather directory: %w", err)
	}
	
	// Save current weather
	currentPath := filepath.Join(weatherDir, "current.json")
	data, err := json.MarshalIndent(weather, "", "  ")
	if err != nil {
		return err
	}
	
	if err := os.WriteFile(currentPath, data, 0644); err != nil {
		return err
	}
	
	// Archive snapshot
	historyDir := filepath.Join(weatherDir, "history")
	if err := os.MkdirAll(historyDir, 0755); err != nil {
		return fmt.Errorf("failed to create history directory: %w", err)
	}
	
	snapshotPath := filepath.Join(historyDir, fmt.Sprintf("%d.json", time.Now().Unix()))
	return os.WriteFile(snapshotPath, data, 0644)
}

// copyGardens creates a copy of garden info map
func (o *FarmOrchestrator) copyGardens() map[string]GardenInfo {
	result := make(map[string]GardenInfo)
	for k, v := range o.Gardens {
		result[k] = *v
	}
	return result
}

// mapKeys extracts keys from a map
func mapKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
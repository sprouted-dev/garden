package weather

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

// EventType represents the type of event being emitted
type EventType string

const (
	EventCommit        EventType = "commit"
	EventDocumentation EventType = "documentation"
	EventConversation  EventType = "conversation"
	EventDecision      EventType = "decision"
	EventWeatherUpdate EventType = "weather_update"
)

// Event represents a garden-level event to be processed by the farm orchestrator
type Event struct {
	EventID          string                 `json:"event_id"`
	Timestamp        time.Time              `json:"timestamp"`
	Garden           string                 `json:"garden"`
	EventType        EventType              `json:"event_type"`
	Payload          map[string]interface{} `json:"payload"`
	Context          EventContext           `json:"context"`
	CorrelationHints []string               `json:"correlation_hints,omitempty"`
}

// EventContext provides additional context about the event
type EventContext struct {
	Branch      string  `json:"branch,omitempty"`
	WeatherTemp float64 `json:"weather_temp,omitempty"`
	SessionID   string  `json:"session_id,omitempty"`
	UserID      string  `json:"user_id,omitempty"`
}

// EventEmitter handles event creation and emission to the farm orchestrator
type EventEmitter struct {
	gardenName string
	farmPath   string
	weather    *Weather
}

// NewEventEmitter creates a new event emitter for the current garden
func NewEventEmitter(weather *Weather) (*EventEmitter, error) {
	gardenName := filepath.Base(weather.RepoPath)
	farmPath := filepath.Dir(weather.RepoPath)
	
	return &EventEmitter{
		gardenName: gardenName,
		farmPath:   farmPath,
		weather:    weather,
	}, nil
}

// Emit creates and queues an event for farm processing
func (e *EventEmitter) Emit(eventType EventType, payload map[string]interface{}, hints ...string) error {
	event := Event{
		EventID:          uuid.New().String(),
		Timestamp:        time.Now(),
		Garden:           e.gardenName,
		EventType:        eventType,
		Payload:          payload,
		CorrelationHints: hints,
		Context: EventContext{
			Branch:      e.weather.CurrentBranch,
			WeatherTemp: float64(e.weather.Context.Weather.Temperature),
		},
	}
	
	return e.QueueEvent(event)
}

// QueueEvent writes an event to the farm event queue
func (e *EventEmitter) QueueEvent(event Event) error {
	// Ensure farm events directory exists
	eventsDir := filepath.Join(e.farmPath, ".farm", "events", "pending")
	if err := os.MkdirAll(eventsDir, 0755); err != nil {
		// If we can't create farm directory, queue locally
		return e.queueLocalEvent(event)
	}
	
	// Create event file
	filename := fmt.Sprintf("%d-%s-%s.json", 
		time.Now().Unix(), 
		event.EventType, 
		event.EventID[:8])
	
	eventPath := filepath.Join(eventsDir, filename)
	
	// Marshal event
	data, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	
	// Write event file
	if err := os.WriteFile(eventPath, data, 0644); err != nil {
		// Fall back to local queue
		return e.queueLocalEvent(event)
	}
	
	return nil
}

// queueLocalEvent queues events locally when farm is inaccessible
func (e *EventEmitter) queueLocalEvent(event Event) error {
	localQueueDir := filepath.Join(e.weather.RepoPath, ".garden", "event-queue")
	if err := os.MkdirAll(localQueueDir, 0755); err != nil {
		return fmt.Errorf("failed to create local event queue: %w", err)
	}
	
	filename := fmt.Sprintf("%d-%s.json", time.Now().Unix(), event.EventID[:8])
	eventPath := filepath.Join(localQueueDir, filename)
	
	data, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	
	return os.WriteFile(eventPath, data, 0644)
}

// EmitCommitEvent emits a commit event with relevant metadata
func (e *EventEmitter) EmitCommitEvent(hash, message string, filesChanged int) error {
	payload := map[string]interface{}{
		"commit_hash":   hash,
		"message":       message,
		"files_changed": filesChanged,
		"author":        os.Getenv("USER"),
	}
	
	// Extract correlation hints from commit message
	hints := extractCorrelationHints(message)
	
	return e.Emit(EventCommit, payload, hints...)
}

// EmitConversationEvent emits a conversation capture event
func (e *EventEmitter) EmitConversationEvent(conversation ConversationCapture) error {
	payload := map[string]interface{}{
		"user_prompt":       conversation.UserPrompt,
		"conversation_type": conversation.ConversationType,
		"key_insights":      conversation.KeyInsights,
		"should_persist":    conversation.ShouldPersist,
	}
	
	return e.Emit(EventConversation, payload, conversation.RelatedFiles...)
}

// extractCorrelationHints extracts potential correlation hints from text
func extractCorrelationHints(text string) []string {
	var hints []string
	
	// Extract issue numbers (#123)
	// Extract feature names (feature-x, feat/y)
	// Extract cross-references (see garden/, references website/)
	// This is a simplified version - real implementation would be more sophisticated
	
	return hints
}

// ConversationCapture represents a captured conversation
type ConversationCapture struct {
	Timestamp         time.Time `json:"timestamp"`
	UserPrompt        string    `json:"user_prompt"`
	AIResponse        string    `json:"ai_response"`
	ConversationType  string    `json:"conversation_type"`
	KeyInsights       []string  `json:"key_insights"`
	Outcome           string    `json:"outcome"`
	ShouldPersist     bool      `json:"should_persist"`
	RelatedFiles      []string  `json:"related_files"`
	Context           string    `json:"context"`
}
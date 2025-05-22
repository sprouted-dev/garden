package weather

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestContextManager_LoadContext_NewGarden(t *testing.T) {
	// Create temporary directory for testing
	tempDir, err := os.MkdirTemp("", "weather_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cm := NewContextManager(tempDir)
	
	// Load context should create default context
	context, err := cm.LoadContext()
	if err != nil {
		t.Fatalf("Failed to load context: %v", err)
	}
	
	// Verify default context properties
	if context.CurrentFocus.Area != "initializing garden" {
		t.Errorf("Expected focus area 'initializing garden', got '%s'", context.CurrentFocus.Area)
	}
	
	if context.Weather.Condition != WeatherSunny {
		t.Errorf("Expected sunny weather, got %s", context.Weather.Condition)
	}
	
	if context.Version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got '%s'", context.Version)
	}
	
	// Verify context file was created
	contextPath := cm.GetContextPath()
	if _, err := os.Stat(contextPath); os.IsNotExist(err) {
		t.Error("Context file was not created")
	}
}

func TestContextManager_SaveAndLoadContext(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "weather_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cm := NewContextManager(tempDir)
	
	// Create test context
	context := &WeatherContext{
		Updated:    time.Now(),
		SessionID:  "test_session",
		GardenPath: tempDir,
		Version:    "1.0.0",
		CurrentFocus: FocusArea{
			Area:       "authentication system",
			Confidence: 0.95,
			LastActive: time.Now(),
		},
		Weather: WeatherConditions{
			Temperature: 75,
			Condition:   WeatherPartlyCloudy,
			Pressure:    60,
			LastUpdate:  time.Now(),
		},
	}
	
	// Save context
	if err := cm.SaveContext(context); err != nil {
		t.Fatalf("Failed to save context: %v", err)
	}
	
	// Load context
	loadedContext, err := cm.LoadContext()
	if err != nil {
		t.Fatalf("Failed to load context: %v", err)
	}
	
	// Verify loaded context matches saved context
	if loadedContext.SessionID != context.SessionID {
		t.Errorf("Expected session ID '%s', got '%s'", context.SessionID, loadedContext.SessionID)
	}
	
	if loadedContext.CurrentFocus.Area != context.CurrentFocus.Area {
		t.Errorf("Expected focus area '%s', got '%s'", context.CurrentFocus.Area, loadedContext.CurrentFocus.Area)
	}
	
	if loadedContext.Weather.Temperature != context.Weather.Temperature {
		t.Errorf("Expected temperature %d, got %d", context.Weather.Temperature, loadedContext.Weather.Temperature)
	}
}

func TestContextManager_UpdateContext(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "weather_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cm := NewContextManager(tempDir)
	
	// Load initial context
	_, err = cm.LoadContext()
	if err != nil {
		t.Fatalf("Failed to load initial context: %v", err)
	}
	
	// Update context
	newFocusArea := "user interface components"
	err = cm.UpdateContext(func(ctx *WeatherContext) {
		ctx.CurrentFocus.Area = newFocusArea
		ctx.Weather.Temperature = 85
	})
	if err != nil {
		t.Fatalf("Failed to update context: %v", err)
	}
	
	// Verify update was persisted
	updatedContext, err := cm.LoadContext()
	if err != nil {
		t.Fatalf("Failed to load updated context: %v", err)
	}
	
	if updatedContext.CurrentFocus.Area != newFocusArea {
		t.Errorf("Expected focus area '%s', got '%s'", newFocusArea, updatedContext.CurrentFocus.Area)
	}
	
	if updatedContext.Weather.Temperature != 85 {
		t.Errorf("Expected temperature 85, got %d", updatedContext.Weather.Temperature)
	}
}

func TestContextManager_CorruptedFileRecovery(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "weather_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cm := NewContextManager(tempDir)
	contextPath := cm.GetContextPath()
	
	// Create .garden directory
	gardenDir := filepath.Dir(contextPath)
	if err := os.MkdirAll(gardenDir, 0755); err != nil {
		t.Fatalf("Failed to create .garden directory: %v", err)
	}
	
	// Write corrupted JSON to context file
	corruptedJSON := `{"this": "is not valid json`
	if err := os.WriteFile(contextPath, []byte(corruptedJSON), 0644); err != nil {
		t.Fatalf("Failed to write corrupted file: %v", err)
	}
	
	// Load context should recover gracefully
	context, err := cm.LoadContext()
	if err != nil {
		t.Fatalf("Failed to recover from corrupted context: %v", err)
	}
	
	// Should create default context
	if context.CurrentFocus.Area != "initializing garden" {
		t.Errorf("Recovery should create default context, got focus area: %s", context.CurrentFocus.Area)
	}
}

func TestWeatherContext_Validate(t *testing.T) {
	tests := []struct {
		name    string
		context WeatherContext
		wantErr bool
	}{
		{
			name: "valid context",
			context: WeatherContext{
				Version:    "1.0.0",
				GardenPath: "/path/to/garden",
				CurrentFocus: FocusArea{
					Confidence: 0.8,
				},
				Weather: WeatherConditions{
					Temperature: 70,
					Pressure:    50,
				},
			},
			wantErr: false,
		},
		{
			name: "missing version",
			context: WeatherContext{
				GardenPath: "/path/to/garden",
			},
			wantErr: true,
		},
		{
			name: "invalid confidence",
			context: WeatherContext{
				Version:    "1.0.0",
				GardenPath: "/path/to/garden",
				CurrentFocus: FocusArea{
					Confidence: 1.5, // Invalid - should be 0-1
				},
			},
			wantErr: true,
		},
		{
			name: "invalid temperature",
			context: WeatherContext{
				Version:    "1.0.0",
				GardenPath: "/path/to/garden",
				CurrentFocus: FocusArea{
					Confidence: 0.8,
				},
				Weather: WeatherConditions{
					Temperature: 150, // Invalid - should be 0-100
				},
			},
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.context.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("WeatherContext.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeatherContext_ToJSON(t *testing.T) {
	context := WeatherContext{
		Version:    "1.0.0",
		SessionID:  "test_session",
		GardenPath: "/test/path",
		CurrentFocus: FocusArea{
			Area:       "testing system",
			Confidence: 0.9,
		},
	}
	
	jsonData, err := context.ToJSON()
	if err != nil {
		t.Fatalf("Failed to convert to JSON: %v", err)
	}
	
	// Verify it's valid JSON by unmarshaling
	var unmarshaled WeatherContext
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Fatalf("Generated JSON is not valid: %v", err)
	}
	
	// Verify key fields
	if unmarshaled.SessionID != context.SessionID {
		t.Errorf("JSON round-trip failed for SessionID")
	}
	
	if unmarshaled.CurrentFocus.Area != context.CurrentFocus.Area {
		t.Errorf("JSON round-trip failed for CurrentFocus.Area")
	}
}

func TestWeatherContext_ToAIContext(t *testing.T) {
	now := time.Now()
	context := WeatherContext{
		CurrentFocus: FocusArea{
			Area:       "authentication system",
			Confidence: 0.95,
		},
		RecentProgress: ProgressSummary{
			Summary:  "Implemented JWT validation",
			Momentum: 80,
		},
		Git: GitContext{
			CurrentBranch:      "feature/auth",
			UncommittedChanges: true,
			LastCommit: GitCommit{
				Message:      "Add JWT middleware",
				SmartSummary: "Added JWT authentication middleware",
				FilesChanged: []string{"auth/jwt.go", "middleware/auth.go"},
			},
		},
		Weather: WeatherConditions{
			Temperature: 75,
			Condition:   WeatherPartlyCloudy,
			Pressure:    60,
			LastUpdate:  now,
		},
		Updated: now,
	}
	
	aiContext := context.ToAIContext()
	
	// Verify AI context structure
	projectStatus, ok := aiContext["project_status"].(map[string]interface{})
	if !ok {
		t.Fatal("AI context missing project_status")
	}
	
	if projectStatus["current_focus"] != "authentication system" {
		t.Errorf("AI context current_focus incorrect: %v", projectStatus["current_focus"])
	}
	
	devContext, ok := aiContext["development_context"].(map[string]interface{})
	if !ok {
		t.Fatal("AI context missing development_context")
	}
	
	if devContext["current_branch"] != "feature/auth" {
		t.Errorf("AI context current_branch incorrect: %v", devContext["current_branch"])
	}
	
	weatherConditions, ok := aiContext["weather_conditions"].(map[string]interface{})
	if !ok {
		t.Fatal("AI context missing weather_conditions")
	}
	
	if weatherConditions["temperature"] != 75 {
		t.Errorf("AI context temperature incorrect: %v", weatherConditions["temperature"])
	}
}

func TestPerformanceRequirements(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "weather_perf_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cm := NewContextManager(tempDir)
	
	// Test that context operations complete within performance requirements
	
	// Load context should complete in <50ms
	start := time.Now()
	_, err = cm.LoadContext()
	duration := time.Since(start)
	if err != nil {
		t.Fatalf("Failed to load context: %v", err)
	}
	if duration > 50*time.Millisecond {
		t.Errorf("LoadContext took %v, should be <50ms", duration)
	}
	
	// Save context should complete in <50ms
	context := &WeatherContext{
		Version:    "1.0.0",
		GardenPath: tempDir,
		Updated:    time.Now(),
	}
	
	start = time.Now()
	err = cm.SaveContext(context)
	duration = time.Since(start)
	if err != nil {
		t.Fatalf("Failed to save context: %v", err)
	}
	if duration > 50*time.Millisecond {
		t.Errorf("SaveContext took %v, should be <50ms", duration)
	}
}
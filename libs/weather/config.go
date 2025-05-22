package weather

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Configuration for weather system behavior
type Config struct {
	// Git monitoring settings
	GitHooksEnabled    bool          `json:"gitHooksEnabled"`
	CommitLookbackDays int           `json:"commitLookbackDays"`
	MaxCommitsToTrack  int           `json:"maxCommitsToTrack"`
	
	// Context persistence settings
	ContextFileName    string        `json:"contextFileName"`
	BackupRetention    int           `json:"backupRetention"`
	SaveInterval       time.Duration `json:"saveInterval"`
	
	// Inference settings
	ConfidenceThreshold   float64 `json:"confidenceThreshold"`
	MomentumDecayHours    int     `json:"momentumDecayHours"`
	WeatherUpdateInterval int     `json:"weatherUpdateIntervalMinutes"`
	
	// Performance settings
	MaxResponseTimeMs int  `json:"maxResponseTimeMs"`
	EnableDebugMode   bool `json:"enableDebugMode"`
}

// DefaultConfig returns the default weather system configuration
func DefaultConfig() Config {
	return Config{
		GitHooksEnabled:       true,
		CommitLookbackDays:    7,
		MaxCommitsToTrack:     10,
		ContextFileName:       "weather-context.json",
		BackupRetention:       5,
		SaveInterval:          time.Minute * 5,
		ConfidenceThreshold:   0.7,
		MomentumDecayHours:    2,
		WeatherUpdateInterval: 30,
		MaxResponseTimeMs:     200,
		EnableDebugMode:       false,
	}
}

// ConfigManager handles loading and saving configuration
type ConfigManager struct {
	gardenPath string
	config     *Config
}

// NewConfigManager creates a new configuration manager
func NewConfigManager(gardenPath string) *ConfigManager {
	return &ConfigManager{
		gardenPath: gardenPath,
		config:     nil,
	}
}

// LoadConfig loads configuration from disk or creates default
func (cm *ConfigManager) LoadConfig() (*Config, error) {
	if cm.config != nil {
		return cm.config, nil
	}
	
	configPath := cm.getConfigPath()
	
	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create default config
		defaultConfig := DefaultConfig()
		if err := cm.SaveConfig(&defaultConfig); err != nil {
			return nil, fmt.Errorf("failed to create default config: %w", err)
		}
		cm.config = &defaultConfig
		return &defaultConfig, nil
	}
	
	// Read existing config
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		// Config corrupted, use defaults
		fmt.Printf("Warning: corrupted weather config, using defaults\n")
		defaultConfig := DefaultConfig()
		cm.config = &defaultConfig
		return &defaultConfig, nil
	}
	
	cm.config = &config
	return &config, nil
}

// SaveConfig saves configuration to disk
func (cm *ConfigManager) SaveConfig(config *Config) error {
	configPath := cm.getConfigPath()
	
	// Ensure .garden directory exists
	gardenDir := filepath.Dir(configPath)
	if err := os.MkdirAll(gardenDir, 0755); err != nil {
		return fmt.Errorf("failed to create .garden directory: %w", err)
	}
	
	// Marshal to JSON with indentation
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	
	// Atomic write
	tempPath := configPath + ".tmp"
	if err := os.WriteFile(tempPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write temporary config file: %w", err)
	}
	
	if err := os.Rename(tempPath, configPath); err != nil {
		os.Remove(tempPath)
		return fmt.Errorf("failed to rename temporary config file: %w", err)
	}
	
	cm.config = config
	return nil
}

// UpdateConfig updates configuration with validation
func (cm *ConfigManager) UpdateConfig(updater func(*Config)) error {
	config, err := cm.LoadConfig()
	if err != nil {
		return err
	}
	
	// Create a copy to avoid modifying the original if validation fails
	updatedConfig := *config
	updater(&updatedConfig)
	
	// Validate updated config
	if err := cm.validateConfig(&updatedConfig); err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}
	
	return cm.SaveConfig(&updatedConfig)
}

// getConfigPath returns the path to the config file
func (cm *ConfigManager) getConfigPath() string {
	return filepath.Join(cm.gardenPath, ".garden", "weather-config.json")
}

// validateConfig validates configuration values
func (cm *ConfigManager) validateConfig(config *Config) error {
	if config.CommitLookbackDays < 1 || config.CommitLookbackDays > 365 {
		return fmt.Errorf("commitLookbackDays must be between 1 and 365")
	}
	
	if config.MaxCommitsToTrack < 1 || config.MaxCommitsToTrack > 100 {
		return fmt.Errorf("maxCommitsToTrack must be between 1 and 100")
	}
	
	if config.ConfidenceThreshold < 0 || config.ConfidenceThreshold > 1 {
		return fmt.Errorf("confidenceThreshold must be between 0 and 1")
	}
	
	if config.MomentumDecayHours < 1 || config.MomentumDecayHours > 168 { // 1 week max
		return fmt.Errorf("momentumDecayHours must be between 1 and 168")
	}
	
	if config.MaxResponseTimeMs < 50 || config.MaxResponseTimeMs > 5000 {
		return fmt.Errorf("maxResponseTimeMs must be between 50 and 5000")
	}
	
	return nil
}
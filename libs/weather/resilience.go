// Package weather provides automatic context preservation and intelligence
// resilience.go implements disaster recovery and backup mechanisms
package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// ResilienceConfig defines backup and recovery settings
type ResilienceConfig struct {
	EnableShadowCopies  bool          `json:"enable_shadow_copies"`
	ShadowCopyCount     int           `json:"shadow_copy_count"`
	BackupInterval      time.Duration `json:"backup_interval"`
	BackupRetention     time.Duration `json:"backup_retention"`
	EnableJournal       bool          `json:"enable_journal"`
	RecoveryStrategies  []string      `json:"recovery_strategies"`
}

// DefaultResilienceConfig returns sensible defaults
func DefaultResilienceConfig() ResilienceConfig {
	return ResilienceConfig{
		EnableShadowCopies:  true,
		ShadowCopyCount:     3,
		BackupInterval:      1 * time.Hour,
		BackupRetention:     7 * 24 * time.Hour, // 7 days
		EnableJournal:       true,
		RecoveryStrategies:  []string{"shadow", "journal", "git-reconstruct"},
	}
}

// ResilienceManager handles backup and recovery operations
type ResilienceManager struct {
	config     ResilienceConfig
	gardenPath string
}

// NewResilienceManager creates a new resilience manager
func NewResilienceManager(gardenPath string, config ResilienceConfig) *ResilienceManager {
	return &ResilienceManager{
		config:     config,
		gardenPath: gardenPath,
	}
}

// CreateBackup creates a full backup of weather data
func (rm *ResilienceManager) CreateBackup() error {
	timestamp := time.Now().Format("20060102-150405")
	backupDir := filepath.Join(rm.gardenPath, ".garden", "backups", timestamp)
	
	// Create backup directory
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return fmt.Errorf("failed to create backup directory: %w", err)
	}
	
	// Files to backup
	filesToBackup := []string{
		filepath.Join(rm.gardenPath, ".garden", "weather-context.json"),
		filepath.Join(rm.gardenPath, "weather.md"),
	}
	
	// Backup each file
	for _, src := range filesToBackup {
		if _, err := os.Stat(src); err == nil {
			dst := filepath.Join(backupDir, filepath.Base(src))
			if err := copyFile(src, dst); err != nil {
				return fmt.Errorf("failed to backup %s: %w", src, err)
			}
		}
	}
	
	// Create backup metadata
	metadata := BackupMetadata{
		Timestamp:   time.Now(),
		Version:     "1.0.0",
		GardenPath:  rm.gardenPath,
		Files:       filesToBackup,
		BackupType:  "full",
		Checksum:    rm.calculateChecksum(),
	}
	
	metadataPath := filepath.Join(backupDir, "backup-metadata.json")
	if err := rm.saveJSON(metadataPath, metadata); err != nil {
		return fmt.Errorf("failed to save backup metadata: %w", err)
	}
	
	// Clean old backups
	if err := rm.cleanOldBackups(); err != nil {
		// Log but don't fail
		fmt.Printf("Warning: failed to clean old backups: %v\n", err)
	}
	
	return nil
}

// CreateShadowCopy creates a shadow copy of critical files
func (rm *ResilienceManager) CreateShadowCopy() error {
	if !rm.config.EnableShadowCopies {
		return nil
	}
	
	shadowDir := filepath.Join(rm.gardenPath, ".garden", "shadows")
	if err := os.MkdirAll(shadowDir, 0755); err != nil {
		return fmt.Errorf("failed to create shadow directory: %w", err)
	}
	
	// Rotate existing shadows
	rm.rotateShadows(shadowDir)
	
	// Create new shadow copies
	weatherContext := filepath.Join(rm.gardenPath, ".garden", "weather-context.json")
	if _, err := os.Stat(weatherContext); err == nil {
		shadowPath := filepath.Join(shadowDir, "weather-context.shadow.json")
		if err := copyFile(weatherContext, shadowPath); err != nil {
			return fmt.Errorf("failed to create shadow copy: %w", err)
		}
	}
	
	// Shadow copy of weather.md
	weatherMd := filepath.Join(rm.gardenPath, "weather.md")
	if _, err := os.Stat(weatherMd); err == nil {
		shadowPath := filepath.Join(shadowDir, "weather.shadow.md")
		if err := copyFile(weatherMd, shadowPath); err != nil {
			return fmt.Errorf("failed to create weather.md shadow: %w", err)
		}
	}
	
	return nil
}

// RecoverFromDisaster attempts to recover weather data
func (rm *ResilienceManager) RecoverFromDisaster() (*RecoveryResult, error) {
	result := &RecoveryResult{
		Timestamp: time.Now(),
		Attempts:  []RecoveryAttempt{},
	}
	
	// Try each recovery strategy in order
	for _, strategy := range rm.config.RecoveryStrategies {
		attempt := RecoveryAttempt{
			Strategy:  strategy,
			Timestamp: time.Now(),
		}
		
		var err error
		switch strategy {
		case "shadow":
			err = rm.recoverFromShadow()
		case "journal":
			err = rm.recoverFromJournal()
		case "git-reconstruct":
			err = rm.recoverFromGit()
		case "backup":
			err = rm.recoverFromBackup()
		default:
			err = fmt.Errorf("unknown recovery strategy: %s", strategy)
		}
		
		if err == nil {
			attempt.Success = true
			attempt.Message = fmt.Sprintf("Successfully recovered using %s strategy", strategy)
			result.Success = true
			result.Attempts = append(result.Attempts, attempt)
			break
		} else {
			attempt.Success = false
			attempt.Error = err.Error()
			result.Attempts = append(result.Attempts, attempt)
		}
	}
	
	return result, nil
}

// recoverFromShadow attempts recovery from shadow copies
func (rm *ResilienceManager) recoverFromShadow() error {
	shadowDir := filepath.Join(rm.gardenPath, ".garden", "shadows")
	
	// Recover weather-context.json
	shadowContext := filepath.Join(shadowDir, "weather-context.shadow.json")
	if _, err := os.Stat(shadowContext); err == nil {
		targetPath := filepath.Join(rm.gardenPath, ".garden", "weather-context.json")
		if err := copyFile(shadowContext, targetPath); err != nil {
			return fmt.Errorf("failed to restore from shadow: %w", err)
		}
	} else {
		return fmt.Errorf("no shadow copy found")
	}
	
	// Recover weather.md if needed
	shadowMd := filepath.Join(shadowDir, "weather.shadow.md")
	if _, err := os.Stat(shadowMd); err == nil {
		targetPath := filepath.Join(rm.gardenPath, "weather.md")
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			if err := copyFile(shadowMd, targetPath); err != nil {
				return fmt.Errorf("failed to restore weather.md: %w", err)
			}
		}
	}
	
	return nil
}

// recoverFromJournal attempts recovery from event journal
func (rm *ResilienceManager) recoverFromJournal() error {
	journalPath := filepath.Join(rm.gardenPath, ".garden", "journal", "weather-events.jsonl")
	if _, err := os.Stat(journalPath); os.IsNotExist(err) {
		return fmt.Errorf("no journal found")
	}
	
	// Read journal and reconstruct state
	// This would replay events to rebuild context
	return fmt.Errorf("journal recovery not yet implemented")
}

// recoverFromGit attempts to reconstruct from git history
func (rm *ResilienceManager) recoverFromGit() error {
	// Create a new Weather instance
	w := &Weather{
		RepoPath: rm.gardenPath,
	}
	
	// Get git context
	gc, err := w.GetGitContext()
	if err != nil {
		return fmt.Errorf("failed to get git context: %w", err)
	}
	
	// Create basic weather context from git
	context := WeatherContext{
		Updated:    time.Now(),
		Version:    "1.0.0",
		GardenPath: rm.gardenPath,
		CurrentFocus: FocusContext{
			Area:       "unknown",
			Confidence: 0.5,
			LastActive: time.Now(),
		},
	}
	
	// Infer from recent commits
	if len(gc.RecentCommits) > 0 {
		recentCommit := gc.RecentCommits[0]
		context.CurrentFocus.InferredFrom = "git history recovery"
		
		// Simple inference from commit message
		if strings.Contains(strings.ToLower(recentCommit.Message), "doc") {
			context.CurrentFocus.Area = "documentation"
		} else if strings.Contains(strings.ToLower(recentCommit.Message), "test") {
			context.CurrentFocus.Area = "testing"
		} else {
			context.CurrentFocus.Area = "development"
		}
	}
	
	// Save reconstructed context
	contextPath := filepath.Join(rm.gardenPath, ".garden", "weather-context.json")
	if err := rm.saveJSON(contextPath, context); err != nil {
		return fmt.Errorf("failed to save reconstructed context: %w", err)
	}
	
	// Generate new weather.md
	w.context = &context
	summary, err := w.GenerateWeatherSummary("# Weather Report\n\nReconstructed from git history due to data loss.")
	if err != nil {
		return fmt.Errorf("failed to generate weather summary: %w", err)
	}
	
	weatherPath := filepath.Join(rm.gardenPath, "weather.md")
	if err := os.WriteFile(weatherPath, []byte(summary), 0644); err != nil {
		return fmt.Errorf("failed to write weather.md: %w", err)
	}
	
	return nil
}

// recoverFromBackup attempts recovery from full backup
func (rm *ResilienceManager) recoverFromBackup() error {
	backupsDir := filepath.Join(rm.gardenPath, ".garden", "backups")
	
	// Find most recent backup
	entries, err := os.ReadDir(backupsDir)
	if err != nil {
		return fmt.Errorf("failed to read backups directory: %w", err)
	}
	
	if len(entries) == 0 {
		return fmt.Errorf("no backups found")
	}
	
	// Sort to get most recent
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() > entries[j].Name()
	})
	
	// Restore from most recent backup
	backupDir := filepath.Join(backupsDir, entries[0].Name())
	
	// Restore weather-context.json
	backupContext := filepath.Join(backupDir, "weather-context.json")
	if _, err := os.Stat(backupContext); err == nil {
		targetPath := filepath.Join(rm.gardenPath, ".garden", "weather-context.json")
		if err := copyFile(backupContext, targetPath); err != nil {
			return fmt.Errorf("failed to restore context from backup: %w", err)
		}
	}
	
	// Restore weather.md
	backupMd := filepath.Join(backupDir, "weather.md")
	if _, err := os.Stat(backupMd); err == nil {
		targetPath := filepath.Join(rm.gardenPath, "weather.md")
		if err := copyFile(backupMd, targetPath); err != nil {
			return fmt.Errorf("failed to restore weather.md from backup: %w", err)
		}
	}
	
	return nil
}

// rotateShadows manages shadow copy rotation
func (rm *ResilienceManager) rotateShadows(shadowDir string) {
	// Rotate numbered shadows
	for i := rm.config.ShadowCopyCount - 1; i > 0; i-- {
		oldPath := filepath.Join(shadowDir, fmt.Sprintf("weather-context.shadow.%d.json", i-1))
		newPath := filepath.Join(shadowDir, fmt.Sprintf("weather-context.shadow.%d.json", i))
		
		if _, err := os.Stat(oldPath); err == nil {
			os.Rename(oldPath, newPath)
		}
	}
	
	// Move current shadow to .0
	currentShadow := filepath.Join(shadowDir, "weather-context.shadow.json")
	if _, err := os.Stat(currentShadow); err == nil {
		rotatedPath := filepath.Join(shadowDir, "weather-context.shadow.0.json")
		os.Rename(currentShadow, rotatedPath)
	}
}

// cleanOldBackups removes backups older than retention period
func (rm *ResilienceManager) cleanOldBackups() error {
	backupsDir := filepath.Join(rm.gardenPath, ".garden", "backups")
	cutoffTime := time.Now().Add(-rm.config.BackupRetention)
	
	entries, err := os.ReadDir(backupsDir)
	if err != nil {
		return err
	}
	
	for _, entry := range entries {
		if entry.IsDir() {
			// Parse timestamp from directory name
			timestamp, err := time.Parse("20060102-150405", entry.Name())
			if err == nil && timestamp.Before(cutoffTime) {
				backupPath := filepath.Join(backupsDir, entry.Name())
				os.RemoveAll(backupPath)
			}
		}
	}
	
	return nil
}

// calculateChecksum calculates a checksum of critical files
func (rm *ResilienceManager) calculateChecksum() string {
	// Simple implementation - in production would use proper checksums
	return fmt.Sprintf("%d", time.Now().Unix())
}

// saveJSON saves data as JSON
func (rm *ResilienceManager) saveJSON(path string, data interface{}) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()
	
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	
	_, err = io.Copy(destination, source)
	return err
}

// BackupMetadata stores backup information
type BackupMetadata struct {
	Timestamp  time.Time `json:"timestamp"`
	Version    string    `json:"version"`
	GardenPath string    `json:"garden_path"`
	Files      []string  `json:"files"`
	BackupType string    `json:"backup_type"`
	Checksum   string    `json:"checksum"`
}

// RecoveryResult stores recovery attempt results
type RecoveryResult struct {
	Success   bool               `json:"success"`
	Timestamp time.Time          `json:"timestamp"`
	Attempts  []RecoveryAttempt  `json:"attempts"`
}

// RecoveryAttempt stores individual recovery attempt
type RecoveryAttempt struct {
	Strategy  string    `json:"strategy"`
	Success   bool      `json:"success"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message,omitempty"`
	Error     string    `json:"error,omitempty"`
}
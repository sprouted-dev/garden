package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// ShadowCopyManager handles backup copies of the weather context
type ShadowCopyManager struct {
	gardenPath string
}

// NewShadowCopyManager creates a new shadow copy manager
func NewShadowCopyManager(gardenPath string) *ShadowCopyManager {
	return &ShadowCopyManager{
		gardenPath: gardenPath,
	}
}

// CreateShadowCopy creates a backup of the current weather context
func (s *ShadowCopyManager) CreateShadowCopy() error {
	contextPath := filepath.Join(s.gardenPath, ".garden", "weather-context.json")
	shadowPath := filepath.Join(s.gardenPath, ".garden", "weather-context.shadow.json")
	
	// Read current context
	src, err := os.Open(contextPath)
	if err != nil {
		return fmt.Errorf("failed to open context: %w", err)
	}
	defer src.Close()
	
	// Create shadow copy
	dst, err := os.Create(shadowPath)
	if err != nil {
		return fmt.Errorf("failed to create shadow copy: %w", err)
	}
	defer dst.Close()
	
	// Copy content
	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("failed to copy context: %w", err)
	}
	
	// Also create a timestamped backup
	timestamp := time.Now().Format("20060102-150405")
	backupDir := filepath.Join(s.gardenPath, ".garden", "backups")
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return fmt.Errorf("failed to create backup dir: %w", err)
	}
	
	backupPath := filepath.Join(backupDir, fmt.Sprintf("weather-context-%s.json", timestamp))
	if _, err := src.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to reset file pointer: %w", err)
	}
	
	backup, err := os.Create(backupPath)
	if err != nil {
		return fmt.Errorf("failed to create timestamped backup: %w", err)
	}
	defer backup.Close()
	
	if _, err := io.Copy(backup, src); err != nil {
		return fmt.Errorf("failed to copy to backup: %w", err)
	}
	
	// Clean old backups (keep last 10)
	if err := s.cleanOldBackups(); err != nil {
		// Non-critical error, just log it
		fmt.Printf("Warning: failed to clean old backups: %v\n", err)
	}
	
	return nil
}

// RestoreFromShadow restores context from shadow copy
func (s *ShadowCopyManager) RestoreFromShadow() error {
	shadowPath := filepath.Join(s.gardenPath, ".garden", "weather-context.shadow.json")
	contextPath := filepath.Join(s.gardenPath, ".garden", "weather-context.json")
	
	// Verify shadow copy exists and is valid
	if err := s.verifyShadowCopy(); err != nil {
		return fmt.Errorf("shadow copy verification failed: %w", err)
	}
	
	// Create backup of current (possibly corrupted) context
	corruptPath := filepath.Join(s.gardenPath, ".garden", "weather-context.corrupt.json")
	if err := os.Rename(contextPath, corruptPath); err != nil {
		// If rename fails, try to remove the corrupt file
		os.Remove(contextPath)
	}
	
	// Copy shadow to main
	src, err := os.Open(shadowPath)
	if err != nil {
		return fmt.Errorf("failed to open shadow copy: %w", err)
	}
	defer src.Close()
	
	dst, err := os.Create(contextPath)
	if err != nil {
		return fmt.Errorf("failed to create restored context: %w", err)
	}
	defer dst.Close()
	
	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("failed to restore from shadow: %w", err)
	}
	
	return nil
}

// VerifyContextIntegrity checks if the weather context is valid
func (s *ShadowCopyManager) VerifyContextIntegrity() error {
	contextPath := filepath.Join(s.gardenPath, ".garden", "weather-context.json")
	
	data, err := os.ReadFile(contextPath)
	if err != nil {
		return fmt.Errorf("failed to read context: %w", err)
	}
	
	var context WeatherContext
	if err := json.Unmarshal(data, &context); err != nil {
		return fmt.Errorf("invalid JSON structure: %w", err)
	}
	
	// Basic integrity checks
	if context.Version == "" {
		return fmt.Errorf("missing version field")
	}
	
	if context.SessionID == "" {
		return fmt.Errorf("missing session ID")
	}
	
	if context.Updated.IsZero() {
		return fmt.Errorf("invalid update timestamp")
	}
	
	return nil
}

// verifyShadowCopy checks if shadow copy is valid
func (s *ShadowCopyManager) verifyShadowCopy() error {
	shadowPath := filepath.Join(s.gardenPath, ".garden", "weather-context.shadow.json")
	
	data, err := os.ReadFile(shadowPath)
	if err != nil {
		return fmt.Errorf("failed to read shadow copy: %w", err)
	}
	
	var context WeatherContext
	if err := json.Unmarshal(data, &context); err != nil {
		return fmt.Errorf("invalid shadow copy JSON: %w", err)
	}
	
	return nil
}

// cleanOldBackups removes old backup files, keeping only the most recent 10
func (s *ShadowCopyManager) cleanOldBackups() error {
	backupDir := filepath.Join(s.gardenPath, ".garden", "backups")
	
	entries, err := os.ReadDir(backupDir)
	if err != nil {
		return err
	}
	
	// Filter weather context backups
	var backups []os.DirEntry
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			backups = append(backups, entry)
		}
	}
	
	// Keep only last 10
	if len(backups) > 10 {
		// Remove oldest backups
		for i := 0; i < len(backups)-10; i++ {
			oldBackup := filepath.Join(backupDir, backups[i].Name())
			if err := os.Remove(oldBackup); err != nil {
				return err
			}
		}
	}
	
	return nil
}

// GetBackupList returns list of available backups
func (s *ShadowCopyManager) GetBackupList() ([]string, error) {
	backupDir := filepath.Join(s.gardenPath, ".garden", "backups")
	
	// Ensure backup directory exists
	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		return []string{}, nil
	}
	
	entries, err := os.ReadDir(backupDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read backup directory: %w", err)
	}
	
	var backups []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			backups = append(backups, entry.Name())
		}
	}
	
	return backups, nil
}

// RestoreFromBackup restores from a specific backup file
func (s *ShadowCopyManager) RestoreFromBackup(backupName string) error {
	backupPath := filepath.Join(s.gardenPath, ".garden", "backups", backupName)
	contextPath := filepath.Join(s.gardenPath, ".garden", "weather-context.json")
	
	// Verify backup exists
	if _, err := os.Stat(backupPath); os.IsNotExist(err) {
		return fmt.Errorf("backup file not found: %s", backupName)
	}
	
	// Verify backup is valid JSON
	data, err := os.ReadFile(backupPath)
	if err != nil {
		return fmt.Errorf("failed to read backup: %w", err)
	}
	
	var context WeatherContext
	if err := json.Unmarshal(data, &context); err != nil {
		return fmt.Errorf("invalid backup JSON: %w", err)
	}
	
	// Create current backup before restore
	if err := s.CreateShadowCopy(); err != nil {
		// Non-critical, continue with restore
		fmt.Printf("Warning: failed to create pre-restore backup: %v\n", err)
	}
	
	// Restore from backup
	if err := os.WriteFile(contextPath, data, 0644); err != nil {
		return fmt.Errorf("failed to restore from backup: %w", err)
	}
	
	return nil
}
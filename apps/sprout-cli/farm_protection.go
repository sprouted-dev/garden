package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// FarmProtectionManager handles Farm-level backup and recovery
type FarmProtectionManager struct {
	farmPath string
}

// NewFarmProtectionManager creates a new manager
func NewFarmProtectionManager() (*FarmProtectionManager, error) {
	// Find farm path by looking for parent of .garden
	gardenPath, err := findGardenPath()
	if err != nil {
		return nil, fmt.Errorf("not in a garden: %w", err)
	}
	
	farmPath := filepath.Dir(gardenPath)
	return &FarmProtectionManager{farmPath: farmPath}, nil
}

// BackupFarm creates a Farm-level backup
func (fpm *FarmProtectionManager) BackupFarm() error {
	backupScript := filepath.Join(fpm.farmPath, ".farm-protection", "backup-farm.sh")
	
	// Check if protection is set up
	if _, err := os.Stat(backupScript); os.IsNotExist(err) {
		return fmt.Errorf("farm protection not set up. Run: setup-farm-protection.sh")
	}
	
	// Run backup script
	cmd := exec.Command(backupScript)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("backup failed: %w\n%s", err, output)
	}
	
	fmt.Print(string(output))
	return nil
}

// CheckFarmProtection verifies Farm-level protection status
func (fpm *FarmProtectionManager) CheckFarmProtection() {
	fmt.Println("🚜 Farm-Level Protection Status")
	fmt.Println()
	
	// Check for protection directory
	protectionDir := filepath.Join(fpm.farmPath, ".farm-protection")
	if _, err := os.Stat(protectionDir); os.IsNotExist(err) {
		fmt.Println("❌ Farm protection NOT set up")
		fmt.Println()
		fmt.Println("⚠️  WARNING: Your Farm-level files are at risk!")
		fmt.Println("   This includes:")
		fmt.Println("   • Private business documentation")
		fmt.Println("   • Cross-repo coordination files")
		fmt.Println("   • Non-git-tracked critical files")
		fmt.Println()
		fmt.Println("💡 To set up protection:")
		fmt.Println("   ./tools/setup-farm-protection.sh")
		return
	}
	
	fmt.Println("✅ Farm protection is set up")
	
	// Check for backups
	backupDir := filepath.Join(protectionDir, "backups")
	entries, err := os.ReadDir(backupDir)
	if err == nil {
		backupCount := 0
		var latestBackup time.Time
		
		for _, entry := range entries {
			if !entry.IsDir() && filepath.Ext(entry.Name()) == ".gz" {
				backupCount++
				if info, err := entry.Info(); err == nil {
					if info.ModTime().After(latestBackup) {
						latestBackup = info.ModTime()
					}
				}
			}
		}
		
		fmt.Printf("📦 Backups available: %d\n", backupCount)
		if !latestBackup.IsZero() {
			fmt.Printf("🕒 Latest backup: %s\n", latestBackup.Format("2006-01-02 15:04:05"))
			
			// Warn if backup is old
			if time.Since(latestBackup) > 24*time.Hour {
				fmt.Println("⚠️  Latest backup is over 24 hours old!")
			}
		}
	}
	
	// Check what's protected
	fmt.Println()
	fmt.Println("📁 Protected directories:")
	protectedDirs := []string{"docs", ".farm"}
	for _, dir := range protectedDirs {
		fullPath := filepath.Join(fpm.farmPath, dir)
		if _, err := os.Stat(fullPath); err == nil {
			fmt.Printf("   ✅ %s/\n", dir)
		} else {
			fmt.Printf("   ⭕ %s/ (not found)\n", dir)
		}
	}
	
	fmt.Println()
	fmt.Println("💡 To create a backup now:")
	fmt.Println("   sprout farm backup")
}
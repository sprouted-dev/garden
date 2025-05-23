package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sprouted/garden/libs/weather"
)

// cmdBackup creates a backup of weather data
func cmdBackup(args []string) error {
	gardenPath, err := findGardenRoot()
	if err != nil {
		return err
	}
	
	config := weather.DefaultResilienceConfig()
	rm := weather.NewResilienceManager(gardenPath, config)
	
	fmt.Println("🔒 Creating backup of Weather System data...")
	
	// Create shadow copies first
	if err := rm.CreateShadowCopy(); err != nil {
		fmt.Printf("⚠️  Warning: Failed to create shadow copies: %v\n", err)
	} else {
		fmt.Println("✅ Shadow copies created")
	}
	
	// Create full backup
	if err := rm.CreateBackup(); err != nil {
		return fmt.Errorf("backup failed: %w", err)
	}
	
	fmt.Println("✅ Backup completed successfully")
	fmt.Println("📁 Backup location: .garden/backups/")
	
	return nil
}

// cmdRecover attempts to recover from data loss
func cmdRecover(args []string) error {
	gardenPath, err := findGardenRoot()
	if err != nil {
		return err
	}
	
	// Check what's missing
	weatherContext := filepath.Join(gardenPath, ".garden", "weather-context.json")
	weatherMd := filepath.Join(gardenPath, "weather.md")
	
	contextMissing := false
	mdMissing := false
	
	if _, err := os.Stat(weatherContext); os.IsNotExist(err) {
		contextMissing = true
		fmt.Println("❌ weather-context.json is missing")
	}
	
	if _, err := os.Stat(weatherMd); os.IsNotExist(err) {
		mdMissing = true
		fmt.Println("❌ weather.md is missing")
	}
	
	if !contextMissing && !mdMissing {
		fmt.Println("✅ No recovery needed - all files present")
		return nil
	}
	
	fmt.Println("\n🚨 Disaster Recovery Mode")
	fmt.Println("========================")
	
	config := weather.DefaultResilienceConfig()
	rm := weather.NewResilienceManager(gardenPath, config)
	
	result, err := rm.RecoverFromDisaster()
	if err != nil {
		return fmt.Errorf("recovery failed: %w", err)
	}
	
	// Display recovery results
	fmt.Println("\n📊 Recovery Results:")
	for _, attempt := range result.Attempts {
		if attempt.Success {
			fmt.Printf("✅ %s: %s\n", attempt.Strategy, attempt.Message)
		} else {
			fmt.Printf("❌ %s: %s\n", attempt.Strategy, attempt.Error)
		}
	}
	
	if result.Success {
		fmt.Println("\n✅ Recovery successful!")
		fmt.Println("🌤️  Weather System data has been restored")
		
		// Run weather to verify
		fmt.Println("\nVerifying recovery...")
		return cmdWeather([]string{})
	} else {
		fmt.Println("\n❌ Recovery failed")
		fmt.Println("💡 Try running 'sprout weather init' to reinitialize")
		return fmt.Errorf("could not recover weather data")
	}
}

// cmdResilienceStatus shows resilience system status
func cmdResilienceStatus(args []string) error {
	gardenPath, err := findGardenRoot()
	if err != nil {
		return err
	}
	
	fmt.Println("🛡️  Resilience System Status")
	fmt.Println("===========================")
	
	// Check shadow copies
	shadowDir := filepath.Join(gardenPath, ".garden", "shadows")
	if entries, err := os.ReadDir(shadowDir); err == nil {
		fmt.Printf("📋 Shadow Copies: %d found\n", len(entries))
		for _, entry := range entries {
			if !entry.IsDir() {
				info, _ := entry.Info()
				fmt.Printf("   - %s (%.1f KB)\n", entry.Name(), float64(info.Size())/1024)
			}
		}
	} else {
		fmt.Println("📋 Shadow Copies: None")
	}
	
	// Check backups
	backupsDir := filepath.Join(gardenPath, ".garden", "backups")
	if entries, err := os.ReadDir(backupsDir); err == nil {
		fmt.Printf("\n💾 Backups: %d found\n", len(entries))
		count := 0
		for _, entry := range entries {
			if entry.IsDir() && count < 5 {
				fmt.Printf("   - %s\n", entry.Name())
				count++
			}
		}
		if len(entries) > 5 {
			fmt.Printf("   ... and %d more\n", len(entries)-5)
		}
	} else {
		fmt.Println("\n💾 Backups: None")
	}
	
	// Check current file status
	fmt.Println("\n📄 Current Files:")
	files := []struct {
		name string
		path string
	}{
		{"weather-context.json", filepath.Join(gardenPath, ".garden", "weather-context.json")},
		{"weather.md", filepath.Join(gardenPath, "weather.md")},
	}
	
	for _, file := range files {
		if info, err := os.Stat(file.path); err == nil {
			fmt.Printf("   ✅ %s (%.1f KB, modified %s)\n", 
				file.name, 
				float64(info.Size())/1024,
				info.ModTime().Format("2006-01-02 15:04"))
		} else {
			fmt.Printf("   ❌ %s (missing)\n", file.name)
		}
	}
	
	// Recovery strategies
	fmt.Println("\n🔄 Recovery Strategies (in order):")
	config := weather.DefaultResilienceConfig()
	for i, strategy := range config.RecoveryStrategies {
		fmt.Printf("   %d. %s\n", i+1, strategy)
	}
	
	return nil
}

// cmdProtect enables automatic protection
func cmdProtect(args []string) error {
	gardenPath, err := findGardenRoot()
	if err != nil {
		return err
	}
	
	fmt.Println("🛡️  Setting up Weather System Protection")
	fmt.Println("======================================")
	
	// Create git hook for automatic shadow copies
	hookPath := filepath.Join(gardenPath, ".git", "hooks", "post-commit")
	hookContent := `#!/bin/bash
# Weather System Protection Hook
# Creates shadow copies after each commit

# Check if sprout is available
if command -v sprout &> /dev/null; then
    sprout backup --shadow-only &
fi
`
	
	// Check if hook already exists
	if _, err := os.Stat(hookPath); err == nil {
		fmt.Println("⚠️  Git hook already exists, skipping...")
	} else {
		if err := os.WriteFile(hookPath, []byte(hookContent), 0755); err != nil {
			return fmt.Errorf("failed to create git hook: %w", err)
		}
		fmt.Println("✅ Git hook installed for automatic shadow copies")
	}
	
	// Create initial backup
	config := weather.DefaultResilienceConfig()
	rm := weather.NewResilienceManager(gardenPath, config)
	
	if err := rm.CreateShadowCopy(); err != nil {
		fmt.Printf("⚠️  Warning: Failed to create initial shadow copies: %v\n", err)
	} else {
		fmt.Println("✅ Initial shadow copies created")
	}
	
	if err := rm.CreateBackup(); err != nil {
		fmt.Printf("⚠️  Warning: Failed to create initial backup: %v\n", err)
	} else {
		fmt.Println("✅ Initial backup created")
	}
	
	fmt.Println("\n🎯 Protection Enabled:")
	fmt.Println("- Shadow copies on every commit")
	fmt.Println("- Manual backups with 'sprout backup'")
	fmt.Println("- Disaster recovery with 'sprout recover'")
	
	return nil
}
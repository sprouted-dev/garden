package weather

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestGitMonitor_ParseGitLog(t *testing.T) {
	gm := NewGitMonitor("/tmp")
	
	// Mock git log output
	gitLogOutput := `abc123|feat: add user authentication|John Doe|1640995200
auth/login.go
auth/middleware.go

def456|fix: resolve login validation bug|Jane Smith|1640998800
auth/validation.go
auth/login.go`
	
	commits, err := gm.parseGitLog(gitLogOutput)
	if err != nil {
		t.Fatalf("Failed to parse git log: %v", err)
	}
	
	if len(commits) != 2 {
		t.Errorf("Expected 2 commits, got %d", len(commits))
	}
	
	// Check first commit
	commit1 := commits[0]
	if commit1.Hash != "abc123" {
		t.Errorf("Expected hash 'abc123', got '%s'", commit1.Hash)
	}
	if commit1.Message != "feat: add user authentication" {
		t.Errorf("Expected message 'feat: add user authentication', got '%s'", commit1.Message)
	}
	if commit1.Author != "John Doe" {
		t.Errorf("Expected author 'John Doe', got '%s'", commit1.Author)
	}
	if len(commit1.FilesChanged) != 2 {
		t.Errorf("Expected 2 files changed, got %d", len(commit1.FilesChanged))
	}
	if commit1.FilesChanged[0] != "auth/login.go" {
		t.Errorf("Expected first file 'auth/login.go', got '%s'", commit1.FilesChanged[0])
	}
}

func TestGitMonitor_GenerateSmartSummary(t *testing.T) {
	gm := NewGitMonitor("/tmp")
	
	tests := []struct {
		name     string
		message  string
		expected string
	}{
		{
			name:     "conventional commit feat",
			message:  "feat: add user authentication",
			expected: "Added add user authentication",
		},
		{
			name:     "conventional commit fix",
			message:  "fix: resolve login validation bug",
			expected: "Fixed resolve login validation bug",
		},
		{
			name:     "conventional commit with scope",
			message:  "feat(auth): add JWT middleware",
			expected: "Added add JWT middleware (auth)",
		},
		{
			name:     "conventional commit refactor",
			message:  "refactor: improve error handling",
			expected: "Refactored improve error handling",
		},
		{
			name:     "regular commit message",
			message:  "Update README with installation instructions",
			expected: "Update README with installation instructions",
		},
		{
			name:     "long commit message",
			message:  "This is a very long commit message that should be truncated because it exceeds the maximum length limit that we want to enforce for smart summaries",
			expected: "This is a very long commit message that should be truncated because it exceeds the maximum length...",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gm.generateSmartSummary(tt.message)
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestGitMonitor_InferScope(t *testing.T) {
	gm := NewGitMonitor("/tmp")
	
	tests := []struct {
		name     string
		files    []string
		expected string
	}{
		{
			name:     "authentication files",
			files:    []string{"auth/login.go", "auth/middleware.go"},
			expected: "authentication",
		},
		{
			name:     "API files",
			files:    []string{"api/handlers/user.go", "api/routes.go"},
			expected: "api",
		},
		{
			name:     "frontend files",
			files:    []string{"frontend/components/Login.tsx", "ui/styles.css"},
			expected: "frontend",
		},
		{
			name:     "test files",
			files:    []string{"auth_test.go", "api/handlers_test.go"},
			expected: "testing",
		},
		{
			name:     "documentation files",
			files:    []string{"README.md", "docs/api.md"},
			expected: "documentation", // Both README.md and docs/api.md match .md pattern
		},
		{
			name:     "mixed files with auth majority",
			files:    []string{"auth/login.go", "auth/middleware.go", "config.json"},
			expected: "authentication",
		},
		{
			name:     "directory-based inference", 
			files:    []string{"backend/server.go", "backend/config.go"},
			expected: "backend", // directory name wins over individual file patterns
		},
		{
			name:     "no files",
			files:    []string{},
			expected: "general",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := gm.inferScope(tt.files)
			if result != tt.expected {
				t.Errorf("Expected scope '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestGitMonitor_UpdateCurrentFocus(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "weather_git_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	gm := NewGitMonitor(tempDir)
	
	// Create test context
	ctx := &WeatherContext{
		CurrentFocus: FocusArea{
			Area: "initial area",
		},
	}
	
	// Test commit with authentication scope
	commit := GitCommit{
		Message:       "feat: add user authentication",
		FilesChanged:  []string{"auth/login.go", "auth/middleware.go"},
		InferredScope: "authentication",
	}
	
	gm.updateCurrentFocus(ctx, commit)
	
	if ctx.CurrentFocus.Area != "authentication system" {
		t.Errorf("Expected focus area 'authentication system', got '%s'", ctx.CurrentFocus.Area)
	}
	
	if ctx.CurrentFocus.Confidence != 0.85 {
		t.Errorf("Expected confidence 0.85, got %f", ctx.CurrentFocus.Confidence)
	}
	
	if ctx.CurrentFocus.InferredFrom != "recent commit to authentication" {
		t.Errorf("Expected inferred from 'recent commit to authentication', got '%s'", ctx.CurrentFocus.InferredFrom)
	}
}

func TestGitMonitor_UpdateWeatherConditions(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "weather_git_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	gm := NewGitMonitor(tempDir)
	
	tests := []struct {
		name      string
		momentum  int
		commits   []GitCommit
		expectedTemp int
		expectedCond WeatherCondition
	}{
		{
			name:     "high momentum with features",
			momentum: 80,
			commits: []GitCommit{
				{Message: "feat: add new feature"},
				{Message: "feat: improve UI"},
			},
			expectedTemp: 95, // min(20+80, 95)
			expectedCond: WeatherSunny,
		},
		{
			name:     "medium momentum with fixes",
			momentum: 40,
			commits: []GitCommit{
				{Message: "fix: resolve bug"},
				{Message: "fix: handle edge case"},
			},
			expectedTemp: 60, // 20+40
			expectedCond: WeatherCloudy,
		},
		{
			name:     "low momentum mixed commits",
			momentum: 20,
			commits: []GitCommit{
				{Message: "feat: add feature"},
				{Message: "fix: resolve issue"},
			},
			expectedTemp: 40, // 20+20
			expectedCond: WeatherPartlyCloudy,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &WeatherContext{
				RecentProgress: ProgressSummary{
					Momentum: tt.momentum,
					Commits:  tt.commits,
				},
				Weather: WeatherConditions{},
			}
			
			gm.updateWeatherConditions(ctx)
			
			if ctx.Weather.Temperature != tt.expectedTemp {
				t.Errorf("Expected temperature %d, got %d", tt.expectedTemp, ctx.Weather.Temperature)
			}
			
			if ctx.Weather.Condition != tt.expectedCond {
				t.Errorf("Expected condition %s, got %s", tt.expectedCond, ctx.Weather.Condition)
			}
		})
	}
}

func TestGitMonitor_UpdateNextSteps(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "weather_git_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	gm := NewGitMonitor(tempDir)
	
	tests := []struct {
		name           string
		focusArea      string
		momentum       int
		expectedSuggestions []string
	}{
		{
			name:      "authentication focus",
			focusArea: "authentication system",
			momentum:  60,
			expectedSuggestions: []string{
				"Add JWT token validation",
				"Implement user session management",
				"Add password reset functionality",
			},
		},
		{
			name:      "API focus",
			focusArea: "API development",
			momentum:  40,
			expectedSuggestions: []string{
				"Add API endpoint validation",
				"Implement error handling middleware", 
				"Add API documentation",
			},
		},
		{
			name:      "frontend focus",
			focusArea: "user interface",
			momentum:  80,
			expectedSuggestions: []string{
				"Add responsive styling",
				"Implement form validation",
				"Add loading states and error handling",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &WeatherContext{
				CurrentFocus: FocusArea{
					Area: tt.focusArea,
				},
				RecentProgress: ProgressSummary{
					Momentum: tt.momentum,
				},
			}
			
			gm.updateNextSteps(ctx)
			
			if len(ctx.NextSteps.Suggestions) != len(tt.expectedSuggestions) {
				t.Errorf("Expected %d suggestions, got %d", len(tt.expectedSuggestions), len(ctx.NextSteps.Suggestions))
			}
			
			for i, expected := range tt.expectedSuggestions {
				if i < len(ctx.NextSteps.Suggestions) && ctx.NextSteps.Suggestions[i] != expected {
					t.Errorf("Expected suggestion '%s', got '%s'", expected, ctx.NextSteps.Suggestions[i])
				}
			}
			
			expectedPriority := min(50+tt.momentum/2, 90)
			if ctx.NextSteps.Priority != expectedPriority {
				t.Errorf("Expected priority %d, got %d", expectedPriority, ctx.NextSteps.Priority)
			}
		})
	}
}

// Integration test that requires a real git repository
func TestGitMonitor_Integration(t *testing.T) {
	// Skip if git is not available
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not available, skipping integration test")
	}
	
	// Skip in CI environment where git might not be fully configured
	if os.Getenv("CI") == "true" {
		t.Skip("Skipping git integration test in CI environment")
	}
	
	tempDir, err := os.MkdirTemp("", "weather_git_integration")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)
	
	// Initialize git repository
	cmd := exec.Command("git", "init")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to init git repo: %v", err)
	}
	
	// Configure git user
	cmd = exec.Command("git", "config", "user.email", "test@example.com")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to set git user.email: %v", err)
	}
	
	cmd = exec.Command("git", "config", "user.name", "Test User")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to set git user.name: %v", err)
	}
	
	// Create a test file and commit
	testFile := filepath.Join(tempDir, "auth", "login.go")
	if err := os.MkdirAll(filepath.Dir(testFile), 0755); err != nil {
		t.Fatalf("Failed to create directory: %v", err)
	}
	
	if err := os.WriteFile(testFile, []byte("package auth\n// Login functionality"), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}
	
	// Add and commit file
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to add files: %v", err)
	}
	
	cmd = exec.Command("git", "commit", "-m", "feat: add user authentication")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to commit: %v", err)
	}
	
	// Test git monitor
	gm := NewGitMonitor(tempDir)
	
	// Get recent commits
	commits, err := gm.GetRecentCommits(1)
	if err != nil {
		t.Fatalf("Failed to get recent commits: %v", err)
	}
	
	if len(commits) != 1 {
		t.Errorf("Expected 1 commit, got %d", len(commits))
	}
	
	commit := commits[0]
	if commit.Message != "feat: add user authentication" {
		t.Errorf("Expected commit message 'feat: add user authentication', got '%s'", commit.Message)
	}
	
	if !strings.Contains(commit.FilesChanged[0], "auth/login.go") {
		t.Errorf("Expected file 'auth/login.go' to be changed, got %v", commit.FilesChanged)
	}
	
	// Test git hook installation
	if err := gm.InstallGitHooks(); err != nil {
		t.Fatalf("Failed to install git hooks: %v", err)
	}
	
	// Verify hooks were created
	postCommitHook := filepath.Join(tempDir, ".git", "hooks", "post-commit")
	if _, err := os.Stat(postCommitHook); os.IsNotExist(err) {
		t.Error("Post-commit hook was not created")
	}
	
	postCheckoutHook := filepath.Join(tempDir, ".git", "hooks", "post-checkout")
	if _, err := os.Stat(postCheckoutHook); os.IsNotExist(err) {
		t.Error("Post-checkout hook was not created")
	}
}
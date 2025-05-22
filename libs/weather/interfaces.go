// Package weather provides automatic context preservation and intelligence
// for development workflows, serving as the heartbeat of the Sprouted ecosystem.
package weather

// No imports needed for interfaces file

// Repository interfaces for better testability and separation of concerns

// GitRepository abstracts git operations for better testing and modularity
type GitRepository interface {
	GetCurrentBranch() (string, error)
	GetCommitInfo(hash string) (*GitCommit, error)
	GetRecentCommits(limit int) ([]GitCommit, error)
	HasUncommittedChanges() (bool, error)
	InstallHooks(sproutPath string) error
}

// ContextRepository abstracts context persistence
type ContextRepository interface {
	Load() (*WeatherContext, error)
	Save(context *WeatherContext) error
	GetPath() string
}

// InferenceEngine abstracts smart analysis and inference
type InferenceEngine interface {
	InferScope(files []string) string
	GenerateSmartSummary(message string) string
	InferFocusArea(commit GitCommit, currentBranch string) FocusArea
	CalculateWeatherConditions(commits []GitCommit) WeatherConditions
	GenerateNextSteps(focus FocusArea, progress ProgressSummary) NextStepsSuggestion
}

// WeatherService orchestrates all weather operations
type WeatherService interface {
	GetContext() (*WeatherContext, error)
	UpdateFromCommit(commitHash string) error
	UpdateFromBranchChange(prevHead, newHead, branchFlag string) error
	InstallGitHooks() error
}
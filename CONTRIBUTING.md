# Contributing to Garden Weather System

Thank you for your interest in contributing to the Garden Weather System! This project aims to eliminate context loss and enable seamless AI-assisted development.

## 🌱 Quick Start

1. **Fork and clone** the repository
2. **Install dependencies**: `make setup`
3. **Run tests**: `make test`
4. **Check weather context**: `sprout weather --onboard-ai`

## 📋 Development Process

We follow a **spec-driven development** methodology:

1. **Vision** → **Specs** → **Tasks** → **Implementation**
2. All significant changes should have corresponding specs
3. Use the Weather System to track your development context
4. Follow our [Agentic Development Workflow](docs/workflows/agentic-development.md)

## 🛠️ Setting Up Development Environment

### Prerequisites

**Required**: Git, Go 1.20+, GPG for signed commits

### 1. GPG Setup for Signed Commits

This project requires signed commits for security. Set up GPG signing:

```bash
# Generate a new GPG key (if you don't have one)
gpg --full-generate-key
# Choose: RSA and RSA, 4096 bits, no expiration, your name and GitHub email

# List your GPG keys to get the key ID
gpg --list-secret-keys --keyid-format=long
# Copy the key ID from sec line (after rsa4096/)

# Export your public key (replace KEY_ID with your actual key ID)
gpg --armor --export KEY_ID

# Add the exported key to your GitHub account:
# Settings > SSH and GPG keys > New GPG key
```

**Configure Git for signed commits:**
```bash
# Set your signing key (replace KEY_ID with your actual key ID)
git config --global user.signingkey KEY_ID

# Enable signed commits by default
git config --global commit.gpgsign true

# Test your setup
git commit --allow-empty -m "test: verify GPG signing works"
```

**Troubleshooting GPG:**
- **macOS**: Install GPG Suite or use `brew install gnupg`
- **Linux**: Install `gnupg2` package
- **Windows**: Install Git for Windows (includes GPG) or GPG4Win
- **VS Code**: Install "Git Commit Plugin" extension for GUI signing

### 2. Development Environment

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/garden.git
cd garden

# Set up development environment
make setup

# Build and test
make build
make test

# Install weather hooks for context tracking
./apps/sprout-cli/sprout weather --install-hooks

# Check your development context
./apps/sprout-cli/sprout weather --onboard-ai
```

### 3. Verify Setup

```bash
# Verify GPG signing works
git log --show-signature -1

# Should show "Good signature" for your commits
```

## 📝 How to Contribute

### 🐛 Bug Reports

Before submitting a bug report:
- Search existing issues to avoid duplicates
- Update to the latest version
- Test with a minimal reproduction case

**Use our bug report template** when creating issues.

### 💡 Feature Requests

For new features:
- Check if it aligns with our [project vision](docs/vision/weather-context-preservation.md)
- Create a spec document following our [spec template](templates/spec-template.md)
- Discuss in GitHub Discussions before implementing

### 🔧 Code Contributions

#### Pull Request Process

1. **Create a feature branch** from `main`
2. **Write/update specs** for significant changes
3. **Follow coding standards** (run `make lint` and `make fmt`)
4. **Add tests** for new functionality
5. **Update documentation** as needed
6. **Ensure CI passes** (all tests and linting)
7. **Create pull request** with descriptive title and description

#### Coding Standards

- **Go**: Follow standard Go conventions, use `gofmt` and `golangci-lint`
- **Documentation**: Update relevant specs and user documentation
- **Tests**: Write tests for new functionality (target: >80% coverage)
- **Commits**: Use conventional commit format (see below)

#### Commit Message Format

```
type(scope): description

[optional body]

[optional footer]
```

**Types**: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`

**Examples**:
```
feat(weather): add git branch monitoring
fix(cli): resolve weather context loading error
docs(api): update weather command documentation
```

## 🏗️ Project Structure

```
garden/
├── apps/           # Applications (sprout-cli)
├── libs/           # Libraries (weather system)
├── docs/           # Documentation hierarchy
│   ├── vision/     # Long-term vision documents
│   ├── specs/      # Detailed specifications
│   ├── tasks/      # Implementation tasks
│   └── workflows/  # Development processes
├── templates/      # Document templates
└── tools/          # Development tools
```

## 🧪 Testing

```bash
# Run all tests
make test

# Run specific component tests
cd libs/weather && go test -v
cd apps/sprout-cli && go test -v

# Run with coverage
make test-coverage
```

## 📖 Documentation

- **User documentation**: Update relevant docs in `docs/`
- **API documentation**: Use Go doc comments
- **Examples**: Add examples to demonstrate new features
- **Specs**: Create/update specification documents for significant changes

## 🌦️ Weather System Integration

This project uses its own Weather System for development:

- **Track your progress**: `sprout weather`
- **AI onboarding**: `sprout weather --onboard-ai`
- **Context sharing**: Share weather context when reporting issues

## 🎯 Focus Areas

Current development priorities:

1. **Weather MVP**: Core context preservation features
2. **Farm Architecture**: Multi-repository workspace support
3. **AI Integration**: Enhanced AI assistant collaboration
4. **Community Features**: Seed sharing and pattern discovery

See [active specs](docs/specs/) and [tasks](docs/tasks/active/) for current work.

## 🤝 Community Guidelines

- **Be respectful**: Follow our [Code of Conduct](CODE_OF_CONDUCT.md)
- **Be helpful**: Support other contributors and users
- **Be patient**: Maintainers review contributions as time allows
- **Be collaborative**: Work together to improve the project

## 💬 Getting Help

- **GitHub Discussions**: For questions and community discussion
- **GitHub Issues**: For bug reports and feature requests
- **Documentation**: Check our comprehensive [docs](docs/)
- **Weather Context**: Use `sprout weather --onboard-ai` for project context

## 🏷️ Issue Labels

- `good first issue`: Good for newcomers
- `help wanted`: Extra attention needed
- `bug`: Something isn't working
- `enhancement`: New feature or request
- `documentation`: Improvements to docs
- `question`: Further information requested

## 📊 Project Roadmap

**Phase 1**: Weather MVP (Current)
- Core context preservation
- Git integration
- AI assistant onboarding

**Phase 2**: Farm Architecture
- Multi-repository coordination
- Event-based weather system
- Workspace intelligence

**Phase 3**: Community Features
- Seed sharing and discovery
- Pattern libraries
- Collaborative development tools

## 🙏 Recognition

Contributors are recognized in:
- GitHub contributors list
- Release notes for significant contributions
- Community showcase for innovative usage

Thank you for contributing to the Garden Weather System! Together we're building the future of context-aware development. 🌱

---

**Questions?** Open a discussion or check our [documentation](docs/). The Weather System itself is the best way to understand our development approach!
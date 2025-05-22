# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 0.1.x   | :white_check_mark: |

## Reporting a Vulnerability

We take security seriously. If you discover a security vulnerability, please follow these steps:

### 1. **Do Not** Open a Public Issue

Please do not report security vulnerabilities through public GitHub issues.

### 2. Report Privately

Send an email to security@sprouted.dev with:
- Description of the vulnerability
- Steps to reproduce the issue
- Possible impact
- Any suggested fixes

### 3. Response Timeline

- **Initial Response**: Within 48 hours
- **Status Update**: Within 1 week
- **Resolution**: Coordinated disclosure timeline

### 4. Coordinated Disclosure

We follow responsible disclosure practices:
- Work with you to understand and fix the issue
- Provide credit for the discovery (if desired)
- Coordinate public disclosure timing

## Security Considerations

### Weather Context Data

The Weather System stores development context locally in JSON files. This data may include:
- File paths and names
- Git commit messages and metadata
- Development patterns and focus areas

**Important**: Weather context files should not contain sensitive information like:
- Passwords or API keys
- Personal or confidential data
- Production environment details

### Installation Security

- Install scripts verify checksums and signatures
- Go modules use secure HTTPS downloads
- Dependencies are regularly updated for security

### Data Privacy

- Weather context is stored locally by default
- No data is transmitted to external servers without explicit configuration
- Users have full control over their context data

## Best Practices

When using the Garden Weather System:

1. **Review context data** before sharing with AI assistants
2. **Use .gitignore** to exclude sensitive files from weather tracking
3. **Regular updates** to get latest security patches
4. **Secure your development environment** as weather preserves context

## Security Updates

Security updates are released as patch versions and announced through:
- GitHub Security Advisories
- Release notes
- Project documentation updates

## Contact

For security concerns: security@sprouted.dev
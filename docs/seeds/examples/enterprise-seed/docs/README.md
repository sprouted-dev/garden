# Enterprise Seed Example

A comprehensive documentation structure for enterprise teams with compliance and governance requirements.

## How We Work

Enterprise development with focus on:
- Architectural governance
- Compliance documentation
- Standardized processes
- Knowledge transfer

## Directory Structure

```
docs/
├── README.md              # You are here
├── architecture/          # System architecture
│   ├── overview.md       # High-level architecture
│   ├── decisions/        # ADRs with enterprise review
│   ├── diagrams/         # Architecture diagrams
│   └── reviews/          # Architecture review board notes
├── compliance/           # Regulatory compliance
│   ├── requirements/     # Compliance requirements
│   ├── audits/          # Audit reports
│   └── controls/        # Security controls
├── processes/           # Development processes
│   ├── sdlc.md         # Software Development Lifecycle
│   ├── code-review.md  # Review standards
│   ├── deployment.md   # Deployment procedures
│   └── incident.md     # Incident response
├── standards/          # Coding and design standards
│   ├── coding/         # Language-specific standards
│   ├── api/            # API design standards
│   ├── database/       # Data standards
│   └── security/       # Security standards
└── training/           # Onboarding and training
    ├── onboarding/     # New developer guides
    ├── workshops/      # Training materials
    └── certifications/ # Required certifications
```

## Key Concepts

### Governance Model
- Architecture Review Board approval required
- Compliance checkpoints at each phase
- Standardized tooling and processes
- Regular audits and assessments

### Documentation Requirements
- Every feature requires architectural review
- Security assessment for all changes
- Compliance impact analysis
- Training materials for new capabilities

### Quality Gates
- Code review against standards
- Architecture conformance check
- Security scan results
- Performance benchmarks

## Weather Integration Points

- **Compliance Tracking**: Weather monitors required documentation
- **Process Adherence**: Detects deviation from standards
- **Training Needs**: Identifies knowledge gaps
- **Audit Readiness**: Maintains compliance evidence

## Getting Started

1. Complete onboarding in `training/onboarding/`
2. Review relevant standards in `standards/`
3. Understand architecture via `architecture/overview.md`
4. Follow processes defined in `processes/`

## Compliance Notes

This structure supports:
- SOC2 Type II requirements
- ISO 27001 controls
- GDPR documentation needs
- Industry-specific regulations

## Enterprise Weather Configuration

```json
{
  "weather_hints": {
    "compliance_indicators": ["compliance/", "architecture/reviews/"],
    "process_checkpoints": ["processes/", "standards/"],
    "training_tracking": ["training/certifications/"],
    "audit_trails": {
      "required": true,
      "retention": "7 years",
      "format": "immutable"
    }
  }
}
```
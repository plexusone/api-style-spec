# API Style Spec

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Docs][docs-mkdoc-svg]][docs-mkdoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/plexusone/api-style-spec/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/plexusone/api-style-spec/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/plexusone/api-style-spec/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/plexusone/api-style-spec/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/plexusone/api-style-spec/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/plexusone/api-style-spec/actions/workflows/go-sast-codeql.yaml
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/plexusone/api-style-spec
 [docs-godoc-url]: https://pkg.go.dev/github.com/plexusone/api-style-spec
 [docs-mkdoc-svg]: https://img.shields.io/badge/Go-dev%20guide-blue.svg
 [docs-mkdoc-url]: https://plexusone.dev/api-style-spec
 [viz-svg]: https://img.shields.io/badge/Go-visualizaton-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=plexusone%2Fapi-style-spec
 [loc-svg]: https://tokei.rs/b1/github/plexusone/api-style-spec
 [repo-url]: https://github.com/plexusone/api-style-spec
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/plexusone/api-style-spec/blob/main/LICENSE

Machine-readable API style specification format that generates human documentation, linting rules, LLM evaluation rubrics, and AI agent instructions from a single source of truth.

## Overview

API style guides from Microsoft, Google, and Zalando have become industry standards, but they exist only as human-readable documents. **api-style-spec** creates a machine-readable specification format that serves as the canonical source, generating all artifacts from one definition.

```
api-style-spec (source of truth)
    ├── Human Style Guide (Markdown)
    ├── Deterministic Linters (Spectral/vacuum)
    ├── LLM Review Rubrics
    ├── AI Agent Instructions (Claude Code, Kiro)
    └── MCP Server Tools
```

## Features

- 📋 **Unified Specification** - Define rules once, generate all artifacts
- ✅ **Deterministic Linting** - Fast, CI-friendly checks via [vacuum](https://github.com/daveshanley/vacuum)
- 🧠 **LLM Evaluation** - Semantic analysis for rules that can't be linted
- 🏢 **Industry Profiles** - Pre-built profiles based on Microsoft, Google, Zalando guidelines
- 🏆 **Conformance Levels** - Graduated compliance (bronze/silver/gold)
- 🌐 **Multi-Platform** - CLI, Web UI, MCP server, AI agents

## Installation

```bash
go install github.com/plexusone/api-style-spec/cmd/api-style@latest
```

## Quick Start

```bash
# Lint an OpenAPI specification
api-style lint openapi.yaml

# Lint with a specific profile
api-style lint openapi.yaml --profile azure

# Lint multiple files with glob patterns
api-style lint api/*.yaml --recursive

# Watch mode for continuous linting
api-style lint openapi.yaml --watch

# Combined lint + LLM evaluation
api-style analyze openapi.yaml --profile azure --level silver

# Generate human-readable style guide
api-style generate guide --spec my-style.json --output docs/
```

## Configuration

Create `.api-style.yaml` in your project root:

```yaml
# .api-style.yaml
profile: azure
level: silver

include:
  - "openapi.yaml"
  - "**/api.yaml"

exclude:
  - "**/generated/**"

exceptions:
  - rule: URI-001
    paths: ["/legacy/**"]
    reason: "Legacy API cannot be changed"

severity-overrides:
  URI-002: warn
```

See [.api-style.yaml.example](.api-style.yaml.example) for a complete example.

## Specification Format

```json
{
  "$schema": "https://api-style-spec.dev/schema/v1/api-style-spec.schema.json",
  "version": "1.0.0",
  "name": "my-api-style",
  "extends": ["default"],

  "rules": [
    {
      "id": "URI-001",
      "title": "Use plural resource names",
      "category": "uri-design",
      "severity": "error",
      "rationale": "Plural resources improve consistency.",
      "examples": {
        "good": ["/users", "/orders"],
        "bad": ["/user", "/order"]
      },
      "enforcement": {
        "type": "spectral",
        "function": "pattern",
        "options": {"match": "^/[a-z]+s(/|$)"}
      }
    }
  ]
}
```

## CLI Commands

| Command | Description |
|---------|-------------|
| `api-style lint` | Deterministic linting (supports glob patterns, `--watch`, `--recursive`) |
| `api-style evaluate` | LLM-based evaluation |
| `api-style analyze` | Combined lint + evaluate |
| `api-style score-profile` | Score a style profile using LLM evaluation |
| `api-style generate guide` | Generate Markdown documentation |
| `api-style generate mkdocs` | Generate MkDocs multi-page site |
| `api-style generate spectral` | Generate Spectral ruleset |
| `api-style generate rubric` | Generate LLM evaluation rubric |
| `api-style hooks` | Generate AI assistant hooks |
| `api-style hooks init` | Install git pre-commit hook |
| `api-style diff` | Breaking change detection |
| `api-style serve mcp` | Start MCP server |
| `api-style serve web` | Start Web UI |

## Built-in Profiles

| Profile | Rules | Categories | Focus |
|---------|-------|------------|-------|
| `default` | 106 | 27 | Industry-leading, SDK-optimized (ogen) |
| `comprehensive` | 88 | 26 | Full coverage, all best practices |
| `zalando` | 147 | 13 | E-commerce, events |
| `microsoft-rest` | 123 | 15 | Enterprise REST APIs |
| `microsoft-graph` | 82 | 12 | OData/Graph APIs |
| `azure` | 23 | 9 | Azure cloud services |
| `google` | 20 | 7 | Resource-oriented design |
| `minimal` | 29 | 7 | Basic API hygiene |

**Default Profile Highlights:**

- 100% evaluable with LLM-as-Judge criteria
- 34% deterministic Spectral enforcement
- SDK-optimized for ogen, openapi-generator
- Multi-tenancy patterns with `~` alias
- RFC 9457 Problem Details for errors
- Discriminated unions for polymorphism

## Documentation

- [Getting Started](docs/guide/getting-started.md)
- [Automated API Governance](docs/guide/automated-api-governance.md) - AI-first API design workflow
- [CI/CD Integration](docs/guide/ci-cd.md) - Pipeline integration and pre-commit hooks
- [Documentation Generation](docs/guide/documentation-generation.md) - Generate Markdown and MkDocs sites
- [Creating Custom Profiles](docs/guide/profiles.md)
- [Writing Custom Rules](docs/guide/custom-rules.md)
- [MRD](docs/specs/inception/MRD.md) | [PRD](docs/specs/inception/PRD.md) | [TRD](docs/specs/inception/TRD.md) | [Roadmap](docs/specs/inception/ROADMAP.md)

## Related Projects

- [vacuum](https://github.com/daveshanley/vacuum) - Fast OpenAPI linter (used internally)
- [structured-evaluation](https://github.com/plexusone/structured-evaluation) - LLM evaluation framework
- [multi-agent-spec](https://github.com/plexusone/multi-agent-spec) - Agent definitions
- [assistantkit](https://github.com/plexusone/assistantkit) - AI assistant file generation

## License

MIT

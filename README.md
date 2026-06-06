# api-style-spec

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

- **Unified Specification** - Define rules once, generate all artifacts
- **Deterministic Linting** - Fast, CI-friendly checks via [vacuum](https://github.com/daveshanley/vacuum)
- **LLM Evaluation** - Semantic analysis for rules that can't be linted
- **Industry Profiles** - Pre-built profiles based on Microsoft, Google, Zalando guidelines
- **Conformance Levels** - Graduated compliance (bronze/silver/gold)
- **Multi-Platform** - CLI, Web UI, MCP server, AI agents

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

# Combined lint + LLM evaluation
api-style analyze openapi.yaml --profile azure --level silver

# Generate human-readable style guide
api-style generate guide --spec my-style.json --output docs/
```

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
| `api-style lint` | Deterministic linting |
| `api-style evaluate` | LLM-based evaluation |
| `api-style analyze` | Combined lint + evaluate |
| `api-style generate guide` | Generate Markdown documentation |
| `api-style generate spectral` | Generate Spectral ruleset |
| `api-style generate agent` | Generate AI agent files |
| `api-style diff` | Breaking change detection |
| `api-style serve mcp` | Start MCP server |
| `api-style serve web` | Start Web UI |

## Built-in Profiles

| Profile | Based On | Rules |
|---------|----------|-------|
| `default` | Common best practices | ~30 |
| `azure` | Microsoft/Azure guidelines | ~80 |
| `google` | Google API Design Guide | ~60 |
| `zalando` | Zalando RESTful Guidelines | ~70 |

## Documentation

- [Getting Started](docs/guide/getting-started.md)
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

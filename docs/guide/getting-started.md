# Getting Started

This guide walks you through installing api-style-spec and linting your first OpenAPI specification.

## Installation

### From Source (Go)

```bash
go install github.com/plexusone/api-style-spec/cmd/api-style@latest
```

### Verify Installation

```bash
api-style version
```

## Your First Lint

Create a simple OpenAPI specification:

```yaml
# openapi.yaml
openapi: "3.1.0"
info:
  title: My API
  version: "1.0.0"
paths:
  /user:
    get:
      responses:
        "200":
          description: OK
```

Run the linter:

```bash
api-style lint openapi.yaml
```

You'll see output like:

```
API Style Lint Report
=====================

Status: FAIL
Errors: 1, Warnings: 2, Info: 0, Hints: 0

Errors:
  - [URI-001] Use plural resource names
    $.paths['/user']

Warnings:
  - [operation-description] operation method `GET` is missing a description
    $.paths['/user'].get
```

## Understanding Violations

Each violation includes:

| Field | Description |
|-------|-------------|
| Rule ID | Unique identifier (e.g., `URI-001`) |
| Message | Human-readable explanation |
| Path | JSON path to the violation location |
| Severity | `error`, `warn`, `info`, or `hint` |

## Fixing Violations

Update your spec to fix the violations:

```yaml
# openapi.yaml
openapi: "3.1.0"
info:
  title: My API
  version: "1.0.0"
  description: A sample API
paths:
  /users:  # Changed from /user to /users (plural)
    get:
      summary: List users
      description: Returns a list of all users
      operationId: listUsers
      responses:
        "200":
          description: OK
```

Run again:

```bash
api-style lint openapi.yaml
```

```
API Style Lint Report
=====================

Status: PASS
Errors: 0, Warnings: 0, Info: 0, Hints: 0

No violations found.
```

## Using Different Profiles

api-style-spec includes profiles based on industry guidelines:

```bash
# Use Azure/Microsoft guidelines
api-style lint openapi.yaml --profile azure

# Use Google API Design Guide
api-style lint openapi.yaml --profile google

# Use Zalando RESTful Guidelines
api-style lint openapi.yaml --profile zalando
```

See [Using Profiles](profiles.md) for details on each profile.

## Output Formats

### Text (default)

```bash
api-style lint openapi.yaml
```

### JSON

```bash
api-style lint openapi.yaml --format json
```

### SARIF (for IDE integration)

```bash
api-style lint openapi.yaml --format sarif --output report.sarif
```

## Combined Analysis

For comprehensive review including LLM-based semantic analysis:

```bash
# Requires ANTHROPIC_API_KEY
export ANTHROPIC_API_KEY=sk-ant-...
api-style analyze openapi.yaml --profile azure
```

This combines:

1. Deterministic linting (fast, CI-friendly)
2. LLM evaluation (semantic analysis of design quality)
3. GO/NO-GO decision for production readiness

## Next Steps

- [Using Profiles](profiles.md) - Understand built-in style profiles
- [Custom Rules](custom-rules.md) - Create your own rules
- [CLI Reference](../reference/cli.md) - Complete command documentation

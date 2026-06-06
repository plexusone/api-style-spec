# CLI Reference

Complete reference for the `api-style` command-line interface.

## Global Options

| Option | Description |
|--------|-------------|
| `--help`, `-h` | Show help for any command |
| `--version` | Show version information |

## Commands

### lint

Lint an OpenAPI specification against style rules.

```bash
api-style lint <openapi-spec> [flags]
```

**Arguments:**

| Argument | Description |
|----------|-------------|
| `openapi-spec` | Path to the OpenAPI specification file |

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | `-f` | `text` | Output format: `text`, `json`, `sarif` |
| `--output` | `-o` | stdout | Output file path |
| `--profile` | `-p` | `default` | Style profile to use |
| `--level` | `-l` | | Conformance level: `bronze`, `silver`, `gold` |

**Examples:**

```bash
# Basic lint
api-style lint openapi.yaml

# JSON output
api-style lint openapi.yaml --format json

# Use Azure profile
api-style lint openapi.yaml --profile azure

# Save SARIF report
api-style lint openapi.yaml --format sarif --output report.sarif

# Check silver conformance
api-style lint openapi.yaml --profile default --level silver
```

**Exit Codes:**

| Code | Meaning |
|------|---------|
| 0 | No blocking violations |
| 1 | Blocking violations found (errors) |

---

### evaluate

LLM-based semantic evaluation of an API specification.

```bash
api-style evaluate <openapi-spec> [flags]
```

**Requires:** `ANTHROPIC_API_KEY` environment variable

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | `-f` | `text` | Output format: `text`, `json` |
| `--output` | `-o` | stdout | Output file path |
| `--profile` | `-p` | `default` | Style profile to use |
| `--categories` | `-c` | all | Categories to evaluate (comma-separated) |

**Examples:**

```bash
# Evaluate with default profile
api-style evaluate openapi.yaml

# Evaluate specific categories
api-style evaluate openapi.yaml --categories uri-design,documentation
```

---

### analyze

Combined lint + evaluate with GO/NO-GO decision.

```bash
api-style analyze <openapi-spec> [flags]
```

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | `-f` | `text` | Output format: `text`, `json` |
| `--output` | `-o` | stdout | Output file path |
| `--profile` | `-p` | `default` | Style profile to use |
| `--lint-only` | | `false` | Skip LLM evaluation |

**Examples:**

```bash
# Full analysis
api-style analyze openapi.yaml --profile azure

# Lint only (no LLM)
api-style analyze openapi.yaml --lint-only
```

---

### generate

Generate artifacts from a style profile.

#### generate guide

Generate Markdown documentation from a profile.

```bash
api-style generate guide [flags]
```

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--profile` | `-p` | `default` | Style profile to use |
| `--output` | `-o` | stdout | Output directory or file |

#### generate spectral

Generate a Spectral ruleset from a profile.

```bash
api-style generate spectral [flags]
```

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--profile` | `-p` | `default` | Style profile to use |
| `--output` | `-o` | stdout | Output file path |

---

### hooks

Generate AI assistant hooks configuration.

```bash
api-style hooks [flags]
api-style hooks generate [flags]
api-style hooks list
```

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | `-f` | `claude` | Output format: `claude`, `cursor`, `windsurf`, `all` |
| `--output` | `-o` | auto | Output file path |
| `--profile` | `-p` | `default` | Style profile for linting |
| `--auto-lint` | | `true` | Enable auto-linting on file save |
| `--inject-context` | | `false` | Inject style context before prompts |
| `--list` | | | List supported formats |

**Examples:**

```bash
# Generate Claude Code hooks
api-style hooks --format claude

# Generate for all assistants
api-style hooks --format all

# List supported formats
api-style hooks list

# Custom output path
api-style hooks --format claude --output .claude/settings.json
```

**Default Output Paths:**

| Format | Default Path |
|--------|--------------|
| `claude` | `.claude/settings.json` |
| `cursor` | `.cursor/hooks.json` |
| `windsurf` | `.windsurf/hooks.json` |

---

### version

Print version information.

```bash
api-style version
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `ANTHROPIC_API_KEY` | Required for `evaluate` and `analyze` commands |

## Output Formats

### Text

Human-readable output with colors and formatting.

### JSON

Machine-readable JSON output:

```json
{
  "status": "fail",
  "violations": [
    {
      "rule_id": "URI-001",
      "severity": "error",
      "message": "Use plural resource names",
      "path": "$.paths['/user']",
      "line": 5
    }
  ],
  "summary": {
    "errors": 1,
    "warnings": 0,
    "total": 1
  }
}
```

### SARIF

Static Analysis Results Interchange Format for IDE integration:

```json
{
  "$schema": "https://raw.githubusercontent.com/oasis-tcs/sarif-spec/master/Schemata/sarif-schema-2.1.0.json",
  "version": "2.1.0",
  "runs": [
    {
      "tool": {
        "driver": {
          "name": "api-style",
          "version": "0.1.0"
        }
      },
      "results": [...]
    }
  ]
}
```

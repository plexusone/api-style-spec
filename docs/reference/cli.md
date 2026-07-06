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
api-style lint <openapi-spec> [<openapi-spec>...] [flags]
```

**Arguments:**

| Argument | Description |
|----------|-------------|
| `openapi-spec` | Path(s) to OpenAPI specification file(s). Supports glob patterns. |

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | `-f` | `text` | Output format: `text`, `json`, `sarif` |
| `--output` | `-o` | stdout | Output file path |
| `--profile` | `-p` | from config | Style profile to use |
| `--level` | `-l` | | Conformance level: `bronze`, `silver`, `gold` |
| `--config` | `-c` | `.api-style.yaml` | Config file path |
| `--recursive` | `-r` | `false` | Search directories recursively |
| `--watch` | `-w` | `false` | Watch files for changes and re-lint |

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

# Lint multiple files with glob pattern
api-style lint api/*.yaml

# Lint directory recursively
api-style lint . --recursive

# Watch mode for continuous linting
api-style lint openapi.yaml --watch

# Use explicit config file
api-style lint openapi.yaml --config .api-style.yaml
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

### score-profile

Evaluate a style guide profile against quality criteria using LLM-as-a-Judge.

```bash
api-style score-profile <profile-name-or-file> [flags]
```

**Requires:** `ANTHROPIC_API_KEY` environment variable

This command assesses how complete and well-structured a style guide is,
evaluating categories like content coverage, rule quality, examples, and more.

**Arguments:**

| Argument | Description |
|----------|-------------|
| `profile-name-or-file` | Built-in profile name (e.g., `default`, `azure`) or path to a custom profile file |

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--format` | `-f` | `text` | Output format: `text`, `json` |
| `--output` | `-o` | stdout | Output file path |
| `--model` | `-m` | `claude-3-5-haiku` | LLM model to use |

**Examples:**

```bash
# Score the default profile
api-style score-profile default

# Score Azure profile with JSON output
api-style score-profile azure --format json

# Score a custom profile file
api-style score-profile ./custom-profile.json --output scores.json

# Use a different model
api-style score-profile zalando --model claude-sonnet-4
```

---

### generate

Generate artifacts from a style profile.

#### generate guide

Generate single-page Markdown documentation from a profile.

```bash
api-style generate guide [flags]
```

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--profile` | `-p` | `default` | Style profile to use |
| `--output` | `-o` | stdout | Output file path |
| `--no-toc` | | `false` | Exclude table of contents |
| `--no-examples` | | `false` | Exclude examples |
| `--no-patterns` | | `false` | Exclude design patterns section |
| `--no-principles` | | `false` | Exclude design principles section |
| `--no-glossary` | | `false` | Exclude glossary |
| `--emojis` | | `false` | Use emojis for severity indicators |

**Examples:**

```bash
# Generate full guide to stdout
api-style generate guide --profile zalando

# Save to file
api-style generate guide --profile azure --output azure-guide.md

# Minimal output (rules only)
api-style generate guide --profile azure --no-patterns --no-principles --no-glossary
```

#### generate mkdocs

Generate a complete MkDocs documentation site from a profile.

```bash
api-style generate mkdocs [flags]
```

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--profile` | `-p` | `default` | Style profile to use |
| `--output` | `-o` | `./docs` | Output directory |
| `--site-name` | | from profile | MkDocs site name |
| `--site-url` | | | Base URL for the site |
| `--repo-url` | | from profile | Repository URL |
| `--theme` | | `material` | MkDocs theme |
| `--split-patterns` | | `false` | Create separate pages per pattern |
| `--no-split-categories` | | `false` | Keep all rules in one page |
| `--no-search` | | `false` | Disable search plugin |

**Examples:**

```bash
# Generate MkDocs site
api-style generate mkdocs --profile zalando --output ./zalando-docs

# Custom site name and URL
api-style generate mkdocs --profile azure \
  --site-name "My API Guidelines" \
  --site-url "https://api.example.com/guidelines"

# Split patterns into individual pages
api-style generate mkdocs --profile azure --split-patterns

# Build and serve the generated site
cd zalando-docs && pip install mkdocs-material && mkdocs serve
```

**Generated Structure:**

```
output/
├── mkdocs.yml              # MkDocs configuration
└── docs/
    ├── index.md            # Home page
    ├── introduction.md     # Introduction (if defined)
    ├── principles.md       # Design principles
    ├── patterns.md         # Design patterns
    ├── conformance.md      # Conformance levels
    ├── glossary.md         # Glossary
    └── rules/
        ├── index.md        # Rules overview
        └── {category}.md   # One file per category
```

#### generate rubric

Generate a structured-evaluation rubric for LLM-as-Judge evaluation.

```bash
api-style generate rubric [flags]
```

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--profile` | `-p` | `default` | Style profile to use |
| `--output` | `-o` | stdout | Output file path |

**Examples:**

```bash
# Generate rubric to stdout
api-style generate rubric --profile azure

# Save rubric to file
api-style generate rubric --profile zalando --output zalando.rubric.json

# Use with structured-evaluation
api-style generate rubric --profile azure --output rubric.json
structured-eval evaluate openapi.yaml --rubric rubric.json
```

The generated rubric:

- Groups rules by category into rubric categories
- Converts rule severity to required/optional status
- Includes pass/partial/fail criteria from rule definitions
- Provides few-shot examples for each category
- Compatible with the [structured-evaluation](https://github.com/plexusone/structured-evaluation) framework

---

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

**Examples:**

```bash
# Generate Spectral ruleset
api-style generate spectral --profile azure --output .spectral.yaml

# Use with vacuum
vacuum lint openapi.yaml -r .spectral.yaml
```

---

### hooks

Generate AI assistant hooks configuration and git pre-commit hooks.

```bash
api-style hooks [flags]
api-style hooks generate [flags]
api-style hooks list
api-style hooks init [flags]
```

#### hooks generate

Generate AI assistant hooks for Claude Code, Cursor, or Windsurf.

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

#### hooks init

Install a git pre-commit hook that lints staged OpenAPI files before each commit.

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--profile` | `-p` | `default` | Style profile for linting |
| `--level` | `-l` | | Conformance level to enforce |
| `--force` | | `false` | Overwrite existing pre-commit hook |

**Examples:**

```bash
# Install pre-commit hook with defaults
api-style hooks init

# Install with Azure profile
api-style hooks init --profile azure

# Install with silver conformance level
api-style hooks init --level silver

# Overwrite existing hook
api-style hooks init --force
```

The hook:

- Runs on `git commit`
- Lints staged OpenAPI/Swagger files (`.yaml`, `.yml`, `.json`)
- Blocks commit if errors are found
- Can be bypassed with `git commit --no-verify`

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

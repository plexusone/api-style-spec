# AI Assistant Hooks

Configure automatic API linting when working with AI coding assistants.

## Overview

Hooks enable automatic API style checking when you save OpenAPI specification files. They integrate with AI assistants like Claude Code, Cursor, and Windsurf.

## Supported Assistants

| Assistant | Config File | Events |
|-----------|-------------|--------|
| Claude Code | `.claude/settings.json` | 16 events |
| Cursor | `.cursor/hooks.json` | 12 events |
| Windsurf | `.windsurf/hooks.json` | 9 events |

## Quick Setup

```bash
# Generate hooks for Claude Code
api-style hooks --format claude

# Generate for all assistants
api-style hooks --format all

# See supported formats
api-style hooks list
```

## Hook Types

### Auto-Lint on Save

Automatically lint OpenAPI files when they're saved:

```bash
api-style hooks --format claude --auto-lint
```

This creates a hook that:

1. Triggers when files are written (Edit/Write tools)
2. Checks if the file matches OpenAPI patterns
3. Runs `mcp-api-style lint` if it matches
4. Returns violations to the assistant

**Matched Patterns:**

- `openapi.yaml`, `openapi.yml`, `openapi.json`
- `swagger.yaml`, `swagger.yml`, `swagger.json`
- `**/openapi.yaml`, `**/api.yaml`, etc.

### Context Injection

Inject API style guidelines before prompts:

```bash
api-style hooks --format claude --inject-context
```

This adds context about the active style profile to help the AI assistant write better APIs.

## Configuration Options

| Option | Default | Description |
|--------|---------|-------------|
| `--format` | `claude` | Target assistant |
| `--profile` | `default` | Style profile for linting |
| `--auto-lint` | `true` | Enable auto-lint on save |
| `--inject-context` | `false` | Inject style context |
| `--output` | auto | Output file path |

## Generated Configuration

### Claude Code

`.claude/settings.json`:

```json
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Write|Edit",
        "hooks": [
          {
            "type": "command",
            "command": "#!/bin/bash\nFILE=\"$CLAUDE_FILE_PATH\"\n...",
            "timeout": 30
          }
        ]
      }
    ]
  }
}
```

### Cursor

`.cursor/hooks.json`:

```json
{
  "hooks": {
    "after_file_write": [
      {
        "pattern": "Write|Edit",
        "actions": [
          {
            "type": "shell",
            "command": "mcp-api-style lint ...",
            "timeout": 30000
          }
        ]
      }
    ]
  }
}
```

### Windsurf

`.windsurf/hooks.json`:

```json
{
  "hooks": {
    "onFileSave": [
      {
        "glob": "**/*.yaml",
        "command": "mcp-api-style lint ..."
      }
    ]
  }
}
```

## Prerequisites

For hooks to work, the `mcp-api-style` binary must be in your PATH:

```bash
go install github.com/plexusone/api-style-spec/cmd/mcp-api-style@latest
```

## Manual Configuration

### Claude Code

Add to `.claude/settings.json`:

```json
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Write|Edit",
        "hooks": [
          {
            "type": "command",
            "command": "mcp-api-style lint -f \"$CLAUDE_FILE_PATH\" -p default 2>&1 | head -50",
            "timeout": 30
          }
        ]
      }
    ]
  }
}
```

### Using with MCP Server

For full integration, combine hooks with the MCP server:

1. Configure MCP server in settings
2. Generate hooks for auto-linting
3. AI assistant can use both automatic checks and on-demand tools

## Workflow Example

1. **Setup:**
   ```bash
   api-style hooks --format claude
   ```

2. **Working with Claude:**
   - Ask Claude to create an OpenAPI spec
   - When Claude writes the file, hooks run automatically
   - Violations appear in the conversation
   - Claude can fix issues immediately

3. **On-Demand Tools:**
   - "Use the lint tool to check this against Azure guidelines"
   - "Explain rule URI-001"
   - "Run analyze for a GO/NO-GO decision"

## Troubleshooting

### Hooks Not Triggering

1. Verify the config file exists in the correct location
2. Check that `mcp-api-style` is in PATH
3. Ensure the file matches OpenAPI patterns

### Timeout Issues

Increase the timeout in the generated config:

```json
{
  "timeout": 60
}
```

### Profile Not Found

Ensure the profile name is valid:

```bash
mcp-api-style list-profiles
```

## Next Steps

- [MCP Server](../guide/mcp-server.md) - Full MCP integration
- [CLI Reference](cli.md) - Command documentation

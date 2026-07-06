# Minimal Profile

The minimal profile provides basic API hygiene rules based on common REST best practices. A lightweight option for simple APIs or getting started.

## Overview

| Property | Value |
|----------|-------|
| Name | `minimal` |
| Version | 1.0.0 |
| Rules | 29 |
| Categories | 7 |
| Focus | Basic API hygiene |

## Philosophy

The minimal profile emphasizes:

- **Simplicity** - Only essential rules
- **Quick Adoption** - Easy to pass, low barrier to entry
- **Foundation** - Starting point for custom profiles
- **Universal** - Applicable to any REST API

## Categories

| Category | Rules | Description |
|----------|-------|-------------|
| uri-design | 5 | URL structure and naming |
| http-methods | 5 | Proper HTTP method usage |
| naming | 4 | Property and parameter naming |
| errors | 3 | Error response format |
| responses | 3 | Response structure |
| documentation | 5 | API documentation |
| security | 4 | API security basics |

## All Rules

### URI Design

| ID | Title | Severity |
|----|-------|----------|
| URI-001 | Use plural resource names | error |
| URI-002 | Use kebab-case for path segments | error |
| URI-003 | Avoid verbs in paths | warn |
| URI-004 | Use path parameters for resource identifiers | warn |
| URI-005 | Limit path nesting depth | warn |

### HTTP Methods

| ID | Title | Severity |
|----|-------|----------|
| HTTP-001 | GET requests must not have request body | error |
| HTTP-002 | DELETE requests should not have request body | warn |
| HTTP-003 | POST should return 201 for resource creation | warn |
| HTTP-004 | Use PATCH for partial updates | info |
| HTTP-005 | Operations must have unique operationId | error |

### Naming Conventions

| ID | Title | Severity |
|----|-------|----------|
| NAMING-001 | Use camelCase for JSON properties | error |
| NAMING-002 | Use camelCase for query parameters | warn |
| NAMING-003 | Boolean properties should use is/has/can prefix | info |
| NAMING-004 | Use consistent date/time property naming | info |

### Error Handling

| ID | Title | Severity |
|----|-------|----------|
| ERR-001 | Define error responses for operations | warn |
| ERR-002 | Use consistent error response schema | warn |
| ERR-003 | Include 401 response for authenticated endpoints | warn |

### Responses

| ID | Title | Severity |
|----|-------|----------|
| RESP-001 | Responses must have descriptions | error |
| RESP-002 | Success responses should define content schema | warn |
| RESP-003 | Use consistent pagination format | info |

### Documentation

| ID | Title | Severity |
|----|-------|----------|
| DOC-001 | API must have info description | error |
| DOC-002 | Operations must have summary or description | warn |
| DOC-003 | Parameters must have descriptions | warn |
| DOC-004 | Schema properties should have descriptions | info |
| DOC-005 | Provide examples for schemas | info |

### Security

| ID | Title | Severity |
|----|-------|----------|
| SEC-001 | Define security schemes | warn |
| SEC-002 | Apply security to operations | warn |
| SEC-003 | Avoid API keys in URL | error |
| SEC-004 | Use HTTPS for server URLs | error |

## Usage

```bash
# Lint with minimal profile
api-style lint openapi.yaml --profile minimal

# Quick check for basic hygiene
api-style lint openapi.yaml --profile minimal --level minimum

# List all minimal rules
api-style list-rules --profile minimal
```

## When to Use

Use the minimal profile when:

- Getting started with API linting
- Building simple internal APIs
- Want a low-friction starting point
- Creating a base for a custom profile

## Comparison to Default

| Aspect | Minimal | Default |
|--------|---------|---------|
| Rules | 29 | 106 |
| Categories | 7 | 27 |
| Focus | Basics | SDK-optimized |
| Learning curve | Low | Medium |
| Strictness | Relaxed | Strict |

## Extending

Build on the minimal profile:

```json
{
  "name": "my-extended-minimal",
  "extends": ["minimal"],
  "rules": [
    {
      "id": "CUSTOM-001",
      "title": "My additional rule",
      "category": "naming",
      "severity": "warn"
    }
  ]
}
```

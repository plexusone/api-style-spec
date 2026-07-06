# Comprehensive Profile

The comprehensive profile provides 100% category coverage synthesized from Microsoft REST, Microsoft Graph, Zalando, Google Cloud, Azure, and PayPal guidelines.

## Overview

| Property | Value |
|----------|-------|
| Name | `comprehensive` |
| Version | 1.0.0 |
| Rules | 88 |
| Categories | 26 |
| Focus | Full coverage across all API design areas |

## Philosophy

The comprehensive profile emphasizes:

- **Complete Coverage** - Rules for every API design category
- **Industry Synthesis** - Best practices from major API providers
- **Balanced Strictness** - Mix of error, warn, and info severities
- **Educational** - Learn API design by understanding all categories

## Categories

| Category | Description |
|----------|-------------|
| general | OpenAPI specification requirements |
| naming | Property and parameter naming conventions |
| urls | URL structure and path design |
| http-methods | HTTP method semantics |
| http-status | Status code usage |
| request-response | Request/response patterns |
| headers | HTTP header conventions |
| errors | Error response format |
| pagination | Collection pagination patterns |
| filtering | Query parameter filtering |
| versioning | API versioning strategy |
| compatibility | Breaking change prevention |
| deprecation | Deprecation patterns |
| security | Authentication and authorization |
| documentation | API documentation requirements |
| json | JSON formatting conventions |
| schema | Schema design patterns |
| collections | Collection resource patterns |
| long-running | Async operation patterns |
| conditional | Conditional request handling |
| performance | Performance optimization |
| hypermedia | HATEOAS patterns |
| batch | Batch operation patterns |
| events | Event/webhook patterns |
| actions | Non-CRUD action patterns |
| throttling | Rate limiting patterns |

## Key Rules

### General

| ID | Title | Severity |
|----|-------|----------|
| GEN-001 | Provide OpenAPI specification | error |
| GEN-002 | Provide API info | error |
| GEN-003 | Use semantic versioning | warn |

### Naming

| ID | Title | Severity |
|----|-------|----------|
| NAME-001 | Use camelCase for properties | error |
| NAME-002 | Use plural nouns for collections | error |

### HTTP Methods

| ID | Title | Severity |
|----|-------|----------|
| HTTP-001 | GET must be safe and idempotent | error |
| HTTP-002 | PUT must be idempotent | error |
| HTTP-003 | DELETE must be idempotent | error |

### Errors

| ID | Title | Severity |
|----|-------|----------|
| ERR-001 | Define standard error schema | error |
| ERR-002 | Include error code in responses | warn |
| ERR-003 | Provide actionable error messages | info |

## Usage

```bash
# Lint with comprehensive profile
api-style lint openapi.yaml --profile comprehensive

# Analyze for full compliance
api-style analyze openapi.yaml --profile comprehensive

# List all rules
api-style list-rules --profile comprehensive

# Filter by category
api-style list-rules --profile comprehensive --category errors
```

## When to Use

Use the comprehensive profile when:

- Learning API design principles
- Auditing existing APIs for completeness
- Building internal API style guides
- Need coverage across all design areas

## Comparison to Default

| Aspect | Default | Comprehensive |
|--------|---------|---------------|
| Rules | 106 | 88 |
| Categories | 27 | 26 |
| Focus | SDK-optimized | Full coverage |
| SDK patterns | High priority | Balanced |
| Multi-tenancy | Yes | No |

## Sources

Rules synthesized from:

- Microsoft REST API Guidelines
- Microsoft Graph API Guidelines
- Zalando RESTful API Guidelines
- Google Cloud API Design Guide
- Azure REST API Guidelines
- PayPal API Standards

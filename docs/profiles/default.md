# Default Profile

The default profile is an industry-leading API style guide optimized for SDK generation with tools like ogen and openapi-generator. It synthesizes best practices from Stripe, GitHub, Microsoft, and Zalando.

## Overview

| Property | Value |
|----------|-------|
| Name | `default` |
| Version | 2.3.0 |
| Rules | 106 |
| Categories | 27 |
| Focus | SDK-optimized API design |

## Philosophy

The default profile emphasizes:

- **SDK Generation** - Clean generated code with discriminated unions, named schemas, explicit nullability
- **Consistency** - Uniform patterns for operationIds, error handling, pagination
- **Multi-tenancy** - Support for `~` alias patterns for current user/org context
- **Modern Standards** - RFC 9457 Problem Details, OpenAPI 3.1

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
| errors | Error response format (RFC 9457) |
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
| multi-tenancy | Multi-tenant API patterns |
| throttling | Rate limiting patterns |

## Key Rules

### General (PO-001)

| ID | Title | Severity |
|----|-------|----------|
| PO-001 | Provide OpenAPI 3.1 specification | error |

### Naming (PO-002)

| ID | Title | Severity |
|----|-------|----------|
| PO-002 | Use camelCase for property names | error |

### URLs (PO-003, PO-004)

| ID | Title | Severity |
|----|-------|----------|
| PO-003 | Use kebab-case for URL paths | error |
| PO-004 | Use plural nouns for collection resources | error |

### Versioning (PO-005)

| ID | Title | Severity |
|----|-------|----------|
| PO-005 | Version APIs with URI prefix | error |

### Errors (PO-024)

| ID | Title | Severity |
|----|-------|----------|
| PO-024 | Use RFC 9457 Problem Details for errors | error |

### Multi-tenancy (PO-103, PO-104, PO-105)

| ID | Title | Severity |
|----|-------|----------|
| PO-103 | Support ~ alias for current context | warn |
| PO-104 | Provide tenant-scoped endpoints | warn |
| PO-105 | Document tenant isolation model | info |

## Conformance Levels

### Minimum

Basic API structure and methods.

- Required rules: Core naming, URL structure, HTTP methods
- Focus: API hygiene

### Standard

Production-ready APIs.

- Includes Minimum requirements
- Required rules: Error handling, documentation, security
- Focus: Deployable APIs

### Exemplary

Best-in-class APIs.

- Includes Standard requirements
- Required rules: All categories
- Focus: SDK-optimized, fully documented

## Usage

```bash
# Lint with default profile
api-style lint openapi.yaml

# Lint with specific conformance level
api-style lint openapi.yaml --level standard

# Combined lint + LLM evaluation
api-style analyze openapi.yaml

# Generate human-readable style guide
api-style generate guide --output docs/

# List all rules
api-style list-rules --profile default
```

## Enforcement Stats

- **Deterministic (Spectral)**: ~34% of rules
- **LLM-evaluable**: 100% of rules have judge criteria
- **SDK-optimized**: Focus on ogen, openapi-generator compatibility

## Comparison to Other Profiles

| Aspect | Default | Azure | Zalando |
|--------|---------|-------|---------|
| Rules | 106 | 23 | 147 |
| Versioning | URI prefix | Date-based | Header |
| Errors | RFC 9457 | Azure format | Problem+JSON |
| Pagination | Cursor-based | OData | Cursor-based |
| SDK Focus | High | Medium | Medium |

## Customization

Extend the default profile with custom rules:

```json
{
  "name": "my-style",
  "extends": ["default"],
  "rules": [
    {
      "id": "CUSTOM-001",
      "title": "My custom rule",
      "category": "naming",
      "severity": "warn"
    }
  ],
  "overrides": {
    "PO-003": { "severity": "warn" }
  }
}
```

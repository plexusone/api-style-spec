# Default Profile

The default profile provides common REST API best practices applicable to most API projects.

## Overview

| Property | Value |
|----------|-------|
| Name | `default` |
| Version | 1.0.0 |
| Rules | 29 |
| Focus | General REST API design |

## Categories

| Category | Rules | Description |
|----------|-------|-------------|
| URI Design | 5 | URL structure and naming |
| HTTP Methods | 5 | Proper HTTP method usage |
| Naming | 4 | Property and parameter naming |
| Errors | 3 | Error response format |
| Responses | 3 | Response structure |
| Documentation | 5 | API documentation |
| Security | 4 | API security |

## Key Rules

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

## Conformance Levels

### Bronze (Minimum Viable)

Basic structure and methods.

- Required categories: `uri-design`, `http-methods`
- Max errors: 0
- Max warnings: 10

### Silver (Production-Ready)

Proper naming and error handling.

- Includes Bronze requirements
- Required categories: `naming`, `errors`
- Max errors: 0
- Max warnings: 5

### Gold (Exemplary)

Complete documentation and security.

- Includes Silver requirements
- Required categories: `documentation`, `security`
- Max errors: 0
- Max warnings: 0

## Usage

```bash
# Lint with default profile
api-style lint openapi.yaml --profile default

# Check silver conformance
api-style lint openapi.yaml --profile default --level silver

# List all rules
api-style list-rules --profile default
```

## Customization

Extend the default profile with custom rules:

```json
{
  "name": "my-style",
  "extends": ["default"],
  "rules": [
    // Additional rules
  ],
  "overrides": {
    "URI-001": { "severity": "warn" }
  }
}
```

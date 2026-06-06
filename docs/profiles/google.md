# Google Profile

The Google profile implements Google's API Design Guide for resource-oriented APIs.

## Overview

| Property | Value |
|----------|-------|
| Name | `google` |
| Based On | [Google API Design Guide](https://cloud.google.com/apis/design) |
| Rules | ~20 |
| Focus | Resource-oriented design |

## Philosophy

Google's API guidelines emphasize:

- **Resource-Oriented Design** - Everything is a resource
- **Standard Methods** - Consistent CRUD operations
- **Predictability** - Developers can guess API behavior

## Key Principles

### Resource-Oriented Design

APIs are modeled as resource hierarchies:

```
/projects/{project}/locations/{location}/instances/{instance}
```

### Standard Methods

| Method | HTTP | Description |
|--------|------|-------------|
| List | GET collection | List resources |
| Get | GET resource | Get single resource |
| Create | POST collection | Create resource |
| Update | PUT/PATCH resource | Update resource |
| Delete | DELETE resource | Delete resource |

### Naming Conventions

- Collection IDs are plural (e.g., `users`, `projects`)
- Resource IDs are singular (e.g., `user-123`)
- Field names use `snake_case`

### Error Model

```json
{
  "error": {
    "code": 404,
    "message": "Resource not found",
    "status": "NOT_FOUND",
    "details": []
  }
}
```

## Categories

| Category | Focus |
|----------|-------|
| Resources | Resource hierarchy and naming |
| Methods | Standard method patterns |
| Fields | Field naming conventions |
| Errors | Google error model |
| Documentation | API documentation standards |

## Notable Rules

### GCP-RES-001: Use Resource Hierarchy

```
Good: /projects/{projectId}/instances/{instanceId}
Bad:  /getInstanceById?id={instanceId}
```

### GCP-FIELD-001: Use snake_case for Fields

```json
Good: {"user_name": "john", "created_at": "..."}
Bad:  {"userName": "john", "createdAt": "..."}
```

### GCP-METHOD-001: Standard Method Mapping

| Operation | HTTP Method | URL Pattern |
|-----------|-------------|-------------|
| List | GET | /users |
| Get | GET | /users/{id} |
| Create | POST | /users |
| Update | PATCH | /users/{id} |
| Delete | DELETE | /users/{id} |

### GCP-ERR-001: Google Error Format

```json
{
  "error": {
    "code": 400,
    "message": "Invalid argument",
    "status": "INVALID_ARGUMENT",
    "details": [
      {
        "@type": "type.googleapis.com/google.rpc.BadRequest",
        "fieldViolations": [...]
      }
    ]
  }
}
```

## Usage

```bash
# Lint with Google profile
api-style lint openapi.yaml --profile google

# List Google rules
api-style list-rules --profile google
```

## When to Use

Use the Google profile when:

- Building Google Cloud Platform APIs
- Following resource-oriented design patterns
- Prefer `snake_case` field naming
- Working with gRPC/Protobuf (field naming consistency)

## Comparison to Default

| Aspect | Default | Google |
|--------|---------|--------|
| Field case | camelCase | snake_case |
| Design | REST | Resource-oriented |
| Methods | HTTP verbs | Standard methods |
| Errors | RFC 7807 | Google error model |

## References

- [Google API Design Guide](https://cloud.google.com/apis/design)
- [Google Cloud API Style Guide](https://cloud.google.com/apis/design/naming_convention)
- [API Improvement Proposals (AIPs)](https://google.aip.dev/)

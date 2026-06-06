# Azure Profile

The Azure profile implements Microsoft's REST API Guidelines for enterprise cloud services.

## Overview

| Property | Value |
|----------|-------|
| Name | `azure` |
| Based On | [Microsoft REST API Guidelines](https://github.com/microsoft/api-guidelines) |
| Rules | ~25 |
| Focus | Enterprise cloud APIs |

## Philosophy

Microsoft's API guidelines emphasize:

- **Consistency** - Uniform interface across all Azure services
- **Compatibility** - Long-term API stability and evolution
- **Clarity** - Self-documenting APIs with clear semantics

## Key Principles

### Versioning

- Date-based API versions (e.g., `2023-01-01`)
- Version in URL path or query parameter
- Support multiple versions simultaneously

### Error Handling

- Standard error response format
- Error codes for programmatic handling
- Detailed error messages for debugging

### Pagination

- Consistent `value` array for collections
- `nextLink` for continuation
- `@odata.count` for total count

### Long-Running Operations

- 202 Accepted for async operations
- `Operation-Location` header for status polling
- Consistent operation status schema

## Categories

| Category | Focus |
|----------|-------|
| Versioning | Date-based version management |
| URI Design | Azure resource hierarchy |
| Responses | Azure response patterns |
| Errors | Microsoft error format |
| Pagination | OData-style pagination |
| Long-Running | Async operation patterns |

## Notable Rules

### AZ-VER-001: Use Date-Based Versioning

```
Good: /api/2023-07-01/users
Bad:  /api/v1/users
```

### AZ-ERR-001: Use Azure Error Format

```json
{
  "error": {
    "code": "ResourceNotFound",
    "message": "The specified resource does not exist.",
    "target": "subscriptionId",
    "details": [],
    "innererror": {}
  }
}
```

### AZ-PAG-001: Use OData-Style Pagination

```json
{
  "value": [...],
  "@odata.count": 100,
  "nextLink": "https://api.example.com/users?$skiptoken=xxx"
}
```

### AZ-LRO-001: Long-Running Operation Pattern

```
POST /subscriptions/{id}/resourceGroups
Response: 202 Accepted
Operation-Location: /operations/{operationId}
```

## Usage

```bash
# Lint with Azure profile
api-style lint openapi.yaml --profile azure

# Analyze for Azure compliance
api-style analyze openapi.yaml --profile azure
```

## When to Use

Use the Azure profile when:

- Building APIs for Azure services
- Following Microsoft enterprise standards
- Need consistent async operation patterns
- Working with Azure-compatible tooling

## Comparison to Default

| Aspect | Default | Azure |
|--------|---------|-------|
| Versioning | Semantic (v1) | Date-based |
| Case style | kebab-case URLs | camelCase fields |
| Pagination | Generic | OData-style |
| Errors | RFC 7807 | Azure error format |
| Async | Not specified | LRO pattern |

## References

- [Microsoft REST API Guidelines](https://github.com/microsoft/api-guidelines)
- [Azure REST API Guidelines](https://azure.github.io/azure-sdk/general_design.html)
- [Azure API Design Best Practices](https://docs.microsoft.com/azure/architecture/best-practices/api-design)

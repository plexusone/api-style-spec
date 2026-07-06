# Microsoft Graph Profile

The Microsoft Graph profile implements rules based on Microsoft Graph API Guidelines for building APIs compatible with Microsoft 365 services.

## Overview

| Property | Value |
|----------|-------|
| Name | `microsoft-graph` |
| Version | 1.0.0 |
| Rules | 82 |
| Categories | 12 |
| Focus | OData-based APIs for Microsoft 365 integration |

## Philosophy

Microsoft Graph API Guidelines emphasize:

- **OData Conventions** - Standard query parameters and response formats
- **Navigation Properties** - Entity relationships via navigation
- **Actions and Functions** - Non-CRUD operations with clear semantics
- **Delta Queries** - Efficient change tracking

## Categories

| Category | Description |
|----------|-------------|
| uri-design | Graph endpoint structure and versioning |
| odata | OData query parameters and conventions |
| navigation | Navigation properties and relationships |
| actions-functions | Non-CRUD operations |
| delta | Change tracking and delta queries |
| batch | Batch request patterns |
| permissions | Scope-based authorization |
| throttling | Rate limiting and retry patterns |
| extensions | Open and schema extensions |
| webhooks | Change notifications |
| naming | Entity and property naming |
| types | Type definitions and enums |

## Key Rules

### URI Design

| ID | Title | Severity |
|----|-------|----------|
| GRAPH-URI-001 | Use graph.microsoft.com endpoint | error |
| GRAPH-URI-002 | Support /me alias | warn |
| GRAPH-URI-003 | Use v1.0 for production, beta for preview | error |

### OData Conventions

| ID | Title | Severity |
|----|-------|----------|
| GRAPH-ODATA-001 | Support standard OData query parameters | error |
| GRAPH-ODATA-002 | Use OData filter syntax | warn |
| GRAPH-ODATA-003 | Support $select for field selection | error |
| GRAPH-ODATA-004 | Support $expand for relationships | warn |
| GRAPH-ODATA-005 | Support $orderby for sorting | warn |

### Navigation Properties

| ID | Title | Severity |
|----|-------|----------|
| GRAPH-NAV-001 | Define navigation properties for relationships | error |
| GRAPH-NAV-002 | Support $ref for relationship management | warn |
| GRAPH-NAV-003 | Use consistent navigation property naming | warn |

### Delta Queries

| ID | Title | Severity |
|----|-------|----------|
| GRAPH-DELTA-001 | Support delta function for change tracking | warn |
| GRAPH-DELTA-002 | Include @odata.deltaLink in responses | warn |
| GRAPH-DELTA-003 | Handle delta token expiration | info |

### Actions and Functions

| ID | Title | Severity |
|----|-------|----------|
| GRAPH-ACT-001 | Use POST for actions, GET for functions | error |
| GRAPH-ACT-002 | Prefix actions with microsoft.graph | warn |
| GRAPH-ACT-003 | Document action parameters | warn |

## OData Query Examples

```
# Select specific fields
GET /users?$select=displayName,mail

# Filter results
GET /users?$filter=startswith(displayName,'J')

# Expand relationships
GET /users/{id}?$expand=memberOf

# Order results
GET /users?$orderby=displayName desc

# Pagination
GET /users?$top=10&$skip=20
```

## Delta Query Pattern

```
# Initial sync
GET /users/delta
Response: {
  "value": [...],
  "@odata.deltaLink": "https://graph.microsoft.com/v1.0/users/delta?$deltatoken=xxx"
}

# Get changes
GET /users/delta?$deltatoken=xxx
Response: {
  "value": [/* changed items */],
  "@odata.deltaLink": "..."
}
```

## Usage

```bash
# Lint with Microsoft Graph profile
api-style lint openapi.yaml --profile microsoft-graph

# Analyze for Graph compliance
api-style analyze openapi.yaml --profile microsoft-graph

# List OData rules
api-style list-rules --profile microsoft-graph --category odata

# List all rules
api-style list-rules --profile microsoft-graph
```

## When to Use

Use the Microsoft Graph profile when:

- Building APIs that integrate with Microsoft 365
- Implementing OData-compliant services
- Need navigation properties and relationships
- Supporting delta queries for sync scenarios
- Building Graph-compatible connectors

## Comparison to Microsoft REST

| Aspect | Microsoft Graph | Microsoft REST |
|--------|-----------------|----------------|
| Rules | 82 | 123 |
| Focus | OData/Graph | General REST |
| Versioning | v1.0/beta | YYYY-MM-DD |
| Query style | OData | Custom |
| Delta support | Yes | No |

## References

- [Microsoft Graph API Guidelines](https://github.com/microsoftgraph/msgraph-sdk-design)
- [Microsoft Graph Documentation](https://docs.microsoft.com/graph)
- [OData Specification](https://www.odata.org/documentation/)

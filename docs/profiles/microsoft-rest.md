# Microsoft REST Profile

The Microsoft REST profile implements comprehensive rules based on Microsoft Azure REST API Guidelines and Microsoft REST API Guidelines.

## Overview

| Property | Value |
|----------|-------|
| Name | `microsoft-rest` |
| Version | 1.0.0 |
| Rules | 123 |
| Categories | 15 |
| Focus | Enterprise REST APIs following Microsoft patterns |

## Philosophy

Microsoft REST API Guidelines emphasize:

- **Consistency** - Uniform interface across all services
- **Compatibility** - Long-term API stability and evolution
- **Date-based Versioning** - YYYY-MM-DD version format
- **Long-Running Operations** - Standardized async patterns

## Categories

| Category | Description |
|----------|-------------|
| versioning | Date-based version management (api-version parameter) |
| uri-design | Resource hierarchy and path structure |
| http-methods | HTTP method semantics and idempotency |
| request-response | Request/response patterns |
| errors | Microsoft error format |
| pagination | OData-style pagination |
| filtering | OData query parameters |
| long-running | Async operation patterns (LRO) |
| naming | Property and parameter naming |
| security | Authentication and authorization |
| headers | HTTP header conventions |
| schema-design | Schema design patterns |
| collections | Collection resource patterns |
| conditional | ETags and conditional requests |
| compatibility | Breaking change prevention |

## Key Rules

### Versioning

| ID | Title | Severity |
|----|-------|----------|
| MS-VER-001 | Use api-version query parameter | error |
| MS-VER-002 | Version format YYYY-MM-DD | error |
| MS-VER-003 | Support multiple API versions | warn |
| MS-VER-004 | Document version differences | info |

### URI Design

| ID | Title | Severity |
|----|-------|----------|
| MS-URI-001 | Use lowercase paths | error |
| MS-URI-002 | Use camelCase for path segments | error |
| MS-URI-003 | Resource hierarchy reflects ownership | warn |

### Long-Running Operations

| ID | Title | Severity |
|----|-------|----------|
| MS-LRO-001 | Return 202 for async operations | error |
| MS-LRO-002 | Include Operation-Location header | error |
| MS-LRO-003 | Provide operation status endpoint | error |
| MS-LRO-004 | Support operation cancellation | warn |

### Error Handling

| ID | Title | Severity |
|----|-------|----------|
| MS-ERR-001 | Use Microsoft error format | error |
| MS-ERR-002 | Include error code | error |
| MS-ERR-003 | Include error message | error |
| MS-ERR-004 | Support innererror for debugging | warn |

### Pagination

| ID | Title | Severity |
|----|-------|----------|
| MS-PAG-001 | Use value array for collections | error |
| MS-PAG-002 | Include nextLink for continuation | error |
| MS-PAG-003 | Support @odata.count | warn |

## Error Response Format

```json
{
  "error": {
    "code": "ResourceNotFound",
    "message": "The specified resource does not exist.",
    "target": "subscriptionId",
    "details": [],
    "innererror": {
      "code": "InternalCode",
      "message": "Additional details"
    }
  }
}
```

## Long-Running Operation Pattern

```
POST /subscriptions/{id}/resourceGroups
Response: 202 Accepted
Operation-Location: /operations/{operationId}
Retry-After: 30

GET /operations/{operationId}
Response: 200 OK
{
  "status": "Running" | "Succeeded" | "Failed",
  "percentComplete": 50
}
```

## Usage

```bash
# Lint with Microsoft REST profile
api-style lint openapi.yaml --profile microsoft-rest

# Analyze for Microsoft compliance
api-style analyze openapi.yaml --profile microsoft-rest

# List all rules
api-style list-rules --profile microsoft-rest

# Filter by category
api-style list-rules --profile microsoft-rest --category long-running
```

## When to Use

Use the Microsoft REST profile when:

- Building APIs for Azure services
- Following Microsoft enterprise standards
- Need standardized async operation patterns
- Working with Microsoft-compatible tooling
- Building enterprise B2B APIs

## Comparison to Azure Profile

| Aspect | Microsoft REST | Azure |
|--------|----------------|-------|
| Rules | 123 | 23 |
| Scope | Comprehensive | Focused |
| Categories | 15 | 9 |
| Detail level | Exhaustive | Essential |

## Evaluation Report

This profile has been evaluated against the `api-style-guide-quality` rubric.

### Summary

| Metric | Value |
|--------|-------|
| Overall Decision | **PASS** |
| Categories | 9 pass, 0 partial, 0 fail |
| Findings | 0 critical, 0 high, 0 medium, 3 low |

### Category Scores

| Category | Score | Assessment |
|----------|-------|------------|
| Content Coverage | 5/5 🟢 | Covers all 6 domains plus enterprise patterns: REST/OData, OpenAPI, camelCase naming, structured error envelope, OAuth2/RBAC, api-version parameter |
| Structure & Navigation | 5/5 🟢 | Excellent structure: TOC, 15 categories, MS-XXX rule IDs, conformance levels, design patterns |
| Rule Quality & Clarity | 5/5 🟢 | Clear titles, detailed descriptions, severity levels, comprehensive rationale with RFC 2119 keywords |
| Examples & Code Samples | 5/5 🟢 | HTTP request/response samples, mermaid sequence diagrams for LRO/pagination flows, good/bad contrasts |
| Enforceability & Tooling | 4/5 🟡 | Distinguishes spectral-lintable vs. judge-only rules; some rules lack explicit Spectral configurations |
| Guide Versioning & Evolution | 4/5 🟡 | Versioned on GitHub with compatibility section; could improve rule-level change tracking |
| Completeness & Depth | 5/5 🟢 | Exceptional depth: LRO state machines, conditional requests, ETag strategies, pagination patterns |
| Internal Consistency | 5/5 🟢 | Consistent terminology, value/nextLink pagination pattern, standardized error format |
| Accessibility & Tone | 4/5 🟡 | Professional tone appropriate for enterprise developers; some Azure-specific terms need glossary expansion |

### Improvement Opportunities

| Finding | Category | Recommendation |
|---------|----------|----------------|
| Some rules lack explicit Spectral configurations | Enforceability | Publish complete Spectral ruleset |
| No per-rule version or deprecation tracking | Versioning | Add version field to each rule |
| Azure-specific terminology may confuse non-Azure developers | Accessibility | Expand glossary with Azure concepts |

!!! note "Evaluation Metadata"
    - **Rubric**: api-style-guide-quality v1.0.0
    - **Evaluated**: 2025-06-17
    - **Evaluator**: Claude Opus 4.5 (LLM-as-Judge)

## References

- [Microsoft REST API Guidelines](https://github.com/microsoft/api-guidelines)
- [Azure REST API Guidelines](https://azure.github.io/azure-sdk/general_design.html)
- [Microsoft Cloud API Design](https://docs.microsoft.com/azure/architecture/best-practices/api-design)

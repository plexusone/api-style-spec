# Zalando Profile

The Zalando profile implements Zalando's RESTful API Guidelines for pragmatic API design.

## Overview

| Property | Value |
|----------|-------|
| Name | `zalando` |
| Based On | [Zalando RESTful API Guidelines](https://opensource.zalando.com/restful-api-guidelines/) |
| Rules | 55 |
| Categories | 13 |
| Patterns | 3 |
| Focus | Pragmatic e-commerce APIs |

## Philosophy

Zalando's API guidelines emphasize:

- **Pragmatic REST** - Practical over dogmatic
- **API First** - Design before implementation
- **Evolution** - Compatible changes over breaking changes

## Key Principles

### API First Development

- Design OpenAPI spec before coding
- Use spec for documentation and validation
- Generate client SDKs from spec

### Hypermedia (HATEOAS)

- Include links in responses where helpful
- Use standard link relations
- Don't require hypermedia for basic operations

### Compatibility

- Additive changes are safe
- Never remove or rename fields
- Use deprecation before removal

### Events

- Async events complement REST
- Use CloudEvents format
- Event-driven architecture patterns

## Categories

| Category | Rules | Focus |
|----------|-------|-------|
| General | 4 | API-first principles |
| Compatibility | 4 | Evolution and versioning |
| JSON | 5 | Property naming and formats |
| Data Formats | 3 | Date/time, enums, money |
| URLs | 4 | URL naming and structure |
| HTTP Methods | 6 | Method semantics |
| HTTP Status | 5 | Response codes |
| Headers | 3 | Standard headers |
| Hypermedia | 3 | HATEOAS patterns |
| Pagination | 4 | Cursor-based patterns |
| Performance | 4 | Caching, compression |
| Security | 4 | OAuth2, TLS |
| Deprecation | 3 | Sunset headers |

## Notable Rules

### ZAL-URL-001: Use kebab-case for URLs

```
Good: /shopping-carts/{cart-id}/line-items
Bad:  /shoppingCarts/{cartId}/lineItems
```

### ZAL-JSON-001: Use snake_case for Properties

```json
Good: {"order_id": "123", "created_at": "..."}
Bad:  {"orderId": "123", "createdAt": "..."}
```

### ZAL-ERR-001: Use RFC 7807 Problem Details

```json
{
  "type": "https://api.zalando.com/problems/out-of-stock",
  "title": "Product Out of Stock",
  "status": 422,
  "detail": "Product SKU-123 is currently unavailable",
  "instance": "/orders/456"
}
```

### ZAL-PAG-001: Cursor-Based Pagination

```json
{
  "items": [...],
  "cursor": "eyJpZCI6MTIzfQ==",
  "self": "https://api.example.com/orders?cursor=abc",
  "next": "https://api.example.com/orders?cursor=def"
}
```

### ZAL-COMPAT-001: Compatibility Rules

- **MUST** add new fields as optional
- **MUST NOT** remove or rename fields
- **SHOULD** use deprecation headers
- **MAY** version via URL for breaking changes

## Usage

```bash
# Lint with Zalando profile
api-style lint openapi.yaml --profile zalando

# Analyze for Zalando compliance
api-style analyze openapi.yaml --profile zalando
```

## Generate Documentation

Generate a human-readable style guide from the Zalando profile:

```bash
# Single-page Markdown (25KB+)
api-style generate guide --profile zalando --output zalando-guide.md

# MkDocs site (20 pages)
api-style generate mkdocs --profile zalando --output ./zalando-docs

# Build and serve MkDocs site
cd zalando-docs && pip install mkdocs-material && mkdocs serve
```

The generated documentation includes all 55 rules organized by category, design patterns (cursor pagination, Problem+JSON), principles, and glossary.

## When to Use

Use the Zalando profile when:

- Building e-commerce APIs
- Prefer `snake_case` field naming
- Want pragmatic REST over strict REST
- Need evolution-friendly API design
- Using event-driven patterns

## Comparison to Default

| Aspect | Default | Zalando |
|--------|---------|---------|
| URL case | kebab-case | kebab-case |
| Field case | camelCase | snake_case |
| Pagination | Generic | Cursor-based |
| Errors | Generic | RFC 7807 |
| Hypermedia | Not required | Encouraged |

## Evaluation Report

This profile has been evaluated against the `api-style-guide-quality` rubric.

### Summary

| Metric | Value |
|--------|-------|
| Overall Decision | **PASS** |
| Categories | 9 pass, 0 partial, 0 fail |
| Findings | 0 critical, 0 high, 0 medium, 4 low |

### Category Scores

| Category | Score | Assessment |
|----------|-------|------------|
| Content Coverage | 5/5 🟢 | Covers all 6 domains: REST/API First, OpenAPI 3.1, snake_case naming, RFC 7807 errors, OAuth2 security, semantic versioning |
| Structure & Navigation | 5/5 🟢 | Clear TOC, numbered rules (Z-100 through Z-255), 13 logical categories, cross-references |
| Rule Quality & Clarity | 5/5 🟢 | Well-structured with severity levels, rationale, RFC 2119 keywords (MUST/SHOULD/MAY) |
| Examples & Code Samples | 4/5 🟡 | Good/bad examples for naming, errors, pagination; ~40% of rules lack examples |
| Enforceability & Tooling | 5/5 🟢 | Zally linter provides automated enforcement; clear spectral vs. judge-only distinction |
| Guide Versioning & Evolution | 4/5 🟡 | Version numbers and GitHub history; no changelog showing rule evolution |
| Completeness & Depth | 5/5 🟢 | Deep coverage of pagination, error handling, compatibility, deprecation timelines |
| Internal Consistency | 5/5 🟢 | Consistent terminology; examples align with rules; no contradictions |
| Accessibility & Tone | 4/5 🟡 | Appropriate technical level; glossary limited to 8 terms for 147 rules |

### Improvement Opportunities

| Finding | Category | Recommendation |
|---------|----------|----------------|
| ~40% of rules lack concrete examples | Examples | Add JSON/HTTP examples to rules without them |
| Some rules lack detailed implementation guidance | Rule Quality | Add step-by-step HOW guidance for complex rules |
| No changelog showing rule additions/changes | Versioning | Add CHANGELOG.md tracking rule evolution |
| Glossary has only 8 terms | Accessibility | Expand glossary to 20+ key concepts |

!!! note "Evaluation Metadata"
    - **Rubric**: api-style-guide-quality v1.0.0
    - **Evaluated**: 2025-06-17
    - **Evaluator**: Claude Opus 4.5 (LLM-as-Judge)

## References

- [Zalando RESTful API Guidelines](https://opensource.zalando.com/restful-api-guidelines/)
- [API Linter (Zally)](https://github.com/zalando/zally)
- [Problem Details RFC 7807](https://tools.ietf.org/html/rfc7807)

# Zalando Profile

The Zalando profile implements Zalando's RESTful API Guidelines for pragmatic API design.

## Overview

| Property | Value |
|----------|-------|
| Name | `zalando` |
| Based On | [Zalando RESTful API Guidelines](https://opensource.zalando.com/restful-api-guidelines/) |
| Rules | ~30 |
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

| Category | Focus |
|----------|-------|
| General | API-first principles |
| URLs | URL naming and structure |
| JSON | JSON property guidelines |
| HTTP | HTTP method and status usage |
| Errors | Problem details (RFC 7807) |
| Pagination | Cursor-based pagination |
| Hypermedia | HATEOAS patterns |
| Compatibility | Evolution strategies |

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

## References

- [Zalando RESTful API Guidelines](https://opensource.zalando.com/restful-api-guidelines/)
- [API Linter (Zally)](https://github.com/zalando/zally)
- [Problem Details RFC 7807](https://tools.ietf.org/html/rfc7807)

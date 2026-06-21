# PlexusOne REST API Guidelines

Opinionated best practices for REST APIs, synthesized from Microsoft, Zalando, and PayPal guidelines. Uses camelCase for properties, kebab-case for URLs, and RFC 7807 for errors.

## Metadata

- **Author:** PlexusOne
- **Last Updated:** 2025-06-18
- **License:** Apache-2.0

## Table of Contents

- [Conformance Levels](#conformance-levels)
- **Rules**
  - [General Guidelines](#general-guidelines)
  - [Naming Conventions](#naming-conventions)
  - [URL Design](#url-design)
  - [HTTP Methods](#http-methods)
  - [HTTP Status Codes](#http-status-codes)
  - [Request & Response](#request-&-response)
  - [Error Handling](#error-handling)
  - [Pagination](#pagination)
  - [Filtering & Sorting](#filtering-&-sorting)
  - [Versioning](#versioning)
  - [Security](#security)
  - [Documentation](#documentation)
- [Glossary](#glossary)

## Conformance Levels

## General Guidelines

Fundamental requirements for all APIs

### PO-001: Provide OpenAPI 3.1 specification

**Severity:** error

All APIs MUST have an OpenAPI 3.1 specification. The spec is the source of truth and MUST be maintained alongside code.

**Rationale:** OpenAPI enables tooling, documentation generation, client SDK generation, and contract testing. Version 3.1 aligns with JSON Schema draft 2020-12.

**Examples:**

Good:

- `openapi: 3.1.0`

Bad:

- `swagger: 2.0`
- `No specification provided`

---

## Naming Conventions

Consistent naming for resources, properties, and URLs

### PO-002: Use camelCase for property names

**Severity:** error

All JSON property names MUST use camelCase (e.g., firstName, createdAt, userId).

**Rationale:** camelCase is native to JavaScript/JSON and aligns with OpenAPI conventions. It is the standard for Google, Microsoft, and GitHub APIs, and matches the PlexusOne codebase.

**Examples:**

Good:

- `firstName`
- `lastName`
- `createdAt`
- `userId`
- `isActive`

Bad:

- `first_name (snake_case)`
- `FirstName (PascalCase)`
- `FIRST_NAME (SCREAMING_SNAKE)`

---

## URL Design

Resource-oriented URL patterns

### PO-003: Use kebab-case for URL paths

**Severity:** error

URL path segments MUST use kebab-case (lowercase with hyphens).

**Rationale:** Kebab-case is the standard for URLs. It's readable and avoids issues with case sensitivity in different systems.

**Examples:**

Good:

- `/api/v1/user-accounts`
- `/api/v1/order-items/{orderId}`

Bad:

- `/api/v1/userAccounts`
- `/api/v1/user_accounts`
- `/api/v1/UserAccounts`

---

### PO-004: Use plural nouns for collection resources

**Severity:** error

Collection resource names MUST be plural nouns. Use /users not /user, /orders not /order.

**Rationale:** Plural nouns make it clear you're dealing with a collection. GET /users returns multiple users, POST /users creates one user in the collection.

**Examples:**

Good:

- `/users`
- `/orders`
- `/products`
- `/user-accounts`

Bad:

- `/user`
- `/order`
- `/product`
- `/userAccount`

---

## HTTP Methods

Proper use of HTTP verbs and semantics

### PO-008: Use standard HTTP methods correctly

**Severity:** error

Use HTTP methods according to their semantics: GET (read), POST (create), PUT (replace), PATCH (partial update), DELETE (remove).

**Rationale:** Correct HTTP method usage enables caching, idempotency guarantees, and predictable behavior.

**Examples:**

Good:

- `GET /v1/users/{id} - retrieve user`
- `POST /v1/users - create user`
- `PATCH /v1/users/{id} - update user fields`
- `DELETE /v1/users/{id} - delete user`

Bad:

- `POST /v1/users/{id}/delete`
- `GET /v1/users/create?name=John`

---

### PO-009: GET and DELETE must be idempotent

**Severity:** error

GET, PUT, and DELETE operations MUST be idempotent. Multiple identical requests must have the same effect as a single request.

**Rationale:** Idempotency enables safe retries and improves reliability. Clients can retry failed requests without causing duplicate operations.

---

### PO-038: Use PATCH for partial updates

**Severity:** error

Use PATCH (not PUT) for partial resource updates. Only send changed fields.

**Rationale:** PATCH is designed for partial updates. PUT replaces the entire resource and requires the client to send all fields.

**Examples:**

Good:

- `PATCH /v1/users/123 with {"name": "New Name"}`

Bad:

- `PUT /v1/users/123 with all fields for a name change`

---

## HTTP Status Codes

Appropriate status code usage

### PO-010: Return 201 Created with Location header for POST

**Severity:** error

Successful POST requests that create resources MUST return 201 Created with a Location header pointing to the new resource.

**Rationale:** 201 explicitly indicates creation. The Location header enables clients to fetch the created resource without parsing the response body.

**Examples:**

Good:

- `HTTP/1.1 201 Created\nLocation: /v1/users/12345`

Bad:

- `HTTP/1.1 200 OK (for creation)`

---

### PO-011: Use 204 No Content for successful DELETE

**Severity:** warn

Successful DELETE requests SHOULD return 204 No Content with an empty body.

**Rationale:** The resource no longer exists, so there's nothing to return. 204 clearly indicates success without a body.

---

### PO-012: Use 400 Bad Request for validation errors

**Severity:** error

Use 400 Bad Request for client errors: malformed syntax, invalid parameters, validation failures.

**Rationale:** 400 is the generic client error. It tells the client the request was wrong and should not be retried without modification.

---

### PO-013: Use 401 Unauthorized for missing/invalid authentication

**Severity:** error

Use 401 Unauthorized when authentication is required but not provided or invalid.

**Rationale:** 401 specifically indicates an authentication problem. The client should authenticate and retry.

---

### PO-014: Use 403 Forbidden for authorization failures

**Severity:** error

Use 403 Forbidden when the user is authenticated but not authorized to perform the action.

**Rationale:** 403 distinguishes authorization from authentication. The user is known but not permitted.

---

### PO-015: Use 404 Not Found for missing resources

**Severity:** error

Use 404 Not Found when the requested resource does not exist.

**Rationale:** 404 is clear and well-understood. Don't use 200 with an empty body or error message.

---

### PO-016: Use 409 Conflict for state conflicts

**Severity:** warn

Use 409 Conflict when the request conflicts with current resource state (e.g., duplicate creation, invalid state transition).

**Rationale:** 409 indicates a conflict that the client might resolve. Useful for optimistic locking and state machines.

---

### PO-017: Use 422 Unprocessable Entity for semantic errors

**Severity:** warn

Use 422 Unprocessable Entity when the request is syntactically correct but semantically invalid (business rule violations).

**Rationale:** 422 distinguishes semantic errors from syntax errors (400). The request was understood but cannot be processed.

---

### PO-018: Use 429 Too Many Requests for rate limiting

**Severity:** error

Use 429 Too Many Requests when rate limiting is applied. Include Retry-After header.

**Rationale:** 429 tells clients they're being throttled. Retry-After helps them back off appropriately.

---

## Request & Response

JSON payload structure and conventions

### PO-020: Use ISO 8601 for dates and times

**Severity:** error

All date and time values MUST use ISO 8601 format: YYYY-MM-DDTHH:mm:ssZ for timestamps, YYYY-MM-DD for dates.

**Rationale:** ISO 8601 is unambiguous and sortable. Always use UTC (Z suffix) for timestamps to avoid timezone confusion.

**Examples:**

Good:

- `2024-01-15T14:30:00Z`
- `2024-01-15`

Bad:

- `01/15/2024`
- `1705329000`
- `Jan 15, 2024`

---

### PO-021: Use string for IDs

**Severity:** error

Resource identifiers MUST be strings, not integers. Use UUIDs or prefixed IDs (e.g., userAbc123).

**Rationale:** String IDs are more flexible than integers. They support UUIDs, prefixed IDs, and don't have integer overflow issues. Prefixed IDs (userAbc123) are self-documenting.

**Examples:**

Good:

- `userAbc123`
- `550e8400-e29b-41d4-a716-446655440000`
- `ord_xyz789`

Bad:

- `12345`
- `9007199254740993`

---

### PO-025: Include request ID in responses

**Severity:** warn

All responses SHOULD include a unique request ID header (X-Request-ID) for tracing and debugging.

**Rationale:** Request IDs enable correlation between client and server logs, essential for debugging distributed systems.

**Examples:**

Good:

- `X-Request-ID: 550e8400-e29b-41d4-a716-446655440000`

---

### PO-032: Use ETag for caching and concurrency

**Severity:** warn

GET responses SHOULD include ETag header. PUT/PATCH SHOULD support If-Match for optimistic locking.

**Rationale:** ETags enable efficient caching and prevent lost updates in concurrent modifications.

**Examples:**

Good:

- `ETag: "abc123"`
- `If-Match: "abc123"`

---

### PO-034: Use consistent envelope for collections

**Severity:** error

Collection responses MUST use a consistent envelope: {"items": [...], "nextCursor": "...", "hasMore": true}

**Rationale:** A consistent envelope enables generic client handling and clearly separates data from metadata.

**Examples:**

Good:

- `{"items": [{"id": "userAbc", "name": "Alice"}], "nextCursor": "eyJpZCI6MTIzfQ", "hasMore": true}`

Bad:

- `Raw array without envelope`
- `{"users": [...], "meta": {...}}`

---

### PO-039: Return created/updated resource in response

**Severity:** warn

POST and PATCH SHOULD return the created/updated resource in the response body.

**Rationale:** Returning the resource saves a round-trip GET and shows server-computed fields (id, timestamps).

---

## Error Handling

Consistent error responses using RFC 7807

### PO-006: Use RFC 7807 Problem Details for errors

**Severity:** error

Error responses MUST use the RFC 7807 Problem Details format with Content-Type: application/problem+json.

**Rationale:** RFC 7807 is an IETF standard for error responses. It provides a consistent structure that clients can parse programmatically.

**Examples:**

Good:

- `{"type": "https://api.example.com/problems/insufficient-funds", "title": "Insufficient Funds", "status": 422, "detail": "Account balance is $10.00, but transaction requires $25.00", "instance": "/accounts/12345/transactions/67890"}`

Bad:

- `{"error": "Something went wrong"}`
- `{"code": 500, "message": "Internal error"}`

---

### PO-019: Never expose internal errors to clients

**Severity:** error

5xx responses MUST NOT expose internal details (stack traces, database errors, internal paths). Use generic messages.

**Rationale:** Internal details are security risks and not useful to clients. Log details server-side, return generic messages to clients.

---

## Pagination

Cursor-based pagination patterns

### PO-007: Use cursor-based pagination

**Severity:** warn

Collection endpoints SHOULD use cursor-based pagination with `cursor` and `limit` parameters, returning `nextCursor` in the response.

**Rationale:** Cursor-based pagination is more scalable than offset-based. It handles real-time data changes gracefully and performs well on large datasets.

**Examples:**

Good:

- `GET /v1/users?limit=20&cursor=eyJpZCI6MTIzfQ`
- `{"items": [...], "nextCursor": "eyJpZCI6MTQzfQ", "hasMore": true}`

Bad:

- `GET /v1/users?page=5&per_page=20`
- `{"items": [...], "page": 5, "total_pages": 100}`

---

### PO-033: Limit response size

**Severity:** warn

Collection endpoints MUST limit response size. Default to 20-50 items, max 100.

**Rationale:** Unbounded responses cause performance issues and timeout errors. Always paginate.

**Examples:**

Good:

- `?limit=20 (default)`
- `?limit=100 (max)`

Bad:

- `Returning 10000 items in one response`

---

## Filtering & Sorting

Query parameter conventions

### PO-026: Use standard filtering syntax

**Severity:** warn

Use query parameters for filtering: ?status=active&created_after=2024-01-01. Avoid complex query languages.

**Rationale:** Simple query parameters are easy to understand and implement. Complex filters can use POST with a body.

**Examples:**

Good:

- `GET /v1/orders?status=pending&created_after=2024-01-01`
- `GET /v1/users?role=admin&sort=createdAt`

Bad:

- `GET /v1/orders?filter=status eq 'pending'`
- `GET /v1/users?q={"role":"admin"}`

---

### PO-027: Use sort parameter with field and direction

**Severity:** info

Use ?sort=field or ?sort=-field for sorting. Prefix with - for descending order.

**Rationale:** This pattern is common and concise. Multiple fields can be comma-separated: ?sort=-createdAt,name

**Examples:**

Good:

- `GET /v1/users?sort=createdAt`
- `GET /v1/users?sort=-createdAt`
- `GET /v1/users?sort=-createdAt,name`

Bad:

- `GET /v1/users?sort=createdAt&order=desc`
- `GET /v1/users?orderBy=createdAt&sortDir=DESC`

---

### PO-035: Support partial responses with fields parameter

**Severity:** info

Endpoints MAY support ?fields=id,name,email to return only specified fields.

**Rationale:** Partial responses reduce bandwidth and improve performance for mobile clients.

**Examples:**

Good:

- `GET /v1/users/123?fields=id,name,email`

---

## Versioning

API versioning and deprecation

### PO-005: Version APIs with URI prefix

**Severity:** error

APIs MUST be versioned using a URI prefix: /v1/, /v2/, etc. Major version only.

**Rationale:** URI versioning is simple, explicit, and cacheable. Including version in the URL makes it obvious which version a client is using. Only major versions need to be in the URL since minor/patch changes are backward compatible.

**Examples:**

Good:

- `/v1/users`
- `/v2/orders`

Bad:

- `/users?version=1`
- `/users with Accept header versioning`

---

### PO-030: Deprecate before removing

**Severity:** error

Endpoints MUST be deprecated for at least 6 months before removal. Use Sunset header to communicate removal date.

**Rationale:** Graceful deprecation gives clients time to migrate. Sudden removal breaks integrations.

**Examples:**

Good:

- `Deprecation: true`
- `Sunset: Sat, 01 Jan 2025 00:00:00 GMT`

---

### PO-031: Never remove fields, only add

**Severity:** error

Collection responses MUST use a consistent envelope: {"items": [...], "nextCursor": "...", "hasMore": true}

**Rationale:** Removing fields breaks existing clients. Add new fields and deprecate old ones if needed.

---

## Security

Authentication, authorization, and security headers

### PO-022: Use Bearer token authentication

**Severity:** error

Collection endpoints SHOULD use cursor-based pagination with `cursor` and `limit` parameters, returning `nextCursor` in the response.

**Rationale:** Bearer tokens are the standard for API authentication. They work with OAuth 2.0, JWT, and API keys.

**Examples:**

Good:

- `Authorization: Bearer eyJhbGciOiJIUzI1NiIs...`

Bad:

- `X-API-Key: secret123`
- `?api_key=secret123`

---

### PO-023: Document security requirements per endpoint

**Severity:** error

Every endpoint MUST document its security requirements in the OpenAPI spec.

**Rationale:** Explicit security documentation prevents accidental exposure and enables security testing.

---

### PO-024: Support CORS for browser clients

**Severity:** warn

APIs consumed by browsers SHOULD support CORS with appropriate Access-Control-Allow-* headers.

**Rationale:** CORS enables secure cross-origin requests from web applications.

---

### PO-036: Use HTTPS only

**Severity:** error

All API endpoints MUST be served over HTTPS. HTTP MUST redirect to HTTPS or return 400.

**Rationale:** HTTPS protects data in transit. There's no reason to use HTTP for production APIs.

---

## Documentation

OpenAPI specification requirements

### PO-028: Document all endpoints with summary and description

**Severity:** error

Every endpoint MUST have a summary (one line) and description (detailed explanation) in the OpenAPI spec.

**Rationale:** Good documentation reduces support burden and improves developer experience. Summary appears in tooling, description provides detail.

---

### PO-029: Provide examples for all schemas

**Severity:** warn

All request/response schemas SHOULD include realistic examples.

**Rationale:** Examples help developers understand expected data formats. They're used in documentation and testing.

---

### PO-037: Document rate limits

**Severity:** warn

API documentation SHOULD specify rate limits and include RateLimit-* headers in responses.

**Rationale:** Clients need to know limits to implement proper backoff. Headers enable dynamic adjustment.

**Examples:**

Good:

- `RateLimit-Limit: 100`
- `RateLimit-Remaining: 95`
- `RateLimit-Reset: 1640000000`

---

### PO-040: Use descriptive operation IDs

**Severity:** warn

OpenAPI operationId SHOULD follow the pattern: verbResource (e.g., getUser, createOrder, listUsers).

**Rationale:** Consistent operation IDs enable better SDK generation and documentation.

**Examples:**

Good:

- `getUser`
- `createOrder`
- `listUsers`
- `updateUserProfile`
- `deleteOrder`

Bad:

- `user_get`
- `order-create`
- `UsersGET`

---

## Glossary

**API First**
: Design approach where the API specification is created before implementation

**Bearer Token**
: An opaque token passed in the Authorization header for authentication

**Cursor**
: An opaque string encoding pagination position, more scalable than offset

**ETag**
: Entity tag for caching and optimistic concurrency control

**Idempotent**
: An operation that produces the same result regardless of how many times it's executed

**kebab-case**
: Naming convention using lowercase words separated by hyphens: user-accounts

**OpenAPI**
: Specification format for describing REST APIs (formerly Swagger)

**Problem Details**
: RFC 7807 standard format for HTTP API error responses

**Resource**
: A domain object exposed through the API, identified by a URL

**snake_case**
: Naming convention using lowercase words separated by underscores (e.g., user_id). PlexusOne uses camelCase instead.

**URI**
: Uniform Resource Identifier, the address of a resource


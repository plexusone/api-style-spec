# PayPal REST API Guidelines

Comprehensive style rules based on PayPal RESTful API Guidelines. These guidelines define standards for designing REST APIs following PayPal's internal best practices. Key differentiators include snake_case property naming, HATEOAS-driven API design, and structured error handling with error catalogs.

## Metadata

- **Author:** PayPal
- **Source:** [https://github.com/paypal/api-standards](https://github.com/paypal/api-standards)
- **Last Updated:** 2024-12-01
- **License:** CC0-1.0

## Table of Contents

- [Introduction](#introduction)
- [Design Principles](#design-principles)
- [Conformance Levels](#conformance-levels)
- [Design Patterns](#design-patterns)
- **Rules**
  - [General Guidelines](#general-guidelines)
  - [HTTP Fundamentals](#http-fundamentals)
  - [Hypermedia](#hypermedia)
  - [Naming Conventions](#naming-conventions)
  - [URI Design](#uri-design)
  - [JSON Schema](#json-schema)
  - [JSON Types](#json-types)
  - [Error Handling](#error-handling)
  - [API Versioning](#api-versioning)
  - [Deprecation](#deprecation)
  - [Compatibility](#compatibility)
  - [Security](#security)
- [Glossary](#glossary)

## Introduction

The PayPal RESTful API Guidelines provide a comprehensive set of best practices for designing REST APIs. They emphasize:
- HATEOAS (Hypermedia As The Engine Of Application State)
- Consistent naming conventions (snake_case for fields, hyphen-case for URIs)
- Structured error handling with error catalogs
- Strong versioning and deprecation policies
- JSON Schema draft-04 for type definitions

## Design Principles

### Loose Coupling

Services and consumers must be loosely coupled from each other. Changes to implementation should not break service consumers.

**Related Rules:** PP-COMPAT-001, PP-COMPAT-002

### Encapsulation

A domain service can access data directly only from its own data sources. Direct database access is forbidden for consumers.

**Related Rules:** PP-DESIGN-001

### Stability

Service contracts must be stable. Only backward-compatible changes may be made to service contracts.

**Related Rules:** PP-COMPAT-001, PP-DEPREC-001

### Reusable

Services should be developed to be reusable across multiple contexts and consumers.

**Related Rules:** PP-DESIGN-002

### Contract-Based

Functionality and technical attributes exposed through service contracts are explicit and unambiguous.

**Related Rules:** PP-SPEC-001, PP-SPEC-002

### Consistency

Services must follow a common set of guidelines, standards, and conventions.

**Related Rules:** PP-NAMING-001, PP-HTTP-001

### Ease of Use

Services should be easy for developers to use without reference to documentation. Follow the principle of least astonishment.

**Related Rules:** PP-DESIGN-003

### Externalizable

Services should be designed for external use. Consider data sensitivity and API context.

**Related Rules:** PP-SECURITY-001

## Conformance Levels

### Bronze

Essential rules for basic API functionality

**Required Rules:**

- PP-SPEC-001
- PP-HTTP-001
- PP-HTTP-002
- PP-HTTP-003
- PP-NAMING-001
- PP-NAMING-003
- PP-ERROR-001
- PP-COMPAT-001

### Silver

Recommended rules for production APIs

**Required Rules:**

- PP-SPEC-001
- PP-SPEC-002
- PP-HTTP-001
- PP-HTTP-002
- PP-HTTP-003
- PP-HTTP-004
- PP-HATEOAS-001
- PP-HATEOAS-002
- PP-NAMING-001
- PP-NAMING-002
- PP-NAMING-003
- PP-NAMING-004
- PP-URI-001
- PP-URI-004
- PP-JSON-001
- PP-JSON-003
- PP-ERROR-001
- PP-ERROR-002
- PP-VERSION-001
- PP-COMPAT-001
- PP-COMPAT-002

### Gold

Full compliance with PayPal REST guidelines

**Required Rules:**

- PP-SPEC-001
- PP-SPEC-002
- PP-HTTP-001
- PP-HTTP-002
- PP-HTTP-003
- PP-HTTP-004
- PP-HTTP-005
- PP-HATEOAS-001
- PP-HATEOAS-002
- PP-HATEOAS-003
- PP-HATEOAS-004
- PP-HATEOAS-005
- PP-NAMING-001
- PP-NAMING-002
- PP-NAMING-003
- PP-NAMING-004
- PP-NAMING-005
- PP-NAMING-006
- PP-URI-001
- PP-URI-002
- PP-URI-003
- PP-URI-004
- PP-URI-005
- PP-URI-006
- PP-QUERY-001
- PP-JSON-001
- PP-JSON-002
- PP-JSON-003
- PP-JSON-004
- PP-JSON-005
- PP-JSON-006
- PP-JSON-007
- PP-JSON-008
- PP-JSON-009
- PP-COMMON-001
- PP-COMMON-002
- PP-COMMON-003
- PP-COMMON-004
- PP-COMMON-005
- PP-COMMON-006
- PP-ERROR-001
- PP-ERROR-002
- PP-ERROR-003
- PP-ERROR-004
- PP-VERSION-001
- PP-VERSION-002
- PP-VERSION-003
- PP-COMPAT-001
- PP-COMPAT-002
- PP-COMPAT-003
- PP-DEPREC-001
- PP-DEPREC-002
- PP-DEPREC-003
- PP-EOL-001
- PP-SECURITY-001

## Design Patterns

### HATEOAS Link Navigation

Use HATEOAS links for state transitions

**Problem:** Clients hardcoding URIs creates tight coupling and makes API evolution difficult.

**Solution:** Return links in responses that guide clients through available state transitions. Clients should follow links rather than construct URIs.

**When to Use:** For all resources that support multiple operations or state transitions

**HATEOAS Response (Correct)**

```json
{
  "id": "ALT-JFWXHGUV7VI",
  "first_name": "John",
  "last_name": "Doe",
  "links": [
    {"href": "https://api.foo.com/v1/customer/users/ALT-JFWXHGUV7VI", "rel": "self"},
    {"href": "https://api.foo.com/v1/customer/users/ALT-JFWXHGUV7VI", "rel": "edit", "method": "PATCH"},
    {"href": "https://api.foo.com/v1/customer/users/ALT-JFWXHGUV7VI", "rel": "delete", "method": "DELETE"}
  ]
}
```

**Related Rules:** PP-HATEOAS-001, PP-HATEOAS-002, PP-HATEOAS-003

---

### Error Catalog Pattern

Use error catalogs for consistent error handling

**Problem:** Hardcoded error messages in code are difficult to localize and maintain.

**Solution:** Externalize error specifications in an error catalog. The catalog contains error names, messages, issues, and suggested actions.

**When to Use:** For all APIs that need consistent, localizable error handling

**Error Catalog Entry (Correct)**

```json
{
  "namespace": "payments",
  "language": "en-US",
  "errors": [{
    "error_spec": {
      "name": "VALIDATION_ERROR",
      "message": "Invalid request - see details",
      "log_level": "ERROR",
      "http_status_codes": [400],
      "issues": [{"id": "InvalidCreditCardType", "issue": "Value is invalid"}]
    }
  }]
}
```

**Related Rules:** PP-ERROR-002, PP-ERROR-003

---

### Snake Case Field Names

Use snake_case for JSON field names

**Problem:** Inconsistent field naming creates confusion across APIs.

**Solution:** Use lowercase words separated by underscores for all JSON field names.

**When to Use:** Always when defining JSON property names

**Snake Case Fields (Correct)**

```json
{
  "first_name": "John",
  "last_name": "Doe",
  "billing_address": {
    "street_address": "123 Main St"
  }
}
```

**Camel Case (Incorrect) (Incorrect)**

```json
{
  "firstName": "John",
  "lastName": "Doe"
}
```

**Related Rules:** PP-NAMING-003

---

## General Guidelines

API specification and fundamental requirements

### PP-SPEC-001: Use OpenAPI for API contract description

**Severity:** error

Use OpenAPI (formerly Swagger) to define API specifications. The OpenAPI Schema Object is based on JSON Schema draft-04 with extensions for documentation.

**Rationale:** OpenAPI provides a vendor-neutral API description format that enables tooling, documentation, and client generation.

**Examples:**

Good:

- `OpenAPI 2.0 or 3.x specification in YAML/JSON`

Bad:

- `No specification`
- `Proprietary documentation format`

**References:**

- [OpenAPI Specification](https://swagger.io/specification/)

---

### PP-DESIGN-001: Encapsulate domain data access

**Severity:** error

A domain service can access data directly only from its own data sources. Consumers MUST NOT have direct database access and MUST go through the service API.

**Rationale:** Direct database access creates tight coupling and bypasses business logic.

---

### PP-DESIGN-002: Design reusable services

**Severity:** info

Services should be developed to be reusable across multiple contexts and consumers. Avoid tight coupling to specific client needs.

**Rationale:** Reusable services reduce duplication and enable broader adoption.

---

### PP-DESIGN-003: Follow principle of least astonishment

**Severity:** info

Services should be easy for developers to use without extensive documentation. API behavior should be predictable and consistent with common expectations.

**Rationale:** Predictable APIs are easier to use and require less documentation.

---

## HTTP Fundamentals

HTTP methods, headers, and status codes

### PP-HTTP-001: Use standard HTTP methods correctly

**Severity:** error

Use HTTP methods according to their semantics:
- GET: Retrieve resource (safe, cacheable)
- POST: Create resource (unsafe, non-idempotent)
- PUT: Replace resource (idempotent)
- PATCH: Partial update (may be idempotent)
- DELETE: Remove resource (idempotent)

**Rationale:** Correct HTTP method usage ensures predictable API behavior and enables caching, idempotency, and safety guarantees.

**Examples:**

Good:

- `GET /users/123 to retrieve a user`
- `POST /users to create a user`
- `PUT /users/123 to replace a user`

Bad:

- `POST /getUser to retrieve a user`
- `GET /deleteUser/123 to delete`

---

### PP-HTTP-002: Use appropriate HTTP status codes

**Severity:** error

Return appropriate HTTP status codes:
- 200 OK: Successful GET/PUT/PATCH/DELETE
- 201 Created: Successful POST creating a resource
- 202 Accepted: Async processing accepted
- 204 No Content: Successful with no response body
- 400 Bad Request: Invalid request syntax or validation errors
- 401 Unauthorized: Missing or invalid authentication
- 403 Forbidden: Valid auth but insufficient permissions
- 404 Not Found: Resource does not exist
- 422 Unprocessable Entity: Semantic validation errors
- 500 Internal Server Error: Server-side error

**Rationale:** Correct status codes enable clients to handle responses appropriately and enable proper caching behavior.

---

### PP-HTTP-003: Use JSON as the primary data format

**Severity:** error

APIs MUST support application/json as the primary content type for requests and responses. Processing MUST be strictly JSON-based with proper serialization/deserialization.

**Rationale:** JSON is the standard data format for REST APIs with wide tooling and library support.

**Examples:**

Good:

- `Content-Type: application/json`
- `Accept: application/json`

Bad:

- `Content-Type: text/plain for JSON data`
- `XML-only APIs`

---

### PP-HTTP-004: Use standard HTTP headers

**Severity:** warn

Use standard HTTP headers for common functionality:
- Content-Type: Media type of request/response body
- Accept: Acceptable media types for response
- Authorization: Authentication credentials
- Content-Language: Language of the content
- Content-Length: Size of the entity-body

**Rationale:** Standard headers ensure interoperability with HTTP infrastructure like caches and proxies.

---

### PP-HTTP-005: Use custom headers sparingly

**Severity:** info

Custom HTTP headers should be used sparingly. If needed, use a consistent prefix (e.g., Foo-Request-Id). Note that HTTP Location header should NOT be used to provide links - use response body instead.

**Rationale:** Custom headers can cause compatibility issues with proxies and CDNs.

**Examples:**

Good:

- `Foo-Request-Id: abc-123-def`

Bad:

- `Using Location header for HATEOAS links`
- `Unprefixed custom headers`

---

## Hypermedia

HATEOAS links and navigation

### PP-HATEOAS-001: Include HATEOAS links in responses

**Severity:** error

All API responses SHOULD include links to guide clients through available state transitions. Links allow clients to navigate the API without hardcoding URIs. The client should only know a single entry point URI.

**Rationale:** HATEOAS enables loose coupling between clients and servers, allowing API evolution without breaking clients.

**Examples:**

Good:

- `Response includes links array with self, edit, delete relations`

Bad:

- `Response has no links`
- `Client required to construct URIs`

---

### PP-HATEOAS-002: Use Link Description Object schema

**Severity:** error

Links MUST use the Link Description Object (LDO) schema with required properties:
- href: Absolute URI template (RFC 6570)
- rel: Link relation type
- method: HTTP verb (defaults to GET)
- title: Optional helpful description

**Rationale:** Consistent link structure enables tooling and client-side link processing.

**Examples:**

Good:

- `{"href": "https://api.foo.com/v1/users/123", "rel": "self", "method": "GET"}`

Bad:

- `Using relative URLs`
- `Missing rel property`

---

### PP-HATEOAS-003: Use absolute URIs in links

**Severity:** error

Use ONLY absolute URIs as values for the href property. The value from the incoming Host header MUST be used as the host field. Clients bookmark absolute URIs for later use.

**Rationale:** Absolute URIs allow clients to bookmark links and make requests without knowing the API base URL.

**Examples:**

Good:

- `https://api.foo.com/v1/users/ALT-JFWXHGUV7VI`

Bad:

- `/v1/users/ALT-JFWXHGUV7VI`
- `users/ALT-JFWXHGUV7VI`

---

### PP-HATEOAS-004: Use standard link relation types

**Severity:** warn

When semantics match, use IANA standardized link relations:
- self: Link to the resource itself
- create: Link to create a new resource
- edit: Link to partially update (PATCH)
- delete: Link to delete the resource
- replace: Link to fully update (PUT)
- first/last/next/prev: Pagination links
- collection: Link to parent collection
- search: Link to search endpoint

For controller actions, use action name as relation (e.g., activate, cancel, refund).

**Rationale:** Standard relation types from IANA ensure consistent semantics across APIs.

---

### PP-HATEOAS-005: Include links array in schema

**Severity:** warn

The links array MUST be declared within the 'properties' keyword of the schema object. All possible links MUST be declared outside properties using URI templates. This enables code generators to add setter/getter methods.

**Rationale:** Declaring links in schema enables code generators to create link accessor methods.

---

## Naming Conventions

URI, field, and resource naming

### PP-NAMING-001: Use lowercase letters in URIs

**Severity:** error

URIs MUST start with a letter and use only lowercase letters. Literals/expressions in URI paths SHOULD be separated using hyphens (-).

**Rationale:** Lowercase URIs are easier to read and avoid case-sensitivity issues.

**Examples:**

Good:

- `https://api.foo.com/v1/credit-cards`
- `https://api.foo.com/v1/billing-agreements`

Bad:

- `https://api.foo.com/v1/CreditCards`
- `https://api.foo.com/v1/credit_cards`

---

### PP-NAMING-002: Use plural nouns for collections

**Severity:** error

Plural nouns SHOULD be used in URIs to identify collections. Resource names MUST be singular for singletons and plural for collections. Use nouns, not verbs.

**Rationale:** Plural nouns clearly indicate that an endpoint returns a collection of resources.

**Examples:**

Good:

- `/invoices`
- `/users`
- `/credit-cards`

Bad:

- `/invoice`
- `/getUsers`
- `/creditCard`

---

### PP-NAMING-003: Use snake_case for field names

**Severity:** error

Key names MUST be lowercase words separated by underscore (_). Fields representing arrays SHOULD use plural nouns. Avoid boolean prefixes like is_ or has_.

**Rationale:** Consistent field naming improves API usability and enables automated processing.

**Examples:**

Good:

- `first_name`
- `billing_address`
- `order_items`

Bad:

- `firstName`
- `FirstName`
- `is_active`

---

### PP-NAMING-004: Use underscore for query parameters

**Severity:** warn

Query parameter names MUST start with a letter and SHOULD be all lowercase. Only alpha characters, digits, and underscore (_) SHALL be used. Values MUST be percent-encoded.

**Rationale:** Consistent query parameter naming across all APIs improves developer experience.

**Examples:**

Good:

- `?start_date=2024-01-01`
- `?page_size=20`

Bad:

- `?startDate=2024-01-01`
- `?page-size=20`

---

### PP-NAMING-005: Use uppercase for enum values

**Severity:** warn

Enum entries SHOULD be composed of only uppercase alphanumeric characters and underscore (_). Industry standards may require exceptions.

**Rationale:** Uppercase enum values are easily distinguishable from other string values.

**Examples:**

Good:

- `FIELD_10`
- `NOT_EQUAL`
- `PAYMENT_COMPLETED`

Bad:

- `field_10`
- `notEqual`
- `PaymentCompleted`

---

### PP-NAMING-006: Use hyphens in URI paths

**Severity:** error

Literals/expressions in URI paths SHOULD be separated using hyphens (-). This is the only place where hyphens are used as word separators. Use underscores (_) everywhere else.

**Rationale:** Hyphens are the standard word separator for URIs and are easier to read than underscores.

**Examples:**

Good:

- `/billing-agreements`
- `/credit-cards`

Bad:

- `/billing_agreements`
- `/creditCards`

---

## URI Design

Resource paths and query parameters

### PP-URI-001: Include version in URI path

**Severity:** error

API endpoints MUST include the major version number in the URI path. Format: /v{major-version}/namespace/resource. Example: /v1/vault/credit-cards

**Rationale:** Version in URI makes API version explicit and enables hosting multiple versions.

**Examples:**

Good:

- `/v1/payments/orders`
- `/v2/customer/users`

Bad:

- `/payments/orders`
- `/api/payments/orders`

---

### PP-URI-002: Include namespace in URI path

**Severity:** error

URI paths MUST include a namespace identifier after the version. Namespaces are determined by logical boundaries in the business capability model.

**Rationale:** Namespaces provide context and scope for resources and prevent naming conflicts.

**Examples:**

Good:

- `/v1/vault/credit-cards`
- `/v1/payments/orders`

Bad:

- `/v1/credit-cards`
- `/v1/orders`

---

### PP-URI-003: Limit sub-resource depth

**Severity:** warn

No more than two levels of sub-resources SHOULD be supported. If specific sub-resource retrieval is needed and IDs are unique, prefer top-level resources.

**Rationale:** Deep nesting creates complex URLs and indicates potential modeling issues.

**Examples:**

Good:

- `/disputes/ABCD1234/documents`
- `/invoice-items/INV-001`

Bad:

- `/disputes/ABCD1234/documents/102030/attachments/ATT-001`

---

### PP-URI-004: Use UUID or HMAC-based resource identifiers

**Severity:** error

A UUID or Hashed Id (HMAC-based) is preferred as resource identifier. APIs MUST NOT use database sequence numbers as resource identifiers. Resource identifiers MUST be owned by the domain model.

**Rationale:** Predictable sequential IDs expose information and enable enumeration attacks.

**Examples:**

Good:

- `CARD-7LT50814996943336KESEVWA`
- `550e8400-e29b-41d4-a716-446655440000`

Bad:

- `/users/1`
- `/orders/12345`

---

### PP-URI-005: Avoid consecutive resource identifiers

**Severity:** error

There MUST NOT be two resource identifiers one after the other in a URI path.

**Rationale:** Consecutive identifiers create ambiguous URIs and indicate poor resource modeling.

**Examples:**

Good:

- `/payments/12345/items/ABC`

Bad:

- `/payments/payments/12345/102030`

---

### PP-URI-006: Scope sub-resource identifiers

**Severity:** error

For security and data integrity, all sub-resource IDs MUST be scoped within the parent resource only. Even if a sub-resource exists, it MUST NOT be returned unless it belongs to the specified parent.

**Rationale:** Scoping prevents unauthorized access to resources through ID guessing.

**Examples:**

Good:

- `/users/1234/linked-accounts/ABCD returns ABCD only if linked to user 1234`

Bad:

- `/users/1234/linked-accounts/ABCD returns ABCD even if not linked to user 1234`

---

### PP-QUERY-001: Use query parameters for filtering

**Severity:** warn

Query parameters SHOULD be used only for restricting collections or as search/filter criteria. Resource identifiers should be in the URI path, not query parameters.

**Rationale:** Query parameters are the standard mechanism for filtering collections.

**Examples:**

Good:

- `/users?status=ACTIVE&created_after=2024-01-01`

Bad:

- `/users?id=123`
- `/getUsers?userId=123`

---

### PP-QUERY-002: Support repeatable query parameters

**Severity:** info

It is RECOMMENDED to pass multiple values by repeating the query parameter (e.g., ?status=CLOSED&status=INVALID). The parameter MUST be marked as repeatable in specifications. Comma-separated values are NOT RECOMMENDED due to additional complexity.

**Rationale:** Repeatable parameters are HTTP-standard and supported by most client libraries.

**Examples:**

Good:

- `?status=CLOSED&status=INVALID`

Bad:

- `?statuses=CLOSED,INVALID`

---

### PP-QUERY-003: Support sorting with sort parameter

**Severity:** info

Default sort order SHOULD be undefined and non-deterministic. For explicit sorting, use the sort query parameter with syntax: {field_name}|{asc|desc},{field_name}|{asc|desc}

**Rationale:** Consistent sort parameter syntax enables client-side sorting controls.

**Examples:**

Good:

- `/accounts?sort=date_of_birth|asc,zip_code|desc`

Bad:

- `/accounts?orderBy=dateOfBirth&order=asc`

---

### PP-PAGINATION-001: Use pagination for collections

**Severity:** warn

Collection responses should support pagination. Use standard link relations (first, last, next, prev) for navigation. Default sort order should be undefined.

**Rationale:** Unbounded collections cause performance and memory issues.

**Examples:**

Good:

- `Response includes next/prev links`
- `total_items and total_pages provided`

Bad:

- `Returning entire collection`
- `No pagination links`

---

## JSON Schema

JSON Schema usage and best practices

### PP-SPEC-002: Use JSON Schema draft-04

**Severity:** warn

Use JSON Schema draft-04 for defining request/response models. If using links, media, or hyper-schema features, use http://json-schema.org/draft-04/hyper-schema#.

**Rationale:** JSON Schema draft-04 is the most widely supported version with mature tooling.

**Examples:**

Good:

- `"$schema": "http://json-schema.org/draft-04/schema#"`

Bad:

- `Using draft-03 or earlier`
- `No $schema declaration`

---

### PP-JSON-008: Use allOf for schema extension

**Severity:** warn

The allOf keyword MUST only be used for extending objects. This replaces the deprecated extends keyword from draft-03.

**Rationale:** allOf provides clear inheritance semantics supported by most tools.

**Examples:**

Good:

- `{"allOf": [{"$ref": "address.json"}, {"properties": {"type": {"enum": ["residential", "business"]}}}]}`

Bad:

- `Using extends keyword`
- `Duplicating base schema properties`

---

### PP-JSON-009: Avoid anyOf and oneOf

**Severity:** warn

The anyOf and oneOf keywords SHOULD NOT be used. Code generators cannot accurately create models. Statically typed languages require custom deserializers. Use flat structures with optional fields instead.

**Rationale:** anyOf/oneOf create problems with code generation, documentation, and deserialization in typed languages.

**Examples:**

Good:

- `{"payment": {...}, "money_request": {...}} with both optional`

Bad:

- `{"extensions": {"oneOf": [{"$ref": "payment.json"}, {"$ref": "money_request.json"}]}}`

---

### PP-JSON-010: Use readOnly for immutable fields

**Severity:** info

When resources contain immutable fields, use the readOnly property to indicate them. PUT/PATCH operations can still update the resource but must not change readOnly fields.

**Rationale:** readOnly clearly indicates which fields cannot be modified by clients.

**Examples:**

Good:

- `{"id": {"type": "string", "readOnly": true}}`

Bad:

- `No indication of immutable fields`

---

## JSON Types

Primitive types and common types

### PP-JSON-001: Define string min and max length

**Severity:** error

Strings SHOULD always explicitly define minLength and maxLength. Without maximum length, database columns cannot be reliably defined. Without minimum, clients may send empty strings inappropriately.

**Rationale:** Length constraints enable database column sizing and backwards-compatible evolution.

**Examples:**

Good:

- `{"type": "string", "minLength": 1, "maxLength": 255}`

Bad:

- `{"type": "string"}`

---

### PP-JSON-002: Use string for enum-like values

**Severity:** warn

The enum keyword SHOULD only be used when values are fixed forever. For extensible values, use string type with documented values and pattern constraint. Set maxLength to 255 and minLength to 1 unless there's a technical reason not to.

**Rationale:** JSON Schema enum cannot be safely extended without breaking backwards compatibility.

**Examples:**

Good:

- `{"type": "string", "minLength": 1, "maxLength": 255, "pattern": "^[A-Z_]+$", "description": "Possible values: OPTION_ONE, OPTION_TWO"}`

Bad:

- `{"enum": ["OPTION_ONE", "OPTION_TWO"]}`

---

### PP-JSON-003: Use string for decimal numbers

**Severity:** error

Never use JSON Schema number type. Use string to represent decimal values. For integers, only use integer type for values that fit in 32-bit signed integer (-2^31 to 2^31-1). Larger values must use string.

**Rationale:** JSON number precision varies by language. JavaScript only supports 64-bit floats, losing precision for large integers.

**Examples:**

Good:

- `{"type": "string", "pattern": "^(-?[0-9]+|-?([0-9]+)?[.][0-9]+)$", "maxLength": 32}`

Bad:

- `{"type": "number"}`

---

### PP-JSON-004: Define integer bounds

**Severity:** error

When using integer type, always provide explicit minimum and maximum. Default to 32-bit signed integer bounds: minimum -2147483648, maximum 2147483647 (or 0 for non-negative).

**Rationale:** Unbounded integers may overflow in some programming languages.

**Examples:**

Good:

- `{"type": "integer", "minimum": 0, "maximum": 2147483647}`

Bad:

- `{"type": "integer"}`

---

### PP-JSON-005: Define array bounds

**Severity:** warn

maxItems SHOULD always be defined. SHOULD NOT exceed 32767 (16-bit signed integer). minItems SHOULD also be defined, typically 0 or 1. Do not use maxItems to communicate page size.

**Rationale:** Unbounded arrays may cause memory issues and should be limited for pagination.

**Examples:**

Good:

- `{"type": "array", "minItems": 0, "maxItems": 100}`

Bad:

- `{"type": "array"}`

---

### PP-JSON-006: Never use null values

**Severity:** error

APIs MUST NOT produce or consume null values. Omit properties rather than setting them to null. Many languages cannot distinguish between undefined and null, leading to serialization issues.

**Rationale:** Null handling varies across languages. Undefined (absent) is clearer than null.

**Examples:**

Good:

- `Omit optional fields: {}`

Bad:

- `{"my_property": null}`

---

### PP-JSON-007: Allow additional properties

**Severity:** error

Schema authors MUST NOT explicitly set additionalProperties to false. This breaks clients that validate responses against schemas. Let the API implementation enforce conformance at runtime instead.

**Rationale:** Setting additionalProperties: false breaks backwards compatibility when new fields are added.

**Examples:**

Good:

- `Omit additionalProperties or set to true`

Bad:

- `{"type": "object", "additionalProperties": false}`

---

### PP-COMMON-001: Use ISO country codes

**Severity:** error

All APIs and services MUST use the ISO 3166-1 alpha-2 two-letter country code standard.

**Rationale:** ISO 3166-1 alpha-2 is the universal standard for country identification.

**Examples:**

Good:

- `US`
- `GB`
- `DE`

Bad:

- `USA`
- `United States`
- `840`

---

### PP-COMMON-002: Use ISO currency codes

**Severity:** error

Currency type MUST use the three-letter currency code as defined in ISO 4217.

**Rationale:** ISO 4217 is the universal standard for currency identification.

**Examples:**

Good:

- `USD`
- `EUR`
- `GBP`

Bad:

- `$`
- `Dollars`
- `840`

---

### PP-COMMON-003: Use BCP-47 language tags

**Severity:** error

Language type MUST use BCP-47 language tags.

**Rationale:** BCP-47 is the standard for language identification used across web technologies.

**Examples:**

Good:

- `en`
- `en-US`
- `zh-Hans`

Bad:

- `English`
- `ENG`
- `1033`

---

### PP-COMMON-004: Use RFC 3339 date-time format

**Severity:** error

Date and time strings MUST conform to RFC 3339 date-time format. APIs MUST only emit UTC time in responses. Accept timezone offsets in requests but convert to UTC.

**Rationale:** RFC 3339 is the standard date-time format for REST APIs.

**Examples:**

Good:

- `2024-01-15T14:30:00.000Z`
- `2024-01-15T14:30:00.000+05:00`

Bad:

- `01/15/2024`
- `January 15, 2024`

---

### PP-COMMON-005: Use IANA timezone identifiers

**Severity:** warn

Timezone strings MUST use IANA timezone database identifiers (aka Olson/tzdata). Do not derive timezone from UTC offset as offsets can map to multiple timezones.

**Rationale:** IANA timezone database is the authoritative source for timezone information.

**Examples:**

Good:

- `America/Los_Angeles`
- `Europe/Berlin`

Bad:

- `PST`
- `UTC-8`
- `-08:00`

---

### PP-COMMON-006: Use money common type

**Severity:** error

Use the money common type for monetary amounts. Both currency_code and value MUST exist. Amounts MUST NOT be negative. Some currencies like JPY have no sub-currency.

**Rationale:** Consistent money representation prevents currency conversion errors.

**Examples:**

Good:

- `{"currency_code": "USD", "value": "10.00"}`

Bad:

- `{"amount": 10.00}`
- `{"currency": "$", "value": -5}`

---

### PP-COMMON-007: Represent percentages as decimal values

**Severity:** warn

Percentages and interest rates MUST be represented as percentages, not decimals. 19.99% should be serialized as 19.99, not 0.1999. Use the percentage common type.

**Rationale:** Consistent percentage representation prevents calculation errors.

**Examples:**

Good:

- `{"interest_rate": "19.99"}`

Bad:

- `{"interest_rate": 0.1999}`

---

## Error Handling

Error schema and error catalog

### PP-ERROR-001: Use error.json schema for errors

**Severity:** error

APIs MUST return JSON error representations conforming to error.json schema with required fields:
- name: Unique human-readable error name
- details: Array of field-level errors (required for 4xx)
- debug_id: Unique server-generated ID for correlation
- message: Human-readable description of the problem
- links: HATEOAS links for error resolution

**Rationale:** Consistent error format enables machine-readable error handling.

**Examples:**

Good:

- `{"name": "VALIDATION_ERROR", "details": [{"field": "/credit_card/expire_month", "issue": "Required field is missing", "location": "body"}], "debug_id": "123456789", "message": "Invalid data provided"}`

Bad:

- `{"error": "Something went wrong"}`
- `{"code": 400, "message": "Bad request"}`

---

### PP-ERROR-002: Use JSON Pointer for field errors

**Severity:** warn

The field property in error details SHOULD use JSON Pointer (RFC 6901) to identify the field in error. Include location (query, path, or body) to indicate where the field is.

**Rationale:** JSON Pointer provides unambiguous field identification in nested structures.

**Examples:**

Good:

- `{"field": "/credit_card/expire_month", "location": "body"}`

Bad:

- `{"field": "credit_card.expire_month"}`
- `{"field": "expireMonth"}`

---

### PP-ERROR-003: Use error catalogs

**Severity:** warn

Create error catalogs to externalize error specifications. The catalog contains error_spec entries with name, message, log_level, http_status_codes, issues, and suggested actions. This enables localization and keeps documentation in sync.

**Rationale:** Error catalogs enable localization, consistency, and separation of error definitions from code.

---

### PP-ERROR-004: Distinguish validation errors from semantic errors

**Severity:** error

Use appropriate status codes for validation:
- 400 Bad Request: Not well-formed JSON or validation errors client can fix
- 422 Unprocessable Entity: Well-formed but semantically invalid (e.g., insufficient balance)

**Rationale:** Different error types require different client handling strategies.

**Examples:**

Good:

- `400 for missing required field`
- `422 for insufficient balance`

Bad:

- `400 for all validation errors`
- `500 for validation errors`

---

## API Versioning

Lifecycle and versioning policy

### PP-VERSION-001: Follow semantic versioning

**Severity:** error

API specifications MUST follow versioning scheme: v{major}.{minor}. Major versions start at 1 for first LIVE release. Minor versions maintain backward compatibility within major version.

**Rationale:** Semantic versioning communicates the impact of changes to consumers.

**Examples:**

Good:

- `v1.0`
- `v2.3`

Bad:

- `v0.1`
- `version1`
- `1.0.0`

---

### PP-VERSION-002: Only expose major version in URI

**Severity:** error

API endpoints MUST only reflect the major version. Minor version changes should be transparent to clients as they maintain backward compatibility.

**Rationale:** Minor versions should be backward compatible and not require URI changes.

**Examples:**

Good:

- `/v1/payments/orders`
- `/v2/users`

Bad:

- `/v1.2/payments/orders`
- `/v1-beta/users`

---

### PP-VERSION-003: Track API lifecycle state

**Severity:** warn

APIs should track lifecycle states:
- PLANNED: Scheduled for development
- BETA: Available to selected subscribers for testing
- LIVE: Available to new subscribers, fully supported
- DEPRECATED: Available to existing subscribers, not accepting new
- RETIRED: Unpublished, removed from production

**Rationale:** Clear lifecycle states enable proper API governance and consumer planning.

---

### PP-EOL-001: Provide migration path before deprecation

**Severity:** error

A major API version MUST NOT be DEPRECATED until a replacement is LIVE with clear migration path. This SHOULD include documentation, migration tools, and sample code.

**Rationale:** Consumers need a clear migration path before deprecated versions are removed.

---

### PP-EOL-002: Allow adequate deprecation period

**Severity:** warn

Deprecated API versions MUST remain in DEPRECATED state for a minimum period to give clients adequate migration notice. External clients may require longer periods.

**Rationale:** Adequate notice enables consumers to plan and execute migrations.

---

## Deprecation

Deprecation annotations and runtime

### PP-DEPREC-001: Use x-deprecated annotation

**Severity:** warn

Use the x-deprecated annotation to mark deprecated API elements in OpenAPI specifications. Include since_version and see (replacement) properties. Supported for resources, methods, parameters, and schema properties.

**Rationale:** Deprecation annotations enable tooling to highlight deprecated elements.

**Examples:**

Good:

- `{"x-deprecated": {"since_version": "1.4", "see": "new-endpoint"}}`

Bad:

- `Deprecation only in description text`
- `No deprecation annotation`

---

### PP-DEPREC-002: Return deprecation header at runtime

**Severity:** warn

The API server MUST respond with a custom deprecation header when deprecated elements are used in request or response. Applications should check for header existence and log warnings.

**Rationale:** Runtime deprecation headers enable automated client notifications.

**Examples:**

Good:

- `Foo-Deprecated: {}`

Bad:

- `No deprecation header`
- `Deprecation only in response body`

---

### PP-DEPREC-003: Support deprecated elements for major version lifetime

**Severity:** error

Deprecated API elements MUST remain supported for the life of the major version or until customers are no longer using them.

**Rationale:** Premature removal of deprecated features breaks existing consumers.

---

## Compatibility

Backward compatibility rules

### PP-COMPAT-001: Make only additive and optional changes

**Severity:** error

Backward-compatible changes MUST be:
1. All changes additive
2. All changes optional
3. Semantics unchanged
4. Parameters unordered
5. Additional functionality as optional extensions or new child resources

**Rationale:** Backward compatibility maintains consumer trust and avoids breaking deployments.

---

### PP-COMPAT-002: Recognize all previous valid values

**Severity:** error

A service MUST recognize all previously valid values for a parameter and SHOULD NOT throw errors when used. For enum types, existing enumerated values and their meanings MUST NOT change.

**Rationale:** Removing valid enum values breaks existing clients.

---

### PP-COMPAT-003: Preserve response structure

**Severity:** error

Existing properties MUST continue to be returned with same name and JSON type. Array content types MUST NOT change. New properties MAY be added but MUST NOT alter meaning of existing properties and MUST NOT be mandatory.

**Rationale:** Clients depend on response structure for deserialization.

---

## Security

Security and data protection

### PP-SECURITY-001: Design for external use

**Severity:** warn

Services should be designed as if they will be exposed externally. Consider data sensitivity, access controls, and appropriate context for all API operations.

**Rationale:** APIs designed for external use have better security posture.

---

## Glossary

**HATEOAS**
: Hypermedia As The Engine Of Application State - REST architectural constraint where clients navigate APIs via hyperlinks.

**LDO**
: Link Description Object - JSON object describing a hyperlink with href, rel, method, and title properties.

**snake_case**
: Naming convention using lowercase letters with underscores between words (e.g., first_name).

**hyphen-case**
: Naming convention for URIs using lowercase letters with hyphens between words (e.g., credit-cards).

**Error Catalog**
: Centralized collection of error specifications for a namespace, supporting localization and consistency.

**Namespace**
: A collection of related capabilities for a domain, forming the top-level segment of an API path.

**Capability**
: A business-oriented abstraction representing a set of related resources.

**Resource Identifier**
: A unique identifier for a resource, typically a UUID or HMAC-based hash.


# Zalando REST API Guidelines

Comprehensive style rules based on Zalando RESTful API Guidelines. These guidelines define standards for designing APIs at Zalando and are used by the Zally API linter. Key differentiators include snake_case property naming, cursor-based pagination, and Problem+JSON error handling.

## Metadata

- **Author:** Zalando SE
- **Source:** [https://opensource.zalando.com/restful-api-guidelines/](https://opensource.zalando.com/restful-api-guidelines/)
- **Last Updated:** 2024-12-01
- **License:** CC-BY-4.0

## Table of Contents

- [Introduction](#introduction)
- [Design Principles](#design-principles)
- [Conformance Levels](#conformance-levels)
- [Design Patterns](#design-patterns)
- **Rules**
  - [General Guidelines](#general-guidelines)
  - [Compatibility](#compatibility)
  - [JSON Guidelines](#json-guidelines)
  - [Data Formats](#data-formats)
  - [URL Design](#url-design)
  - [HTTP Methods](#http-methods)
  - [HTTP Status Codes](#http-status-codes)
  - [Pagination](#pagination)
  - [HTTP Headers](#http-headers)
  - [Hypermedia](#hypermedia)
  - [Deprecation](#deprecation)
  - [Security](#security)
  - [Performance](#performance)
  - [Events](#events)
- [Glossary](#glossary)

## Introduction

The Zalando RESTful API Guidelines are a set of best practices for designing REST APIs. They have been developed over several years of API development at Zalando and incorporate lessons learned from building large-scale e-commerce platforms.

These guidelines emphasize:
- API First development approach
- Consistent naming conventions (snake_case)
- Cursor-based pagination for performance
- RFC 7807 Problem+JSON for errors
- Strong compatibility rules for evolution

## Design Principles

### API First

Design APIs first, before implementing them. Use OpenAPI as the specification language and ensure API specifications are the source of truth.

**Related Rules:** Z-100, Z-101

### Consistency

APIs should look like they came from the same hand. Use consistent naming, error handling, and patterns across all APIs.

**Related Rules:** Z-118, Z-129, Z-151

### Backward Compatibility

Change APIs without breaking consumers. Follow rules for compatible extensions and proper deprecation.

**Related Rules:** Z-106, Z-107, Z-108

### Robustness (Postel's Law)

Be conservative in what you do, be liberal in what you accept from others. Design APIs that are tolerant of client variations.

**Related Rules:** Z-109, Z-108

## Conformance Levels

### Bronze

Basic Zalando compliance

**Required Rules:**

- Z-100
- Z-101
- Z-118
- Z-129
- Z-134
- Z-167
- Z-176
- Z-243

### Silver

Standard Zalando compliance

**Required Rules:**

- Z-100
- Z-101
- Z-103
- Z-106
- Z-107
- Z-118
- Z-120
- Z-125
- Z-129
- Z-134
- Z-136
- Z-148
- Z-150
- Z-151
- Z-159
- Z-160
- Z-167
- Z-169
- Z-171
- Z-176
- Z-243
- Z-104
- Z-227

### Gold

Full Zalando compliance

**Required Rules:**

- Z-100
- Z-101
- Z-102
- Z-103
- Z-106
- Z-107
- Z-108
- Z-109
- Z-112
- Z-118
- Z-120
- Z-122
- Z-125
- Z-110
- Z-129
- Z-134
- Z-135
- Z-136
- Z-137
- Z-148
- Z-149
- Z-150
- Z-151
- Z-154
- Z-159
- Z-160
- Z-161
- Z-248
- Z-167
- Z-169
- Z-171
- Z-172
- Z-176
- Z-182
- Z-183
- Z-187
- Z-189
- Z-219
- Z-227
- Z-229
- Z-238
- Z-243
- Z-252
- Z-253
- Z-104
- Z-105

## Design Patterns

### Cursor-Based Pagination

Use cursor-based pagination for large collections

**Problem:** Offset-based pagination performs poorly on large datasets and sharded databases.

**Solution:** Use opaque cursors that encode page position. Return cursors in pagination links.

**When to Use:** For collections that may contain many items or use NoSQL databases

**Paginated Response (Correct)**

```json
{
  "items": [
    {"id": "abc123", "name": "Item 1"},
    {"id": "def456", "name": "Item 2"}
  ],
  "self": "https://api.example.com/items?cursor=abc",
  "next": "https://api.example.com/items?cursor=xyz"
}
```

**Related Rules:** Z-159, Z-160, Z-161, Z-248

---

### Problem+JSON Error Responses

Use RFC 7807 Problem+JSON for error responses

**Problem:** Inconsistent error formats make error handling difficult for clients.

**Solution:** Use standard Problem+JSON format with type, title, status, detail, and instance fields.

**When to Use:** For all 4xx and 5xx error responses

**Problem Response (Correct)**

```json
{
  "type": "https://zalando.github.io/problem/constraint-violation",
  "title": "Constraint Violation",
  "status": 400,
  "detail": "The 'name' field is required.",
  "instance": "/orders/12345"
}
```

**Related Rules:** Z-176, Z-177

---

### Snake Case Naming

Use snake_case for all JSON property names

**Problem:** Inconsistent naming conventions create confusion across APIs.

**Solution:** Use snake_case for all JSON properties: lowercase letters, digits, and underscores.

**When to Use:** Always when defining JSON property names

**Snake Case Properties (Correct)**

```json
{
  "customer_number": "12345",
  "sales_order_number": "SO-001",
  "billing_address": {
    "street_name": "Example St"
  }
}
```

**Camel Case (Incorrect) (Incorrect)**

```json
{
  "customerNumber": "12345",
  "salesOrderNumber": "SO-001"
}
```

**Related Rules:** Z-118

---

## General Guidelines

API First principle and fundamental requirements

### Z-100: Follow API first principle

**Severity:** error

You must follow the API First principle:

1. Define APIs first, before coding implementation, using OpenAPI as specification language
2. Design APIs consistently with these guidelines
3. Use Zally for automated rule checks
4. Call for early review feedback from peers and client developers

**Rationale:** API First ensures APIs are designed intentionally before implementation, leading to better API quality and consistency.

**References:**

- [Zalando API First](https://opensource.zalando.com/restful-api-guidelines/#100)

---

### Z-101: Provide API specifications using OpenAPI

**Severity:** error

Use the OpenAPI standard to define API specifications. Provide the specification using a single self-contained YAML file. We encourage OpenAPI 3.1 for new APIs.

**Rationale:** OpenAPI is the standard for REST API specifications and enables tooling, documentation, and client generation.

**Examples:**

Good:

- `openapi: 3.1.0 in YAML format`

Bad:

- `No specification`
- `Multiple scattered files without proper references`

---

### Z-102: Provide API user manual

**Severity:** warn

In addition to the API specification, provide a user manual covering:
- API scope, purpose, and use cases
- Concrete examples of API usage
- Edge cases, error details, and repair hints
- Architecture context and dependencies

**Rationale:** An API user manual helps developers understand how to use the API effectively beyond what the specification provides.

---

### Z-103: Write APIs using U.S. English

**Severity:** error

Using consistent English ensures APIs are understandable by a global audience.

---

### Z-234: Only use durable and immutable remote references

**Severity:** error

API specification files must be self-contained. Remote references are only allowed to:
- Zalando's API repository (immutable revisions)
- Guideline-defined fragments at opensource.zalando.com

**Rationale:** Remote references to mutable content can cause API semantics to change unexpectedly.

---

### Z-219: Define API audience

**Severity:** error

Use x-api-audience to specify API visibility:
- component-internal: Within component
- business-unit-internal: Within business unit
- company-internal: Within company
- external-partner: Partners
- external-public: Public

**Rationale:** API audience determines review requirements and documentation needs.

---

### Z-116: Use semantic versioning

**Severity:** error

Use Semantic Versioning 2.0 (MAJOR.MINOR.PATCH) for API specification version in #/info/version. Increment MAJOR for incompatible changes, MINOR for backwards-compatible additions, PATCH for bug fixes.

**Rationale:** Semantic versioning provides clear expectations about API changes and their impact.

**Examples:**

Good:

- `1.0.0`
- `2.3.1`
- `0.1.0`

Bad:

- `1.0`
- `v1`
- `1.0.0-beta`

**References:**

- [Zalando Rule 116](https://opensource.zalando.com/restful-api-guidelines/#116)

---

### Z-215: Provide API identifiers

**Severity:** error

Each API specification must have a globally unique and immutable API identifier via x-api-id extension. Use UUID format to avoid conflicts.

**Rationale:** Globally unique API identifiers support lifecycle management and version tracking.

**Examples:**

Good:

- `x-api-id: d0184f38-b98d-11e7-9c56-68f728c1ba70`

Bad:

- `No x-api-id defined`
- `x-api-id changed between versions`

**References:**

- [Zalando Rule 215](https://opensource.zalando.com/restful-api-guidelines/#215)

---

### Z-218: Contain API meta information

**Severity:** error

API specifications must contain: info/title, info/version, info/description, info/contact, x-api-id, and x-audience.

**Rationale:** Complete API metadata ensures discoverability and proper documentation.

**References:**

- [Zalando Rule 218](https://opensource.zalando.com/restful-api-guidelines/#218)

---

### Z-223: Use functional naming schema

**Severity:** warn

Use functional naming for hostnames and event names: <functional-domain>-<functional-component>. Required for external APIs, recommended for internal APIs.

**Rationale:** Functional naming aligns global resources and keeps APIs stable through organizational changes.

**References:**

- [Zalando Rule 223](https://opensource.zalando.com/restful-api-guidelines/#223)

---

### Z-224: Follow naming convention for hostnames

**Severity:** warn

Hostnames should follow functional naming: <functional-name>.zalandoapis.com for external APIs.

**Rationale:** Consistent hostname conventions improve discoverability and maintainability.

**References:**

- [Zalando Rule 224](https://opensource.zalando.com/restful-api-guidelines/#224)

---

### Z-192: Publish OpenAPI specification for APIs

**Severity:** error

All service applications must publish API specifications of their external APIs. While optional for component-internal APIs, it is recommended to profit from API documentation and review tooling.

**Rationale:** Published API specifications enable discoverability and client generation.

**References:**

- [Zalando Rule 192](https://opensource.zalando.com/restful-api-guidelines/#192)

---

### Z-193: Monitor API usage

**Severity:** warn

Owners of APIs used in production should monitor API usage to identify clients. This helps find review partners for API changes and track deprecation migration progress.

**Rationale:** API usage monitoring identifies clients and supports change management.

**References:**

- [Zalando Rule 193](https://opensource.zalando.com/restful-api-guidelines/#193)

---

## Compatibility

Backward compatibility and API evolution

### Z-106: Not break backward compatibility

**Severity:** error

Change APIs, but keep all consumers running. There are two techniques:
1. Follow rules for compatible extensions
2. Introduce new API versions and support older versions with deprecation

We strongly encourage compatible extensions and discourage versioning.

**Rationale:** Breaking changes disrupt API consumers and erode trust. APIs are contracts that cannot be broken unilaterally.

---

### Z-107: Prefer compatible extensions

**Severity:** warn

Follow these rules for backward-compatible changes:

For input schemas:
- Add optional fields only, never mandatory
- Make mandatory fields optional, not vice-versa
- Don't remove fields
- Don't make validation more restrictive

For output schemas:
- Add fields (mandatory or optional)
- Make optional fields mandatory, not vice-versa
- Don't remove fields
- Enum ranges cannot be extended

**Rationale:** Compatible extensions allow API evolution without versioning complexity.

---

### Z-108: Prepare clients to accept compatible API extensions

**Severity:** error

Clients must be prepared to accept:
- New fields in response payloads (ignore unknown fields)
- New enum values (handle gracefully)
- New optional request fields

**Rationale:** Clients must be tolerant of API extensions to enable server evolution.

---

### Z-109: Design APIs conservatively

**Severity:** warn

Be conservative and accurate in what you accept from clients:
- Unknown input fields should return HTTP 400 (not silently ignored)
- Be accurate in defining input constraints
- Prefer being more specific and restrictive

**Rationale:** Being strict about input prevents errors and makes APIs more predictable.

---

### Z-112: Use open-ended list of values (x-extensible-enum)

**Severity:** warn

For enumerations used in output that may be extended, use x-extensible-enum instead of enum. This signals to clients that they should handle unknown values.

**Rationale:** Closed enums prevent backward-compatible additions; x-extensible-enum signals that new values may be added.

**Examples:**

Good:

- `x-extensible-enum: [PENDING, APPROVED, REJECTED]`

Bad:

- `enum: [PENDING, APPROVED, REJECTED] for output fields that may grow`

---

### Z-111: Treat OpenAPI specification as open for extension by default

**Severity:** error

OpenAPI object definitions are considered open for extension by default. API clients must not assume objects are closed for extension in the absence of an additionalProperties declaration. API formats must not declare additionalProperties to be false.

**Rationale:** OpenAPI specifications should be extensible by default to support API evolution without breaking clients.

**References:**

- [Zalando Rule 111](https://opensource.zalando.com/restful-api-guidelines/#111)

---

### Z-113: Avoid versioning

**Severity:** warn

Avoid generating additional API versions. If incompatible changes are necessary, prefer creating new resources or new service endpoints over versioning.

**Rationale:** Multiple API versions significantly complicate understanding, testing, maintaining, and operating systems.

**References:**

- [Zalando Rule 113](https://opensource.zalando.com/restful-api-guidelines/#113)

---

### Z-114: Use media type versioning

**Severity:** error

When API versioning is unavoidable, use media type versioning. Version information is provided via Content-Type header (e.g., application/x.zalando.cart+json;version=2).

**Rationale:** Media type versioning is less tightly coupled and supports content negotiation.

**Examples:**

Good:

- `Content-Type: application/x.zalando.cart+json;version=2`

Bad:

- `/v1/customers`
- `/api/v2/orders`

**References:**

- [Zalando Rule 114](https://opensource.zalando.com/restful-api-guidelines/#114)

---

### Z-115: Not use URL versioning

**Severity:** error

Do not include version numbers in URLs (e.g., /v1/customers). Use media type versioning with content negotiation instead.

**Rationale:** URL versioning creates tight coupling and complex release management.

**Examples:**

Good:

- `Accept: application/x.zalando.cart+json;version=2`

Bad:

- `/v1/customers`
- `/api/v2/orders`

**References:**

- [Zalando Rule 115](https://opensource.zalando.com/restful-api-guidelines/#115)

---

## JSON Guidelines

JSON payload structure and naming

### Z-167: Use JSON as payload data interchange format

**Severity:** error

Use JSON (RFC 7159) to represent structured data in requests and responses. The JSON payload must:
- Use a JSON object as top-level structure (not an array)
- Use UTF-8 encoding
- Consist of valid Unicode strings
- Contain only unique member names

**Rationale:** JSON is the standard for REST APIs and has wide tooling support.

---

### Z-118: Property names must be snake_case

**Severity:** error

Property names must use ASCII snake_case matching regex `^[a-z_][a-z_0-9]*$`. The first character must be a lowercase letter or underscore.

**Rationale:** Consistent naming makes APIs predictable. snake_case is preferred by many tech companies including GitHub, Twitter, and Stack Exchange.

**Examples:**

Good:

- `customer_number`
- `sales_order_number`
- `billing_address`

Bad:

- `customerNumber`
- `SalesOrderNumber`
- `billingAddress`

---

### Z-120: Pluralize array names

**Severity:** warn

Plural names clearly indicate arrays; singular names indicate single objects.

**Examples:**

Good:

- `items: []`
- `orders: []`
- `addresses: []`

Bad:

- `item: []`
- `order: []`

---

### Z-125: Declare enum values using UPPER_SNAKE_CASE

**Severity:** warn

UPPER_SNAKE_CASE clearly distinguishes enum values from properties.

**Examples:**

Good:

- `PENDING`
- `IN_PROGRESS`
- `COMPLETED`

Bad:

- `pending`
- `inProgress`
- `Completed`

---

### Z-122: Null values should have their fields removed

**Severity:** warn

APIs should not return fields with null values. Instead, omit the field entirely from the response.

**Rationale:** Omitting null fields reduces payload size and avoids ambiguity between null and missing.

**Examples:**

Good:

- `{"name": "John"}`

Bad:

- `{"name": "John", "middle_name": null}`

---

### Z-110: Response payloads must be JSON objects, not arrays

**Severity:** error

Response payloads must use a JSON object as the top-level structure. Collections should be wrapped in an object with an 'items' property.

**Rationale:** Top-level arrays prevent future extension; objects allow adding metadata fields.

**Examples:**

Good:

- `{"items": [{"id": 1}, {"id": 2}]}`

Bad:

- `[{"id": 1}, {"id": 2}]`

---

### Z-172: Use standard media types

**Severity:** warn

Use application/json for JSON payloads and application/problem+json for errors. Avoid custom media types.

**Rationale:** Standard media types are well-supported by clients and tools.

---

### Z-252: Design single schema for reading and writing

**Severity:** warn

Use the same schema for reading and writing resources. Mark properties as readOnly (only in response) or writeOnly (only in request) as needed.

**Rationale:** Single schemas reduce complexity and cognitive load for API consumers.

---

### Z-123: Use same semantics for null and absent properties

**Severity:** error

Both null values and absent properties must have the same meaning. Exception: JSON Merge Patch uses null to indicate property deletion.

**Rationale:** Different semantics for null and absent values creates confusion and implementation complexity.

**References:**

- [Zalando Rule 123](https://opensource.zalando.com/restful-api-guidelines/#123)

---

### Z-124: Not use null for empty arrays

**Severity:** warn

Use empty array [] instead of null to represent empty collections.

**Rationale:** Empty arrays can be unambiguously represented as [], avoiding null ambiguity.

**Examples:**

Good:

- `{"items": []}`

Bad:

- `{"items": null}`

**References:**

- [Zalando Rule 124](https://opensource.zalando.com/restful-api-guidelines/#124)

---

### Z-168: Pass non-JSON media types using data specific standard formats

**Severity:** info

Non-JSON media types may be supported if using business object specific standard formats (e.g., image/png, application/pdf). Generic formats like XML or CSV should only be provided additionally to JSON.

**Rationale:** Standard formats for non-JSON data ensure interoperability and proper client handling.

**References:**

- [Zalando Rule 168](https://opensource.zalando.com/restful-api-guidelines/#168)

---

### Z-173: Use the common money object

**Severity:** error

Use the standard Money structure with amount (decimal) and currency (ISO 4217) fields. Treat Money as a closed data type, not for inheritance.

**Rationale:** Consistent money representation prevents precision issues and enables library support.

**Examples:**

Good:

- `{"amount": 19.99, "currency": "EUR"}`

Bad:

- `{"price": 19.99}`
- `{"amount": 19.99, "currency": "EUR", "discount": 5.00}`

**References:**

- [Zalando Rule 173](https://opensource.zalando.com/restful-api-guidelines/#173)

---

### Z-174: Use common field names and semantics

**Severity:** error

Use standard field names: id (opaque string identifier), xyz_id (reference to another object), e_tag (ETag for embedded sub-resources), created_at, modified_at.

**Rationale:** Common field names create consistency across APIs and improve understanding.

**Examples:**

Good:

- `id`
- `customer_id`
- `created_at`
- `modified_at`

Bad:

- `ID`
- `customerId`
- `creation_date`

**References:**

- [Zalando Rule 174](https://opensource.zalando.com/restful-api-guidelines/#174)

---

### Z-216: Define maps using additionalProperties

**Severity:** warn

Define maps (string key to value mappings) using additionalProperties with a schema defining the value type. Map keys don't need to follow snake_case naming.

**Rationale:** Maps with string keys should use additionalProperties for proper schema definition.

**References:**

- [Zalando Rule 216](https://opensource.zalando.com/restful-api-guidelines/#216)

---

### Z-235: Use naming convention for date/time properties

**Severity:** warn

Date and time property names should contain 'date', 'time', 'timestamp', or end with '_at' suffix (e.g., created_at, modified_at, campaign_start_time).

**Rationale:** Consistent date/time naming makes properties easily identifiable.

**Examples:**

Good:

- `created_at`
- `modified_at`
- `campaign_start_time`
- `arrival_date`

Bad:

- `created`
- `modified`
- `campaign_start`

**References:**

- [Zalando Rule 235](https://opensource.zalando.com/restful-api-guidelines/#235)

---

### Z-249: Use the common address fields

**Severity:** warn

Use standard address fields: street, additional, city, zip, country_code (ISO 3166-1 alpha-2). For addressees: salutation, first_name, last_name, business_name.

**Rationale:** Standard address structures ensure consistency across APIs handling addresses.

**References:**

- [Zalando Rule 249](https://opensource.zalando.com/restful-api-guidelines/#249)

---

### Z-250: Be aware of services not fully supporting JSON/unicode

**Severity:** warn

Services forwarding JSON content to other tools should verify those tools fully support JSON/unicode. Postgres cannot handle \u0000 in jsonb/text types.

**Rationale:** Some downstream services (e.g., Postgres) cannot handle all valid JSON/unicode characters.

**References:**

- [Zalando Rule 250](https://opensource.zalando.com/restful-api-guidelines/#250)

---

### Z-240: Declare enum values using UPPER_SNAKE_CASE string

**Severity:** warn

Enum values should use UPPER_SNAKE_CASE format. Use string type for enums. For extensible enums, use x-extensible-enum or examples with description.

**Rationale:** Consistent enum value formatting improves readability and code generation.

**Examples:**

Good:

- `ACTIVE`
- `PENDING_APPROVAL`
- `OUT_OF_STOCK`

Bad:

- `active`
- `PendingApproval`
- `out-of-stock`

**References:**

- [Zalando Rule 240/125](https://opensource.zalando.com/restful-api-guidelines/#240)

---

## Data Formats

Standard data types and formats

### Z-238: Use standard data formats

**Severity:** error

Use OpenAPI/JSON Schema standard formats for data types:
- Integers: int32, int64, bigint
- Numbers: float, double, decimal
- Dates: date, time, date-time
- Strings: email, uri, uuid, etc.

**Rationale:** Standard formats ensure interoperability and enable proper client code generation.

---

### Z-171: Define a format for number and integer types

**Severity:** error

Always provide format (int32, int64, bigint, float, double, decimal) for number and integer types.

**Rationale:** Without explicit format, clients may guess precision incorrectly.

**Examples:**

Good:

- `type: integer, format: int64`

Bad:

- `type: integer (no format)`

---

### Z-169: Use standard formats for date and time

**Severity:** error

Use string typed formats: date, time, date-time, duration, or period. Times should be UTC with uppercase T separator and Z suffix.

**Rationale:** RFC 3339/ISO 8601 dates are unambiguous and widely supported.

**Examples:**

Good:

- `2024-01-15T10:30:00Z`

Bad:

- `01/15/2024`
- `1705315800`

---

### Z-127: Define format for duration and period

**Severity:** warn

ISO 8601 duration format is standard and unambiguous.

**Examples:**

Good:

- `P1DT3H4S (1 day, 3 hours, 4 seconds)`

Bad:

- `86400 (seconds)`
- `1 day`

---

### Z-144: Use UUID for resource IDs

**Severity:** warn

Use UUID format for resource identifiers where applicable.

**Rationale:** UUIDs are globally unique and don't leak information about resource count.

---

### Z-239: Encode binary data in base64url

**Severity:** error

Binary data must be defined as string with binary format using base64url encoding.

**Rationale:** Base64url encoding is the standard for binary data in JSON/REST APIs.

**References:**

- [Zalando Rule 239](https://opensource.zalando.com/restful-api-guidelines/#239)

---

### Z-244: Use content negotiation

**Severity:** warn

Support content negotiation via Accept, Accept-Language, Accept-Encoding headers when serving different representations of a resource.

**Rationale:** Content negotiation allows clients to choose the best representation.

**References:**

- [Zalando Rule 244](https://opensource.zalando.com/restful-api-guidelines/#244)

---

### Z-255: Use appropriate formats for date and time properties

**Severity:** warn

Use date-time for exact points in time, date for date-only (implies local timezone), time-local/date-time-local for local times supplemented with timezone field.

**Rationale:** Choosing the right date/time format prevents misinterpretation and timezone issues.

**References:**

- [Zalando Rule 255](https://opensource.zalando.com/restful-api-guidelines/#255)

---

### Z-170: Use standard formats for country, language and currency properties

**Severity:** error

Use ISO standards: ISO 3166-1 alpha-2 for country codes (format: iso-3166-alpha-2), ISO 639-1 for language codes (format: iso-639-1), BCP 47 for language tags (format: bcp47), ISO 4217 for currency codes (format: iso-4217).

**Rationale:** Standard codes ensure interoperability and prevent ambiguity.

**Examples:**

Good:

- `DE`
- `en`
- `en-US`
- `EUR`

Bad:

- `Germany`
- `English`
- `dollars`

**References:**

- [Zalando Rule 170](https://opensource.zalando.com/restful-api-guidelines/#170)

---

### Z-126: Use standard formats for date and time properties

**Severity:** error

Use string typed formats: date, time-local, date-time-local, time, date-time, duration, or period for date/time properties. Date-time must be UTC with Z suffix (ISO 8601).

**Rationale:** Standard date/time formats ensure interoperability and parsing consistency.

**Examples:**

Good:

- `2024-01-15`
- `2024-01-15T14:30:00Z`
- `14:30:00`
- `P1DT2H30M`

Bad:

- `01/15/2024`
- `2024-01-15T14:30:00+00:00`

**References:**

- [Zalando Rule 126/169](https://opensource.zalando.com/restful-api-guidelines/#126)

---

### Z-128: Use standard formats for country, language and currency

**Severity:** error

Use ISO standards: ISO 3166-1 alpha-2 for countries, ISO 639-1 for languages, BCP 47 for language tags, ISO 4217 for currencies.

**Rationale:** Standard codes ensure interoperability and prevent ambiguity.

**Examples:**

Good:

- `DE`
- `en`
- `en-US`
- `EUR`

Bad:

- `Germany`
- `English`
- `dollars`

**References:**

- [Zalando Rule 128/170](https://opensource.zalando.com/restful-api-guidelines/#128)

---

## URL Design

Resource paths and query parameters

### Z-129: Use lowercase separate words with hyphens for paths

**Severity:** error

URL paths must use lowercase letters with hyphens as word separators (kebab-case).

**Rationale:** Kebab-case paths are readable and consistent with URL conventions.

**Examples:**

Good:

- `/shopping-carts`
- `/order-items`

Bad:

- `/shoppingCarts`
- `/order_items`
- `/OrderItems`

---

### Z-134: Pluralize resource names

**Severity:** error

Plural names indicate collections; consistent naming improves API predictability.

**Examples:**

Good:

- `/customers`
- `/orders`
- `/products`

Bad:

- `/customer`
- `/order`
- `/product`

---

### Z-135: Not use /api as base path

**Severity:** warn

/api is redundant when the entire service is an API.

**Examples:**

Good:

- `/orders`
- `/customers`

Bad:

- `/api/orders`
- `/api/v1/customers`

---

### Z-136: Use normalized paths without trailing slashes

**Severity:** error

Trailing slashes cause inconsistency and caching issues.

**Examples:**

Good:

- `/orders`
- `/orders/{id}`

Bad:

- `/orders/`
- `/orders/{id}/`

---

### Z-137: Stick to conventional query parameters

**Severity:** warn

Use conventional names for common query parameters:
- q: search query
- sort: sorting (e.g., sort=+name,-date)
- fields: sparse fieldsets
- embed: embedded resources
- offset/cursor: pagination
- limit: page size

**Rationale:** Standard query parameter names improve API consistency.

---

### Z-138: Allow optional embedding of sub-resources

**Severity:** info

Support optional embedding via 'embed' query parameter.

**Rationale:** Embedding reduces round trips while keeping default responses lightweight.

**Examples:**

Good:

- `/orders/{id}?embed=(items,customer)`

---

### Z-228: Use URL-friendly resource identifiers

**Severity:** error

Resource IDs must match regex [a-zA-Z0-9:._\-/]* - only ASCII letters, numbers, underscore, minus, colon, period, and slash (for compound keys).

**Rationale:** URL-friendly identifiers simplify encoding and prevent URL parsing issues.

**References:**

- [Zalando Rule 228](https://opensource.zalando.com/restful-api-guidelines/#228)

---

### Z-130: Use snake_case for query parameters

**Severity:** error

Query parameter names must use snake_case to match JSON property naming conventions.

**Rationale:** Consistent naming between query parameters and JSON properties improves developer experience.

**Examples:**

Good:

- `?customer_id=123`
- `?sort_order=asc`

Bad:

- `?customerId=123`
- `?sortOrder=asc`

**References:**

- [Zalando Rule 130](https://opensource.zalando.com/restful-api-guidelines/#130)

---

### Z-139: Model complete business processes

**Severity:** warn

An API should contain all resources representing a complete business process, preventing services from being thin wrappers around databases.

**Rationale:** Complete business process modeling enables client understanding and prevents thin database wrappers.

**References:**

- [Zalando Rule 139](https://opensource.zalando.com/restful-api-guidelines/#139)

---

### Z-140: Define useful resources

**Severity:** warn

Resources should contain as much information as necessary but as little as possible. Support filtering and embedding for edge cases.

**Rationale:** Resources should cover 90% of use cases with appropriate information density.

**References:**

- [Zalando Rule 140](https://opensource.zalando.com/restful-api-guidelines/#140)

---

### Z-141: Keep URLs verb-free

**Severity:** error

Use only nouns in URLs. Instead of verbs like 'cancel', model as resources (e.g., POST /cancellations).

**Rationale:** URLs describe resources, not actions; HTTP methods indicate actions.

**Examples:**

Good:

- `/orders/{id}/cancellation`
- `/cancellations`

Bad:

- `/orders/{id}/cancel`
- `/cancelOrder`

**References:**

- [Zalando Rule 141](https://opensource.zalando.com/restful-api-guidelines/#141)

---

### Z-142: Use domain-specific resource names

**Severity:** error

Use domain-specific nomenclature for resource names (e.g., 'sales-order-items' instead of generic 'items').

**Rationale:** Domain-specific names improve understanding and reduce documentation needs.

**Examples:**

Good:

- `/sales-order-items`
- `/shipment-orders`

Bad:

- `/items`
- `/orders`

**References:**

- [Zalando Rule 142](https://opensource.zalando.com/restful-api-guidelines/#142)

---

### Z-145: Consider using (non-) nested URLs

**Severity:** info

Use nested URLs for sub-resources with lifecycle coupled to parent. Expose resources with globally unique IDs at top level.

**Rationale:** URL structure should reflect resource relationships and access patterns.

**References:**

- [Zalando Rule 145](https://opensource.zalando.com/restful-api-guidelines/#145)

---

### Z-146: Limit number of resource types

**Severity:** warn

Follow functional segmentation and separation of concerns. Well-defined APIs typically have 4-8 resource types.

**Rationale:** Too many resource types increases complexity and maintenance burden.

**References:**

- [Zalando Rule 146](https://opensource.zalando.com/restful-api-guidelines/#146)

---

### Z-147: Limit number of sub-resource levels

**Severity:** warn

Use <= 3 sub-resource nesting levels. Remember URL length limits (some browsers cap at 2000 characters).

**Rationale:** Deep nesting increases complexity and URL length.

**Examples:**

Good:

- `/orders/{id}/items/{item-id}`

Bad:

- `/customers/{id}/orders/{oid}/items/{iid}/details/{did}`

**References:**

- [Zalando Rule 147](https://opensource.zalando.com/restful-api-guidelines/#147)

---

### Z-241: Expose compound keys as resource identifiers

**Severity:** info

Compound keys can be exposed in URLs using slashes (e.g., /shopping-carts/{country}/{session-id}). Apply consistently and provide compound key abstraction.

**Rationale:** Compound keys can simplify URL structure for resources with natural multi-part identifiers.

**References:**

- [Zalando Rule 241](https://opensource.zalando.com/restful-api-guidelines/#241)

---

## HTTP Methods

Correct use of GET, POST, PUT, PATCH, DELETE

### Z-148: Use HTTP methods correctly

**Severity:** error

Use HTTP methods according to RFC 9110 semantics:
- GET: Read resources (no body, cacheable)
- POST: Create resources or execute operations
- PUT: Full resource replacement
- PATCH: Partial update (use merge-patch or JSON Patch)
- DELETE: Remove resources

**Rationale:** Correct HTTP method semantics ensure APIs work predictably with caching and intermediaries.

---

### Z-149: Consider idempotency for write operations

**Severity:** warn

PUT and DELETE must be idempotent. POST should be idempotent when possible using idempotency keys.

**Rationale:** Idempotent operations can be safely retried, improving reliability.

---

### Z-154: Use JSON merge patch for partial updates

**Severity:** warn

For PATCH operations, prefer JSON Merge Patch with content-type application/merge-patch+json.

**Rationale:** JSON Merge Patch (RFC 7396) is simple and widely supported.

---

### Z-229: Use idempotency key for POST

**Severity:** warn

Support X-Idempotency-Key header for POST operations that create resources.

**Rationale:** Idempotency keys prevent duplicate resource creation on retries.

---

### Z-143: Return created resources with Location header

**Severity:** warn

POST that creates a resource should return 201 with Location header pointing to the new resource.

**Rationale:** Location header enables clients to find the created resource.

---

### Z-231: Use secondary key for idempotent POST design

**Severity:** warn

Design POST as idempotent using secondary unique key provided by client. Service returns 200 OK if resource exists, 201 Created for new creation.

**Rationale:** Secondary keys enable idempotent resource creation without client-generated IDs.

**References:**

- [Zalando Rule 231](https://opensource.zalando.com/restful-api-guidelines/#231)

---

### Z-236: Design simple query languages using query parameters

**Severity:** warn

For simple queries, use standard query parameters with patterns like: field={value}, field={op}:{value} (where op: eq, ne, lt, le, gt, ge), multi-select via field={value},{value}.

**Rationale:** Query parameters work well for simple filtering use cases.

**Examples:**

Good:

- `?status=active`
- `?price=gt:100`
- `?category=books,music`

Bad:

- `?filter=status eq 'active'`

**References:**

- [Zalando Rule 236](https://opensource.zalando.com/restful-api-guidelines/#236)

---

### Z-237: Design complex query languages using JSON

**Severity:** info

For complex queries with deep structures, use GET-with-body operations with JSON request bodies. Document using schema that mirrors property filter semantics.

**Rationale:** JSON bodies enable complex queries that don't fit in query strings.

**References:**

- [Zalando Rule 237](https://opensource.zalando.com/restful-api-guidelines/#237)

---

## HTTP Status Codes

Status codes and error handling

### Z-243: Use official HTTP status codes

**Severity:** error

Only use official HTTP status codes from RFC standards (RFC 9110, RFC 6585) consistently with their defined semantics.

**Rationale:** Official status codes are well understood and properly handled by clients and intermediaries.

---

### Z-151: Specify success and error responses

**Severity:** error

Define all success and service-specific error responses in the API specification. Standard errors (401, 403, 404, 500, 503) can use default description.

**Rationale:** Clear response documentation helps clients handle both success and error cases.

---

### Z-150: Use most common HTTP status codes

**Severity:** warn

Prefer common status codes:
- Success: 200, 201, 202, 204
- Client errors: 400, 401, 403, 404, 409, 429
- Server errors: 500, 503

**Rationale:** Common status codes are better understood; obscure codes cause confusion.

**Common HTTP Status Codes**

| Code | Name | When to Use |
| --- | --- | --- |
| 200 | OK | General success with response body |
| 201 | Created | Resource created (POST, PUT) |
| 202 | Accepted | Async processing started |
| 204 | No Content | Success without body (DELETE) |
| 400 | Bad Request | Invalid request format |
| 401 | Unauthorized | Authentication required |
| 403 | Forbidden | Permission denied |
| 404 | Not Found | Resource not found |
| 409 | Conflict | Resource state conflict |
| 429 | Too Many Requests | Rate limit exceeded |
| 500 | Internal Server Error | Unexpected server error |
| 503 | Service Unavailable | Temporary unavailability |

---

### Z-176: Use Problem JSON for errors

**Severity:** error

Error responses (4xx, 5xx) must use application/problem+json media type with RFC 7807 Problem structure:
- type: URI identifying error type
- title: Human-readable summary
- status: HTTP status code
- detail: Human-readable explanation
- instance: URI identifying specific occurrence

**Rationale:** RFC 7807 Problem+JSON provides a standard, machine-readable error format.

**Examples:**

Good:

- `{"type": "...", "title": "Bad Request", "status": 400, "detail": "..."}`

Bad:

- `{"error": "Something went wrong"}`

---

### Z-152: Use multi-status responses for batch operations

**Severity:** warn

For batch operations that can partially fail, return 207 Multi-Status with per-item status codes.

**Rationale:** Batch operations may partially succeed; multi-status response provides per-item status.

---

### Z-253: Use 202 for async resource creation

**Severity:** warn

When resource creation won't complete within the request, return 202 Accepted with a location to check status.

**Rationale:** 202 Accepted indicates the request was accepted but not yet completed.

---

### Z-220: Use most specific HTTP status codes

**Severity:** warn

Use the most specific HTTP status code that accurately describes the outcome. E.g., use 409 for conflicts instead of generic 400, use 422 for semantic errors.

**Rationale:** Specific status codes provide better error handling information to clients.

**References:**

- [Zalando Rule 220](https://opensource.zalando.com/restful-api-guidelines/#220)

---

### Z-177: Not expose stack traces

**Severity:** error

Error responses must not contain stack traces or internal exception details. Return safe, user-friendly error messages with appropriate detail levels.

**Rationale:** Stack traces expose internal implementation details and potential security vulnerabilities.

**References:**

- [Zalando Rule 177](https://opensource.zalando.com/restful-api-guidelines/#177)

---

### Z-251: Not use redirection codes

**Severity:** warn

Avoid using 3xx redirect status codes. Provide the correct URL in documentation. Use 301 only for permanent URL changes with communication to clients.

**Rationale:** Redirects can cause unexpected behavior and caching issues with REST clients.

**References:**

- [Zalando Rule 251](https://opensource.zalando.com/restful-api-guidelines/#251)

---

### Z-153: Use code 429 with headers for rate limits

**Severity:** error

Return 429 Too Many Requests for rate limiting with X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset, and Retry-After headers.

**Rationale:** Proper rate limiting responses help clients implement backoff strategies.

**References:**

- [Zalando Rule 153](https://opensource.zalando.com/restful-api-guidelines/#153)

---

## Pagination

Cursor-based and offset pagination patterns

### Z-159: Support pagination

**Severity:** error

All collection endpoints that may return many items must support pagination. Use cursor or offset-based pagination with consistent query parameters.

**Rationale:** Pagination protects services from overload and enables client-side iteration.

---

### Z-160: Prefer cursor-based pagination

**Severity:** warn

Prefer cursor-based pagination over offset-based pagination. Cursors are opaque tokens that encode page position. Offset pagination is acceptable when:
- Jumping to specific pages is required
- Dataset is small and stable

**Rationale:** Cursor-based pagination performs better on large datasets and NoSQL databases.

---

### Z-248: Use pagination response page object

**Severity:** warn

Use consistent pagination response with fields:
- items: Array of resources
- self: Current page link/cursor
- next: Next page link/cursor (omit on last page)
- prev: Previous page link/cursor (omit on first page)
- first, last: Optional boundary links

**Rationale:** Consistent pagination response structure improves usability.

**Examples:**

Good:

- `{"items": [...], "self": "...", "next": "..."}`

---

### Z-161: Use pagination links where applicable

**Severity:** warn

Replace plain cursor values with full pagination links when possible.

**Rationale:** Full URLs are easier for clients to use than constructing URLs from cursors.

---

### Z-254: Avoid a total result count

**Severity:** warn

Avoid providing total result counts in pagination responses. If required, support via Prefer: return=total-count header.

**Rationale:** Total counts are expensive to compute and rarely needed.

**References:**

- [Zalando Rule 254](https://opensource.zalando.com/restful-api-guidelines/#254)

---

## HTTP Headers

Standard and custom header usage

### Z-183: Use only headers with defined semantics

**Severity:** warn

Use standard headers for their intended purpose. Avoid using headers for operation-specific data that belongs in the URL or body.

**Rationale:** Headers should only carry general context information, not operation-specific data.

---

### Z-180: Use standardized proprietary header names

**Severity:** warn

Custom headers should use prefixes:
- X-Flow-ID: Request tracing
- X-Tenant-ID: Multi-tenant identification
- X-Sales-Channel: Business context

**Rationale:** Standard prefixes help identify custom headers and their scope.

---

### Z-182: Support conditional requests with ETag

**Severity:** warn

Support ETag header for resources that can be updated. Use If-Match/If-None-Match for conditional requests.

**Rationale:** ETags enable conditional requests and optimistic locking.

---

### Z-132: Use kebab-case with uppercase separate words for HTTP headers

**Severity:** warn

HTTP header names should use Kebab-Case-With-Uppercase-Separate-Words (e.g., If-Modified-Since, Content-Type).

**Rationale:** Consistent header naming following standard conventions improves readability.

**Examples:**

Good:

- `If-Modified-Since`
- `Accept-Encoding`
- `Content-ID`

Bad:

- `if-modified-since`
- `ACCEPT-ENCODING`
- `contentId`

**References:**

- [Zalando Rule 132](https://opensource.zalando.com/restful-api-guidelines/#132)

---

### Z-133: Use standard headers

**Severity:** info

Use standard HTTP headers defined in non-obsolete RFCs where applicable. Custom headers should be documented explicitly.

**Rationale:** Standard HTTP headers are widely understood and supported by infrastructure.

**References:**

- [Zalando Rule 133](https://opensource.zalando.com/restful-api-guidelines/#133)

---

### Z-178: Use Content-* headers correctly

**Severity:** error

Use Content-* headers (Content-Type, Content-Length, Content-Encoding, Content-Language, Content-Location, Content-Range, Content-Disposition) correctly to describe response body content.

**Rationale:** Content headers describe the body and must be used consistently.

**References:**

- [Zalando Rule 178](https://opensource.zalando.com/restful-api-guidelines/#178)

---

### Z-179: Use Content-Location header

**Severity:** info

Content-Location can indicate the actual location of the transmitted resource. Useful for content negotiation and cache support.

**Rationale:** Content-Location guides caching and signals actual resource location.

**References:**

- [Zalando Rule 179](https://opensource.zalando.com/restful-api-guidelines/#179)

---

### Z-181: Consider to support Prefer header

**Severity:** info

Support RFC 7240 Prefer header for client preferences like return=minimal, respond-async, wait, handling=lenient/strict.

**Rationale:** Prefer header enables client-requested processing behaviors.

**References:**

- [Zalando Rule 181](https://opensource.zalando.com/restful-api-guidelines/#181)

---

### Z-230: Consider to support Idempotency-Key header

**Severity:** info

Support Idempotency-Key header for POST/PATCH operations to ensure idempotent behavior. Store key with response for 24 hours; return cached response on retry.

**Rationale:** Idempotency-Key enables safe retries for non-idempotent operations.

**References:**

- [Zalando Rule 230](https://opensource.zalando.com/restful-api-guidelines/#230)

---

### Z-184: Propagate proprietary headers

**Severity:** error

All Zalando proprietary headers (X-Flow-ID, X-Tenant-ID, etc.) are end-to-end and must be propagated to downstream services unchanged.

**Rationale:** End-to-end headers must be propagated to enable tracing and context passing.

**References:**

- [Zalando Rule 184](https://opensource.zalando.com/restful-api-guidelines/#184)

---

### Z-233: Support X-Flow-ID

**Severity:** error

Services must support X-Flow-ID header for request correlation. Generate new Flow-ID if not provided. Propagate to all downstream calls and events.

**Rationale:** Flow IDs enable request tracing across distributed services.

**References:**

- [Zalando Rule 233](https://opensource.zalando.com/restful-api-guidelines/#233)

---

## Hypermedia

Links and HATEOAS patterns

### Z-217: Use hypermedia for navigation

**Severity:** info

Include hypermedia links to related resources and available actions.

**Rationale:** Hypermedia links enable discoverable APIs and reduce client coupling.

---

### Z-162: Use REST maturity level 2

**Severity:** error

Implement REST Maturity Level 2: resource-oriented APIs using HTTP verbs correctly (GET for retrieval, POST for creation, etc.) and proper status codes.

**Rationale:** REST level 2 enables resource-oriented APIs with proper HTTP semantics.

**References:**

- [Zalando Rule 162](https://opensource.zalando.com/restful-api-guidelines/#162)

---

### Z-163: Consider REST maturity level 3 - HATEOAS

**Severity:** info

HATEOAS is not generally recommended due to added complexity without clear benefits in SOA contexts. May be useful for specific human-facing scenarios.

**Rationale:** HATEOAS adds complexity with limited value in typical SOA contexts.

**References:**

- [Zalando Rule 163](https://opensource.zalando.com/restful-api-guidelines/#163)

---

### Z-164: Use common hypertext controls

**Severity:** error

Use common hypertext control objects with href attribute for resource links. Follow IANA link relations, converting hyphens to underscores.

**Rationale:** Consistent hypertext controls enable reliable link handling.

**Examples:**

Good:

- `{"href": "https://api.example.com/orders/123"}`

Bad:

- `{"link": "https://api.example.com/orders/123"}`

**References:**

- [Zalando Rule 164](https://opensource.zalando.com/restful-api-guidelines/#164)

---

### Z-165: Use simple hypertext controls for pagination and self-references

**Severity:** warn

For pagination and self-references, use simple URI strings with standard link relations (next, prev, first, last, self).

**Rationale:** Simple URI values reduce overhead for common pagination patterns.

**References:**

- [Zalando Rule 165](https://opensource.zalando.com/restful-api-guidelines/#165)

---

### Z-166: Not use link headers with JSON entities

**Severity:** error

Do not use RFC 8288 Link headers with JSON responses. Embed links directly in the JSON payload for flexibility and precision.

**Rationale:** Embedded links in JSON are more flexible than HTTP Link headers.

**References:**

- [Zalando Rule 166](https://opensource.zalando.com/restful-api-guidelines/#166)

---

## Deprecation

Deprecation and sunset handling

### Z-187: Reflect deprecation in API specifications

**Severity:** error

Mark deprecated elements with OpenAPI 'deprecated: true' and provide migration guidance in description.

**Rationale:** Clear deprecation signals help clients migrate before removal.

---

### Z-189: Add Sunset header for deprecated APIs

**Severity:** warn

Include Sunset HTTP header with the date when the API will be removed.

**Rationale:** Sunset header (RFC 8594) programmatically signals deprecation timeline.

**Examples:**

Good:

- `Sunset: Sat, 31 Dec 2024 23:59:59 GMT`

---

### Z-185: Obtain approval of clients before API shut down

**Severity:** error

Before shutting down an API, obtain consent from all clients on a sunset date. Help consumers migrate with migration manuals.

**Rationale:** Uncoordinated shutdowns break clients and cause production incidents.

**References:**

- [Zalando Rule 185](https://opensource.zalando.com/restful-api-guidelines/#185)

---

### Z-186: Collect external partner consent on deprecation time span

**Severity:** error

For APIs consumed by external partners, define and communicate a reasonable after-deprecation-life-span. Partners must consent before using the API.

**Rationale:** External partners need adequate time to migrate.

**References:**

- [Zalando Rule 186](https://opensource.zalando.com/restful-api-guidelines/#186)

---

### Z-188: Monitor usage of deprecated API scheduled for sunset

**Severity:** error

Monitor usage of deprecated APIs to track migration progress and avoid breaking clients still using the sunset API.

**Rationale:** Usage monitoring prevents breaking production traffic during sunset.

**References:**

- [Zalando Rule 188](https://opensource.zalando.com/restful-api-guidelines/#188)

---

### Z-190: Add monitoring for Deprecation and Sunset header

**Severity:** warn

Clients should monitor Deprecation and Sunset headers in responses and alert on their presence.

**Rationale:** Clients should proactively detect API deprecations.

**References:**

- [Zalando Rule 190](https://opensource.zalando.com/restful-api-guidelines/#190)

---

### Z-191: Not start using deprecated APIs

**Severity:** error

Clients must not start using deprecated APIs, API versions, or API features.

**Rationale:** New integrations with deprecated APIs create unnecessary migration work.

**References:**

- [Zalando Rule 191](https://opensource.zalando.com/restful-api-guidelines/#191)

---

## Security

Authentication and authorization

### Z-104: Secure endpoints with authentication

**Severity:** error

All API endpoints must be secured with proper authentication. Use OAuth 2.0 with appropriate flows.

**Rationale:** APIs must be secured to protect data and prevent unauthorized access.

---

### Z-105: Define and assign permissions (scopes)

**Severity:** error

Define OAuth scopes and assign them to API operations. Document required permissions for each endpoint.

**Rationale:** Scopes enable fine-grained access control and follow least-privilege principle.

---

### Z-227: Use HTTPS

**Severity:** error

HTTPS protects data in transit from interception.

---

### Z-225: Follow naming convention for permissions (scopes)

**Severity:** error

Permission names must follow pattern: <application-id>.<access-mode> or <application-id>.<resource-name>.<access-mode>. Use 'uid' for unrestricted access.

**Rationale:** Consistent permission naming enables authorization management.

**Examples:**

Good:

- `order-management.read`
- `order-management.sales-order.write`

Bad:

- `read_orders`
- `ORDER_MANAGEMENT_READ`

**References:**

- [Zalando Rule 225](https://opensource.zalando.com/restful-api-guidelines/#225)

---

## Performance

Performance considerations

### Z-226: Design for performance

**Severity:** warn

Consider performance in API design:
- Support partial responses ($fields)
- Enable caching with ETags
- Use pagination for collections
- Consider compression

**Rationale:** Performance-conscious API design improves user experience and reduces costs.

---

### Z-155: Reduce bandwidth needs and improve responsiveness

**Severity:** warn

Support bandwidth reduction techniques: compression, field filtering, ETag/If-Match headers, Prefer header with return=minimal, pagination, and caching.

**Rationale:** Bandwidth optimization is critical for mobile clients.

**References:**

- [Zalando Rule 155](https://opensource.zalando.com/restful-api-guidelines/#155)

---

### Z-156: Use gzip compression

**Severity:** warn

Servers and clients should support gzip content encoding via Accept-Encoding and Content-Encoding headers. Document compression support.

**Rationale:** Compression significantly reduces payload size for text-based responses.

**References:**

- [Zalando Rule 156](https://opensource.zalando.com/restful-api-guidelines/#156)

---

### Z-157: Support partial responses via filtering

**Severity:** warn

Support 'fields' query parameter for partial responses. Syntax: fields=(name,friends(name)) for nested filtering.

**Rationale:** Field filtering reduces payload size when clients need subset of data.

**Examples:**

Good:

- `?fields=(name,email)`
- `?fields=(id,items(sku,price))`

Bad:

- `?select=name,email`

**References:**

- [Zalando Rule 157](https://opensource.zalando.com/restful-api-guidelines/#157)

---

### Z-158: Allow optional embedding of sub-resources

**Severity:** warn

Support 'embed' query parameter for sub-resource expansion. Syntax: embed=(items) to include order items with order response.

**Rationale:** Resource embedding reduces N+1 query problems.

**Examples:**

Good:

- `?embed=(items)`
- `?embed=(customer,items)`

Bad:

- `?include=items`

**References:**

- [Zalando Rule 158](https://opensource.zalando.com/restful-api-guidelines/#158)

---

## Events

### Z-194: Treat events as part of the service interface

**Severity:** error

Events are part of a service's interface equivalent to REST APIs. Design events with API-first principle in mind.

**Rationale:** Events are a first-class contract with the same care as REST APIs.

**References:**

- [Zalando Rule 194](https://opensource.zalando.com/restful-api-guidelines/#194)

---

### Z-195: Make event schema available for review

**Severity:** error

Services publishing events must make event schema and event type definition available for consumer review.

**Rationale:** Event schemas must be reviewable by consumers.

**References:**

- [Zalando Rule 195](https://opensource.zalando.com/restful-api-guidelines/#195)

---

### Z-196: Ensure event schema conforms to OpenAPI schema object

**Severity:** error

Event schemas should use OpenAPI Schema Object specification (extended JSON Schema Draft 4). Avoid additionalItems, contains, patternProperties, dependencies, propertyNames, const, not, oneOf.

**Rationale:** OpenAPI schema alignment enables tooling consistency.

**References:**

- [Zalando Rule 196](https://opensource.zalando.com/restful-api-guidelines/#196)

---

### Z-197: Specify and register events as event types

**Severity:** error

Register events as Event Types with: name, category (data/general), owning_application, schema, compatibility_mode, and audience.

**Rationale:** Event type registration enables discoverability and validation.

**References:**

- [Zalando Rule 197](https://opensource.zalando.com/restful-api-guidelines/#197)

---

### Z-198: Ensure that events conform to an event category

**Severity:** error

Events must conform to either General Event (business processes) or Data Change Event (entity mutations) category with required metadata.

**Rationale:** Event categories provide standard structure and metadata.

**References:**

- [Zalando Rule 198](https://opensource.zalando.com/restful-api-guidelines/#198)

---

### Z-199: Ensure that events define useful business resources

**Severity:** warn

Events should be based on business resources and processes. Avoid explosion of event types; prefer abstract/generic types valuable for multiple use cases.

**Rationale:** Events should represent meaningful business value.

**References:**

- [Zalando Rule 199](https://opensource.zalando.com/restful-api-guidelines/#199)

---

### Z-200: Avoid writing sensitive data to events

**Severity:** warn

Avoid writing sensitive data (personal data, PII) to events unless required for business purposes.

**Rationale:** Sensitive data in events increases compliance and security risks.

**References:**

- [Zalando Rule 200](https://opensource.zalando.com/restful-api-guidelines/#200)

---

### Z-201: Use general events to signal steps in business processes

**Severity:** error

Use General Event category for business process steps. Include business process ID, ordering information (parent_eids for causality), and only new information per step.

**Rationale:** General events enable business process tracking and replay.

**References:**

- [Zalando Rule 201](https://opensource.zalando.com/restful-api-guidelines/#201)

---

### Z-202: Use data change events to signal mutations

**Severity:** error

Use Data Change Event category for entity mutations (create, update, delete). Provide complete entity data, entity identifier, and data_op field.

**Rationale:** Data change events enable change data capture and replication.

**References:**

- [Zalando Rule 202](https://opensource.zalando.com/restful-api-guidelines/#202)

---

### Z-203: Provide explicit event ordering for general events

**Severity:** warn

Provide ordering_key_fields and optionally ordering_instance_ids in event type definition. Use monotonically increasing versions or sequence counters.

**Rationale:** Ordering enables event stream reconstruction and replay.

**References:**

- [Zalando Rule 203](https://opensource.zalando.com/restful-api-guidelines/#203)

---

### Z-204: Use the hash partition strategy for data change events

**Severity:** warn

Use hash partition strategy with entity identifier as key. This ensures all events for an entity go to the same partition for ordered processing.

**Rationale:** Hash partitioning ensures ordered delivery per entity.

**References:**

- [Zalando Rule 204](https://opensource.zalando.com/restful-api-guidelines/#204)

---

### Z-205: Ensure data change events match APIs resources

**Severity:** warn

Data change event entity representation should correspond to REST API resource representation where possible.

**Rationale:** Consistent representations reduce consumer complexity.

**References:**

- [Zalando Rule 205](https://opensource.zalando.com/restful-api-guidelines/#205)

---

### Z-207: Indicate ownership of event types

**Severity:** error

Event types must have clear ownership via owning_application field. Owner is responsible for definition and evolution.

**Rationale:** Clear ownership enables accountability and coordination.

**References:**

- [Zalando Rule 207](https://opensource.zalando.com/restful-api-guidelines/#207)

---

### Z-208: Define events compliant with overall API guidelines

**Severity:** error

Events must follow API guidelines for general conventions, JSON guidelines, data formats, and hypermedia principles.

**Rationale:** Events should follow same quality standards as REST APIs.

**References:**

- [Zalando Rule 208](https://opensource.zalando.com/restful-api-guidelines/#208)

---

### Z-209: Maintain backwards compatibility for events

**Severity:** error

Event changes must be additive and backward compatible. Compatible changes: new optional fields, field order changes. Incompatible: removing required fields, changing types.

**Rationale:** Events need even stricter compatibility than REST APIs due to async nature.

**References:**

- [Zalando Rule 209](https://opensource.zalando.com/restful-api-guidelines/#209)

---

### Z-210: Avoid additionalProperties in event type schemas

**Severity:** warn

Event schemas should avoid additionalProperties: true. Define new optional fields explicitly. Consumers must ignore unknown fields.

**Rationale:** Wildcard extensions hinder schema evolution.

**References:**

- [Zalando Rule 210](https://opensource.zalando.com/restful-api-guidelines/#210)

---

### Z-211: Provide unique event identifiers

**Severity:** error

Provide eid (event identifier) as UUID in event metadata. Use same eid for retries of the same event.

**Rationale:** Event identifiers enable deduplication and idempotent processing.

**References:**

- [Zalando Rule 211](https://opensource.zalando.com/restful-api-guidelines/#211)

---

### Z-212: Design for idempotent out-of-order processing

**Severity:** warn

Design events for idempotent out-of-order processing by including entity identifier, monotonically increasing ordering key, and state after change.

**Rationale:** Idempotent processing enables resilient event consumption.

**References:**

- [Zalando Rule 212](https://opensource.zalando.com/restful-api-guidelines/#212)

---

### Z-213: Follow naming convention for event type names

**Severity:** error

Event type names must follow pattern: <functional-name>.<event-name>[.<version>]. Use functional naming for external events.

**Rationale:** Consistent event naming enables discoverability.

**Examples:**

Good:

- `transactions-order.order-cancelled`
- `customer-personal-data.email-changed.v2`

Bad:

- `order_cancelled`
- `OrderCancelledEvent`

**References:**

- [Zalando Rule 213](https://opensource.zalando.com/restful-api-guidelines/#213)

---

### Z-214: Be robust against duplicates when consuming events

**Severity:** error

Event consumers must be robust against duplicates. Use eid for deduplication or leverage data keys and ordering for CDC.

**Rationale:** Duplicate events are inherent in distributed systems.

**References:**

- [Zalando Rule 214](https://opensource.zalando.com/restful-api-guidelines/#214)

---

### Z-242: Provide explicit event ordering for data change events

**Severity:** error

Data change events must provide explicit ordering information via ordering_key_fields and ordering_instance_ids.

**Rationale:** Data change ordering is critical for CDC and replication.

**References:**

- [Zalando Rule 242](https://opensource.zalando.com/restful-api-guidelines/#242)

---

### Z-245: Carefully define the compatibility mode

**Severity:** error

Choose compatibility mode carefully: none (any change), forward (consumers can read with old schema), compatible (all versions validate).

**Rationale:** Compatibility mode affects schema evolution flexibility.

**References:**

- [Zalando Rule 245](https://opensource.zalando.com/restful-api-guidelines/#245)

---

### Z-246: Use semantic versioning of event type schemas

**Severity:** error

Event schemas must follow semantic versioning (MAJOR.MINOR.PATCH). PATCH for description changes, MINOR for optional fields, MAJOR for breaking changes.

**Rationale:** Semantic versioning communicates impact of schema changes.

**References:**

- [Zalando Rule 246](https://opensource.zalando.com/restful-api-guidelines/#246)

---

### Z-247: Provide mandatory event metadata

**Severity:** error

Events must include metadata with eid (event identifier) and occurred_at (event creation timestamp). Optional: event_type, version, parent_eids, flow_id, partition.

**Rationale:** Metadata enables event tracking and processing.

**References:**

- [Zalando Rule 247](https://opensource.zalando.com/restful-api-guidelines/#247)

---

## Glossary

**API First**
: Development approach where APIs are designed before implementation, using specifications as the source of truth.

**OpenAPI**
: The standard specification format for REST APIs, formerly known as Swagger.

**snake_case**
: Naming convention using lowercase letters with underscores between words (e.g., customer_number).

**Cursor**
: An opaque token that identifies a position in a paginated collection.

**Problem+JSON**
: RFC 7807 standard for machine-readable error responses (application/problem+json).

**Zally**
: Zalando's API linter that validates OpenAPI specifications against these guidelines.

**x-extensible-enum**
: OpenAPI extension for enums that may have additional values added in the future.

**Sunset Header**
: HTTP header indicating when an API or endpoint will be deprecated (RFC 8594).


# Using Profiles

Profiles are collections of API style rules. api-style-spec includes profiles based on industry-standard guidelines.

## Built-in Profiles

| Profile | Based On | Rules | Focus |
|---------|----------|-------|-------|
| `default` | Common best practices | ~30 | General REST API design |
| `azure` | Microsoft/Azure guidelines | ~25 | Enterprise cloud APIs |
| `google` | Google API Design Guide | ~20 | Resource-oriented design |
| `zalando` | Zalando RESTful Guidelines | ~30 | E-commerce APIs |

## Selecting a Profile

Use the `--profile` flag with any command:

```bash
# Lint with Azure profile
api-style lint openapi.yaml --profile azure

# Analyze with Google profile
api-style analyze openapi.yaml --profile google

# Generate documentation for Zalando profile
api-style generate guide --profile zalando
```

## Profile Details

### Default Profile

The default profile covers fundamental REST API best practices:

- **URI Design** - Plural resources, lowercase, hyphens
- **HTTP Methods** - Proper verb usage
- **Status Codes** - Appropriate response codes
- **Versioning** - API version strategies
- **Documentation** - Required descriptions

```bash
api-style lint openapi.yaml --profile default
```

### Azure Profile

Based on [Microsoft Azure REST API Guidelines](https://github.com/microsoft/api-guidelines):

- **Consistency** - Uniform interface patterns
- **Versioning** - Date-based API versions
- **Error Handling** - Standard error response format
- **Pagination** - Consistent collection handling
- **Long-running Operations** - Async patterns

```bash
api-style lint openapi.yaml --profile azure
```

### Google Profile

Based on [Google API Design Guide](https://cloud.google.com/apis/design):

- **Resource-Oriented Design** - Everything is a resource
- **Standard Methods** - List, Get, Create, Update, Delete
- **Naming Conventions** - Consistent field names
- **Error Model** - Structured error details
- **Documentation** - Comprehensive descriptions

```bash
api-style lint openapi.yaml --profile google
```

### Zalando Profile

Based on [Zalando RESTful API Guidelines](https://opensource.zalando.com/restful-api-guidelines/):

- **Pragmatic REST** - Practical over dogmatic
- **Hypermedia** - HATEOAS where appropriate
- **Events** - Async event patterns
- **Compatibility** - Evolution strategies
- **Security** - OAuth2 patterns

```bash
api-style lint openapi.yaml --profile zalando
```

## Listing Profile Rules

View all rules in a profile:

```bash
# List all rules
api-style list-rules --profile azure

# Filter by category
api-style list-rules --profile azure --category uri-design

# Filter by severity
api-style list-rules --profile azure --severity error
```

## Comparing Profiles

Different profiles may have different opinions on the same topic:

| Topic | Default | Azure | Google | Zalando |
|-------|---------|-------|--------|---------|
| Resource names | Plural | Plural | Plural | Plural |
| Case style | kebab-case | camelCase | snake_case | kebab-case |
| Versioning | URL path | URL path | URL path | URL/Header |
| Date format | ISO 8601 | ISO 8601 | RFC 3339 | ISO 8601 |

## Conformance Levels

Profiles support graduated compliance levels:

| Level | Description |
|-------|-------------|
| Bronze | Minimum viable API quality |
| Silver | Production-ready APIs |
| Gold | Best-in-class API design |

```bash
# Check against silver level
api-style lint openapi.yaml --profile azure --level silver
```

Rules are tagged with minimum conformance levels. Higher levels include all rules from lower levels.

## Next Steps

- [Custom Rules](custom-rules.md) - Create your own rules
- [Profile Reference](../profiles/default.md) - Detailed rule documentation

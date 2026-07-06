# API Style Guide Comparison

This document compares the built-in API style guide profiles, highlighting their strengths, coverage gaps, and recommended use cases.

## Quick Comparison

| Profile | Rules | Categories | Best For |
|---------|-------|------------|----------|
| **zalando** | 147 | 13 | E-commerce, event-driven |
| **microsoft-rest** | 123 | 15 | Enterprise REST APIs |
| **comprehensive** | 88 | 26 | Full coverage, new APIs |
| **microsoft-graph** | 82 | 12 | OData/Graph APIs |
| **default** | 79 | 26 | General-purpose, recommended |
| **minimal** | 29 | 7 | Basic hygiene |
| **azure** | 23 | 9 | Azure cloud services |
| **google** | 20 | 7 | Resource-oriented design |

## Coverage Matrix

The following matrix shows rule count by category for each profile. A dash (-) indicates no coverage.

| Category | comprehensive | zalando | ms-rest | ms-graph | default | minimal | azure | google |
|----------|--------------|---------|---------|----------|---------|---------|-------|--------|
| **general** | 3 | 13 | - | - | 1 | - | - | - |
| **naming** | 5 | - | 4 | 3 | 1 | 4 | 2 | 3 |
| **urls** | 4 | 16 | - | - | 2 | - | - | - |
| **http-methods** | 6 | 8 | 13 | - | 3 | 5 | 3 | - |
| **http-status** | 9 | 10 | - | - | 9 | - | - | - |
| **request-response** | 3 | - | 11 | - | 6 | - | 3 | - |
| **headers** | 3 | 11 | 12 | - | 3 | - | - | - |
| **errors** | 3 | - | 10 | - | 2 | 3 | 2 | 2 |
| **pagination** | 3 | 5 | 6 | - | 2 | - | 2 | - |
| **filtering** | 3 | - | 6 | - | 3 | - | - | - |
| **versioning** | 2 | - | 4 | - | 3 | - | 2 | 2 |
| **compatibility** | 3 | 9 | 4 | - | 3 | - | - | - |
| **deprecation** | 3 | 7 | - | - | 3 | - | - | - |
| **security** | 4 | 4 | 9 | - | 4 | 4 | 3 | - |
| **documentation** | 4 | - | - | - | 4 | 5 | - | 3 |
| **json** | 3 | 18 | - | - | 3 | - | - | - |
| **schema** | 3 | - | - | - | 3 | - | - | - |
| **collections** | 2 | - | 5 | - | 2 | - | - | - |
| **long-running** | 3 | - | 7 | - | 3 | - | 2 | - |
| **conditional** | 3 | - | 5 | - | 3 | - | - | - |
| **performance** | 3 | 5 | - | - | 3 | - | - | - |
| **hypermedia** | 2 | 6 | - | - | 2 | - | - | - |
| **batch** | 3 | - | - | 5 | 3 | - | - | - |
| **events** | 3 | 24 | - | - | 3 | - | - | - |
| **actions** | 2 | - | - | - | 2 | - | - | - |
| **throttling** | 3 | - | - | 5 | 3 | - | - | - |

## Profile Strengths & Gaps

### Comprehensive Profile

**Strengths:**

- Full category coverage (26/26 categories)
- Balanced rule distribution
- Synthesized best practices from all major guides
- Clear conformance levels (minimum/standard/exemplary)
- Machine-enforceable rules with Spectral

**Gaps:**

- Fewer rules per category than specialized profiles
- Less depth in domain-specific areas (events, OData)

**Best for:** New API projects needing full coverage without bias toward specific ecosystem.

---

### Zalando Profile

**Strengths:**

- Deepest coverage of events/webhooks (24 rules)
- Strong JSON conventions (18 rules)
- Excellent URL design guidance (16 rules)
- Comprehensive compatibility rules (9 rules)
- Well-defined deprecation process (7 rules)

**Gaps:**

- No request-response rules
- No error handling rules
- No filtering/sorting rules
- No documentation rules
- Limited versioning guidance

**Best for:** E-commerce APIs, event-driven architectures, APIs requiring strong backward compatibility.

---

### Microsoft REST Profile

**Strengths:**

- Most comprehensive HTTP methods coverage (13 rules)
- Excellent headers guidance (12 rules)
- Strong request-response patterns (11 rules)
- Detailed error handling (10 rules)
- Good security coverage (9 rules)
- Long-running operation patterns (7 rules)

**Gaps:**

- No URL design rules
- No http-status rules
- No JSON convention rules
- No documentation rules
- No events/webhooks guidance

**Best for:** Enterprise REST APIs, Microsoft ecosystem, complex CRUD operations.

---

### Microsoft Graph Profile

**Strengths:**

- Specialized OData conventions
- Batch operation patterns (5 rules)
- Rate limiting/throttling (5 rules)
- Delta query patterns
- Navigation properties

**Gaps:**

- Very narrow focus (OData only)
- No coverage of most standard REST categories
- Not suitable for non-OData APIs

**Best for:** Microsoft Graph-compatible APIs, OData services only.

---

### Default Profile (PlexusOne)

**Strengths:**

- Full category coverage (26/26 categories)
- Synthesized best practices from Microsoft, Zalando, Google, PayPal
- Opinionated defaults: camelCase, kebab-case URLs, cursor pagination, RFC 9457
- Complete lifecycle coverage (versioning, deprecation, compatibility)
- Events/webhooks with CloudEvents format
- Long-running operations with LRO pattern
- Rate limiting and throttling guidance
- Clear conformance levels (minimum/standard/exemplary)

**Trade-offs:**

- Fewer rules per category than specialized profiles
- Less depth in domain-specific areas (e-commerce events, OData)
- Opinionated choices may differ from team conventions

**Best for:** General-purpose REST APIs, new projects, teams wanting full coverage with sensible defaults.

---

### Minimal Profile

**Strengths:**

- Lightweight, fast to validate
- Covers essential hygiene
- Good starting point for custom profiles

**Gaps:**

- Only 7 categories covered
- Missing most advanced topics
- No guidance for complex scenarios

**Best for:** Quick validation, CI pipelines, extending with custom rules.

---

### Azure Profile

**Strengths:**

- Date-based versioning patterns
- Long-running operation (LRO) patterns
- Azure-specific error format
- OData-style pagination

**Gaps:**

- Limited category coverage (9 categories)
- No URL design rules
- No headers rules
- No JSON rules
- No events guidance

**Best for:** Azure cloud services, ARM-compatible APIs.

---

### Google Profile

**Strengths:**

- Resource-oriented design philosophy
- Standard methods (List, Get, Create, Update, Delete)
- Clear naming conventions
- Good documentation requirements

**Gaps:**

- Minimal coverage (7 categories, 20 rules)
- No HTTP methods rules
- No http-status rules
- No security rules
- No pagination rules

**Best for:** Google Cloud-style APIs, resource-oriented services.

## Recommendations by Use Case

| Use Case | Recommended Profile | Alternative |
|----------|---------------------|-------------|
| New API project | `comprehensive` | `default` |
| E-commerce platform | `zalando` | `comprehensive` |
| Enterprise SaaS | `microsoft-rest` | `comprehensive` |
| Microsoft ecosystem | `microsoft-graph` | `azure` |
| Azure services | `azure` | `microsoft-rest` |
| Google Cloud style | `google` | `comprehensive` |
| Quick validation | `minimal` | `default` |
| Event-driven APIs | `zalando` | `comprehensive` |
| Public APIs | `comprehensive` | `zalando` |

## Combining Profiles

For maximum coverage, consider:

1. Start with `comprehensive` as base
2. Add domain-specific rules from `zalando` for events
3. Add `microsoft-rest` patterns for enterprise features
4. Use `score-profile` to evaluate coverage

```bash
# Check profile coverage
api-style score-profile comprehensive

# Compare profiles
api-style score-profile zalando --format json
api-style score-profile microsoft-rest --format json
```

## See Also

- [Using Profiles](profiles.md) - Basic profile usage
- [Custom Rules](custom-rules.md) - Creating custom rules
- [Profile Scoring](profile-scoring.md) - Evaluating profile quality

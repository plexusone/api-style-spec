# Bug: Lint Command False Positives for Pattern-Based Rules

**Component**: `pkg/lint/` or `cmd/api-style/lint.go` — regex evaluation during linting  
**Severity**: High — causes false-positive lint failures for valid paths  
**Found by**: Saviynt API Style Guide team (linting `microservices-health-api-spec`)  
**Version**: Current HEAD (api-style v0.4.x)

## Problem

The `api-style lint` command reports violations for pattern-based rules (SAV-007, SAV-012, SAV-020) even when the spec is conformant. The generated spectral YAML is correct (verified by unit tests), so the issue is in how the lint command evaluates regex patterns against actual values.

## Reproduction

```bash
./api-style lint /path/to/openapi.yaml --profile /path/to/saviynt-api-style-guide.json
```

Reports:
```
- [SAV-012] Use kebab-case for path segments
  $.paths['/api/observe/v1/alive'] (line 32)
- [SAV-007] Use semantic versioning
  $.info.version (line 11)    # version: "1.0.0" — clearly valid semver
- [SAV-020] GET must not have request body: `value` must be falsy
  $.paths['/api/observe/v1/alive'].get   # spec has NO requestBody
```

All three are false positives:
- `/api/observe/v1/alive` IS valid kebab-case (all lowercase, hyphens/digits only)
- `1.0.0` IS valid semver matching `^[0-9]+\.[0-9]+\.[0-9]+$`  
- The spec has no `requestBody` on GET operations

## Root Cause Hypothesis

The lint command uses **vacuum** (`github.com/daveshanley/vacuum`) as its linting engine. The `buildVacuumRuleSet()` function in `pkg/lint/vacuum.go` translates `types.APIStyleSpec` rules into vacuum `model.Rule` objects.

The false positives likely occur in this translation layer — when regex patterns from the profile are passed to vacuum's `model.Rule.Then.FunctionOptions`, they may be double-escaped or improperly formatted for vacuum's `pattern` function.

The spectral YAML generator is confirmed correct (unit tests pass). The issue is specifically in how `buildVacuumRuleSet()` constructs the `functionOptions` map that vacuum evaluates.

## Test Case

The following test (added to `pkg/generate/spectral_regex_test.go`) passes, confirming the YAML generation is correct:

```go
func TestSpectral_RegexRoundtrip_MatchesRealPaths(t *testing.T) {
    // JSON profile → unmarshal → extract regex → compile → match real paths
    // Verifies that /api/observe/v1/alive matches the kebab-case regex
    // after going through the JSON deserialization pipeline.
}
```

The bug is downstream — in the lint evaluation, not in the spectral generation.

## Expected Fix Location

Look at `cmd/api-style/lint.go` or `pkg/lint/` for where `options.match` patterns are evaluated against actual OpenAPI spec values. The regex is likely being double-escaped or not compiled correctly before matching.

## Validation

After fix, this should produce ZERO violations for SAV-007 and SAV-012:

```bash
echo '{"openapi":"3.1.0","info":{"title":"T","description":"D","version":"1.0.0","contact":{"name":"X"},"x-api-id":"abc","x-audience":"internal"},"paths":{"/api/users/v1/accounts":{"get":{"operationId":"list","description":"D","responses":{"200":{"description":"OK"}}}}}}' > /tmp/test.json
./api-style lint /tmp/test.json --profile saviynt-rest
```

## Additional: SAV-016 and SAV-020 False Positives

**SAV-016** (no query params in path): Fires on paths like `/api/observe/v1/alive` which have no `?` — the linter may be matching the regex incorrectly (the `notMatch: "\\?"` pattern should only fire if there IS a `?` in the path).

**SAV-020** (GET no body): The `falsy` function with `options.match: "requestBody"` fires even when no `requestBody` key exists in the operation. The hand-maintained spectral uses `field: "requestBody"` + `function: falsy` (no functionOptions) which works correctly in Spectral. The generated approach using `functionOptions.match` may have different semantics in the lint engine.

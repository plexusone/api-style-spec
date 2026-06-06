// Package types defines the core data types for api-style-spec.
//
// These Go types are the source of truth for the api-style-spec format.
// JSON Schema is generated from these types using invopop/jsonschema.
//
// Main types:
//   - APIStyleSpec: Root type for a style specification
//   - Rule: Individual style rule with enforcement and judge criteria
//   - LintReport: Results from deterministic linting
//   - Violation: A single rule violation
package types

//go:generate go run ../../cmd/schemagen/main.go ../../schema

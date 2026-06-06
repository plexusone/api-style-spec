package types

// Rule defines a single API style guideline.
type Rule struct {
	// ID is a unique identifier for the rule (e.g., "URI-001").
	ID string `json:"id"`

	// Title is a short, descriptive name for the rule.
	Title string `json:"title"`

	// Category groups related rules (e.g., "uri-design", "naming", "security").
	Category string `json:"category"`

	// Severity indicates the importance of violations.
	Severity Severity `json:"severity"`

	// Scope defines what part of the spec this rule applies to.
	Scope Scope `json:"scope,omitempty"`

	// Rationale explains why this rule exists and its benefits.
	Rationale string `json:"rationale,omitempty"`

	// Examples provides good and bad usage patterns.
	Examples *Examples `json:"examples,omitempty"`

	// Enforcement defines deterministic checking configuration.
	Enforcement *Enforcement `json:"enforcement,omitempty"`

	// Judge defines LLM evaluation criteria for this rule.
	Judge *JudgeCriteria `json:"judge,omitempty"`

	// References links to external documentation.
	References []Reference `json:"references,omitempty"`

	// Tags are labels for filtering and grouping rules.
	Tags []string `json:"tags,omitempty"`

	// Recommended indicates if this rule is part of the recommended set.
	Recommended bool `json:"recommended,omitempty"`
}

// Examples provides good and bad usage patterns for a rule.
type Examples struct {
	// Good shows correct usage patterns.
	Good []string `json:"good,omitempty"`
	// Bad shows incorrect usage patterns.
	Bad []string `json:"bad,omitempty"`
}

// EnforcementType defines how a rule is enforced.
type EnforcementType string

const (
	// EnforcementSpectral uses Spectral/vacuum for linting.
	EnforcementSpectral EnforcementType = "spectral"
	// EnforcementCustom uses a custom Go function.
	EnforcementCustom EnforcementType = "custom"
	// EnforcementRegex uses regular expression matching.
	EnforcementRegex EnforcementType = "regex"
	// EnforcementNone means the rule is LLM-only (no deterministic check).
	EnforcementNone EnforcementType = "none"
)

// Enforcement defines deterministic rule checking configuration.
type Enforcement struct {
	// Type is the enforcement mechanism.
	Type EnforcementType `json:"type"`

	// Function is the Spectral function name (for type=spectral).
	Function string `json:"function,omitempty"`

	// Options are function-specific configuration options.
	Options *EnforcementOptions `json:"options,omitempty"`

	// Given is the JSONPath expression(s) for targeting nodes (Spectral-style).
	// Can be a single string or array of strings.
	Given *GivenPaths `json:"given,omitempty"`

	// Then defines the assertion to apply (Spectral-style).
	Then *SpectralThen `json:"then,omitempty"`

	// Pattern is the regex pattern (for type=regex).
	Pattern string `json:"pattern,omitempty"`

	// CustomFunction is the name of a custom Go function (for type=custom).
	CustomFunction string `json:"customFunction,omitempty"`
}

// SpectralThen defines a Spectral-compatible assertion.
type SpectralThen struct {
	// Field is the field to check within the matched node.
	Field string `json:"field,omitempty"`
	// Function is the assertion function to apply.
	Function string `json:"function,omitempty"`
	// FunctionOptions are options for the function.
	FunctionOptions map[string]string `json:"functionOptions,omitempty"`
}

// GivenPaths represents JSONPath expressions for Spectral rules.
// Can be marshaled as a single string or array of strings.
type GivenPaths struct {
	// Paths contains one or more JSONPath expressions.
	Paths []string `json:"paths"`
}

// NewGivenPath creates a GivenPaths with a single path.
func NewGivenPath(path string) *GivenPaths {
	return &GivenPaths{Paths: []string{path}}
}

// NewGivenPaths creates a GivenPaths with multiple paths.
func NewGivenPaths(paths ...string) *GivenPaths {
	return &GivenPaths{Paths: paths}
}

// EnforcementOptions contains common options for enforcement functions.
type EnforcementOptions struct {
	// Match is a regex pattern to match against (for pattern function).
	Match string `json:"match,omitempty"`
	// NotMatch is a regex pattern that should not match.
	NotMatch string `json:"notMatch,omitempty"`
	// Min is a minimum value (for length function).
	Min *int `json:"min,omitempty"`
	// Max is a maximum value (for length function).
	Max *int `json:"max,omitempty"`
	// Values is a list of allowed values (for enumeration function).
	Values []string `json:"values,omitempty"`
	// Type specifies the expected casing type (for casing function).
	// Values: flat, camel, pascal, kebab, cobol, snake, macro
	Type string `json:"type,omitempty"`
	// Separator is used for casing validation.
	Separator string `json:"separator,omitempty"`
	// Schema is a JSON Schema for validation (for schema function).
	Schema string `json:"schema,omitempty"`
}

// JudgeCriteria defines LLM evaluation parameters for a rule.
type JudgeCriteria struct {
	// Prompt is the evaluation instruction for the LLM.
	Prompt string `json:"prompt"`

	// Weight influences scoring (0.0-1.0, default 1.0).
	Weight float64 `json:"weight,omitempty"`

	// RequiresContext indicates if broader context is needed for evaluation.
	RequiresContext bool `json:"requiresContext,omitempty"`

	// Category overrides the rule's category for evaluation grouping.
	Category string `json:"category,omitempty"`
}

// Reference links to external documentation.
type Reference struct {
	// Title is the display text for the reference.
	Title string `json:"title"`
	// URL is the link to the external resource.
	URL string `json:"url"`
}

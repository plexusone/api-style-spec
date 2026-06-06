package types

// Lexicon defines approved and forbidden terminology for API design.
type Lexicon struct {
	// Approved lists terms that should be used.
	Approved []string `json:"approved,omitempty"`

	// Forbidden lists terms that should not be used, with replacements.
	Forbidden []ForbiddenTerm `json:"forbidden,omitempty"`

	// Aliases maps equivalent terms.
	Aliases map[string]string `json:"aliases,omitempty"`

	// CasingRules defines naming conventions for different contexts.
	CasingRules *CasingRules `json:"casingRules,omitempty"`
}

// ForbiddenTerm defines a term that should not be used.
type ForbiddenTerm struct {
	// Term is the forbidden word or phrase.
	Term string `json:"term"`
	// ReplaceWith suggests the preferred alternative.
	ReplaceWith string `json:"replaceWith,omitempty"`
	// Reason explains why this term is forbidden.
	Reason string `json:"reason,omitempty"`
}

// CasingRules defines naming conventions for different API elements.
type CasingRules struct {
	// Paths defines casing for URL paths (e.g., "kebab-case").
	Paths string `json:"paths,omitempty"`
	// Parameters defines casing for query/path parameters.
	Parameters string `json:"parameters,omitempty"`
	// Properties defines casing for JSON properties.
	Properties string `json:"properties,omitempty"`
	// Headers defines casing for HTTP headers.
	Headers string `json:"headers,omitempty"`
}

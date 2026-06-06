package types

// APIStyleSpec is the root type for an API style specification.
// This is the source of truth from which JSON Schema is generated.
type APIStyleSpec struct {
	// Schema is the JSON Schema URI for validation.
	Schema string `json:"$schema,omitempty"`

	// Version is the semantic version of this specification.
	Version string `json:"version"`

	// Name is a unique identifier for this style spec.
	Name string `json:"name"`

	// Description provides context about this style specification.
	Description string `json:"description,omitempty"`

	// Extends lists parent profiles to inherit rules from.
	Extends []string `json:"extends,omitempty"`

	// Rules are the style rules defined in this specification.
	Rules []Rule `json:"rules"`

	// Overrides modify inherited rules from extended profiles.
	Overrides map[string]RuleOverride `json:"overrides,omitempty"`

	// Lexicon defines approved and forbidden terminology.
	Lexicon *Lexicon `json:"lexicon,omitempty"`

	// ConformanceLevels define graduated compliance tiers.
	ConformanceLevels map[string]ConformanceLevel `json:"conformanceLevels,omitempty"`

	// Exceptions are approved rule waivers.
	Exceptions []Exception `json:"exceptions,omitempty"`

	// Categories defines available rule categories with metadata.
	Categories []Category `json:"categories,omitempty"`

	// Metadata contains additional specification information.
	Metadata *SpecMetadata `json:"metadata,omitempty"`
}

// RuleOverride modifies an inherited rule.
type RuleOverride struct {
	// Severity overrides the rule's severity.
	Severity *Severity `json:"severity,omitempty"`
	// Disabled completely disables the rule.
	Disabled bool `json:"disabled,omitempty"`
	// Rationale provides context for the override.
	Rationale string `json:"rationale,omitempty"`
}

// Category defines a grouping for related rules.
type Category struct {
	// ID is the category identifier (e.g., "uri-design").
	ID string `json:"id"`
	// Title is the display name.
	Title string `json:"title"`
	// Description explains what this category covers.
	Description string `json:"description,omitempty"`
	// Order determines display ordering (lower = first).
	Order int `json:"order,omitempty"`
}

// SpecMetadata contains additional specification information.
type SpecMetadata struct {
	// Author is the creator of this specification.
	Author string `json:"author,omitempty"`
	// License is the license for this specification.
	License string `json:"license,omitempty"`
	// Repository is the source repository URL.
	Repository string `json:"repository,omitempty"`
	// Website is a documentation website URL.
	Website string `json:"website,omitempty"`
	// URL is an alias for website/repository for external reference.
	URL string `json:"url,omitempty"`
	// Contact is contact information.
	Contact string `json:"contact,omitempty"`
	// LastUpdated is when the specification was last modified.
	LastUpdated string `json:"lastUpdated,omitempty"`
}

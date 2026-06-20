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

	// Introduction provides detailed introductory content (Markdown).
	Introduction string `json:"introduction,omitempty"`

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

	// Patterns defines reusable API design patterns.
	Patterns []Pattern `json:"patterns,omitempty"`

	// Sections defines document structure for navigation.
	Sections []Section `json:"sections,omitempty"`

	// Glossary defines terminology used in the specification.
	Glossary []GlossaryTerm `json:"glossary,omitempty"`

	// Principles defines high-level design principles.
	Principles []Principle `json:"principles,omitempty"`

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

// RulesForCategory returns all rules with the given category ID.
func (s *APIStyleSpec) RulesForCategory(categoryID string) []Rule {
	var rules []Rule
	for _, r := range s.Rules {
		if r.Category == categoryID {
			rules = append(rules, r)
		}
	}
	return rules
}

// GetCategory returns a category by ID, or nil if not found.
func (s *APIStyleSpec) GetCategory(categoryID string) *Category {
	for i := range s.Categories {
		if s.Categories[i].ID == categoryID {
			return &s.Categories[i]
		}
	}
	return nil
}

// GetPattern returns a pattern by ID, or nil if not found.
func (s *APIStyleSpec) GetPattern(patternID string) *Pattern {
	for i := range s.Patterns {
		if s.Patterns[i].ID == patternID {
			return &s.Patterns[i]
		}
	}
	return nil
}

// GetSection returns a section by ID, or nil if not found.
func (s *APIStyleSpec) GetSection(sectionID string) *Section {
	for i := range s.Sections {
		if s.Sections[i].ID == sectionID {
			return &s.Sections[i]
		}
	}
	return nil
}

// GetRule returns a rule by ID, or nil if not found.
func (s *APIStyleSpec) GetRule(ruleID string) *Rule {
	for i := range s.Rules {
		if s.Rules[i].ID == ruleID {
			return &s.Rules[i]
		}
	}
	return nil
}

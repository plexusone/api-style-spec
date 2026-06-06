package types

// ConformanceLevel defines a graduated compliance tier.
type ConformanceLevel struct {
	// Description explains what this level represents.
	Description string `json:"description,omitempty"`

	// RequiredCategories lists category IDs that must pass.
	RequiredCategories []string `json:"requiredCategories,omitempty"`

	// RequiredRules lists specific rule IDs that must pass.
	RequiredRules []string `json:"requiredRules,omitempty"`

	// MaxErrors is the maximum allowed error-severity violations.
	MaxErrors int `json:"maxErrors"`

	// MaxWarnings is the maximum allowed warning-severity violations.
	MaxWarnings int `json:"maxWarnings"`

	// Extends inherits requirements from another level.
	Extends string `json:"extends,omitempty"`
}

// DefaultConformanceLevels returns the standard bronze/silver/gold levels.
func DefaultConformanceLevels() map[string]ConformanceLevel {
	return map[string]ConformanceLevel{
		"bronze": {
			Description:        "Minimum viable API",
			RequiredCategories: []string{"uri-design", "http-methods"},
			MaxErrors:          0,
			MaxWarnings:        10,
		},
		"silver": {
			Description:        "Production-ready API",
			RequiredCategories: []string{"naming", "errors", "pagination"},
			MaxErrors:          0,
			MaxWarnings:        5,
			Extends:            "bronze",
		},
		"gold": {
			Description:        "Exemplary API",
			RequiredCategories: []string{"security", "documentation", "versioning"},
			MaxErrors:          0,
			MaxWarnings:        0,
			Extends:            "silver",
		},
	}
}

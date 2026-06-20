package types

// MigrationGuidance provides instructions for fixing rule violations.
type MigrationGuidance struct {
	// Summary is a brief description of the migration.
	Summary string `json:"summary"`

	// Steps are ordered migration steps.
	Steps []MigrationStep `json:"steps,omitempty"`

	// AutoFixAvailable indicates if an automated fix is available.
	AutoFixAvailable bool `json:"autoFixAvailable,omitempty"`

	// Effort estimates the work required.
	Effort string `json:"effort,omitempty"` // "low", "medium", "high"

	// BreakingChange indicates if this migration is a breaking change.
	BreakingChange bool `json:"breakingChange,omitempty"`
}

// MigrationStep is a single step in a migration process.
type MigrationStep struct {
	// Order is the step sequence number.
	Order int `json:"order"`

	// Description explains what to do in this step.
	Description string `json:"description"`

	// Code provides example code for the step.
	Code string `json:"code,omitempty"`

	// Language specifies the code language.
	Language string `json:"language,omitempty"`
}

// DeprecationInfo tracks when and why a rule was deprecated.
type DeprecationInfo struct {
	// Version is when the rule was deprecated.
	Version string `json:"version"`

	// Message explains why the rule is deprecated.
	Message string `json:"message"`

	// ReplacedBy lists rule IDs that replace this rule.
	ReplacedBy []string `json:"replacedBy,omitempty"`

	// RemovalVersion indicates when the rule will be removed.
	RemovalVersion string `json:"removalVersion,omitempty"`
}

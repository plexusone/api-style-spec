package types

// FixSuggestion represents a proposed fix for a violation.
type FixSuggestion struct {
	// RuleID is the rule this fix addresses.
	RuleID string `json:"ruleId"`

	// Path is the JSONPath to the element to fix.
	Path string `json:"path"`

	// CurrentValue is the existing value (may be empty for missing fields).
	CurrentValue string `json:"currentValue,omitempty"`

	// SuggestedValue is the proposed replacement.
	SuggestedValue string `json:"suggestedValue"`

	// Diff shows the change in unified diff format.
	Diff string `json:"diff,omitempty"`

	// Confidence indicates certainty that this fix is correct (0.0-1.0).
	Confidence float64 `json:"confidence"`

	// Reasoning explains why this fix is suggested.
	Reasoning string `json:"reasoning,omitempty"`

	// Breaking indicates if this fix could break existing clients.
	Breaking bool `json:"breaking,omitempty"`

	// BreakingReason explains why the fix is breaking.
	BreakingReason string `json:"breakingReason,omitempty"`
}

// FixReport contains all fix suggestions for a spec.
type FixReport struct {
	// Suggestions are the proposed fixes.
	Suggestions []FixSuggestion `json:"suggestions"`

	// PatchOperations are JSON Patch operations (RFC 6902).
	PatchOperations []PatchOperation `json:"patchOperations,omitempty"`

	// FixedCount is how many violations have suggestions.
	FixedCount int `json:"fixedCount"`

	// UnfixedCount is how many violations could not be auto-fixed.
	UnfixedCount int `json:"unfixedCount"`

	// UnfixedRules lists rules that couldn't be auto-fixed.
	UnfixedRules []string `json:"unfixedRules,omitempty"`
}

// PatchOperation represents a JSON Patch operation (RFC 6902).
type PatchOperation struct {
	// Op is the operation: "add", "remove", "replace", "move", "copy", "test".
	Op string `json:"op"`

	// Path is the JSON Pointer to the target location.
	Path string `json:"path"`

	// From is the source location for move/copy operations.
	From string `json:"from,omitempty"`

	// Value is the value for add/replace/test operations.
	Value any `json:"value,omitempty"`
}

// ConformancePath shows the path to reach a conformance level.
type ConformancePath struct {
	// CurrentLevel is the current conformance level (or "none").
	CurrentLevel string `json:"currentLevel"`

	// TargetLevel is the requested conformance level.
	TargetLevel string `json:"targetLevel"`

	// Blockers are errors that must be fixed to reach target.
	Blockers []ConformanceBlocker `json:"blockers"`

	// Warnings are issues that should be addressed.
	Warnings []ConformanceBlocker `json:"warnings,omitempty"`

	// ProgressToTarget is a percentage (0.0-1.0) of completion.
	ProgressToTarget float64 `json:"progressToTarget"`

	// EstimatedFixes is the count of changes needed.
	EstimatedFixes int `json:"estimatedFixes"`
}

// ConformanceBlocker describes a barrier to conformance.
type ConformanceBlocker struct {
	// RuleID is the blocking rule.
	RuleID string `json:"ruleId"`

	// Count is how many violations of this rule exist.
	Count int `json:"count"`

	// Priority is the fix order (1 = fix first).
	Priority int `json:"priority"`

	// FixInstructions provides guidance for resolution.
	FixInstructions string `json:"fixInstructions"`
}

// DesignCheck provides pre-generation guidance.
type DesignCheck struct {
	// Checklist is ordered rules to follow.
	Checklist []DesignCheckItem `json:"checklist"`

	// Template is an OpenAPI skeleton to use.
	Template map[string]any `json:"template,omitempty"`

	// Warnings are potential issues to consider.
	Warnings []string `json:"warnings,omitempty"`
}

// DesignCheckItem is a single design guidance item.
type DesignCheckItem struct {
	// RuleID is the relevant rule.
	RuleID string `json:"ruleId"`

	// Instruction is what to do.
	Instruction string `json:"instruction"`

	// Priority determines order (100 = first).
	Priority int `json:"priority"`

	// Required indicates if this is mandatory.
	Required bool `json:"required"`
}

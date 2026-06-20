package types

// RuleApplicability defines when a rule applies.
type RuleApplicability struct {
	// APITypes specifies which API types this rule applies to.
	APITypes []string `json:"apiTypes,omitempty"` // "rest", "graphql", "grpc"

	// OpenAPIVersions specifies which OpenAPI versions this rule applies to.
	OpenAPIVersions []string `json:"openAPIVersions,omitempty"` // "3.0", "3.1"

	// HTTPMethods specifies which HTTP methods this rule applies to.
	HTTPMethods []string `json:"httpMethods,omitempty"` // "GET", "POST", "PUT", "DELETE", etc.

	// Contexts specifies which contexts this rule applies to.
	Contexts []string `json:"contexts,omitempty"` // "public", "internal", "partner"

	// IncludePatterns are JSONPath expressions for targeted application.
	IncludePatterns []string `json:"includePatterns,omitempty"`

	// ExcludePatterns are JSONPath expressions for exclusion.
	ExcludePatterns []string `json:"excludePatterns,omitempty"`
}

// Condition defines if/then/unless logic for rule application.
type Condition struct {
	// When is a natural language description of when this condition applies.
	When string `json:"when"`

	// Expression is a JSONPath or CEL expression for evaluation.
	Expression string `json:"expression,omitempty"`

	// Then describes what should happen when the condition is met.
	Then string `json:"then"`

	// Unless describes exceptions to the condition.
	Unless string `json:"unless,omitempty"`
}

// RuleRelation defines relationships between rules.
type RuleRelation struct {
	// RuleID is the related rule identifier.
	RuleID string `json:"ruleId"`

	// Type specifies the relationship type.
	Type string `json:"type"` // "requires", "conflicts", "supersedes", "related"

	// Description explains the relationship.
	Description string `json:"description,omitempty"`
}

// DecisionTable provides structured decision guidance.
type DecisionTable struct {
	// Title is the table title.
	Title string `json:"title"`

	// Description explains the table's purpose.
	Description string `json:"description,omitempty"`

	// Headers are the column headers.
	Headers []string `json:"headers"`

	// Rows are the table data rows.
	Rows []DecisionRow `json:"rows"`
}

// DecisionRow is a single row in a decision table.
type DecisionRow struct {
	// Values are the cell values in order.
	Values []string `json:"values"`

	// Highlight indicates if this row should be emphasized.
	Highlight bool `json:"highlight,omitempty"`
}

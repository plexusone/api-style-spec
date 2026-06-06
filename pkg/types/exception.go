package types

import "time"

// Exception defines an approved waiver for a specific rule violation.
type Exception struct {
	// ID is a unique identifier for this exception.
	ID string `json:"id"`

	// RuleID is the rule being waived.
	RuleID string `json:"ruleId"`

	// AppliesTo defines the scope of the exception.
	AppliesTo *ExceptionScope `json:"appliesTo,omitempty"`

	// Reason explains why this exception was granted.
	Reason string `json:"reason"`

	// ApprovedBy identifies who approved the exception.
	ApprovedBy string `json:"approvedBy,omitempty"`

	// ApprovedOn is when the exception was granted.
	ApprovedOn *time.Time `json:"approvedOn,omitempty"`

	// ExpiresOn is when the exception expires (nil = never).
	ExpiresOn *time.Time `json:"expiresOn,omitempty"`

	// Ticket links to an issue tracker for tracking.
	Ticket string `json:"ticket,omitempty"`
}

// ExceptionScope defines where an exception applies.
type ExceptionScope struct {
	// API limits the exception to a specific API name.
	API string `json:"api,omitempty"`
	// Path limits the exception to specific paths (glob supported).
	Path string `json:"path,omitempty"`
	// Operation limits the exception to specific operations.
	Operation string `json:"operation,omitempty"`
	// Paths lists multiple paths (alternative to single Path).
	Paths []string `json:"paths,omitempty"`
}

// IsExpired returns true if the exception has expired.
func (e *Exception) IsExpired() bool {
	if e.ExpiresOn == nil {
		return false
	}
	return time.Now().After(*e.ExpiresOn)
}

// Matches returns true if this exception applies to the given context.
func (e *Exception) Matches(ruleID, api, path, operation string) bool {
	if e.RuleID != ruleID {
		return false
	}
	if e.IsExpired() {
		return false
	}
	if e.AppliesTo == nil {
		return true // No scope means applies everywhere
	}

	scope := e.AppliesTo
	if scope.API != "" && scope.API != api {
		return false
	}
	if scope.Path != "" && scope.Path != path {
		// TODO: Support glob matching
		return false
	}
	if scope.Operation != "" && scope.Operation != operation {
		return false
	}

	return true
}

package types

// Severity represents the severity level of a rule violation.
type Severity string

const (
	// SeverityError indicates a critical violation that should block approval.
	SeverityError Severity = "error"
	// SeverityWarn indicates a significant issue that should be addressed.
	SeverityWarn Severity = "warn"
	// SeverityInfo indicates an informational finding.
	SeverityInfo Severity = "info"
	// SeverityHint indicates a suggestion for improvement.
	SeverityHint Severity = "hint"
)

// IsBlocking returns true if this severity level blocks approval.
func (s Severity) IsBlocking() bool {
	return s == SeverityError
}

// Weight returns the numeric weight for aggregation purposes.
func (s Severity) Weight() int {
	switch s {
	case SeverityError:
		return 100
	case SeverityWarn:
		return 10
	case SeverityInfo:
		return 1
	case SeverityHint:
		return 0
	default:
		return 0
	}
}

// String returns the string representation.
func (s Severity) String() string {
	return string(s)
}

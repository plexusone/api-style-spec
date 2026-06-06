package lint

import (
	"time"

	vacuumModel "github.com/daveshanley/vacuum/model"

	"github.com/plexusone/api-style-spec/pkg/types"
)

// convertVacuumResults converts vacuum results to our LintReport format.
func convertVacuumResults(
	results []vacuumModel.RuleFunctionResult,
	spec *types.APIStyleSpec,
	opts *Options,
	duration time.Duration,
) *types.LintReport {
	report := types.NewLintReport()

	// Build rule lookup for enrichment
	ruleLookup := make(map[string]*types.Rule)
	for i := range spec.Rules {
		ruleLookup[spec.Rules[i].ID] = &spec.Rules[i]
	}

	// Convert each result to a violation
	for i := range results {
		violation := convertVacuumResult(&results[i], ruleLookup)
		report.AddViolation(violation)
	}

	// Set metadata
	report.Metadata = &types.ReportMetadata{
		SpecFile:       opts.FileName,
		Profile:        opts.Profile,
		Duration:       duration,
		DurationMS:     duration.Milliseconds(),
		Timestamp:      time.Now(),
		RulesEvaluated: len(spec.Rules),
	}

	return report
}

// convertVacuumResult converts a single vacuum result to a Violation.
func convertVacuumResult(
	result *vacuumModel.RuleFunctionResult,
	ruleLookup map[string]*types.Rule,
) types.Violation {
	violation := types.Violation{
		RuleID:   result.RuleId,
		Message:  result.Message,
		Path:     result.Path,
		Severity: convertSeverity(result.RuleSeverity),
	}

	// Set location if available
	if result.StartNode != nil {
		violation.Line = result.StartNode.Line
		violation.Column = result.StartNode.Column
	}
	if result.EndNode != nil {
		violation.EndLine = result.EndNode.Line
		violation.EndColumn = result.EndNode.Column
	}

	// Enrich with rule information
	if rule, ok := ruleLookup[result.RuleId]; ok {
		violation.RuleTitle = rule.Title
		violation.Category = rule.Category
		if rule.Examples != nil && len(rule.Examples.Good) > 0 {
			violation.Suggestion = "Consider: " + rule.Examples.Good[0]
		}
	}

	return violation
}

// convertSeverity maps vacuum severity strings to our Severity type.
func convertSeverity(vacuumSeverity string) types.Severity {
	switch vacuumSeverity {
	case "error":
		return types.SeverityError
	case "warn", "warning":
		return types.SeverityWarn
	case "info", "information":
		return types.SeverityInfo
	case "hint":
		return types.SeverityHint
	default:
		return types.SeverityWarn
	}
}

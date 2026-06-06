// Package analyze provides combined API style analysis.
//
// This package orchestrates both deterministic linting (via vacuum)
// and LLM-based evaluation (via judge) to produce comprehensive
// API quality assessments.
//
// Key concepts:
//
//   - Analyzer: Orchestrates lint and evaluate operations
//   - AnalysisReport: Combined results from all analysis types
//   - Decision: GO/NO-GO recommendation based on configured thresholds
//
// Usage:
//
//	spec, _ := profile.Load("azure")
//	analyzer := analyze.New(spec, provider)
//	report, err := analyzer.Analyze(ctx, openAPISpec, opts)
//	if report.Decision == analyze.DecisionNoGo {
//	    // Handle blocking issues
//	}
//
// The analyzer can run lint-only, evaluate-only, or combined analysis
// based on the options provided.
package analyze

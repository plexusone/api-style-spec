// Package judge provides LLM-as-Judge evaluation for API style rules.
//
// This package integrates with structured-evaluation to enable
// AI-powered assessment of API specifications beyond what deterministic
// linting can achieve.
//
// Key concepts:
//
//   - Evaluator: The main interface for running LLM evaluations
//   - RubricSet: A collection of evaluation criteria built from APIStyleSpec rules
//   - EvaluationResult: The outcome of evaluating an API spec against criteria
//
// Usage:
//
//	spec, _ := profile.Load("azure")
//	evaluator := judge.NewEvaluator(provider, spec)
//	result, err := evaluator.Evaluate(ctx, openAPISpec)
//
// The evaluator builds prompts from rule JudgeCriteria and uses an LLM
// provider to assess compliance. Results include scores, reasoning, and
// specific findings mapped back to rule IDs.
package judge

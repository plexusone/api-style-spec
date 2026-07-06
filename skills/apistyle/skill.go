// Package apistyle provides an omniskill Skill for API style specification linting and evaluation.
//
// This package exposes the core api-style-spec functionality as MCP tools:
//   - lint: Lint an OpenAPI spec against style rules
//   - evaluate: LLM-based semantic evaluation
//   - analyze: Combined lint + evaluate with GO/NO-GO decision
//   - list_rules: List all rules from a profile
//   - list_profiles: List available style profiles
//   - explain_rule: Get detailed explanation of a rule
package apistyle

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/plexusone/api-style-spec/pkg/analyze"
	"github.com/plexusone/api-style-spec/pkg/judge"
	"github.com/plexusone/api-style-spec/pkg/lint"
	"github.com/plexusone/api-style-spec/pkg/profile"
	"github.com/plexusone/api-style-spec/pkg/types"
	"github.com/plexusone/omniskill/skill"
)

// Skill provides API style specification tools.
type Skill struct {
	// anthropicAPIKey is used for LLM evaluation (optional).
	anthropicAPIKey string
}

// Option configures the Skill.
type Option func(*Skill)

// WithAnthropicAPIKey sets the Anthropic API key for LLM evaluation.
func WithAnthropicAPIKey(key string) Option {
	return func(s *Skill) {
		s.anthropicAPIKey = key
	}
}

// New creates a new API style skill.
func New(opts ...Option) *Skill {
	s := &Skill{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Name returns the skill identifier.
func (s *Skill) Name() string {
	return "apistyle"
}

// Description returns what this skill does.
func (s *Skill) Description() string {
	return "Lint and evaluate OpenAPI specifications against style guidelines using deterministic rules and LLM-based semantic analysis"
}

// Init initializes the skill (no-op for this skill).
func (s *Skill) Init(_ context.Context) error {
	return nil
}

// Close releases resources (no-op for this skill).
func (s *Skill) Close() error {
	return nil
}

// Tools returns all tools provided by this skill.
func (s *Skill) Tools() []skill.Tool {
	return []skill.Tool{
		s.lintTool(),
		s.evaluateTool(),
		s.analyzeTool(),
		s.listRulesTool(),
		s.listProfilesTool(),
		s.explainRuleTool(),
	}
}

// Ensure Skill implements skill.Skill.
var _ skill.Skill = (*Skill)(nil)

func (s *Skill) lintTool() skill.Tool {
	return skill.NewTool(
		"lint",
		"Lint an OpenAPI specification against API style rules using deterministic checks. Returns violations with severity, rule ID, and location.",
		map[string]skill.Parameter{
			"openapi_spec": {
				Type:        "string",
				Description: "The OpenAPI specification content (YAML or JSON)",
				Required:    true,
			},
			"profile": {
				Type:        "string",
				Description: "Style profile to use (default, minimal, comprehensive, azure, google, microsoft-rest, microsoft-graph, zalando)",
				Required:    false,
				Default:     "default",
			},
		},
		func(ctx context.Context, params map[string]any) (any, error) {
			specContent, _ := params["openapi_spec"].(string)
			if specContent == "" {
				return nil, fmt.Errorf("openapi_spec is required")
			}

			profileName, _ := params["profile"].(string)
			if profileName == "" {
				profileName = "default"
			}

			// Load the style profile
			styleSpec, err := profile.Load(profileName)
			if err != nil {
				return nil, fmt.Errorf("failed to load profile %q: %w", profileName, err)
			}

			// Create linter and run
			linter := lint.NewVacuumLinter(styleSpec)

			report, err := linter.Lint(ctx, []byte(specContent), nil)
			if err != nil {
				return nil, fmt.Errorf("linting failed: %w", err)
			}

			return formatLintReport(report, profileName), nil
		},
	)
}

func (s *Skill) evaluateTool() skill.Tool {
	return skill.NewTool(
		"evaluate",
		"Evaluate an OpenAPI specification using LLM-based semantic analysis. Requires ANTHROPIC_API_KEY. Returns findings with confidence scores.",
		map[string]skill.Parameter{
			"openapi_spec": {
				Type:        "string",
				Description: "The OpenAPI specification content (YAML or JSON)",
				Required:    true,
			},
			"profile": {
				Type:        "string",
				Description: "Style profile to use (default, minimal, comprehensive, azure, google, microsoft-rest, microsoft-graph, zalando)",
				Required:    false,
				Default:     "default",
			},
			"categories": {
				Type:        "array",
				Description: "Categories to evaluate (empty = all)",
				Required:    false,
				Items:       &skill.Parameter{Type: "string"},
			},
		},
		func(ctx context.Context, params map[string]any) (any, error) {
			if s.anthropicAPIKey == "" {
				return nil, fmt.Errorf("ANTHROPIC_API_KEY not configured; LLM evaluation unavailable")
			}

			specContent, _ := params["openapi_spec"].(string)
			if specContent == "" {
				return nil, fmt.Errorf("openapi_spec is required")
			}

			profileName, _ := params["profile"].(string)
			if profileName == "" {
				profileName = "default"
			}

			// Load the style profile
			styleSpec, err := profile.Load(profileName)
			if err != nil {
				return nil, fmt.Errorf("failed to load profile %q: %w", profileName, err)
			}

			// Create evaluator
			provider := judge.NewAnthropicProvider(s.anthropicAPIKey, nil)
			evaluator := judge.NewClaudeEvaluator(provider, styleSpec)

			opts := &judge.Options{}

			// Handle categories
			if cats, ok := params["categories"].([]any); ok && len(cats) > 0 {
				categories := make([]string, len(cats))
				for i, c := range cats {
					categories[i], _ = c.(string)
				}
				opts.Categories = categories
			}

			report, err := evaluator.Evaluate(ctx, []byte(specContent), opts)
			if err != nil {
				return nil, fmt.Errorf("evaluation failed: %w", err)
			}

			return formatEvaluationReport(report), nil
		},
	)
}

func (s *Skill) analyzeTool() skill.Tool {
	return skill.NewTool(
		"analyze",
		"Combined analysis: deterministic linting + LLM evaluation with GO/NO-GO decision. Best for comprehensive API review.",
		map[string]skill.Parameter{
			"openapi_spec": {
				Type:        "string",
				Description: "The OpenAPI specification content (YAML or JSON)",
				Required:    true,
			},
			"profile": {
				Type:        "string",
				Description: "Style profile to use (default, minimal, comprehensive, azure, google, microsoft-rest, microsoft-graph, zalando)",
				Required:    false,
				Default:     "default",
			},
			"lint_only": {
				Type:        "boolean",
				Description: "Skip LLM evaluation, only run deterministic linting",
				Required:    false,
				Default:     false,
			},
		},
		func(ctx context.Context, params map[string]any) (any, error) {
			specContent, _ := params["openapi_spec"].(string)
			if specContent == "" {
				return nil, fmt.Errorf("openapi_spec is required")
			}

			profileName, _ := params["profile"].(string)
			if profileName == "" {
				profileName = "default"
			}

			lintOnly, _ := params["lint_only"].(bool)

			// Load the style profile
			styleSpec, err := profile.Load(profileName)
			if err != nil {
				return nil, fmt.Errorf("failed to load profile %q: %w", profileName, err)
			}

			// Create analyzer
			var provider judge.Provider
			if !lintOnly && s.anthropicAPIKey != "" {
				provider = judge.NewAnthropicProvider(s.anthropicAPIKey, nil)
			}

			analyzer := analyze.New(styleSpec, provider)

			opts := &analyze.Options{
				EnableLint:     true,
				EnableEvaluate: !lintOnly && s.anthropicAPIKey != "",
			}

			report, err := analyzer.Analyze(ctx, []byte(specContent), opts)
			if err != nil {
				return nil, fmt.Errorf("analysis failed: %w", err)
			}

			return formatAnalysisReport(report), nil
		},
	)
}

func (s *Skill) listRulesTool() skill.Tool {
	return skill.NewTool(
		"list_rules",
		"List all rules from a style profile with their IDs, titles, categories, and severities.",
		map[string]skill.Parameter{
			"profile": {
				Type:        "string",
				Description: "Style profile to use (default, minimal, comprehensive, azure, google, microsoft-rest, microsoft-graph, zalando)",
				Required:    false,
				Default:     "default",
			},
			"category": {
				Type:        "string",
				Description: "Filter rules by category",
				Required:    false,
			},
			"severity": {
				Type:        "string",
				Description: "Filter rules by severity (error, warn, info, hint)",
				Required:    false,
			},
		},
		func(_ context.Context, params map[string]any) (any, error) {
			profileName, _ := params["profile"].(string)
			if profileName == "" {
				profileName = "default"
			}

			categoryFilter, _ := params["category"].(string)
			severityFilter, _ := params["severity"].(string)

			// Load the style profile
			styleSpec, err := profile.Load(profileName)
			if err != nil {
				return nil, fmt.Errorf("failed to load profile %q: %w", profileName, err)
			}

			var rules []map[string]any
			for _, rule := range styleSpec.Rules {
				// Apply filters
				if categoryFilter != "" && rule.Category != categoryFilter {
					continue
				}
				if severityFilter != "" && string(rule.Severity) != severityFilter {
					continue
				}

				rules = append(rules, map[string]any{
					"id":       rule.ID,
					"title":    rule.Title,
					"category": rule.Category,
					"severity": rule.Severity,
				})
			}

			return map[string]any{
				"profile":     profileName,
				"rule_count":  len(rules),
				"total_rules": len(styleSpec.Rules),
				"rules":       rules,
			}, nil
		},
	)
}

func (s *Skill) listProfilesTool() skill.Tool {
	return skill.NewTool(
		"list_profiles",
		"List all available style profiles with their descriptions and rule counts.",
		map[string]skill.Parameter{},
		func(_ context.Context, _ map[string]any) (any, error) {
			names, err := profile.ListBuiltin()
			if err != nil {
				return nil, fmt.Errorf("listing profiles: %w", err)
			}

			result := make([]map[string]any, 0, len(names))
			for _, name := range names {
				spec, err := profile.Load(name)
				if err != nil {
					continue
				}

				result = append(result, map[string]any{
					"name":        name,
					"description": spec.Description,
					"version":     spec.Version,
					"rule_count":  len(spec.Rules),
				})
			}

			return map[string]any{
				"profiles": result,
			}, nil
		},
	)
}

func (s *Skill) explainRuleTool() skill.Tool {
	return skill.NewTool(
		"explain_rule",
		"Get detailed explanation of a specific rule including rationale, examples, and references.",
		map[string]skill.Parameter{
			"rule_id": {
				Type:        "string",
				Description: "The rule ID to explain (e.g., PO-001)",
				Required:    true,
			},
			"profile": {
				Type:        "string",
				Description: "Style profile to use (default, minimal, comprehensive, azure, google, microsoft-rest, microsoft-graph, zalando)",
				Required:    false,
				Default:     "default",
			},
		},
		func(_ context.Context, params map[string]any) (any, error) {
			ruleID, _ := params["rule_id"].(string)
			if ruleID == "" {
				return nil, fmt.Errorf("rule_id is required")
			}

			profileName, _ := params["profile"].(string)
			if profileName == "" {
				profileName = "default"
			}

			// Load the style profile
			styleSpec, err := profile.Load(profileName)
			if err != nil {
				return nil, fmt.Errorf("failed to load profile %q: %w", profileName, err)
			}

			// Find the rule
			var rule *types.Rule
			for i := range styleSpec.Rules {
				if strings.EqualFold(styleSpec.Rules[i].ID, ruleID) {
					rule = &styleSpec.Rules[i]
					break
				}
			}

			if rule == nil {
				return nil, fmt.Errorf("rule %q not found in profile %q", ruleID, profileName)
			}

			result := map[string]any{
				"id":        rule.ID,
				"title":     rule.Title,
				"category":  rule.Category,
				"severity":  rule.Severity,
				"rationale": rule.Rationale,
			}

			if rule.Examples != nil {
				result["examples"] = map[string]any{
					"good": rule.Examples.Good,
					"bad":  rule.Examples.Bad,
				}
			}

			if len(rule.References) > 0 {
				refs := make([]map[string]string, len(rule.References))
				for i, ref := range rule.References {
					refs[i] = map[string]string{
						"title": ref.Title,
						"url":   ref.URL,
					}
				}
				result["references"] = refs
			}

			if rule.Enforcement != nil {
				result["enforcement"] = map[string]any{
					"type":     rule.Enforcement.Type,
					"function": rule.Enforcement.Function,
				}
			}

			if rule.Judge != nil {
				result["llm_judgeable"] = true
			}

			return result, nil
		},
	)
}

// formatLintReport converts a lint report to a map for JSON serialization.
func formatLintReport(report *types.LintReport, profileName string) map[string]any {
	violations := make([]map[string]any, len(report.Violations))
	for i, v := range report.Violations {
		violations[i] = map[string]any{
			"rule_id":  v.RuleID,
			"severity": v.Severity,
			"message":  v.Message,
			"path":     v.Path,
			"line":     v.Line,
		}
	}

	return map[string]any{
		"status":          report.Status,
		"violation_count": len(report.Violations),
		"violations":      violations,
		"profile":         profileName,
		"summary": map[string]any{
			"errors":   report.Summary.Errors,
			"warnings": report.Summary.Warnings,
			"infos":    report.Summary.Infos,
			"hints":    report.Summary.Hints,
			"total":    report.Summary.Total,
		},
	}
}

// formatEvaluationReport converts an evaluation report to a map for JSON serialization.
func formatEvaluationReport(report *judge.EvaluationReport) map[string]any {
	findings := make([]map[string]any, len(report.Findings))
	for i, f := range report.Findings {
		findings[i] = map[string]any{
			"rule_id":     f.RuleID,
			"passed":      f.Passed,
			"score":       f.Score,
			"reasoning":   f.Reasoning,
			"suggestions": f.Suggestions,
		}
	}

	categories := make([]map[string]any, len(report.Categories))
	for i, c := range report.Categories {
		categories[i] = map[string]any{
			"name":  c.Name,
			"score": c.Score,
		}
	}

	return map[string]any{
		"status":        report.Status,
		"summary":       report.Summary,
		"categories":    categories,
		"findings":      findings,
		"finding_count": len(report.Findings),
	}
}

// formatAnalysisReport converts an analysis report to a map for JSON serialization.
func formatAnalysisReport(report *analyze.AnalysisReport) map[string]any {
	result := map[string]any{
		"decision": report.Decision,
		"summary":  report.Summary,
	}

	if report.LintReport != nil {
		result["lint"] = formatLintReport(report.LintReport, "")
	}

	if report.EvaluationReport != nil {
		result["evaluation"] = formatEvaluationReport(report.EvaluationReport)
	}

	return result
}

// MarshalJSON implements json.Marshaler for Skill metadata.
func (s *Skill) MarshalJSON() ([]byte, error) {
	tools := s.Tools()
	toolNames := make([]string, len(tools))
	for i, t := range tools {
		toolNames[i] = t.Name()
	}

	return json.Marshal(map[string]any{
		"name":        s.Name(),
		"description": s.Description(),
		"tools":       toolNames,
	})
}

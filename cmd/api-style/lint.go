package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/plexusone/api-style-spec/pkg/lint"
	"github.com/plexusone/api-style-spec/pkg/profile"
	"github.com/plexusone/api-style-spec/pkg/sarif"
	"github.com/plexusone/api-style-spec/pkg/types"
)

var (
	lintFormat  string
	lintOutput  string
	lintProfile string
	lintLevel   string
)

var lintCmd = &cobra.Command{
	Use:   "lint <openapi-spec>",
	Short: "Lint an OpenAPI specification",
	Long: `Lint an OpenAPI specification against style rules.

Uses vacuum (a fast, Spectral-compatible linter) to check the specification
against the configured style rules.

Examples:
  api-style lint openapi.yaml
  api-style lint openapi.yaml --format json
  api-style lint openapi.yaml --format json --output report.json
  api-style lint openapi.yaml --profile azure --level silver`,
	Args: cobra.ExactArgs(1),
	RunE: runLint,
}

func init() {
	lintCmd.Flags().StringVarP(&lintFormat, "format", "f", "text", "Output format: text, json, sarif")
	lintCmd.Flags().StringVarP(&lintOutput, "output", "o", "", "Output file (default: stdout)")
	lintCmd.Flags().StringVarP(&lintProfile, "profile", "p", "default", "Style profile to use")
	lintCmd.Flags().StringVarP(&lintLevel, "level", "l", "", "Target conformance level (bronze, silver, gold)")
}

func runLint(_ *cobra.Command, args []string) error {
	specFile := args[0]

	// Read the OpenAPI spec
	specBytes, err := os.ReadFile(specFile)
	if err != nil {
		return fmt.Errorf("reading spec file: %w", err)
	}

	ctx := context.Background()
	opts := &lint.Options{
		FileName:         specFile,
		Profile:          lintProfile,
		ConformanceLevel: lintLevel,
	}

	var report *types.LintReport
	var styleSpec *types.APIStyleSpec

	// Load profile if specified
	if lintProfile != "" && lintProfile != "vacuum" {
		loadedSpec, loadErr := profile.Load(lintProfile)
		if loadErr != nil {
			// Fall back to vacuum defaults if profile not found
			fmt.Fprintf(os.Stderr, "Warning: profile %q not found, using vacuum defaults\n", lintProfile)
			report, err = lint.WithDefaults(ctx, specBytes, opts)
			if err != nil {
				return fmt.Errorf("linting: %w", err)
			}
		} else {
			styleSpec = loadedSpec
			// Use custom profile
			linter := lint.NewVacuumLinter(styleSpec)
			report, err = linter.Lint(ctx, specBytes, opts)
			if err != nil {
				return fmt.Errorf("linting: %w", err)
			}
		}
	} else {
		// Use vacuum's default rules
		report, err = lint.WithDefaults(ctx, specBytes, opts)
		if err != nil {
			return fmt.Errorf("linting: %w", err)
		}
	}

	// Format output
	output, err := formatReport(report, styleSpec, lintFormat)
	if err != nil {
		return fmt.Errorf("formatting report: %w", err)
	}

	// Write output
	if lintOutput != "" {
		if err := os.WriteFile(lintOutput, []byte(output), 0o600); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
	} else {
		fmt.Print(output)
	}

	// Exit with code 1 if there are blocking violations
	if report.HasBlockingViolations() {
		os.Exit(1)
	}

	return nil
}

func formatReport(report *types.LintReport, styleSpec *types.APIStyleSpec, format string) (string, error) {
	switch strings.ToLower(format) {
	case "json":
		return formatJSON(report)
	case "sarif":
		return formatSARIF(report, styleSpec)
	case "text":
		return formatText(report), nil
	default:
		return "", fmt.Errorf("unknown format: %s", format)
	}
}

func formatJSON(report *types.LintReport) (string, error) {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data) + "\n", nil
}

func formatText(report *types.LintReport) string {
	var sb strings.Builder

	// Header
	sb.WriteString("API Style Lint Report\n")
	sb.WriteString("=====================\n\n")

	// Summary
	fmt.Fprintf(&sb, "Status: %s\n", strings.ToUpper(string(report.Status)))
	fmt.Fprintf(&sb, "Errors: %d, Warnings: %d, Info: %d, Hints: %d\n\n",
		report.Summary.Errors,
		report.Summary.Warnings,
		report.Summary.Infos,
		report.Summary.Hints,
	)

	if len(report.Violations) == 0 {
		sb.WriteString("No violations found.\n")
		return sb.String()
	}

	// Group by severity
	byLevel := map[types.Severity][]types.Violation{
		types.SeverityError: {},
		types.SeverityWarn:  {},
		types.SeverityInfo:  {},
		types.SeverityHint:  {},
	}

	for _, v := range report.Violations {
		byLevel[v.Severity] = append(byLevel[v.Severity], v)
	}

	// Print violations by severity
	printViolations(&sb, "Errors", byLevel[types.SeverityError])
	printViolations(&sb, "Warnings", byLevel[types.SeverityWarn])
	printViolations(&sb, "Info", byLevel[types.SeverityInfo])
	printViolations(&sb, "Hints", byLevel[types.SeverityHint])

	return sb.String()
}

func printViolations(sb *strings.Builder, level string, violations []types.Violation) {
	if len(violations) == 0 {
		return
	}

	fmt.Fprintf(sb, "%s:\n", level)
	for _, v := range violations {
		location := v.Path
		if v.Line > 0 {
			location = fmt.Sprintf("%s (line %d)", v.Path, v.Line)
		}
		fmt.Fprintf(sb, "  - [%s] %s\n    %s\n", v.RuleID, v.Message, location)
	}
	sb.WriteString("\n")
}

// formatSARIF outputs in SARIF format for IDE integration
func formatSARIF(report *types.LintReport, styleSpec *types.APIStyleSpec) (string, error) {
	opts := sarif.DefaultOptions()

	// Add rule metadata if we have a style spec
	if styleSpec != nil {
		opts.Rules = make(map[string]*types.Rule)
		for i := range styleSpec.Rules {
			rule := &styleSpec.Rules[i]
			opts.Rules[rule.ID] = rule
		}
	}

	return sarif.FormatLintReport(report, opts)
}

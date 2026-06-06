package generate

import (
	"strings"
	"testing"

	"github.com/plexusone/api-style-spec/pkg/types"
)

func TestMarkdown_Basic(t *testing.T) {
	spec := &types.APIStyleSpec{
		Name:        "Test Guidelines",
		Version:     "1.0.0",
		Description: "Test API style guidelines.",
		Categories: []types.Category{
			{ID: "uri", Title: "URI Design", Description: "URL conventions"},
		},
		Rules: []types.Rule{
			{
				ID:        "URI-001",
				Title:     "Use plural nouns",
				Category:  "uri",
				Severity:  types.SeverityError,
				Rationale: "Plural nouns are clearer.",
				Examples: &types.Examples{
					Good: []string{"/users", "/orders"},
					Bad:  []string{"/user", "/order"},
				},
			},
		},
	}

	md, err := Markdown(spec, nil)
	if err != nil {
		t.Fatalf("Markdown() error = %v", err)
	}

	// Check title
	if !strings.Contains(md, "# Test Guidelines") {
		t.Error("Missing title")
	}

	// Check description
	if !strings.Contains(md, "Test API style guidelines.") {
		t.Error("Missing description")
	}

	// Check rule
	if !strings.Contains(md, "URI-001") {
		t.Error("Missing rule ID")
	}

	if !strings.Contains(md, "Use plural nouns") {
		t.Error("Missing rule title")
	}

	// Check examples
	if !strings.Contains(md, "/users") {
		t.Error("Missing good example")
	}

	if !strings.Contains(md, "/user") {
		t.Error("Missing bad example")
	}
}

func TestMarkdown_WithConformanceLevels(t *testing.T) {
	spec := &types.APIStyleSpec{
		Name: "Test",
		ConformanceLevels: map[string]types.ConformanceLevel{
			"bronze": {
				Description:   "Basic compliance",
				RequiredRules: []string{"RULE-001"},
			},
			"silver": {
				Description:   "Standard compliance",
				RequiredRules: []string{"RULE-001", "RULE-002"},
				Extends:       "bronze",
			},
		},
		Rules: []types.Rule{
			{ID: "RULE-001", Title: "Rule 1", Category: "test"},
			{ID: "RULE-002", Title: "Rule 2", Category: "test"},
		},
	}

	opts := DefaultMarkdownOptions()
	opts.IncludeConformance = true

	md, err := Markdown(spec, opts)
	if err != nil {
		t.Fatalf("Markdown() error = %v", err)
	}

	if !strings.Contains(md, "Conformance Levels") {
		t.Error("Missing conformance levels section")
	}

	if !strings.Contains(md, "Bronze") {
		t.Error("Missing bronze level")
	}

	if !strings.Contains(md, "Silver") {
		t.Error("Missing silver level")
	}
}

func TestMarkdown_WithMetadata(t *testing.T) {
	spec := &types.APIStyleSpec{
		Name: "Test",
		Metadata: &types.SpecMetadata{
			Author:      "Test Author",
			URL:         "https://example.com",
			LastUpdated: "2024-01-01",
		},
		Rules: []types.Rule{},
	}

	opts := DefaultMarkdownOptions()
	opts.IncludeMetadata = true

	md, err := Markdown(spec, opts)
	if err != nil {
		t.Fatalf("Markdown() error = %v", err)
	}

	if !strings.Contains(md, "Test Author") {
		t.Error("Missing author")
	}

	if !strings.Contains(md, "https://example.com") {
		t.Error("Missing URL")
	}
}

func TestSpectral_Basic(t *testing.T) {
	spec := &types.APIStyleSpec{
		Name:    "Test Guidelines",
		Version: "1.0.0",
		Rules: []types.Rule{
			{
				ID:       "URI-001",
				Title:    "Use lowercase in URIs",
				Category: "uri",
				Severity: types.SeverityError,
				Enforcement: &types.Enforcement{
					Type:     types.EnforcementSpectral,
					Function: "pattern",
					Given:    types.NewGivenPath("$.paths[*]~"),
					Options: &types.EnforcementOptions{
						Match: "^[a-z0-9/{}.-]*$",
					},
				},
			},
		},
	}

	yaml, err := Spectral(spec, nil)
	if err != nil {
		t.Fatalf("Spectral() error = %v", err)
	}

	// Check header
	if !strings.Contains(yaml, "Generated from: Test Guidelines") {
		t.Error("Missing generated header")
	}

	// Check extends
	if !strings.Contains(yaml, "extends:") {
		t.Error("Missing extends section")
	}

	// Check rule
	if !strings.Contains(yaml, "uri-001:") {
		t.Error("Missing rule ID (should be kebab-case)")
	}

	if !strings.Contains(yaml, "severity: error") {
		t.Error("Missing severity")
	}

	if !strings.Contains(yaml, "function: pattern") {
		t.Error("Missing function")
	}

	if !strings.Contains(yaml, "match:") {
		t.Error("Missing functionOptions")
	}
}

func TestSpectral_SkipsNonEnforceable(t *testing.T) {
	spec := &types.APIStyleSpec{
		Name: "Test",
		Rules: []types.Rule{
			{
				ID:       "LLM-001",
				Title:    "LLM-only rule",
				Category: "test",
				Severity: types.SeverityWarn,
				// No enforcement - LLM only
				Judge: &types.JudgeCriteria{
					Prompt: "Check this rule",
				},
			},
		},
	}

	opts := DefaultSpectralOptions()
	opts.SkipNonEnforceable = true

	yaml, err := Spectral(spec, opts)
	if err != nil {
		t.Fatalf("Spectral() error = %v", err)
	}

	// Should not contain the rule
	if strings.Contains(yaml, "llm-001:") {
		t.Error("Should skip non-enforceable rules")
	}
}

func TestSpectral_MultipleGivenPaths(t *testing.T) {
	spec := &types.APIStyleSpec{
		Name: "Test",
		Rules: []types.Rule{
			{
				ID:       "SEC-001",
				Title:    "Define security",
				Category: "security",
				Severity: types.SeverityError,
				Enforcement: &types.Enforcement{
					Type:     types.EnforcementSpectral,
					Function: "truthy",
					Given:    types.NewGivenPaths("$.security", "$.components.securitySchemes"),
				},
			},
		},
	}

	yaml, err := Spectral(spec, nil)
	if err != nil {
		t.Fatalf("Spectral() error = %v", err)
	}

	// Should have multiple given paths
	if !strings.Contains(yaml, "given:") {
		t.Error("Missing given section")
	}

	if !strings.Contains(yaml, "$.security") {
		t.Error("Missing first given path")
	}

	if !strings.Contains(yaml, "$.components.securitySchemes") {
		t.Error("Missing second given path")
	}
}

func TestGroupRulesByCategory(t *testing.T) {
	rules := []types.Rule{
		{ID: "A-001", Category: "cat-a"},
		{ID: "A-002", Category: "cat-a"},
		{ID: "B-001", Category: "cat-b"},
		{ID: "C-001", Category: ""},
	}

	groups := groupRulesByCategory(rules)

	if len(groups["cat-a"]) != 2 {
		t.Errorf("cat-a should have 2 rules, got %d", len(groups["cat-a"]))
	}

	if len(groups["cat-b"]) != 1 {
		t.Errorf("cat-b should have 1 rule, got %d", len(groups["cat-b"]))
	}

	if len(groups["other"]) != 1 {
		t.Errorf("other should have 1 rule, got %d", len(groups["other"]))
	}
}

func TestToKebabCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"URI-001", "uri-001"},
		{"HTTP-METHOD-002", "http-method-002"},
		{"simple", "simple"},
	}

	for _, tt := range tests {
		result := toKebabCase(tt.input)
		if result != tt.expected {
			t.Errorf("toKebabCase(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

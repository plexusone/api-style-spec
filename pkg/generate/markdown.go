package generate

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/plexusone/api-style-spec/pkg/types"
)

var titleCaser = cases.Title(language.English)

// MarkdownOptions configures Markdown generation.
type MarkdownOptions struct {
	// IncludeTOC adds a table of contents.
	IncludeTOC bool

	// IncludeExamples includes good/bad examples.
	IncludeExamples bool

	// IncludeRationale includes rule rationales.
	IncludeRationale bool

	// IncludeReferences includes external references.
	IncludeReferences bool

	// IncludeConformance includes conformance level details.
	IncludeConformance bool

	// IncludeMetadata includes spec metadata.
	IncludeMetadata bool

	// SeverityEmojis uses emojis for severity indicators.
	SeverityEmojis bool
}

// DefaultMarkdownOptions returns options with all features enabled.
func DefaultMarkdownOptions() *MarkdownOptions {
	return &MarkdownOptions{
		IncludeTOC:         true,
		IncludeExamples:    true,
		IncludeRationale:   true,
		IncludeReferences:  true,
		IncludeConformance: true,
		IncludeMetadata:    true,
		SeverityEmojis:     false,
	}
}

// Markdown generates Markdown documentation from an APIStyleSpec.
func Markdown(spec *types.APIStyleSpec, opts *MarkdownOptions) (string, error) {
	if opts == nil {
		opts = DefaultMarkdownOptions()
	}

	var sb strings.Builder

	// Title
	sb.WriteString("# ")
	sb.WriteString(spec.Name)
	sb.WriteString("\n\n")

	// Description
	if spec.Description != "" {
		sb.WriteString(spec.Description)
		sb.WriteString("\n\n")
	}

	// Metadata
	if opts.IncludeMetadata && spec.Metadata != nil {
		writeMetadata(&sb, spec.Metadata)
	}

	// Table of Contents
	if opts.IncludeTOC {
		writeTOC(&sb, spec, opts)
	}

	// Conformance Levels
	if opts.IncludeConformance && len(spec.ConformanceLevels) > 0 {
		writeConformanceLevels(&sb, spec)
	}

	// Rules by Category
	writeRulesByCategory(&sb, spec, opts)

	return sb.String(), nil
}

func writeMetadata(sb *strings.Builder, meta *types.SpecMetadata) {
	sb.WriteString("## Metadata\n\n")

	if meta.Author != "" {
		fmt.Fprintf(sb, "- **Author:** %s\n", meta.Author)
	}
	if meta.URL != "" {
		fmt.Fprintf(sb, "- **Source:** [%s](%s)\n", meta.URL, meta.URL)
	}
	if meta.LastUpdated != "" {
		fmt.Fprintf(sb, "- **Last Updated:** %s\n", meta.LastUpdated)
	}
	if meta.License != "" {
		fmt.Fprintf(sb, "- **License:** %s\n", meta.License)
	}

	sb.WriteString("\n")
}

func writeTOC(sb *strings.Builder, spec *types.APIStyleSpec, opts *MarkdownOptions) {
	sb.WriteString("## Table of Contents\n\n")

	if opts.IncludeConformance && len(spec.ConformanceLevels) > 0 {
		sb.WriteString("- [Conformance Levels](#conformance-levels)\n")
	}

	// Group rules by category
	categories := groupRulesByCategory(spec.Rules)
	categoryOrder := getCategoryOrder(spec.Categories, categories)

	for _, catID := range categoryOrder {
		catName := getCategoryName(spec.Categories, catID)
		anchor := strings.ToLower(strings.ReplaceAll(catName, " ", "-"))
		fmt.Fprintf(sb, "- [%s](#%s)\n", catName, anchor)
	}

	sb.WriteString("\n")
}

func writeConformanceLevels(sb *strings.Builder, spec *types.APIStyleSpec) {
	sb.WriteString("## Conformance Levels\n\n")

	// Sort levels: bronze, silver, gold
	levelOrder := []string{"bronze", "silver", "gold"}

	for _, levelName := range levelOrder {
		level, ok := spec.ConformanceLevels[levelName]
		if !ok {
			continue
		}

		title := titleCaser.String(levelName)
		fmt.Fprintf(sb, "### %s\n\n", title)

		if level.Description != "" {
			sb.WriteString(level.Description)
			sb.WriteString("\n\n")
		}

		if len(level.RequiredRules) > 0 {
			sb.WriteString("**Required Rules:**\n\n")
			for _, ruleID := range level.RequiredRules {
				fmt.Fprintf(sb, "- %s\n", ruleID)
			}
			sb.WriteString("\n")
		}
	}
}

func writeRulesByCategory(sb *strings.Builder, spec *types.APIStyleSpec, opts *MarkdownOptions) {
	categories := groupRulesByCategory(spec.Rules)
	categoryOrder := getCategoryOrder(spec.Categories, categories)

	for _, catID := range categoryOrder {
		rules := categories[catID]
		if len(rules) == 0 {
			continue
		}

		catName := getCategoryName(spec.Categories, catID)
		catDesc := getCategoryDescription(spec.Categories, catID)

		fmt.Fprintf(sb, "## %s\n\n", catName)

		if catDesc != "" {
			sb.WriteString(catDesc)
			sb.WriteString("\n\n")
		}

		for _, rule := range rules {
			writeRule(sb, rule, opts)
		}
	}
}

func writeRule(sb *strings.Builder, rule types.Rule, opts *MarkdownOptions) {
	// Rule header
	fmt.Fprintf(sb, "### %s: %s\n\n", rule.ID, rule.Title)

	// Severity badge
	severityStr := formatSeverity(rule.Severity, opts.SeverityEmojis)
	fmt.Fprintf(sb, "**Severity:** %s\n\n", severityStr)

	// Rationale
	if opts.IncludeRationale && rule.Rationale != "" {
		sb.WriteString(rule.Rationale)
		sb.WriteString("\n\n")
	}

	// Examples
	if opts.IncludeExamples && rule.Examples != nil {
		writeExamples(sb, rule.Examples)
	}

	// References
	if opts.IncludeReferences && len(rule.References) > 0 {
		writeReferences(sb, rule.References)
	}

	sb.WriteString("---\n\n")
}

func writeExamples(sb *strings.Builder, examples *types.Examples) {
	if len(examples.Good) == 0 && len(examples.Bad) == 0 {
		return
	}

	sb.WriteString("**Examples:**\n\n")

	if len(examples.Good) > 0 {
		sb.WriteString("Good:\n")
		for _, ex := range examples.Good {
			fmt.Fprintf(sb, "- `%s`\n", ex)
		}
		sb.WriteString("\n")
	}

	if len(examples.Bad) > 0 {
		sb.WriteString("Bad:\n")
		for _, ex := range examples.Bad {
			fmt.Fprintf(sb, "- `%s`\n", ex)
		}
		sb.WriteString("\n")
	}
}

func writeReferences(sb *strings.Builder, refs []types.Reference) {
	sb.WriteString("**References:**\n\n")
	for _, ref := range refs {
		fmt.Fprintf(sb, "- [%s](%s)\n", ref.Title, ref.URL)
	}
	sb.WriteString("\n")
}

func formatSeverity(sev types.Severity, useEmoji bool) string {
	if useEmoji {
		switch sev {
		case types.SeverityError:
			return "Error"
		case types.SeverityWarn:
			return "Warning"
		case types.SeverityInfo:
			return "Info"
		case types.SeverityHint:
			return "Hint"
		default:
			return string(sev)
		}
	}
	return string(sev)
}

func groupRulesByCategory(rules []types.Rule) map[string][]types.Rule {
	categories := make(map[string][]types.Rule)
	for _, rule := range rules {
		cat := rule.Category
		if cat == "" {
			cat = "other"
		}
		categories[cat] = append(categories[cat], rule)
	}
	return categories
}

func getCategoryOrder(categories []types.Category, ruleCategories map[string][]types.Rule) []string {
	// First add categories in defined order
	var order []string
	seen := make(map[string]bool)

	for _, cat := range categories {
		if _, hasRules := ruleCategories[cat.ID]; hasRules {
			order = append(order, cat.ID)
			seen[cat.ID] = true
		}
	}

	// Then add any categories not in the defined list
	var extra []string
	for catID := range ruleCategories {
		if !seen[catID] {
			extra = append(extra, catID)
		}
	}
	sort.Strings(extra)
	order = append(order, extra...)

	return order
}

func getCategoryName(categories []types.Category, id string) string {
	for _, cat := range categories {
		if cat.ID == id {
			return cat.Title
		}
	}
	// Fallback: title case the ID
	return titleCaser.String(strings.ReplaceAll(id, "-", " "))
}

func getCategoryDescription(categories []types.Category, id string) string {
	for _, cat := range categories {
		if cat.ID == id {
			return cat.Description
		}
	}
	return ""
}

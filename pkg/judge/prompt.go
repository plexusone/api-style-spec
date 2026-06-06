package judge

import (
	"fmt"
	"strings"
)

// PromptBuilder constructs evaluation prompts for LLM evaluation.
type PromptBuilder struct {
	// SystemPrompt is the base system instruction.
	SystemPrompt string
}

// NewPromptBuilder creates a new prompt builder with default system prompt.
func NewPromptBuilder() *PromptBuilder {
	return &PromptBuilder{
		SystemPrompt: defaultSystemPrompt,
	}
}

const defaultSystemPrompt = `You are an expert API designer and reviewer. Your task is to evaluate OpenAPI specifications against specific style rules and guidelines.

For each evaluation:
1. Analyze the OpenAPI specification carefully
2. Check if the specified rule is followed throughout the spec
3. Provide a score from 0.0 to 1.0:
   - 1.0 = Fully compliant, no issues found
   - 0.8-0.9 = Mostly compliant, minor issues
   - 0.5-0.7 = Partially compliant, some issues
   - 0.2-0.4 = Mostly non-compliant, significant issues
   - 0.0-0.1 = Not compliant, major violations
4. Explain your reasoning
5. Provide specific examples from the spec
6. Suggest improvements if applicable

Be thorough but fair. Consider context and practical constraints.`

// BuildSingleEvaluation creates a prompt for evaluating a single criterion.
func (pb *PromptBuilder) BuildSingleEvaluation(criterion *Criterion, specContent string) string {
	var sb strings.Builder

	sb.WriteString("## Rule to Evaluate\n\n")
	fmt.Fprintf(&sb, "**ID:** %s\n", criterion.RuleID)
	fmt.Fprintf(&sb, "**Title:** %s\n", criterion.RuleTitle)
	fmt.Fprintf(&sb, "**Category:** %s\n", criterion.Category)
	fmt.Fprintf(&sb, "**Severity:** %s\n\n", criterion.Severity)

	sb.WriteString("### Evaluation Criteria\n\n")
	sb.WriteString(criterion.Prompt)
	sb.WriteString("\n\n")

	if criterion.Rationale != "" {
		sb.WriteString("### Rationale\n\n")
		sb.WriteString(criterion.Rationale)
		sb.WriteString("\n\n")
	}

	if len(criterion.GoodExamples) > 0 || len(criterion.BadExamples) > 0 {
		sb.WriteString("### Examples\n\n")
		if len(criterion.GoodExamples) > 0 {
			sb.WriteString("**Good:**\n")
			for _, ex := range criterion.GoodExamples {
				fmt.Fprintf(&sb, "- `%s`\n", ex)
			}
			sb.WriteString("\n")
		}
		if len(criterion.BadExamples) > 0 {
			sb.WriteString("**Bad:**\n")
			for _, ex := range criterion.BadExamples {
				fmt.Fprintf(&sb, "- `%s`\n", ex)
			}
			sb.WriteString("\n")
		}
	}

	sb.WriteString("## OpenAPI Specification\n\n")
	sb.WriteString("```yaml\n")
	sb.WriteString(specContent)
	sb.WriteString("\n```\n\n")

	sb.WriteString("## Your Evaluation\n\n")
	sb.WriteString("Provide your evaluation in the following JSON format:\n\n")
	sb.WriteString("```json\n")
	sb.WriteString(`{
  "score": 0.0,
  "passed": false,
  "reasoning": "Your detailed explanation",
  "examples": ["Specific examples from the spec"],
  "suggestions": ["Improvement recommendations"],
  "locations": ["$.paths./users", "$.components.schemas.User"]
}`)
	sb.WriteString("\n```\n")

	return sb.String()
}

// BuildCategoryEvaluation creates a prompt for evaluating multiple criteria in a category.
func (pb *PromptBuilder) BuildCategoryEvaluation(category string, criteria []*Criterion, specContent string) string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "## Category: %s\n\n", category)
	sb.WriteString("Evaluate the following rules in this category:\n\n")

	for i, c := range criteria {
		fmt.Fprintf(&sb, "### Rule %d: %s (%s)\n\n", i+1, c.RuleTitle, c.RuleID)
		fmt.Fprintf(&sb, "**Severity:** %s\n", c.Severity)
		fmt.Fprintf(&sb, "**Criteria:** %s\n\n", c.Prompt)

		if c.Rationale != "" {
			fmt.Fprintf(&sb, "**Rationale:** %s\n\n", c.Rationale)
		}
	}

	sb.WriteString("## OpenAPI Specification\n\n")
	sb.WriteString("```yaml\n")
	sb.WriteString(specContent)
	sb.WriteString("\n```\n\n")

	sb.WriteString("## Your Evaluation\n\n")
	sb.WriteString("Provide your evaluation for each rule in the following JSON format:\n\n")
	sb.WriteString("```json\n")
	sb.WriteString(`{
  "findings": [
    {
      "ruleId": "RULE-001",
      "score": 0.0,
      "passed": false,
      "reasoning": "Your detailed explanation",
      "examples": ["Specific examples from the spec"],
      "suggestions": ["Improvement recommendations"],
      "locations": ["$.paths./users"]
    }
  ]
}`)
	sb.WriteString("\n```\n")

	return sb.String()
}

// BuildBatchEvaluation creates a prompt for evaluating all rules at once.
// Use sparingly - can hit token limits on large specs.
func (pb *PromptBuilder) BuildBatchEvaluation(rubricSet *RubricSet, specContent string) string {
	var sb strings.Builder

	sb.WriteString("## API Style Evaluation\n\n")
	fmt.Fprintf(&sb, "Evaluating against: **%s**\n\n", rubricSet.Name)

	for category, criteria := range rubricSet.Categories {
		fmt.Fprintf(&sb, "### %s\n\n", category)
		for _, c := range criteria {
			fmt.Fprintf(&sb, "- **%s** (%s, %s): %s\n",
				c.RuleTitle, c.RuleID, c.Severity, c.Prompt)
		}
		sb.WriteString("\n")
	}

	sb.WriteString("## OpenAPI Specification\n\n")
	sb.WriteString("```yaml\n")
	sb.WriteString(specContent)
	sb.WriteString("\n```\n\n")

	sb.WriteString("## Your Evaluation\n\n")
	sb.WriteString("Evaluate each rule and provide results in JSON format.\n")

	return sb.String()
}

// TruncateSpec truncates the spec content if it exceeds maxChars.
// Returns the truncated content and whether truncation occurred.
func TruncateSpec(content string, maxChars int) (string, bool) {
	if len(content) <= maxChars {
		return content, false
	}

	truncated := content[:maxChars]

	// Try to truncate at a line boundary
	if lastNewline := strings.LastIndex(truncated, "\n"); lastNewline > maxChars/2 {
		truncated = truncated[:lastNewline]
	}

	return truncated + "\n\n# ... (truncated) ...\n", true
}

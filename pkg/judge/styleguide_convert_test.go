package judge

import (
	"testing"
)

func TestStyleGuideReportToEvaluationReport(t *testing.T) {
	sg := &StyleGuideReport{
		ProfileName:  "test-profile",
		Status:       "partial",
		OverallScore: 3.5,
		Categories: []StyleGuideCategoryResult{
			{
				Category:     "content_coverage",
				Name:         "Content Coverage",
				Score:        "pass",
				NumericScore: 5,
				Weight:       0.2,
				Required:     true,
				Reasoning:    "Covers all essential areas.",
				Strengths:    []string{"Broad rule set", "Good examples"},
			},
			{
				Category:     "rule_quality",
				Name:         "Rule Quality",
				Score:        "partial",
				NumericScore: 3,
				Reasoning:    "Some rules lack rationale.",
				Weaknesses:   []string{"Missing rationale on 4 rules", "No severity on 2 rules"},
			},
			{
				Category:     "examples",
				Name:         "Examples",
				Score:        "fail",
				NumericScore: 1,
				Reasoning:    "No examples provided.",
				Weaknesses:   []string{"No good/bad examples"},
			},
		},
		Summary: StyleGuideSummary{
			TotalCategories:   3,
			PassedCategories:  1,
			PartialCategories: 1,
			FailedCategories:  1,
			WeightedScore:     3.5,
		},
		Metadata: StyleGuideReportMetadata{
			RubricName:    "style-guide-quality",
			RubricVersion: "1.0.0",
			Model:         "claude-sonnet-5",
			Timestamp:     "2026-08-30T12:00:00Z",
		},
	}

	eval := sg.ToEvaluationReport()

	if eval.ReviewType != "style-guide-quality" {
		t.Errorf("ReviewType = %q, want style-guide-quality", eval.ReviewType)
	}
	if eval.RubricID != "style-guide-quality" || eval.RubricVersion != "1.0.0" {
		t.Errorf("rubric = %q/%q, want style-guide-quality/1.0.0", eval.RubricID, eval.RubricVersion)
	}
	if eval.OverallDecision != "partial" {
		t.Errorf("OverallDecision = %q, want partial", eval.OverallDecision)
	}
	if eval.Metadata == nil || eval.Metadata.Document != "test-profile" {
		t.Fatalf("Metadata.Document not mapped: %+v", eval.Metadata)
	}
	if eval.Metadata.GeneratedAt.IsZero() {
		t.Error("GeneratedAt not parsed from timestamp")
	}

	if len(eval.Categories) != 3 {
		t.Fatalf("len(Categories) = %d, want 3", len(eval.Categories))
	}
	if eval.Categories[0].Reasoning == sg.Categories[0].Reasoning {
		t.Error("strengths not folded into reasoning")
	}

	// Weaknesses become findings with score-derived severity.
	if len(eval.Findings) != 3 {
		t.Fatalf("len(Findings) = %d, want 3", len(eval.Findings))
	}
	if eval.Findings[0].Severity != "medium" {
		t.Errorf("partial-category finding severity = %q, want medium", eval.Findings[0].Severity)
	}
	if eval.Findings[2].Severity != "high" {
		t.Errorf("fail-category finding severity = %q, want high", eval.Findings[2].Severity)
	}

	if eval.Decision == nil {
		t.Fatal("Decision is nil")
	}
	cc := eval.Decision.CategoryCounts
	if cc == nil || cc.Pass != 1 || cc.Partial != 1 || cc.Fail != 1 || cc.Total != 3 {
		t.Errorf("CategoryCounts = %+v, want 1/1/1/3", cc)
	}
	fc := eval.Decision.FindingCounts
	if fc == nil || fc.Medium != 2 || fc.High != 1 || fc.Total != 3 {
		t.Errorf("FindingCounts = %+v, want medium=2 high=1 total=3", fc)
	}
}

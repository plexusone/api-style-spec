package types

import "testing"

func TestLintReportAddViolation(t *testing.T) {
	r := NewLintReport()
	if r.Status != ReportStatusPass {
		t.Fatalf("new report status = %q, want pass", r.Status)
	}

	r.AddViolation(Violation{RuleID: "R1", Severity: SeverityWarn})
	r.AddViolation(Violation{RuleID: "R2", Severity: SeverityInfo})
	r.AddViolation(Violation{RuleID: "R3", Severity: SeverityHint})

	if r.Status != ReportStatusPass {
		t.Errorf("status after non-error violations = %q, want pass", r.Status)
	}
	if r.HasBlockingViolations() {
		t.Error("HasBlockingViolations() = true without errors")
	}

	r.AddViolation(Violation{RuleID: "R4", Severity: SeverityError})

	if r.Status != ReportStatusFail {
		t.Errorf("status after error violation = %q, want fail", r.Status)
	}
	if !r.HasBlockingViolations() {
		t.Error("HasBlockingViolations() = false with an error")
	}

	s := r.Summary
	if s.Errors != 1 || s.Warnings != 1 || s.Infos != 1 || s.Hints != 1 || s.Total != 4 {
		t.Errorf("summary = %+v, want 1/1/1/1 total 4", s)
	}
}

func TestMultiLintReportAggregation(t *testing.T) {
	pass := NewLintReport()
	pass.AddViolation(Violation{RuleID: "R1", Severity: SeverityWarn})

	fail := NewLintReport()
	fail.AddViolation(Violation{RuleID: "R2", Severity: SeverityError})

	m := NewMultiLintReport()
	m.AddFileReport("a.yaml", pass)
	if m.Status != ReportStatusPass {
		t.Errorf("status after passing file = %q, want pass", m.Status)
	}

	m.AddFileReport("b.yaml", fail)
	if m.Status != ReportStatusFail {
		t.Errorf("status after failing file = %q, want fail", m.Status)
	}
	if !m.HasBlockingViolations() {
		t.Error("HasBlockingViolations() = false with a failing file")
	}
	if m.FileCount() != 2 {
		t.Errorf("FileCount() = %d, want 2", m.FileCount())
	}
	if m.FailedFileCount() != 1 {
		t.Errorf("FailedFileCount() = %d, want 1", m.FailedFileCount())
	}
	if m.Summary.Total != 2 || m.Summary.Errors != 1 || m.Summary.Warnings != 1 {
		t.Errorf("aggregate summary = %+v, want total 2, errors 1, warnings 1", m.Summary)
	}
}

func TestSeverity(t *testing.T) {
	if !SeverityError.IsBlocking() {
		t.Error("error severity should block")
	}
	for _, s := range []Severity{SeverityWarn, SeverityInfo, SeverityHint} {
		if s.IsBlocking() {
			t.Errorf("%s severity should not block", s)
		}
	}
	if SeverityError.Weight() <= SeverityWarn.Weight() {
		t.Error("error should weigh more than warn")
	}
	if SeverityWarn.Weight() <= SeverityInfo.Weight() {
		t.Error("warn should weigh more than info")
	}
}

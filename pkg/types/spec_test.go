package types

import "testing"

func testSpec() *APIStyleSpec {
	return &APIStyleSpec{
		Name: "test",
		Categories: []Category{
			{ID: "naming", Title: "Naming"},
			{ID: "errors", Title: "Errors"},
		},
		Sections: []Section{
			{ID: "intro", Title: "Introduction"},
		},
		Patterns: []Pattern{
			{ID: "pagination", Name: "Pagination"},
		},
		Rules: []Rule{
			{ID: "R1", Category: "naming"},
			{ID: "R2", Category: "naming"},
			{ID: "R3", Category: "errors"},
		},
	}
}

func TestSpecLookups(t *testing.T) {
	s := testSpec()

	if got := len(s.RulesForCategory("naming")); got != 2 {
		t.Errorf("RulesForCategory(naming) = %d rules, want 2", got)
	}
	if got := len(s.RulesForCategory("missing")); got != 0 {
		t.Errorf("RulesForCategory(missing) = %d rules, want 0", got)
	}

	if c := s.GetCategory("errors"); c == nil || c.Title != "Errors" {
		t.Errorf("GetCategory(errors) = %+v", c)
	}
	if s.GetCategory("missing") != nil {
		t.Error("GetCategory(missing) should be nil")
	}

	if r := s.GetRule("R3"); r == nil || r.Category != "errors" {
		t.Errorf("GetRule(R3) = %+v", r)
	}
	if s.GetRule("missing") != nil {
		t.Error("GetRule(missing) should be nil")
	}

	if p := s.GetPattern("pagination"); p == nil {
		t.Error("GetPattern(pagination) = nil")
	}
	if sec := s.GetSection("intro"); sec == nil {
		t.Error("GetSection(intro) = nil")
	}
}

func TestGivenPathsConstructors(t *testing.T) {
	if g := NewGivenPath("$.info"); len(g.Paths) != 1 || g.Paths[0] != "$.info" {
		t.Errorf("NewGivenPath = %+v", g)
	}
	if g := NewGivenPaths("$.a", "$.b"); len(g.Paths) != 2 {
		t.Errorf("NewGivenPaths = %+v", g)
	}
}

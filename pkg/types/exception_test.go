package types

import (
	"testing"
	"time"
)

func TestExceptionMatchesPathGlobs(t *testing.T) {
	tests := []struct {
		name  string
		scope *ExceptionScope
		path  string
		want  bool
	}{
		{"nil scope handled by caller; empty scope matches all", &ExceptionScope{}, "/v1/users", true},
		{"exact match", &ExceptionScope{Path: "/v1/users"}, "/v1/users", true},
		{"exact mismatch", &ExceptionScope{Path: "/v1/users"}, "/v1/orders", false},
		{"single star within segment", &ExceptionScope{Path: "/v1/*/status"}, "/v1/users/status", true},
		{"single star does not cross segments", &ExceptionScope{Path: "/v1/*"}, "/v1/users/status", false},
		{"double star crosses segments", &ExceptionScope{Path: "/v1/**"}, "/v1/users/status", true},
		{"double star mid-pattern", &ExceptionScope{Path: "/v1/**/status"}, "/v1/a/b/status", true},
		{"question mark matches one char", &ExceptionScope{Path: "/v?/users"}, "/v1/users", true},
		{"question mark does not match slash", &ExceptionScope{Path: "/v1?users"}, "/v1/users", false},
		{"regex metachars are literal", &ExceptionScope{Path: "/v1/users.json"}, "/v1/usersXjson", false},
		{"paths list matches any", &ExceptionScope{Paths: []string{"/v1/a", "/v1/b"}}, "/v1/b", true},
		{"paths list no match", &ExceptionScope{Paths: []string{"/v1/a", "/v1/b"}}, "/v1/c", false},
		{"path and paths combined", &ExceptionScope{Path: "/v1/a", Paths: []string{"/v1/b"}}, "/v1/b", true},
		{"glob in paths list", &ExceptionScope{Paths: []string{"/legacy/**"}}, "/legacy/x/y", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exc := &Exception{RuleID: "R1", AppliesTo: tt.scope}
			if got := exc.Matches("R1", "", tt.path, ""); got != tt.want {
				t.Errorf("Matches(path=%q) with scope %+v = %v, want %v", tt.path, tt.scope, got, tt.want)
			}
		})
	}
}

func TestExceptionMatchesScopeFields(t *testing.T) {
	past := time.Now().Add(-time.Hour)
	future := time.Now().Add(time.Hour)

	tests := []struct {
		name string
		exc  *Exception
		rule string
		api  string
		op   string
		want bool
	}{
		{"rule mismatch", &Exception{RuleID: "R1"}, "R2", "", "", false},
		{"nil scope applies everywhere", &Exception{RuleID: "R1"}, "R1", "any", "any", true},
		{"api mismatch", &Exception{RuleID: "R1", AppliesTo: &ExceptionScope{API: "billing"}}, "R1", "users", "", false},
		{"operation mismatch", &Exception{RuleID: "R1", AppliesTo: &ExceptionScope{Operation: "GET"}}, "R1", "", "POST", false},
		{"expired exception never matches", &Exception{RuleID: "R1", ExpiresOn: &past}, "R1", "", "", false},
		{"unexpired exception matches", &Exception{RuleID: "R1", ExpiresOn: &future}, "R1", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.exc.Matches(tt.rule, tt.api, "/v1/x", tt.op); got != tt.want {
				t.Errorf("Matches(%q, api=%q, op=%q) = %v, want %v", tt.rule, tt.api, tt.op, got, tt.want)
			}
		})
	}
}

package schema

import "testing"

func TestEmbeddedSchemasParse(t *testing.T) {
	for _, name := range []string{"api-style-spec.schema.json", "lint-report.schema.json"} {
		t.Run(name, func(t *testing.T) {
			m, err := Map(name)
			if err != nil {
				t.Fatalf("Map(%q) error = %v", name, err)
			}
			if m["$schema"] == nil {
				t.Errorf("schema %q missing $schema key", name)
			}
			if m["properties"] == nil && m["$defs"] == nil {
				t.Errorf("schema %q has neither properties nor $defs", name)
			}
		})
	}
}

func TestSchemaAccessors(t *testing.T) {
	if data, err := APIStyleSpecSchema(); err != nil || len(data) == 0 {
		t.Errorf("APIStyleSpecSchema() = %d bytes, err %v", len(data), err)
	}
	if data, err := LintReportSchema(); err != nil || len(data) == 0 {
		t.Errorf("LintReportSchema() = %d bytes, err %v", len(data), err)
	}
	if _, err := Schema("nonexistent.schema.json"); err == nil {
		t.Error("Schema(nonexistent) should return an error")
	}
}

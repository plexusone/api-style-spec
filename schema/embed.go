// Package schema embeds JSON Schema files for runtime access.
package schema

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed *.schema.json
var schemaFS embed.FS

// APIStyleSpecSchema returns the JSON Schema for APIStyleSpec.
func APIStyleSpecSchema() ([]byte, error) {
	return schemaFS.ReadFile("api-style-spec.schema.json")
}

// LintReportSchema returns the JSON Schema for LintReport.
func LintReportSchema() ([]byte, error) {
	return schemaFS.ReadFile("lint-report.schema.json")
}

// Schema returns a schema by filename.
func Schema(name string) ([]byte, error) {
	data, err := schemaFS.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("schema %q not found: %w", name, err)
	}
	return data, nil
}

// Map returns a schema parsed as a map.
func Map(name string) (map[string]any, error) {
	data, err := Schema(name)
	if err != nil {
		return nil, err
	}

	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("parsing schema %q: %w", name, err)
	}
	return m, nil
}

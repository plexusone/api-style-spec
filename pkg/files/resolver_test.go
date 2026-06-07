package files

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsOpenAPIFile(t *testing.T) {
	tests := []struct {
		path     string
		expected bool
	}{
		{"openapi.yaml", true},
		{"openapi.yml", true},
		{"openapi.json", true},
		{"swagger.yaml", true},
		{"swagger.json", true},
		{"api.yaml", true},
		{"api.yml", true},
		{"api-spec.yaml", true},
		{"my-api-spec.json", true},
		{"random.yaml", false},
		{"config.json", false},
		{"README.md", false},
		{"api/openapi.yaml", true},
		{"specs/swagger.json", true},
	}

	for _, tc := range tests {
		t.Run(tc.path, func(t *testing.T) {
			result := IsOpenAPIFile(tc.path)
			if result != tc.expected {
				t.Errorf("IsOpenAPIFile(%q) = %v, want %v", tc.path, result, tc.expected)
			}
		})
	}
}

func TestMatchGlob(t *testing.T) {
	tests := []struct {
		path    string
		pattern string
		want    bool
	}{
		{"api.yaml", "*.yaml", true},
		{"api.json", "*.yaml", false},
		{"api/openapi.yaml", "**/openapi.yaml", true},
		{"deep/nested/openapi.yaml", "**/openapi.yaml", true},
		{"openapi.yaml", "openapi.yaml", true},
		{"other.yaml", "openapi.yaml", false},
	}

	for _, tc := range tests {
		t.Run(tc.path+"_"+tc.pattern, func(t *testing.T) {
			got, err := MatchGlob(tc.path, tc.pattern)
			if err != nil {
				t.Fatalf("MatchGlob error: %v", err)
			}
			if got != tc.want {
				t.Errorf("MatchGlob(%q, %q) = %v, want %v", tc.path, tc.pattern, got, tc.want)
			}
		})
	}
}

func TestResolveSpecs(t *testing.T) {
	// Create test directory structure
	tmpDir := t.TempDir()

	// Create files
	files := []string{
		"openapi.yaml",
		"swagger.json",
		"api/openapi.yaml",
		"api/v2/openapi.yaml",
		"config.yaml",
		"random.json",
	}

	for _, f := range files {
		path := filepath.Join(tmpDir, f)
		_ = os.MkdirAll(filepath.Dir(path), 0o755)
		_ = os.WriteFile(path, []byte("test"), 0o600)
	}

	t.Run("single file", func(t *testing.T) {
		result, err := ResolveSpecs([]string{filepath.Join(tmpDir, "openapi.yaml")}, false, nil, nil)
		if err != nil {
			t.Fatalf("ResolveSpecs error: %v", err)
		}
		if len(result) != 1 {
			t.Errorf("Expected 1 file, got %d", len(result))
		}
	})

	t.Run("glob pattern", func(t *testing.T) {
		result, err := ResolveSpecs([]string{filepath.Join(tmpDir, "*.yaml")}, false, nil, nil)
		if err != nil {
			t.Fatalf("ResolveSpecs error: %v", err)
		}
		// Should match openapi.yaml and config.yaml, but config.yaml isn't OpenAPI
		// Since we're using glob directly, it returns both
		if len(result) < 1 {
			t.Errorf("Expected at least 1 file, got %d", len(result))
		}
	})

	t.Run("directory non-recursive", func(t *testing.T) {
		result, err := ResolveSpecs([]string{tmpDir}, false, nil, nil)
		if err != nil {
			t.Fatalf("ResolveSpecs error: %v", err)
		}
		// Should find openapi.yaml and swagger.json in root only
		if len(result) != 2 {
			t.Errorf("Expected 2 files, got %d: %v", len(result), result)
		}
	})

	t.Run("directory recursive", func(t *testing.T) {
		result, err := ResolveSpecs([]string{tmpDir}, true, nil, nil)
		if err != nil {
			t.Fatalf("ResolveSpecs error: %v", err)
		}
		// Should find all OpenAPI files
		if len(result) != 4 {
			t.Errorf("Expected 4 files, got %d: %v", len(result), result)
		}
	})

	t.Run("with exclude", func(t *testing.T) {
		result, err := ResolveSpecs([]string{tmpDir}, true, nil, []string{"**/v2/**"})
		if err != nil {
			t.Fatalf("ResolveSpecs error: %v", err)
		}
		// Should exclude api/v2/openapi.yaml
		if len(result) != 3 {
			t.Errorf("Expected 3 files, got %d: %v", len(result), result)
		}
	})
}

func TestShouldInclude(t *testing.T) {
	tests := []struct {
		path    string
		include []string
		exclude []string
		want    bool
	}{
		// No filters
		{"api.yaml", nil, nil, true},
		// Include only
		{"api.yaml", []string{"*.yaml"}, nil, true},
		{"api.json", []string{"*.yaml"}, nil, false},
		// Exclude only
		{"api.yaml", nil, []string{"*.json"}, true},
		{"api.json", nil, []string{"*.json"}, false},
		// Both
		{"api.yaml", []string{"*.yaml"}, []string{"test*.yaml"}, true},
		{"test.yaml", []string{"*.yaml"}, []string{"test*.yaml"}, false},
	}

	for _, tc := range tests {
		t.Run(tc.path, func(t *testing.T) {
			got := shouldInclude(tc.path, tc.include, tc.exclude)
			if got != tc.want {
				t.Errorf("shouldInclude(%q, %v, %v) = %v, want %v",
					tc.path, tc.include, tc.exclude, got, tc.want)
			}
		})
	}
}

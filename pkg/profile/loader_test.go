package profile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_Default(t *testing.T) {
	spec, err := Load("default")
	if err != nil {
		t.Fatalf("Load(default) failed: %v", err)
	}

	if spec.Name != "default" {
		t.Errorf("Name = %q, want %q", spec.Name, "default")
	}

	if spec.Version != "2.3.0" {
		t.Errorf("Version = %q, want %q", spec.Version, "2.3.0")
	}

	// Should have 100+ rules (industry-leading profile with SDK focus)
	if len(spec.Rules) < 100 {
		t.Errorf("len(Rules) = %d, want >= 100", len(spec.Rules))
	}

	t.Logf("Loaded %d rules from default profile", len(spec.Rules))

	// Check some expected rules exist (PO-xxx for PlexusOne rules)
	ruleIDs := make(map[string]bool)
	for _, rule := range spec.Rules {
		ruleIDs[rule.ID] = true
	}

	expectedRules := []string{"PO-001", "PO-002", "PO-003", "PO-006", "PO-008", "PO-022"}
	for _, id := range expectedRules {
		if !ruleIDs[id] {
			t.Errorf("Expected rule %q not found", id)
		}
	}
}

func TestLoad_NotFound(t *testing.T) {
	_, err := Load("nonexistent-profile")
	if err == nil {
		t.Error("Load(nonexistent) should fail")
	}
}

func TestListBuiltin(t *testing.T) {
	names, err := ListBuiltin()
	if err != nil {
		t.Fatalf("ListBuiltin failed: %v", err)
	}

	if len(names) == 0 {
		t.Error("ListBuiltin returned no profiles")
	}

	t.Logf("Built-in profiles: %v", names)

	// Should include 'default'
	found := false
	for _, name := range names {
		if name == "default" {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected 'default' in built-in profiles")
	}
}

func TestDefaultProfile_Categories(t *testing.T) {
	spec, err := Load("default")
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	expectedCategories := []string{
		"general",
		"naming",
		"urls",
		"http-methods",
		"http-status",
		"errors",
		"security",
		"documentation",
	}

	categoryMap := make(map[string]bool)
	for _, cat := range spec.Categories {
		categoryMap[cat.ID] = true
	}

	for _, expected := range expectedCategories {
		if !categoryMap[expected] {
			t.Errorf("Expected category %q not found", expected)
		}
	}
}

func TestDefaultProfile_ConformanceLevels(t *testing.T) {
	spec, err := Load("default")
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	expectedLevels := []string{"minimum", "standard", "exemplary"}

	for _, level := range expectedLevels {
		if _, ok := spec.ConformanceLevels[level]; !ok {
			t.Errorf("Expected conformance level %q not found", level)
		}
	}

	// Standard should have more required rules than minimum
	standard := spec.ConformanceLevels["standard"]
	minimum := spec.ConformanceLevels["minimum"]
	if len(standard.RequiredRules) <= len(minimum.RequiredRules) {
		t.Errorf("Standard should have more required rules than minimum")
	}
}

func TestLoad_Azure(t *testing.T) {
	spec, err := Load("azure")
	if err != nil {
		t.Fatalf("Load(azure) failed: %v", err)
	}

	if spec.Name != "Azure REST API Guidelines" {
		t.Errorf("Name = %q, want %q", spec.Name, "Azure REST API Guidelines")
	}

	// Should have ~23 rules
	if len(spec.Rules) < 20 {
		t.Errorf("len(Rules) = %d, want >= 20", len(spec.Rules))
	}

	t.Logf("Loaded %d rules from azure profile", len(spec.Rules))

	// Check some expected rules exist
	ruleIDs := make(map[string]bool)
	for _, rule := range spec.Rules {
		ruleIDs[rule.ID] = true
	}

	expectedRules := []string{"AZ-VER-001", "AZ-URI-001", "AZ-REQ-001", "AZ-SEC-001"}
	for _, id := range expectedRules {
		if !ruleIDs[id] {
			t.Errorf("Expected rule %q not found", id)
		}
	}
}

func TestLoad_Google(t *testing.T) {
	spec, err := Load("google")
	if err != nil {
		t.Fatalf("Load(google) failed: %v", err)
	}

	if spec.Name != "Google Cloud API Design Guide" {
		t.Errorf("Name = %q, want %q", spec.Name, "Google Cloud API Design Guide")
	}

	// Should have ~20 rules
	if len(spec.Rules) < 15 {
		t.Errorf("len(Rules) = %d, want >= 15", len(spec.Rules))
	}

	t.Logf("Loaded %d rules from google profile", len(spec.Rules))

	// Check some expected rules exist
	ruleIDs := make(map[string]bool)
	for _, rule := range spec.Rules {
		ruleIDs[rule.ID] = true
	}

	expectedRules := []string{"GCP-RES-001", "GCP-NAME-001", "GCP-ERR-001"}
	for _, id := range expectedRules {
		if !ruleIDs[id] {
			t.Errorf("Expected rule %q not found", id)
		}
	}
}

func TestLoad_Zalando(t *testing.T) {
	spec, err := Load("zalando")
	if err != nil {
		t.Fatalf("Load(zalando) failed: %v", err)
	}

	if spec.Name != "Zalando REST API Guidelines" {
		t.Errorf("Name = %q, want %q", spec.Name, "Zalando REST API Guidelines")
	}

	// Should have ~147 rules (comprehensive version)
	if len(spec.Rules) < 100 {
		t.Errorf("len(Rules) = %d, want >= 100", len(spec.Rules))
	}

	t.Logf("Loaded %d rules from zalando profile", len(spec.Rules))

	// Check some expected rules exist (official Zalando Z-XXX IDs)
	ruleIDs := make(map[string]bool)
	for _, rule := range spec.Rules {
		ruleIDs[rule.ID] = true
	}

	expectedRules := []string{"Z-100", "Z-101", "Z-118", "Z-176"}
	for _, id := range expectedRules {
		if !ruleIDs[id] {
			t.Errorf("Expected rule %q not found", id)
		}
	}
}

func TestLoadFile_JSON(t *testing.T) {
	// Create a temporary JSON profile
	tmpDir, err := os.MkdirTemp("", "profile-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	profileContent := `{
		"name": "test-json",
		"version": "1.0.0",
		"description": "Test JSON profile",
		"rules": [
			{
				"id": "TEST-001",
				"title": "Test Rule",
				"category": "test",
				"severity": "error"
			}
		]
	}`

	profilePath := filepath.Join(tmpDir, "test.json")
	if err := os.WriteFile(profilePath, []byte(profileContent), 0o600); err != nil {
		t.Fatalf("failed to write profile: %v", err)
	}

	spec, err := LoadFile(profilePath)
	if err != nil {
		t.Fatalf("LoadFile failed: %v", err)
	}

	if spec.Name != "test-json" {
		t.Errorf("Name = %q, want %q", spec.Name, "test-json")
	}

	if len(spec.Rules) != 1 {
		t.Errorf("len(Rules) = %d, want 1", len(spec.Rules))
	}

	if spec.Rules[0].ID != "TEST-001" {
		t.Errorf("Rule ID = %q, want %q", spec.Rules[0].ID, "TEST-001")
	}
}

func TestLoadFile_YAML(t *testing.T) {
	// Create a temporary YAML profile
	tmpDir, err := os.MkdirTemp("", "profile-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	profileContent := `
name: test-yaml
version: "1.0.0"
description: Test YAML profile
rules:
  - id: TEST-001
    title: Test Rule
    category: test
    severity: error
    rationale: This is a test rule
  - id: TEST-002
    title: Another Test Rule
    category: test
    severity: warn
`

	profilePath := filepath.Join(tmpDir, "test.yaml")
	if err := os.WriteFile(profilePath, []byte(profileContent), 0o600); err != nil {
		t.Fatalf("failed to write profile: %v", err)
	}

	spec, err := LoadFile(profilePath)
	if err != nil {
		t.Fatalf("LoadFile failed: %v", err)
	}

	if spec.Name != "test-yaml" {
		t.Errorf("Name = %q, want %q", spec.Name, "test-yaml")
	}

	if len(spec.Rules) != 2 {
		t.Errorf("len(Rules) = %d, want 2", len(spec.Rules))
	}

	if spec.Rules[0].ID != "TEST-001" {
		t.Errorf("Rule ID = %q, want %q", spec.Rules[0].ID, "TEST-001")
	}

	if spec.Rules[1].Severity != "warn" {
		t.Errorf("Rule severity = %q, want %q", spec.Rules[1].Severity, "warn")
	}
}

func TestLoadFile_YML(t *testing.T) {
	// Create a temporary YML profile (alternate extension)
	tmpDir, err := os.MkdirTemp("", "profile-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	profileContent := `
name: test-yml
version: "1.0.0"
rules: []
`

	profilePath := filepath.Join(tmpDir, "test.yml")
	if err := os.WriteFile(profilePath, []byte(profileContent), 0o600); err != nil {
		t.Fatalf("failed to write profile: %v", err)
	}

	spec, err := LoadFile(profilePath)
	if err != nil {
		t.Fatalf("LoadFile failed: %v", err)
	}

	if spec.Name != "test-yml" {
		t.Errorf("Name = %q, want %q", spec.Name, "test-yml")
	}
}

func TestLoadFile_ValidationError_MissingName(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "profile-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Profile missing required name
	profileContent := `{
		"version": "1.0.0",
		"rules": []
	}`

	profilePath := filepath.Join(tmpDir, "invalid.json")
	if err := os.WriteFile(profilePath, []byte(profileContent), 0o600); err != nil {
		t.Fatalf("failed to write profile: %v", err)
	}

	_, err = LoadFile(profilePath)
	if err == nil {
		t.Error("LoadFile should fail for profile missing name")
	}

	if err.Error() == "" || !contains(err.Error(), "name is required") {
		t.Errorf("expected error about missing name, got: %v", err)
	}
}

func TestLoadFile_ValidationError_MissingVersion(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "profile-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Profile missing required version
	profileContent := `{
		"name": "test",
		"rules": []
	}`

	profilePath := filepath.Join(tmpDir, "invalid.json")
	if err := os.WriteFile(profilePath, []byte(profileContent), 0o600); err != nil {
		t.Fatalf("failed to write profile: %v", err)
	}

	_, err = LoadFile(profilePath)
	if err == nil {
		t.Error("LoadFile should fail for profile missing version")
	}

	if !contains(err.Error(), "version is required") {
		t.Errorf("expected error about missing version, got: %v", err)
	}
}

func TestLoader_SearchPaths(t *testing.T) {
	// Create a temp directory structure with a profile
	tmpDir, err := os.MkdirTemp("", "profile-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	profilesDir := filepath.Join(tmpDir, "profiles")
	if err := os.MkdirAll(profilesDir, 0o755); err != nil {
		t.Fatalf("failed to create profiles dir: %v", err)
	}

	profileContent := `{
		"name": "custom-profile",
		"version": "2.0.0",
		"rules": []
	}`

	profilePath := filepath.Join(profilesDir, "custom.api-style.json")
	if err := os.WriteFile(profilePath, []byte(profileContent), 0o600); err != nil {
		t.Fatalf("failed to write profile: %v", err)
	}

	// Create loader with custom search paths
	loader := &Loader{
		SearchPaths: []string{profilesDir},
	}

	spec, err := loader.Load("custom")
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if spec.Name != "custom-profile" {
		t.Errorf("Name = %q, want %q", spec.Name, "custom-profile")
	}

	if spec.Version != "2.0.0" {
		t.Errorf("Version = %q, want %q", spec.Version, "2.0.0")
	}
}

func TestLoader_SearchPaths_YAML(t *testing.T) {
	// Create a temp directory structure with a YAML profile
	tmpDir, err := os.MkdirTemp("", "profile-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	profileContent := `
name: yaml-profile
version: "1.5.0"
description: A YAML profile found via search path
rules:
  - id: YAML-001
    title: YAML Test
    category: test
    severity: info
`

	profilePath := filepath.Join(tmpDir, "myprofile.api-style.yaml")
	if err := os.WriteFile(profilePath, []byte(profileContent), 0o600); err != nil {
		t.Fatalf("failed to write profile: %v", err)
	}

	loader := &Loader{
		SearchPaths: []string{tmpDir},
	}

	spec, err := loader.Load("myprofile")
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if spec.Name != "yaml-profile" {
		t.Errorf("Name = %q, want %q", spec.Name, "yaml-profile")
	}

	if len(spec.Rules) != 1 {
		t.Errorf("len(Rules) = %d, want 1", len(spec.Rules))
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsAt(s, substr, 0))
}

func containsAt(s, substr string, start int) bool {
	for i := start; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Exemplar tests

func TestListExemplars(t *testing.T) {
	exemplars, err := ListExemplars()
	if err != nil {
		t.Fatalf("ListExemplars failed: %v", err)
	}

	if len(exemplars) == 0 {
		t.Error("ListExemplars returned no exemplars")
	}

	t.Logf("Found %d exemplars", len(exemplars))

	// Verify structure of exemplars
	for _, e := range exemplars {
		if e.Name == "" {
			t.Error("Exemplar has empty name")
		}
		if e.Profile == "" {
			t.Error("Exemplar has empty profile")
		}
		if len(e.Content) == 0 {
			t.Errorf("Exemplar %q has no content", e.Name)
		}
		t.Logf("  - %s (profile: %s, %d bytes)", e.Name, e.Profile, len(e.Content))
	}
}

func TestListExemplars_ContainsExpected(t *testing.T) {
	exemplars, err := ListExemplars()
	if err != nil {
		t.Fatalf("ListExemplars failed: %v", err)
	}

	expectedNames := map[string]bool{
		"default-minimal": false,
		"default-crud":    false,
	}

	for _, e := range exemplars {
		if _, ok := expectedNames[e.Name]; ok {
			expectedNames[e.Name] = true
		}
	}

	for name, found := range expectedNames {
		if !found {
			t.Errorf("Expected exemplar %q not found", name)
		}
	}
}

func TestListExemplarsForProfile(t *testing.T) {
	exemplars, err := ListExemplarsForProfile("default")
	if err != nil {
		t.Fatalf("ListExemplarsForProfile failed: %v", err)
	}

	if len(exemplars) == 0 {
		t.Error("No exemplars found for 'default' profile")
	}

	for _, e := range exemplars {
		if e.Profile != "default" {
			t.Errorf("Exemplar %q has profile %q, expected 'default'", e.Name, e.Profile)
		}
	}
}

func TestListExemplarsForProfile_Empty(t *testing.T) {
	exemplars, err := ListExemplarsForProfile("nonexistent")
	if err != nil {
		t.Fatalf("ListExemplarsForProfile failed: %v", err)
	}

	if len(exemplars) != 0 {
		t.Errorf("Expected 0 exemplars for nonexistent profile, got %d", len(exemplars))
	}
}

func TestGetExemplar(t *testing.T) {
	exemplar, err := GetExemplar("default-minimal")
	if err != nil {
		t.Fatalf("GetExemplar failed: %v", err)
	}

	if exemplar.Name != "default-minimal" {
		t.Errorf("Name = %q, want %q", exemplar.Name, "default-minimal")
	}

	if exemplar.Profile != "default" {
		t.Errorf("Profile = %q, want %q", exemplar.Profile, "default")
	}

	if len(exemplar.Content) == 0 {
		t.Error("Content is empty")
	}

	// Verify it's valid YAML/OpenAPI
	if !contains(string(exemplar.Content), "openapi:") {
		t.Error("Content should contain 'openapi:' declaration")
	}

	if exemplar.Description == "" {
		t.Error("Description should be extracted from content")
	}
}

func TestGetExemplar_NotFound(t *testing.T) {
	_, err := GetExemplar("nonexistent-exemplar")
	if err == nil {
		t.Error("GetExemplar should fail for nonexistent exemplar")
	}
}

func TestExemplar_Description(t *testing.T) {
	exemplar, err := GetExemplar("default-minimal")
	if err != nil {
		t.Fatalf("GetExemplar failed: %v", err)
	}

	// Description should be extracted from OpenAPI info.description
	if exemplar.Description == "" {
		t.Error("Description should not be empty")
	}

	// Should contain something about "minimal" or "conformant"
	desc := exemplar.Description
	if !contains(desc, "minimal") && !contains(desc, "Minimal") {
		t.Logf("Description: %s", desc)
		// This is a warning, not an error - description content may vary
	}
}

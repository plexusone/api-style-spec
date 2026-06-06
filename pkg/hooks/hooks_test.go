package hooks

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.Profile != "default" {
		t.Errorf("expected profile 'default', got %q", cfg.Profile)
	}

	if !cfg.AutoLint {
		t.Error("expected AutoLint to be true")
	}

	if cfg.InjectContext {
		t.Error("expected InjectContext to be false by default")
	}

	if len(cfg.AutoLintPatterns) == 0 {
		t.Error("expected non-empty AutoLintPatterns")
	}

	// Check for common OpenAPI patterns
	patterns := strings.Join(cfg.AutoLintPatterns, " ")
	if !strings.Contains(patterns, "openapi.yaml") {
		t.Error("expected openapi.yaml in patterns")
	}
	if !strings.Contains(patterns, "swagger.json") {
		t.Error("expected swagger.json in patterns")
	}
}

func TestConfig_Generate(t *testing.T) {
	cfg := &Config{
		Profile:  "azure",
		AutoLint: true,
	}

	hooksCfg := cfg.Generate()

	if hooksCfg == nil {
		t.Fatal("expected non-nil hooks config")
	}
}

func TestConfig_Generate_WithContextInjection(t *testing.T) {
	cfg := &Config{
		Profile:       "google",
		AutoLint:      true,
		InjectContext: true,
	}

	hooksCfg := cfg.Generate()

	if hooksCfg == nil {
		t.Fatal("expected non-nil hooks config")
	}
}

func TestConfig_Generate_NoAutoLint(t *testing.T) {
	cfg := &Config{
		Profile:  "default",
		AutoLint: false,
	}

	hooksCfg := cfg.Generate()

	if hooksCfg == nil {
		t.Fatal("expected non-nil hooks config")
	}
}

func TestConfig_MarshalFormat_Claude(t *testing.T) {
	cfg := &Config{
		Profile:  "default",
		AutoLint: true,
	}

	data, err := cfg.MarshalFormat("claude")
	if err != nil {
		t.Fatalf("MarshalFormat failed: %v", err)
	}

	if len(data) == 0 {
		t.Error("expected non-empty output")
	}

	// Verify it's valid JSON
	var parsed map[string]any
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}

	// Check for hooks key
	if _, ok := parsed["hooks"]; !ok {
		t.Error("expected 'hooks' key in output")
	}
}

func TestConfig_MarshalFormat_Cursor(t *testing.T) {
	cfg := &Config{
		Profile:  "azure",
		AutoLint: true,
	}

	data, err := cfg.MarshalFormat("cursor")
	if err != nil {
		t.Fatalf("MarshalFormat failed: %v", err)
	}

	if len(data) == 0 {
		t.Error("expected non-empty output")
	}

	// Verify it's valid JSON
	var parsed map[string]any
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}
}

func TestConfig_MarshalFormat_Windsurf(t *testing.T) {
	cfg := &Config{
		Profile:  "zalando",
		AutoLint: true,
	}

	data, err := cfg.MarshalFormat("windsurf")
	if err != nil {
		t.Fatalf("MarshalFormat failed: %v", err)
	}

	if len(data) == 0 {
		t.Error("expected non-empty output")
	}

	// Verify it's valid JSON
	var parsed map[string]any
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}
}

func TestConfig_MarshalFormat_UnsupportedFormat(t *testing.T) {
	cfg := &Config{
		Profile:  "default",
		AutoLint: true,
	}

	_, err := cfg.MarshalFormat("unsupported")
	if err == nil {
		t.Error("expected error for unsupported format")
	}

	if !strings.Contains(err.Error(), "unsupported format") {
		t.Errorf("expected 'unsupported format' in error, got: %v", err)
	}
}

func TestConfig_WriteToFile(t *testing.T) {
	cfg := &Config{
		Profile:  "default",
		AutoLint: true,
	}

	// Create temp directory
	tmpDir, err := os.MkdirTemp("", "hooks-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	// Create subdirectory (WriteToFile doesn't create directories)
	claudeDir := filepath.Join(tmpDir, ".claude")
	if err := os.MkdirAll(claudeDir, 0o755); err != nil {
		t.Fatalf("failed to create .claude dir: %v", err)
	}

	outputPath := filepath.Join(claudeDir, "settings.json")

	err = cfg.WriteToFile(outputPath, "claude")
	if err != nil {
		t.Fatalf("WriteToFile failed: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Error("expected output file to exist")
	}

	// Read and verify contents
	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	var parsed map[string]any
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}

	if _, ok := parsed["hooks"]; !ok {
		t.Error("expected 'hooks' key in output")
	}
}

func TestConfig_WriteToFile_UnsupportedFormat(t *testing.T) {
	cfg := &Config{
		Profile:  "default",
		AutoLint: true,
	}

	tmpDir, err := os.MkdirTemp("", "hooks-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	outputPath := filepath.Join(tmpDir, "output.json")

	err = cfg.WriteToFile(outputPath, "unsupported")
	if err == nil {
		t.Error("expected error for unsupported format")
	}
}

func TestDefaultPaths(t *testing.T) {
	paths := DefaultPaths()

	if len(paths) == 0 {
		t.Error("expected non-empty paths map")
	}

	// Check expected formats
	expectedFormats := []string{"claude", "cursor", "windsurf"}
	for _, format := range expectedFormats {
		if _, ok := paths[format]; !ok {
			t.Errorf("expected path for format %q", format)
		}
	}

	// Verify Claude path
	if paths["claude"] != filepath.Join(".claude", "settings.json") {
		t.Errorf("unexpected claude path: %s", paths["claude"])
	}
}

func TestSupportedFormats(t *testing.T) {
	formats := SupportedFormats()

	if len(formats) == 0 {
		t.Error("expected non-empty formats list")
	}

	// Check expected formats exist
	formatMap := make(map[string]bool)
	for _, f := range formats {
		formatMap[f] = true
	}

	expectedFormats := []string{"claude", "cursor", "windsurf"}
	for _, expected := range expectedFormats {
		if !formatMap[expected] {
			t.Errorf("expected format %q in supported formats", expected)
		}
	}
}

func TestEventSupport(t *testing.T) {
	support := EventSupport()

	if len(support) == 0 {
		t.Error("expected non-empty event support map")
	}

	// Each format should have some events
	for format, events := range support {
		if len(events) == 0 {
			t.Errorf("expected events for format %q", format)
		}
	}

	// Claude should have more events than others typically
	if len(support["claude"]) == 0 {
		t.Error("expected Claude to have events")
	}
}

func TestBuildLintCommand(t *testing.T) {
	cfg := &Config{
		Profile: "azure",
	}

	cmd := cfg.buildLintCommand()

	if cmd == "" {
		t.Error("expected non-empty command")
	}

	// Check for key components
	if !strings.Contains(cmd, "CLAUDE_FILE_PATH") {
		t.Error("expected CLAUDE_FILE_PATH in command")
	}

	if !strings.Contains(cmd, "mcp-api-style") {
		t.Error("expected mcp-api-style in command")
	}

	if !strings.Contains(cmd, "azure") {
		t.Error("expected profile name 'azure' in command")
	}

	if !strings.Contains(cmd, "openapi.yaml") {
		t.Error("expected openapi.yaml pattern in command")
	}
}

func TestBuildLintCommand_DifferentProfiles(t *testing.T) {
	profiles := []string{"default", "azure", "google", "zalando"}

	for _, profile := range profiles {
		t.Run(profile, func(t *testing.T) {
			cfg := &Config{Profile: profile}
			cmd := cfg.buildLintCommand()

			if !strings.Contains(cmd, profile) {
				t.Errorf("expected profile %q in command", profile)
			}
		})
	}
}

package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFile(t *testing.T) {
	// Create a temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, ".api-style.yaml")

	configContent := `
profile: azure
level: silver
exceptions:
  - rule: URI-001
    paths:
      - /legacy/**
    reason: Legacy API cannot be changed
severity-overrides:
  URI-002: hint
include:
  - "api/**/*.yaml"
exclude:
  - "**/generated/**"
`

	if err := os.WriteFile(configPath, []byte(configContent), 0o600); err != nil {
		t.Fatalf("Failed to write test config: %v", err)
	}

	cfg, err := LoadFile(configPath)
	if err != nil {
		t.Fatalf("LoadFile failed: %v", err)
	}

	if cfg.Profile != "azure" {
		t.Errorf("Expected profile 'azure', got %q", cfg.Profile)
	}
	if cfg.Level != "silver" {
		t.Errorf("Expected level 'silver', got %q", cfg.Level)
	}
	if len(cfg.Exceptions) != 1 {
		t.Errorf("Expected 1 exception, got %d", len(cfg.Exceptions))
	}
	if cfg.Exceptions[0].RuleID != "URI-001" {
		t.Errorf("Expected exception rule 'URI-001', got %q", cfg.Exceptions[0].RuleID)
	}
	if cfg.SeverityOverrides["URI-002"] != "hint" {
		t.Errorf("Expected severity override 'hint' for URI-002, got %q", cfg.SeverityOverrides["URI-002"])
	}
}

func TestLoadNoConfig(t *testing.T) {
	// Create a temp dir with no config
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	defer func() { _ = os.Chdir(origDir) }()
	_ = os.Chdir(tmpDir)

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if cfg != nil {
		t.Error("Expected nil config when no file exists")
	}
}

func TestLoadOrDefault(t *testing.T) {
	tmpDir := t.TempDir()
	origDir, _ := os.Getwd()
	defer func() { _ = os.Chdir(origDir) }()
	_ = os.Chdir(tmpDir)

	cfg, err := LoadOrDefault()
	if err != nil {
		t.Fatalf("LoadOrDefault failed: %v", err)
	}
	if cfg == nil {
		t.Fatal("Expected non-nil default config")
	}
	if cfg.Profile != "default" {
		t.Errorf("Expected default profile 'default', got %q", cfg.Profile)
	}
}

func TestConfigMerge(t *testing.T) {
	cfg := &Config{
		Profile: "default",
		Level:   "bronze",
	}

	cfg.Merge("azure", "gold", nil, nil)

	if cfg.Profile != "azure" {
		t.Errorf("Expected profile 'azure', got %q", cfg.Profile)
	}
	if cfg.Level != "gold" {
		t.Errorf("Expected level 'gold', got %q", cfg.Level)
	}
}

func TestConfigToExceptions(t *testing.T) {
	cfg := &Config{
		Exceptions: []ExceptionConfig{
			{
				RuleID: "URI-001",
				Paths:  []string{"/legacy/**"},
				Reason: "Legacy API",
			},
		},
	}

	exceptions := cfg.ToExceptions()
	if len(exceptions) != 1 {
		t.Fatalf("Expected 1 exception, got %d", len(exceptions))
	}
	if exceptions[0].RuleID != "URI-001" {
		t.Errorf("Expected rule 'URI-001', got %q", exceptions[0].RuleID)
	}
	if exceptions[0].Reason != "Legacy API" {
		t.Errorf("Expected reason 'Legacy API', got %q", exceptions[0].Reason)
	}
}

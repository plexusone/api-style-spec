package apistyle

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	skill := New()

	if skill == nil {
		t.Fatal("expected non-nil skill")
	}

	if skill.Name() != "apistyle" {
		t.Errorf("expected name 'apistyle', got %q", skill.Name())
	}
}

func TestNew_WithAnthropicAPIKey(t *testing.T) {
	skill := New(WithAnthropicAPIKey("test-key"))

	if skill.anthropicAPIKey != "test-key" {
		t.Error("expected API key to be set")
	}
}

func TestSkill_Name(t *testing.T) {
	skill := New()

	if skill.Name() != "apistyle" {
		t.Errorf("expected 'apistyle', got %q", skill.Name())
	}
}

func TestSkill_Description(t *testing.T) {
	skill := New()

	desc := skill.Description()
	if desc == "" {
		t.Error("expected non-empty description")
	}

	if !strings.Contains(desc, "OpenAPI") {
		t.Error("expected description to mention OpenAPI")
	}
}

func TestSkill_Init(t *testing.T) {
	skill := New()

	err := skill.Init(context.Background())
	if err != nil {
		t.Errorf("Init failed: %v", err)
	}
}

func TestSkill_Close(t *testing.T) {
	skill := New()

	err := skill.Close()
	if err != nil {
		t.Errorf("Close failed: %v", err)
	}
}

func TestSkill_Tools(t *testing.T) {
	skill := New()

	tools := skill.Tools()
	if len(tools) == 0 {
		t.Fatal("expected non-empty tools list")
	}

	// Check expected tools
	toolNames := make(map[string]bool)
	for _, tool := range tools {
		toolNames[tool.Name()] = true
	}

	expectedTools := []string{
		"lint",
		"evaluate",
		"analyze",
		"list_rules",
		"list_profiles",
		"explain_rule",
	}

	for _, expected := range expectedTools {
		if !toolNames[expected] {
			t.Errorf("expected tool %q", expected)
		}
	}
}

func TestSkill_MarshalJSON(t *testing.T) {
	skill := New()

	data, err := skill.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}

	var parsed map[string]any
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}

	if parsed["name"] != "apistyle" {
		t.Errorf("expected name 'apistyle', got %v", parsed["name"])
	}

	tools, ok := parsed["tools"].([]any)
	if !ok {
		t.Fatal("expected tools array")
	}

	if len(tools) != 6 {
		t.Errorf("expected 6 tools, got %d", len(tools))
	}
}

func TestLintTool(t *testing.T) {
	skill := New()
	_ = skill.Init(context.Background())
	defer func() { _ = skill.Close() }()

	// Find lint tool
	var lintTool interface {
		Call(context.Context, map[string]any) (any, error)
	}
	for _, tool := range skill.Tools() {
		if tool.Name() == "lint" {
			lintTool = tool
			break
		}
	}

	if lintTool == nil {
		t.Fatal("lint tool not found")
	}

	// Test with valid spec
	spec := `
openapi: "3.1.0"
info:
  title: Test API
  version: "1.0.0"
  description: A test API
paths:
  /users:
    get:
      operationId: listUsers
      summary: List users
      responses:
        "200":
          description: OK
`

	result, err := lintTool.Call(context.Background(), map[string]any{
		"openapi_spec": spec,
		"profile":      "default",
	})

	if err != nil {
		t.Fatalf("lint failed: %v", err)
	}

	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Fatal("expected map result")
	}

	if _, ok := resultMap["status"]; !ok {
		t.Error("expected 'status' in result")
	}

	if _, ok := resultMap["violations"]; !ok {
		t.Error("expected 'violations' in result")
	}
}

func TestLintTool_MissingSpec(t *testing.T) {
	skill := New()
	_ = skill.Init(context.Background())
	defer func() { _ = skill.Close() }()

	var lintTool interface {
		Call(context.Context, map[string]any) (any, error)
	}
	for _, tool := range skill.Tools() {
		if tool.Name() == "lint" {
			lintTool = tool
			break
		}
	}

	_, err := lintTool.Call(context.Background(), map[string]any{
		"profile": "default",
	})

	if err == nil {
		t.Error("expected error for missing spec")
	}

	if !strings.Contains(err.Error(), "openapi_spec") {
		t.Errorf("expected error to mention openapi_spec, got: %v", err)
	}
}

func TestLintTool_InvalidProfile(t *testing.T) {
	skill := New()
	_ = skill.Init(context.Background())
	defer func() { _ = skill.Close() }()

	var lintTool interface {
		Call(context.Context, map[string]any) (any, error)
	}
	for _, tool := range skill.Tools() {
		if tool.Name() == "lint" {
			lintTool = tool
			break
		}
	}

	spec := `openapi: "3.1.0"
info:
  title: Test
  version: "1.0.0"
paths: {}`

	_, err := lintTool.Call(context.Background(), map[string]any{
		"openapi_spec": spec,
		"profile":      "nonexistent-profile",
	})

	if err == nil {
		t.Error("expected error for invalid profile")
	}
}

func TestListProfilesTool(t *testing.T) {
	skill := New()
	_ = skill.Init(context.Background())
	defer func() { _ = skill.Close() }()

	var listProfilesTool interface {
		Call(context.Context, map[string]any) (any, error)
	}
	for _, tool := range skill.Tools() {
		if tool.Name() == "list_profiles" {
			listProfilesTool = tool
			break
		}
	}

	if listProfilesTool == nil {
		t.Fatal("list_profiles tool not found")
	}

	result, err := listProfilesTool.Call(context.Background(), map[string]any{})
	if err != nil {
		t.Fatalf("list_profiles failed: %v", err)
	}

	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Fatal("expected map result")
	}

	profiles, ok := resultMap["profiles"].([]map[string]any)
	if !ok {
		t.Fatal("expected profiles array")
	}

	if len(profiles) < 4 {
		t.Errorf("expected at least 4 profiles, got %d", len(profiles))
	}

	// Check for expected profiles
	profileNames := make(map[string]bool)
	for _, p := range profiles {
		if name, ok := p["name"].(string); ok {
			profileNames[name] = true
		}
	}

	expectedProfiles := []string{"default", "azure", "google", "zalando"}
	for _, expected := range expectedProfiles {
		if !profileNames[expected] {
			t.Errorf("expected profile %q", expected)
		}
	}
}

func TestListRulesTool(t *testing.T) {
	skill := New()
	_ = skill.Init(context.Background())
	defer func() { _ = skill.Close() }()

	var listRulesTool interface {
		Call(context.Context, map[string]any) (any, error)
	}
	for _, tool := range skill.Tools() {
		if tool.Name() == "list_rules" {
			listRulesTool = tool
			break
		}
	}

	if listRulesTool == nil {
		t.Fatal("list_rules tool not found")
	}

	result, err := listRulesTool.Call(context.Background(), map[string]any{
		"profile": "default",
	})

	if err != nil {
		t.Fatalf("list_rules failed: %v", err)
	}

	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Fatal("expected map result")
	}

	if resultMap["profile"] != "default" {
		t.Errorf("expected profile 'default', got %v", resultMap["profile"])
	}

	ruleCount, ok := resultMap["rule_count"].(int)
	if !ok {
		t.Fatal("expected rule_count")
	}

	if ruleCount == 0 {
		t.Error("expected non-zero rule count")
	}
}

func TestListRulesTool_WithFilters(t *testing.T) {
	skill := New()
	_ = skill.Init(context.Background())
	defer func() { _ = skill.Close() }()

	var listRulesTool interface {
		Call(context.Context, map[string]any) (any, error)
	}
	for _, tool := range skill.Tools() {
		if tool.Name() == "list_rules" {
			listRulesTool = tool
			break
		}
	}

	// Filter by category
	result, err := listRulesTool.Call(context.Background(), map[string]any{
		"profile":  "default",
		"category": "uri-design",
	})

	if err != nil {
		t.Fatalf("list_rules failed: %v", err)
	}

	resultMap := result.(map[string]any)
	rules := resultMap["rules"].([]map[string]any)

	for _, rule := range rules {
		if rule["category"] != "uri-design" {
			t.Errorf("expected category 'uri-design', got %v", rule["category"])
		}
	}
}

func TestExplainRuleTool(t *testing.T) {
	skill := New()
	_ = skill.Init(context.Background())
	defer func() { _ = skill.Close() }()

	var explainRuleTool interface {
		Call(context.Context, map[string]any) (any, error)
	}
	for _, tool := range skill.Tools() {
		if tool.Name() == "explain_rule" {
			explainRuleTool = tool
			break
		}
	}

	if explainRuleTool == nil {
		t.Fatal("explain_rule tool not found")
	}

	result, err := explainRuleTool.Call(context.Background(), map[string]any{
		"rule_id": "PO-001",
		"profile": "default",
	})

	if err != nil {
		t.Fatalf("explain_rule failed: %v", err)
	}

	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Fatal("expected map result")
	}

	if resultMap["id"] != "PO-001" {
		t.Errorf("expected id 'PO-001', got %v", resultMap["id"])
	}

	if _, ok := resultMap["title"]; !ok {
		t.Error("expected 'title' in result")
	}

	if _, ok := resultMap["rationale"]; !ok {
		t.Error("expected 'rationale' in result")
	}
}

func TestExplainRuleTool_NotFound(t *testing.T) {
	skill := New()
	_ = skill.Init(context.Background())
	defer func() { _ = skill.Close() }()

	var explainRuleTool interface {
		Call(context.Context, map[string]any) (any, error)
	}
	for _, tool := range skill.Tools() {
		if tool.Name() == "explain_rule" {
			explainRuleTool = tool
			break
		}
	}

	_, err := explainRuleTool.Call(context.Background(), map[string]any{
		"rule_id": "NONEXISTENT-999",
		"profile": "default",
	})

	if err == nil {
		t.Error("expected error for nonexistent rule")
	}

	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("expected 'not found' in error, got: %v", err)
	}
}

func TestExplainRuleTool_MissingRuleID(t *testing.T) {
	skill := New()
	_ = skill.Init(context.Background())
	defer func() { _ = skill.Close() }()

	var explainRuleTool interface {
		Call(context.Context, map[string]any) (any, error)
	}
	for _, tool := range skill.Tools() {
		if tool.Name() == "explain_rule" {
			explainRuleTool = tool
			break
		}
	}

	_, err := explainRuleTool.Call(context.Background(), map[string]any{
		"profile": "default",
	})

	if err == nil {
		t.Error("expected error for missing rule_id")
	}
}

func TestEvaluateTool_NoAPIKey(t *testing.T) {
	skill := New() // No API key
	_ = skill.Init(context.Background())
	defer func() { _ = skill.Close() }()

	var evaluateTool interface {
		Call(context.Context, map[string]any) (any, error)
	}
	for _, tool := range skill.Tools() {
		if tool.Name() == "evaluate" {
			evaluateTool = tool
			break
		}
	}

	spec := `openapi: "3.1.0"
info:
  title: Test
  version: "1.0.0"
paths: {}`

	_, err := evaluateTool.Call(context.Background(), map[string]any{
		"openapi_spec": spec,
	})

	if err == nil {
		t.Error("expected error when API key not configured")
	}

	if !strings.Contains(err.Error(), "ANTHROPIC_API_KEY") {
		t.Errorf("expected error to mention API key, got: %v", err)
	}
}

func TestAnalyzeTool_LintOnly(t *testing.T) {
	skill := New()
	_ = skill.Init(context.Background())
	defer func() { _ = skill.Close() }()

	var analyzeTool interface {
		Call(context.Context, map[string]any) (any, error)
	}
	for _, tool := range skill.Tools() {
		if tool.Name() == "analyze" {
			analyzeTool = tool
			break
		}
	}

	spec := `
openapi: "3.1.0"
info:
  title: Test API
  version: "1.0.0"
  description: A test API
paths:
  /users:
    get:
      operationId: listUsers
      summary: List users
      responses:
        "200":
          description: OK
`

	result, err := analyzeTool.Call(context.Background(), map[string]any{
		"openapi_spec": spec,
		"lint_only":    true,
	})

	if err != nil {
		t.Fatalf("analyze failed: %v", err)
	}

	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Fatal("expected map result")
	}

	if _, ok := resultMap["decision"]; !ok {
		t.Error("expected 'decision' in result")
	}

	if _, ok := resultMap["lint"]; !ok {
		t.Error("expected 'lint' in result")
	}
}

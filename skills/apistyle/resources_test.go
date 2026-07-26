package apistyle

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newReadResourceRequest creates a test ReadResourceRequest with the given URI.
func newReadResourceRequest(uri string) *mcp.ReadResourceRequest {
	return &mcp.ReadResourceRequest{
		Params: &mcp.ReadResourceParams{
			URI: uri,
		},
	}
}

func TestHandleProfileList(t *testing.T) {
	req := newReadResourceRequest("apistyle://profiles")

	result, err := handleProfileList(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.Contents, 1)

	content := result.Contents[0]
	assert.Equal(t, "apistyle://profiles", content.URI)
	assert.Equal(t, "application/json", content.MIMEType)
	assert.NotEmpty(t, content.Text)

	// Parse and verify structure
	var data map[string]any
	err = json.Unmarshal([]byte(content.Text), &data)
	require.NoError(t, err)

	profiles, ok := data["profiles"].([]any)
	require.True(t, ok, "should have profiles array")
	assert.Greater(t, len(profiles), 0, "should have at least one profile")

	// Check first profile structure
	firstProfile, ok := profiles[0].(map[string]any)
	require.True(t, ok)
	assert.NotEmpty(t, firstProfile["name"])
	assert.NotEmpty(t, firstProfile["version"])
}

func TestHandleExemplarList(t *testing.T) {
	req := newReadResourceRequest("apistyle://exemplars")

	result, err := handleExemplarList(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.Contents, 1)

	content := result.Contents[0]
	assert.Equal(t, "apistyle://exemplars", content.URI)
	assert.Equal(t, "application/json", content.MIMEType)

	// Parse and verify structure
	var data map[string]any
	err = json.Unmarshal([]byte(content.Text), &data)
	require.NoError(t, err)

	exemplars, ok := data["exemplars"].([]any)
	require.True(t, ok, "should have exemplars array")
	assert.Greater(t, len(exemplars), 0, "should have at least one exemplar")
}

func TestHandleProfile(t *testing.T) {
	req := newReadResourceRequest("apistyle://profile/default")

	result, err := handleProfile(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.Contents, 1)

	content := result.Contents[0]
	assert.Equal(t, "apistyle://profile/default", content.URI)
	assert.Equal(t, "application/json", content.MIMEType)

	// Parse and verify it's a valid profile
	var profile map[string]any
	err = json.Unmarshal([]byte(content.Text), &profile)
	require.NoError(t, err)

	assert.Equal(t, "default", profile["name"])
	assert.NotEmpty(t, profile["version"])
	assert.NotNil(t, profile["rules"])
}

func TestHandleProfile_NotFound(t *testing.T) {
	req := newReadResourceRequest("apistyle://profile/nonexistent")

	_, err := handleProfile(context.Background(), req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nonexistent")
}

func TestHandleExemplar(t *testing.T) {
	req := newReadResourceRequest("apistyle://exemplar/default-minimal")

	result, err := handleExemplar(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.Contents, 1)

	content := result.Contents[0]
	assert.Equal(t, "apistyle://exemplar/default-minimal", content.URI)
	assert.Equal(t, "application/x-yaml", content.MIMEType)

	// Verify it's valid OpenAPI YAML
	assert.True(t, strings.Contains(content.Text, "openapi:"))
	assert.True(t, strings.Contains(content.Text, "paths:"))
}

func TestHandleExemplar_NotFound(t *testing.T) {
	req := newReadResourceRequest("apistyle://exemplar/nonexistent")

	_, err := handleExemplar(context.Background(), req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nonexistent")
}

func TestHandlePatternList(t *testing.T) {
	req := newReadResourceRequest("apistyle://patterns/default")

	result, err := handlePatternList(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.Contents, 1)

	content := result.Contents[0]
	assert.Equal(t, "application/json", content.MIMEType)

	// Parse and verify structure
	var data map[string]any
	err = json.Unmarshal([]byte(content.Text), &data)
	require.NoError(t, err)

	assert.Equal(t, "default", data["profile"])
	patterns, ok := data["patterns"].([]any)
	require.True(t, ok, "should have patterns array")
	assert.Greater(t, len(patterns), 0, "should have at least one pattern")
}

func TestHandlePattern(t *testing.T) {
	req := newReadResourceRequest("apistyle://pattern/default/cursor-pagination")

	result, err := handlePattern(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.Contents, 1)

	content := result.Contents[0]
	assert.Equal(t, "application/json", content.MIMEType)

	// Parse and verify structure
	var pattern map[string]any
	err = json.Unmarshal([]byte(content.Text), &pattern)
	require.NoError(t, err)

	assert.Equal(t, "cursor-pagination", pattern["id"])
	assert.NotEmpty(t, pattern["name"])
	assert.NotEmpty(t, pattern["summary"])
}

func TestHandlePattern_NotFound(t *testing.T) {
	req := newReadResourceRequest("apistyle://pattern/default/nonexistent")

	_, err := handlePattern(context.Background(), req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nonexistent")
}

func TestHandleRubric_Evaluation(t *testing.T) {
	req := newReadResourceRequest("apistyle://rubric/default/evaluation")

	result, err := handleRubric(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.Contents, 1)

	content := result.Contents[0]
	assert.Equal(t, "application/json", content.MIMEType)

	// Verify it's valid JSON
	var rubric map[string]any
	err = json.Unmarshal([]byte(content.Text), &rubric)
	require.NoError(t, err)
	assert.NotEmpty(t, rubric)
}

func TestHandleRubric_Generation(t *testing.T) {
	req := newReadResourceRequest("apistyle://rubric/default/generation")

	result, err := handleRubric(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Len(t, result.Contents, 1)

	content := result.Contents[0]
	assert.Equal(t, "application/json", content.MIMEType)

	// Verify it's valid JSON with expected structure
	var rubric map[string]any
	err = json.Unmarshal([]byte(content.Text), &rubric)
	require.NoError(t, err)

	assert.Equal(t, "default", rubric["name"])
	assert.NotNil(t, rubric["phases"])
}

func TestHandleRubric_InvalidMode(t *testing.T) {
	req := newReadResourceRequest("apistyle://rubric/default/invalid")

	_, err := handleRubric(context.Background(), req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid")
}

func TestExtractURIParam(t *testing.T) {
	tests := []struct {
		uri      string
		prefix   string
		expected string
	}{
		{"apistyle://profile/default", "apistyle://profile/", "default"},
		{"apistyle://exemplar/default-minimal", "apistyle://exemplar/", "default-minimal"},
		{"apistyle://pattern/default/cursor-pagination", "apistyle://pattern/", "default/cursor-pagination"},
		{"other://profile/default", "apistyle://profile/", ""},
	}

	for _, tt := range tests {
		t.Run(tt.uri, func(t *testing.T) {
			result := extractURIParam(tt.uri, tt.prefix)
			assert.Equal(t, tt.expected, result)
		})
	}
}

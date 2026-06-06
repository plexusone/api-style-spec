package judge

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	anthropicAPIURL     = "https://api.anthropic.com/v1/messages"
	anthropicAPIVersion = "2023-06-01"
)

// AnthropicProvider implements Provider using the Anthropic Claude HTTP API.
// Uses HTTP directly to avoid SDK dependency conflicts.
type AnthropicProvider struct {
	apiKey       string
	httpClient   *http.Client
	defaultModel string
	baseURL      string
}

// NewAnthropicProvider creates a new Anthropic provider.
// If apiKey is empty, it reads from ANTHROPIC_API_KEY environment variable.
func NewAnthropicProvider(apiKey string, config *ProviderConfig) *AnthropicProvider {
	if apiKey == "" {
		apiKey = os.Getenv("ANTHROPIC_API_KEY")
	}

	if config == nil {
		config = DefaultProviderConfig()
	}

	timeout := time.Duration(config.Timeout) * time.Second
	if timeout == 0 {
		timeout = 60 * time.Second
	}

	defaultModel := config.DefaultModel
	if defaultModel == "" {
		defaultModel = ModelClaudeHaiku
	}

	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = anthropicAPIURL
	}

	return &AnthropicProvider{
		apiKey:       apiKey,
		httpClient:   &http.Client{Timeout: timeout},
		defaultModel: defaultModel,
		baseURL:      baseURL,
	}
}

// Model constants for Anthropic Claude models.
const (
	ModelClaudeOpus   = "claude-opus-4-20250514"
	ModelClaudeSonnet = "claude-sonnet-4-20250514"
	ModelClaudeHaiku  = "claude-3-5-haiku-20241022"
)

// anthropicRequest represents the request body for the Messages API.
type anthropicRequest struct {
	Model       string             `json:"model"`
	MaxTokens   int                `json:"max_tokens"`
	Messages    []anthropicMessage `json:"messages"`
	System      string             `json:"system,omitempty"`
	Temperature *float64           `json:"temperature,omitempty"`
	StopSeqs    []string           `json:"stop_sequences,omitempty"`
}

type anthropicMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// anthropicResponse represents the response from the Messages API.
type anthropicResponse struct {
	ID           string             `json:"id"`
	Type         string             `json:"type"`
	Role         string             `json:"role"`
	Content      []anthropicContent `json:"content"`
	Model        string             `json:"model"`
	StopReason   string             `json:"stop_reason"`
	StopSequence string             `json:"stop_sequence,omitempty"`
	Usage        anthropicUsage     `json:"usage"`
	Error        *anthropicError    `json:"error,omitempty"`
}

type anthropicContent struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

type anthropicUsage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

type anthropicError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Complete sends a completion request to the Claude API.
func (p *AnthropicProvider) Complete(ctx context.Context, req *CompletionRequest) (*CompletionResponse, error) {
	model := req.Model
	if model == "" {
		model = p.defaultModel
	}

	maxTokens := req.MaxTokens
	if maxTokens == 0 {
		maxTokens = 2048
	}

	apiReq := anthropicRequest{
		Model:     model,
		MaxTokens: maxTokens,
		Messages: []anthropicMessage{
			{Role: "user", Content: req.UserPrompt},
		},
		System:   req.SystemPrompt,
		StopSeqs: req.StopSequences,
	}

	if req.Temperature > 0 {
		apiReq.Temperature = &req.Temperature
	}

	body, err := json.Marshal(apiReq)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", p.baseURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-api-key", p.apiKey)
	httpReq.Header.Set("anthropic-version", anthropicAPIVersion)

	resp, err := p.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp anthropicResponse
		if json.Unmarshal(respBody, &errResp) == nil && errResp.Error != nil {
			return nil, fmt.Errorf("API error (%d): %s - %s",
				resp.StatusCode, errResp.Error.Type, errResp.Error.Message)
		}
		return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, string(respBody))
	}

	var apiResp anthropicResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	// Extract text content
	var content string
	for _, block := range apiResp.Content {
		if block.Type == "text" {
			content += block.Text
		}
	}

	return &CompletionResponse{
		Content: content,
		Model:   apiResp.Model,
		Usage: &TokenUsage{
			InputTokens:  apiResp.Usage.InputTokens,
			OutputTokens: apiResp.Usage.OutputTokens,
			TotalTokens:  apiResp.Usage.InputTokens + apiResp.Usage.OutputTokens,
		},
		FinishReason: apiResp.StopReason,
	}, nil
}

// Name returns the provider name.
func (p *AnthropicProvider) Name() string {
	return "anthropic"
}

// DefaultModel returns the configured default model.
func (p *AnthropicProvider) DefaultModel() string {
	return p.defaultModel
}

// SetDefaultModel updates the default model.
func (p *AnthropicProvider) SetDefaultModel(model string) {
	p.defaultModel = model
}

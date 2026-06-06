package judge

import (
	"context"
)

// Provider defines the interface for LLM providers.
type Provider interface {
	// Complete sends a completion request to the LLM.
	Complete(ctx context.Context, req *CompletionRequest) (*CompletionResponse, error)

	// Name returns the provider name (e.g., "anthropic", "openai").
	Name() string

	// DefaultModel returns the provider's default model.
	DefaultModel() string
}

// CompletionRequest contains parameters for an LLM completion.
type CompletionRequest struct {
	// SystemPrompt sets the system/context instruction.
	SystemPrompt string

	// UserPrompt is the main user message/question.
	UserPrompt string

	// Model specifies the model to use (provider-specific).
	Model string

	// Temperature controls randomness (0.0-1.0).
	Temperature float64

	// MaxTokens limits response length.
	MaxTokens int

	// StopSequences are strings that stop generation.
	StopSequences []string

	// Metadata is optional context for logging/tracing.
	Metadata map[string]string
}

// CompletionResponse contains the LLM's response.
type CompletionResponse struct {
	// Content is the generated text.
	Content string

	// Model is the model that was used.
	Model string

	// Usage contains token usage statistics.
	Usage *TokenUsage

	// FinishReason indicates why generation stopped.
	FinishReason string
}

// TokenUsage tracks token consumption.
type TokenUsage struct {
	// InputTokens is the number of tokens in the request.
	InputTokens int

	// OutputTokens is the number of tokens in the response.
	OutputTokens int

	// TotalTokens is the sum of input and output tokens.
	TotalTokens int
}

// ProviderConfig holds common provider configuration.
type ProviderConfig struct {
	// APIKey is the authentication key.
	APIKey string

	// BaseURL overrides the default API endpoint.
	BaseURL string

	// DefaultModel sets the default model for requests.
	DefaultModel string

	// Timeout in seconds for API calls.
	Timeout int

	// MaxRetries for failed requests.
	MaxRetries int

	// RetryDelay in milliseconds between retries.
	RetryDelay int
}

// DefaultProviderConfig returns a config with sensible defaults.
func DefaultProviderConfig() *ProviderConfig {
	return &ProviderConfig{
		Timeout:    60,
		MaxRetries: 3,
		RetryDelay: 1000,
	}
}

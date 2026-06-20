package types

// DetailedExample provides a rich example with annotations and context.
type DetailedExample struct {
	// Title is a brief description of the example.
	Title string `json:"title"`

	// Description provides additional context for the example.
	Description string `json:"description,omitempty"`

	// Type indicates whether this is a good, bad, or context example.
	Type string `json:"type"` // "good", "bad", "context"

	// Language specifies the code language (e.g., "openapi", "json", "http").
	Language string `json:"language,omitempty"`

	// Code is the example code or content.
	Code string `json:"code"`

	// Annotations highlight specific parts of the code.
	Annotations []CodeAnnotation `json:"annotations,omitempty"`

	// Before shows the state before migration (for migration examples).
	Before string `json:"before,omitempty"`

	// After shows the state after migration (for migration examples).
	After string `json:"after,omitempty"`
}

// CodeAnnotation highlights a specific part of code in an example.
type CodeAnnotation struct {
	// Line is the starting line number (1-indexed).
	Line int `json:"line"`

	// EndLine is the ending line number (optional, for multi-line annotations).
	EndLine int `json:"endLine,omitempty"`

	// Text is the annotation message.
	Text string `json:"text"`

	// Type indicates the annotation severity or purpose.
	Type string `json:"type,omitempty"` // "info", "warning", "error"
}

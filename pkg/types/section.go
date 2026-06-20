package types

// Section represents a document section for navigation and organization.
type Section struct {
	// ID is a unique identifier for the section.
	ID string `json:"id"`

	// Title is the section heading.
	Title string `json:"title"`

	// Description provides a brief summary of the section.
	Description string `json:"description,omitempty"`

	// Order determines display ordering (lower = first).
	Order int `json:"order,omitempty"`

	// ParentID references a parent section for hierarchy.
	ParentID string `json:"parentId,omitempty"`

	// Rules lists rule IDs contained in this section.
	Rules []string `json:"rules,omitempty"`

	// Patterns lists pattern IDs contained in this section.
	Patterns []string `json:"patterns,omitempty"`

	// Introduction is introductory content for the section (Markdown).
	Introduction string `json:"introduction,omitempty"`

	// Content is the main section content (Markdown).
	Content string `json:"content,omitempty"`
}

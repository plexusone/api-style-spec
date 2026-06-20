package types

// GlossaryTerm defines a term in the API style glossary.
type GlossaryTerm struct {
	// Term is the word or phrase being defined.
	Term string `json:"term"`

	// Definition explains the term.
	Definition string `json:"definition"`

	// Aliases are alternative names for the term.
	Aliases []string `json:"aliases,omitempty"`
}

// Principle defines a high-level design principle.
type Principle struct {
	// ID is a unique identifier for the principle.
	ID string `json:"id"`

	// Title is the principle name.
	Title string `json:"title"`

	// Description explains the principle in detail (Markdown).
	Description string `json:"description"`

	// RelatedRules lists rule IDs that implement this principle.
	RelatedRules []string `json:"relatedRules,omitempty"`
}

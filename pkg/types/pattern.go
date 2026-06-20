package types

// Pattern defines a reusable API design pattern.
type Pattern struct {
	// ID is a unique identifier for the pattern.
	ID string `json:"id"`

	// Name is the human-readable name.
	Name string `json:"name"`

	// Category groups related patterns (e.g., "collections", "errors", "versioning").
	Category string `json:"category,omitempty"`

	// Summary is a brief one-line description.
	Summary string `json:"summary"`

	// Description provides extended prose explanation (Markdown).
	Description string `json:"description,omitempty"`

	// Problem describes what issue this pattern addresses.
	Problem string `json:"problem,omitempty"`

	// Solution describes how the pattern solves the problem.
	Solution string `json:"solution,omitempty"`

	// When describes when to use this pattern.
	When string `json:"when,omitempty"`

	// Examples provides detailed usage examples.
	Examples []DetailedExample `json:"examples,omitempty"`

	// RelatedRules lists rule IDs that implement or relate to this pattern.
	RelatedRules []string `json:"relatedRules,omitempty"`

	// RelatedPatterns lists other pattern IDs that work with this one.
	RelatedPatterns []string `json:"relatedPatterns,omitempty"`

	// References links to external documentation.
	References []Reference `json:"references,omitempty"`

	// Diagrams provides visual representations of the pattern.
	Diagrams []Diagram `json:"diagrams,omitempty"`
}

// Diagram provides visual representation of a concept.
type Diagram struct {
	// Title is the diagram title.
	Title string `json:"title"`

	// Type specifies the diagram format.
	Type string `json:"type"` // "mermaid", "plantuml", "url"

	// Content is the diagram content (code for mermaid/plantuml, URL for url type).
	Content string `json:"content"`

	// Alt is alternative text for accessibility.
	Alt string `json:"alt,omitempty"`
}

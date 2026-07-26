package types

// CategoryGroup defines a semantic grouping of related categories.
type CategoryGroup struct {
	// ID is the group identifier (e.g., "resource-design").
	ID string `json:"id"`
	// Title is the display name.
	Title string `json:"title"`
	// Description explains what this group covers.
	Description string `json:"description,omitempty"`
	// Categories lists category IDs belonging to this group, in display order.
	Categories []string `json:"categories"`
	// Order determines display ordering (lower = first).
	Order int `json:"order,omitempty"`
}

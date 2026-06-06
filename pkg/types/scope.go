package types

// Scope defines what part of an OpenAPI specification a rule applies to.
type Scope string

const (
	// ScopePath applies to path definitions.
	ScopePath Scope = "path"
	// ScopeOperation applies to individual operations (GET, POST, etc.).
	ScopeOperation Scope = "operation"
	// ScopeParameter applies to parameters.
	ScopeParameter Scope = "parameter"
	// ScopeSchema applies to schema definitions.
	ScopeSchema Scope = "schema"
	// ScopeResponse applies to response definitions.
	ScopeResponse Scope = "response"
	// ScopeInfo applies to the info section.
	ScopeInfo Scope = "info"
	// ScopeSecurity applies to security definitions.
	ScopeSecurity Scope = "security"
	// ScopeGlobal applies to the entire specification.
	ScopeGlobal Scope = "global"
)

// String returns the string representation.
func (s Scope) String() string {
	return string(s)
}

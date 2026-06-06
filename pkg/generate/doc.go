// Package generate provides output generation from APIStyleSpec.
//
// This package can generate various outputs from a style specification:
//
//   - Markdown documentation (human-readable style guide)
//   - Spectral YAML (linting ruleset for Spectral/vacuum)
//
// Usage:
//
//	spec, _ := profile.Load("azure")
//	markdown, err := generate.Markdown(spec, nil)
//	spectral, err := generate.Spectral(spec, nil)
//
// The generated outputs can be written to files or served via API.
package generate

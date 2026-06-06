package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/plexusone/api-style-spec/pkg/generate"
	"github.com/plexusone/api-style-spec/pkg/profile"
)

var (
	generateOutput  string
	generateProfile string
)

var generateCmd = &cobra.Command{
	Use:   "generate <type>",
	Short: "Generate outputs from a style profile",
	Long: `Generate various outputs from an API style specification profile.

Supported types:
  guide    - Generate Markdown documentation (style guide)
  spectral - Generate Spectral YAML ruleset

Examples:
  api-style generate guide --profile azure
  api-style generate guide --profile default --output style-guide.md
  api-style generate spectral --profile zalando --output .spectral.yaml`,
}

var generateGuideCmd = &cobra.Command{
	Use:   "guide",
	Short: "Generate Markdown style guide",
	Long: `Generate a Markdown documentation file from a style profile.

The generated guide includes:
- Table of contents
- Rule categories and descriptions
- Good and bad examples
- Conformance level requirements

Examples:
  api-style generate guide
  api-style generate guide --profile azure
  api-style generate guide --profile google --output api-guidelines.md`,
	RunE: runGenerateGuide,
}

var generateSpectralCmd = &cobra.Command{
	Use:   "spectral",
	Short: "Generate Spectral ruleset",
	Long: `Generate a Spectral-compatible YAML ruleset from a style profile.

The generated ruleset can be used with Spectral or vacuum for linting.
Only rules with Spectral enforcement configuration are included.

Examples:
  api-style generate spectral
  api-style generate spectral --profile azure
  api-style generate spectral --profile default --output .spectral.yaml`,
	RunE: runGenerateSpectral,
}

func init() {
	generateCmd.PersistentFlags().StringVarP(&generateOutput, "output", "o", "", "Output file (default: stdout)")
	generateCmd.PersistentFlags().StringVarP(&generateProfile, "profile", "p", "default", "Style profile to use")

	generateCmd.AddCommand(generateGuideCmd)
	generateCmd.AddCommand(generateSpectralCmd)
}

func runGenerateGuide(_ *cobra.Command, _ []string) error {
	// Load profile
	spec, err := profile.Load(generateProfile)
	if err != nil {
		return fmt.Errorf("loading profile %q: %w", generateProfile, err)
	}

	// Generate Markdown
	opts := generate.DefaultMarkdownOptions()
	md, err := generate.Markdown(spec, opts)
	if err != nil {
		return fmt.Errorf("generating markdown: %w", err)
	}

	// Write output
	return writeOutput(md)
}

func runGenerateSpectral(_ *cobra.Command, _ []string) error {
	// Load profile
	spec, err := profile.Load(generateProfile)
	if err != nil {
		return fmt.Errorf("loading profile %q: %w", generateProfile, err)
	}

	// Generate Spectral YAML
	opts := generate.DefaultSpectralOptions()
	yaml, err := generate.Spectral(spec, opts)
	if err != nil {
		return fmt.Errorf("generating spectral: %w", err)
	}

	// Write output
	return writeOutput(yaml)
}

func writeOutput(content string) error {
	// Ensure trailing newline
	if !strings.HasSuffix(content, "\n") {
		content += "\n"
	}

	if generateOutput != "" {
		if err := os.WriteFile(generateOutput, []byte(content), 0o644); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
		fmt.Printf("Generated: %s\n", generateOutput)
	} else {
		fmt.Print(content)
	}

	return nil
}

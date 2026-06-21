package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/plexusone/api-style-spec/pkg/generate"
	"github.com/plexusone/api-style-spec/pkg/judge"
	"github.com/plexusone/api-style-spec/pkg/profile"
	"github.com/plexusone/api-style-spec/pkg/types"
)

var (
	generateOutput  string
	generateProfile string
	generateFile    string

	// MkDocs-specific flags
	mkdocsSiteName        string
	mkdocsSiteURL         string
	mkdocsRepoURL         string
	mkdocsTheme           string
	mkdocsSplitPatterns   bool
	mkdocsNoSplitCategory bool
	mkdocsNoSearch        bool
)

var generateCmd = &cobra.Command{
	Use:   "generate <type>",
	Short: "Generate outputs from a style profile",
	Long: `Generate various outputs from an API style specification profile.

Supported types:
  guide    - Generate single-page Markdown documentation (style guide)
  mkdocs   - Generate MkDocs multi-page documentation site
  spectral - Generate Spectral YAML ruleset
  rubric   - Generate structured-evaluation rubric for LLM-as-Judge

Examples:
  api-style generate guide --profile azure
  api-style generate guide --file examples/zalando-rest.api-style.json --output zalando.md
  api-style generate mkdocs --profile zalando --output ./docs
  api-style generate spectral --profile zalando --output .spectral.yaml
  api-style generate rubric --profile zalando --output zalando.rubric.json`,
}

var generateGuideCmd = &cobra.Command{
	Use:   "guide",
	Short: "Generate single-page Markdown style guide",
	Long: `Generate a Markdown documentation file from a style profile.

The generated guide includes:
- Table of contents
- Introduction and design principles
- Design patterns with examples
- Rule categories and descriptions
- Good and bad examples
- Conformance level requirements
- Glossary

Examples:
  api-style generate guide
  api-style generate guide --profile azure
  api-style generate guide --file examples/zalando-rest.api-style.json --output zalando.md`,
	RunE: runGenerateGuide,
}

var generateMkDocsCmd = &cobra.Command{
	Use:   "mkdocs",
	Short: "Generate MkDocs multi-page documentation site",
	Long: `Generate a complete MkDocs documentation site from a style profile.

The generated site includes:
- mkdocs.yml configuration
- index.md with overview
- Separate pages for principles, patterns, rules (by category), glossary
- Material theme with search, dark mode, and code highlighting

Examples:
  api-style generate mkdocs --profile zalando --output ./docs
  api-style generate mkdocs --file examples/azure.api-style.json --output ./azure-docs
  api-style generate mkdocs --profile azure --site-name "My API Guidelines"`,
	RunE: runGenerateMkDocs,
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

var generateRubricCmd = &cobra.Command{
	Use:   "rubric",
	Short: "Generate structured-evaluation rubric",
	Long: `Generate a structured-evaluation rubric from a style profile.

The generated rubric can be used with LLM-as-Judge for semantic
evaluation of OpenAPI specifications against the style guide.

The rubric includes:
- Categories derived from style profile categories
- Pass/partial/fail criteria from rule judge configurations
- Few-shot examples for calibration
- Evaluation prompts for each category

Examples:
  api-style generate rubric
  api-style generate rubric --profile azure
  api-style generate rubric --file examples/zalando-rest.api-style.json --output zalando.rubric.json`,
	RunE: runGenerateRubric,
}

func init() {
	generateCmd.PersistentFlags().StringVarP(&generateOutput, "output", "o", "", "Output file/directory (default: stdout for guide/spectral, ./docs for mkdocs)")
	generateCmd.PersistentFlags().StringVarP(&generateProfile, "profile", "p", "", "Style profile name to use (built-in or from search paths)")
	generateCmd.PersistentFlags().StringVarP(&generateFile, "file", "f", "", "Path to style profile JSON/YAML file")

	// MkDocs-specific flags
	generateMkDocsCmd.Flags().StringVar(&mkdocsSiteName, "site-name", "", "MkDocs site name (default: from profile)")
	generateMkDocsCmd.Flags().StringVar(&mkdocsSiteURL, "site-url", "", "MkDocs site URL")
	generateMkDocsCmd.Flags().StringVar(&mkdocsRepoURL, "repo-url", "", "Repository URL (default: from profile)")
	generateMkDocsCmd.Flags().StringVar(&mkdocsTheme, "theme", "material", "MkDocs theme")
	generateMkDocsCmd.Flags().BoolVar(&mkdocsSplitPatterns, "split-patterns", false, "Create separate pages per pattern")
	generateMkDocsCmd.Flags().BoolVar(&mkdocsNoSplitCategory, "no-split-categories", false, "Keep all rules in one page")
	generateMkDocsCmd.Flags().BoolVar(&mkdocsNoSearch, "no-search", false, "Disable search plugin")

	generateCmd.AddCommand(generateGuideCmd)
	generateCmd.AddCommand(generateMkDocsCmd)
	generateCmd.AddCommand(generateSpectralCmd)
	generateCmd.AddCommand(generateRubricCmd)
}

func loadProfile() (*types.APIStyleSpec, error) {
	// Prefer --file over --profile
	if generateFile != "" {
		spec, err := profile.LoadFile(generateFile)
		if err != nil {
			return nil, fmt.Errorf("loading file %q: %w", generateFile, err)
		}
		return spec, nil
	}

	// Use profile name (default to "default" if neither specified)
	profileName := generateProfile
	if profileName == "" {
		profileName = "default"
	}

	spec, err := profile.Load(profileName)
	if err != nil {
		return nil, fmt.Errorf("loading profile %q: %w", profileName, err)
	}
	return spec, nil
}

func runGenerateGuide(_ *cobra.Command, _ []string) error {
	// Load profile
	spec, err := loadProfile()
	if err != nil {
		return err
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

func runGenerateMkDocs(_ *cobra.Command, _ []string) error {
	// Load profile
	spec, err := loadProfile()
	if err != nil {
		return err
	}

	// Configure options
	opts := generate.DefaultMkDocsOptions()
	if mkdocsSiteName != "" {
		opts.SiteName = mkdocsSiteName
	}
	if mkdocsSiteURL != "" {
		opts.SiteURL = mkdocsSiteURL
	}
	if mkdocsRepoURL != "" {
		opts.RepoURL = mkdocsRepoURL
	}
	if mkdocsTheme != "" {
		opts.Theme = mkdocsTheme
	}
	opts.SplitPatterns = mkdocsSplitPatterns
	opts.SplitCategories = !mkdocsNoSplitCategory
	opts.IncludeSearch = !mkdocsNoSearch

	// Generate MkDocs site
	result, err := generate.MkDocs(spec, opts)
	if err != nil {
		return fmt.Errorf("generating mkdocs: %w", err)
	}

	// Determine output directory
	outputDir := generateOutput
	if outputDir == "" {
		outputDir = "./docs"
	}

	// Write to filesystem
	if err := generate.WriteMkDocs(result, outputDir); err != nil {
		return fmt.Errorf("writing mkdocs: %w", err)
	}

	fmt.Printf("Generated MkDocs site: %s\n", outputDir)
	fmt.Printf("  - %d pages created\n", len(result.Pages))
	fmt.Println("\nTo build and serve:")
	fmt.Printf("  cd %s && pip install mkdocs-material && mkdocs serve\n", outputDir)

	return nil
}

func runGenerateSpectral(_ *cobra.Command, _ []string) error {
	// Load profile
	spec, err := loadProfile()
	if err != nil {
		return err
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

func runGenerateRubric(_ *cobra.Command, _ []string) error {
	// Load profile
	spec, err := loadProfile()
	if err != nil {
		return err
	}

	// Generate structured-evaluation rubric
	rubricSet := judge.GenerateRubricSet(spec)

	// Serialize to JSON
	jsonBytes, err := judge.RubricSetToJSON(rubricSet)
	if err != nil {
		return fmt.Errorf("serializing rubric: %w", err)
	}

	// Write output
	return writeOutput(string(jsonBytes))
}

func writeOutput(content string) error {
	// Ensure trailing newline
	if !strings.HasSuffix(content, "\n") {
		content += "\n"
	}

	if generateOutput != "" {
		if err := os.WriteFile(generateOutput, []byte(content), 0o600); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
		fmt.Printf("Generated: %s\n", generateOutput)
	} else {
		fmt.Print(content)
	}

	return nil
}

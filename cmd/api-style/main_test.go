package main

import "testing"

func TestRootCommandRegistration(t *testing.T) {
	want := []string{
		"lint", "evaluate", "analyze", "generate", "exemplar", "pattern",
		"hooks", "serve", "score-profile", "suggest-fixes", "version",
	}

	registered := make(map[string]bool)
	for _, c := range rootCmd.Commands() {
		registered[c.Name()] = true
	}

	for _, name := range want {
		if !registered[name] {
			t.Errorf("command %q not registered on root", name)
		}
	}
}

func TestGenerateSubcommands(t *testing.T) {
	want := []string{
		"guide", "guide-html", "mkdocs", "spectral", "rubric",
		"report", "gap-analysis",
	}

	registered := make(map[string]bool)
	for _, c := range generateCmd.Commands() {
		registered[c.Name()] = true
	}

	for _, name := range want {
		if !registered[name] {
			t.Errorf("generate subcommand %q not registered", name)
		}
	}
}

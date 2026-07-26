// Package files provides file resolution and pattern matching for OpenAPI specs.
package files

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// OpenAPIExtensions are file extensions that may contain OpenAPI specs.
var OpenAPIExtensions = []string{".yaml", ".yml", ".json"}

// OpenAPIPatterns are filename patterns that typically indicate OpenAPI specs.
var OpenAPIPatterns = []string{
	"openapi",
	"swagger",
	"api-spec",
	"api_spec",
}

// ResolveSpecs resolves file arguments to a list of OpenAPI spec files.
// It handles:
//   - Explicit file paths
//   - Glob patterns (e.g., "api/*.yaml")
//   - Directories (with recursive flag)
//
// The include/exclude patterns filter the results.
func ResolveSpecs(args []string, recursive bool, include, exclude []string) ([]string, error) {
	var files []string
	seen := make(map[string]bool)

	for _, arg := range args {
		// Check if arg is a glob pattern
		if containsGlob(arg) {
			matches, err := filepath.Glob(arg)
			if err != nil {
				return nil, err
			}
			for _, m := range matches {
				if !seen[m] && shouldInclude(m, include, exclude) {
					files = append(files, m)
					seen[m] = true
				}
			}
			continue
		}

		// Check if arg is a directory
		info, err := os.Stat(arg)
		if err != nil {
			if os.IsNotExist(err) {
				// Might be a glob that didn't match
				continue
			}
			return nil, err
		}

		if info.IsDir() {
			dirFiles, err := resolveDir(arg, recursive, include, exclude)
			if err != nil {
				return nil, err
			}
			for _, f := range dirFiles {
				if !seen[f] {
					files = append(files, f)
					seen[f] = true
				}
			}
			continue
		}

		// It's a regular file - when explicitly provided, only check exclude patterns
		// (bypass include patterns since user explicitly requested this file)
		if !seen[arg] && !isExcluded(arg, exclude) {
			files = append(files, arg)
			seen[arg] = true
		}
	}

	// Sort for consistent ordering
	sort.Strings(files)
	return files, nil
}

// resolveDir walks a directory and returns matching OpenAPI files.
func resolveDir(dir string, recursive bool, include, exclude []string) ([]string, error) {
	var files []string

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories in non-recursive mode
		if info.IsDir() {
			if path != dir && !recursive {
				return filepath.SkipDir
			}
			return nil
		}

		// Check if it looks like an OpenAPI file
		if IsOpenAPIFile(path) && shouldInclude(path, include, exclude) {
			files = append(files, path)
		}

		return nil
	}

	if err := filepath.Walk(dir, walkFn); err != nil {
		return nil, err
	}

	return files, nil
}

// IsOpenAPIFile returns true if the path looks like an OpenAPI specification.
func IsOpenAPIFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	hasValidExt := false
	for _, e := range OpenAPIExtensions {
		if ext == e {
			hasValidExt = true
			break
		}
	}
	if !hasValidExt {
		return false
	}

	// Check if filename matches OpenAPI patterns
	base := strings.ToLower(filepath.Base(path))
	name := strings.TrimSuffix(base, ext)

	for _, pattern := range OpenAPIPatterns {
		if strings.Contains(name, pattern) {
			return true
		}
	}

	// Also accept files named simply "api.yaml", etc.
	if name == "api" {
		return true
	}

	return false
}

// isExcluded checks if a file matches any exclude pattern.
func isExcluded(path string, exclude []string) bool {
	for _, pattern := range exclude {
		if m, _ := matchGlob(path, pattern); m {
			return true
		}
	}
	return false
}

// shouldInclude checks if a file should be included based on patterns.
func shouldInclude(path string, include, exclude []string) bool {
	// If include patterns are specified, file must match at least one
	if len(include) > 0 {
		matched := false
		for _, pattern := range include {
			if m, _ := matchGlob(path, pattern); m {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}

	// If exclude patterns are specified, file must not match any
	for _, pattern := range exclude {
		if m, _ := matchGlob(path, pattern); m {
			return false
		}
	}

	return true
}

// matchGlob matches a path against a glob pattern.
// Supports ** for recursive matching.
func matchGlob(path, pattern string) (bool, error) {
	// Normalize path separators
	path = filepath.ToSlash(path)
	pattern = filepath.ToSlash(pattern)

	// Handle ** patterns
	if strings.Contains(pattern, "**") {
		return matchDoubleStarGlob(path, pattern), nil
	}

	// Fall back to standard filepath.Match
	return filepath.Match(pattern, filepath.Base(path))
}

// matchDoubleStarGlob handles ** glob patterns.
func matchDoubleStarGlob(path, pattern string) bool {
	// Handle patterns like **/dirname/** (matches any path containing dirname)
	if strings.HasPrefix(pattern, "**/") && strings.HasSuffix(pattern, "/**") {
		middle := pattern[3 : len(pattern)-3]
		return strings.Contains(path, "/"+middle+"/") ||
			strings.HasPrefix(path, middle+"/") ||
			strings.HasSuffix(path, "/"+middle)
	}

	// Handle patterns like **/suffix
	if strings.HasPrefix(pattern, "**/") {
		suffix := pattern[3:]
		// Check if path ends with the suffix
		if strings.HasSuffix(path, "/"+suffix) || path == suffix {
			return true
		}
		// Also try filepath.Match on the base name
		if matched, _ := filepath.Match(suffix, filepath.Base(path)); matched {
			return true
		}
		return false
	}

	// Handle patterns like prefix/**
	if strings.HasSuffix(pattern, "/**") {
		prefix := pattern[:len(pattern)-3]
		return strings.HasPrefix(path, prefix+"/") || path == prefix
	}

	// Handle patterns with ** in the middle
	parts := strings.Split(pattern, "**")
	if len(parts) == 2 {
		prefix := strings.TrimSuffix(parts[0], "/")
		suffix := strings.TrimPrefix(parts[1], "/")

		prefixMatch := prefix == "" || strings.HasPrefix(path, prefix)
		suffixMatch := suffix == "" || strings.HasSuffix(path, suffix)

		return prefixMatch && suffixMatch
	}

	return false
}

// containsGlob returns true if the string contains glob metacharacters.
func containsGlob(s string) bool {
	return strings.ContainsAny(s, "*?[")
}

// MatchGlob is the exported version of matchGlob.
func MatchGlob(path, pattern string) (bool, error) {
	return matchGlob(path, pattern)
}

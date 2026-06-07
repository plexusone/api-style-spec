package watch

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestWatcher(t *testing.T) {
	// Create temp directory and file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "openapi.yaml")
	if err := os.WriteFile(testFile, []byte("test"), 0o600); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	cfg := &Config{
		Paths:    []string{tmpDir},
		Debounce: 50 * time.Millisecond,
	}

	w, err := New(cfg)
	if err != nil {
		t.Fatalf("New failed: %v", err)
	}
	defer w.Close()

	// Track events
	events := make(chan Event, 10)
	w.OnChange(func(e Event) error {
		events <- e
		return nil
	})

	// Start watching in background
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	go func() {
		_ = w.Watch(ctx)
	}()

	// Give watcher time to start
	time.Sleep(50 * time.Millisecond)

	// Modify the file
	if err := os.WriteFile(testFile, []byte("updated"), 0o600); err != nil {
		t.Fatalf("Failed to update file: %v", err)
	}

	// Wait for event
	select {
	case e := <-events:
		if e.Path != testFile {
			t.Errorf("Expected path %s, got %s", testFile, e.Path)
		}
	case <-ctx.Done():
		t.Error("Timed out waiting for event")
	}
}

func TestShouldWatch(t *testing.T) {
	w := &Watcher{
		cfg: &Config{
			Include: []string{"*.yaml", "*.yml"},
			Exclude: []string{"*.tmp"},
		},
	}

	tests := []struct {
		path string
		want bool
	}{
		{"openapi.yaml", true},
		{"api.yml", true},
		{"config.json", false},
		{"test.tmp", false},
	}

	for _, tc := range tests {
		t.Run(tc.path, func(t *testing.T) {
			got := w.shouldWatch(tc.path)
			if got != tc.want {
				t.Errorf("shouldWatch(%q) = %v, want %v", tc.path, got, tc.want)
			}
		})
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.Debounce != 100*time.Millisecond {
		t.Errorf("Expected debounce 100ms, got %v", cfg.Debounce)
	}
}

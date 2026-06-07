// Package watch provides file watching for automatic re-linting.
package watch

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Event represents a file change event.
type Event struct {
	// Path is the path to the changed file.
	Path string

	// Op is the operation that triggered the event.
	Op fsnotify.Op

	// Time is when the event occurred.
	Time time.Time
}

// Config configures the watcher.
type Config struct {
	// Paths are the files or directories to watch.
	Paths []string

	// Recursive watches directories recursively.
	Recursive bool

	// Debounce is the duration to wait before triggering callback.
	// Multiple events within this window are coalesced.
	Debounce time.Duration

	// Include patterns for files to watch.
	Include []string

	// Exclude patterns for files to ignore.
	Exclude []string
}

// DefaultConfig returns default watcher configuration.
func DefaultConfig() *Config {
	return &Config{
		Debounce: 100 * time.Millisecond,
	}
}

// Watcher watches files for changes and triggers callbacks.
type Watcher struct {
	cfg      *Config
	fsw      *fsnotify.Watcher
	callback func(Event) error

	mu             sync.Mutex
	pending        map[string]time.Time
	debounceCtx    context.Context
	debounceCancel context.CancelFunc
}

// New creates a new file watcher.
func New(cfg *Config) (*Watcher, error) {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	fsw, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("creating watcher: %w", err)
	}

	return &Watcher{
		cfg:     cfg,
		fsw:     fsw,
		pending: make(map[string]time.Time),
	}, nil
}

// OnChange sets the callback function for file changes.
func (w *Watcher) OnChange(fn func(Event) error) {
	w.callback = fn
}

// Watch starts watching files and blocks until context is cancelled.
func (w *Watcher) Watch(ctx context.Context) error {
	// Add paths to watch
	for _, path := range w.cfg.Paths {
		if err := w.addPath(path); err != nil {
			return fmt.Errorf("adding path %s: %w", path, err)
		}
	}

	// Start debounce goroutine
	w.debounceCtx, w.debounceCancel = context.WithCancel(ctx)
	go w.debounceLoop()

	// Process events
	for {
		select {
		case <-ctx.Done():
			w.debounceCancel()
			return ctx.Err()

		case event, ok := <-w.fsw.Events:
			if !ok {
				return nil
			}
			w.handleEvent(event)

		case err, ok := <-w.fsw.Errors:
			if !ok {
				return nil
			}
			// Log error but continue watching
			fmt.Printf("Watch error: %v\n", err)
		}
	}
}

// addPath adds a path to the watcher.
func (w *Watcher) addPath(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	return w.fsw.Add(absPath)
}

// handleEvent processes a file system event.
func (w *Watcher) handleEvent(event fsnotify.Event) {
	// Only care about write and create events
	if event.Op&(fsnotify.Write|fsnotify.Create) == 0 {
		return
	}

	// Check include/exclude patterns
	if !w.shouldWatch(event.Name) {
		return
	}

	// Add to pending with current time
	w.mu.Lock()
	w.pending[event.Name] = time.Now()
	w.mu.Unlock()
}

// shouldWatch returns true if the file should be watched.
func (w *Watcher) shouldWatch(path string) bool {
	// Check exclude patterns first
	for _, pattern := range w.cfg.Exclude {
		if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			return false
		}
	}

	// If include patterns are set, file must match at least one
	if len(w.cfg.Include) > 0 {
		for _, pattern := range w.cfg.Include {
			if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
				return true
			}
		}
		return false
	}

	return true
}

// debounceLoop processes pending events after debounce period.
func (w *Watcher) debounceLoop() {
	ticker := time.NewTicker(w.cfg.Debounce / 2)
	defer ticker.Stop()

	for {
		select {
		case <-w.debounceCtx.Done():
			return
		case <-ticker.C:
			w.processPending()
		}
	}
}

// processPending processes events that have been pending long enough.
func (w *Watcher) processPending() {
	w.mu.Lock()
	now := time.Now()
	var ready []string

	for path, t := range w.pending {
		if now.Sub(t) >= w.cfg.Debounce {
			ready = append(ready, path)
		}
	}

	for _, path := range ready {
		delete(w.pending, path)
	}
	w.mu.Unlock()

	// Trigger callbacks outside of lock
	for _, path := range ready {
		if w.callback != nil {
			event := Event{
				Path: path,
				Op:   fsnotify.Write,
				Time: now,
			}
			if err := w.callback(event); err != nil {
				fmt.Printf("Callback error for %s: %v\n", path, err)
			}
		}
	}
}

// Close closes the watcher.
func (w *Watcher) Close() error {
	if w.debounceCancel != nil {
		w.debounceCancel()
	}
	return w.fsw.Close()
}

package internal

import (
	"os"
	"path/filepath"
)

// PluginsDir is the directory for built plugins.
func PluginsDir() string {
	return filepath.Join(LibDir(), "plugins")
}

// LibDir is the directory for caddy plugin loader resources.
//
// TODO: choose one of the following explanations...
//
// Option 1:
//
// It is searched in this order:
// 1. Take it directly from the $CADDYPLUG_LIB environment variable.
// 2. Use an XDG Base Directory style path if appropriate:
//    a. If $XDG_CACHE_HOME is set: use $XDG_CACHE_HOME/caddyplug
//    b. If $XDG_RUNTIME_DIR is set: use $HOME/.cache/caddyplug
// 3. Otherwise, use the backward-compatible default: $HOME/caddy/lib
//
// Option 2:
//
// Use the first of the following that exists:
// 1. $CADDYPLUG_LIB environment variable
// 2. If $XDG_CACHE_HOME is non-empty: $XDG_CACHE_HOME/caddyplug
// 3. If $XDG_RUNTIME_DIR is non-empty: $HOME/.cache/caddyplug
// 4. $HOME/caddy/lib
//
// Option 3: Just rely on the code comments.
func LibDir() string {
	// Use CADDYPLUG_LIB if set.
	fromEnv := os.Getenv("CADDYPLUG_LIB")
	if fromEnv != "" {
		return fromEnv
	}
	// Use XDG_CACHE_HOME if set, or use its default value ($HOME/.cache) if
	// XDG_RUNTIME_DIR is set (i.e., the system looks like it wants XDG Base
	// Directories).
	cache := os.Getenv("XDG_CACHE_HOME")
	if cache == "" && os.Getenv("XDG_RUNTIME_DIR") != "" {
		cache = filepath.Join(os.Getenv("HOME"), ".cache")
	}
	if cache != "" {
		return filepath.Join(cache, "caddyplug")
	}
	// Fallback to previous default.
	return filepath.Join(os.Getenv("HOME"), "lib", "caddy")
}

// Package config handles loading and managing engram configuration.
package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// DefaultConfigDir is the default directory for engram configuration.
const DefaultConfigDir = ".engram"

// Config holds the engram application configuration.
type Config struct {
	// StoragePath is the directory where engram stores its data.
	StoragePath string `json:"storage_path"`

	// MaxChunkSize is the maximum size (in bytes) of a single chunk file.
	MaxChunkSize int64 `json:"max_chunk_size"`

	// CompressionEnabled controls whether chunk files are gzip compressed.
	CompressionEnabled bool `json:"compression_enabled"`

	// IndexFile is the name of the manifest/index file.
	IndexFile string `json:"index_file"`
}

// DefaultConfig returns a Config populated with sensible defaults.
func DefaultConfig() *Config {
	return &Config{
		StoragePath:        DefaultConfigDir,
		MaxChunkSize:       1024 * 1024 * 10, // 10 MB
		CompressionEnabled: true,
		IndexFile:          "manifest.json",
	}
}

// Load reads the config file at the given path and returns a Config.
// If the file does not exist, the default config is returned.
func Load(path string) (*Config, error) {
	cfg := DefaultConfig()

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return cfg, nil
		}
		return nil, err
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Save writes the config to the given path, creating directories as needed.
func (c *Config) Save(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// ManifestPath returns the full path to the manifest file.
func (c *Config) ManifestPath() string {
	return filepath.Join(c.StoragePath, c.IndexFile)
}

// ChunksDir returns the full path to the chunks directory.
func (c *Config) ChunksDir() string {
	return filepath.Join(c.StoragePath, "chunks")
}

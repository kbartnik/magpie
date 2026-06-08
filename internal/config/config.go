// Package config loads configuration files from a config file
// and returns a populated `Config` object
package config

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config holds magpie configuration. All fields are optional.
type Config struct {
	DefaultVault string            `yaml:"default_vault"`
	Plugins      map[string]string `yaml:"plugins"`

	InboxPath   string `yaml:"inbox_path"`
	ArchivePath string `yaml:"archive_path"`
	LogPath     string `yaml:"log_path"`
}

// Load reads a Config from path. Returns an empty Config (not an error) if the
// file does not exist.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return &Config{}, nil
		}
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// configPath returns the path to the global config file, honoring XDG_CONFIG_HOME.
func configPath() string {
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return filepath.Join(xdg, "magpie", "config.yaml")
	}
	base, err := os.UserConfigDir()
	if err != nil {
		base = filepath.Join(os.Getenv("HOME"), ".config")
	}

	return filepath.Join(base, "magpie", "config.yaml")
}

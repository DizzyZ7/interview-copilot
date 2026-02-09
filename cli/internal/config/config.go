package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	BaseURL string `json:"base_url"`
	Token   string `json:"token"`
}

func configPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".copilot", "config.json")
}

func Load() (*Config, error) {
	path := configPath()
	b, err := os.ReadFile(path)
	if err != nil {
		return &Config{BaseURL: "http://localhost:8080"}, nil
	}
	var c Config
	err = json.Unmarshal(b, &c)
	if c.BaseURL == "" {
		c.BaseURL = "http://localhost:8080"
	}
	return &c, err
}

func Save(c *Config) error {
	path := configPath()
	os.MkdirAll(filepath.Dir(path), 0755)
	b, _ := json.MarshalIndent(c, "", "  ")
	return os.WriteFile(path, b, 0644)
}

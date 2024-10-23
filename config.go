package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func getConfigFilePath() string {
	var configDir string
	switch runtime.GOOS {
	case "linux":
		configDir = filepath.Join(os.Getenv("HOME"), ".config")
	case "darwin": // macOS
		configDir = filepath.Join(os.Getenv("HOME"), "Library", "Application Support")
	case "windows":
		configDir = os.Getenv("APPDATA")
	default:
		configDir = os.TempDir() // Fallback to a temporary directory
	}
	author := "gVguy"
	app := "BalanceTimer"
	filename := "config.json"
	return filepath.Join(configDir, author, app, filename)
}

var configPath = getConfigFilePath()

type Config struct {
	Intervals IntervalsMap `json:"intervals"`
}

type IntervalsMap = map[SessionState]time.Duration

func NewConfig() *Config {
	return &Config{
		// default values (used as fallback when a different config doesn't exist in the fs, or failed to load)
		Intervals: IntervalsMap{
			Working: 30 * time.Minute,
			Resting: 10 * time.Minute,
			WorkEnd: 20 * time.Second,
		},
	}
}

// Load reads the configuration from the JSON file
func (c *Config) Load() *Config {

	// Read the file content
	byteValue, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Printf("LoadConfig(): ReadFile error, %v\n", err)
		return c
	}

	// Unmarshal the JSON data into the Config struct
	if err := json.Unmarshal(byteValue, c); err != nil {
		fmt.Printf("LoadConfig(): Unmarshal error, %v\n", err)
	}

	return c
}

// Write writes the configuration to the JSON file
func (c *Config) Write() error {
	// Ensure the directory exists
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	// Marshal the Config struct to JSON
	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}
	// Write the JSON data to the config file
	return os.WriteFile(configPath, jsonData, 0644) // 0644 is the permission
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

const ConfigFilename = "follow-liked.json"

type Config struct {
	SpotifyAuthToken string
}

func main() {
	log.SetFlags(0)

	config, err := loadExistingOrEmptyConfig(ConfigFilename)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	if len(config.SpotifyAuthToken) == 0 {
		log.Println("Before using this tool, you must authenticate with Spotify.")
	}

	if err := writeConfigToFile(config, ConfigFilename); err != nil {
		log.Printf("Error %v", err)
	}
}

func loadExistingOrEmptyConfig(path string) (Config, error) {
	var cfg Config

	fileContents, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return cfg, nil
	} else if err != nil {
		return cfg, fmt.Errorf("loading existing config from %s: %w", path, err)
	}

	if err := json.Unmarshal(fileContents, &cfg); err != nil {
		return cfg, fmt.Errorf("unmarshaling JSON from %s: %w", path, err)
	}

	return cfg, nil
}

func writeConfigToFile(cfg Config, path string) error {
	formattedConfig, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return fmt.Errorf("marshaling config data for saving to %s: %w", path, err)
	}

	if err := os.WriteFile(path, formattedConfig, fs.ModePerm); err != nil {
		return fmt.Errorf("saving config to %s: %w", path, err)
	}

	return nil
}

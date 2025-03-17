package config

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

const configPathEnvKey = "BUILDER_WITH_AI_CONFIG"

type Config struct {
	RepositoryDir string `json:"repository_dir"`
}

func MustLoad() *Config {
	path := os.Getenv(configPathEnvKey)
	if path == "" {
		log.Fatalf("path to config not set into environment variable: %s\n", configPathEnvKey)
		return nil
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("not correct path (\"%s\") was set into environment variable: %s\n", path, configPathEnvKey)
	}
	rawFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config from %s: %s\n", path, err)
	}
	var config Config
	if err := json.NewDecoder(bytes.NewBuffer(rawFile)).Decode(&config); err != nil {
		log.Fatalf("failed to parse config from %s: %s\n", path, err)
	}
	return &config
}

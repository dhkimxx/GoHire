package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Database struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Host      string `json:"host"`
		Port      int    `json:"port"`
		Dbname    string `json:"dbname"`
		Charset   string `json:"charset"`
		ParseTime bool   `json:"parseTime"`
		Loc       string `json:"loc"`
	} `json:"database"`
}

var GlobalConfig Config

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("unable to get current working directory: %w", err))
	}

	configPath := filepath.Join(wd, "config.json")

	file, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Errorf("unable to open config file at %s: %w", configPath, err))
	}

	if err = json.Unmarshal(file, &GlobalConfig); err != nil {
		panic(fmt.Errorf("unable to decode config file: %w", err))
	}
}

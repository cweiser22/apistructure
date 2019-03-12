package config

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteConfig struct {
	URI string `json:"URI"`
}

type Constants struct {
	PORT   string       `json:"PORT"`
	Sqlite SqliteConfig `json:"Sqlite"`
}

//Config is the main app configuration.
type Config struct {
	Constants
	Database *sql.DB
}

// NewConfig is used to generate a configuration instance which will be passed around the codebase
func New() (*Config, error) {
	config := Config{}
	var constants Constants
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	jsonData, _ := ioutil.ReadAll(configFile)
	json.Unmarshal(jsonData, &constants)
	config.Constants = constants
	if err != nil {
		return &config, err
	}
	db, err := sql.Open("sqlite3", config.Constants.Sqlite.URI)
	if err != nil {
		return &config, err
	}
	config.Database = db
	return &config, err
}

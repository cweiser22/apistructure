package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

type Constants struct {
	PORT   string
	Sqlite struct {
		URI string
	}
}

type Config struct {
	Constants
	Database *sql.DB
}

// NewConfig is used to generate a configuration instance which will be passed around the codebase
func New() (*Config, error) {
	config := Config{}
	constants, err := initViper()
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

func initViper() (Constants, error) {
	viper.SetConfigName("config") // Configuration fileName without the .TOML or .YAML extension
	viper.AddConfigPath(".")      // Search the root directory for the configuration file
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		return Constants{}, err
	}
	viper.WatchConfig() // Watch for changes to the configuration file and recompile
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.SetDefault("PORT", "8080")
	if err = viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}

package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

var GlobalConfig *Config

// LoadConfig loads the configuration from YAML file and environment variables
func LoadConfig(configPath string) (*Config, error) {
	// Read from YAML
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: error reading config file: %v", err)
	}

	// Read from .env
	viper.SetConfigFile(".env")
	if err := viper.MergeInConfig(); err != nil {
		log.Printf("Note: .env file not found or could not be read, using defaults/environment variables")
	}

	viper.AutomaticEnv()

	// Mapping environment variables manually for clarity (optional but good for .env)
	_ = viper.BindEnv("database.host", "DB_HOST")
	_ = viper.BindEnv("database.port", "DB_PORT")
	_ = viper.BindEnv("database.user", "DB_USER")
	_ = viper.BindEnv("database.password", "DB_PASSWORD")
	_ = viper.BindEnv("database.dbname", "DB_NAME")
	_ = viper.BindEnv("database.sslmode", "DB_SSLMODE")

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	GlobalConfig = &config
	log.Printf("Configuration loaded successfully from %s", configPath)
	return GlobalConfig, nil
}

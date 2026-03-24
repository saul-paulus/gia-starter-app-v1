package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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

// findFile searches for a file in the current directory and its parent directories.
// It returns the absolute path to the file if found, otherwise an empty string.
func findFile(filename string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current working directory: %v", err)
		return ""
	}

	for {
		filePath := filepath.Join(currentDir, filename)
		if _, err := os.Stat(filePath); err == nil {
			return filePath
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir { // Reached root directory
			break
		}
		currentDir = parentDir
	}
	return ""
}

// LoadConfig loads the configuration from YAML file and environment variables
func LoadConfig(configPath string) (*Config, error) {
	// 1. Configure Viper for Environment Variables
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// 2. Try to find the YAML config file by searching upwards
	foundConfigPath := findFile(configPath)
	if foundConfigPath == "" {
		log.Printf("Warning: could not find config file %s in current or parent directories", configPath)
		viper.SetConfigFile(configPath) // Fallback to original path
	} else {
		viper.SetConfigFile(foundConfigPath)
		log.Printf("Configuration found at %s", foundConfigPath)
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: error reading config file: %v", err)
	}

	// 3. Try to find .env file by searching upwards
	foundEnvPath := findFile(".env")
	if foundEnvPath == "" {
		log.Printf("Note: .env file not found in current or parent directories")
	} else {
		viper.SetConfigFile(foundEnvPath)
		viper.SetConfigType("env") // Explicitly set type to env
		if err := viper.MergeInConfig(); err != nil {
			log.Printf("Warning: error merging .env file: %v", err)
		} else {
			log.Printf(".env file loaded from %s", foundEnvPath)
		}
	}

	// 4. Mapping environment variables manually to ensure Unmarshal picks them up
	// We use the keys from .env directly
	if host := viper.GetString("DB_HOST"); host != "" {
		viper.Set("database.host", host)
	}
	if port := viper.GetInt("DB_PORT"); port != 0 {
		viper.Set("database.port", port)
	}
	if user := viper.GetString("DB_USER"); user != "" {
		viper.Set("database.user", user)
	}
	if pass := viper.GetString("DB_PASSWORD"); pass != "" {
		viper.Set("database.password", pass)
	}
	if name := viper.GetString("DB_NAME"); name != "" {
		viper.Set("database.dbname", name)
	}
	if ssl := viper.GetString("DB_SSLMODE"); ssl != "" {
		viper.Set("database.sslmode", ssl)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	GlobalConfig = &config
	return GlobalConfig, nil
}

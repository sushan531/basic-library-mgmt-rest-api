package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	PGDATABASE string `yaml:"PGDATABASE"`
	Host       string `yaml:"HOST"`
	User       string `yaml:"USER"`
	Password   string `yaml:"PASSWORD"`
	Port       string `yaml:"PORT"`
	JWTSecret  string `yaml:"JWT_SECRET"`
	JWTExpiry  string `yaml:"JWT_EXPIRY"`
}

// ReadConfig reads the configuration for a given environment.
// It takes a string parameter 'env' and returns a Config struct and an error.
func ReadConfig(env string) (Config, error) {
	// Initialize an empty map to store Config objects
	obj := make(map[string]Config)

	// Get the current working directory
	dir, _ := os.Getwd()
	// Construct the path to the configuration file
	path := filepath.Join(dir, "env.yaml")

	// Read the YAML file
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		// If the file does not exist, use environment variables to populate the configuration
		if strings.Contains(err.Error(), "no such file or directory") {
			conf := Config{
				PGDATABASE: os.Getenv("PGDATABASE"),
				Host:       os.Getenv("PGHOST"),
				User:       os.Getenv("PGUSER"),
				Password:   os.Getenv("PGPASSWORD"),
				Port:       os.Getenv("PGPORT"),
			}

			// Check if any required field is empty
			if conf.Port == "" || conf.User == "" || conf.Password == "" || conf.Host == "" || conf.JWTSecret == "" || conf.JWTExpiry == "" || conf.PGDATABASE == "" {
				fmt.Printf("Please Set ENV variables or have the env.yaml file in the root of the project")
				return Config{}, err
			}
			return conf, nil
		} else {
			// Handle uncaught exceptions
			return Config{}, err
		}
	}

	// Unmarshal the YAML file into the map of Config objects
	err = yaml.Unmarshal(yamlFile, obj)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
		return Config{}, err
	}

	// Retrieve the configuration based on the specified environment
	conf := obj[env]
	if conf.Password == "" || conf.User == "" || conf.Host == "" || conf.Port == "" {
		fmt.Printf("Please provide all the env variables")
		os.Exit(1)
	}

	return conf, nil
}

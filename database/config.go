package database

import "os"

// Config represents the database configuration
type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

// GetConfig returns the database configuration
func GetConfig() Config {
	return Config{
		DBUsername: os.Getenv("DBUserName"),
		DBPassword: os.Getenv("DBPassword"),
		DBHost:     os.Getenv("DBHost"),
		DBPort:     os.Getenv("DBPort"),
		DBName:     os.Getenv("DBName"),
	}
}

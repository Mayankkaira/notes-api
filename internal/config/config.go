package config

import (
	"fmt"
	"net"
	"net/url"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

type envField struct {
	name  string
	value string
}

func LoadConfig() (Config, error) {
	config := Config{DBHost: os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
	validationSlice := []envField{
		{name: "DB_HOST", value: config.DBHost},
		{name: "DB_PORT", value: config.DBPort},
		{name: "DB_USER", value: config.DBUser},
		{name: "DB_PASSWORD", value: config.DBPassword},
		{name: "DB_NAME", value: config.DBName},
	}
	for _, value := range validationSlice {
		if value.value == "" {
			return Config{}, fmt.Errorf("%s is required", value.name)
		}
	}
	return config, nil
}

func (c Config) DSN() string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.DBUser, c.DBPassword),
		Host:   net.JoinHostPort(c.DBHost, c.DBPort),
		Path:   "/" + c.DBName,
	}
	connectionString := u.String()
	return connectionString
}

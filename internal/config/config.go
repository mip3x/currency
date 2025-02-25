package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	defaultPort       = "8000"
	defaultService    = "currency"
	defaultAuthor     = "unknown"
	defaultVersion    = "unknown"
	defaultUserAgent  = "Currency-App"
	defaultConfigName = "config"
	defaultConfigType = "yaml"
	defaultConfigPath = "."
)

type Config struct {
	Port      string `mapstructure:"port"`
	Service   string `mapstructure:"service"`
	Author    string `mapstructure:"author"`
	Version   string `mapstructure:"version"`
	URL       string `mapstructure:"url"`
	UserAgent string `mapstructure:"useragent"`
}

func LoadConfig() (*Config, error) {
	viper.SetDefault("port", defaultPort)
	viper.SetDefault("service", defaultService)
	viper.SetDefault("author", defaultAuthor)
	viper.SetDefault("version", defaultVersion)
	viper.SetDefault("useragent", defaultUserAgent)

	viper.BindEnv("port", "PORT")
	viper.BindEnv("version", "VERSION")

	viper.SetConfigName(defaultConfigName)
	viper.SetConfigType(defaultConfigType)
	viper.AddConfigPath(defaultConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("Error unmarshaling config: %v", err)
	}

	return &cfg, nil
}

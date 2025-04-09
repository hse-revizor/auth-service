package config

import (
	"errors"

	"github.com/spf13/viper"
)

const (
	ENV_MODE_DEVELOPMENT = iota + 1
	ENV_MODE_PRODUCTION  = iota + 1
	ENV_MODE_STAGE       = iota + 1
)

const (
	envModeDevelopmentStr = "development"
	envModeProductionStr  = "production"
	envModeStageStr       = "stage"
)

type Config struct {
	EnvMode uint8
	Port    uint16
	Host    string
	DB      DataBaseConfig
	GitHub  GitHubConfig
}

type DataBaseConfig struct {
	Host     string
	Port     uint16
	Username string
	Name     string
	Password string
	SSLMode  string
}

type GitHubConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func LoadConfig(envMode string) (*Config, error) {
	mode, err := validateEnvMode(envMode)
	if err != nil {
		return nil, err
	}

	config := new(Config)
	viper.AutomaticEnv()

	viper.SetDefault("PORT", 8383)
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_NAME", "postgres")
	viper.SetDefault("DB_USER", "user")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.SetDefault("GITHUB_CLIENT_ID", "Ov23licEPfBjzQdvCOIl")
	viper.SetDefault("GITHUB_CLIENT_SECRET", "ec1cedd8b1638df6d95e86c93f0a52012a897c15")
	viper.SetDefault("GITHUB_REDIRECT_URL", "http://localhost:8383/api/v1/auth/github/callback")

	config.Port = uint16(viper.GetInt("PORT"))
	config.Host = viper.GetString("HOST")
	config.DB = DataBaseConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     uint16(viper.GetInt("DB_PORT")),
		Name:     viper.GetString("DB_NAME"),
		Username: viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		SSLMode:  viper.GetString("DB_SSLMODE"),
	}
	config.GitHub = GitHubConfig{
		ClientID:     viper.GetString("GITHUB_CLIENT_ID"),
		ClientSecret: viper.GetString("GITHUB_CLIENT_SECRET"),
		RedirectURL:  viper.GetString("GITHUB_REDIRECT_URL"),
	}

	config.EnvMode = mode

	return config, nil
}

func MustLoadConfig(envMode string) *Config {
	config, err := LoadConfig(envMode)
	if err != nil {
		panic(err)
	}

	return config
}
func validateEnvMode(envMode string) (uint8, error) {
	var mode uint8
	switch envMode {
	case envModeDevelopmentStr:
		mode = ENV_MODE_DEVELOPMENT
	case envModeProductionStr:
		mode = ENV_MODE_PRODUCTION
	case envModeStageStr:
		mode = ENV_MODE_STAGE
	default:
		return mode, errors.New("Unknown environment mode")
	}

	return mode, nil
}

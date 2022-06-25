package configs

import (
	"github.com/piTch-time/pitch-backend/infrastructure/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	typeEXT      = "yaml"
	defaultPhase = "dev"
)

// Config ...
type Config struct {
	DBConfig DBConfig `mapstructure:"db-config"`
}

// DBConfig ...
type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
}

// Load ...
func Load(path string) (Config, error) {
	phase := viper.GetString("PHASE")
	logger.Info("viper config is loading...", zap.String("phase", phase))
	config := Config{}
	viper.AddConfigPath(path)
	viper.SetConfigName(phase)
	viper.SetConfigType(typeEXT)

	err := viper.ReadInConfig()

	if err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}

// DatabaseConfig ...
func DatabaseConfig() *DBConfig {
	return &DBConfig{
		Host:     viper.GetString("db-config.host"),
		Port:     viper.GetInt("db-config.port"),
		User:     viper.GetString("db-config.user"),
		Name:     viper.GetString("db-config.name"),
		Password: viper.GetString("db-config.password"),
	}
}

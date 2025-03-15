package util

import (
	"time"

	"github.com/spf13/viper"
)

// config stores all configurations of the application
// Values are read by Viper from config file or ENV variables
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	CloudName           string        `mapstructure:"CLOUD_NAME"`
	CloudApiKey         string        `mapstructure:"CLOUD_API_KEY"`
	CloudApiSecret      string        `mapstructure:"CLOUD_API_SECRET"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// loads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

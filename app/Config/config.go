package config

import "github.com/spf13/viper"

type Config struct {
	DBUrl         string `mapstructure:"DB_URL"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	TokenKey      string `mapstructure:"TOKEN_KEY"`
	ServerPort    int    `mapstructure:"SERVER_PORT"`
}

func LoadConfig(path string, fileName string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

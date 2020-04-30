package cfg

import (
	"log"

	"github.com/spf13/viper"
)

// InitEnv reads and set default values for environment variables.
func InitEnv() {
	viper.AutomaticEnv()
	viper.SetConfigFile("config.yml")

	configErr := viper.ReadInConfig()
	if configErr != nil {
		log.Println("Config file failed to load. Defaulting to env.")
	}

	viper.WatchConfig()
}

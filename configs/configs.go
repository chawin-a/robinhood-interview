package configs

import (
	"log"
	"strings"
	"sync"

	"github.com/chawin-a/robinhood-interview/internal/postgres"
	"github.com/spf13/viper"
)

var (
	configOnce sync.Once
	config     *Config
)

type Config struct {
	App      App             `mapstructure:"app"`
	Postgres postgres.Config `mapstructure:"postgres"`
}

type App struct {
	Port string `mapstructure:"port"`
}

func InitConfig() *Config {
	configOnce.Do(func() {
		configPath := "./configs"
		configName := "configs"

		viper.SetConfigName(configName)
		viper.AddConfigPath(configPath)

		// set default values
		viper.SetDefault("app.port", "8000")

		viper.SetDefault("postgres.host", "localhost")
		viper.SetDefault("postgres.port", "5432")
		viper.SetDefault("postgres.password", "password")
		viper.SetDefault("postgres.dbname", "db")

		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			log.Println("config file not found. using default/env config: " + err.Error())
		}
		viper.AutomaticEnv()

		viper.WatchConfig()
		if err := viper.Unmarshal(&config); err != nil {
			panic(err)
		}
	})
	return config
}

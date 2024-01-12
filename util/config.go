package util

import "github.com/spf13/viper"

type Config struct {
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBName        string `mapstructure:"DB_NAME"`
	SSLMode       string `mapstructure:"SSL_MODE"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	MigrationURL  string `mapstructure:"MIGRATION_URL"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func Loadconfig(path string) (config Config, err error) {
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

package env

import (
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort    string
	OrgsService *OrgsServiceConfig
}

type OrgsServiceConfig struct {
	Host string
	Port string
}

var Conf *Config

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	viper.BindEnv("HTTP_PORT")
	viper.BindEnv("APP_ORGS_SERVICE_HOST")
	viper.BindEnv("APP_ORGS_SERVICE_PORT")

	Conf = &Config{
		HttpPort: viper.GetString("HTTP_PORT"),
		OrgsService: &OrgsServiceConfig{
			Host: viper.GetString("ORGS_SERVICE_HOST"),
			Port: viper.GetString("ORGS_SERVICE_PORT"),
		},
	}
}

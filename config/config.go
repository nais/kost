package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port              string `mapstructure:"PORT"`
	GCPBillingProject string `mapstructure:"GCP_BILLING_PROJECT"`
}

var AppConfig Config

func Load() {
	viper.SetDefault("PORT", "8080")
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
}

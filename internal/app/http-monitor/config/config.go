package config

import (
	"bytes"
	"log"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

const (
	app       = "virtual-box"
	cfgFile   = "config.yaml"
	cfgPrefix = "virtual-box"
)

type (
	Config struct {
		JWT    JWT    `mapstructure:"jwt"`
		Server Server `mapstructure:"server"`
	}

	JWT struct {
		Secret     string `mapstructure:"secret"`
		Expiration int    `mapstructure:"expiration"`
	}

	Server struct {
		Address string `mapstructure:"address"`
	}
)

func (c Config) Validate() error {
	return validator.New().Struct(c)
}

// Init initializes application configuration.
func Init() Config {
	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(bytes.NewReader([]byte(defaultConfig))); err != nil {
		log.Fatalf("error loading default configs: %s", err.Error())
	}
	v.SetEnvPrefix(cfgPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	var cfg Config
	if err := v.UnmarshalExact(&cfg); err != nil {
		log.Fatalf("failed to unmarshal config into struct: %s", err.Error())
	}
	print(cfg.Server.Address)
	return cfg
}

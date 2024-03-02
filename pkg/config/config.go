package config

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Sqlite *struct {
		Path string `mapstructure:"path"`
	} `mapstructure:"sqlite"`

	PostgresQL *struct {
		User       string `mapstructure:"user"`
		Password   string `mapstructure:"password"`
		Host       string `mapstructure:"host"`
		Port       int64  `mapstructure:"port"`
		DBName     string `mapstructure:"dbname"`
		DisableSSL bool   `mapstructure:"disable_ssl"`
	} `mapstructure:"postgresql"`

	MySQL *struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     int64  `mapstructure:"port"`
		DBName   string `mapstructure:"dbname"`
	} `mapstructure:"mysql"`
}

func LoadConfig() error {
	for _, opt := range AllConfigOptions {
		if opt.Default != nil {
			// fmt.Printf("setting default value for '%s': %+v\n", opt.ViperKey(), opt.Default)
			viper.SetDefault(opt.ViperKey(), opt.Default)
		}
	}

	// Read config from the current directory.
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("WATGBRIDGE")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	for _, opt := range AllConfigOptions {
		key := opt.ViperKey()
		isSet := viper.IsSet(key)
		value := viper.Get(key)

		if opt.Required && !isSet {
			return fmt.Errorf("'%s' is a required key", key)
		}

		if !isSet {
			continue
		}

		if opt.Type.Kind() != reflect.Struct && opt.Type != reflect.TypeOf(value) {
			return fmt.Errorf("'%s' is not of expected type: wanted %s, got %T", key, opt.Type, value)
		}

		if opt.Validator != nil {
			if err := opt.Validator(opt, value); err != nil {
				return err
			}
		}
	}

	return nil
}

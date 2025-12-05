package config

import "github.com/spf13/viper"

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUser         string `mapstructure:"POSTGRES_USER"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	DBName         string `mapstructure:"POSTGRES_DATABASE"`
	DBPassword     string `mapstructure:"POSTGRES_PASSWORD"`
	Port           string `mapstructure:"PORT"`
	MigrationsPath string `mapstructure:"MIGRATION_PATH"`
	ConnString     string `mapstructure:"CONN_DB_POSTGRES"`
}

var cfg *Config

func Load() error {
	v := viper.New()
	v.AutomaticEnv()

	if err := v.BindEnv("POSTGRES_HOST"); err != nil {
		return err
	}

	if err := v.BindEnv("POSTGRES_PORT"); err != nil {
		return err
	}

	if err := v.BindEnv("POSTGRES_PASSWORD"); err != nil {
		return err
	}

	if err := v.BindEnv("POSTGRES_DATABASE"); err != nil {
		return err
	}

	if err := v.BindEnv("POSTGRES_USER"); err != nil {
		return err
	}

	if err := v.BindEnv("PORT"); err != nil {
		return err
	}

	if err := v.BindEnv("CONN_DB_POSTGRES"); err != nil {
		return err
	}

	if err := v.BindEnv("MIGRATION_PATH"); err != nil {
		return err
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}

	return nil

}

func Get() *Config {
	return cfg
}

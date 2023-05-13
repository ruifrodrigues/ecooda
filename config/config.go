package config

import (
	"github.com/spf13/viper"
	"os"
)

type OptionsFunc func(conf *Config)

type Config struct {
	Values   *Values
	Database IDatabase
	Cache    string
	Queue    string
	Logger   string
	Elastic  string
}

func NewConfig(opts ...OptionsFunc) (Config, error) {
	setDefaults()

	conf := Config{}

	if err := configureViper(); err != nil {
		return conf, err
	}

	conf.loadValues()
	conf.loadOptions(opts...)

	return conf, nil
}

func setDefaults() {
	// we need to define defaults to make it possible to read from ENV when no .env file provided
	// viper "feature", ref - https://github.com/spf13/viper/issues/584
	viper.SetDefault("DB_DRIVER", "mysql")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_TEST", "test.db")
}

func configureViper() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	viper.AddConfigPath(wd)
	viper.SetConfigType("dotenv")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	return nil
}

func (c *Config) loadValues() {
	v, err := NewValues()
	if err != nil {
		panic(err)
	}

	c.Values = v
}

func (c *Config) loadOptions(opts ...OptionsFunc) {
	for _, fn := range opts {
		fn(c)
	}
}

func WithMySQL(conf *Config) {
	dialer := NewMySQLDialer(*conf)
	conf.Database = NewDatabase(dialer)
}

func WithCache(conf *Config) {
	conf.Cache = "to be Implemented"
}

func WithQueue(conf *Config) {
	conf.Queue = "to be implemented"
}

func WithLogger(conf *Config) {
	conf.Logger = "to be implemented"
}

func WithElastic(conf *Config) {
	conf.Elastic = "to be implemented"
}

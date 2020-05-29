package config

import "github.com/jinzhu/configor"

type Config struct {
	AppName string `default:"main"`
	Port    string `default:"8080"`

	DB struct {
		Use      string `default:"postgres"`
		Postgres struct {
			Enabled  bool   `default:"true"`
			Host     string `default:"localhost"`
			Port     string `default:"5432"`
			UserName string `default:"oseifrimpong"`
			Password string `default:"Frimpong.com20"`
			Database string `default:"dartscoredb"`
		}
	}
}

func NewConfig() (*Config, error) {
	c := &Config{}
	err := configor.Load(c, "../config/config.yml")
	if err != nil {
		return nil, err
	}
	return c, nil
}

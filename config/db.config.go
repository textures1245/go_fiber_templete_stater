package config

import "github.com/spf13/viper"

type DBconfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func LoadDBconfig() DBconfig {
	return DBconfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetInt("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASS"),
		Name:     viper.GetString("DB_NAME"),
	}
}

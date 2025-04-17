package config

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)


type Config struct {
 Server  ServerConfig
 Postgres PostgresConfig
 Redis    RedisConfig
}

type ServerConfig struct {
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SSL      bool
}

type RedisConfig struct {
	Host              string
	Port              string
	Password          string
	Db                string
	MinIdleConnection int
	PoolSize          int
	PoolTimeout       int
}
func GetConfig() *Config{
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v , err :=LoadConfig(cfgPath,"yaml")
	if err != nil {
		log.Fatalf("Cant load config file: %v", err)
		return nil
		
	}
	cfg, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("Cant parse config file: %v", err)
		return nil
		
	}
	return cfg

}


func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Cant unmarshal config file: %v", err)
		return nil, err
	}
	return &cfg, nil

}


func LoadConfig(filename string,filetype string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType(filetype)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Cant read config file: %v", err)
		if _ , ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found ")
		}
		return nil, err
	}
	return v, nil
}




func getConfigPath(env string) string{
	if env == "docker" {
		return "config/config-docker"
	}else if env == "production" {
		return "config/config-production"
	}else {
		return "../config/config-development"
	}


}
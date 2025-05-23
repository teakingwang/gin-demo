package config

import (
	"github.com/spf13/viper"
	"os"
)

var Config *config

type config struct {
	Server   ServerConfig   `yaml:"server" json:"server"`
	Database DatabaseConfig `yaml:"database" json:"database"`
	Redis    RedisConfig    `yaml:"redis" json:"redis"`
	SMS      SMSConfig      `yaml:"sms" json:"sms"`
}

type ServerConfig struct {
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}

type DatabaseConfig struct {
	Dialect  string `yaml:"dialect" json:"dialect"`
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Database string `yaml:"database" json:"database"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
	Schema   string `yaml:"schema" json:"schema"`
	Level    string `yaml:"level" json:"level"`
}

type RedisConfig struct {
	Host         string `yaml:"host" json:"host"`
	Port         int    `yaml:"port" json:"port"`
	Password     string `yaml:"password" json:"password"`
	DB           int    `yaml:"db" json:"db"`
	PoolSize     int    `yaml:"poolSize" json:"poolSize"`
	MinIdleConns int    `yaml:"minIdleConns" json:"minIdleConns"`
}

type SMSConfig struct {
	Provider          string `yaml:"provider" json:"provider"`
	AccessKeyID       string `yaml:"accessKeyId" json:"accessKeyId"`
	AccessKeySecret   string `yaml:"accessKeySecret" json:"accessKeySecret"`
	SignName          string `yaml:"signName" json:"signName"`
	TemplateCode      string `yaml:"templateCode" json:"templateCode"`
	Region            string `yaml:"region" json:"region"`
	CodeExpireSeconds int    `yaml:"codeExpireSeconds" json:"codeExpireSeconds"`
}

func LoadConfig() *config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	path := os.Getenv("MY_APP_CONFIG_PATH")
	if path == "" {
		path = "./resources"
	}
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	Config = &cfg

	return Config
}

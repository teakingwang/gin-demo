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
	JWT      JWTConfig      `yaml:"jwt" json:"jwt"`
	RocketMQ RocketMQConfig `yaml:"rocketmq" json:"rocketmq"`
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

type JWTConfig struct {
	Secret     string `yaml:"secret" json:"secret"`
	TTLSeconds int    `yaml:"ttlSeconds" json:"ttlSeconds"`
}

type RocketMQConfig struct {
	NameServer    string `yaml:"nameServer" json:"nameServer"`       // RocketMQ NameServer 地址
	ProducerGroup string `yaml:"producerGroup" json:"producerGroup"` // Producer 分组
	ProducerTopic string `yaml:"producerTopic" json:"producerTopic"` // Producer 使用的 topic
	ConsumerGroup string `yaml:"consumerGroup" json:"consumerGroup"` // Consumer 分组
	ConsumerTopic string `yaml:"consumerTopic" json:"consumerTopic"` // Consumer 订阅的 topic
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

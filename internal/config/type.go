package config

import "time"

type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
	Telegram Telegram `yaml:"telegram"`
}

type Server struct {
	Name             string        `yaml:"name"`
	Port             int           `yaml:"port"`
	Url              string        `yaml:"external_url"`
	HealthCheckTimer time.Duration `yaml:"health_check_timer"`
}

type Database struct {
	DNS string `yaml:"internal_dns"`
}

type Telegram struct {
	Token string `yaml:"token"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
}

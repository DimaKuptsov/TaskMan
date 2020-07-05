package config

import cfg "github.com/Yalantis/go-config"

const (
	DevelopmentEnvironment = "dev"
	ProductionEnvironment  = "prod"
)

type Config struct {
	AppName     string   `json:"app_name"   envconfig:"APP_NAME"   default:"task-man"`
	Environment string   `json:"env"        envconfig:"ENV"        default:"dev"`
	ListenURL   string   `json:"listen_url" envconfig:"LISTEN_URL" default:":8080"`
	Postgres    Postgres `json:"postgres"`
}

var config Config

func GetConfig() *Config {
	return &config
}

func InitFromFile(configFilePath string) error {
	return cfg.Init(&config, configFilePath)
}

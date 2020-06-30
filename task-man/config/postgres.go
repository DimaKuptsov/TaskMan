package config

import cfg "github.com/Yalantis/go-config"

type Postgres struct {
	Host         string       `json:"host"          envconfig:"POSTGRES_HOST"          default:"localhost"`
	Port         string       `json:"port"          envconfig:"POSTGRES_PORT"          default:"5432"`
	Database     string       `json:"database"      envconfig:"POSTGRES_DB_NAME"       default:"task-man"`
	User         string       `json:"user"          envconfig:"POSTGRES_USER"          default:"root"`
	Password     string       `json:"password"      envconfig:"POSTGRES_PASSWORD"      default:"root"`
	PoolSize     int          `json:"pool_size"     envconfig:"POSTGRES_POOL_SIZE"     default:"10"`
	MaxRetries   int          `json:"max_retries"   envconfig:"POSTGRES_MAX_RETRIES"   default:"5"`
	ReadTimeout  cfg.Duration `json:"read_timeout"  envconfig:"POSTGRES_READ_TIMEOUT"  default:"10s"`
	WriteTimeout cfg.Duration `json:"write_timeout" envconfig:"POSTGRES_WRITE_TIMEOUT" default:"10s"`
}

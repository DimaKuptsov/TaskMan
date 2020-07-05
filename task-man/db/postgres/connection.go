package postgres

import (
	"github.com/DimaKuptsov/task-man/config"
	"github.com/go-pg/pg/v10"
	"golang.org/x/net/context"
	"sync"
	"time"
)

var connectionInstance *pg.DB
var once sync.Once

func GetConnection() (*pg.DB, error) {
	once.Do(func() {
		cfg := config.GetConfig().Postgres
		connectionInstance = pg.Connect(&pg.Options{
			Addr:         cfg.Host + ":" + cfg.Port,
			User:         cfg.User,
			Password:     cfg.Password,
			Database:     cfg.Database,
			PoolSize:     cfg.PoolSize,
			WriteTimeout: time.Duration(cfg.WriteTimeout),
			ReadTimeout:  time.Duration(cfg.ReadTimeout),
			MaxRetries:   cfg.MaxRetries,
		})
	})
	err := connectionInstance.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return connectionInstance, nil
}

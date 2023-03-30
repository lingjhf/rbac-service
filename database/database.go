package database

import (
	"errors"
	"rbac-service/config"
	"rbac-service/utils"
	"time"

	"gorm.io/gorm"
)

type Database[T any] interface {
	Connect() (T, error)
}

func New[R any](config Database[R]) Database[R] {
	return config
}

func ConnectDatabase(config *config.Config, opts ...gorm.Option) (*gorm.DB, error) {
	var db *gorm.DB
	var databaseConfig Database[*gorm.DB]
	switch config.DB_TYPE {
	case "sqlite":
		databaseConfig = NewSqliteWithConfig(config, opts...)
	case "postgresql":
		databaseConfig = NewPostgresqlWithConfig(config, opts...)
	default:
		return nil, errors.New("DB_TYPE必须是sqlite或者postgresql")
	}
	err := utils.Retry(func() error {
		var err error
		db, err = New(databaseConfig).Connect()
		if err != nil {
			return err
		}
		return nil
	}, config.DB_CONNECT_RETRY, time.Duration(config.DB_CONNECT_TIMEOUT)*time.Millisecond)
	if err != nil {
		return nil, err
	}
	return db, nil
}

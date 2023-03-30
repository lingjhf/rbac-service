package database

import (
	"fmt"
	"rbac-service/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteWithConfig(config *config.Config, opts ...gorm.Option) Database[*gorm.DB] {
	return &Sqlite{
		Name: config.DB_NAME,
		Opts: opts,
	}
}

type Sqlite struct {
	Name string
	Opts []gorm.Option
}

func (s *Sqlite) Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", s.Name)), s.Opts...)
}

package database

import (
	"fmt"
	"rbac-service/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteWithConfig(config *config.Config) Database[*gorm.DB] {
	return &Sqlite{
		Name: config.DB_NAME,
	}
}

type Sqlite struct {
	Name string
}

func (s *Sqlite) Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", s.Name)))
}

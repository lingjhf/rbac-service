package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	Name string
}

func (s *Sqlite) Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", s.Name)))
}

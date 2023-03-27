package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgresql struct {
	Name     string
	Host     string
	Port     uint
	User     string
	Password string
}

func (p *Postgresql) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d  dbname=%s user=%s password=%s  sslmode=disable TimeZone=Asia/Shanghai",
		p.Host, p.Port, p.Name, p.User, p.Password,
	)
	return gorm.Open(postgres.Open(dsn))
}

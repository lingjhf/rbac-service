package database

import (
	"fmt"
	"rbac-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresqlWithConfig(config *config.Config, opts ...gorm.Option) Database[*gorm.DB] {
	return &Postgresql{
		Host:     config.DB_HOST,
		Port:     config.DB_PORT,
		Name:     config.DB_NAME,
		User:     config.DB_USER,
		Password: config.DB_PASSWD,
		Opts:     opts,
	}
}

type Postgresql struct {
	Name     string
	Host     string
	Port     uint
	User     string
	Password string
	Opts     []gorm.Option
}

func (p *Postgresql) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d  dbname=%s user=%s password=%s  sslmode=disable TimeZone=Asia/Shanghai",
		p.Host, p.Port, p.Name, p.User, p.Password,
	)
	return gorm.Open(postgres.Open(dsn), p.Opts...)
}

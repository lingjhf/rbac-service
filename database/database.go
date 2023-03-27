package database

import "gorm.io/gorm"

type Database interface {
	Connect() (*gorm.DB, error)
}

func New[T Database](config T) T {
	return config
}

package dao

import "gorm.io/gorm"

type Dao interface {
	GetTenancyById(id string)

	GetUserById(id string)

	GetRoleById(id string)

	GetPermissionById(id string)
}

type DatabaseDao struct {
	DB *gorm.DB
}

func New[T Dao](dao T) T {
	return dao
}

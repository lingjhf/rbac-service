package dao

import (
	"rbac-service/tables"

	"gorm.io/gorm"
)

type Dao interface {

	//租户数据库操作
	CreateTenancy(*tables.Tenancy) error
	GetTenancyById(id string) (*tables.Tenancy, error)
	GetTenancyByName(name string) (*tables.Tenancy, error)

	//用户数据库操作
	CreateUser(*tables.User) error
	UpdateUser(user *tables.User, values map[string]any) error
	GetUserById(id string) (*tables.User, error)
	GetUserByUsername(username string) (*tables.User, error)
	GetUserByEmail(email string) (*tables.User, error)
	GetUserByPhone(phone string) (*tables.User, error)
	GetUserByIdWithTenancy(id, tenancyId string) (*tables.User, error)
	GetUserByUsernameWithTenancy(username, tenancyId string) (*tables.User, error)

	//角色数据库操作
	CreateRole(*tables.Role) error
	GetRoleById(id, tenancyId string) (*tables.Role, error)
	GetRoleByName(name, tenancyId string) (*tables.Role, error)

	//权限数据操作
	CreatePermission(*tables.Permission) error
	GetPermissionById(id, tenancyId string) (*tables.Permission, error)
	GetPermissionByName(name, tenancyId string) (*tables.Permission, error)
}

type DatabaseDao struct {
	DB *gorm.DB
}

func New[T Dao](dao T) T {
	return dao
}

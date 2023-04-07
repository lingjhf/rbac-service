package dao

import (
	"rbac-service/tables"

	"gorm.io/gorm"
)

type Dao interface {
	CreateTenant(*tables.Tenant) error
	CreateUser(*tables.User) error
	CreateUserTenant(*tables.UserTenant) error
	CreateUserAndJoinTenant(user *tables.User, userTenant *tables.UserTenant) error
	CreateUserRole(userRole *tables.UserRole) error
	CreateRole(*tables.Role) error
	CreateRolePermission(rolePermission *tables.RolePermission) error
	CreatePermission(*tables.Permission) error

	UpdateTenant(tenant *tables.Tenant, values map[string]any) error
	UpdateUser(user *tables.User, values map[string]any) error
	UpdateRole(role *tables.Role, values map[string]any) error
	UpdatePermission(permission *tables.Permission, values map[string]any) error

	GetTenantById(id string) (*tables.Tenant, error)
	GetTenantByIdWithOwner(id, owner string) (*tables.Tenant, error)
	GetTenantByName(name string) (*tables.Tenant, error)
	GetTenantByNameWithParent(name string, parentId *string) (*tables.Tenant, error)
	GetTenantByNameWithParentAndOwner(name, owner string, parentId *string) (*tables.Tenant, error)
	GetTenantListByOwner(owner string, offset uint, limit uint) ([]*tables.Tenant, error)
	GetTenantCountByOwner(owner string) (int64, error)
	GetTenantJoinTenantTreeOnDescendant(ancestorId, descendantId string) (*tables.Tenant, error)
	GetTenantTree(ancestorId, descendantId string) (*tables.TenantTree, error)
	GetUserById(id string) (*tables.User, error)
	GetUserByUsername(username string) (*tables.User, error)
	GetUserByEmail(email string) (*tables.User, error)
	GetUserByPhone(phone string) (*tables.User, error)
	GetUserOnTenantById(id, tenantId string) (*tables.User, error)
	GetUserOnTenantByUsername(username, tenantId string) (*tables.User, error)
	GetUserOnTenantTreeById(id, tenantId string) (*tables.User, error)
	GetUserListByTenant(tenantId string, offset uint, limit uint) ([]*tables.User, error)
	GetUserCountByTenant(tenantId string) (int64, error)
	GetUserRoleByUnique(userId, RoleId, tenantId string) (*tables.UserRole, error)
	GetUserTenantByUnique(userId, tenantId string) (*tables.UserTenant, error)
	GetRoleById(id, tenantId string) (*tables.Role, error)
	GetRoleByName(name, tenantId string) (*tables.Role, error)
	GetRoleList(tenantId string, offset uint, limit uint) ([]*tables.Role, error)
	GetRoleCount(tenantId string) (int64, error)
	GetRolePermissionById(id, tenantId string) (*tables.RolePermission, error)
	GetRolePermissionByUnique(roleId, permissionId, tenantId string) (*tables.RolePermission, error)
	GetPermissionById(id, tenantId string) (*tables.Permission, error)
	GetPermissionByName(name, tenantId string) (*tables.Permission, error)
	GetPermissionByCode(code, tenantId string) (*tables.Permission, error)
	GetPermissionList(tenantId string, offset uint, limit uint) ([]*tables.Permission, error)
	GetPermissionCount(tenantId string) (int64, error)
	GetPermissionOnTenantTreeById(id, tenantId string) (*tables.Permission, error)
}

type DatabaseDao struct {
	DB *gorm.DB
}

func New[T Dao](dao T) T {
	return dao
}

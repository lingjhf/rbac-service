package dao

import (
	"rbac-service/tables"

	"gorm.io/gorm"
)

// CreateRole 创建角色
func (d *DatabaseDao) CreateRole(role *tables.Role) error {
	return d.DB.Create(role).Error
}

func (d *DatabaseDao) CreateRolePermission(rolePermission *tables.RolePermission) error {
	return d.DB.Create(rolePermission).Error
}

// UpdateRole 更新角色
func (d *DatabaseDao) UpdateRole(role *tables.Role, values map[string]any) error {
	role.UpdateTime()
	values["update_at"] = role.UpdateAt
	return d.DB.Model(role).Updates(values).Error
}

// GetRoleById 通过id获取角色
func (d *DatabaseDao) GetRoleById(id, tenantId string) (*tables.Role, error) {
	role := &tables.Role{}
	err := d.DB.Model(role).Take(role, "id = ? and tenant_id = ?", id, tenantId).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return role, err
}

// GetRoleByName 通过name获取角色
func (d *DatabaseDao) GetRoleByName(name, tenantId string) (*tables.Role, error) {
	role := &tables.Role{}
	err := d.DB.Model(role).Take(role, "name = ? and tenant_id = ?", name, tenantId).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return role, err
}

func (d *DatabaseDao) GetRoleList(tenantId string, offset uint, limit uint) ([]*tables.Role, error) {
	roleList := []*tables.Role{}
	err := d.DB.
		Where("tenant_id = ?", tenantId).
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&roleList).Error
	if err != nil {
		return nil, err
	}
	return roleList, nil
}

func (d *DatabaseDao) GetRoleCount(tenantId string) (int64, error) {
	var count int64
	err := d.DB.Model(&tables.Role{}).Where("tenant_id = ?", tenantId).Select("id").Count(&count).Error
	return count, err
}

func (d *DatabaseDao) GetRolePermissionById(id, tenantId string) (*tables.RolePermission, error) {
	rolePermission := &tables.RolePermission{}
	err := d.DB.Model(rolePermission).Take(rolePermission, "id = ? and tenant_id = ?", id, tenantId).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return rolePermission, err
}

func (d *DatabaseDao) GetRolePermissionByUnique(roleId, permissionId, tenantId string) (*tables.RolePermission, error) {
	rolePermission := &tables.RolePermission{}
	err := d.DB.
		Model(rolePermission).
		Take(rolePermission, "role_id = ? and permission_id = ? and tenant_id = ?", roleId, permissionId, tenantId).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return rolePermission, err
}

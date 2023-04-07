package dao

import (
	"fmt"
	"rbac-service/tables"

	"gorm.io/gorm"
)

func (d *DatabaseDao) CreatePermission(permission *tables.Permission) error {
	return d.DB.Create(permission).Error
}

func (d *DatabaseDao) UpdatePermission(permission *tables.Permission, values map[string]any) error {
	permission.UpdateTime()
	values["update_at"] = permission.UpdateAt
	return d.DB.Model(permission).Updates(values).Error
}

func (d *DatabaseDao) GetPermissionById(id, tenantId string) (*tables.Permission, error) {
	permission := &tables.Permission{}
	err := d.DB.Model(permission).Take(permission, "id = ? and tenant_id = ?", id, tenantId).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return permission, err
}

func (d *DatabaseDao) GetPermissionByName(name, tenantId string) (*tables.Permission, error) {
	permission := &tables.Permission{}
	err := d.DB.Model(permission).Take(permission, "name = ? and tenant_id = ?", name, tenantId).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return permission, err
}

func (d *DatabaseDao) GetPermissionByCode(code, tenantId string) (*tables.Permission, error) {
	permission := &tables.Permission{}
	err := d.DB.Model(permission).Take(permission, "code = ? and tenant_id = ?", code, tenantId).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return permission, err
}

func (d *DatabaseDao) GetPermissionList(tenantId string, offset uint, limit uint) ([]*tables.Permission, error) {
	permissionList := []*tables.Permission{}
	err := d.DB.
		Where("tenant_id = ?", tenantId).
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&permissionList).Error
	if err != nil {
		return nil, err
	}
	return permissionList, nil
}

func (d *DatabaseDao) GetPermissionCount(tenantId string) (int64, error) {
	var count int64
	err := d.DB.Model(&tables.Permission{}).Where("tenant_id = ?", tenantId).Select("id").Count(&count).Error
	return count, err
}

func (d *DatabaseDao) GetPermissionOnTenantTreeById(id, tenantId string) (*tables.Permission, error) {
	permission := &tables.Permission{}
	tenantTree := tables.TenantTree{}
	permissionTableName := permission.TableName()
	tenantTreeTableName := tenantTree.TableName()
	err := d.DB.Model(permission).
		Joins(
			fmt.Sprintf(
				"join %s on %s.descendant_id=%s.tenant_id",
				tenantTreeTableName,
				tenantTreeTableName,
				permissionTableName,
			),
		).
		Where(
			fmt.Sprintf("%s.id = ? and %s.ancestor_id = ?", permissionTableName, tenantTreeTableName),
			id,
			tenantId,
		).
		Take(permission).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return permission, err
}

package dao

import (
	"fmt"
	"rbac-service/tables"

	"gorm.io/gorm"
)

func (d *DatabaseDao) CreateUser(user *tables.User) error {
	return d.DB.Create(user).Error
}

func (d *DatabaseDao) CreateUserTenant(userTenant *tables.UserTenant) error {
	return d.DB.Create(userTenant).Error
}

func (d *DatabaseDao) CreateUserAndJoinTenant(user *tables.User, userTenant *tables.UserTenant) error {
	return d.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		if err := tx.Create(userTenant).Error; err != nil {
			return err
		}
		return nil
	})
}

func (d *DatabaseDao) CreateUserRole(userRole *tables.UserRole) error {
	return d.DB.Create(userRole).Error
}

func (d *DatabaseDao) UpdateUser(user *tables.User, values map[string]any) error {
	user.UpdateTime()
	values["update_at"] = user.UpdateAt
	return d.DB.Model(user).Updates(values).Error
}

func (d *DatabaseDao) GetUserById(id string) (*tables.User, error) {
	user := &tables.User{}
	err := d.DB.Model(user).Take(user, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, err
}

func (d *DatabaseDao) GetUserByUsername(username string) (*tables.User, error) {
	user := &tables.User{}
	err := d.DB.Model(user).Take(user, "username = ?", username).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, err
}

func (d *DatabaseDao) GetUserByEmail(email string) (*tables.User, error) {
	user := &tables.User{}
	err := d.DB.Model(user).Take(user, "email = ?", email).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, err
}

func (d *DatabaseDao) GetUserByPhone(phone string) (*tables.User, error) {
	user := &tables.User{}
	err := d.DB.Model(user).Take(user, "phone = ?", phone).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, err
}

func (d *DatabaseDao) GetUserListByTenant(tenantId string, offset uint, limit uint) ([]*tables.User, error) {
	user := &tables.User{}
	userTenant := tables.UserTenant{}
	userTableName := user.TableName()
	userTenantTableName := userTenant.TableName()
	userList := []*tables.User{}
	err := d.DB.
		Joins(
			fmt.Sprintf(
				"join %s on %s.user_id=%s.id",
				userTenantTableName,
				userTenantTableName,
				userTableName,
			),
		).
		Where(fmt.Sprintf("%s.tenant_id = ?", userTenantTableName), tenantId).
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}
func (d *DatabaseDao) GetUserCountByTenant(tenantId string) (int64, error) {
	user := &tables.User{}
	userTenant := tables.UserTenant{}
	userTableName := user.TableName()
	userTenantTableName := userTenant.TableName()
	var count int64
	err := d.DB.Debug().Model(user).
		Joins(
			fmt.Sprintf(
				"join %s on %s.user_id=%s.id",
				userTenantTableName,
				userTenantTableName,
				userTableName,
			),
		).
		Where(fmt.Sprintf("%s.tenant_id = ?", userTenantTableName), tenantId).
		Select(fmt.Sprintf("%s.id", userTableName)).
		Count(&count).Error
	return count, err
}

func (d *DatabaseDao) GetUserOnTenantById(id, tenantId string) (*tables.User, error) {
	user := &tables.User{}
	userTenant := tables.UserTenant{}
	userTableName := user.TableName()
	userTenantTableName := userTenant.TableName()
	err := d.DB.Model(user).
		Joins(
			fmt.Sprintf(
				"join %s on %s.user_id=%s.id",
				userTenantTableName,
				userTenantTableName,
				userTableName,
			),
		).
		Where(
			fmt.Sprintf("%s.id = ? and %s.tenant_id = ?", userTableName, userTenantTableName),
			id,
			tenantId,
		).
		Take(user).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, err
}

func (d *DatabaseDao) GetUserOnTenantByUsername(username, tenantId string) (*tables.User, error) {
	user := &tables.User{}
	userTenant := tables.UserTenant{}
	userTableName := user.TableName()
	userTenantTableName := userTenant.TableName()
	err := d.DB.Model(user).
		Joins(
			fmt.Sprintf(
				"join %s on %s.user_id=%s.id",
				userTenantTableName,
				userTenantTableName,
				userTableName,
			),
		).
		Where(
			fmt.Sprintf("%s.username = ? and %s.tenant_id = ?", userTableName, userTenantTableName),
			username,
			tenantId,
		).
		Take(user).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, err
}

func (d *DatabaseDao) GetUserOnTenantTreeById(id, tenantId string) (*tables.User, error) {
	user := &tables.User{}
	userTenant := &tables.UserTenant{}
	tenantTree := &tables.TenantTree{}
	userTableName := user.TableName()
	userTenantTableName := userTenant.TableName()
	tenantTreeTableName := tenantTree.TableName()
	err := d.DB.Debug().Model(user).
		Joins(
			fmt.Sprintf(
				"join %s on %s.user_id=%s.id",
				userTenantTableName, userTenantTableName, userTableName,
			),
		).
		Joins(
			fmt.Sprintf(
				"join %s on %s.descendant_id=%s.tenant_id",
				tenantTreeTableName, tenantTreeTableName, userTenantTableName,
			),
		).
		Where(
			fmt.Sprintf("%s.id = ? and %s.ancestor_id = ?", userTableName, tenantTreeTableName),
			id,
			tenantId,
		).
		Take(user).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, err
}

func (d *DatabaseDao) GetUserRoleByUnique(userId, RoleId, tenantId string) (*tables.UserRole, error) {
	userRole := &tables.UserRole{}
	err := d.DB.
		Model(userRole).
		Take(userRole, "user_id = ? and role_id = ? and tenant_id = ?", userId, RoleId, tenantId).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return userRole, err
}

func (d *DatabaseDao) GetUserTenantByUnique(userId, tenantId string) (*tables.UserTenant, error) {
	userTenant := &tables.UserTenant{}
	err := d.DB.
		Model(userTenant).
		Take(userTenant, "user_id = ? and tenant_id = ?", userId, tenantId).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return userTenant, err
}

package dao

import (
	"fmt"
	"rbac-service/tables"

	"gorm.io/gorm"
)

func createTenantTree(tx *gorm.DB, tenant *tables.Tenant) error {
	tenantTree := tables.TenantTree{}
	tenantTreeTableName := tenantTree.TableName()
	if tenant.ParentId == nil {
		return tx.Exec(
			fmt.Sprintf(
				"insert into %s (ancestor_id, descendant_id, distance) select ?,?,0",
				tenantTreeTableName,
			),
			tenant.Id,
			tenant.Id,
		).Error
	} else {
		return tx.Exec(
			fmt.Sprintf(
				"insert into %s (ancestor_id, descendant_id, distance) select ancestor_id,?,distance+1 from %s where descendant_id = ? union all select ?,?,0",
				tenantTreeTableName,
				tenantTreeTableName,
			),
			tenant.Id,
			tenant.ParentId,
			tenant.Id,
			tenant.Id,
		).Error
	}
}

func (d *DatabaseDao) CreateTenant(tenant *tables.Tenant) error {
	return d.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(tenant).Error; err != nil {
			return err
		}
		userTenant := &tables.UserTenant{UserId: tenant.Owner, TenantId: tenant.Id}
		userTenant.Init()
		if err := tx.Create(userTenant).Error; err != nil {
			return err
		}
		return createTenantTree(tx, tenant)
	})
}

func (d *DatabaseDao) UpdateTenant(tenant *tables.Tenant, values map[string]any) error {
	return d.DB.Transaction(func(tx *gorm.DB) error {
		tenant.UpdateTime()
		values["update_at"] = tenant.UpdateAt
		if err := tx.Model(tenant).Updates(values).Error; err != nil {
			return err
		}
		if parentId, ok := values["parent_id"]; ok {
			tenantTree := tables.TenantTree{}
			tenantTreeTableName := tenantTree.TableName()
			if err := tx.Delete(
				&tables.TenantTree{},
				fmt.Sprintf(`
				descendant_id in (select descendant_id from %s where ancestor_id = ?) 
					and
				ancestor_id in (select ancestor_id from %s where descendant_id = ? and ancestor_id != descendant_id)
				`,
					tenantTreeTableName,
					tenantTreeTableName,
				),
				tenant.Id,
				tenant.Id,
			).Error; err != nil {
				return err
			}
			return tx.Exec(fmt.Sprintf(`
			insert into
 				%s (ancestor_id, descendant_id, distance)
			select
  				super_tree.ancestor_id,
  				sub_tree.descendant_id,
  				super_tree.distance + sub_tree.distance + 1 as distance
			from
  				%s as super_tree
  				cross join %s as sub_tree
			where
  				super_tree.descendant_id = ?
  				and sub_tree.ancestor_id = ?
			`,
				tenantTreeTableName,
				tenantTreeTableName,
				tenantTreeTableName,
			),
				parentId,
				tenant.Id,
			).Error
		}
		return nil
	})
}

func (d *DatabaseDao) GetTenantById(id string) (*tables.Tenant, error) {
	tenant := &tables.Tenant{}
	err := d.DB.Model(tenant).Take(tenant, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return tenant, err
}

func (d *DatabaseDao) GetTenantByIdWithOwner(id, owner string) (*tables.Tenant, error) {
	tenant := &tables.Tenant{}
	err := d.DB.Model(tenant).
		Take(tenant, "id = ? and  owner = ?", id, owner).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return tenant, err
}

func (d *DatabaseDao) GetTenantByName(name string) (*tables.Tenant, error) {
	tenant := &tables.Tenant{}
	err := d.DB.Model(tenant).Take(tenant, "name = ?", name).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return tenant, err
}

func (d *DatabaseDao) GetTenantByNameWithParent(name string, parentId *string) (*tables.Tenant, error) {
	var err error
	tenant := &tables.Tenant{}
	if parentId == nil {
		err = d.DB.Model(tenant).
			Take(tenant, "name = ? and parent_id is null", name).
			Error
	} else {
		err = d.DB.Model(tenant).
			Take(tenant, "name = ? and parent_id = ?", name, parentId).
			Error
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return tenant, err
}

func (d *DatabaseDao) GetTenantByNameWithParentAndOwner(name, owner string, parentId *string) (*tables.Tenant, error) {
	var err error
	tenant := &tables.Tenant{}
	if parentId == nil {
		err = d.DB.Model(tenant).
			Take(tenant, "name = ? and owner = ? and parent_id is null", name, owner).
			Error
	} else {
		err = d.DB.Model(tenant).
			Take(tenant, "name = ? and owner = ? and parent_id = ?", name, owner, parentId).
			Error
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return tenant, err
}

func (d *DatabaseDao) GetTenantListByOwner(owner string, offset, limit uint) ([]*tables.Tenant, error) {
	tenantList := []*tables.Tenant{}
	err := d.DB.
		Where("owner = ?", owner).
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&tenantList).Error
	if err != nil {
		return nil, err
	}
	return tenantList, nil
}

func (d *DatabaseDao) GetTenantCountByOwner(owner string) (int64, error) {
	var count int64
	err := d.DB.Model(&tables.Tenant{}).Where("owner = ?", owner).Select("id").Count(&count).Error
	return count, err
}

func (d *DatabaseDao) GetTenantJoinTenantTreeOnDescendant(ancestorId, descendantId string) (*tables.Tenant, error) {
	tenant := &tables.Tenant{}
	tenantTree := tables.TenantTree{}
	tenantTableName := tenant.TableName()
	tenantTreeTableName := tenantTree.TableName()
	err := d.DB.Model(tenant).
		Joins(
			fmt.Sprintf(
				"join %s on %s.descendant_id = %s.id and %s.ancestor_id = ? and %s.descendant_id = ?",
				tenantTreeTableName,
				tenantTreeTableName,
				tenantTableName,
				tenantTreeTableName,
				tenantTreeTableName,
			),
			ancestorId,
			descendantId,
		).
		Take(tenant).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return tenant, err
}

func (d *DatabaseDao) GetTenantTree(ancestorId, descendantId string) (*tables.TenantTree, error) {
	tenantTree := &tables.TenantTree{}
	err := d.DB.Model(tenantTree).Take(tenantTree, "ancestor_id = ? and descendant_id = ?", ancestorId, descendantId).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return tenantTree, err

}

func (d *DatabaseDao) GetRootTenantListByUser(userId string, offset, limit uint) ([]*tables.Tenant, error) {
	tenant := &tables.Tenant{}
	userTenant := tables.UserTenant{}
	tenantTableName := tenant.TableName()
	userTenantTablelName := userTenant.TableName()
	tenantList := []*tables.Tenant{}
	err := d.DB.
		Model(tenant).
		Joins(
			fmt.Sprintf(
				"join %s on %s.tenant_id = %s.id",
				userTenantTablelName,
				userTenantTablelName,
				tenantTableName,
			),
		).
		Where(
			fmt.Sprintf(
				"%s.parent_id is null and  %s.user_id = ?",
				tenantTableName,
				userTenantTablelName,
			),
			userId,
		).
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&tenantList).Error
	if err != nil {
		return nil, err
	}
	return tenantList, nil

}

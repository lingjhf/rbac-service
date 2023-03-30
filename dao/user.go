package dao

import (
	"rbac-service/tables"

	"gorm.io/gorm"
)

func (d *DatabaseDao) CreateUser(user *tables.User) error {
	return d.DB.Create(user).Error
}

func (d *DatabaseDao) UpdateUser(user *tables.User, values map[string]any) error {
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

func (d *DatabaseDao) GetUserByIdWithTenancy(id, tenancyId string) (*tables.User, error) {
	user := &tables.User{}
	err := d.DB.Model(user).Take(user, "id = ? and tenancy_id = ?", id, tenancyId).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, err
}

func (d *DatabaseDao) GetUserByUsernameWithTenancy(username, tenancyId string) (*tables.User, error) {
	user := &tables.User{}
	err := d.DB.Model(user).Take(user, "username = ?  and tenancy_id = ?", username, tenancyId).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, err
}

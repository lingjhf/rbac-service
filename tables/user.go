package tables

type User struct {
	Base
	Name     string `gorm:"unique;"`
	Password string `gorm:"not null;"`
	Email    string `gorm:"unique;default:'';"`
	Phone    string `gorm:"unique;default:'';"`
}

type UserTenancy struct {
	Base
	UserId    string  `gorm:"type:varchar(32);not null;uniqueIndex:idx_user_tenancy;"`
	TenancyId string  `gorm:"type:varchar(32);not null;uniqueIndex:idx_user_tenancy;"`
	User      User    `gorm:"foreignKey:Id;references:UserId;constraint:OnDelete:CASCADE;"`
	Tenancy   Tenancy `gorm:"foreignKey:Id;references:TenancyId;constraint:OnDelete:CASCADE;"`
}

type UserRole struct {
	Base
	UserId    string  `gorm:"type:varchar(32);not null;uniqueIndex:idx_user_role;"`
	RoleId    string  `gorm:"type:varchar(32);not null;uniqueIndex:idx_user_role;"`
	TenancyId string  `gorm:"type:varchar(32);not null;uniqueIndex:idx_user_role;"`
	User      User    `gorm:"foreignKey:Id;references:UserId;constraint:OnDelete:CASCADE;"`
	Role      Role    `gorm:"foreignKey:Id;references:RoleId;constraint:OnDelete:CASCADE;"`
	Tenancy   Tenancy `gorm:"foreignKey:Id;references:TenancyId;constraint:OnDelete:CASCADE;"`
}

func (User) TableName() string {
	return "rbac_user"
}

func (UserRole) TableName() string {
	return "rbac_user_role"
}

func (UserTenancy) TableName() string {
	return "rbac_user_tenancy"
}

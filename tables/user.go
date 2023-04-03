package tables

type User struct {
	Base
	Username string `gorm:"type:varchar(32);index;unique;not null;"`
	Password string `gorm:"type:varchar(70);not null;"`
	Email    string `gorm:"type:varchar(254);index;not null;default:'';"`
	Phone    string `gorm:"type:varchar(20);index;not null;default:'';"`
}

type UserTenant struct {
	Base
	UserId   string `gorm:"type:varchar(36);not null;uniqueIndex:idx_user_tenant;"`
	TenantId string `gorm:"type:varchar(36);not null;uniqueIndex:idx_user_tenant;"`
	User     User   `gorm:"foreignKey:Id;references:UserId;constraint:OnDelete:CASCADE;"`
	Tenant   Tenant `gorm:"foreignKey:Id;references:TenantId;constraint:OnDelete:CASCADE;"`
}

type UserRole struct {
	Base
	UserId   string `gorm:"type:varchar(36);not null;uniqueIndex:idx_user_role;"`
	RoleId   string `gorm:"type:varchar(36);not null;uniqueIndex:idx_user_role;"`
	TenantId string `gorm:"type:varchar(36);not null;uniqueIndex:idx_user_role;"`
	User     User   `gorm:"foreignKey:Id;references:UserId;constraint:OnDelete:CASCADE;"`
	Role     Role   `gorm:"foreignKey:Id;references:RoleId;constraint:OnDelete:CASCADE;"`
	Tenant   Tenant `gorm:"foreignKey:Id;references:TenantId;constraint:OnDelete:CASCADE;"`
}

func (User) TableName() string {
	return "rbac_user"
}

func (UserRole) TableName() string {
	return "rbac_user_role"
}

func (UserTenant) TableName() string {
	return "rbac_user_tenant"
}

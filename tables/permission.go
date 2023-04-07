package tables

type Permission struct {
	Base
	Name        string `gorm:"type:varchar(64);not null;uniqueIndex:idx_permission_name;"`
	Code        string `gorm:"type:varchar(64);not null;uniqueIndex:idx_permission_code;"`
	Description string `gorm:"type:varchar(128);default:'';"`
	TenantId    string `gorm:"type:varchar(36);not null;uniqueIndex:idx_permission_name;uniqueIndex:idx_permission_code;"`
	Tenant      Tenant `gorm:"foreignKey:Id;references:TenantId;constraint:OnDelete:CASCADE;"`
}

func (Permission) TableName() string {
	return "rbac_permission"
}

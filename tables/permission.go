package tables

type Permission struct {
	Base
	Name        string  `gorm:"type:varchar(64);not null;"`
	Code        string  `gorm:"type:varchar(64);not null;"`
	Description string  `gorm:"type:varchar(128);default:'';"`
	TenancyId   string  `gorm:"type:varchar(32);not null;"`
	Tenancy     Tenancy `gorm:"foreignKey:Id;references:TenancyId;constraint:OnDelete:CASCADE;"`
}

func (Permission) TableName() string {
	return "rbac_permission"
}

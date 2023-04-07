package tables

type Tenant struct {
	Base
	Name     string  `gorm:"type:varchar(32);not null;"`
	Owner    string  `gorm:"type:varchar(36);not null;"`
	ParentId *string `gorm:"varchar(36);"`
	User     *User   `gorm:"foreignKey:Id;references:Owner;constraint:OnDelete:CASCADE;"`
	Parent   *Tenant `gorm:"foreignkey:ParentId;constraint:OnDelete:CASCADE;"`
}

type TenantTree struct {
	AncestorId       string `gorm:"type:varchar(36);not null;uniqueIndex:idx_tenant_tree;"` //祖先节点
	DescendantId     string `gorm:"type:varchar(36);not null;uniqueIndex:idx_tenant_tree;"` //子孙节点
	Distance         uint64 `gorm:"not null;uniqueIndex:idx_tenant_tree;"`                  //子孙到祖先的距离
	AncestorTenant   Tenant `gorm:"foreignKey:Id;references:AncestorId;constraint:OnDelete:CASCADE;"`
	DescendantTenant Tenant `gorm:"foreignKey:Id;references:DescendantId;constraint:OnDelete:CASCADE;"`
}

func (Tenant) TableName() string {
	return "rbac_tenant"
}

func (TenantTree) TableName() string {
	return "rbac_tenant_tree"
}

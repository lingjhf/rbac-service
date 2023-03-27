package tables

type Tenancy struct {
	Base
	Name     string   `gorm:"not null"`
	ParentId *string  `gorm:"varchar(32)"`
	Parent   *Tenancy `gorm:"foreignkey:ParentId"`
}

type TenancyTree struct {
	Base
	AncestorId        string  `gorm:"type:varchar(32);not null;uniqueIndex:idx_tenancy_tree;"` //祖先节点
	DescendantId      string  `gorm:"type:varchar(32);not null;uniqueIndex:idx_tenancy_tree;"` //子孙节点
	Distance          uint64  `gorm:"not null;uniqueIndex:idx_tenancy_tree;"`                  //子孙到祖先的距离
	AncestorTenancy   Tenancy `gorm:"foreignKey:Id;references:AncestorId;constraint:OnDelete:CASCADE;"`
	DescendantTenancy Tenancy `gorm:"foreignKey:Id;references:DescendantId;constraint:OnDelete:CASCADE;"`
}

func (Tenancy) TableName() string {
	return "rbac_tenancy"
}

func (TenancyTree) TableName() string {
	return "rbac_tenancy_tree"
}

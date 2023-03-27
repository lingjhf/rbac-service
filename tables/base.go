package tables

type Base struct {
	Id       string `gorm:"primaryKey;type:varchar(32);"`
	UpdateAt int    `gorm:"not null;"`
	CreateAt int    `gorm:"not null;"`
}

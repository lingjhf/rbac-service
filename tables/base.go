package tables

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	Id       string `gorm:"primaryKey;type:varchar(36);"`
	UpdateAt int64  `gorm:"not null;"`
	CreateAt int64  `gorm:"not null;"`
}

func (b *Base) Init() {
	t := time.Now().UnixMilli()
	b.Id = uuid.New().String()
	b.CreateAt = t
	b.UpdateAt = t
}

func (b *Base) UpdateTime() {
	b.UpdateAt = time.Now().UnixMilli()
}

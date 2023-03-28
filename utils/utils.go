package utils

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

// Retry 重试函数
func Retry(fn func() error, n uint, d time.Duration) (err error) {
	var count uint
	for {
		err = fn()
		if err != nil {
			time.Sleep(d)
			count++
		} else {
			err = nil
			break
		}
		if count > n {
			break
		}
	}
	return
}

// GenerateCaptcha 随机生成4位数验证码
func GenerateCaptcha() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}

// CreateTablesIfNotExists 判断表如果不存在就创建表
func CreateTablesIfNotExists(db *gorm.DB, tables ...interface{}) error {
	for _, table := range tables {
		if !db.Migrator().HasTable(table) {
			if err := db.Migrator().CreateTable(table); err != nil {
				return err
			}
		}
	}
	return nil
}

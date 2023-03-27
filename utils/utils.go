package utils

import (
	"fmt"
	"math/rand"
	"time"
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

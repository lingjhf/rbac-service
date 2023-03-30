package config

import (
	"io"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	//服务配置
	HTTP_PORT string
	RETRY     uint
	TIMEOUT   uint

	//数据库配置
	DB_TYPE            string //sqlite, postgresql, mysql
	DB_HOST            string
	DB_PORT            uint
	DB_NAME            string //数据库名称
	DB_USER            string
	DB_PASSWD          string
	DB_CONNECT_RETRY   uint
	DB_CONNECT_TIMEOUT uint //毫秒为单位

	//缓存配置
	CACHE_TYPE   string //redis
	CACHE_HOST   string
	CACHE_PORT   uint
	CACHE_DB     uint
	CACHE_USER   string
	CACHE_PASSWD string

	//jwt
	JWT_SECRET_KEY string
}

func New() (*Config, error) {
	file, _ := os.Open("config.yml")
	c := &Config{}
	c.readFromFile(file)
	return c, nil
}

// readFromFile 读取配置文件
func (c *Config) readFromFile(in io.Reader) error {
	v := viper.New()
	v.SetConfigType("yaml")
	err := v.ReadConfig(in)
	if err != nil {
		return err
	}
	err = v.Unmarshal(c)
	if err != nil {
		return err
	}
	return nil
}

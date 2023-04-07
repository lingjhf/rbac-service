package main

import (
	"fmt"
	"rbac-service/cache"
	"rbac-service/config"
	"rbac-service/dao"
	"rbac-service/database"
	"rbac-service/services"
	"rbac-service/tables"
	"rbac-service/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := readConfig()
	dao := createDao(config)
	cache := createCache(config)
	runService(&services.Service{App: fiber.New(), Config: config, Dao: dao, Cache: cache})
}

func runService(service *services.Service) {
	services.New(service).
		NewAuthService().
		NewTenantService().
		NewUserService().
		NewRoleService().
		NewPermissionService()
	utils.Retry(func() error {
		err := service.App.Listen(fmt.Sprintf(":%s", service.Config.HTTP_PORT))
		if err != nil {
			return err
		}
		return nil
	}, service.Config.RETRY, time.Duration(service.Config.TIMEOUT)*time.Millisecond)
}

func readConfig() *config.Config {
	config, _ := config.New()
	return config
}

func createDao(config *config.Config) dao.Dao {
	db, err := database.ConnectDatabase(config)
	if err != nil {
		panic("数据库连接失败")
	}
	utils.CreateTablesIfNotExists(
		db,
		&tables.User{},
		&tables.Tenant{},
		&tables.TenantTree{},
		&tables.Role{},
		&tables.Permission{},
		&tables.UserTenant{},
		&tables.UserRole{},
		&tables.RolePermission{},
	)
	return &dao.DatabaseDao{DB: db}
}

func createCache(config *config.Config) cache.Cache {
	client, err := database.NewRedisWithConfig(config).Connect()
	if err != nil {
		panic("缓存连接失败")
	}

	return &cache.RedisCache{Client: client}
}

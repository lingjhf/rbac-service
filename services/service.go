package services

import (
	"rbac-service/cache"
	"rbac-service/config"
	"rbac-service/dao"

	"github.com/gofiber/fiber/v2"
)

type Servicer interface {
	NewTenancyService() Servicer
	NewUserService() Servicer
	NewRoleService() Servicer
	NewPermissionService() Servicer
}

type Service struct {
	Config *config.Config
	App    *fiber.App
	Dao    dao.Dao
	Cache  cache.Cache
}

func New[T Servicer](s T) T {
	return s
}

package services

import (
	"rbac-service/errors"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) NewPermissionService() Servicer {
	s.App.Route("/permission", func(router fiber.Router) {
		router.Get("/:id", s.GetPermissionById).Name("get_permission_by_id")
		router.Post("/create", s.CreatePermission).Name("create_permission")
	}, "permission.")
	return s
}

func (s *Service) GetPermissionById(c *fiber.Ctx) error {

	return c.JSON(&errors.JSONResponse{})
}

func (s *Service) CreatePermission(c *fiber.Ctx) error {

	return c.JSON(&errors.JSONResponse{})
}

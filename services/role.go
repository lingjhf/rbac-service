package services

import (
	"rbac-service/errors"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) NewRoleService() Servicer {
	s.App.Route("/role", func(router fiber.Router) {
		router.Get("/:id", s.GetRoleById).Name("get_role_by_id")
		router.Post("/create", s.CreateRole).Name("create_role")
	}, "role.")
	return s
}

func (s *Service) GetRoleById(c *fiber.Ctx) error {

	return c.JSON(&errors.JSONResponse{})
}

func (s *Service) CreateRole(c *fiber.Ctx) error {

	return c.JSON(&errors.JSONResponse{})
}

package services

import (
	"rbac-service/errors"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) NewTenancyService() Servicer {
	s.App.Route("/tenancy", func(router fiber.Router) {
		router.Get("/:id", s.GetTenancyById).Name("get_tenancy_by_id")
		router.Post("/create", s.CreateTenancy).Name("create_tenancy")
	}, "tenancy.")
	return s
}

func (s *Service) GetTenancyById(c *fiber.Ctx) error {

	return c.JSON(&errors.JSONResponse{})
}

func (s *Service) CreateTenancy(c *fiber.Ctx) error {

	return c.JSON(&errors.JSONResponse{})
}

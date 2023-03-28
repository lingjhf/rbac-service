package services

import (
	"rbac-service/errors"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) NewUserService() Servicer {
	s.App.Route("/user", func(router fiber.Router) {
		router.Get("/:id", s.GetUserById).Name("get_user_by_id")
		router.Post("/create", s.CreateUser).Name("create_user")
	}, "user.")
	return s
}

func (s *Service) GetUserById(c *fiber.Ctx) error {

	return c.JSON(&errors.JSONResponse{})
}

func (s *Service) CreateUser(c *fiber.Ctx) error {

	return c.JSON(&errors.JSONResponse{})
}

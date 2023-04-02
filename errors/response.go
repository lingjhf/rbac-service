package errors

import "github.com/gofiber/fiber/v2"

type JSONResponse struct {
	Code    string `json:"code"`
	Message any    `json:"message"`
	Data    any    `json:"data"`
}

func Succeeded(c *fiber.Ctx) error {
	return c.JSON(&JSONResponse{
		Code:    ErrSucceeded.Code,
		Message: ErrSucceeded.Message,
	})
}

func SucceededWithData(c *fiber.Ctx, data any) error {
	return c.JSON(&JSONResponse{
		Code:    ErrSucceeded.Code,
		Message: ErrSucceeded.Message,
		Data:    data,
	})
}

func BodyParserError(c *fiber.Ctx, message any) error {
	return c.JSON(&JSONResponse{
		Code:    ErrBodyParser.Code,
		Message: message,
	})
}

func QueryParserError(c *fiber.Ctx, message any) error {
	return c.JSON(&JSONResponse{
		Code:    ErrQueryParser.Code,
		Message: message,
	})
}

func ParameterError(c *fiber.Ctx, message any) error {
	return c.JSON(&JSONResponse{
		Code:    ErrParameter.Code,
		Message: message,
	})
}

func DatabaseError(c *fiber.Ctx) error {
	return c.JSON(&JSONResponse{
		Code:    ErrDatabase.Code,
		Message: ErrDatabase.Message,
	})
}

func SignupError(c *fiber.Ctx) error {
	return c.JSON(&JSONResponse{
		Code:    ErrSignin.Code,
		Message: ErrSignin.Message,
	})
}

func SigninError(c *fiber.Ctx) error {
	return c.JSON(&JSONResponse{
		Code:    ErrSignin.Code,
		Message: ErrSignin.Message,
	})
}

func UnauthorizedError(c *fiber.Ctx) error {
	return c.JSON(&JSONResponse{
		Code:    ErrUnauthorized.Code,
		Message: ErrUnauthorized.Message,
	})
}

func TokenExpiredError(c *fiber.Ctx) error {
	return c.JSON(&JSONResponse{
		Code:    ErrTokenExpired.Code,
		Message: ErrTokenExpired.Message,
	})
}

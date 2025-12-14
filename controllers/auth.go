package controllers

import (
	"github.com/gofiber/fiber/v3"
	"govault/models"
	"govault/repository"
	"govault/util"
)

type AuthController interface {
	SignUp(ctx fiber.Ctx) error
	SignIn(ctx fiber.Ctx) error
	GetUser(ctx fiber.Ctx) error
	GetUsers(ctx fiber.Ctx) error
	PutUser(ctx fiber.Ctx) error
	DeleteUser(ctx fiber.Ctx) error
}

type authController struct {
	usersRepo repository.UsersRepository
}

func NewAuthController(usersRepo repository.UsersRepository) AuthController {
	return &authController{usersRepo: usersRepo}
}

func (c *authController) SignUp(ctx fiber.Ctx) error {
	newUser := new(models.User)

	if err := ctx.Bind().Body(newUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(util.NewRequestError(err))
	}

	newUser.Email = util.NormalizeEmail(newUser.Email)
	// validate email
	// https://github.com/EricLau1/go-fiber-auth-api/blob/master/controllers/auth.go

	return nil
}

func (c *authController) SignIn(ctx fiber.Ctx) error {
	return nil
}

func (c *authController) GetUser(ctx fiber.Ctx) error {
	return nil
}

func (c *authController) GetUsers(ctx fiber.Ctx) error {
	return nil
}

func (c *authController) PutUser(ctx fiber.Ctx) error {
	return nil
}

func (c *authController) DeleteUser(ctx fiber.Ctx) error {
	return nil
}

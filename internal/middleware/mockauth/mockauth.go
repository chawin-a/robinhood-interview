package middleware_mockauth

import (
	"context"
	"log"
	"math/rand"

	usecase_user "github.com/chawin-a/robinhood-interview/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
)

func NewMockAuth(UserUsecase *usecase_user.Usecase) fiber.Handler {
	userUsecase := UserUsecase
	return func(c *fiber.Ctx) error {
		users, err := userUsecase.List(c.UserContext())
		if err != nil {
			log.Println(err)
			return err
		}
		ctx := context.WithValue(c.UserContext(), "userId", users[rand.Intn(len(users))].Id)
		c.SetUserContext(ctx)
		return c.Next()
	}
}

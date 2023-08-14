package handler_user

import (
	"log"
	"net/http"

	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type GetRequestBody struct {
	Id uuid.UUID `params:"id"`
}

type GetResponseBody struct {
	User *entity.User `json:"user"`
}

func (h *Handler) Get(ctx *fiber.Ctx) error {
	req := &GetRequestBody{}
	if err := ctx.ParamsParser(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}
	user, err := h.userUsecase.Get(ctx.UserContext(), req.Id)
	if err != nil {
		log.Println(ctx.Method(), ctx.Path(), err)
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{"error": "something went wrong"})
	}

	return ctx.JSON(&GetResponseBody{
		User: user,
	})
}

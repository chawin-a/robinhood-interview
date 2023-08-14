package handler_comment

import (
	"log"
	"net/http"

	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateRequestBody struct {
	InterviewId uuid.UUID `params:"interviewId"`
	UserId      uuid.UUID `json:"user_id"`
	Comment     string    `json:"comment"`
}

type CreateResponseBody struct {
	Comment *entity.Comment `json:"comment"`
}

func (h *Handler) Create(ctx *fiber.Ctx) error {
	req := &CreateRequestBody{}
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}
	if err := ctx.ParamsParser(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	comment, err := h.commentUsecase.Create(ctx.UserContext(), &entity.Comment{
		UserId:      req.UserId,
		InterviewId: req.InterviewId,
		Comment:     req.Comment,
	})
	if err != nil {
		log.Println(ctx.Method(), ctx.Path(), err)
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{"error": "something went wrong"})
	}

	return ctx.JSON(&CreateResponseBody{
		Comment: comment,
	})
}

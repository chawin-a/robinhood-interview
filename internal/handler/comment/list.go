package handler_comment

import (
	"log"
	"net/http"

	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ListRequestBody struct {
	InterviewId uuid.UUID `params:"interviewId"`
}

type ListResponseBody struct {
	Comments []*entity.Comment `json:"comments"`
}

func (h *Handler) List(ctx *fiber.Ctx) error {
	req := &ListRequestBody{}
	if err := ctx.ParamsParser(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	comments, err := h.commentUsecase.List(ctx.UserContext(), req.InterviewId)
	if err != nil {
		log.Println(ctx.Method(), ctx.Path(), err)
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{"error": "something went wrong"})
	}

	return ctx.JSON(&ListResponseBody{
		Comments: comments,
	})
}

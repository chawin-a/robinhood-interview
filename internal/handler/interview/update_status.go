package handler_interview

import (
	"log"
	"net/http"

	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UpdateStatusRequestBody struct {
	Id     uuid.UUID              `params:"id" json:"-"`
	Status entity.InterviewStatus `params:"-" json:"status"`
}

type UpdateStatusResponseBody struct {
	Interview *entity.Interview `json:"interview"`
}

func (h *Handler) UpdateStatus(ctx *fiber.Ctx) error {
	req := &UpdateStatusRequestBody{}
	if err := ctx.ParamsParser(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	interview, err := h.interviewUsecase.UpdateStatus(ctx.UserContext(), req.Id, req.Status)
	if err != nil {
		log.Println(ctx.Method(), ctx.Path(), err)
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{"error": "something went wrong"})
	}

	return ctx.JSON(&UpdateStatusResponseBody{
		Interview: interview,
	})
}

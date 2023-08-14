package handler_interview

import (
	"log"
	"net/http"
	"time"

	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type ListRequestBody struct {
	Limit           int       `json:"limit"`
	LatestTimestamp time.Time `json:"latest_timestamp"`
}

type ListResponseBody struct {
	Interviews      []*entity.Interview `json:"interviews"`
	LatestTimestamp time.Time           `json:"latest_timestamp"`
}

func (h *Handler) List(ctx *fiber.Ctx) error {
	req := &ListRequestBody{}
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	interviews, err := h.interviewUsecase.List(ctx.UserContext(), req.Limit, req.LatestTimestamp)
	if err != nil {
		log.Println(ctx.Method(), ctx.Path(), err)
		return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{"error": "something went wrong"})
	}

	lastTimestamp := time.Now()
	if len(interviews) > 0 {
		lastTimestamp = interviews[len(interviews)-1].CreatedAt
	}

	return ctx.JSON(&ListResponseBody{
		Interviews:      interviews,
		LatestTimestamp: lastTimestamp,
	})
}

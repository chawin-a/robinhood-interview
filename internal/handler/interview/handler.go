package handler_interview

import (
	usecase_interview "github.com/chawin-a/robinhood-interview/internal/usecase/interview"
)

type Handler struct {
	interviewUsecase *usecase_interview.Usecase
}

func NewHandler(uc *usecase_interview.Usecase) *Handler {
	return &Handler{
		interviewUsecase: uc,
	}
}

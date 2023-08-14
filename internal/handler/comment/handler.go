package handler_comment

import usecase_comment "github.com/chawin-a/robinhood-interview/internal/usecase/comment"

type Handler struct {
	commentUsecase *usecase_comment.Usecase
}

func NewHandler(uc *usecase_comment.Usecase) *Handler {
	return &Handler{
		commentUsecase: uc,
	}
}

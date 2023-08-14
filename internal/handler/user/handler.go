package handler_user

import usecase_user "github.com/chawin-a/robinhood-interview/internal/usecase/user"

type Handler struct {
	userUsecase *usecase_user.Usecase
}

func NewHandler(uc *usecase_user.Usecase) *Handler {
	return &Handler{
		userUsecase: uc,
	}
}

package usecase_user

import (
	"context"

	"github.com/chawin-a/robinhood-interview/internal/datagateway"
	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/google/uuid"
)

type Usecase struct {
	userDg datagateway.User
}

func NewUsecase(userDg datagateway.User) *Usecase {
	return &Usecase{
		userDg: userDg,
	}
}

func (u *Usecase) Get(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return u.userDg.Get(ctx, id)
}

func (u *Usecase) List(ctx context.Context) ([]*entity.User, error) {
	return u.userDg.List(ctx)
}

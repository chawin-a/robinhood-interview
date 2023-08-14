package usecase_interview

import (
	"context"
	"time"

	"github.com/chawin-a/robinhood-interview/internal/datagateway"
	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/google/uuid"
)

type Usecase struct {
	interviewDg datagateway.Interview
}

func NewUsecase(interviewDg datagateway.Interview) *Usecase {
	return &Usecase{
		interviewDg: interviewDg,
	}
}

func (u *Usecase) List(ctx context.Context, limit int, t time.Time) ([]*entity.Interview, error) {
	return u.interviewDg.List(ctx, limit, t)
}

func (u *Usecase) Get(ctx context.Context, id uuid.UUID) (*entity.Interview, error) {
	return u.interviewDg.Get(ctx, id)
}

func (u *Usecase) UpdateStatus(ctx context.Context, id uuid.UUID, status entity.InterviewStatus) (*entity.Interview, error) {
	return u.interviewDg.UpdateStatus(ctx, id, status)
}

func (u *Usecase) Archive(ctx context.Context, id uuid.UUID) (*entity.Interview, error) {
	return u.interviewDg.Archive(ctx, id)
}

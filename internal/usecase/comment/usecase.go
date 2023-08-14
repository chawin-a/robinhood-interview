package usecase_comment

import (
	"context"

	"github.com/chawin-a/robinhood-interview/internal/datagateway"
	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/google/uuid"
)

type Usecase struct {
	commentDg datagateway.Comment
}

func NewUsecase(commentDg datagateway.Comment) *Usecase {
	return &Usecase{
		commentDg: commentDg,
	}
}

func (u *Usecase) List(ctx context.Context, interviewId uuid.UUID) ([]*entity.Comment, error) {
	return u.commentDg.ListByInterviewId(ctx, interviewId)
}

func (u *Usecase) Create(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {
	return u.commentDg.Create(ctx, comment)
}

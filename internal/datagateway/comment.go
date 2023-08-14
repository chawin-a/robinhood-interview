package datagateway

import (
	"context"

	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/google/uuid"
)

type Comment interface {
	Create(ctx context.Context, comment *entity.Comment) (*entity.Comment, error)
	ListByInterviewId(ctx context.Context, interviewId uuid.UUID) ([]*entity.Comment, error)
}

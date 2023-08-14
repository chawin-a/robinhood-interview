package datagateway

import (
	"context"

	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/google/uuid"
)

type User interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Get(ctx context.Context, id uuid.UUID) (*entity.User, error)
}

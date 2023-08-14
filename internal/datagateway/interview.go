package datagateway

import (
	"context"
	"time"

	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/google/uuid"
)

type Interview interface {
	// List interview records older than t.
	List(ctx context.Context, limit int, t time.Time) ([]*entity.Interview, error)
	Get(ctx context.Context, id uuid.UUID) (*entity.Interview, error)
	Create(ctx context.Context, interview *entity.Interview) (*entity.Interview, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status entity.InterviewStatus) (*entity.Interview, error)
	Archive(ctx context.Context, id uuid.UUID) (*entity.Interview, error)
}

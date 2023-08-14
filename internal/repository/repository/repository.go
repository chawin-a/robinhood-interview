package repository_interview

import (
	"context"
	"time"

	"github.com/chawin-a/robinhood-interview/internal/datagateway"
	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type InterviewRepository struct {
	db *bun.DB
}

// Create implements datagateway.Interview.
func (r *InterviewRepository) Create(ctx context.Context, interview *entity.Interview) (*entity.Interview, error) {
	interviewModel := fromEntity(interview)
	_, err := r.db.NewInsert().
		Model(interviewModel).
		ExcludeColumn("id", "created_at", "updated_at").
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return interviewModel.ToEntity(), nil
}

// Get implements datagateway.Interview.
func (*InterviewRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Interview, error) {
	panic("unimplemented")
}

// List implements datagateway.Interview.
func (*InterviewRepository) List(ctx context.Context, limit int, t time.Time) ([]*entity.Interview, error) {
	panic("unimplemented")
}

func NewInterviewRepository(db *bun.DB) datagateway.Interview {
	return &InterviewRepository{
		db: db,
	}
}

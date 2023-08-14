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

// Archive implements datagateway.Interview.
func (r *InterviewRepository) Archive(ctx context.Context, id uuid.UUID) (*entity.Interview, error) {
	var model Interview
	if _, err := r.db.NewUpdate().
		Model(&model).
		Where("id = ?", id).
		SetColumn("is_archived", "?", true).
		SetColumn("updated_at", "?", time.Now().UTC()).
		Returning("*").
		Exec(ctx); err != nil {
		return nil, err
	}
	return model.ToEntity()
}

// UpdateStatus implements datagateway.Interview.
func (r *InterviewRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status entity.InterviewStatus) (*entity.Interview, error) {
	var model Interview
	newStatus, err := fromEntityInterviewStatus(status)
	if err != nil {
		return nil, err
	}
	if _, err := r.db.NewUpdate().
		Model(&model).
		Where("id = ?", id).
		SetColumn("status", "?", newStatus).
		SetColumn("updated_at", "?", time.Now().UTC()).
		Returning("*").
		Exec(ctx); err != nil {
		return nil, err
	}
	return model.ToEntity()
}

// Create implements datagateway.Interview.
func (r *InterviewRepository) Create(ctx context.Context, interview *entity.Interview) (*entity.Interview, error) {
	model, err := fromEntity(interview)
	if err != nil {
		return nil, err
	}
	if _, err := r.db.NewInsert().
		Model(model).
		ExcludeColumn("id", "created_at", "updated_at").
		Returning("*").
		Exec(ctx); err != nil {
		return nil, err
	}
	return model.ToEntity()
}

// Get implements datagateway.Interview.
func (r *InterviewRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Interview, error) {
	var model Interview
	if err := r.db.NewSelect().
		Model(&model).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, err
	}
	return model.ToEntity()
}

// List implements datagateway.Interview.
func (r *InterviewRepository) List(ctx context.Context, limit int, t time.Time) ([]*entity.Interview, error) {
	var interviewModels []Interview
	if err := r.db.NewSelect().
		Model(&interviewModels).
		Where("is_archived = false AND created_at > ?", t).
		Order("created_at ASC").
		Limit(limit).
		Scan(ctx); err != nil {
		return nil, err
	}
	var res []*entity.Interview
	for _, model := range interviewModels {
		e, err := model.ToEntity()
		if err != nil {
			return nil, err
		}
		res = append(res, e)
	}
	return res, nil
}

func NewInterviewRepository(db *bun.DB) datagateway.Interview {
	return &InterviewRepository{
		db: db,
	}
}

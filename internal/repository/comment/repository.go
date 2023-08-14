package repository_comment

import (
	"context"

	"github.com/chawin-a/robinhood-interview/internal/datagateway"
	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CommentRepository struct {
	db *bun.DB
}

// ListByInterviewId implements datagateway.Comment.
func (r *CommentRepository) ListByInterviewId(ctx context.Context, interviewId uuid.UUID) ([]*entity.Comment, error) {
	var models []Comment
	if err := r.db.NewSelect().
		Model(&models).
		Where("interview_id = ?", interviewId).
		Order("created_at DESC").
		Scan(ctx); err != nil {
		return nil, err
	}
	var res []*entity.Comment
	for _, model := range models {
		res = append(res, model.ToEntity())
	}
	return res, nil
}

// Create implements datagateway.Comment.
func (r *CommentRepository) Create(ctx context.Context, comment *entity.Comment) (*entity.Comment, error) {
	commentModel := fromEntity(comment)
	_, err := r.db.NewInsert().
		Model(commentModel).
		ExcludeColumn("id", "created_at").
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return commentModel.ToEntity(), nil
}

func NewCommentRepository(db *bun.DB) datagateway.Comment {
	return &CommentRepository{
		db: db,
	}
}

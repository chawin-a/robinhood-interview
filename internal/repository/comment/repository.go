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

// ListByUserId implements datagateway.Comment.
func (*CommentRepository) ListByUserId(ctx context.Context, userId uuid.UUID) ([]*entity.Comment, error) {
	panic("unimplemented")
}

func NewCommentRepository(db *bun.DB) datagateway.Comment {
	return &CommentRepository{
		db: db,
	}
}

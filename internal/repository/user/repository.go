package repository_user

import (
	"context"

	"github.com/chawin-a/robinhood-interview/internal/datagateway"
	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

// Create implements datagateway.User.
func (r *UserRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	userModel := fromEntity(user)
	_, err := r.db.NewInsert().
		Model(userModel).
		ExcludeColumn("id", "created_at").
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return userModel.ToEntity(), nil
}

// Get implements datagateway.User.
func (*UserRepository) Get(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	panic("unimplemented")
}

// List implements datagateway.User.
func (*UserRepository) List(ctx context.Context, ids []uuid.UUID) ([]*entity.User, error) {
	panic("unimplemented")
}

func NewUserRepository(db *bun.DB) datagateway.User {
	return &UserRepository{
		db: db,
	}
}

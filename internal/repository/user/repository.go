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
func (r *UserRepository) Get(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var model User
	if err := r.db.NewSelect().
		Model(&model).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, err
	}
	return model.ToEntity(), nil
}

func NewUserRepository(db *bun.DB) datagateway.User {
	return &UserRepository{
		db: db,
	}
}

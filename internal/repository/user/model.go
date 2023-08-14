package repository_user

import (
	"time"

	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Username  string
	Name      string
	Email     string
	CreatedAt time.Time
}

func fromEntity(c *entity.User) *User {
	return &User{
		Id:        c.Id,
		Username:  c.Username,
		Name:      c.Name,
		Email:     c.Email,
		CreatedAt: c.CreatedAt,
	}
}

func (c *User) ToEntity() *entity.User {
	return &entity.User{
		Id:        c.Id,
		Username:  c.Username,
		Name:      c.Name,
		Email:     c.Email,
		CreatedAt: c.CreatedAt,
	}
}

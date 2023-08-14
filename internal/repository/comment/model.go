package repository_comment

import (
	"time"

	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/google/uuid"
)

type Comment struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	InterviewId uuid.UUID
	Comment     string
	CreatedAt   time.Time
}

func fromEntity(c *entity.Comment) *Comment {
	return &Comment{
		Id:          c.Id,
		UserId:      c.UserId,
		InterviewId: c.InterviewId,
		Comment:     c.Comment,
		CreatedAt:   c.CreatedAt,
	}
}

func (c *Comment) ToEntity() *entity.Comment {
	return &entity.Comment{
		Id:          c.Id,
		UserId:      c.UserId,
		InterviewId: c.InterviewId,
		Comment:     c.Comment,
		CreatedAt:   c.CreatedAt,
	}
}

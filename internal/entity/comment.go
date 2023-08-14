package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	InterviewId uuid.UUID
	Comment     string
	CreatedAt   time.Time
}

package entity

import (
	"time"

	"github.com/google/uuid"
)

type InterviewStatus string

const (
	TO_DO       InterviewStatus = "To Do"
	IN_PROGRESS                 = "In Progress"
	DONE                        = "Done"
)

type Interview struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	Description string
	Status      InterviewStatus
	IsArchived  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

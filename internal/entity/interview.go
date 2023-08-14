package entity

import (
	"time"

	"github.com/google/uuid"
)

type InterviewStatus int

const (
	TO_DO InterviewStatus = iota
	IN_PROGRESS
	DONE
)

func (s InterviewStatus) String() string {
	switch s {
	case TO_DO:
		return "To Do"
	case IN_PROGRESS:
		return "In Progress"
	case DONE:
		return "Done"
	}
	panic("invalid interview status")
}

type Interview struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	Description string
	Status      InterviewStatus
	IsArchived  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

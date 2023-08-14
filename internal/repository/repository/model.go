package repository_interview

import (
	"time"

	"github.com/chawin-a/robinhood-interview/internal/entity"
	"github.com/google/uuid"
)

type InterviewStatus string

const (
	TO_DO       InterviewStatus = "TO_DO"
	IN_PROGRESS InterviewStatus = "IN_PROGRESS"
	DONE        InterviewStatus = "DONE"
)

func (s InterviewStatus) ToEntityInterviewStatus() entity.InterviewStatus {
	switch s {
	case TO_DO:
		return entity.TO_DO
	case IN_PROGRESS:
		return entity.IN_PROGRESS
	case DONE:
		return entity.DONE
	}
	panic("invalid interview status from model to entity")
}

func fromEntityInterviewStatus(s entity.InterviewStatus) InterviewStatus {
	switch s {
	case entity.TO_DO:
		return TO_DO
	case entity.IN_PROGRESS:
		return IN_PROGRESS
	case entity.DONE:
		return DONE
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

func fromEntity(c *entity.Interview) *Interview {
	return &Interview{
		Id:          c.Id,
		UserId:      c.UserId,
		Description: c.Description,
		Status:      fromEntityInterviewStatus(c.Status),
		IsArchived:  c.IsArchived,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func (c *Interview) ToEntity() *entity.Interview {
	return &entity.Interview{
		Id:          c.Id,
		UserId:      c.UserId,
		Description: c.Description,
		Status:      c.Status.ToEntityInterviewStatus(),
		IsArchived:  c.IsArchived,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

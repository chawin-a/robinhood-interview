package repository_interview

import (
	"errors"
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

func (s InterviewStatus) ToEntityInterviewStatus() (entity.InterviewStatus, error) {
	switch s {
	case TO_DO:
		return entity.TO_DO, nil
	case IN_PROGRESS:
		return entity.IN_PROGRESS, nil
	case DONE:
		return entity.DONE, nil
	}
	return "", errors.New("invalid interview status from model to entity")
}

func fromEntityInterviewStatus(s entity.InterviewStatus) (InterviewStatus, error) {
	switch s {
	case entity.TO_DO:
		return TO_DO, nil
	case entity.IN_PROGRESS:
		return IN_PROGRESS, nil
	case entity.DONE:
		return DONE, nil
	}
	return "", errors.New("invalid interview status from entity to model")
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

func fromEntity(c *entity.Interview) (*Interview, error) {
	status, err := fromEntityInterviewStatus(c.Status)
	if err != nil {
		return nil, err
	}
	return &Interview{
		Id:          c.Id,
		UserId:      c.UserId,
		Description: c.Description,
		Status:      status,
		IsArchived:  c.IsArchived,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}, nil
}

func (c *Interview) ToEntity() (*entity.Interview, error) {
	status, err := c.Status.ToEntityInterviewStatus()
	if err != nil {
		return nil, err
	}
	return &entity.Interview{
		Id:          c.Id,
		UserId:      c.UserId,
		Description: c.Description,
		Status:      status,
		IsArchived:  c.IsArchived,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}, nil
}

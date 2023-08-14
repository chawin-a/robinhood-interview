package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Username  string
	Name      string
	Email     string
	CreatedAt time.Time
}

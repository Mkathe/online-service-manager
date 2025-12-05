package model

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Id        uuid.UUID
	Name      string
	Price     int
	UserId    uuid.UUID
	StartDate time.Time
	EndDate   time.Time
}

package model

import (
	"time"

	"github.com/google/uuid"
)

type TotalCostFilter struct {
	From        time.Time
	To          time.Time
	UserID      uuid.UUID
	ServiceName string
}

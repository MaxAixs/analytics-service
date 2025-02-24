package analytics

import (
	"github.com/google/uuid"
	"time"
)

type TaskModel struct {
	UserId    uuid.UUID
	Email     string
	ItemId    int
	CreatedAt time.Time
}

type CompletedTaskModel struct {
	UserId uuid.UUID `json:"userId"`
	Email  string    `json:"email"`
	Count  int32     `json:"count"`
}

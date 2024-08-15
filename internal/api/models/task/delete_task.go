package task

import (
	"github.com/google/uuid"
	"time"
)

type RequestDeleteTask struct {
	ID   *uuid.UUID `json:"id"`
	Name *string    `json:"name"`
}

type ResponseDeleteTask struct {
	ID      uuid.UUID `json:"id"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

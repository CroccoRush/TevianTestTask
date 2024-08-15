package task

import (
	"github.com/google/uuid"
	"time"
)

type RequestProcessTask struct {
	ID   *uuid.UUID `json:"id,omitempty"`
	Name *string    `json:"name,omitempty"`
}

type ResponseProcessTask struct {
	ID      uuid.UUID `json:"id"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

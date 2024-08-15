package image

import (
	"github.com/google/uuid"
	"time"
)

type RequestUploadImage struct {
	TaskID uuid.UUID `json:"task_id"`
	Name   string    `json:"name"`
}

type ResponseUploadImage struct {
	ID      uuid.UUID `json:"id"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

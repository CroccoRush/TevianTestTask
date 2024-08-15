package models

import (
	"github.com/google/uuid"
	"log"
	"time"
)

type RequestCommon struct {
	ID   *uuid.UUID `json:"id"`
	Name *string    `json:"name"`
}

type ResponseCommon struct {
	ID      uuid.UUID `json:"id"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

func (r *ResponseCommon) Create(message string) {

	UUID, err := uuid.NewUUID()
	if err != nil {
		log.Print(err)
	}
	r.ID = UUID
	r.Time = time.Now()
	r.Message = message
}

type ResponseError struct {
	ResponseCommon

	Error string `json:"error"`
}

func (r *ResponseError) Create(message string, error string) {

	r.ResponseCommon.Create(message)
	r.Error = error
}

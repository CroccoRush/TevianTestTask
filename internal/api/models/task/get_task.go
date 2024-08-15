package task

import "github.com/google/uuid"

type RequestGetTask struct {
	ID   *uuid.UUID `json:"id,omitempty"`
	Name *string    `json:"name,omitempty"`
}

type ResponseGetTask struct {
	ID        uuid.UUID   `json:"id"`
	Status    Status      `json:"status"`
	Images    []ImageData `json:"images"`
	Statistic Statistic   `json:"statistic"`
}

type Status string

const (
	Forming    Status = "forming"
	Processing        = "processing"
	Completed         = "completed"
	Error             = "error"
)

type ImageData struct {
	Name  string     `json:"name"`
	Faces []FaceData `json:"faces"`
}

type FaceData struct {
	BoundingBox string `json:"bounding_box"`
	Sex         Sex    `json:"sex"`
	Age         int    `json:"age"`
}

type Sex string

const (
	Male   Sex = "male"
	Female     = "female"
)

type Statistic struct {
	FaceCount        int `json:"face_count"`
	MaleCount        int `json:"male_count"`
	FemaleCount      int `json:"female_count"`
	AverageMaleAge   int `json:"avg_male_age"`
	AverageFemaleAge int `json:"avg_female_age"`
}

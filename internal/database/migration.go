package database

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID        uuid.UUID   `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time   `gorm:"" json:"-"`
	UpdatedAt time.Time   `gorm:"" json:"-"`
	Name      string      `gorm:"unique" json:"name"`
	Status    Status      `gorm:"type:status" json:"status"`
	Images    []ImageData `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE" json:"images"`
	Statistic Statistic   `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE" json:"statistic"`
}

type Status string

const (
	Forming    Status = "forming"
	Processing        = "processing"
	Completed         = "completed"
	Error             = "error"
)

type ImageData struct {
	ID        uuid.UUID   `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time   `gorm:"" json:"-"`
	UpdatedAt time.Time   `gorm:"" json:"-"`
	TaskID    uuid.UUID   `gorm:"uniqueIndex:idx_task_name" json:"-"`
	Name      string      `gorm:"uniqueIndex:idx_task_name" json:"name"`
	Status    ImageStatus `gorm:"type:image_status;index" json:"status"`
	Faces     *[]FaceData `gorm:"foreignKey:ImageID;constraint:OnDelete:CASCADE" json:"faces,omitempty"`
}

type ImageStatus string

const (
	ImageUntouched ImageStatus = "untouched"
	ImageProcessed             = "processed"
	ImageError                 = "error"
)

type FaceData struct {
	ID          uuid.UUID `gorm:"primaryKey" json:"-"`
	CreatedAt   time.Time `gorm:"" json:"-"`
	UpdatedAt   time.Time `gorm:"" json:"-"`
	ImageID     uuid.UUID `gorm:"" json:"-"`
	BoundingBox string    `gorm:"" json:"bounding_box"`
	Sex         Sex       `gorm:"type:sex" json:"sex"`
	Age         float64   `gorm:"" json:"age"`
}

type Sex string

const (
	Male   Sex = "male"
	Female     = "female"
)

type Statistic struct {
	ID               uuid.UUID `gorm:"primaryKey" json:"-"`
	CreatedAt        time.Time `gorm:"" json:"-"`
	UpdatedAt        time.Time `gorm:"" json:"-"`
	TaskID           uuid.UUID `gorm:"" json:"-"`
	FaceCount        int       `json:"face_count"`
	MaleCount        int       `json:"male_count"`
	FemaleCount      int       `json:"female_count"`
	AverageMaleAge   float64   `json:"avg_male_age"`
	AverageFemaleAge float64   `json:"avg_female_age"`
}

func (class Task) GetID() uuid.UUID {
	return class.ID
}

func (class ImageData) GetID() uuid.UUID {
	return class.ID
}

func (class FaceData) GetID() uuid.UUID {
	return class.ID
}

func (class Statistic) GetID() uuid.UUID {
	return class.ID
}

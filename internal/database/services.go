package database

import (
	iError "TevianTestTask/internal/internal_errors"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"time"
)

func (db *DBServ) NewTask(name string) (task *Task, err error) {

	UUID, err := uuid.NewUUID()
	if err != nil {
		err = errors.Wrap(err, "UUID was not created successfully")
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	task = &Task{
		ID:        UUID,
		CreatedAt: time.Now(),
		Name:      name,
		Status:    Forming,
	}

	if err = db.Save(task).Error; err != nil {
		err = errors.Wrap(err, "saving task failed")
		if errors.As(err, &gorm.ErrDuplicatedKey) {
			err = errors.Wrap(iError.ErrDuplicateKey, err.Error())
		} else {
			err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		}
		return nil, err
	}

	statistic, err := db.NewStatistic(UUID)
	if err != nil {
		err = errors.Wrapf(err, "statistic for Task %s --- %s was not created successfully", name, UUID)
	}

	task.Statistic = *statistic

	return
}

func (db *DBServ) NewStatistic(taskID uuid.UUID) (statistic *Statistic, err error) {

	UUID, err := uuid.NewUUID()
	if err != nil {
		err = errors.Wrap(err, "UUID was not created successfully")
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	statistic = &Statistic{
		ID:               UUID,
		CreatedAt:        time.Now(),
		TaskID:           taskID,
		FaceCount:        0,
		MaleCount:        0,
		FemaleCount:      0,
		AverageMaleAge:   0,
		AverageFemaleAge: 0,
	}

	if err = db.Save(statistic).Error; err != nil {
		err = errors.Wrap(err, "saving statistic failed")
		err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		return nil, err
	}

	return statistic, nil
}

func (db *DBServ) NewImage(taskID uuid.UUID, name string) (image *ImageData, err error) {

	UUID, err := uuid.NewUUID()
	if err != nil {
		err = errors.Wrap(err, "UUID was not created successfully")
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	image = &ImageData{
		ID:        UUID,
		CreatedAt: time.Now(),
		TaskID:    taskID,
		Name:      name,
		Status:    ImageUntouched,
	}

	if err = db.Save(image).Error; err != nil {
		err = errors.Wrap(err, "saving image failed")
		if errors.As(err, &gorm.ErrDuplicatedKey) {
			err = errors.Wrap(iError.ErrDuplicateKey, err.Error())
		} else {
			err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		}
		return nil, err
	}

	return
}

func (db *DBServ) NewFace(imageID uuid.UUID) (face *FaceData, err error) {

	UUID, err := uuid.NewUUID()
	if err != nil {
		err = errors.Wrap(err, "UUID was not created successfully")
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	face = &FaceData{
		ID:        UUID,
		CreatedAt: time.Now(),
		ImageID:   imageID,
	}

	if err = db.Save(face).Error; err != nil {
		err = errors.Wrap(err, "saving face failed")
		err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		return nil, err
	}

	return
}

func (db *DBServ) AddFace(imageID uuid.UUID, bbox string, age float64, sex string) (face *FaceData, err error) {

	UUID, err := uuid.NewUUID()
	if err != nil {
		err = errors.Wrap(err, "UUID was not created successfully")
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	face = &FaceData{
		ID:          UUID,
		CreatedAt:   time.Now(),
		ImageID:     imageID,
		BoundingBox: bbox,
		Age:         age,
		Sex:         Sex(sex),
	}

	if err = db.Save(&face).Error; err != nil {
		err = errors.Wrap(err, "saving face failed")
		err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		return nil, err
	}

	return
}

func (db *DBServ) DeleteTask(id *uuid.UUID, name *string) (task *Task, err error) {

	task = &Task{}
	if id != nil {
		task.ID = *id
	} else if name != nil {
		task.Name = *name
	} else {
		err = errors.New("no ID or name is given")
		return nil, errors.Wrap(iError.ErrInvalidParams, err.Error())
	}

	err = db.Delete(task).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.Wrap(iError.ErrNotFound, err.Error())
		return nil, err
	} else if err != nil {
		err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		return nil, err
	}

	return
}

func (db *DBServ) GetTask(id *uuid.UUID, name *string) (task *Task, err error) {

	task = &Task{}
	if id != nil {
		task.ID = *id
	} else if name != nil {
		task.Name = *name
	} else {
		err = errors.New("no ID or name is given")
		return nil, errors.Wrap(iError.ErrInvalidParams, err.Error())
	}

	err = db.Where("id = ? or name = ?", task.ID, task.Name).First(task).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.Wrap(iError.ErrNotFound, err.Error())
		return nil, err
	} else if err != nil {
		err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		return nil, err
	}

	return
}

func (db *DBServ) GetFullTask(id *uuid.UUID, name *string) (task *Task, err error) {

	task, err = db.GetTask(id, name)
	if err != nil {
		err = errors.Wrapf(err, "failed to find task with id %v or name %v", id, name)
		return nil, err
	}

	statistic, err := db.GetTaskStatistic(&task.ID)
	if err != nil {
		err = errors.Wrapf(err, "failed to get statistic for task with id %s or name %s", task.ID, task.Name)
		return nil, err
	}
	task.Statistic = *statistic

	images, err := db.GetTaskFullImages(&task.ID)
	if err != nil {
		err = errors.Wrapf(err, "failed to get images for task with id %s or name %s", task.ID, task.Name)
		return nil, err
	}
	task.Images = *images

	return

}

func (db *DBServ) GetTaskStatistic(taskID *uuid.UUID) (statistic *Statistic, err error) {

	statistic = &Statistic{}
	if taskID != nil {
		statistic.TaskID = *taskID
	} else {
		err = errors.New("no taskID is given")
		err = errors.Wrap(iError.ErrInvalidParams, err.Error())
		return nil, err
	}

	err = db.Where("task_id = ?", statistic.TaskID).First(&statistic).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.Wrap(iError.ErrNotFound, err.Error())
		return nil, err
	} else if err != nil {
		err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		return nil, err
	}

	return
}

func (db *DBServ) GetStatistic(id *uuid.UUID) (statistic *Statistic, err error) {

	statistic = &Statistic{}
	if id != nil {
		statistic.ID = *id
	} else {
		err = errors.New("no ID is given")
		err = errors.Wrap(iError.ErrInvalidParams, err.Error())
		return nil, err
	}

	err = db.Where("is = ?", statistic.ID).First(&statistic).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.Wrap(iError.ErrNotFound, err.Error())
		return nil, err
	} else if err != nil {
		err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		return nil, err
	}

	return
}

func (db *DBServ) ProcessStatistic(taskID *uuid.UUID) (err error) {

	var statistic, oldStatistic *Statistic

	oldStatistic, err = db.GetTaskStatistic(taskID)
	if err != nil {
		err = errors.Wrap(err, "failed to get statistic")
		return
	}

	err = db.Raw(`
		SELECT
			tasks.id AS task_id,
			COUNT(face_data.id) AS face_count,
			SUM(CASE WHEN face_data.sex = 'male' THEN 1 ELSE 0 END) AS male_count,
			SUM(CASE WHEN face_data.sex = 'female' THEN 1 ELSE 0 END) AS female_count,
			AVG(CASE WHEN face_data.sex = 'male' THEN face_data.age END) AS average_male_age,
			AVG(CASE WHEN face_data.sex = 'female' THEN face_data.age END) AS average_female_age
		FROM
			tasks
				LEFT JOIN
			image_data ON tasks.id = image_data.task_id
				LEFT JOIN
			face_data ON image_data.id = face_data.image_id
		WHERE
			tasks.id = ?
		GROUP BY
			tasks.id
    `, *taskID).Scan(&statistic).Error
	if err != nil {
		err = errors.Wrap(err, "failed to process statistic")
		err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		return
	}

	statistic.ID = oldStatistic.ID
	statistic.CreatedAt = oldStatistic.CreatedAt
	if err = db.Save(&statistic).Error; err != nil {
		err = errors.Wrap(err, "failed to save processed statistic")
		if errors.As(err, &gorm.ErrDuplicatedKey) {
			err = errors.Wrap(iError.ErrDuplicateKey, err.Error())
		} else {
			err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		}
		return
	}

	return
}

func (db *DBServ) GetTaskImages(taskID *uuid.UUID, all bool) (images *[]ImageData, err error) {

	images = &[]ImageData{}
	if taskID == nil {
		err = errors.New("no taskID is given")
		err = errors.Wrap(iError.ErrInvalidParams, err.Error())
		return nil, err
	}

	if all {
		err = db.Where("task_id = ?", *taskID).Find(images).Error
	} else {
		err = db.Where("task_id = ? and status != ?", *taskID, ImageProcessed).Find(images).Error
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.Wrap(iError.ErrNotFound, err.Error())
		return nil, err
	} else if err != nil {
		err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		return nil, err
	}

	return
}

func (db *DBServ) GetTaskFullImages(taskID *uuid.UUID) (images *[]ImageData, err error) {

	images, err = db.GetTaskImages(taskID, true)
	if err != nil {
		err = errors.Wrapf(err, "failed to find image with task_id %v", *taskID)
		return nil, err
	}

	for i, image := range *images {
		faces, err := db.GetImageFaces(&image.ID)
		if err != nil {
			err = errors.Wrapf(err, "failed to get faces for image with id %s", image.ID)
			log.Println(err)
			continue
		}
		(*images)[i].Faces = faces
	}

	return
}

func (db *DBServ) GetImage(id *uuid.UUID, name *string) (image *ImageData, err error) {

	image = &ImageData{}
	if id != nil {
		image.ID = *id
	} else if name != nil {
		image.Name = *name
	} else {
		err = errors.New("no ID or name is given")
		err = errors.Wrap(iError.ErrInvalidParams, err.Error())
		return nil, err
	}

	err = db.Where("id = ? or name = ?", image.ID, image.Name).First(image).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.Wrap(iError.ErrNotFound, err.Error())
		return nil, err
	} else if err != nil {
		err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		return nil, err
	}

	return
}

func (db *DBServ) GetImageFaces(imageID *uuid.UUID) (faces *[]FaceData, err error) {

	faces = &[]FaceData{}
	if imageID == nil {
		err = errors.New("no imageID is given")
		err = errors.Wrap(iError.ErrInvalidParams, err.Error())
		return nil, err
	}

	err = db.Where("image_id = ?", *imageID).Find(&faces).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.Wrap(iError.ErrNotFound, err.Error())
		return nil, err
	} else if err != nil {
		err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		return nil, err
	}

	return
}

func (db *DBServ) GetFace(id *uuid.UUID) (face *FaceData, err error) {

	face = &FaceData{}
	if id != nil {
		face.ID = *id
	} else {
		err = errors.New("no ID is given")
		err = errors.Wrap(iError.ErrInvalidParams, err.Error())
		return
	}

	err = db.Where("is = ?", face.ID).First(&face).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.Wrap(iError.ErrNotFound, err.Error())
		return nil, err
	} else if err != nil {
		err = errors.Wrap(iError.ErrUnexpectedDB, err.Error())
		return nil, err
	}

	return
}

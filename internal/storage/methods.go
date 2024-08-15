package storage

import (
	iError "TevianTestTask/internal/internal_errors"
	"TevianTestTask/pkg/storage_manager"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io"
	"log"
	"mime/multipart"
)

func DeleteTask(taskID *uuid.UUID) (err error) {

	path, err := storage_manager.DeletePath(storagePath, taskID.String())
	if err != nil {
		err = errors.Wrapf(err, "failed to remove directory %s", path)
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	log.Printf("Directory %s removed successfully", taskID)
	return
}

func CreateTask(taskID *uuid.UUID) (err error) {

	path, err := storage_manager.InitPath(storagePath, taskID.String())
	if err != nil {
		err = errors.Wrapf(err, "failed to create directory %s", path)
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	log.Printf("Directory %s created successfully", taskID)

	return
}

func UploadImages(taskID *uuid.UUID) (err error) {

	path, err := storage_manager.InitPath(storagePath, taskID.String())
	if err != nil {
		err = errors.Wrapf(err, "failed to remove directory %s", path)
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	log.Printf("Directory %s removed successfully", taskID)

	return
}

func UploadImage(taskID, imageID *uuid.UUID, image *multipart.File) (err error) {

	dst, err := storage_manager.Create(storagePath, taskID.String(), imageID.String())
	if err != nil {
		err = errors.Wrapf(err, "failed to init file %s", imageID)
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, *image); err != nil {
		err = errors.Wrapf(err, "failed to copy file %s", imageID)
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	log.Printf("Image %s upload successfully", imageID)

	return
}

func FindImage(taskID, imageID *uuid.UUID) (path string, err error) {

	path, err = storage_manager.Find(storagePath, taskID.String(), imageID.String())
	if err != nil {
		err = errors.Wrap(iError.ErrNotFound, err.Error())
	}

	return
}

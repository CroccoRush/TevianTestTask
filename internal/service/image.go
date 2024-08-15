package service

import (
	"TevianTestTask/internal/api/models/image"
	"TevianTestTask/internal/database"
	iError "TevianTestTask/internal/internal_errors"
	"TevianTestTask/internal/storage"
	"encoding/json"
	"github.com/pkg/errors"
	"mime/multipart"
	"time"
)

func UploadImage(form *multipart.Form) (rawBody []byte, err error) {

	metaData := form.Value["meta_data"]
	if len(metaData) == 0 {
		err = errors.New("no meta data")
		err = errors.Wrap(iError.ErrInvalidParams, err.Error())
		return
	}
	reqBody := metaData[0]

	images := form.File["image"]
	if len(images) == 0 {
		err = errors.New("no image upload")
		err = errors.Wrap(iError.ErrInvalidParams, err.Error())
		return
	}
	reqImage := images[0]

	reqImageMeta := new(image.RequestUploadImage)

	if err = json.Unmarshal([]byte(reqBody), reqImageMeta); err != nil {
		err = errors.Wrap(iError.ErrInvalidJson, err.Error())
		return
	}

	imagesTask, err := database.DB.GetTask(&reqImageMeta.TaskID, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to get task")
		return
	}

	if imagesTask.Status != database.Forming {
		err = errors.New("images task is not forming")
		err = errors.Wrap(iError.ErrLocked, err.Error())
		return
	}

	// Открываем файл
	rawImage, err := reqImage.Open()
	if err != nil {
		err = errors.Wrap(err, "failed to open file")
		err = errors.Wrap(iError.ErrInvalidParams, err.Error())
		return
	}
	defer rawImage.Close()

	imageData, err := database.DB.NewImage(reqImageMeta.TaskID, reqImageMeta.Name)
	if err != nil {
		err = errors.Wrap(err, "failed to save image data")
		return
	}

	err = storage.UploadImage(&imageData.TaskID, &imageData.ID, &rawImage)
	if err != nil {
		err = errors.Wrap(err, "failed to save image")
		return
	}

	resBody := image.ResponseUploadImage{
		ID:      imageData.ID,
		Time:    time.Now(),
		Message: "The image has been saved",
	}
	rawBody, err = json.Marshal(resBody)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal response")
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	return
}

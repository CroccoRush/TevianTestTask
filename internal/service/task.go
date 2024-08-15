package service

import (
	"TevianTestTask/internal/api/models/task"
	"TevianTestTask/internal/database"
	"TevianTestTask/internal/facecloud"
	iError "TevianTestTask/internal/internal_errors"
	"TevianTestTask/internal/storage"
	retry "TevianTestTask/pkg/retry_function"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log"
	"sync"
	"time"
)

func AddTask(reqBody []byte) (rawBody []byte, err error) {

	reqTask := new(task.RequestAddTask)

	if err = json.Unmarshal(reqBody, reqTask); err != nil {
		err = errors.Wrap(iError.ErrInvalidJson, err.Error())
		return
	}

	resTask, err := database.DB.NewTask(reqTask.Name)
	if err != nil {
		err = errors.Wrap(err, "failed to create task")
		return
	}

	if err = storage.CreateTask(&resTask.ID); err != nil {
		err = errors.Wrap(err, "failed to create task directory")
		return
	}

	rawBody, err = json.Marshal(resTask)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal task")
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	return
}

func GetTask(reqBody []byte) (rawBody []byte, err error) {

	reqTask := new(task.RequestGetTask)

	if err = json.Unmarshal(reqBody, reqTask); err != nil {
		err = errors.Wrap(iError.ErrInvalidJson, err.Error())
		return
	}

	resTask, err := database.DB.GetFullTask(reqTask.ID, reqTask.Name)
	if err != nil {
		err = errors.Wrap(err, "failed to get task")
		return
	}

	rawBody, err = json.Marshal(resTask)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal task")
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	return
}

func ProcessTask(reqBody []byte) (rawBody []byte, err error) {

	reqTask := new(task.RequestProcessTask)

	if err = json.Unmarshal(reqBody, reqTask); err != nil {
		err = errors.Wrap(iError.ErrInvalidJson, err.Error())
		return
	}

	processingTask, err := database.DB.GetTask(reqTask.ID, reqTask.Name)
	if err != nil {
		err = errors.Wrap(err, "failed to get task")
		return
	}
	if processingTask.Status == database.Completed {
		resBody := task.ResponseProcessTask{
			ID:      processingTask.ID,
			Time:    time.Now(),
			Message: "The images have already been processed",
		}
		rawBody, err = json.Marshal(resBody)
		err = iError.ErrAlreadyProcess
		return
	}

	processingTask.Status = database.Processing
	database.DB.Save(processingTask)

	go processImages(&processingTask.ID)

	resBody := task.ResponseProcessTask{
		ID:      processingTask.ID,
		Time:    time.Now(),
		Message: "The images have been sent for processing",
	}
	rawBody, err = json.Marshal(resBody)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal response")
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	return
}

func processImages(taskID *uuid.UUID) {
	imagesTask, err := database.DB.GetTask(taskID, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to get task")
		log.Println(err)
		return
	}

	images, err := database.DB.GetTaskImages(taskID, false)
	if err != nil {
		err = errors.Wrap(err, "failed to get images")
		log.Println(err)
		return
	}

	success := true

	wg := new(sync.WaitGroup)
	wg.Add(len(*images))
	for _, image := range *images {
		go func(image *database.ImageData) {

			log.Printf("image #%s: send to process", image.ID.String())

			err = processImage(image, wg)
			if err != nil {
				image.Status = database.ImageError
				success = false
				log.Printf("image #%s: %s", image.ID.String(), err)
			} else {
				image.Status = database.ImageProcessed
				log.Printf("image #%s: complete the process", image.ID.String())
			}

			database.DB.Save(image)

		}(&image)
	}
	wg.Wait()

	err = processStatistic(taskID)
	if err != nil {
		err = errors.Wrap(err, "failed to update statistic")
		log.Println(err)
	}

	if success {
		imagesTask.Status = database.Completed
	} else {
		imagesTask.Status = database.Error
	}
	database.DB.Save(imagesTask)
}

func processImage(image *database.ImageData, wg *sync.WaitGroup) (err error) {
	defer wg.Done()

	path, err := storage.FindImage(&image.TaskID, &image.ID)
	if err != nil {
		err = errors.Wrap(err, "failed to find image in storage")
		return
	}

	response, err := facecloud.Detect(path)
	if err != nil {
		err = errors.Wrap(err, "failed to facecloud detect image")
		return
	}

	for _, data := range response.Data {
		_, err = database.DB.AddFace(image.ID, data.BBox.String(), data.Demographics.Age.Mean, data.Demographics.Gender)
		if err != nil {
			return errors.Wrap(err, "failed to add face")
		}
	}

	return
}

func processStatistic(taskID *uuid.UUID) (err error) {

	_, err = retry.RetryFunc(database.DB.ProcessStatistic, []interface{}{taskID}, 3, 1*time.Second)
	if err != nil {
		err = errors.Wrap(err, "failed to process statistic")
	}

	return
}

func DeleteTask(reqBody []byte) (rawBody []byte, err error) {

	reqTask := new(task.RequestDeleteTask)

	if err = json.Unmarshal(reqBody, reqTask); err != nil {
		err = errors.Wrap(iError.ErrInvalidJson, err.Error())
		return
	}

	resTask, err := database.DB.GetTask(reqTask.ID, reqTask.Name)
	if err != nil {
		err = errors.Wrap(err, "failed to get task")
		return
	}

	if resTask.Status == database.Processing {
		resBody := task.ResponseProcessTask{
			ID:      *reqTask.ID,
			Time:    time.Now(),
			Message: "The task is in processing, it is forbidden to delete it",
		}
		rawBody, err = json.Marshal(resBody)
		if err != nil {
			err = errors.Wrap(err, "failed to marshal response")
			err = errors.Wrap(iError.ErrInternal, err.Error())
			return
		}

		err = errors.New("The task is in processing, it is forbidden to delete it")
		err = errors.Wrap(iError.ErrLocked, err.Error())
		return
	}

	_, err = database.DB.DeleteTask(reqTask.ID, reqTask.Name)
	if err != nil {
		err = errors.Wrap(err, "failed to delete task from database")
		return
	}

	err = storage.DeleteTask(reqTask.ID)
	if err != nil {
		err = errors.Wrap(err, "failed to delete task from local storage")
		return
	}

	resBody := task.ResponseProcessTask{
		ID:      *reqTask.ID,
		Time:    time.Now(),
		Message: "The task has been completely deleted",
	}
	rawBody, err = json.Marshal(resBody)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal response")
		err = errors.Wrap(iError.ErrInternal, err.Error())
		return
	}

	return
}

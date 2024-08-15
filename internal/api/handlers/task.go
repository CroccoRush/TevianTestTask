package handlers

import (
	"TevianTestTask/internal/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

// AddTask godoc
// @Summary      Add task
// @Description  Creates a task with the passed name
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        request   body     task.RequestAddTask  false "body"
// @Success      200     {object}   task.ResponseAddTask
// @Failure      400 	 {object} 	models.ResponseError
// @Failure      404 	 {object} 	models.ResponseError
// @Failure      409 	 {object} 	models.ResponseError
// @Failure      500 	 {object} 	models.ResponseError
// @Router       /api/task/add      [post]
func AddTask(c *fiber.Ctx) (err error) {

	var res []byte
	status := fiber.StatusOK

	res, err = service.AddTask(c.Body())
	if err != nil {
		log.Print(err)
		res, status = ErrorHandler(err)
	}

	return c.Status(status).Send(res)
}

// GetTask godoc
// @Summary      Get task
// @Description  Returns a task with statistics and a list of images and their data on recognized faces
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        request   body     task.RequestGetTask  false "body"
// @Success      200     {object}   task.ResponseGetTask
// @Failure      400 	 {object} 	models.ResponseError
// @Failure      404 	 {object} 	models.ResponseError
// @Failure      409 	 {object} 	models.ResponseError
// @Failure      500 	 {object} 	models.ResponseError
// @Router       /api/task          [get]
func GetTask(c *fiber.Ctx) (err error) {

	var res []byte
	status := fiber.StatusOK

	res, err = service.GetTask(c.Body())
	if err != nil {
		log.Print(err)
		res, status = ErrorHandler(err)
	}

	return c.Status(status).Send(res)
}

// ProcessTask godoc
// @Summary      Process task
// @Description  Processes images from the task and calculates statistics for the task
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        request   body     task.RequestProcessTask  false "body"
// @Success      202     {object}   task.ResponseProcessTask
// @Success      200     {object}   task.ResponseProcessTask
// @Failure      400 	 {object} 	models.ResponseError
// @Failure      404 	 {object} 	models.ResponseError
// @Failure      409 	 {object} 	models.ResponseError
// @Failure      500 	 {object} 	models.ResponseError
// @Router       /api/task/process  [post]
func ProcessTask(c *fiber.Ctx) (err error) {

	var res []byte
	status := fiber.StatusAccepted

	res, err = service.ProcessTask(c.Body())
	if err != nil {
		log.Print(err)
		res, status = ErrorHandler(err)
	}

	return c.Status(status).Send(res)
}

// DeleteTask godoc
// @Summary      Delete task
// @Description  Delete task by ID
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        request   body     task.RequestDeleteTask  false "body"
// @Success      200     {object}   task.ResponseDeleteTask
// @Failure      400 	 {object} 	models.ResponseError
// @Failure      404 	 {object} 	models.ResponseError
// @Failure      423 	 {object} 	models.ResponseError
// @Failure      500 	 {object} 	models.ResponseError
// @Router       /api/task          [delete]
func DeleteTask(c *fiber.Ctx) (err error) {

	var res []byte
	status := fiber.StatusOK

	res, err = service.DeleteTask(c.Body())
	if err != nil {
		log.Print(err)
		res, status = ErrorHandler(err)
	}

	return c.Status(status).Send(res)
}

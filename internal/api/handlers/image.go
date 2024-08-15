package handlers

import (
	"TevianTestTask/internal/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

// UploadImage godoc
// @Summary      Upload image
// @Description  Uploads the image to the task
// @Tags         image
// @Accept       mpfd
// @Produce      json
// @Param        meta_data formData     image.RequestUploadImage  false "body"
// @Param        file 	   formData     file 					  false "Uploaded image"
// @Success      200       {object}     image.ResponseUploadImage
// @Failure      400 	   {object} 	models.ResponseError
// @Failure      409 	   {object} 	models.ResponseError
// @Failure      423 	   {object} 	models.ResponseError
// @Failure      500 	   {object} 	models.ResponseError
// @Router       /api/image/upload [POST]
func UploadImage(c *fiber.Ctx) (err error) {

	var res []byte
	status := fiber.StatusOK

	data, err := c.MultipartForm()
	if err != nil {
		log.Print(err)
		res, status = ErrorHandler(err)
		return c.Status(status).Send(res)
	}

	res, err = service.UploadImage(data)
	if err != nil {
		log.Print(err)
		res, status = ErrorHandler(err)
	}

	return c.Status(status).Send(res)
}

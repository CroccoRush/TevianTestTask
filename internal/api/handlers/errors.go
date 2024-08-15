package handlers

import (
	"TevianTestTask/internal/api/models"
	iError "TevianTestTask/internal/internal_errors"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func ErrorHandler(err error) (result []byte, status int) {

	if errors.Is(err, iError.ErrAlreadyProcess) {

		status = fiber.StatusOK

		res := models.ResponseCommon{}
		res.Create("All Images have already been processed")
		result, err = json.Marshal(res)

		return
	}

	if errors.Is(err, iError.ErrNotFound) ||
		errors.Is(err, iError.ErrInvalidParams) ||
		errors.Is(err, iError.ErrInvalidJson) {

		status = fiber.StatusBadRequest

	} else if errors.Is(err, iError.ErrDuplicateKey) {

		status = fiber.StatusConflict

	} else if errors.Is(err, iError.ErrLocked) {

		status = fiber.StatusLocked

	} else {

		status = fiber.StatusInternalServerError

	}

	res := models.ResponseError{}
	res.Create("", err.Error())
	result, err = json.Marshal(res)

	return
}

package api

import (
	"backend-test/constants"
	"backend-test/helpers"
	"backend-test/internal/interfaces"
	"backend-test/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IUserHandler struct {
	UserService interfaces.IUserService
}

func (api *IUserHandler) SendMail(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	req := models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	data, token, err := api.UserService.SubmitEmail(c, req)
	if err != nil {
		log.Error("failed to send email: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	helpers.SendResponseWithToken(c, http.StatusOK, constants.SuccessMessage, token, data)
}

func (api *IUserHandler) CekSessionByUUID(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	uuidStr := c.Query("uuid")
	if uuidStr == "" {
		log.Error("UUID tidak boleh kosong")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, "UUID tidak boleh kosong")
		return
	}

	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		log.Error("UUID tidak valid: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, "UUID tidak valid")
		return
	}

	data, err := api.UserService.CekSessionByUUID(c, parsedUUID)
	if err != nil {
		log.Error("Anda tidak memiliki token: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, "Anda tidak memiliki token")
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, data)
}

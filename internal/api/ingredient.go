package api

import (
	"backend-test/constants"
	"backend-test/helpers"
	"backend-test/internal/interfaces"
	"backend-test/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IngredientAPI struct {
	SvcIngredient interfaces.IIngredientService
}

func (api *IngredientAPI) InsertIngredient(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	req := models.Ingredient{}
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
	data, err := api.SvcIngredient.InsertIngredient(c, &req)
	if err != nil {
		log.Error("failed to create ingredient: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, data)
}

func (api *IngredientAPI) UpdateIngredient(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	req := models.Ingredient{}
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
	var inputID models.UriId
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		log.Error("id not valid: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	req.ID = uint(inputID.ID)
	data, err := api.SvcIngredient.UpdateIngredient(c, &req)
	if err != nil {
		log.Error("failed to update ingredient: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, data)
}

func (api *IngredientAPI) DeleteIngredient(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	var inputID models.UriId
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		log.Error("id not valid: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	err = api.SvcIngredient.DeleteIngredient(c, inputID.ID)
	if err != nil {
		log.Error("failed to delete ingredient: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, nil)
}

func (api *IngredientAPI) GetAllIngredient(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	param := c.Query("search")
	data, err := api.SvcIngredient.GetAllIngredient(c, param)
	if err != nil {
		log.Error("failed to delete ingredient: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, data)
}

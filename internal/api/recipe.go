package api

import (
	"backend-test/constants"
	"backend-test/helpers"
	"backend-test/internal/interfaces"
	"backend-test/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReceipeAPi struct {
	SvcRecipe interfaces.IRecipeService
}

func (api *ReceipeAPi) InsertRecipe(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	req := models.Recipe{}
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
	data, err := api.SvcRecipe.InsertRecipe(c, &req)
	if err != nil {
		log.Error("failed to create receipe: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, data)
}

func (api *ReceipeAPi) DeleteRecipe(c *gin.Context) {
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
	err = api.SvcRecipe.DeleteRecipe(c, inputID.ID)
	if err != nil {
		log.Error("failed to delete recipe: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, nil)
}

func (api *ReceipeAPi) GetAllRecipe(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	param := c.Query("search")
	data, err := api.SvcRecipe.GetAllRecipe(c, param)
	if err != nil {
		log.Error("failed to delete recipe: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, data)
}

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

func (api *ReceipeAPi) UpdateRecipe(c *gin.Context) {
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
	var inputID models.UriId
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		log.Error("id not valid: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}
	req.ID = uint(inputID.ID)
	data, err := api.SvcRecipe.UpdateRecipe(c, &req)
	if err != nil {
		log.Error("failed to update recipe: ", err)
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
	objComponent, _ := helpers.ComptServerSidePre(c)
	limit := objComponent.Limit
	if limit == 0 {
		limit = helpers.GetLimitData()
		objComponent.Limit = limit
	}

	count, err := api.SvcRecipe.CountData(c, objComponent)
	if err != nil {
		log.Error("failed to count data inventory: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}
	param := c.Query("search")
	data, err := api.SvcRecipe.GetAllRecipe(c, objComponent, param)
	if err != nil {
		log.Error("failed to get data recipe: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}
	if len(data) == 0 {
		log.Error("failed to get dat inventory: ", err)
		helpers.SendResponseHTTP(c, http.StatusOK, constants.ErrFailedBadRequest, err)
		return
	}
	response := helpers.APIResponseView("Succesfully Get Data!", http.StatusOK, "Succesfully", count, limit, data)
	response.Meta.CurrentPage = count
	c.JSON(http.StatusOK, response)
}

package api

import (
	"backend-test/constants"
	"backend-test/helpers"
	"backend-test/internal/interfaces"
	"backend-test/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InventoryAPI struct {
	InventoryService interfaces.IInventoryService
}

func (api *InventoryAPI) InsertInv(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	req := models.Inventory{}
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

	data, err := api.InventoryService.InsertInv(c, &req)
	if err != nil {
		log.Error("failed to create inventory: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, data)
}

func (api *InventoryAPI) UpdateInv(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	req := models.Inventory{}
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
	data, err := api.InventoryService.UpdateInv(c, &req)
	if err != nil {
		log.Error("failed to update inventory: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, data)
}

func (api *InventoryAPI) DeleteInv(c *gin.Context) {
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
	err = api.InventoryService.DeleteInv(c, inputID.ID)
	if err != nil {
		log.Error("failed to delete inventory: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, nil)
}

func (api *InventoryAPI) GetAllInv(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	objComponent, _ := helpers.ComptServerSidePre(c)
	limit := objComponent.Limit
	if limit == 0 {
		limit = helpers.GetLimitData()
		objComponent.Limit = limit
	}

	count, err := api.InventoryService.CountData(c, objComponent)
	if err != nil {
		log.Error("failed to count data inventory: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	data, err := api.InventoryService.GetAllInv(c, objComponent)
	if err != nil {
		log.Error("failed to delete inventory: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}
	if len(*data) == 0 {
		log.Error("failed to delete inventory: ", err)
		helpers.SendResponseHTTP(c, http.StatusOK, constants.ErrFailedBadRequest, err)
		return
	}
	response := helpers.APIResponseView("Succesfully Get Data!", http.StatusOK, "Succesfully", count, limit, data)
	response.Meta.CurrentPage = count
	c.JSON(http.StatusOK, response)
}

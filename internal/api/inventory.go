package api

import (
	"backend-test/constants"
	"backend-test/helpers"
	"backend-test/internal/interfaces"
	"backend-test/internal/models"
	"net/http"
	"os"

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

func (api *InventoryAPI) GetTemplate(c *gin.Context) {
	var (
		pathFile = "./uploads/movies.xlsx"
		fileName = "movies.xlsx"
	)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File(pathFile)
}

func (api *InventoryAPI) UploadExcel(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
		return
	}

	filePath := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	err = api.InventoryService.InsertFromExcel(c, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := os.Remove(filePath); err != nil {
		log.Println("Failed to delete file:", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "File processed successfully"})
}

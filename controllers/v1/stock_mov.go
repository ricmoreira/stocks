package controllers

import (
	"encoding/json"
	"stocks/models/request"
	"stocks/services"
	"stocks/util/errors"

	"github.com/gin-gonic/gin"
)

type (
	// StockMovController represents the controller for operating on the stocks resource
	StockMovController struct {
		StockMovService services.StockMovServiceContract
	}
)

// NewStockMovController is the constructor of StockMovController
func NewStockMovController(sms *services.StockMovService) *StockMovController {
	return &StockMovController{
		StockMovService: sms,
	}
}

// CreateAction creates a new stock movement
func (pc StockMovController) CreateAction(c *gin.Context) {
	req := mrequest.StockMovCreate{}
	json.NewDecoder(c.Request.Body).Decode(&req)

	e := errors.ValidateRequest(&req)
	if e != nil {
		c.JSON(e.HttpCode, e)
		return
	}

	res, err := pc.StockMovService.CreateOne(&req)

	if err != nil {
		c.JSON(err.HttpCode, err)
		return
	}

	c.JSON(200, res)
}

// ListAction list stock movements
func (pc StockMovController) ListAction(c *gin.Context) {
	validSorts := map[string]string{}
	validSorts["_id"] = "_id"

	validFilters := map[string]string{}
	validFilters["MovementType"] = "MovementType"
	validFilters["DocumentID"] = "DocumentID"
	validFilters["ProductCode"] = "ProductCode"
	validFilters["_id"] = "_id"

	qValues := c.Request.URL.Query()
	req := mrequest.NewListRequest(qValues, validSorts, validFilters)

	res, err := pc.StockMovService.List(req)

	if err != nil {
		c.JSON(err.HttpCode, err)
		return
	}

	c.JSON(200, res)
}

// ListStockMovCountAction aggregates stock results by ProductID, Dir and WharehouseID
func (pc StockMovController) ListStockMovCountAction(c *gin.Context) {
	validSorts := map[string]string{}
	validSorts["_id"] = "_id"

	validFilters := map[string]string{}
	validFilters["Dir"] = "Dir"
	validFilters["ProductCode"] = "ProductCode"
	validFilters["WharehouseID"] = "WharehouseID"
	validFilters["_id"] = "_id"

	qValues := c.Request.URL.Query()
	
	req := mrequest.NewListRequest(qValues, validSorts, validFilters)
	
	res, err := pc.StockMovService.ListStockMovCount(req)

	if err != nil {
		c.JSON(err.HttpCode, err)
		return
	}

	c.JSON(200, res)
}

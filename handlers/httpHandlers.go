package handlers

import (
	"stocks/util/errors"

	"stocks/models/response"

	"github.com/gin-gonic/gin"
)

// HttpHandlers provides generic http handlers
type HttpHandlers struct{}

// NewHttpHandlers is the HttpHandlers constructor
func NewHttpHandlers() *HttpHandlers {
	return &HttpHandlers{}
}

// NotFound responds to the client that the provided route does not exist
func (h *HttpHandlers) NotFound(c *gin.Context) {
	c.JSON(404, errors.HandleErrorResponse(errors.NOT_FOUND, []mresponse.ErrorDetail{}, ""))
}

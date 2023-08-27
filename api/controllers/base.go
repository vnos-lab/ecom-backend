package controller

import (
	"ecom/api/response"
	"ecom/api_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (b *BaseController) Response(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, response.SimpleResponse{
		Data:    data,
		Message: message,
	})
}

func (b *BaseController) ResponseList(c *gin.Context, message string, total *int64, data interface{}) {
	c.JSON(http.StatusOK, response.SimpleResponseList{
		Message: message,
		Data:    data,
		Total:   total,
	})
}

func (b *BaseController) ResponseError(c *gin.Context, err error) {
	status, ok := api_errors.MapErrorStatusCode[err.Error()]
	if !ok {
		status = http.StatusInternalServerError
	}

	c.AbortWithStatusJSON(status, response.ResponseError{
		Message: err.Error(),
		Error:   err,
	})
}

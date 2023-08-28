package controller

import (
	"ecom/api/response"
	"ecom/api_errors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	message := err.Error()
	ginType := gin.ErrorTypePublic
	if !ok {
		status = http.StatusInternalServerError
		ginType = gin.ErrorTypePrivate
		message = api_errors.ErrOcurred.Error()
	}

	c.Errors = append(c.Errors, &gin.Error{
		Err:  err,
		Type: ginType,
		Meta: status,
	})

	c.AbortWithStatusJSON(status, response.ResponseError{
		Message: message,
		Error:   err,
	})
}

func (b *BaseController) ResponseValidationError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		err = errors.New(ve[0].Field() + " is " + ve[0].Tag())
	}

	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.ResponseError{
		Message: err.Error(),
		Error:   err,
	})
}

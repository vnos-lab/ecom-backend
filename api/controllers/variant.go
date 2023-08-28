package controller

import (
	"ecom/api/middlewares"
	"ecom/api_errors"
	"ecom/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type ProductVariantController struct {
	BaseController
	variantService domain.ProductVariantService
	logger         *zap.Logger
}

func NewProductVariantController(c *gin.RouterGroup, logger *zap.Logger, variantService domain.ProductVariantService, middleware *middlewares.GinMiddleware) *ProductVariantController {
	controller := &ProductVariantController{variantService: variantService, logger: logger}
	g := c.Group("/variant")
	g.Use(middleware.TimeOut(5*time.Second, api_errors.ErrRequestTimeout))

	g.POST("/", controller.Create)
	g.PUT("/:id", controller.Update)
	g.DELETE("/:id", controller.Delete)

	return controller
}

func (v *ProductVariantController) Create(c *gin.Context) {
	var req domain.CreateProductVariantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		v.ResponseValidationError(c, err)
		return
	}

	result, err := v.variantService.Create(c.Request.Context(), req)
	if err != nil {
		v.ResponseError(c, err)
		return
	}

	v.Response(c, http.StatusOK, "Variant created successfully", result)
}

func (v *ProductVariantController) Update(c *gin.Context) {
	var req domain.UpdateProductVariantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		v.ResponseValidationError(c, err)
		return
	}

	uuid, _ := uuid.FromString(c.Param("id"))

	result, err := v.variantService.Update(c.Request.Context(), req, uuid)
	if err != nil {
		v.ResponseError(c, err)
		return
	}

	v.Response(c, http.StatusOK, "Variant updated successfully", result)
}

func (v *ProductVariantController) Delete(c *gin.Context) {
	err := v.variantService.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		v.ResponseError(c, err)
		return
	}

	v.Response(c, http.StatusOK, "Variant deleted successfully", nil)
}

package controller

import (
	"ecom/api/middlewares"
	"ecom/api_errors"
	"ecom/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductController struct {
	BaseController
	productService domain.ProductService
	logger         *zap.Logger
}

func NewProductController(c *gin.RouterGroup, productService domain.ProductService, logger *zap.Logger, middleware *middlewares.GinMiddleware) *ProductController {
	controller := &ProductController{
		productService: productService,
		logger:         logger,
	}

	g := c.Group("/product")
	g.Use(middleware.TimeOut(5*time.Second, api_errors.ErrRequestTimeout))

	g.POST("/create", controller.Create)
	g.GET("/:id", controller.GetByID)

	return controller
}

func (p *ProductController) Create(c *gin.Context) {
	var req domain.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		p.ResponseValidationError(c, err)
		return
	}

	result, err := p.productService.Create(c.Request.Context(), req)
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Product created successfully", result)
}

func (p *ProductController) GetByID(c *gin.Context) {
	result, err := p.productService.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		p.ResponseError(c, err)
		return
	}

	p.Response(c, http.StatusOK, "Product retrieved successfully", result)
}

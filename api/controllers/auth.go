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

type AuthController struct {
	BaseController
	authService domain.AuthService
	logger      *zap.Logger
}

func NewAuthController(c *gin.RouterGroup, authService domain.AuthService, logger *zap.Logger, middleWare *middlewares.GinMiddleware) *AuthController {
	controller := &AuthController{
		authService: authService,
		logger:      logger,
	}
	g := c.Group("/auth")
	g.Use(middleWare.TimeOut(5*time.Second, api_errors.ErrRequestTimeout))
	g.POST("/register", controller.Register)
	g.POST("/login", controller.Login)

	return controller
}

func (b *AuthController) Register(c *gin.Context) {
	var req domain.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}

	_, err := b.authService.Register(c.Request.Context(), req)
	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", nil)
}

func (b *AuthController) Login(c *gin.Context) {
	var req domain.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseValidationError(c, err)
		return
	}

	res, err := b.authService.Login(c.Request.Context(), req)

	if err != nil {
		b.ResponseError(c, err)
		return
	}
	b.Response(c, http.StatusOK, "success", res)
}

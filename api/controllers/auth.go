package controller

import (
	"erp/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthController struct {
	AuthService domain.AuthService
	logger      *zap.Logger
	BaseController
}

func NewAuthController(c *gin.RouterGroup, authService domain.AuthService, logger *zap.Logger) *AuthController {
	controller := &AuthController{
		AuthService: authService,
		logger:      logger,
	}
	g := c.Group("/auth")
	g.POST("/register", controller.Register)
	g.POST("/login", controller.Login)

	return controller
}

func (b *AuthController) Register(c *gin.Context) {
	var req domain.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}

	_, err := b.AuthService.Register(c.Request.Context(), req)

	if err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	b.Response(c, http.StatusOK, "success", nil)
}

func (b *AuthController) Login(c *gin.Context) {
	var req domain.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}

	res, err := b.AuthService.Login(c.Request.Context(), req)

	if err != nil {
		b.ResponseError(c, http.StatusBadRequest, []error{err})
		return
	}
	b.Response(c, http.StatusOK, "success", res)
}

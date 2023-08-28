package middlewares

import (
	"context"
	"ecom/api/response"
	"ecom/api_errors"
	config "ecom/config"
	"ecom/infrastructure/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (e *GinMiddleware) JWT(config *config.Config, db *db.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: add public routes in config
		publicRoutes := []string{
			"/api/auth/login",
			"/api/auth/register",
		}

		for _, route := range publicRoutes {
			if route == c.Request.URL.Path {
				c.Next()
				return
			}
		}

		auth := c.Request.Header.Get("Authorization")

		resError := func() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseError{
				Message: "Unauthorized",
				Error:   api_errors.ErrUnauthorizedAccess,
			})
		}

		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseError{
				Message: "Unauthorized",
				Error:   api_errors.ErrUnauthorizedAccess,
			})
			return
		}
		jwtToken := strings.Split(auth, " ")[1]

		if jwtToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseError{
				Message: "JWT token is missing",
				Error:   api_errors.ErrUnauthorizedAccess,
			})
			return
		}

		token, err := parseToken(jwtToken, config.Jwt.Secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseError{
				Message: err.Error(),
				Error:   api_errors.ErrUnauthorizedAccess,
			})
			return
		}

		claims, OK := token.Claims.(jwt.MapClaims)
		if !OK {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseError{
				Message: err.Error(),
				Error:   api_errors.ErrUnauthorizedAccess,
			})
			return
		}

		userID, OK := claims["sub"].(string)
		if !OK {
			resError()
			return
		}
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "x-user-id", userID))

		c.Next()
	}
}

func parseToken(jwtToken string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, api_errors.ErrTokenBadSignedMethod
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, api_errors.ErrUnauthorizedAccess
	}

	return token, nil
}

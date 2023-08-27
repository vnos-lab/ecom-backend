package api_errors

import (
	"errors"
	"net/http"
)

var (
	ErrInternalServerError  = errors.New("internal server error")
	ErrUnauthorizedAccess   = errors.New("unauthorized access")
	ErrTokenBadSignedMethod = errors.New("bad signed method received")
	ErrTokenExpired         = errors.New("token expired")
	ErrTokenInvalid         = errors.New("token invalid")
	ErrTokenMalformed       = errors.New("token malformed")
	ErrUserNotFound         = errors.New("user not found")
)

func GetStatusCode(err error) (int, bool) {
	if v, ok := MapErrorStatusCode[err.Error()]; !ok {
		return http.StatusInternalServerError, false
	} else {
		return v, true
	}
}

var MapErrorStatusCode = map[string]int{
	ErrInternalServerError.Error():  http.StatusInternalServerError,
	ErrUnauthorizedAccess.Error():   http.StatusUnauthorized,
	ErrTokenBadSignedMethod.Error(): http.StatusUnauthorized,
	ErrTokenExpired.Error():         http.StatusUnauthorized,
	ErrTokenInvalid.Error():         http.StatusUnauthorized,
	ErrTokenMalformed.Error():       http.StatusUnauthorized,
	ErrUserNotFound.Error():         http.StatusNotFound,
}

package service

import (
	"context"
	config "ecom/config"
	"ecom/domain"
	models "ecom/models"

	"golang.org/x/crypto/bcrypt"
)

type (
	AuthServiceImpl struct {
		userService domain.UserService
		jwtService  domain.JwtService
		config      *config.Config
	}
)

func NewAuthService(userService domain.UserService, config *config.Config, jwtService domain.JwtService) domain.AuthService {
	return &AuthServiceImpl{
		userService: userService,
		jwtService:  jwtService,
		config:      config,
	}
}

func (a *AuthServiceImpl) Register(ctx context.Context, req domain.RegisterRequest) (user *models.User, err error) {

	// if req.RequestFrom != string(constants.Web) {
	// 	roleKey = constants.RoleStoreOwner
	// }

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return user, err
	}

	req.Password = string(encryptedPassword)

	user, err = a.userService.Create(ctx, models.User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})

	return user, err
}

func (a *AuthServiceImpl) Login(ctx context.Context, req domain.LoginRequest) (res *domain.LoginResponse, err error) {
	user, err := a.userService.GetByEmail(ctx, req.Email)

	// if req.RequestFrom != string(constants.Web) {
	// 	if user.RoleKey != constants.RoleStoreOwner {
	// 		return nil, api_errors.ErrUnauthorizedAccess
	// 	}
	// }

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := a.jwtService.GenerateAuthTokens(user.ID.String())
	if err != nil {
		return nil, err
	}

	res = &domain.LoginResponse{
		User: domain.UserResponse{
			ID:        user.ID.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		},
		Token: domain.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiresIn:    a.config.Jwt.AccessTokenExpiresIn,
		},
	}

	return res, nil
}

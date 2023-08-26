package service

import (
	"context"
	config "ecom/config"
	"ecom/domain"
	"ecom/infrastructure/cache"
	models "ecom/models"
	"ecom/utils"
)

type (
	UserServiceImpl struct {
		userRepo    domain.UserRepository
		cacheClient *cache.Client
		config      *config.Config
	}
)

func NewUserService(itemRepo domain.UserRepository, config *config.Config, cacheClient *cache.Client) domain.UserService {
	return &UserServiceImpl{
		userRepo:    itemRepo,
		cacheClient: cacheClient,
		config:      config,
	}
}

func (u *UserServiceImpl) Create(ctx context.Context, user models.User) (*models.User, error) {
	r, err := u.userRepo.Create(ctx, user)
	return r, err
}

func (u *UserServiceImpl) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		if utils.ErrNoRows(err) {
			return user, err
		}
		return nil, err
	}
	return user, err
}

func (u *UserServiceImpl) GetByID(ctx context.Context, id string) (user *models.User, err error) {
	user, err = u.userRepo.GetByID(ctx, id)
	if err != nil {
		if utils.ErrNoRows(err) {
			return user, err
		}
		return nil, err
	}
	return
}

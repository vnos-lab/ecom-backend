package domain

import (
	"context"
	"erp/models"
)

type UserRepository interface {
	GetByID(ctx context.Context, id string) (res *models.User, err error)
	GetByEmail(ctx context.Context, email string) (res *models.User, err error)
	Create(ctx context.Context, user models.User) (res *models.User, err error)
}

type UserService interface {
	Create(ctx context.Context, user models.User) (*models.User, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

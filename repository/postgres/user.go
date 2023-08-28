package postgres

import (
	"context"
	"ecom/domain"
	"ecom/infrastructure/db"
	"ecom/models"
	"ecom/utils"
	"fmt"

	"github.com/pkg/errors"
)

type userRepository struct {
	*db.Database
}

func NewUserRepository(db *db.Database) domain.UserRepository {
	utils.MustHaveDb(db)
	return &userRepository{db}
}

func (u *userRepository) Create(ctx context.Context, user models.User) (res *models.User, err error) {
	args := []interface{}{user.Email, user.Password, user.FirstName, user.LastName}
	query, _, _ := utils.Psql().Insert("users").
		Columns("email", "password", "first_name", "last_name").
		Values(args...).
		ToSql()

	fmt.Println(query)

	_, err = u.ExecContext(ctx, query, args...)

	return &user, errors.WithStack(err)
}

func (u *userRepository) GetByID(ctx context.Context, id string) (res *models.User, err error) {
	query, _, _ := utils.Psql().Select("*").From("users").Where("id = ?", id).Limit(1).ToSql()

	err = u.SelectContext(ctx, &res, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query, _, _ := utils.Psql().Select("*").From("users").Where("email = ?", email).Limit(1).ToSql()
	uu := models.User{}

	err := u.Get(&uu, query, email)

	return &uu, errors.WithStack(err)
}

package models

type User struct {
	BaseModel
	FirstName string `json:"first_name" db:"first_name" validate:"min=1,max=50"`
	LastName  string `json:"last_name" db:"last_name" validate:"min=1,max=50"`
	Email     string `json:"email" db:"email" validate:"email"`
	Password  string `json:"password" db:"password" validate:"min=6,max=20"`
}

package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type BaseModel struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	UpdaterID *uuid.UUID `json:"updater_id" db:"updater_id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}

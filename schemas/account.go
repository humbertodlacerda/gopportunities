package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	Name    string `json:"name"`
	Balance int64  `json:"balace"`
}

type AccountResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Balance   int64     `json:"balace"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

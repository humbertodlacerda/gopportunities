package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

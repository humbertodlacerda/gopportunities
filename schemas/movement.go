package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Movement struct {
	gorm.Model
	Date        string `json:"date"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Account     string `json:"account"`
	Value       int64  `json:"value"`
	Status      int64  `json:"status"`
}

type MovementResponse struct {
	ID          uint      `json:"id"`
	Date        string    `json:"date"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Account     string    `json:"account"`
	Value       int64     `json:"value"`
	Status      int64     `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}

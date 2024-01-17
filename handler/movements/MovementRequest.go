package movements

import (
	"fmt"
	"github.com/humbertodlacerda/gopportunities/handler"
)

type CreateMovementRequest struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Account     string `json:"account"`
	Value       int64  `json:"value"`
	Status      int64  `json:"status"`
}

func (c *CreateMovementRequest) Validate() error {
	if c.Date == "" && c.Description == "" && c.Category == "" && c.Account == "" && c.Value == 0 && c.Status == 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if c.Date == "" {
		return handler.ErrorParamIsRequired("date", "string")
	}
	if c.Description == "" {
		return handler.ErrorParamIsRequired("description", "string")
	}
	if c.Category == "" {
		return handler.ErrorParamIsRequired("category", "string")
	}
	if c.Account == "" {
		return handler.ErrorParamIsRequired("account", "boolean")
	}
	if c.Value <= 0 {
		return handler.ErrorParamIsRequired("value", "int64")
	}
	if c.Status <= 0 {
		return handler.ErrorParamIsRequired("status", "int64")
	}

	return nil
}

type UpdateMovementRequest struct {
	Date        string `json:"date"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Account     string `json:"account"`
	Value       int64  `json:"value"`
	Status      int64  `json:"status"`
}

func (u *UpdateMovementRequest) Validate() error {
	if u.Date != "" || u.Description != "" || u.Category != "" || u.Account != "" || u.Value >= 0 || u.Status >= 0 {
		return nil
	}

	return fmt.Errorf("request body is empty or malformed")
}

package opening

import (
	"fmt"
	"github.com/humbertodlacerda/gopportunities/handler"
)

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return handler.ErrorParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return handler.ErrorParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return handler.ErrorParamIsRequired("location", "string")
	}
	if r.Remote == nil {
		return handler.ErrorParamIsRequired("remote", "boolean")
	}
	if r.Link == "" {
		return handler.ErrorParamIsRequired("link", "string")
	}
	if r.Salary <= 0 {
		return handler.ErrorParamIsRequired("salary", "int64")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Link != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least on valid field must be provided")
}

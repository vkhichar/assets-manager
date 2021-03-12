package contract

import (
	"errors"
	"strings"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func (req RegisterRequest) Validate() error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}

	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}

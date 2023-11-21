package dtos

import "fmt"

func missingParamError(param, typ string) error {
	return fmt.Errorf("params: %s (type: %s) is required", param, typ)
}

type RequestRegisterPlayerDTO struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	ConfirmPassword string `json:"confirm_password"`
}

type ResponseRegisterPlayerDTO struct {
	PlayerID int `json:"id"`
}

func (v *RequestRegisterPlayerDTO) Validate() error {
	if v.Username == "" && v.Password == "" && v.ConfirmPassword == "" && v.Email == "" {
		return fmt.Errorf("body cannot empty")
	}

	if v.Email == "" {
		return missingParamError("email", "string")
	}

	if v.Username == "" {
		return missingParamError("username", "string")
	}

	if v.Password == "" {
		return missingParamError("password", "string")
	}

	if v.Password != v.ConfirmPassword {
		return fmt.Errorf("passwords not equals")
	}

	return nil
}

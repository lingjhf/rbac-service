package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type (
	SignupWithPhone struct {
		Username string `json:"username"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
	SignupWithEmail struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	SigninWithUsername struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	SigninWithPhone struct {
		Phone string `json:"phone"`
		Code  uint   `json:"code"`
	}
	SigninWithEmail struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	Token struct {
		Token string `json:"token"`
	}
	ResetPassword struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
)

func (s *SignupWithPhone) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Username, validation.Required, validation.Length(1, 32)),
		validation.Field(&s.Phone, validation.Required, validation.Length(1, 20)),
		validation.Field(&s.Password, validation.Required, validation.Length(6, 32)),
	)
}

func (s *SignupWithEmail) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Username, validation.Required, validation.Length(1, 32)),
		validation.Field(&s.Email, validation.Required, is.EmailFormat),
		validation.Field(&s.Password, validation.Required, validation.Length(6, 32)),
	)
}

func (s *SigninWithUsername) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Username, validation.Required, validation.Length(1, 32)),
		validation.Field(&s.Password, validation.Required, validation.Length(6, 32)),
	)
}

func (s *SigninWithPhone) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Phone, validation.Required, validation.Length(1, 20)),
	)
}

func (s *SigninWithEmail) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Email, validation.Required, is.EmailFormat),
		validation.Field(&s.Password, validation.Required, validation.Length(6, 32)),
	)
}

func (s *ResetPassword) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.OldPassword, validation.Required, validation.Length(6, 32)),
		validation.Field(&s.NewPassword, validation.Required, validation.Length(6, 32)),
	)
}

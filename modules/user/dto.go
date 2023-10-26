package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	is "github.com/go-ozzo/ozzo-validation/is"
)

type RegisterRequestDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (r RegisterRequestDto) Validate() error {
	return validation.ValidateStruct(&r,
		// Name cannot be empty, and the length must between 5 and 20
		validation.Field(&r.Name, validation.Required, validation.Length(5, 20)),
		// Email cannot be empty, Should have proper email format and the length must between 5 and 20
		validation.Field(&r.Email, validation.Required, is.Email),
	)
}

type RegisterResponseDto struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type FindUserDto struct {
	Name string `json:"name" form:"name"`
}

func (r FindUserDto) Validate() error {
	return validation.ValidateStruct(&r,
		// Name cannot be empty, and the length must between 5 and 20
		validation.Field(&r.Name, validation.Required, validation.Length(5, 20)),
	)
}

type FindUserResponseDto struct {
	Users []User `json:"users"`
	Total int64  `json:"total"`
}

type LoginUserDto struct {
	Name string `json:"name" form:"name"`
}

func (r LoginUserDto) Validate() error {
	return validation.ValidateStruct(&r,
		// Name cannot be empty, and the length must between 5 and 20
		validation.Field(&r.Name, validation.Required, validation.Length(5, 20)),
	)
}

type LoginUserResponseDto struct {
	Users []User `json:"users"`
	Total int64  `json:"total"`
}

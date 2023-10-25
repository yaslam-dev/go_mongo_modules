package user

type RegisterRequestDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RegisterResponseDto struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type FindUserDto struct {
	Name string `json:"name" form:"name"`
}

type FindUserResponseDto struct {
	Users []User `json:"users"`
	Total int64  `json:"total"`
}

type LoginUserDto struct {
	Name string `json:"name" form:"name"`
}

type LoginUserResponseDto struct {
	Users []User `json:"users"`
	Total int64  `json:"total"`
}

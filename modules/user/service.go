// Package user provide services, repositories and handlers for everything related to user
package user

import (
	"time"

	"github.com/Yasir900Aslam/go_mongo_modules/core"
	"github.com/golang-jwt/jwt"
)

func Register(registerDto *RegisterRequestDto) (*RegisterResponseDto, error) {
	var err error
	var token string
	var fetchedUser *User

	//according to type get user info
	//check if user exist
	user := &User{
		Name:     registerDto.Name,
		Email:    registerDto.Email,
		Bio:      "",
		Image:    "http://dkkdopkeq",
		LoggedBy: "web",
	}

	fetchedUser, err = user.GetByEmail()

	if err != nil {
		return nil, err
	}

	if fetchedUser != nil {
		token, err := generateToken(user)
		if err != nil {
			return nil, err
		}
		return &RegisterResponseDto{
			Token: token,
			User:  *user,
		}, nil
	}

	//if not create user
	err = user.Create()
	if err != nil {
		return nil, err
	}
	token, err = generateToken(user)
	if err != nil {
		return nil, err
	}
	return &RegisterResponseDto{
		Token: token,
		User:  *user,
	}, nil
}

func Search(findUserDto *FindUserDto) (*FindUserResponseDto, error) {
	user := &User{}
	users, err := user.SearchByName()
	if err != nil {
		return &FindUserResponseDto{
			Users: []User{},
			Total: 0,
		}, nil
	}

	return &FindUserResponseDto{
		Users: users,
		Total: int64(len(users)),
	}, nil
}

func generateToken(user *User) (string, error) {
	secret := core.ConfigInstance().JwtSecret
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["image"] = user.Image
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	return token.SignedString([]byte(secret))
}

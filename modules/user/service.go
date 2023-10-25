package user

import (
	"time"

	"github.com/Yasir900Aslam/go_mongo_modules/core"
	"github.com/golang-jwt/jwt"
)

func Register(registerDto *RegisterRequestDto) (*RegisterResponseDto, error) {
	//according to type get user info
	//check if user exist
	user := &User{
		Name:     registerDto.Email,
		Email:    registerDto.Email,
		Bio:      "",
		Image:    "http://dkkdopkeq",
		LoggedBy: "web",
	}

	err := user.GetByEmail()
	// If the error is nill user is found
	// return token
	if err == nil {
		token, err := generateToken(user)
		if err != nil {

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
	token, err := generateToken(user)
	if err != nil {
		return nil, err
	}
	//if exist return it
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
	secret := core.ConfigInstance().JWT_SECRET
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["image"] = user.Image
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	return token.SignedString([]byte(secret))
}

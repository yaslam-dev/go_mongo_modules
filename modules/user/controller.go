package user

import (
	"fmt"

	"github.com/Yasir900Aslam/go_mongo_modules/core"
	"github.com/Yasir900Aslam/go_mongo_modules/home"
	"github.com/gofiber/fiber/v2"
)

type AuthController = core.Controller

func NewAuthController() *AuthController {

	return core.NewController().
		SetName("Auth").
		SetVersion("v1").
		SetPath("/auth").
		AddHandler(core.NewHandler().
			SetMethod(core.HTTP_POST_METHOD).
			SetPath("/register").
			SetHandlerFunc(RegisterUser).
			SetDescription("Register user").
			SetRequestDto(&RegisterRequestDto{}).
			SetResponseDto(&RegisterResponseDto{})).
		AddHandler(core.NewHandler().
			SetMethod(core.HTTP_POST_METHOD).
			SetPath("/login").
			SetHandlerFunc(LoginUser).
			SetDescription("Login user").
			SetRequestDto(&LoginUserDto{}).
			SetResponseDto(&LoginUserResponseDto{})).
		AddHandler(core.NewHandler().
			SetMethod(core.HTTP_GET_METHOD).
			SetPath("/login").
			SetHandlerFunc(LoginUser).
			SetDescription("get token of user").
			SetResponseDto(&LoginUserResponseDto{})).
		AddHandler(core.NewHandler().
			SetMethod(core.HTTP_GET_METHOD).
			SetPath("/:id").
			SetHandlerFunc(GetUser).
			SetDescription("get user").
			SetResponseDto(&User{}))
}

// Login user
func RegisterUser(ctx *fiber.Ctx) error {
	registerDto := &RegisterRequestDto{}
	if err := ctx.BodyParser(registerDto); err != nil {
		if err := ctx.Status(400).JSON(home.ErrorResponse([]error{err}, home.CANNOT_PARSE_BODY)); err != nil {
			return err
		}
		return err
	}
	user, err := Register(registerDto)
	if err != nil {
		err := ctx.Status(401).JSON(home.ErrorResponse([]error{err}, home.OAUTH_TOKEN_NOT_CORRECT))
		if err != nil {
			return err
		}
		return err
	}
	if err := ctx.Status(200).JSON(user); err != nil {
		return err
	}
	return nil
}

func LoginUser(ctx *fiber.Ctx) error {
	findUserDto := &LoginUserDto{}
	err := ctx.BodyParser(findUserDto)
	if err != nil {
		if err := ctx.Status(400).JSON(home.ErrorResponse([]error{err}, home.CANNOT_PARSE_BODY)); err != nil {
			return err
		}
		return err
	}
	fmt.Println(findUserDto)
	return nil
}

type QueryStruct struct {
	Q      string   `json:"q"`
	Filter []string `json:"filter"`
}
type ParamStruct struct {
	ID string `json:"id"`
}

func GetUser(ctx *fiber.Ctx) error {
	var err error
	params := ParamStruct{}
	if err = ctx.ParamsParser(&params); err != nil {
		return err
	}
	query := QueryStruct{}
	if err = ctx.QueryParser(&query); err != nil {
		return err
	}
	if err = ctx.Status(200).JSON(map[string]interface{}{
		"path":  params,
		"query": query,
	}); err != nil {
		return err
	}
	return nil
}

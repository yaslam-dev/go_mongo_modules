package home

import (
	"github.com/Yasir900Aslam/go_mongo_modules/core"
	"github.com/gin-gonic/gin"
)

const (
	CannotParseBody     = 3001
	OAuthTokenIncorrect = 4001
	UserServerError     = 5001
)

func ErrorResponse(errs []error, code int) *gin.H {
	return &gin.H{
		"code":    code,
		"errors":  core.ErrorsToJSON(errs),
		"success": false,
	}
}

package users

import (
	"foods-spider-server/internal/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUserApi interface {
	api.IBase
	GetUserList() func(ctx *gin.Context)
}

type UserApi struct {
}

func (f *UserApi) SetPrefix() string {
	return "/api/v1/users"
}
func (f *UserApi) GetUserList() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "获取用户列表",
		})
	}
}
func NewUsersApi() IUserApi {
	return &UserApi{}
}

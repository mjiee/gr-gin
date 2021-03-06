package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mjiee/grf-gin/app/lib"
	"github.com/mjiee/grf-gin/app/pkg/apperr"
	"github.com/mjiee/grf-gin/app/pkg/response"
)

// UserHandler 用户处理器
type UserHandler struct {
	userSrv *lib.UserService
}

// NewUserHandler 创建用户处理器
func NewUserHandler(userSrv *lib.UserService) *UserHandler {
	return &UserHandler{userSrv}
}

// @Summary 'GetUserInfo'
// @description '获取用户信息'
// @Tags user
// @Security ApiKeyAuth
// @Prodece application/json
// @response default {object} response.Response "响应包装"
// @Success 200 {object} model.User "用户信息"
// @Router /user/getUserInfo [get]
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	user, err := h.userSrv.GetUserInfo(c.MustGet("id").(string), false)
	if err != nil {
		response.Failure(c, apperr.BusinessErr, err.Error())
		return
	}
	response.Success(c, user)
}

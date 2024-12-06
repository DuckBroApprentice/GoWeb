package controller

import (
	"github.com/DuckBroApprentice/GoWeb/dao"
	"github.com/DuckBroApprentice/GoWeb/models"
	"github.com/DuckBroApprentice/GoWeb/tool"
	"github.com/gin-gonic/gin"
)

var UserController struct{}

func (uc *UserController) UserLogin (ctx *gin.Context){
	var userLoginParam param.UserLoginParam

	err := tool.Decode(ctx.Request.Body,&userLoginParam)
	if err != nil {
		tool.Failed(ctx,"參數解析失敗")
		return
	}
	us := service.UserService{}
	user := us.Login(userLoginParam.Name,userLoginParam.Password)
	if user.Id != 0 {
		tool.Success(ctx , &user)
		return
	}
	tool.Failed(ctx,"登錄失敗")
}
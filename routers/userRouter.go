package routers

import (
	"github.com/DuckBroApprentice/GoWeb/controller"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) UserRouter(engine *gin.Engine) {
	engine.POST("/user/login", controller.UserLogin)
}

package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS int = iota
	FAILED
)

func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "成功",
		"data": v,
	})
}

func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"code": FAILED,
		"msg":  "失敗",
	})
}

package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yeeep.cn/yepgo/handler"
)

func User(apiRouter *gin.RouterGroup) {
	rg := apiRouter.Group("/user")
	{
		rg.POST("/register", handler.UserRegister)
		rg.GET("/hi", handler.BDconf)
	}
}

func hi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hi",
	})
}

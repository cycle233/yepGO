package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"yeeep.cn/yepgo/conf"
	"yeeep.cn/yepgo/model"
)

func UserRegister(c *gin.Context) {
	var user model.UserRegister
	passwordAgain := c.PostForm("password-again")
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "数据不合法"})
	} else if passwordAgain != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "两次密码不一致"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"username":      user.Username,
			"password":      user.Password,
			"passwordAgain": passwordAgain,
			"email":         user.Email,
			"msg":           "注册成功",
		})
	}
	logrus.Debug(user, passwordAgain)
}

func BDconf(c *gin.Context) {
	conf := conf.LoadDBConf()
	c.JSON(http.StatusOK, gin.H{
		"db_type":     conf.DBType,
		"db_host":      conf.Host,
		"db_port":     conf.Port,
		"db_username": conf.Username,
		"db_password": conf.Password,
		"db_name":     conf.Name,
	})
}

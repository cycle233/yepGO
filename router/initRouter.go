package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//r.Static("/statics", "./statics")


	apiRouter := r.Group("/api")
	{
		User(apiRouter)
	}

	return r
}

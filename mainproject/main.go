package main

import (
	"mainproject/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化路由
	router := gin.Default()
	//映射静态资源
	router.Static("/home", "./view")
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	r1 := router.Group("api/v1.0")
	{
		r1.GET("/areas", controller.GetArea)
		r1.GET("/test", controller.Test)
		r1.GET("/session", controller.GetSession)
		r1.GET("/imagecode/:uuid", controller.GetImageCode)
		r1.GET("/smscode/:mobile", controller.GetSmscd)
		r1.GET("/user", controller.GetUserInfo)
		r1.POST("/users", controller.PostRet)
		r1.POST("/sessions", controller.PostLogin)
		r1.POST("/user/avatar", controller.PostAvatar)
		r1.DELETE("/session", controller.DeleteSession)
		r1.PUT("/user/name",controller.PutUserName)
	}

	//开启监听
	router.Run(":8080")
}

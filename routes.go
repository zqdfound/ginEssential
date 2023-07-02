package main

import (
	"ginEssential/controller"
	"ginEssential/middleWare"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	//使用中间件保护用户信息接口，相当于拦截器
	r.GET("/api/auth/info", middleWare.AuthMiddleware(), controller.Info)
	return r
}

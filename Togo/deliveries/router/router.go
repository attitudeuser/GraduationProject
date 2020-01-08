package router

import (
	"GraduationProject/deliveries/handler"
	"GraduationProject/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	//开启允许跨域
	router.Use(middleware.CORS())
	//通用路由
	var general = router.Group("/api/general")
	{
		//用户注册功能
		general.POST("/signup", handler.Signup)
		//用户登录功能
		general.POST("/signin", handler.Signin)
		//用户忘记密码 找回 发送邮件接口
		general.GET("/forget", handler.Forget)
		//用户忘记密码 找回 重置密码接口
		general.POST("/reset/:key", handler.Reset)
	}
	//用户路由 这部分路由需要JWT认证 使用JWT中间件
	var user = router.Group("/api/user", middleware.JWTAuth())
	{
		//查找用户的个人信息
		user.GET("/profile", handler.Profile)
		//用户信息更新功能
		user.POST("/update", handler.ModifyInformation)
		//用户读取消息
		user.GET("/message", handler.MessageRead)
		//用户发送消息
		user.POST("/message/send", handler.MessageSend)
	}
	//管理员路由
	var admin = router.Group("/api/admin")
	{
		//查找用户列表
		admin.GET("/userlist", handler.FindMany)
	}
	return router
}

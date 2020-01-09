package router

import (
	"GraduationProject/deliveries/handler"
	"GraduationProject/middleware"
	"github.com/gin-gonic/gin"
)

// Router 路由集中配置 返回路由
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
		general.GET("/reset", handler.Reset)
		general.GET("/profile", handler.Profile)
	}
	//用户路由 这部分路由需要JWT认证 使用JWT中间件
	var user = router.Group("/api/user", middleware.JWTAuth())
	{
		//查找用户的个人信息
		user.GET("/profile", handler.Profile)
		//用户信息更新功能
		user.POST("/update", handler.ModifyInformation)
	}
	//消息路由
	var message = router.Group("/api/message")
	{
		//用户发送消息
		message.POST("/send", handler.MessageSend)
		message.GET("/list", handler.MessageList)
	}
	//课题路由
	var issue = router.Group("/api/issue")
	{
		issue.POST("/publish", handler.IssuePublish)
		issue.GET("/list", handler.IssueList)
	}
	var subject = router.Group("/api/subject")
	{
		subject.POST("/select", handler.SelectSubject)
		subject.GET("/get", handler.GetSubject)
		subject.GET("/delete", handler.DeleteSubject)
		subject.GET("/relate", handler.RelateToIssue)
	}
	return router
}

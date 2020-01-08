package handler

import (
	"GraduationProject/usecases"
	"github.com/gin-gonic/gin"
)

//uuc user use case
var uuc = usecases.NewUserUC()

//用户-注册
func Signup(this *gin.Context) {
	this.JSON(uuc.Signup(this))
}

//用户-登陆
func Signin(this *gin.Context)  {
	this.JSON(uuc.Signin(this))
}

//用户-退出
func Logout(this *gin.Context)  {
	this.JSON(uuc.Logout(this))
}

// Forget 忘记密码 发送邮件功能
func Forget(this *gin.Context) {
	this.JSON(uuc.Forget2SendEmail(this))
}

// Reset 忘记密码 重置密码功能
func Reset(this *gin.Context) {
	this.JSON(uuc.Forget2ResetPassword(this))
}

//用户-个人中心
func Profile(this *gin.Context)  {
	this.JSON(uuc.Profile(this))
}


//查看用户信息
func FindOne(this *gin.Context) {
	this.JSON(uuc.FindOne(this))
}

//查看用户列表
func FindMany(this *gin.Context) {
	r := uuc.FindMany(this)
	this.JSON(r.Code, r)
}

//修改用户
func ModifyInformation(this *gin.Context) {
	this.JSON(uuc.ModifyInformation(this))
}
package usecases

import (
	"github.com/gin-gonic/gin"
)

// UserInterface 是用户模块的接口
type UserInterface interface {
	//用户注册接口
	Signup(ctx *gin.Context) (int, *Response)
	//用户登陆接口
	Signin(ctx *gin.Context) (int, *Response)
	//找回密码 发送邮件功能
	Forget2SendEmail(ctx *gin.Context) (int, *Response)
	//找回密码 重置密码功能
	Forget2ResetPassword(ctx *gin.Context) (int, *Response)
	//用户个人中心
	Profile(ctx *gin.Context) (int, *Response)
	//修改用户信息接口
	ModifyInformation(ctx *gin.Context) (int, *Response)
	//查询单一用户接口
	FindOne(ctx *gin.Context) (int, *Response)
}

// MessageInterface 消息交流接口
type MessageInterface interface {
	//消息发送接口
	Send(*gin.Context) (int, *Response)
	//读消息接口
	Read(*gin.Context) (int, *Response)
	//消息列表
	List(*gin.Context) (int, *Response)
}

// IssueInterface 发布选题接口
type IssueInterface interface {
	//发布选题
	Publish(ctx *gin.Context) (int, *Response)
	//修改选题
	Update(ctx *gin.Context) (int, *Response)
	//删除选题
	Delete(ctx *gin.Context) (int, *Response)
	//查询某个老师的选题
	FindById(ctx *gin.Context) (int, *Response)
	//查询所有选题
	List(ctx *gin.Context) (int, *Response)
}

// 选题接口
type SubjectInterface interface {
	//查看个人选题情况
	Get(ctx *gin.Context) (int, *Response)
	//用户进行选题
	Select(ctx *gin.Context) (int, *Response)
	//删除已选选题
	Delete(ctx *gin.Context) (int, *Response)
	//用户关联的选题信息
	RelateIssue(ctx *gin.Context) (int, *Response)
}

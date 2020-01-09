package usecases

import (
	"GraduationProject/models"
	"GraduationProject/repositories"
	"GraduationProject/utils"
	"github.com/gin-gonic/gin"
)

// issue_module 教师出题管理

var ir = repositories.NewIssueRepo()

type IssueUC struct {
}

func NewIssueUC() IssueInterface {
	return &IssueUC{}
}

func (this IssueUC) Publish(ctx *gin.Context) (int, *Response) {
	var issue = new(models.Issue)
	//检测注册信息是否解析完成
	if err := ctx.Bind(issue); err != nil {
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "解析参数错误",
		}
	}
	if _, err := ir.Publish(issue); err != nil {
		return StatusServerError, &Response{
			Code:    ErrorDatabaseInsert,
			Message: "数据库插入失败",
		}
	}
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
	}
}

func (this IssueUC) Update(ctx *gin.Context) (int, *Response) {
	var issue = new(models.Issue)
	//检测注册信息是否解析完成
	if err := ctx.Bind(issue); err != nil {
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "解析参数错误",
		}
	}
	if _, err := ir.Update(issue); err != nil {
		return StatusServerError, &Response{
			Code:    ErrorDatabaseUpdate,
			Message: "数据库更新失败",
		}
	}
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
	}
}

func (this IssueUC) Delete(ctx *gin.Context) (int, *Response) {
	var issue = new(models.Issue)
	issue.Id = utils.ParseStringToInt64(ctx.Query("id"))
	if _, err := ir.Delete(issue.Id); err != nil {
		return StatusServerError, &Response{
			Code:    ErrorDatabaseDelete,
			Message: "数据库删除失败",
		}
	}
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
	}
}

func (this IssueUC) FindById(ctx *gin.Context) (int, *Response) {
	var user = new(models.User)
	user.Id = utils.ParseStringToInt64(ctx.Query("uid"))
	list := ir.FindById(user.Id)
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "OK",
		Data:    list,
	}
}

func (this IssueUC) List(ctx *gin.Context) (int, *Response) {
	list := ir.List()
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "OK",
		Data:    list,
	}
}

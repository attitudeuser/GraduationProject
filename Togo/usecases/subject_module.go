package usecases

import (
	"GraduationProject/models"
	"GraduationProject/repositories"
	"github.com/gin-gonic/gin"
)

var sr = repositories.NewSubjectRepo()

type SubjectUC struct {
}

func NewSubjectUC() SubjectInterface {
	return &SubjectUC{}
}

func (this SubjectUC) Select(ctx *gin.Context) (int, *Response) {
	var s = new(models.Subject)
	//检测登陆信息是否绑定成功
	if err := ctx.Bind(s); err != nil {
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "解析参数错误",
		}
	}
	//选中
	if _, err := sr.Select(s); err != nil {
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

func (this SubjectUC) Delete(ctx *gin.Context) (int, *Response) {
	var s = new(models.Subject)
	//检测登陆信息是否绑定成功
	if err := ctx.Bind(s); err != nil {
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "解析参数错误",
		}
	}
	//选中
	if _, err := sr.Delete(s); err != nil {
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

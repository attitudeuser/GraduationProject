package usecases

import (
	"GraduationProject/models"
	"GraduationProject/repositories"
	"GraduationProject/utils"
	"github.com/gin-gonic/gin"
)

var sr = repositories.NewSubjectRepo()

type SubjectUC struct {
}

func NewSubjectUC() SubjectInterface {
	return &SubjectUC{}
}

func (this *SubjectUC) Get(ctx *gin.Context) (int, *Response) {
	var s = new(models.Subject)
	//检测信息是否绑定成功
	uid := utils.ParseStringToInt64(ctx.Query("uid"))
	//选中
	s = sr.Get(uid)
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
		Data:    s,
	}
}

func (this *SubjectUC) Select(ctx *gin.Context) (int, *Response) {
	var s = new(models.Subject)
	//检测登陆信息是否绑定成功
	if err := ctx.Bind(s); err != nil {
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "解析参数错误",
		}
	}
	if t := sr.Get(s.Uid); t.Id == 0 {
		//选中
		if _, err := sr.Select(s); err != nil {
			return StatusServerError, &Response{
				Code:    ErrorDatabaseInsert,
				Message: "数据库插入失败",
			}
		} else {
			//课容量-1
			var ir = repositories.IssueRepo{}
			ir.Subtraction(s.Iid)
			return StatusOK, &Response{
				Code:    StatusOK,
				Message: "ok",
			}
		}
	} else {
		return StatusServerError, &Response{
			Code:    ErrorDatabaseInsert,
			Message: "你已选课！请不要重复选择",
		}
	}

}

func (this *SubjectUC) Delete(ctx *gin.Context) (int, *Response) {
	var s = new(models.Subject)
	//检测登陆信息是否绑定成功
	s.Id = utils.ParseStringToInt64(ctx.Query("sid"))
	s.Iid = utils.ParseStringToInt64(ctx.Query("iid"))
	//选中
	if _, err := sr.Delete(s); err != nil {
		return StatusServerError, &Response{
			Code:    ErrorDatabaseDelete,
			Message: "数据库删除失败",
		}
	}
	//课容量+1
	var ir = repositories.IssueRepo{}
	ir.Addition(s.Iid)
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
	}
}

func (this *SubjectUC) RelateIssue(ctx *gin.Context) (int, *Response) {
	issue := sr.RelateIssue(utils.ParseStringToInt64(ctx.Query("iid")))
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
		Data:    issue,
	}
}

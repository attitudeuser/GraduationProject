package usecases

import (
	"GraduationProject/models"
	"GraduationProject/repositories"
	"GraduationProject/utils"
	"github.com/gin-gonic/gin"
)

// sr 是是存储层反转过来的对象实例 与存储层的subject model有关
var sr = repositories.NewSubjectRepo()

// SubjectUC 选中课题用例结构
type SubjectUC struct {
}

// NewSubjectUC 返回一个课题选择情况的实例
func NewSubjectUC() SubjectInterface {
	return &SubjectUC{}
}

// Get 获得某用户的选题情况
func (this *SubjectUC) Get(ctx *gin.Context) (int, *Response) {
	var s = new(models.Subject)
	_s := ctx.Query("uid")
	if _s == "null" || _s == "0" {
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "解析参数错误",
		}
	}
	//检测信息是否绑定成功
	uid := utils.ParseStringToInt64(_s)
	//选中
	s = sr.Get(uid)
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
		Data:    s,
	}
}

// Select 选择某个课题
func (this *SubjectUC) Select(ctx *gin.Context) (int, *Response) {
	var s = new(models.Subject)
	//检测信息是否绑定成功
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

// Delete 删除某个已经选中的课题
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

// RelateIssue 处理和issue实例的关系
func (this *SubjectUC) RelateIssue(ctx *gin.Context) (int, *Response) {
	issue := sr.RelateIssue(utils.ParseStringToInt64(ctx.Query("iid")))
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
		Data:    issue,
	}
}

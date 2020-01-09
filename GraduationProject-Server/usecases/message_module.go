package usecases

import (
	"GraduationProject/models"
	"GraduationProject/repositories"
	"GraduationProject/utils"
	"github.com/gin-gonic/gin"
	"time"
)

var mr = repositories.NewMessageRepo()

type MessageUC struct {
}

// NewMessageUC 生成一个message样例的实例
func NewMessageUC() MessageInterface {
	return &MessageUC{}
}

// Send 发送消息
func (this *MessageUC) Send(ctx *gin.Context) (int, *Response) {
	var message = new(models.Message)
	//检测发送信息是否绑定成功
	if err := ctx.Bind(message); err != nil {
		return StatusClientError, &Response{
			Code:    ErrorParameterParse,
			Message: "解析参数错误",
		}
	}
	//写入数据库
	message.SendTime = time.Now()
	if _, err := mr.Send(message); err != nil {
		return StatusServerError, &Response{
			Code:    ErrorDatabaseInsert,
			Message: "数据库插入出错",
			Data:    err,
		}
	}
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "信息发送成功!",
	}
}

// Read 读取消息
func (this *MessageUC) Read(ctx *gin.Context) (int, *Response) {
	var message = new(models.Message)
	message.Id = utils.ParseStringToInt64(ctx.Query("id"))
	message = mr.Read(message.Id)
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "查询成功",
		Data:    message,
	}
}

// Delete 删除消息
func (this *MessageUC) Delete(ctx *gin.Context) (int, *Response) {
	mid := utils.ParseStringToInt64(ctx.Query("id"))
	if _, err := mr.Delete(mid); err != nil {
		return StatusServerError, &Response{
			Code:    ErrorDatabaseDelete,
			Message: "数据库删除出错",
			Data:    err,
		}
	}
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "删除成功",
		Data:    "",
	}
}

// List
func (this *MessageUC) List(ctx *gin.Context) (int, *Response) {
	list := mr.List()
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "OK",
		Data:    list,
	}
}

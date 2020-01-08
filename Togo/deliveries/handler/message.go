package handler

import (
	"GraduationProject/usecases"
	"github.com/gin-gonic/gin"
)

var muc = usecases.NewMessageUC()

// MessageRead 用户读取消息
func MessageRead(this *gin.Context) {
	this.JSON(muc.Read(this))
}

// MessageSend 用户发送消息
func MessageSend(this *gin.Context) {
	this.JSON(muc.Send(this))
}

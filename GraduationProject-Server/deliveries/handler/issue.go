package handler

import (
	"GraduationProject/usecases"
	"github.com/gin-gonic/gin"
)

var iuc = usecases.NewIssueUC()

// IssuePublish 发布issue
func IssuePublish(this *gin.Context) {
	this.JSON(iuc.Publish(this))
}

// IssueUpdate issue更新
func IssueUpdate(this *gin.Context) {
	this.JSON(iuc.Update(this))
}

// IssueList 所有课题列表
func IssueList(this *gin.Context) {
	this.JSON(iuc.List(this))
}

func IssueListById(this *gin.Context) {
	this.JSON(iuc.FindById(this))
}

func IssueDelete(this *gin.Context) {
	this.JSON(iuc.Delete(this))
}

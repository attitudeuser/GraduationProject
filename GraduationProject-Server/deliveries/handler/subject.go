package handler

import (
	"GraduationProject/usecases"
	"github.com/gin-gonic/gin"
)

var suc = usecases.NewSubjectUC()

// GetSubject 用户读取选课信息
func GetSubject(this *gin.Context) {
	this.JSON(suc.Get(this))
}

// GetSubject 用户读取选课信息
func SelectSubject(this *gin.Context) {
	this.JSON(suc.Select(this))
}

// DeleteSubject 用户删除选课
func DeleteSubject(this *gin.Context) {
	this.JSON(suc.Delete(this))
}

// DeleteSubject 用户删除选课
func RelateToIssue(this *gin.Context) {
	this.JSON(suc.RelateIssue(this))
}

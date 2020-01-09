package models

import "time"

const (
	STUDENT = "student" //学生
	TEACHER = "teacher" //教师
	ADMIN   = "admin"   //管理员
)

// User 用户表结构
type User struct {
	Id         int64     `form:"id" json:"id"`
	Username   string    `form:"username" json:"username" xorm:"varchar(32) notnull unique"` //用户名
	Password   string    `form:"password" json:"-" xorm:"varchar(64) notnull"`               //密码
	Email      string    `form:"email" json:"email" xorm:"varchar(128) notnull unique"`      //邮箱
	Phone      string    `form:"phone" json:"phone" xorm:"varchar(11)"`                      //电话号码
	Age        int       `form:"age" json:"age"`                                             //年龄
	Sid        string    `form:"sid" json:"sid" xorm:"varchar(11)"`                          //学号 工号
	College    string    `form:"college" json:"college" xorm:"varchar(16)"`                  //学院
	Grade      string    `form:"grade" json:"grade" xorm:"varchar(16)"`                      //年级
	Type       string    `form:"type" json:"type" xorm:"varchar(8)"`                         //用户类型
	CreateTime time.Time `form:"-" json:"create_time" xorm:"created"`                        //注册时间
}

package models

type Issue struct {
	Id      int64  `json:"id" form:"-"`                                //主键
	Tid     int64  `json:"tid" form:"-"`                               //发布者ID
	Title   string `json:"title" form:"title" xorm:"varchar(64)"`      //标题
	Type    string `json:"type" form:"type" xorm:"varchar(16)"`        //类型
	Require string `json:"require" form:"require" xorm:"varchar(128)"` //要求
	Cap     int    `json:"cap" form:"-"`                               //课程容量 0|1
	Content string `json:"content" form:"content"`                     //内容
}

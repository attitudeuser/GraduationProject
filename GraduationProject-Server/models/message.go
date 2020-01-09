package models

import "time"

// Message 消息结构
type Message struct {
	Id       int64     `json:"id"`                                         //主键
	Fid      int64     `json:"fid" form:"fid"`                             //发送者ID
	Username string    `json:"username" form:"username"`                   //接受者
	Message  string    `json:"message" form:"message" xorm:"varchar(255)"` //消息内容
	Status   bool      `json:"status"`                                     //消息状态 true已读 false未读
	SendTime time.Time `json:"send_time"`                                  //发送时间
}

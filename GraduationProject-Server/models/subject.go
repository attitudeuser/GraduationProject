package models

type Subject struct {
	Id  int64 `json:"id"`             //主键
	Uid int64 `form:"uid" json:"uid"` //用户ID
	Iid int64 `form:"iid" json:"iid"` //课题ID
}

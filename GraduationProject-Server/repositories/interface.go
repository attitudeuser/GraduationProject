package repositories

import (
	"GraduationProject/models"
)

//用户实体-仓库操作-接口
type UserRepoInterface interface {
	//添加用户接口
	Insert(user *models.User) (int64, error)
	//删除用户接口
	Delete(id int64) (int64, error)
	//修改用户接口
	Update(user *models.User) (int64, error)

	// 忘记密码时 发送邮件 将验证参数写入redis 过期时间expires秒
	SetResetArgs(uuid string, uid, expires int64) error
	// 重置忘记密码时 获得参数 检验是否有效
	GetResetArgs(uuid string) (string, error)

	//根据id查询单个用户接口
	FindOneById(id int64) *models.User
	//检测对象是否存在
	RecordExist(u *models.User) (bool, error)
	//检测对象是否存在 并返回user对象
	RecordGet(u *models.User) (bool, *models.User)
	//查询用户列表接口
	FindMany() []*models.User
}

// MessageInterface 消息交流接口
type MessageInterface interface {
	//消息发送接口
	Send(*models.Message) (int64, error)
	//读消息接口
	Read(int64) *models.Message
	//删除消息接口
	Delete(int64) (int64, error)
	List() []*models.Message
}

// IssueInterface 教师出题接口
type IssueInterface interface {
	//发布
	Publish(issue *models.Issue) (int64, error)
	//修改选题
	Update(issue *models.Issue) (int64, error)
	//删除课题
	Delete(id int64) (int64, error)
	//所有出题列表
	List() []*models.Issue
	//罗列某个老师的出题列表
	FindById(id int64) []*models.Issue
}

// SubjectInterface 学生选题管理接口
type SubjectInterface interface {
	Get(uid int64) *models.Subject
	//选题
	Select(subject *models.Subject) (int64, error)
	//删除选题
	Delete(subject *models.Subject) (int64, error)
	RelateIssue(iid int64) *models.Issue
}

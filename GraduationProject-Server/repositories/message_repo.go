package repositories

import (
	"GraduationProject/models"
	"log"
)

// UserRepo 操作用户模型 实现了UserRepoInterface 接口
type MessageRepo struct {
}

// NewMessageRepo 反转 message_repo 的信息
func NewMessageRepo() MessageInterface {
	return &MessageRepo{}
}

// Send 发送信息
func (this *MessageRepo) Send(m *models.Message) (int64, error) {
	return engine.Insert(m)
}

// Read 阅读信息
func (this *MessageRepo) Read(id int64) *models.Message {
	var m = &models.Message{}
	if _, err := engine.ID(id).Get(m); err != nil {
		log.Println("find user error:", err)
	}
	//将status改为false
	m.Status = false
	if _, err := engine.Id(id).Cols("status").Update(m); err != nil {
		log.Println("find user error:", err)
	}
	return m
}

// Delete 删除信息
func (this *MessageRepo) Delete(id int64) (int64, error) {
	return engine.Delete(&models.Message{Id: id})
}

// List 获得消息列表
func (this *MessageRepo) List() []*models.Message {
	all := make([]*models.Message, 0)
	if err := engine.Find(&all); err != nil {
		log.Fatal("find all user error:", err)
	}
	return all
}

package repositories

import (
	"GraduationProject/models"
	"log"
)

// UserRepo 操作用户模型 实现了UserRepoInterface 接口
type MessageRepo struct {
}

func NewMessageRepo() MessageInterface {
	return &MessageRepo{}
}

func (this *MessageRepo) Send(m *models.Message) (int64, error) {
	return engine.Insert(m)
}

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

func (this *MessageRepo) Delete(id int64) (int64, error) {
	return engine.Delete(&models.Message{Id: id})
}

func (this *MessageRepo) List() []*models.Message {
	all := make([]*models.Message, 0)
	if err := engine.Find(&all); err != nil {
		log.Fatal("find all user error:", err)
	}
	return all
}

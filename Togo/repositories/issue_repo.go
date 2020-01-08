package repositories

import (
	"GraduationProject/models"
	"log"
)

type IssueRepo struct {
}

func NewIssueRepo() IssueInterface {
	return &IssueRepo{}
}

// Publish 发布题目
func (this *IssueRepo) Publish(issue *models.Issue) (int64, error) {
	return engine.Insert(issue)
}

func (this *IssueRepo) Update(issue *models.Issue) (int64, error) {
	return engine.Update(issue)
}

func (this *IssueRepo) Delete(id int64) (int64, error) {
	return engine.Delete(&models.Issue{Id: id})
}

func (this *IssueRepo) List() []*models.Issue {
	all := make([]*models.Issue, 0)
	if err := engine.Find(&all); err != nil {
		log.Fatal("find all user error:", err)
	}
	return all
}

func (this *IssueRepo) FindById(id int64) []*models.Issue {
	all := make([]*models.Issue, 0)
	if err := engine.Where("tid = ?", id).Find(&all); err != nil {
		log.Fatal("find all user error:", err)
	}
	return all
}

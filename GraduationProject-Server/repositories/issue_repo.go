package repositories

import (
	"GraduationProject/models"
	"log"
)

// IssueRepo 仓库结构
type IssueRepo struct {
}

// NewIssueRepo 用于反转
func NewIssueRepo() IssueInterface {
	return &IssueRepo{}
}

// Publish 发布题目
func (this *IssueRepo) Publish(issue *models.Issue) (int64, error) {
	return engine.Insert(issue)
}

// Update 更新课题信息
func (this *IssueRepo) Update(issue *models.Issue) (int64, error) {
	return engine.Update(issue)
}

// Addition 将课题容量+1
func (this *IssueRepo) Addition(id int64) bool {
	if _, err := engine.Id(id).Cols("cap").Update(&models.Issue{Cap: 1}); err != nil {
		return false
	}
	return true
}

// Subtraction 将课题容量-1
func (this *IssueRepo) Subtraction(id int64) bool {
	if _, err := engine.Id(id).Cols("cap").Update(&models.Issue{Cap: 0}); err != nil {
		return false
	}
	return true
}

// Delete 删除课题
func (this *IssueRepo) Delete(id int64) (int64, error) {
	return engine.Delete(&models.Issue{Id: id})
}

// List 课题列表
func (this *IssueRepo) List() []*models.Issue {
	all := make([]*models.Issue, 0)
	if err := engine.Find(&all); err != nil {
		log.Fatal("find all user error:", err)
	}
	return all
}

// FindById 通过教师ID 查看课题列表
func (this *IssueRepo) FindById(id int64) []*models.Issue {
	all := make([]*models.Issue, 0)
	if err := engine.Where("tid = ?", id).Find(&all); err != nil {
		log.Fatal("find all user error:", err)
	}
	return all
}

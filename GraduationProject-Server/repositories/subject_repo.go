package repositories

import "GraduationProject/models"

// SubjectRepo 选课信息仓库
type SubjectRepo struct {
}

// NewSubjectRepo 反转选课信息的存储实例
func NewSubjectRepo() SubjectInterface {
	return &SubjectRepo{}
}

// Get 获得用户的选课信息
func (this *SubjectRepo) Get(uid int64) *models.Subject {
	var s = new(models.Subject)
	s.Uid = uid
	engine.Get(s)
	return s
}

// Select 用户选课
func (this *SubjectRepo) Select(subject *models.Subject) (int64, error) {
	return engine.Insert(subject)
}

// Delete 用户删除自己的选课情况
func (this *SubjectRepo) Delete(subject *models.Subject) (int64, error) {
	//更新选择 (修改选择和移除选择)
	return engine.Delete(subject)
}

// RelateIssue 解决subject对issue的依赖
func (this *SubjectRepo) RelateIssue(iid int64) *models.Issue {
	//更新选择 (修改选择和移除选择)
	var issue = new(models.Issue)
	issue.Id = iid
	engine.Get(issue)
	return issue
}

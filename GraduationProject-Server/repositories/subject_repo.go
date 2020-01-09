package repositories

import "GraduationProject/models"

type SubjectRepo struct {
}

func NewSubjectRepo() SubjectInterface {
	return &SubjectRepo{}
}

func (this *SubjectRepo) Get(uid int64) *models.Subject {
	var s = new(models.Subject)
	s.Uid = uid
	engine.Get(s)
	return s
}

func (this *SubjectRepo) Select(subject *models.Subject) (int64, error) {
	return engine.Insert(subject)
}

func (this *SubjectRepo) Delete(subject *models.Subject) (int64, error) {
	//更新选择 (修改选择和移除选择)
	return engine.Delete(subject)
}

func (this *SubjectRepo) RelateIssue(iid int64) *models.Issue {
	//更新选择 (修改选择和移除选择)
	var issue = new(models.Issue)
	issue.Id = iid
	engine.Get(issue)
	return issue
}

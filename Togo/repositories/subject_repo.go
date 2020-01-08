package repositories

import "GraduationProject/models"

type SubjectRepo struct {
}

func NewSubjectRepo() SubjectInterface {
	return &SubjectRepo{}
}

func (this *SubjectRepo) Select(subject *models.Subject) (int64, error) {
	return engine.Insert(subject)
}

func (this *SubjectRepo) Delete(subject *models.Subject) (int64, error) {
	//更新选择 (修改选择和移除选择)
	return engine.Delete(subject)
}

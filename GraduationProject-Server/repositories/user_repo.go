package repositories

import (
	"GraduationProject/models"
	"log"
	"time"
)

// UserRepo 操作用户模型 实现了UserRepoInterface 接口
type UserRepo struct {
}

// NewUserRepo 返回一个UserRepo对象
func NewUserRepo() UserRepoInterface {
	return &UserRepo{}
}

// Insert 将user信息插入数据库
func (this *UserRepo) Insert(u *models.User) (int64, error) {
	return engine.Insert(u)
}

// Delete 根据ID删除user表信息
func (this *UserRepo) Delete(id int64) (int64, error) {
	return engine.Delete(&models.User{Id: id})
}

// Update 更新user信息
func (this *UserRepo) Update(u *models.User) (int64, error) {
	return engine.Update(u)
}

// FindOneById 根据ID查询一个user信息
func (this *UserRepo) FindOneById(id int64) *models.User {
	var u = &models.User{}
	if _, err := engine.ID(id).Get(u); err != nil {
		log.Fatal("find user error:", err)
	}
	return u
}

// RecordExist 检测拥有该属性的对象是否存在
func (this *UserRepo) RecordExist(u *models.User) (bool, error) {
	return engine.Exist(u)
}

// RecordGet 检测是否存在并返回该对象
func (this *UserRepo) RecordGet(u *models.User) (bool, *models.User) {
	has, _ := engine.Get(u)
	return has, u
}

// FindMany 返回用户列表
func (this *UserRepo) FindMany() []*models.User {
	all := make([]*models.User, 0)
	if err := engine.Find(&all); err != nil {
		log.Fatal("find all user error:", err)
	}
	return all
}

// SetResetArgs 忘记密码时 发送邮件 将验证参数写入redis 设置过期时间1h
func (this *UserRepo) SetResetArgs(uuid string, uid, expires int64) error {
	//过期时间为0 则强制为3600
	if expires == 0 {
		expires = 3600
	}
	err := rgo.SetNX(uuid, uid, time.Duration(expires)*time.Second).Err()
	if err != nil {
		panic("redis:" + err.Error())
	}
	return err
}

// GetResetArgs 重置忘记密码时 获得参数 检验是否有效
func (this *UserRepo) GetResetArgs(uuid string) (string, error) {
	return rgo.Get(uuid).Result()
}

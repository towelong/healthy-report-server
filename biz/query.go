package biz

import (
	"errors"
	"fmt"

	"github.com/towelong/healthy-report-server/dal/model"
	"github.com/towelong/healthy-report-server/dal/query"
	"github.com/towelong/healthy-report-server/module"
)

var q = query.Q

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	UserID    int32  `json:"-"`
	SchoolID  string `json:"school_id"`  // 学校ID
	StudentID string `json:"student_id"` // 学生ID
	Address   string `json:"address"`    // 地址
}

func Login(user User) (string, error) {
	t := q.User
	u, err := t.Where(t.Username.Eq(user.Username)).First()
	if err != nil {
		recordError(err)
		return "", errors.New("用户未找到")
	}
	e := module.NewEncryption(module.WithPassword(user.Password))
	if u != nil && e.VerifyPassword(u.Password) {
		return module.GenreateToken(int(u.ID))
	}
	return "", errors.New("密码错误")
}

func Register(user User) error {
	do := q.User
	us, _ := do.Where(do.Username.Eq(user.Username)).First()
	if us != nil {
		return errors.New("用户名已存在")
	}
	var u = model.User{
		Username: user.Username,
		Password: module.NewEncryption(module.WithPassword(user.Password)).EncodePassword(),
	}
	err := do.Create(&u)
	recordError(err)
	if err != nil {
		return errors.New("注册失败")
	}
	return err
}

func FindUserById(id int) (*model.User, error) {
	u := q.User
	return u.Where(u.ID.Eq(int32(id))).First()
}

func UploadInformation(t *Task) error {
	ts := q.Task
	task, _ := ts.Where(ts.StudentID.Eq(t.StudentID), ts.SchoolID.Eq(t.SchoolID)).First()
	if task != nil {
		return errors.New("该任务已存在")
	}
	err := ts.Create(&model.Task{
		UserID:    t.UserID,
		StudentID: t.StudentID,
		SchoolID:  t.SchoolID,
		Address:   t.Address,
	})
	recordError(err)
	if err != nil {
		return errors.New("信息上传失败")
	}
	return nil
}

func EditUserInfomation(t *Task) error {
	ts := q.Task
	info, err := ts.Where(ts.UserID.Eq(t.UserID)).Updates(&model.Task{
		SchoolID:  t.SchoolID,
		StudentID: t.StudentID,
		Address:   t.Address,
	})
	if err != nil {
		recordError(err)
		return errors.New("更新失败")
	}
	if info.RowsAffected == 0 {
		return nil
	}
	return nil
}

func FindUserInformation(id int32) (*model.Task, error) {
	ts := q.Task
	task, err := ts.Where(ts.UserID.Eq(id)).First()
	if err != nil {
		recordError(err)
		return nil, errors.New("无任务")
	}
	return task, nil
}

func FindTaskList() ([]*model.Task, error) {
	do := q.Task
	return do.Find()
}

func recordError(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

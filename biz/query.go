package biz

import (
	"errors"
	"fmt"

	"github.com/towelong/healthy-report-server/dal/model"
	"github.com/towelong/healthy-report-server/dal/query"
	"github.com/towelong/healthy-report-server/db"
	"github.com/towelong/healthy-report-server/module"
)

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
	var q = query.Use(db.DB).User
	u, err := q.Where(q.Username.Eq(user.Username)).First()
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
	var q = query.Use(db.DB).User
	us, _ := q.Where(q.Username.Eq(user.Username)).First()
	if us != nil {
		return errors.New("用户名已存在")
	}
	var u = model.User{
		Username: user.Username,
		Password: module.NewEncryption(module.WithPassword(user.Password)).EncodePassword(),
	}
	err := q.Create(&u)
	recordError(err)
	if err != nil {
		return errors.New("注册失败")
	}
	return err
}

func FindUserById(id int) (*model.User, error) {
	var q = query.Use(db.DB).User
	return q.Where(q.ID.Eq(int32(id))).First()
}

func UploadInformation(t *Task) error {
	var q = query.Use(db.DB).Task
	task, _ := q.Where(q.StudentID.Eq(t.StudentID), q.SchoolID.Eq(t.SchoolID)).First()
	if task != nil {
		return errors.New("该任务已存在")
	}
	err := q.Create(&model.Task{
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

func recordError(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

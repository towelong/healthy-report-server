package biz

import (
	"github.com/towelong/healthy-report-server/dal/model"
	"github.com/towelong/healthy-report-server/dal/query"
	"github.com/towelong/healthy-report-server/db"
)

func Register() error {
	var q = query.Use(db.DB).User
	return q.Create(&model.User{
		Username: "towelong",
		Password: "123123123",
	})
}

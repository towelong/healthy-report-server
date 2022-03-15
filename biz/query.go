package biz

import (
	"github.com/towelong/healthy-report-server/dal/model"
	"github.com/towelong/healthy-report-server/dal/query"
	"github.com/towelong/healthy-report-server/db"
)

var q = query.Use(db.Conn()).User

func Register() error {
	return q.Create(&model.User{
		Username: "towelong",
		Password: "123123123",
	})
}

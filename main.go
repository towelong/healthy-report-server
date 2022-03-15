package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/towelong/healthy-report-server/biz"
	"github.com/towelong/healthy-report-server/dal/model"
	"github.com/towelong/healthy-report-server/db"
	"github.com/towelong/healthy-report-server/module"
	"github.com/towelong/healthy-report-server/server"
)

var wg sync.WaitGroup

func init() {
	db.Conn()
}

func task() {
	tasks, err := biz.FindTaskList()
	if err != nil {
		fmt.Println("没有进行中的任务")
		return
	}
	for _, t := range tasks {
		wg.Add(1)
		go func(t *model.Task) {
			// "4136013436" "江西省九江市共青城市江西农业大学南昌商学院"
			r := module.NewHealthyReport(t.StudentID, t.SchoolID, t.Address)
			err := r.Report()
			fmt.Println(err)
			wg.Done()
		}(t)
	}
	wg.Wait()
}

func main() {
	l := time.FixedZone("CST", 8*3600)
	s := gocron.NewScheduler(l)
	s.Every(1).Day().At("07:00").Do(task)
	s.StartAsync()
	server.Run()
}

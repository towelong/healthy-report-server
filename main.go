package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/towelong/healthy-report-server/server"
	// "github.com/towelong/healthy-report-server/module"
)

func task() {
	fmt.Println("I am running task.")
}

func main() {
	// r := module.NewHealthyReport("19205116", "4136013436", "江西省九江市共青城市江西农业大学南昌商学院")
	// err := r.Report()
	// fmt.Println(err)
	l, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(l)
	s.Every(1).Seconds().Do(task)
	s.StartAsync()
	server.Run()
}

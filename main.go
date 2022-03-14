package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/towelong/healthy-report-server/module"
	"github.com/towelong/healthy-report-server/server"
)

var wg *sync.WaitGroup

func task() {
	jobs := []string{"19205116", "19205118", "19205133"}
	for _, j := range jobs {
		wg.Add(1)
		go func(job string) {
			r := module.NewHealthyReport(job, "4136013436", "江西省九江市共青城市江西农业大学南昌商学院")
			err := r.Report()
			fmt.Println(err)
			wg.Done()
		}(j)
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

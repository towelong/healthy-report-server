package main

import (
	"fmt"

	"github.com/towelong/healthy-report-server/module"
)

func main() {
	r := module.NewHealthyReport("19205116", "4136013436", "江西省九江市共青城市江西农业大学南昌商学院")
	err := r.Report()
	fmt.Println(err)
}

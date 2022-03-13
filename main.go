package main

import (
	"fmt"

	"github.com/towelong/healthy-report-server/module"
)

func main() {
	r := module.NewHealthyReport("19205116", "4136013436")
	err := r.Report()
	fmt.Println(err)
}

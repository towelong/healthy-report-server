package main

import (
	"github.com/towelong/healthy-report-server/db"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dal/query",
		Mode:          gen.WithoutContext,
		FieldNullable: true,
	})

	g.UseDB(db.Conn())

	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)
	g.GenerateModel("user",
		gen.FieldType("delete_time", "gorm.DeletedAt"),
	)
	g.GenerateModel("task",
		gen.FieldType("delete_time", "gorm.DeletedAt"),
	)
	g.Execute()
}

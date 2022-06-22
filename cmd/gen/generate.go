package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/towelong/healthy-report-server/db"
	"gorm.io/gen"
)

var env string

func init() {
	flag.StringVar(&env, "conf", ".env.development", "conf file, eg: -conf .env.development")
}

func main() {
	flag.Parse()
	if err := godotenv.Load(env); err != nil {
		log.Fatalln("load config error...")
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./dal/query",
		Mode:          gen.WithDefaultQuery | gen.WithoutContext,
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

package main

import (
	"log"
	"net/http"

	"http/example/config"
	"http/example/factory"
	"http/example/migrations"
	"http/example/routers"
)

func main() {
	db := config.PostgreSQLConfig()

	migrations.DBMigrator(db)
	dbSQL, _ := db.DB()
	presenter := factory.InitFactory(dbSQL)
	routers.RouterSetup(presenter)

	log.Println("Program Running ..................................")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

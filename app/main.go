package main

import (
	"app/db"
	"app/handler"
	"app/service"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	const (
		dbpath = "root:mysql@tcp(mysql:3306)/app"
		port   = ":8080"
	)

	db, err := db.NewDBMysql(dbpath)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//set http handlers
	svc := service.NewStationService(db)
	h := handler.NewStationHandler(svc)
	hh := handler.NewHelloHandler()
	mux := http.NewServeMux()
	mux.HandleFunc("/stations", h.ServeHTTP)
	mux.HandleFunc("/hello", hh.ServeHTTP)
	log.Fatal(http.ListenAndServe(port, mux))
}

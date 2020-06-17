package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rajaanova/cleancode/internal/controller"
	"github.com/rajaanova/cleancode/internal/persistent"
	"github.com/rajaanova/cleancode/internal/service"
	"net/http"
)

func main() {

	db, err := sql.Open("mysql", "root:raj@tcp(127.0.0.1:3308)/school")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err!= nil {
		panic(err.Error())
	}
	defer db.Close()


	router := mux.NewRouter()
	dataProvider := persistent.NewDataprovider(db)
	ser := service.NewService(dataProvider)
	router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("{'status':'ok'}"))
		writer.WriteHeader(200)
		return
	})
	studentRouter := router.PathPrefix("/cleancode").Subrouter()
	controller.NewGradeKidsController(ser,studentRouter)
	http.ListenAndServe(":8080", router)
}

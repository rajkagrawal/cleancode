package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rajaanova/cleancode/app/bootstrap"
	"github.com/rajaanova/cleancode/app/controller"
	"github.com/rajaanova/cleancode/app/persistent"
	"github.com/rajaanova/cleancode/app/service"
	"net/http"
)

func main() {
	config, err := bootstrap.BS{}.Boot()
	if err != nil {
		panic("failed to start the service "+err.Error())
	}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",config.DatabaseUser,config.DatabasePassword,config.DatabaseHost,config.DatabasePort,config.DatabaseName))
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
		writer.WriteHeader(200)
		writer.Write([]byte("{'status':'ok'}"))
		return
	})
	studentRouter := router.PathPrefix("/cleancode").Subrouter()
	controller.NewGradeKidsController(ser,studentRouter)
	http.ListenAndServe(":8080", router)
}

package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rajaanova/cleancode/internal/entity"
	"github.com/rajaanova/cleancode/internal/service"
	"net/http"
)

type controller struct {
	service *service.Service
	router *mux.Router
}

func NewGradeKidsController(gradeKidsServcie *service.Service,router *mux.Router) *controller {
	cont := &controller{service:gradeKidsServcie,router:router}
	router.HandleFunc("/v1/students", cont.Create).Methods(http.MethodPost)
	return  cont
}

	func (a *controller) Create(w http.ResponseWriter,r *http.Request) {
	var s entity.Student
	err := json.NewDecoder(r.Body).Decode(&s)
	fmt.Println(s)
	fmt.Println(err)
	if err!=nil {
		fmt.Fprintf(w,"not able to update %#v",s )
		return
	}
		if  a.service.CreateStudent(s) != nil {
			fmt.Fprintf(w,"not able to update %#v",s )
			return
		}
	fmt.Fprint(w,"update the value")
}




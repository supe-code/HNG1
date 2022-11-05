package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"HNG-Tasks/task1"
	"HNG-Tasks/task2"
)

func main(){
	r:= mux.NewRouter()
	r.HandleFunc("/", task1.Task).Methods("GET")
	r.HandleFunc("/", task2.Task).Methods("POST")
	http.Handle("/", r)
	fmt.Println("server running on port: 5000")
	http.ListenAndServe(":5000",nil)
}
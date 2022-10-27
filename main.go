package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type response struct{
	SlackUsername string `json:"slackUsername"`
	Backend bool `json:"backend"`
	Age int `json:"age"`
	Bio string `json:"bio"`
}

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-type","application/json");
		resStruct:=response{
			SlackUsername: "Arigbede Jacob",
			Backend: true,
			Age: 24,
			Bio: "Young and Hopeful software engineer to be",
		}
		res,err:= json.Marshal(resStruct)
		if err!=nil{
			log.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})
	fmt.Println("server running on port: 5000....")
	http.ListenAndServe(":5000",nil)
}
package handler1

import(
	"encoding/json"
	"net/http"
	"log"
	help "HNG-Tasks/task1/helper1"
)

func Response(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json");
	resStruct:= help.Response{
		SlackUsername: "Arigbede Jacbo",
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
}
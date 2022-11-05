package handler2

import (
	help "HNG-Tasks/task2/helper2"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func Response(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	req:= &help.Req
	err:=json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		log.Println(err)
	}
	res := &help.Res
	switch req.Operation{
	case "addition":
		val:= help.Addition(req.X,req.Y)
		res.Result = val
		res.OperationType = "addition"
		res.Key = os.Getenv("open-key")
	case "subtraction":
		val:= help.Subtraction(req.X,req.Y)
		res.Result = val
		res.OperationType = "subtraction"
	case "multiplication":
		val:= help.Multiplication(req.X,req.Y)
		res.Result = val
		res.OperationType = "multiplication"
	default:
		val,err:= help.GPTVal(req.Operation)
		if err!=nil{
			log.Println(err)
		}
		opr,err:= help.GPTOpr(req.Operation)
		res.Result = val
		res.OperationType = opr
	}
	resJson,err:= json.Marshal(res)
	if err!= nil{
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resJson)
}
package task2

import(
	"net/http"
	handler "HNG-Tasks/task2/handler2"
)

func Task(w http.ResponseWriter,r *http.Request){
	handler.Response(w,r)
}
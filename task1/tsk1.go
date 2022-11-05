package task1

import(
	handler "HNG-Tasks/task1/handler1"
	"net/http"
)

func Task(w http.ResponseWriter, r *http.Request){
	handler.Response(w,r)
}
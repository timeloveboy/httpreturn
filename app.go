package main

import (
	"os"
	"net/http"
)
var response string
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(response))
}
func main(){
	http.HandleFunc("/",handler)
	response=os.Args[1]
	err:=http.ListenAndServe(":9000", nil)
	if(err!=nil){
		panic(err)
	}
}
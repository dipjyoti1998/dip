package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main(){
	fmt.Println("Ddd")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello",hello)
	err:=http.ListenAndServe(":8080",mux)
	if err!=nil{
		fmt.Println("failed")
	}

}
func hello(w http.ResponseWriter,req *http.Request){
	fmt.Println("HelloWorld")
	body,_:=ioutil.ReadAll(req.Body)
	responseString:= "Hello "+string(body)
	w.Write([]byte(responseString))
}
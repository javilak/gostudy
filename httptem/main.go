package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hellof(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hellof.tmpl")
	if err != nil {
		fmt.Printf("parse file err %v/n", err)
	}
	t.Execute(w, "你好3")
}
func main() {
	http.HandleFunc("/hello", hellof)
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		fmt.Printf("http server start failed: %v/n", err)
	}
}

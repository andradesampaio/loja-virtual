package main

import (
	"net/http"
	"loja-virtual/internal/controller"
)


func main() {	
	http.HandleFunc("/", controller.Index) 
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/update", controller.Update)
	http.ListenAndServe(":8080", nil)
}
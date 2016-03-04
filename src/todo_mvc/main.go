package main

import (
	"fmt"
	"log"
	"net/http"
	"todo_mvc/controllers"
)

var port = "8888"

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/finish", controllers.Finish)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/undo", controllers.Undo)
	fmt.Println("Welcome to TODO web service... Please access port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe failed: ", err)
	}
}

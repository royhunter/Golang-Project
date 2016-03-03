package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)


type event_item struct {
	Item   string
	Finish bool
	Id     string
}

var event_count = 2

var event_list = map[string] *event_item {
	"0": &event_item{Item: "aaa", Finish: false, Id: "0"},
	"1": &event_item{Item: "bbb", Finish: false, Id: "1"},
}

func req_print(r *http.Request) {
	fmt.Println("Method: ", r.Method)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("Proto: ", r.Proto)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println("Header: ", r.Header)
	fmt.Println("Content-Type: ", r.Header["Content-Type"])
	r.ParseForm()
	fmt.Println("Form: ", r.Form)
	fmt.Println("ContentLength: ", r.ContentLength)
	fmt.Println("Body: ")
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))
}

func render(w http.ResponseWriter, el map[string] *event_item) {
	t, _ := template.ParseFiles("todo.html")
	t.Execute(w, el)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index method: ", r.Method)
	if r.Method == "GET" {
		render(w, event_list)
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New method: ", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		text := r.FormValue("event")
		if text == "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		id := event_count
		event_count++
		event_list[strconv.Itoa(id)] = &event_item{
			Item: text,
			Finish: false,
			Id: strconv.Itoa(id),
		}
		render(w, event_list)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete  method: ", r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		id := r.FormValue("id")
		delete(event_list, id)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func Undo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Undo method: ", r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		id := r.FormValue("id")
		value, ok := event_list[id] 
		if ok {
			value.Finish = false
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}

func Finish(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Finish  method: ", r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		id := r.FormValue("id")
		value, ok := event_list[id] 
		if ok {
			value.Finish = true
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/new", New)
	http.HandleFunc("/finish", Finish)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/undo", Undo)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe failed: ", err)
	}
}

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	//"crypto/md5"
	//"strconv"
	//"time"
	//"io"
)

const (
	INIT   = 1
	FINISH = 2
)

type event_item struct {
	Item   string
	Status int
	Id     string
}

var event_count = 2

var event_list = map[string]event_item{
	"0": event_item{Item: "aaa", Status: INIT, Id: "0"},
	"1": event_item{Item: "bbb", Status: FINISH, Id: "1"},
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

func render(w http.ResponseWriter, el map[string]event_item) {
	t, _ := template.ParseFiles("todo.html")
	t.Execute(w, el)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		render(w, event_list)
	} else {
		//req_print(r)
		r.ParseForm()
		//text := r.FormValue("event")
		text := r.FormValue("event")
		//id := ev.Count
		//ev.Count++
		//ev_it := event_item{Item: text, Status: INIT, Id: id}
		//ev.Event = append(ev.Event, ev_it)
		fmt.Println(text)
		render(w, event_list)
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		text := r.FormValue("event")
		//id := ev.Count
		//ev.Count++
		//ev_it := event_item{Item: text, Status: INIT, Id: id}
		//ev.Event = append(ev.Event, ev_it)
		fmt.Println(text)
		render(w, event_list)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

}

func Undo(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/new", New)
	http.HandleFunc("/del", Delete)
	http.HandleFunc("/undo", Undo)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe failed: ", err)
	}
}

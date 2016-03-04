package controllers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"todo_mvc/models"
)

var db_path = "./data/todo.sqlite"

func showError(w http.ResponseWriter, message string) {
	fmt.Fprintf(w, message)
}

/*
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
*/

func render(w http.ResponseWriter, el map[string]*models.Event_item) {
	t, _ := template.ParseFiles("views/todo.html")
	t.Execute(w, el)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index method: ", r.Method)
	//req_print(r)
	if r.Method == "GET" {
		if r.URL.Path == "/" {
			ev_list := models.QueryAll(db_path)
			render(w, ev_list)
		} else {
			showError(w, "404, unknow link")
		}

	} else {
		showError(w, "Wront method")
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("New method: ", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		text := r.FormValue("event")
		if text == "" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		id := models.Event_count
		models.Event_count++
		eii := models.Event_item{
			Item:   text,
			Finish: 0,
			Id:     strconv.Itoa(id),
		}
		models.InsertItem(eii, db_path)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Delete  method: ", r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		id := r.FormValue("id")
		models.DeleteItem(id, db_path)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func Undo(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Undo method: ", r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		id := r.FormValue("id")
		models.UpdateItem(id, 0, db_path)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func Finish(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Finish  method: ", r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		id := r.FormValue("id")
		models.UpdateItem(id, 1, db_path)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

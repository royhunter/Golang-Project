package controllers

import (
	//"fmt"
	"html/template"
	"net/http"
	"strconv"
	"todo_mvc/models"
)

var db_path = "./data/todo.sqlite"

func render(w http.ResponseWriter, el map[string]*models.Event_item) {
	t, _ := template.ParseFiles("views/todo.html")
	t.Execute(w, el)
}

func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Index method: ", r.Method)
	if r.Method == "GET" {
		ev_list := models.QueryAll(db_path)
		render(w, ev_list)
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

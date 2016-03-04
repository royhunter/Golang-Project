package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type event_item struct {
	Item   string
	Finish int
	Id     string
}

var event_count = 2

func QueryAll() map[string]*event_item {
	db, err := sql.Open("sqlite3", "./todo.sqlite")
	if err != nil {
		log.Fatal("QueryAll db open failed", err)
	}
	defer db.Close()

	//sqlstring := `create table TODO (ID text PRIMARY KEY, ITEM text, FINISH int);`
	//db.Exec(sqlstring)

	sqlx := `SELECT * FROM TODO;`
	rows, err := db.Query(sqlx)
	defer rows.Close()

	event_list := make(map[string]*event_item)
	for rows.Next() {
		var item string
		var finish int
		var id string
		rows.Scan(&id, &item, &finish)
		event_list[id] = &event_item{item, finish, id}
	}
	return event_list
}

func InsertItem(ei event_item) {
	db, err := sql.Open("sqlite3", "./todo.sqlite")
	if err != nil {
		log.Fatal("InsertItem db open failed", err)
	}
	defer db.Close()

	item := ei.Item
	finish := ei.Finish
	id := ei.Id

	sql := `INSERT INTO TODO(ID, ITEM, FINISH) VALUES(` + id + `,` + item + `,` + strconv.Itoa(finish) + `)`
	db.Exec(sql)
}

func DeleteItem(id string) {
	db, err := sql.Open("sqlite3", "./todo.sqlite")
	if err != nil {
		log.Fatal("InsertItem db open failed", err)
	}
	defer db.Close()

	sql := `DELETE FROM TODO WHERE ID=` + id
	db.Exec(sql)
}

func UpdateItem(id string, fin int) {
	db, err := sql.Open("sqlite3", "./todo.sqlite")
	if err != nil {
		log.Fatal("InsertItem db open failed", err)
	}
	defer db.Close()

	sql := `UPDATE TODO SET finish=` + strconv.Itoa(fin) + ` WHERE ID=` + id
	db.Exec(sql)
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

func render(w http.ResponseWriter, el map[string]*event_item) {
	t, _ := template.ParseFiles("todo.html")
	t.Execute(w, el)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index method: ", r.Method)
	if r.Method == "GET" {
		ev_list := QueryAll()
		render(w, ev_list)
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
		eii := event_item{
			Item:   text,
			Finish: 0,
			Id:     strconv.Itoa(id),
		}
		InsertItem(eii)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete  method: ", r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		id := r.FormValue("id")
		DeleteItem(id)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func Undo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Undo method: ", r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		id := r.FormValue("id")
		UpdateItem(id, 0)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func Finish(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Finish  method: ", r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		id := r.FormValue("id")
		UpdateItem(id, 1)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
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

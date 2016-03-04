package models

import (
	"database/sql"
	//"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

type Event_item struct {
	Item   string
	Finish int
	Id     string
}

var Event_count = 2

func QueryAll(dbname string) map[string]*Event_item {
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal("QueryAll db open failed", err)
	}
	defer db.Close()

	//sqlstring := `create table TODO (ID text PRIMARY KEY, ITEM text, FINISH int);`
	//db.Exec(sqlstring)

	sql := `SELECT * FROM TODO;`
	rows, err := db.Query(sql)
	defer rows.Close()

	event_list := make(map[string]*Event_item)
	for rows.Next() {
		var item string
		var finish int
		var id string
		rows.Scan(&id, &item, &finish)
		event_list[id] = &Event_item{item, finish, id}
	}
	return event_list
}

func InsertItem(ei Event_item, dbname string) {
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal("InsertItem db open failed", err)
	}
	defer db.Close()

	item := ei.Item
	finish := ei.Finish
	id := ei.Id

	sql := `INSERT INTO TODO(ID, ITEM, FINISH) VALUES(` + id + `,` + `"` + item + `"` + `,` + strconv.Itoa(finish) + `)`
	//fmt.Println(sql)
	db.Exec(sql)
}

func DeleteItem(id string, dbname string) {
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal("InsertItem db open failed", err)
	}
	defer db.Close()

	sql := `DELETE FROM TODO WHERE ID=` + id
	db.Exec(sql)
}

func UpdateItem(id string, fin int, dbname string) {
	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal("InsertItem db open failed", err)
	}
	defer db.Close()

	sql := `UPDATE TODO SET finish=` + strconv.Itoa(fin) + ` WHERE ID=` + id
	db.Exec(sql)
}

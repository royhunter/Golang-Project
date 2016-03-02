package main 


import (
	"net/http"
	"log"
	"html/template"
	"fmt"
	"io/ioutil"
	//"crypto/md5"
	//"strconv"
	//"time"
	//"io"
)

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

func todo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("todo.html")
		t.Execute(w, nil)
	} else {
		//req_print(r)
		r.ParseForm()
		for _, v := range r.Form["event"] {
			fmt.Println("event: ", v)
		}	
	}
}

func main() {
	http.HandleFunc("/", todo)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe failed: ", err)
	}

}
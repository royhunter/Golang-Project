package main

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "log"
    "net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Method: ", r.Method)
    fmt.Println("path: ", r.URL.Path)
    fmt.Println("Proto: ", r.Proto)
    fmt.Println("scheme: ", r.URL.Scheme)
    fmt.Println("Header: ", r.Header)
    fmt.Println("Content-Type: ", r.Header["Content-Type"])
    r.ParseForm()
    fmt.Println("Form: ", r.Form)
    fmt.Println("ContentLength: ", r.ContentLength)
    w.Header().Add("AtEnd1", "value 1")
    fmt.Println("Body: ")
    defer r.Body.Close()
    b, _ := ioutil.ReadAll(r.Body)
    fmt.Println(string(b))
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Hello, royluo")
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method: ", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        fmt.Println("username: ", r.Form["username"])
        fmt.Println("password: ", r.Form["password"])
    }
}

func main() {
    http.HandleFunc("/", sayhelloName)
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}


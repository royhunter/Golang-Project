package main

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "log"
    "net/http"
    "time"
    "crypto/md5"
    "strconv"
    "io"
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

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    req_print(r)
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

func upload(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method: ", r.Method)
    if r.Method == "GET" {
        req_print(r)
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t, _ := template.ParseFiles("upload.gtpl")
        t.Execute(w, token)
    } else {
        
    }
}

func main() {
    http.HandleFunc("/", sayhelloName)
    http.HandleFunc("/login", login)
    http.HandleFunc("/upload", upload)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}


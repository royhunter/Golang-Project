package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

func resp_print(resp *http.Response) {
    defer resp.Body.Close()
    fmt.Println("Status: ", resp.Status)
    fmt.Println("StatusCode: ", resp.StatusCode)
    fmt.Println("Proto: ", resp.Proto)
    fmt.Println("ContentLength: ", resp.ContentLength)
    fmt.Println("Header: ", resp.Header)
    b, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Body: ", string(b))
}

func client_get() {
    //client := &http.Client{}
    //resp, err := client.Get("http://localhost:9090")
    resp, err := http.Get("http://localhost:9090")
    if err != nil {
        fmt.Println(err)
        log.Fatal("client Get failed")
    }

    resp_print(resp)
}

func client_post() {

    json_string := "[" + "{\"id\": 912345678901," + "\"text\":\"How do I read JSON onAndroid?\"," + "\"geo\":null," + "\"user\":{\"name\":\"android_newb\",\"followers_count\":41}}," + "{\"id\": 912345678902," + "\"text\":\"@android_newb just useandroid.util.JsonReader!\"," + "\"geo\":[50.454722,-104.606667]," + "\"user\":{\"name\":\"jesse\",\"followers_count\":2}}" + "]"

    resp, err := http.Post("http://localhost:9090", "json", strings.NewReader(json_string))
    if err != nil {
        fmt.Println(err)
        log.Fatal("client post failed")
    }

    resp_print(resp)
}

func main() {
    //client_get()
    client_post()
}


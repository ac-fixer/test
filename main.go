package main

import (
    "net/http"
    "log"
    "fmt"
    "os" 
)

func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello there!\n")
    getHost()
}

func getHost() {
    name, err := os.Hostname()
    if err != nil {
        panic(err)
    }
    fmt.Println("hostname:", name)
}

func main(){
    http.HandleFunc("/", myHandler)		//	设置访问路由
    log.Fatal(http.ListenAndServe(":80", nil))
}

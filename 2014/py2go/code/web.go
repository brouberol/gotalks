// Minimal boilerplate to implement a web server in Go

package main

import "net/http"

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello, world!\n"))
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8888", nil)
}

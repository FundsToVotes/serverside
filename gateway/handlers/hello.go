package handlers

import "net/http"

//HelloHandler is a very basic handler for the purpose of getting a bare bones server up and running
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
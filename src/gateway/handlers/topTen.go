package handlers

import "net/http"

//HelloHandler is a very basic handler for the purpose of getting a bare bones server up and running
func NotYetWorkingTopTenHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("No content here yet!"))
}

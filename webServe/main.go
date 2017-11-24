package main

import (
	"net/http"
)

func homeAction(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("<h1>Auto Deploy!</h1>"))
}

func main() {
	http.HandleFunc("/", homeAction)

	http.ListenAndServe(":8000", nil)
}

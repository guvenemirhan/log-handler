package main

import (
	"net/http"
	"real-time-log-analyze/api"
)

func main() {
	http.HandleFunc("/", api.Handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

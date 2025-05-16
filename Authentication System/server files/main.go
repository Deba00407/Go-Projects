package main

import (
	"auth/internal/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/post", auth.MakePostRequestToAPIEndpoint)
	http.ListenAndServe(":4001", router)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/createsong", songExplorerCreate).Methods("POST")
	router.HandleFunc("/updatesong", songExplorerUpdate).Methods("POST")
	router.HandleFunc("/deletesong", songExplorerDelete).Methods("POST")
	http.ListenAndServe(":8000", router)
}

func songExplorerCreate(w http.ResponseWriter, r *http.Request) {

	fmt.Println("WIP")
}

func songExplorerUpdate(w http.ResponseWriter, r *http.Request) {

	fmt.Println("WIP")
}

func songExplorerDelete(w http.ResponseWriter, r *http.Request) {

	fmt.Println("WIP")
}

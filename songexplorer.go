package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database Connection
var dsn = "mike:Password1@tcp(benchmarktest01.mysql.database.azure.com:3306)/benchmarks?charset=utf8mb4"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type SongExplorerModel struct {
	Year     string
	SongName string
}

func main() {

	// fmt.Println(db)

	router := mux.NewRouter()
	router.HandleFunc("/createsong", songExplorerCreate).Methods("POST")
	router.HandleFunc("/updatesong", songExplorerUpdate).Methods("POST")
	router.HandleFunc("/deletesong", songExplorerDelete).Methods("POST")
	router.HandleFunc("/getsong", songExplorerGet).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func songExplorerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("WIP")
}

func songExplorerCreate(w http.ResponseWriter, r *http.Request) {

	todo := &SongExplorerModel{Year: "2021", SongName: "MySong"}
	db.Create(&todo)

	json.NewEncoder(w).Encode(todo)

	fmt.Println("Endoint hit: ", todo)
}

func songExplorerUpdate(w http.ResponseWriter, r *http.Request) {

	fmt.Println("WIP")
}

func songExplorerDelete(w http.ResponseWriter, r *http.Request) {

	fmt.Println("WIP")
}

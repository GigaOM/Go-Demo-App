package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	http.HandleFunc("/createsong", songExplorerCreate)
	http.HandleFunc("/updatesong", songExplorerUpdate)
	http.HandleFunc("/deletesong", songExplorerDelete)
	http.HandleFunc("/getsong", songExplorerGet)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func songExplorerGet(w http.ResponseWriter, r *http.Request) {
	var songs SongExplorerModel

	song := db.Find(&songs)
	fmt.Print(song)
}

func songExplorerCreate(w http.ResponseWriter, r *http.Request) {

	todo := SongExplorerModel{
		Year:     "2021",
		SongName: "MySong"}

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

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
	gorm.Model
	Year     string
	SongName string
}

func main() {

	http.HandleFunc("/createsong", songExplorerCreate)
	http.HandleFunc("/updatesongyear", songExplorerUpdateYear)
	http.HandleFunc("/updatesongname", songExplorerUpdateSongName)
	http.HandleFunc("/deletesong", songExplorerDelete)
	http.HandleFunc("/getsong", songExplorerGet)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func songExplorerGet(w http.ResponseWriter, r *http.Request) {

	song := db.QueryFields
	fmt.Print(song)
}

func songExplorerCreate(w http.ResponseWriter, r *http.Request) {

	songExplorer := SongExplorerModel{
		Year:     "2021",
		SongName: "MySong"}

	db.Create(&songExplorer)

	json.NewEncoder(w).Encode(songExplorer)

	fmt.Println("Fields Added ", songExplorer)
}

func songExplorerUpdateYear(w http.ResponseWriter, r *http.Request) {
	songExplorer := SongExplorerModel{
		Year:     "",
		SongName: "",
	}

	year := ""
	songExplorer.Year = year

}

func songExplorerUpdateSongName(w http.ResponseWriter, r *http.Request) {
	songExplorer := SongExplorerModel{
		Year:     "",
		SongName: "",
	}

	songName := ""
	songExplorer.SongName = songName
}

func songExplorerDelete(w http.ResponseWriter, r *http.Request) {

	fmt.Println("WIP")
}

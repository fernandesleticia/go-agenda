package main

import (
	"encoding/json"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var db, _ = gorm.Open("mysql", "root:root@/agenda?charset=utf8&parseTime=True&loc=Local")

type AgendaItemModel struct {
	Id          int `gorm:primary_key`
	Description string
	Done        bool
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	log.Info("All good with Agenda API")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `I'am alive`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	log.WithFields(log.Fields{"description": description}).Info("Add a new item")
	agenda_item := &AgendaItemModel{Description: description, Done: false}
	db.Create(&agenda_item)
	result := db.Last(&agenda_item)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Value)
}

func main() {
	defer db.Close()

	db.Debug().DropTableIfExists(&AgendaItemModel{})
	db.Debug().AutoMigrate(&AgendaItemModel{})

	log.Info("Starting agenda")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	router.HandleFunc("/item", CreateItem).Methods("POST")
	http.ListenAndServe(":8000", router)
}

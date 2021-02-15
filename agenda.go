package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/fernandesleticia/go-agenda/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var db, _ = gorm.Open("mysql", "root:root@/agenda?charset=utf8&parseTime=True&loc=Local")

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
	item := &database.Item{Description: description, Done: false}
	db.Create(&item)
	result := db.Last(&item)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Value)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	has_item := GetItemByID(id)
	if has_item == false {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"updated": false, "error": Recorde Not Found}`)
	} else {
		done, _ := strconv.ParseBool(r.FormValue("done"))
		log.WithFields(log.Fields{"Id": id, "Done": done}).Info("Updating item")
		item := &database.Item{}
		db.First(&item, id)
		item.Done = done
		db.Save(&item)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"updated": true}`)
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	has_item := GetItemByID(id)
	if has_item == false {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": false, "error": "Record not found"}`)
	} else {
		log.WithFields(log.Fields{"Id": id}).Info("Deleting item")
		item := &database.Item{}
		db.First(&item, id)
		db.Delete(&item)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": true}`)
	}

}

func GetItemByID(Id int) bool {
	item := &database.Item{}
	result := db.First(&item, Id)
	if result.Error != nil {
		log.Warn("Item not found in database")
		return false

	}
	return true
}

func GetDoneItems(w http.ResponseWriter, r *http.Request) {
	log.Info("Getting done items")
	done_items := GetItemsWith(true)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(done_items)
}

func GetPendingItems(w http.ResponseWriter, r *http.Request) {
	log.Info("Getting pending items")
	pending_items := GetItemsWith(false)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pending_items)
}

func GetItemsWith(done bool) interface{} {
	var items []database.Item
	Items := db.Where("done = ?", done).Find(&items).Value
	return Items
}

func main() {
	defer db.Close()

	db.Debug().DropTableIfExists(&database.Item{})
	db.Debug().AutoMigrate(&database.Item{})

	log.Info("Starting agenda")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	router.HandleFunc("/item", CreateItem).Methods("POST")
	router.HandleFunc("/update/{id}", UpdateItem).Methods("POST")
	router.HandleFunc("/delete/{id}", DeleteItem).Methods("DELETE")
	router.HandleFunc("/done", GetDoneItems).Methods("GET")
	router.HandleFunc("/pending", GetPendingItems).Methods("GET")
	http.ListenAndServe(":8000", router)
}

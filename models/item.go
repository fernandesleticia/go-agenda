package models

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/fernandesleticia/go-agenda/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Item struct {
	Id          int `gorm:primary_key`
	Description string
	Done        bool
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	log.WithFields(log.Fields{"description": description}).Info("Add a new item")
	item := &Item{Description: description, Done: false}
	database.MysqlInstance.Create(&item)
	result := database.MysqlInstance.Last(&item)
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
		item := &Item{}
		database.MysqlInstance.First(&item, id)
		item.Done = done
		database.MysqlInstance.Save(&item)
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
		item := &Item{}
		database.MysqlInstance.First(&item, id)
		database.MysqlInstance.Delete(&item)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"deleted": true}`)
	}
}

func GetItemByID(Id int) bool {
	item := &Item{}
	result := database.MysqlInstance.First(&item, Id)
	if result.Error != nil {
		log.Warn("Item not found")
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
	var items []Item
	Items := database.MysqlInstance.Where("done = ?", done).Find(&items).Value
	return Items
}

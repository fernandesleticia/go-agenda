package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/fernandesleticia/go-agenda/database"
	"github.com/fernandesleticia/go-agenda/models"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	log.WithFields(log.Fields{"description": description}).Info("Adding a new item")

	item := &models.Item{Description: description, Done: false}
	database.MysqlInstance.Create(&item)
	result := database.MysqlInstance.Last(&item)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Value)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	w.Header().Set("Content-Type", "application/json")
	hasItem := models.GetItemByID(id)

	if hasItem == false {
		io.WriteString(w, `{"updated": false, "error": Recorde Not Found}`)
	} else {
		done, _ := strconv.ParseBool(r.FormValue("done"))
		log.WithFields(log.Fields{"Id": id, "Done": done}).Info("Updating item")
		item := &models.Item{}
		database.MysqlInstance.First(&item, id)
		item.Done = done
		database.MysqlInstance.Save(&item)

		io.WriteString(w, `{"updated": true}`)
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	w.Header().Set("Content-Type", "application/json")
	hasItem := models.GetItemByID(id)

	if hasItem == false {
		io.WriteString(w, `{"deleted": false, "error": "Record not found"}`)
	} else {
		log.WithFields(log.Fields{"Id": id}).Info("Deleting item")
		item := &models.Item{}
		database.MysqlInstance.First(&item, id)
		database.MysqlInstance.Delete(&item)

		io.WriteString(w, `{"deleted": true}`)
	}
}

func GetDoneItems(w http.ResponseWriter, r *http.Request) {
	log.Info("Getting done items")
	doneItems := models.GetItemsWith(true)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doneItems)
}

func GetPendingItems(w http.ResponseWriter, r *http.Request) {
	log.Info("Getting pending items")
	pendingItems := models.GetItemsWith(false)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pendingItems)
}

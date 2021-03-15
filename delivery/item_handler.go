package delivery

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/fernandesleticia/go-agenda/database"
	"github.com/fernandesleticia/go-agenda/models"
	"github.com/fernandesleticia/go-agenda/usecase"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type ItemHandler interface {
	CreateItem(description string, done bool) ([]*models.Item, error)
	UpdateItem(id int, done bool) ([]*models.Item, error)
	DeleteItem(id int) ([]*models.Item, error)
	GetDoneItems() (bool, error)
	GetPendingItems() (bool, error)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	description := r.FormValue("description")
	log.WithFields(log.Fields{"description": description}).Info("Adding a new item")

	result := usecase.CreateItem(description, false)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Value)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	done, _ := strconv.ParseBool(r.FormValue("done"))
	message := ""

	updated := usecase.UpdateItem(id, done)

	if updated == false {
		message = `{"updated": false, "error": Recorde Not Found}`
	} else {
		message = `{"updated": true}`
	}

	io.WriteString(w, message)
	w.Header().Set("Content-Type", "application/json")
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	message := ""

    deleted = usecase.DeleteItem(id)

	if deleted == false {
		message = `{"deleted": false, "error": "Record not found"}`
	} else {
		message = `{"deleted": true}`
	}
    
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, message)
}

func GetDoneItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usecase.GetDoneItems)
}

func GetPendingItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usecase.GetPendingItems)
}
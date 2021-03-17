package usecase

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

type ItemUsecase interface {
	CreateItem(description string, done bool) ([]*models.Item, error)
	UpdateItem(id int, done bool) ([]*models.Item, error)
	DeleteItem(id int) ([]*models.Item, error)
	GetDoneItems() (bool, error)
	GetPendingItems() (bool, error)
}

func CreateItem(description string, done bool) result {
	item := &models.Item{Description: description, Done: done}
	database.MysqlInstance.Create(&item)
	return database.MysqlInstance.Last(&item)
}

func UpdateItem(id int) updated {
	hasItem := models.GetItemByID(id)

	if hasItem == false {
		return false
	} else {
		log.WithFields(log.Fields{"Id": id, "Done": done}).Info("Updating item")
		item := &models.Item{}
		database.MysqlInstance.First(&item, id)
		item.Done = done
		database.MysqlInstance.Save(&item)

		return true
	}
}

func DeleteItem(id int) deleted {
	hasItem := models.GetItemByID(id)

	if hasItem == false {
		return false
		io.WriteString(w, `{"deleted": false, "error": "Record not found"}`)
	} else {
		log.WithFields(log.Fields{"Id": id}).Info("Deleting item")
		item := &models.Item{}
		database.MysqlInstance.First(&item, id)
		database.MysqlInstance.Delete(&item)

		io.WriteString(w, `{"deleted": true}`)
		return true
	}
}

func GetDoneItems() doneItems {
	log.Info("Getting done items")
	doneItems := models.GetItemsWith(true)
	return doneItems
}

func GetPendingItems() pendingItems {
	log.Info("Getting pending items")
	pendingItems := models.GetItemsWith(false)
    return pendingItems
}

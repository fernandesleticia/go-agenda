package main

import (
	"io"
	"net/http"

	"github.com/fernandesleticia/go-agenda/database"
	"github.com/fernandesleticia/go-agenda/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	log.Info("All good with Agenda API")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `I'am alive`)
}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	defer database.MysqlInstance.Close()

	database.MysqlInstance.Debug().DropTableIfExists(&models.Item{})
	database.MysqlInstance.Debug().AutoMigrate(&models.Item{})

	log.Info("Starting agenda")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	router.HandleFunc("/item", models.CreateItem).Methods("POST")
	router.HandleFunc("/update/{id}", models.UpdateItem).Methods("POST")
	router.HandleFunc("/delete/{id}", models.DeleteItem).Methods("DELETE")
	router.HandleFunc("/done", models.GetDoneItems).Methods("GET")
	router.HandleFunc("/pending", models.GetPendingItems).Methods("GET")
	http.ListenAndServe(":8000", router)
}

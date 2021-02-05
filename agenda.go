package main

import (
	"io"
	"net/http"

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
	log.Info("Starting agenda")
}

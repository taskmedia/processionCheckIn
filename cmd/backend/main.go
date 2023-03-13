package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/processionCheckIn/pkg/backend"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db"
	"github.com/taskmedia/processionCheckIn/pkg/cmd/parameter"
	"github.com/taskmedia/processionCheckIn/pkg/version"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	log.WithFields(log.Fields{
		"version": version.VERSION,
		"commit":  version.COMMIT,
	}).Info("Version information")

	flag.Parse()

	db.InitDbConn()

	log.Info(*parameter.LogLevel)

	err := backend.StartRouter()
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/processionCheckIn/pkg/backend"
	"github.com/taskmedia/processionCheckIn/pkg/cmd/parameter"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	flag.Parse()

	log.Info(*parameter.LogLevel)

	err := backend.StartRouter()
	if err != nil {
		log.Fatal(err)
	}
}

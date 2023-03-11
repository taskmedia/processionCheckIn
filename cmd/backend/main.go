package main

import (
	"fmt"
	"net/http"

	"flag"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

type AppParameters struct {
	host     string
	port     int
	loglevel string
	// fromStdin  bool
}

var parameters AppParameters

func main() {

	flag.StringVar(&parameters.host, "host", "0.0.0.0", "Host to listen on")
	flag.IntVar(&parameters.port, "port", 8080, "Port to listen on")
	// flag.StringVar(&parameters.loglevel, "loglevel", "default", "Log level")
	flag.Parse()

	log.Info(parameters)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf("%s:%d", parameters.host, parameters.port))
}

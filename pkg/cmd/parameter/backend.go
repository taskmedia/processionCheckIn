package parameter

import (
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// server
var Host = flag.String("host", "0.0.0.0", "Host to listen on")
var Port = flag.Int("port", 8080, "Port to listen on")

// logging
var LogLevel = flag.String("loglevel", "default", "Log level")

// PSQL
var PsqlHost = flag.String("psqlhost", "localhost", "PostgreSQL host")
var PsqlPort = flag.Int("psqlport", 5432, "PostgreSQL port")
var PsqlUser = flag.String("psqluser", "postgres", "PostgreSQL user")
var PsqlPassword = flag.String("psqlpassword", "pci-psql-password", "PostgreSQL password")
var PsqlDb = flag.String("psqldb", "postgres", "PostgreSQL database")

func init() {
	// override default values with environment variables or command line flags
	// priority: command line flag > environment variable > default value
	//
	// environment variables are prefixed with PCI_ and are all uppercase
	// example:
	//   $ export PCI_PORT=8081
	flag.VisitAll(func(f *flag.Flag) {
		if env := os.Getenv(
			fmt.Sprintf("PCI_%s", strings.ToUpper(f.Name)),
		); env != "" {
			err := f.Value.Set(env)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err.Error(),
				}).Error("Error setting flag value from environment variable")
			}
		}
	})

	flag.Parse()
}

package parameter

import "flag"

var Host = flag.String("host", "0.0.0.0", "Host to listen on")
var LogLevel = flag.String("loglevel", "default", "Log level")
var Port = flag.Int("port", 8080, "Port to listen on")

func init() {
	flag.Parse()
}

package config

import (
	"flag"
	"os"
	"strconv"
)

const (
	ENV_DEBUG   string = "DEBUG"
	ENV_HOST    string = "LISTEN_ADDR"
	ENV_PORT    string = "LISTEN_PORT"
	ENV_VERSION string = "PROJECT_VERSION"
)

const (
	defaultDebug bool   = false
	defaultHost  string = "localhost"
	defaultPort  int    = 5555
)

var (
	Debug   bool
	Host    string
	Port    int
	Version string
)

func init() {
	flag.BoolVar(&Debug, "debug", defaultDebug, "a boolean, makes the server's output more verbose")
	flag.StringVar(&Host, "host", defaultHost, "a string, a FQDN type of host's name")
	flag.IntVar(&Port, "port", defaultPort, "an integer, a port to listen for incoming requests on")

	flag.Parse()

	//
	//  Try ENV vars.
	//

	var err error

	// Should be non-empty true, otherwise is false.
	if os.Getenv(ENV_DEBUG) != "" {
		Debug, err = strconv.ParseBool(os.Getenv(ENV_DEBUG))
		if err != nil {
			Debug = defaultDebug
		}
	}

	// Should be set to anything non-empty.
	if os.Getenv(ENV_PORT) != "" {
		Port, err = strconv.Atoi(os.Getenv(ENV_PORT))
		if err != nil {
			Port = defaultPort
		}
	}

	if os.Getenv(ENV_VERSION) != "" {
		Version = os.Getenv(ENV_VERSION)
	}
}

package config

type Configuration struct {
	Port int
}

var (
	Debug   bool
	Host    string
	Port    int
	Version string
)

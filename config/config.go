package config

import "flag"

type Config struct {
	Port int
}

func GetConfig() Config {
	port := flag.Int("p", 8000, "The server port")

	flag.Parse()

	return Config{
		Port: *port,
	}
}

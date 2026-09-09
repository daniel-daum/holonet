package config

import (
	"fmt"
	"os"
)

type Settings struct {
	HOST string
	PORT string
	ENV  string
}

func LoadSettings() Settings {
	// HOST
	host, exists := os.LookupEnv("HOST")

	if !exists {
		fmt.Println("Host was not provided, defaulting to 0.0.0.0")
		host = "0.0.0.0"
	}

	// PORT
	port, exists := os.LookupEnv("PORT")

	if !exists {
		fmt.Println("PORT was not provided, defaulting to 7432")
		port = "7432"
	}

	// ENV
	env, exists := os.LookupEnv("ENV")

	if !exists {
		fmt.Println("ENV was not provided, defaulting to development")
		env = "development"
	}

	return Settings{HOST: host, PORT: port, ENV: env}
}

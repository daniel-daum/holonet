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

func getOrDefault(envVar string, defVal string) string {
	variable, exists := os.LookupEnv(envVar)

	if !exists {
		fmt.Printf("%s was not provided, defaulting to %s", variable, defVal)
		variable = defVal
	}
	return variable
}

func LoadSettings() Settings {

	host := getOrDefault("HOST", "0.0.0.0")
	port := getOrDefault("PORT", "7432")
	env := getOrDefault("ENV", "development")

	return Settings{HOST: host, PORT: port, ENV: env}
}

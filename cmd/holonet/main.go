package main

import (
	"git.danieldaum.net/daniel-daum/holonet/internal/config"

	"fmt"
)

func main() {

	settings := config.LoadSettings()

	fmt.Printf("HOST IS SET: %s\n", settings.HOST)
	fmt.Printf("PORT IS SET: %s\n", settings.PORT)
	fmt.Printf("ENV IS SET: %s\n", settings.ENV)
}

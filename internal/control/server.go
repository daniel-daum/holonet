package control

import (
	"net/http"

	"git.danieldaum.net/daniel-daum/holonet/internal/config"
)

// setup control pane http server
// pull and provide configs
func holonet() *http.Server {
	config.GetConfig()

	s := &http.Server{}

	return s
}

// start server
// wrap in graceful shutdown
func Serve() {

	holonet().ListenAndServe()
}

package web

import (
	"fmt"
	"github.com/go-martini/martini"
)

func Start(port int) {
	m := martini.Classic()
	martini.Env = "production"
	m.Post("/server/start", startHandler)
	m.Get("/server/status", statusHandler)
	m.RunOnAddr(fmt.Sprintf(":%d", port))

}

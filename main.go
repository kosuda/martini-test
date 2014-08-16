package main

import (
	"github.com/go-martini/martini"
	"github.com/kosuda/golang-web/conf"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.New()

	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	m.Use(render.Renderer())

	r := conf.Router()
	m.MapTo(r, (*martini.Router)(nil))
	m.Action(r.Handle)

	m.Run()
}

package conf

import (
	"github.com/go-martini/martini"
	"github.com/kosuda/golang-web/controllers"
	"github.com/martini-contrib/binding"
)

// Router setup
func Router() martini.Router {
	router := martini.NewRouter()
	setup(router)
	return router
}

func setup(router martini.Router) {
	router.Get("/user/:id", controllers.UserGet)

	router.Post("/user/:id",
		binding.Json(controllers.User{}),
		binding.ErrorHandler,
		controllers.UserPut)
}

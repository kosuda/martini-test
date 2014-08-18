package controllers

import (
	"github.com/go-martini/martini"
	"github.com/kosuda/golang-web/db"
	"github.com/kosuda/golang-web/models"
	"github.com/martini-contrib/render"
)

// UserGet func
func UserGet(r render.Render, params martini.Params) {

	var user models.User

	// Get all user
	if len(params) == 0 {
	}

	// find by id
	if params["id"] != "" {
		if err := db.Read("user:"+params["id"], &user); err != nil {
			r.JSON(500, map[string]string{"error": err.Error()})
			return
		}
	}

	// error
	if user.Name == "" && user.Tel == "" {
		r.JSON(400, map[string]string{"error": "Argument invalid."})
		return
	}

	r.JSON(200, user)

}

// UserPut func
func UserPut(r render.Render, params martini.Params, user models.User) {

	if len(params) == 0 {
		r.JSON(400, map[string]string{"error": "Argument invalid."})
		return
	}

	if err := db.Write("user:"+params["id"], &user); err != nil {
		r.JSON(500, map[string]string{"error": err.Error()})
		return
	}

	r.JSON(200, user)

}

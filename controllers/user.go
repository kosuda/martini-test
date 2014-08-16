package controllers

import (
	"github.com/go-martini/martini"
	"github.com/kosuda/golang-web/db"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"net/http"
)

// User struct
type User struct {
	Name string `redis:"name" json:"name" binding:"required"`
	Tel  string `redis:"tel" json:"tel" binding:"required"`
}

// Validate fun
func (user User) Validate(errors binding.Errors, req *http.Request) binding.Errors {
	if len(user.Name) < 4 {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"name"},
			Classification: "MinlengthError",
			Message:        "Too short, min length 4 characters.",
		})
	}
	return errors
}

// UserGet func
func UserGet(r render.Render, params martini.Params) {

	var user User

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
func UserPut(r render.Render, params martini.Params, user User) {

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

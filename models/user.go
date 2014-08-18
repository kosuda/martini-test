package models

import (
	"github.com/martini-contrib/binding"
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

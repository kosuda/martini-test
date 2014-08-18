package db

import (
	"github.com/kosuda/golang-web/models"
	"testing"
)

func TestWrite(t *testing.T) {
	data := &models.User{Name: "spam", Tel: "ham"}
	err := Write("test:1", data)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestRead(t *testing.T) {
	var data models.User

	err := Read("test:1", &data)

	if err != nil {
		t.Error(err.Error())
	}

	if data.Name != "spam" || data.Tel != "ham" {
		t.Error("invalid data")
	}
}

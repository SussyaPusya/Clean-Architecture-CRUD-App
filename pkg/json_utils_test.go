package pkg_test

import (
	"App/internal/entity"
	"App/pkg"
	"fmt"
	"testing"
)

func TestParseJson(t *testing.T) {
	bobes := `{"name":"nazik", "pass":"afafs", "auth":"yes"}`
	var bob entity.User

	expected := entity.User{UserName: "nazik", Password: "afafs", Auth: "yes"}

	us := pkg.ParseJson([]byte(bobes), bob)
	fmt.Println(us)

	if us != expected {
		t.Error("КАЛ")
	}

}

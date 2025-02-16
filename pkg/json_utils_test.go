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

func TestValidAuth(t *testing.T) {
	test := []entity.User{
		{UserName: "bob", Auth: "Yes"},
		{UserName: "231", Auth: "No"},
		{UserName: "3", Auth: "yes"},
		{UserName: "12", Auth: "asd"},
	}
	expectedt := []bool{true, false, true, false}

	for i, val := range test {
		if expectedt[i] != pkg.ValidAuth(val) {
			fmt.Println(expectedt[i], pkg.ValidAuth(val))
			t.Error("КААААЛ")
		}

	}
}

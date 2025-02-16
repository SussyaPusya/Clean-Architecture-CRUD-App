package pkg

import (
	"App/internal/entity"
	"encoding/json"
	"log"
	"strings"
)

func ParseJson(bytes []byte, bob entity.User) entity.User {

	err := json.Unmarshal(bytes, &bob)
	if err != nil {
		log.Print("invalid json")
	}
	return bob

}

func ValidAuth(bob entity.User) bool {
	str := strings.ToLower(bob.Auth)

	return str == "yes"
}

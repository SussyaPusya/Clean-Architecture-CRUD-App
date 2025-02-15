package pkg

import (
	"App/internal/entity"
	"encoding/json"
	"log"
)

func ParseJson(bytes []byte, bob entity.User) entity.User {

	err := json.Unmarshal(bytes, &bob)
	if err != nil {
		log.Print("invalid json")
	}
	return bob

}

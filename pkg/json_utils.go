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

func ToJson(user entity.User) ([]byte, error) {
	bytees, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Кал ебаный")
		return nil, err
	}

	return bytees, nil

}

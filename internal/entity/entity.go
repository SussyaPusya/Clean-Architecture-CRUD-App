package entity

type User struct {
	UserName string `json:"name"`
	Password string `json:"pass"`
	Auth     string `json:"auth"`

	BirhDay string `json:"bithday"`
}

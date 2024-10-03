package entity

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type CreateUser struct {
	Name string `json:"name"`
}

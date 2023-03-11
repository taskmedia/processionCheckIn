package model

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

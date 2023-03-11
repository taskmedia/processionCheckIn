package model

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Location struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type User struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

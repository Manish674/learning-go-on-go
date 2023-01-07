package models

type Director struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Age       int      `json:"age"`
	Work      *[]Movie `json:"work"`
}

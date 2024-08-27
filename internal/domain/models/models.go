package models

type User struct {
	UID   string `json:"u_id" validate:"required`
	Name  string `json:"name" validate:"required`
	Login string `json:"login" validate:"required, email`
	Pass  string `json:"pass" validate:"required`
}

type Book struct {
	BID    string `json:"b_id" validate:"required`
	Lable  string `json:"lable" validate:"required`
	Author string `json:"author" validate:"required`
}

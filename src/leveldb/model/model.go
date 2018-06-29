package model

type ListSayhi struct {
	List []Sayhi `json:"sayhi"`
}

type Sayhi struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

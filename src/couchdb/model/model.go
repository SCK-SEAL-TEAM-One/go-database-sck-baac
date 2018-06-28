package model

type Doc struct {
	Id          int    `json:"id"`
	Description string `json:"sayhi"`
}

type Docs struct {
	Detail []Doc `json:"docs"`
}

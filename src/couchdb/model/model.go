package model

type Doc struct {
	IdCoucheDB  string `json:"_id"`
	RevisionID  string `json:"_rev"`
	Id          int    `json:"Id"`
	Description string `json:"Description"`
}

type Docs struct {
	Detail []Doc `json:"docs"`
}

type CouchDbId struct {
	Uuids []string `json:"uuids"`
}
